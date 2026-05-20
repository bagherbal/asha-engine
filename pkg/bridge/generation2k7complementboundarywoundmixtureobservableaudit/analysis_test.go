package generation2k7complementboundarywoundmixtureobservableaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate703InheritanceAndRearrangement(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	inh := a.Inherited
	if !inh.ScalarWallAirlockInherited || !inh.UnitScalarWallGlue || !nearly(inh.EventProbability, pK7) || !nearly(inh.ResponseCoefficient, pK7) || !inh.NonTautologyInherited || !inh.NoNativeScalarWallAirlock || !inh.NoNativeBoundaryHistory || !inh.NoNativeSevenOver72 || inh.Verdict != StatusGate703ScalarWallAirlockInherited {
		t.Fatalf("bad inheritance: %+v", inh)
	}
	r := a.Rearrangement
	if !strings.Contains(r.StartingEquation, "p_K7") || !strings.Contains(r.RearrangedEquation, "p_perp") || !r.LambdaNegative || !nearly(r.KSum, kappaLambda+kappaE) || !nearly(r.PositiveScalarDepth, math.Abs(lambdaLambda12)) || !nearly(r.GaugeWound, r3Minus1) || !nearly(r.WeightedClosureRight, pComplement*math.Abs(lambdaLambda12)+pK7*r3Minus1) || !r.SameAsGate700Residual || !strings.Contains(r.Verdict, StatusGate700ResponseLawRearranged) {
		t.Fatalf("bad rearrangement: %+v", r)
	}
}

func TestProbabilitiesObservableAndExpectation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Probabilities
	if p.K7Rank != 7 || p.ComplementRank != 65 || p.TotalDimension != 72 || !nearly(p.PK7Probability, 7.0/72.0) || !nearly(p.ComplementProb, 65.0/72.0) || !p.ProbabilitiesSumTo1 || !strings.Contains(p.Verdict, Status65Over72ComplementEventProbability) {
		t.Fatalf("bad probabilities: %+v", p)
	}
	o := a.Observable
	if !strings.Contains(o.Observable, "W_boundary") || !nearly(o.K7Payoff, r3Minus1) || !nearly(o.ComplementPayoff, math.Abs(lambdaLambda12)) || !strings.Contains(o.K7PayoffRole, "gauge") || !strings.Contains(o.ComplementPayoffRole, "scalar") || !o.IsTwoPayoffObservable || o.Verdict != StatusTwoPayoffBoundaryWoundObservableDefined {
		t.Fatalf("bad observable: %+v", o)
	}
	e := a.Expectation
	if !strings.Contains(e.Formula, "p_K7") || !strings.Contains(e.Formula, "p_perp") || !nearly(e.ExpectedBoundaryWound, pK7*r3Minus1+pComplement*math.Abs(lambdaLambda12)) || !nearly(e.KSum, kappaLambda+kappaE) || math.Abs(e.Residual-8.5258344e-10) > 1e-16 || !e.ReproducesWeightedClosure || !strings.Contains(e.Verdict, StatusKappaSumNoBiasExpectedBoundaryWound) {
		t.Fatalf("bad expectation: %+v", e)
	}
}

func TestNumericalInterpretationAndEquivalence(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	n := a.Numerical
	if math.Abs(n.KSum-0.0498265972876517) > 5e-16 || math.Abs(n.ExpectedWound-0.0498265964350682) > 5e-16 || math.Abs(n.Residual-8.5258344e-10) > 1e-16 || !n.SameResidual || n.Verdict != StatusNumericalResidualRecorded {
		t.Fatalf("bad numerical audit: %+v", n)
	}
	i := a.Interpretation
	if !strings.Contains(i.K7EventPayoff, "gauge") || !strings.Contains(i.ComplementPayoff, "scalar") || !strings.Contains(i.Observer, "rho_72") || !strings.Contains(i.Output, "K_sum") || !strings.Contains(i.Reading, "expected boundary wound") {
		t.Fatalf("bad interpretation: %+v", i)
	}
	eq := a.Equivalence
	if len(eq.Forms) != 4 || eq.IntroducesNewNumericalRelation || !eq.UpgradesSourceType || eq.Verdict != StatusEquivalenceToPreviousFormsAudited {
		t.Fatalf("bad equivalence audit: %+v", eq)
	}
	for _, form := range eq.Forms {
		if !form.Equivalent {
			t.Fatalf("non-equivalent form: %+v", form)
		}
	}
}

func TestAlternativeMixturesMissingAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	alts := a.Alternatives
	if len(alts.Alternatives) != 5 || !alts.ReversedPayoffRejected || !alts.SupportLocalFormSeparated || !alts.MidpointRejected || !alts.HodgeSignedRejected || !alts.ActiveMixtureAccepted || alts.Verdict != StatusAlternativeMixtureObservablesAudited {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	if !alts.Alternatives[0].Rejected || !strings.Contains(alts.Alternatives[0].Reason, "wrong active orientation") {
		t.Fatalf("bad reversed alternative: %+v", alts.Alternatives[0])
	}
	if !alts.Alternatives[1].Rejected || !strings.Contains(alts.Alternatives[1].Reason, "D_base") {
		t.Fatalf("bad support-local alternative: %+v", alts.Alternatives[1])
	}
	if !alts.Alternatives[2].Rejected || !strings.Contains(alts.Alternatives[2].Reason, "loses event/complement") {
		t.Fatalf("bad midpoint alternative: %+v", alts.Alternatives[2])
	}
	if !alts.Alternatives[3].Rejected || !strings.Contains(alts.Alternatives[3].Reason, "signed polarity") {
		t.Fatalf("bad Hodge-signed alternative: %+v", alts.Alternatives[3])
	}
	if !alts.Alternatives[4].Active || alts.Alternatives[4].Rejected || !nearly(alts.Alternatives[4].Expectation, a.Expectation.ExpectedBoundaryWound) {
		t.Fatalf("bad active mixture: %+v", alts.Alternatives[4])
	}

	m := a.Missing
	if len(m.Missing) != 5 || !strings.Contains(m.Verdict, StatusNoNativeK7ReceivesGaugeWound) || !strings.Contains(m.Verdict, StatusNoNativeComplementReceivesScalarWound) || !strings.Contains(m.Verdict, StatusNoNativeBoundaryWoundMixtureTheorem) || !strings.Contains(m.Verdict, StatusNoNativeHistoryResponseTheorem) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsK7ReceivesGaugeWoundNative || f.ClaimsComplementReceivesScalarWoundNative || f.ClaimsNativeBoundaryWoundMixtureTheorem || f.ClaimsNativeHistoryResponseTheorem || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate704BoundaryWoundMixtureBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
}

func TestTheoremNotesContainExpectedVerdicts(t *testing.T) {
	res := Generation2K7ComplementBoundaryWoundMixtureObservableAuditTheorem().Verify()
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
