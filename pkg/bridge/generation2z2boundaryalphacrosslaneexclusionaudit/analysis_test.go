package generation2z2boundaryalphacrosslaneexclusionaudit

import (
	"strings"
	"testing"
)

func TestGate915InheritedSelectorDoesNotReopenClosedWounds(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SelectorShapeSupported || a.Inherited.ReopensPhaseSign || a.Inherited.ReopensSocketOrder || a.Inherited.ReopensRepresentative || a.Inherited.DerivesAlpha || a.Inherited.UpdatesOfficialLedger {
		t.Fatalf("bad inherited selector: %s", FormatInherited(a.Inherited))
	}
}

func TestGate915ExposureEnclosureTypeSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TypeSep.ExcludesByType || a.TypeSep.NativeTheorem || a.TypeSep.LinearFalseLane.TypeCompatible || a.TypeSep.QuadraticFalseLane.TypeCompatible {
		t.Fatalf("bad type separation: %s", FormatTypeSeparation(a.TypeSep))
	}
	if a.TypeSep.LinearFalseLane.FalseContribution != FalseLinearTerm || a.TypeSep.QuadraticFalseLane.FalseContribution != FalseQuadraticTerm {
		t.Fatalf("bad false terms: %s", FormatTypeSeparation(a.TypeSep))
	}
}

func TestGate915SelectorDeterminismConditionalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Determinism.IsFunction || !a.Determinism.ExcludesFalseTargets || a.Determinism.UniqueNativeSelector {
		t.Fatalf("determinism firewall leaked: %s", FormatDeterminism(a.Determinism))
	}
}

func TestGate915RankContamination(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Contamination.CorrectAlpha, AlphaB) || !a.Contamination.MismatchDetected || a.Contamination.PollutedAlpha <= a.Contamination.CorrectAlpha || a.Contamination.NativeExclusion {
		t.Fatalf("bad contamination check: %s", FormatContamination(a.Contamination))
	}
	if !near(a.Contamination.PollutedCoefficient, float64(143)/float64(360)) {
		t.Fatalf("bad polluted coefficient: %s", FormatContamination(a.Contamination))
	}
}

func TestGate915CumulativeAndZ2Compatibility(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Cumulative.F2OverF0Rank != RankF2OverF0 || a.Cumulative.F2OverF1Rank != RankF2OverF1 || !a.Cumulative.KeepsCumulativeEnclosure || !a.Cumulative.RejectsAssociatedGradedSlice || a.Cumulative.NativeReasonForChoice {
		t.Fatalf("bad cumulative consistency: %s", FormatCumulative(a.Cumulative))
	}
	if !a.Z2.CorrectLanesRepresentativeFree || !a.Z2.FalseLanesRepresentativeFree || !a.Z2.CorrectMapToCorrect || !a.Z2.FalseMapToFalse || a.Z2.NativeExclusionTheorem {
		t.Fatalf("bad z2 compatibility: %s", FormatZ2(a.Z2))
	}
}

func TestGate915GlobalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeZ2CrossLaneTheorem, FailureTypeSeparationNotNativeFunctorTheorem, FailureNoNativeUniqueDegreeSelector, FailureNumericalMismatchNotNativeExclusion, FailureNoNativeCumulativeReason, FailureZ2CompatibilityNotNativeExclusion, FailureAlphaStillSealed, FailureDenominatorsSTransportStillExternal, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s from %v", want, a.FirewallsList())
		}
	}
}

func TestGate915Theorem(t *testing.T) {
	res := Generation2Z2BoundaryAlphaCrossLaneExclusionAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s in notes: %s", want, joined)
		}
	}
}
