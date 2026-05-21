package generation2symbolicdffirstorderjoppositecompatibilityaudit

import (
	"strings"
	"testing"
)

func TestGate850InheritsSymbolicDF(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SymbolicDF.SupportOnly || a.SymbolicDF.NativeFiniteTriple || a.SymbolicDF.ExplicitDFOperator || a.SymbolicDF.NumericalDFMatrix {
		t.Fatalf("symbolic D_F overpromoted: %s", FormatSymbolicDF(a.SymbolicDF))
	}
	if a.SymbolicDF.LeftRank != HLRank || a.SymbolicDF.RightRank != HRMinRank || a.SymbolicDF.TotalRank != ChiralTotalDim || a.SymbolicDF.YRank != 7 || a.SymbolicDF.DFRank != 14 || a.SymbolicDF.KernelDim != 1 {
		t.Fatalf("bad D_F rank anatomy: %s", FormatSymbolicDF(a.SymbolicDF))
	}
	if !a.SymbolicDF.SelfAdjointByBlock || !a.SymbolicDF.ChiralityOddByBlock {
		t.Fatalf("missing block-form support properties: %s", FormatSymbolicDF(a.SymbolicDF))
	}
}

func TestGate850BlocksFirstOrderWithoutData(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Representation.RhoFAvailable || a.Representation.CompleteActionLedger || a.Representation.CompletePackage || a.Representation.JFAvailable || a.Representation.DFOperatorAvailable {
		t.Fatalf("representation data incorrectly certified: %s", FormatRepresentation(a.Representation))
	}
	if a.FirstOrder.CanFormCommutator || a.FirstOrder.CanFormOppositeAction || a.FirstOrder.FirstOrderProven || a.FirstOrder.BimoduleStable || a.FirstOrder.JOppositeCompatible {
		t.Fatalf("first-order proof incorrectly certified: %s", FormatFirstOrder(a.FirstOrder))
	}
	if !containsAll(a.FirstOrder.Failures, []string{FailureNoCompleteRhoFActionLedger, FailureNoJOppositeCompatibilityProof, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof}) {
		t.Fatalf("missing first-order failures: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate850KernelStabilityCandidateFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Representation.Kernel.Name != "h_+ tensor P_1" || !a.Representation.Kernel.Kernel || !a.Representation.Kernel.StableCandidate || a.Representation.Kernel.PhysicalAssignment {
		t.Fatalf("bad kernel candidate: %s", FormatCell(a.Representation.Kernel))
	}
	if a.Representation.KernelStabilityCertified || a.FirstOrder.KernelStable {
		t.Fatalf("kernel stability overpromoted: %s | %s", FormatRepresentation(a.Representation), FormatFirstOrder(a.FirstOrder))
	}
	if a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing {
		t.Fatalf("impact overpromoted: %s", FormatImpact(a.Impact))
	}
}

func TestGate850Theorem(t *testing.T) {
	res := Generation2SymbolicDFFirstOrderJOppositeCompatibilityAuditTheorem().Verify()
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
