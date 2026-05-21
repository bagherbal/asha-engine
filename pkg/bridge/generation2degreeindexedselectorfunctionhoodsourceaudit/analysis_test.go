package generation2degreeindexedselectorfunctionhoodsourceaudit

import (
	"strings"
	"testing"
)

func TestGate923LaneSourceTyping(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedStatus != Gate922ShortStatus {
		t.Fatalf("bad inherited status: %s", a.InheritedStatus)
	}
	if a.DegreeOne.Degree != 1 || a.DegreeOne.Rank != RankExposureFace || !a.DegreeOne.BridgeTyped || a.DegreeOne.Native {
		t.Fatalf("bad degree one lane: %s", FormatLane(a.DegreeOne))
	}
	if a.DegreeTwo.Degree != 2 || a.DegreeTwo.Rank != RankFullEnclosure || !a.DegreeTwo.BridgeTyped || a.DegreeTwo.Native {
		t.Fatalf("bad degree two lane: %s", FormatLane(a.DegreeTwo))
	}
}

func TestGate923CumulativeFunctionhoodCrossLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CumulativeEnclosure.CumulativeRequired || a.CumulativeEnclosure.Native || a.CumulativeEnclosure.RankF2OverF0 != RankFullEnclosure || a.CumulativeEnclosure.RankF2OverF1 != RankAssociatedSlice {
		t.Fatalf("bad cumulative audit: %s", FormatCumulative(a.CumulativeEnclosure))
	}
	if !a.Functionhood.ExposureEnclosureAccepted || !a.Functionhood.FunctionalIfExposureAccepted || a.Functionhood.NativeFunctionhood || a.Functionhood.PrimaryGap != PrimaryGapExposureFunctor {
		t.Fatalf("bad functionhood: %s", FormatFunctionhood(a.Functionhood))
	}
	if !a.CrossLane.ExcludedIfFunctional || a.CrossLane.NativeExclusion {
		t.Fatalf("bad cross lane: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate923Z2MuBAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Z2Compatibility.CommutesWithPhaseFlip || !a.Z2Compatibility.RanksRepresentativeFree || a.Z2Compatibility.NativeZ2Selector {
		t.Fatalf("bad z2 audit: %s", FormatZ2(a.Z2Compatibility))
	}
	if !near(a.MuB.LinearContribution, AlphaLinear) || !near(a.MuB.QuadraticContribution, AlphaQuad) || !near(a.MuB.Alpha, AlphaB) || a.MuB.NativeMuB {
		t.Fatalf("bad mu_B: %s", FormatMuB(a.MuB))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureExposureToF1NotNative, FailureEnclosureToF2NotNative, FailureNoNativeCumulativeTheorem, FailureSelectorDependsOnBridgeRule, FailureCrossLaneNotNativeWithoutSelector, FailureZ2CompatibilityNotNativeSelector, FailureMuBStillNotNativeWithoutSelector, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate923Theorem(t *testing.T) {
	res := Generation2DegreeIndexedSelectorFunctionhoodSourceAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate, PunctureClass, BoundaryResponse, SelectorFormula, SelectorOneFormula, SelectorTwoFormula, MuBFormula, BoundaryAlphaFormula, ExposureEnclosureRule, PrimaryGapExposureFunctor} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
