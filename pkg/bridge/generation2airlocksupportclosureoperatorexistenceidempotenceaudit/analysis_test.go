package generation2airlocksupportclosureoperatorexistenceidempotenceaudit

import (
	"strings"
	"testing"
)

func TestGate930ClosureOperatorLaws(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != Gate929ShortStatus {
		t.Fatalf("bad inherited: %s", a.Inherited)
	}
	if !a.SupportChain.FiniteChain || !a.SupportChain.LeastSupportsExist || a.SupportChain.NativeClosureTheorem {
		t.Fatalf("bad support chain: %s", FormatSupportChain(a.SupportChain))
	}
	if !a.Extensivity.DemandTyped || !a.Extensivity.Cl0ContainsBasepoint || !a.Extensivity.Cl1ContainsExposure || !a.Extensivity.Cl2ContainsEnclosure || a.Extensivity.NativeSubspaceClosure {
		t.Fatalf("bad extensivity: %s", FormatExtensivity(a.Extensivity))
	}
	if !a.Monotonicity.Monotone || !a.Monotonicity.SupportInclusion || a.Monotonicity.NativeActionTheorem {
		t.Fatalf("bad monotonicity: %s", FormatMonotonicity(a.Monotonicity))
	}
	if !a.Idempotence.ImageInAdmissibleFamily || !a.Idempotence.F0Closed || !a.Idempotence.F1Closed || !a.Idempotence.F2Closed || !a.Idempotence.Idempotent || a.Idempotence.NativeClosureTheorem {
		t.Fatalf("bad idempotence: %s", FormatIdempotence(a.Idempotence))
	}
}

func TestGate930MinimalSaturatedZ2RecoveryAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Minimality.DemandK != 1 || a.Minimality.ClosureTarget != "F_1" || !a.Minimality.LeastNonbaseSupport || !a.Minimality.RejectsJumpToF2 || a.Minimality.NativeRule {
		t.Fatalf("bad minimality: %s", FormatMinimality(a.Minimality))
	}
	if a.Saturation.DemandK != 2 || a.Saturation.ClosureTarget != "F_2" || !a.Saturation.FullPairSaturation || !a.Saturation.RejectsCloseToF1 || a.Saturation.NativeRule {
		t.Fatalf("bad saturation: %s", FormatSaturation(a.Saturation))
	}
	if !a.Z2.PhaseFlipCommutes || !a.Z2.DescendsToZ2Class || a.Z2.NativePhaseTheorem {
		t.Fatalf("bad z2: %s", FormatZ2(a.Z2))
	}
	if !a.TargetRecovery.CumulativeTargets || !a.TargetRecovery.RejectsAssociatedGraded || a.TargetRecovery.NativeFixedBaseTheorem || a.TargetRecovery.ThetaOneRank != RankF1OverF0 || a.TargetRecovery.ThetaTwoRank != RankF2OverF0 {
		t.Fatalf("bad target recovery: %s", FormatTargetRecovery(a.TargetRecovery))
	}
	if !a.Measure.RewrittenUsingClosure || !a.Measure.AlphaReconstructed || a.Measure.NativeAlpha {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate930Theorem(t *testing.T) {
	res := Generation2AirlockSupportClosureOperatorExistenceIdempotenceAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, AdmissibleSupportFamily, Z2AdmissibleSupportFamily, BoundaryDemandChain, ClosureOperatorName, ClosureDefinition, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, MuBViaClosure, AlphaViaClosureOperator, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
