package generation2higgsorientedstabilizeralgebrapostorientationlayeraudit

import (
	"strings"
	"testing"
)

func TestGate856WeakFrameStabilizer(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.WeakFrame.FullHActsOnFullDoublet || a.WeakFrame.FullHPreservesIndividualLines || !a.WeakFrame.StabilizerPreservesHPlus || !a.WeakFrame.StabilizerPreservesHMinus || !a.WeakFrame.StabilizerIsComplexSubalgebra || a.WeakFrame.StabilizerIsNativeFullH {
		t.Fatalf("weak frame overpromoted: %s", FormatWeakFrame(a.WeakFrame))
	}
	if !containsAll(a.WeakFrame.Failures, []string{FailureFullHPreservesSocketFrame, FailureFullHNativeSocketEigensplit}) {
		t.Fatalf("missing full-H firewall failures: %+v", a.WeakFrame.Failures)
	}
}

func TestGate856OrientedAlgebraLayer(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Algebra.ContainsFullH || !a.Algebra.ContainsCH || !a.Algebra.ContainsM3C || !a.Algebra.ContainsRightC || !a.Algebra.PostOrientationLayer || a.Algebra.UnbrokenFullAFTheorem || a.Algebra.PhysicalElectroweakTheorem {
		t.Fatalf("bad oriented algebra flags: %s", FormatAlgebra(a.Algebra))
	}
	if a.Algebra.OrientedAlgebra != "A_F^orient=C_R plus C_H plus M_3(C)" {
		t.Fatalf("unexpected oriented algebra: %s", a.Algebra.OrientedAlgebra)
	}
}

func TestGate856ActionAndDFSupport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Action.MinimalCarrierClosureInAForient || a.Action.MinimalCarrierClosureInFullAF || !a.Action.PunctureRemainsOutside || !a.Action.LeftKernelStableCandidate {
		t.Fatalf("action preservation invalid: %s", FormatAction(a.Action))
	}
	if !a.DF.SupportCompatible || a.DF.OperatorTheoremCompatible || a.DF.FullAFCompatible || !a.DF.PostOrientationObject || !a.DF.FirstOrderReadyForGate857 || a.DF.FirstOrderCalculatedThisGate || a.DF.FirstOrderCertified {
		t.Fatalf("D_F support overpromoted: %s", FormatDF(a.DF))
	}
}

func TestGate856CarrierLedgerAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Carrier.HPartMinRank != 15 || a.Carrier.HFMinRank != 30 || a.Carrier.AmbientPartRank != 16 || a.Carrier.AmbientFRank != 32 || a.Carrier.DSymRank != 14 || a.Carrier.KernelRank != 1 {
		t.Fatalf("bad carrier: %s", FormatCarrier(a.Carrier))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate856Theorem(t *testing.T) {
	res := Generation2HiggsOrientedStabilizerAlgebraPostOrientationLayerAuditTheorem().Verify()
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
