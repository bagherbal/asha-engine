package generation2boundaryquotientpayofffunctionalselectionaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate696InheritanceAndPayoffProblem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SupportLocalBernoulliObservableInherited || a.Inherited.Rho72Definition != "rho_72 = I_H72/72" || a.Inherited.SupportLocalObservable != "R_a=aP_K7" || a.Inherited.BoundaryPayoffAssignment != "a=S_split" || a.Inherited.ReconstructedOperator != "R_split=S_split P_K7" || !nearly(a.Inherited.EventPayoff, a.Inherited.SSplit) || a.Inherited.ComplementPayoff != 0 || !a.Inherited.NoExpectationAloneSelection || !a.Inherited.NoNativeSupportLocality || !a.Inherited.NoNativeSSplitPayoff || !a.Inherited.NoNativeSevenOver72 {
		t.Fatalf("bad Gate696 inheritance: %+v", a.Inherited)
	}
	if a.PayoffProblem.RemainingQuestion != "a ?= S_split" || !a.PayoffProblem.SupportLocalityFixesB || !a.PayoffProblem.EventPayoffStillUnfixed || !strings.Contains(a.PayoffProblem.CandidatePayoff, "lambda(Lambda_12)") || a.PayoffProblem.Verdict != StatusPayoffSourceProblemDefined {
		t.Fatalf("bad payoff source problem: %+v", a.PayoffProblem)
	}
}

func TestBoundaryQuotientSigmaDescends(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	q := a.Quotient
	if q.BoundarySpace != "B_boundary=span(lambda,R_3-1)" || q.AntiAlignmentGenerator != [2]float64{-1, 1} || !nearly(q.BoundaryVector[0], lambdaLambda12) || !nearly(q.BoundaryVector[1], r3Minus1) {
		t.Fatalf("bad boundary objects: %+v", q)
	}
	if math.Abs(q.SigmaOnAntiAlignment) > tolerance || !q.KernelMatchesWall || !q.DescendsToQuotient || q.QuotientSpace != "Q_boundary=B_boundary/L_anti" || !nearly(q.SSplit, lambdaLambda12+r3Minus1) || math.Abs(q.SSplit-a.Inherited.SSplit) > 5e-16 || !q.SSplitMatchesSigma {
		t.Fatalf("sigma does not descend correctly: %+v", q)
	}
	for _, want := range []string{StatusBoundaryAntiAlignmentWallDefined, StatusSigmaBoundaryDescendsToQuotientCoordinate, StatusSSplitIdentifiedAsBoundaryQuotientPayoff} {
		if !strings.Contains(q.Verdict, want) {
			t.Fatalf("missing quotient verdict %s in %q", want, q.Verdict)
		}
	}
}

func TestPayoffInterpretationAndAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Interpretation.EventPayoff != "a=S_split=sigma_boundary(b)" || a.Interpretation.Observable != "R_split=sigma_boundary(b)P_K7" || !strings.Contains(a.Interpretation.EquivalentOperator, "R_3-1") || !nearly(a.Interpretation.Payoff, a.Inherited.SSplit) || !a.Interpretation.K7ReceivesBoundaryDefect || !a.Interpretation.DoesNotProveCoupling {
		t.Fatalf("bad payoff interpretation: %+v", a.Interpretation)
	}
	if len(a.Alternatives.Alternatives) != 5 || !a.Alternatives.LambdaOnlyRejected || !a.Alternatives.GaugeOnlyRejected || !a.Alternatives.AntiAlignedRejected || !a.Alternatives.MidpointStressRejected || !a.Alternatives.SplitPayoffAccepted || !a.Alternatives.AllAudited {
		t.Fatalf("bad alternatives audit: %+v", a.Alternatives)
	}
	byName := map[string]BoundaryPayoffAlternative{}
	for _, x := range a.Alternatives.Alternatives {
		byName[x.Name] = x
	}
	if byName["lambda-only payoff"].VanishesOnWall || byName["lambda-only payoff"].Active || byName["gauge-only payoff"].VanishesOnWall || byName["gauge-only payoff"].Active {
		t.Fatalf("one-coordinate payoff not rejected: %+v", byName)
	}
	if byName["anti-aligned magnitude"].MeasuresQuotient || byName["midpoint stress"].MeasuresQuotient {
		t.Fatalf("non-quotient alternatives misclassified: %+v", byName)
	}
	if !byName["split payoff"].VanishesOnWall || !byName["split payoff"].MeasuresQuotient || !byName["split payoff"].Active || !nearly(byName["split payoff"].Value, a.Inherited.SSplit) {
		t.Fatalf("split payoff not accepted: %+v", byName["split payoff"])
	}
}

