package generation2airlockclosureaxiomsourceflaggeneratedminimalityaudit

import (
	"strings"
	"testing"
)

func TestGate929FlagGeneratedClosureSources(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != Gate928ShortStatus {
		t.Fatalf("bad inherited: %s", a.Inherited)
	}
	if !a.Basepoint.PunctureInitiality || !a.Basepoint.EmptyBoundarySubset || !a.Basepoint.FlagGenerated || a.Basepoint.NativeClosureTheorem || a.Basepoint.ClosureTarget != "F_0" {
		t.Fatalf("bad basepoint: %s", FormatBasepoint(a.Basepoint))
	}
	if !a.Monotonicity.SupportInclusion || !a.Monotonicity.OrdersCompatible || !a.Monotonicity.FlagNatural || a.Monotonicity.NativeActionTheorem {
		t.Fatalf("bad monotonicity: %s", FormatMonotonicity(a.Monotonicity))
	}
	if !a.Minimality.SingletonActivation || !a.Minimality.ExposedFace || !a.Minimality.ForcedByMinimalSupport || a.Minimality.NativeClosureTheorem || a.Minimality.ClosureTarget != "F_1" {
		t.Fatalf("bad minimality: %s", FormatMinimality(a.Minimality))
	}
	if !a.Saturation.FullPairActivation || !a.Saturation.TopExteriorDegree || !a.Saturation.ForcedBySaturation || a.Saturation.NativeClosureTheorem || a.Saturation.ClosureTarget != "F_2" {
		t.Fatalf("bad saturation: %s", FormatSaturation(a.Saturation))
	}
}

func TestGate929FixedBaseZ2LedgerMeasureAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FixedBase.UsesFixedBaseF0 || !a.FixedBase.CumulativeQuotient || !a.FixedBase.RejectsAssociatedGraded || a.FixedBase.NativeMeasureTheorem {
		t.Fatalf("bad fixed-base: %s", FormatFixedBase(a.FixedBase))
	}
	if !a.Z2.PhaseFlipExchanges || !a.Z2.ClassLevelClosure || !a.Z2.RanksInvariant || a.Z2.NativeGlobalPhaseTheorem {
		t.Fatalf("bad z2: %s", FormatZ2(a.Z2))
	}
	if !a.Ledger.FlagGenerated || a.Ledger.NativeOperatorExists {
		t.Fatalf("bad ledger: %s", FormatLedger(a.Ledger))
	}
	if !a.Measure.ThetaReconstructed || !a.Measure.TargetsFixed || a.Measure.NativeAlpha || a.Measure.ThetaOneRank != RankF1OverF0 || a.Measure.ThetaTwoRank != RankF2OverF0 {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate929Theorem(t *testing.T) {
	res := Generation2AirlockClosureAxiomSourceFlagGeneratedMinimalityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, BoundarySubsetChain, AirlockFlagChain, Z2PunctureClass, CandidateOperator, LeastSupportRule, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, AlphaViaClosure, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
