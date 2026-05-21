package generation2boundarydegreetoairlockflagtargetfunctornaturalityanduniquenessaudit

import (
	"strings"
	"testing"
)

func TestGate926OrderTypeAndThetaUniqueness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedStatus != Gate925ShortStatus {
		t.Fatalf("bad inherited status: %s", a.InheritedStatus)
	}
	if !a.Order.UniqueUnderOrder || !a.Order.OrderPreserving || !a.Order.SwappedOrderReversing || a.Order.NativeOrderTheorem {
		t.Fatalf("bad order audit: %s", FormatOrder(a.Order))
	}
	if !a.Types.ExposureUnique || !a.Types.EnclosureUnique || !a.Types.ForcesThetaAssignment || a.Types.NativeTypeTheorem {
		t.Fatalf("bad type audit: %s", FormatTypes(a.Types))
	}
}

func TestGate926Z2AndAlternativeRejections(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Z2.RepresentativeIndependent || !a.Z2.CommutesWithFlip || a.Z2.NativeZ2Theorem || a.Z2.RankOne != RankF1OverF0 || a.Z2.RankTwo != RankF2OverF0 {
		t.Fatalf("bad z2 audit: %s", FormatZ2(a.Z2))
	}
	if !a.Graded.RejectedByTypeAndRank || !a.Graded.FailsCumulativeType || !a.Graded.FailsAlphaRank || a.Graded.NativeCumulativeTheorem || a.Graded.AlternativeRank != RankF2OverF1 {
		t.Fatalf("bad graded audit: %s", FormatGraded(a.Graded))
	}
	if !a.Alternatives.DegreeZeroAbsent || !a.Alternatives.ViolatesType || !a.Alternatives.FailsAlphaShape || a.Alternatives.NativeUniquenessTheorem {
		t.Fatalf("bad alternatives: %s", FormatAlternatives(a.Alternatives))
	}
}

func TestGate926MeasureAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Measure.SelectorFunctionhood || !a.Measure.CrossLaneExclusion || !a.Measure.TargetRanksFixed || a.Measure.NativeMeasure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureThetaNotNative, FailureOrderPreservationNotNative, FailureExposureEnclosureUniquenessNative, FailureZ2IndependenceNotNative, FailureF2OverF1RejectionNotNative, FailureAlternativeRejectionNotNative, FailureMuBNotNativeWithoutThetaSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate926Theorem(t *testing.T) {
	res := Generation2BoundaryDegreeToAirlockFlagTargetFunctorNaturalityAndUniquenessAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, SourceChain, TargetChain, ThetaFunctor, ThetaOne, ThetaTwo, SwappedAssignment, AssociatedGradedTarget, ExposureType, EnclosureType, MeasureFormula, AlphaFormula, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
