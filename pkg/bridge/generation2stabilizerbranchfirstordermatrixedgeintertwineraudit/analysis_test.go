package generation2stabilizerbranchfirstordermatrixedgeintertwineraudit

import (
	"strings"
	"testing"
)

func TestGate857StabilizerAlgebraTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Algebra.ContainsFullH || !a.Algebra.ContainsCH || !a.Algebra.ContainsRightC || !a.Algebra.ContainsM3C || !a.Algebra.PostOrientationLayer {
		t.Fatalf("bad oriented algebra: %s", FormatAlgebra(a.Algebra))
	}
	if a.Algebra.OrientedAlgebra != "A_F^orient=C_R plus C_H plus M_3(C)" {
		t.Fatalf("unexpected oriented algebra: %s", a.Algebra.OrientedAlgebra)
	}
	if !containsAll(a.Algebra.Failures, []string{FailureFullAFNotTestedAsTarget, FailureAForientNotFullAF, FailureFullHSocketFirewall}) {
		t.Fatalf("missing algebra firewall: %+v", a.Algebra.Failures)
	}
}

func TestGate857EdgesAreSupportIntertwinersOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Edges) != 3 {
		t.Fatalf("expected three edges: %s", FormatEdges(a.Edges))
	}
	for _, e := range a.Edges {
		if !e.BlockwiseCompatible || e.OperatorIntertwiner || e.CharacterMatchCertified {
			t.Fatalf("edge overpromoted: %+v", e)
		}
		if !containsAll(e.Failures, []string{FailureCharacterMatchSupportOnly, FailureYSupportNotMagnitude}) {
			t.Fatalf("missing edge failures: %+v", e.Failures)
		}
	}
}

func TestGate857FirstOrderSupportButNotOperatorTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.DCommutatorExpectedNonzero || !a.FirstOrder.NonzeroCommutatorAllowed || !a.FirstOrder.OppositeSupportAuditable || !a.FirstOrder.SupportFirstOrderCompatible {
		t.Fatalf("support first-order not enabled: %s", FormatFirstOrder(a.FirstOrder))
	}
	if a.FirstOrder.OppositeOperatorCertified || a.FirstOrder.OperatorFirstOrderCertified || a.FirstOrder.BimoduleCertified {
		t.Fatalf("first-order overpromoted: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate857KernelPunctureAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.NeutralPair.RightPuncture != "e_+ tensor P_1" || a.NeutralPair.LeftKernel != "h_+ tensor P_1" || !a.NeutralPair.RightPuncturePreserved || !a.NeutralPair.LeftKernelPreserved || a.NeutralPair.LeftKernelOperatorStable || a.NeutralPair.PhysicalParticleTheorem {
		t.Fatalf("neutral pair overpromoted: %s", FormatNeutralPair(a.NeutralPair))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger/promote firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate857Theorem(t *testing.T) {
	res := Generation2StabilizerBranchFirstOrderMatrixEdgeIntertwinerAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
