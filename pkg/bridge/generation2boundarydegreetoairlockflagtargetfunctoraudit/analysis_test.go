package generation2boundarydegreetoairlockflagtargetfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate925ChainAndTargets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedStatus != Gate924ShortStatus {
		t.Fatalf("bad inherited status: %s", a.InheritedStatus)
	}
	if !a.ChainMatch.MatchingOrderType || a.ChainMatch.NativeTargetFunctor || a.ChainMatch.SourceLevels != 2 || a.ChainMatch.TargetLevels != 2 {
		t.Fatalf("bad chain match: %s", FormatChain(a.ChainMatch))
	}
	if a.ExposureTarget.Degree != 1 || a.ExposureTarget.Target != "[F_1/F_0]_{Z2}" || a.ExposureTarget.Rank != RankF1OverF0 || a.ExposureTarget.NativeRule {
		t.Fatalf("bad exposure target: %s", FormatTarget(a.ExposureTarget))
	}
	if a.EnclosureTarget.Degree != 2 || a.EnclosureTarget.Target != "[F_2/F_0]_{Z2}" || a.EnclosureTarget.Rank != RankF2OverF0 || a.EnclosureTarget.NativeRule {
		t.Fatalf("bad enclosure target: %s", FormatTarget(a.EnclosureTarget))
	}
}

func TestGate925ThetaSelectorAndMeasure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Graded.TopDegreeCumulative || a.Graded.NativeCumulativeRule || a.Graded.RejectedRank != RankF2OverF1 || a.Graded.CumulativeRank != RankF2OverF0 {
		t.Fatalf("bad graded audit: %s", FormatGraded(a.Graded))
	}
	if !a.Theta.OrderPreserving || !a.Theta.Z2RepresentativeIndependent || !a.Theta.ExposureEnclosureTyped || !a.Theta.CumulativeTopDegree || a.Theta.NativeFunctor {
		t.Fatalf("bad theta: %s", FormatTheta(a.Theta))
	}
	if !a.Selector.IBEqualsTheta || !a.Selector.CrossLaneExcluded || a.Selector.NativeSelectorFunctionhood {
		t.Fatalf("bad selector: %s", FormatSelector(a.Selector))
	}
	if !a.Measure.TargetRanksSupplied || a.Measure.NativeMeasure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 || a.Measure.H10Rank != RankH10 || a.Measure.H72Rank != RankH72 {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
}

func TestGate925Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureOrderTypeNotNative, FailureMinimalExposureLevelNotNative, FailureFullEnclosureLevelNotNative, FailureCumulativeOverGradedNotNative, FailureThetaShapeNotNative, FailureSelectorNonNativeWithoutTheta, FailureMuBNotNativeWithoutTheta, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate925Theorem(t *testing.T) {
	res := Generation2BoundaryDegreeToAirlockFlagTargetFunctorAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, BoundaryDegreeChain, AirlockFlagChain, Z2PunctureClass, ThetaFunctor, ThetaOne, ThetaTwo, ExposureTarget, EnclosureTarget, MeasureFormula, AlphaFormula, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