func TestScaleFirewallAndExpectation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ScaleFirewall.QuotientCoordinateUniqueUpToScale || a.ScaleFirewall.ScaledFunctional != "c*sigma_boundary" || !strings.Contains(a.ScaleFirewall.ActiveNormalizationSource, "Gates 668-670") || a.ScaleFirewall.UnitCoefficientLambda != 1 || a.ScaleFirewall.UnitCoefficientGauge != 1 || a.ScaleFirewall.ClaimsNativePayoffNormalization || a.ScaleFirewall.Verdict != StatusPayoffFunctionalUniqueOnlyUpToNormalization {
		t.Fatalf("bad scale firewall: %+v", a.ScaleFirewall)
	}
	if a.Expectation.Rho72 != "rho_72=I_H72/72" || a.Expectation.ResponseOperator != "R_split=sigma_boundary(b)P_K7" || !nearly(a.Expectation.K7Weight, 7.0/72.0) || !strings.Contains(a.Expectation.ExpectationFormula, "(7/72)S_split") || !nearly(a.Expectation.Expectation, a.Inherited.Expectation) || !nearly(a.Expectation.DBase, a.Inherited.DBase) || math.Abs(a.Expectation.ResidualE1-a.Inherited.ResidualE1) > 1e-18 || !a.Expectation.MatchesInherited {
		t.Fatalf("bad event expectation reconstruction: %+v", a.Expectation)
	}
}

func TestSourceTypesMissingAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.SourceTypes.BoundarySpaceRole, "boundary") || !strings.Contains(a.SourceTypes.SigmaBoundaryRole, "quotient") || !strings.Contains(a.SourceTypes.SSplitRole, "payoff") || !strings.Contains(a.SourceTypes.PK7Role, "Boolean-octonionic") || !strings.Contains(a.SourceTypes.Rho72Role, "maximum-entropy") || !strings.Contains(a.SourceTypes.RSplitRole, "K7 event") {
		t.Fatalf("bad source types: %+v", a.SourceTypes)
	}
	if len(a.Missing.Candidates) != 5 || !strings.Contains(a.Missing.PreciseGap, "boundary anti-alignment quotient") || !strings.Contains(a.Missing.Verdict, StatusNoNativeReasonK7ReceivesBoundaryQuotientPayoff) || !strings.Contains(a.Missing.Verdict, StatusNoNativePayoffCouplingTheorem) || !strings.Contains(a.Missing.Verdict, StatusNoNativeHistoryResponseTheorem) || !strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Firewall.ClaimsK7ReceivesBoundaryPayoffNatively || a.Firewall.ClaimsNativePayoffCouplingTheorem || a.Firewall.ClaimsNativeHistoryResponseTheorem || a.Firewall.ClaimsNativeSevenOver72Theorem || a.Firewall.ClaimsBoundaryStressDerived || a.Firewall.ClaimsScalarRGMatching || a.Firewall.ClaimsHiggsMass || a.Firewall.ClaimsGaugeUnification || a.Firewall.ClaimsFlavorDerivation || a.Firewall.ClaimsCKMPMNS || a.Firewall.Verdict != StatusGate697BoundaryQuotientPayoffBoundary {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryQuotientPayoffFunctionalSelectionAuditTheorem().Verify()
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
