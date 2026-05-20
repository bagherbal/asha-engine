package generation2k7eventweightandbernoulliresponseobservableaudit

import (
	"math"
	"strings"
	"testing"
)

func TestInheritanceAndEventProjector(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.MaximumEntropyObserverInherited || a.Inherited.Rho72Definition != "rho_72 = I_H72/72" || !a.Inherited.Rho72MaximumEntropy || !a.Inherited.Rho72NoBias || a.Inherited.ResponseOperator != "R_split = S_split P_K7" {
		t.Fatalf("bad Gate694 inheritance: %+v", a.Inherited)
	}
	if !a.Inherited.NoNativeMaximumEntropyHistory || !a.Inherited.NoNativeStateSelection || !a.Inherited.NoNativeSevenOver72 {
		t.Fatalf("Gate694 firewall not inherited: %+v", a.Inherited)
	}
	if a.Event.EventProjector != "E_K7 = P_K7" || !a.Event.IsProjector || !a.Event.Rho72TraceOne || math.Abs(a.Event.EventWeight-7.0/72.0) > tolerance || math.Abs(a.Event.ComplementWeight-65.0/72.0) > tolerance || !a.Event.WeightsNormalize || a.Event.EventDimension != k7Dimension || a.Event.ComplementDimension != k7ComplementDim {
		t.Fatalf("bad K7 event audit: %+v", a.Event)
	}
}

func TestBernoulliObservableExpectationAndMoments(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Observable.Observable != "R_split = S_split P_K7" || math.Abs(a.Observable.EventPayoff-a.Inherited.SSplit) > tolerance || a.Observable.ComplementPayoff != 0 || math.Abs(a.Observable.EventProbability-7.0/72.0) > tolerance || math.Abs(a.Observable.ComplementProbability-65.0/72.0) > tolerance || !a.Observable.TwoPointObservable {
		t.Fatalf("bad Bernoulli observable audit: %+v", a.Observable)
	}
	if math.Abs(a.Expectation.Expectation-a.Inherited.ActiveExpectation) > tolerance || math.Abs(a.Expectation.ResidualE1-8.525834398014336e-10) > tolerance || !a.Expectation.ReproducesBridge {
		t.Fatalf("bad expectation audit: %+v", a.Expectation)
	}
	wantSecond := (7.0 / 72.0) * a.Inherited.SSplit * a.Inherited.SSplit
	wantVariance := (7.0 / 72.0) * (65.0 / 72.0) * a.Inherited.SSplit * a.Inherited.SSplit
	if math.Abs(a.Moment.SecondMoment-wantSecond) > tolerance || math.Abs(a.Moment.Gate690F2-a.Moment.SecondMoment) > tolerance || !a.Moment.MatchesGate690F2 || math.Abs(a.Moment.Variance-wantVariance) > tolerance || a.Moment.VarianceIsBridge {
		t.Fatalf("bad moment audit: %+v", a.Moment)
	}
}

func TestAlternativeStateEventWeightsAndInterpretation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Alternatives.States) != 5 || !a.Alternatives.Rho72Active || !a.Alternatives.FiniteRejected || !a.Alternatives.KernelRejected || !a.Alternatives.LocalK7Rejected || !a.Alternatives.BoundaryRejected || !a.Alternatives.AllAlternativesAudited {
		t.Fatalf("bad alternative state weights: %+v", a.Alternatives)
	}
	weights := map[string]float64{}
	for _, s := range a.Alternatives.States {
		weights[s.Name] = s.K7Weight
	}
	if math.Abs(weights["rho_72"]-7.0/72.0) > tolerance || math.Abs(weights["rho_finite"]-7.0/70.0) > tolerance || math.Abs(weights["rho_kernel"]-7.0/71.0) > tolerance || math.Abs(weights["rho_K7"]-1.0) > tolerance || math.Abs(weights["rho_boundary"]) > tolerance {
		t.Fatalf("unexpected alternative weights: %+v", weights)
	}
	if !a.Interpretation.NoBiasExpectedPayoff || !strings.Contains(a.Interpretation.PayoffInterpretation, "S_split") || !strings.Contains(a.Interpretation.ProbabilityInterpretation, "7/72") || !strings.Contains(a.Interpretation.ScalarDefectReading, "D_base") {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
}

func TestMissingTheoremsAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Missing.Missing) != 4 || !strings.Contains(a.Missing.PreciseGap, "why physical history uses rho_72") || !strings.Contains(a.Missing.Verdict, StatusNoNativeReasonK7EventReceivesSSplitPayoff) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Firewall.ClaimsHistoryUsesRho72Natively || a.Firewall.ClaimsK7PhysicalEventNatively || a.Firewall.ClaimsSSplitPayoffNatively || a.Firewall.ClaimsExpectationEqualsDBaseNatively || a.Firewall.ClaimsResidualExplained || a.Firewall.ClaimsNativeSevenOver72Theorem || a.Firewall.ClaimsBoundaryStress || a.Firewall.ClaimsScalarRGMatching || a.Firewall.ClaimsHiggsMass || a.Firewall.ClaimsGaugeUnification || a.Firewall.ClaimsFlavorDerivation || a.Firewall.ClaimsCKMPMNS {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7EventWeightAndBernoulliResponseObservableAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
