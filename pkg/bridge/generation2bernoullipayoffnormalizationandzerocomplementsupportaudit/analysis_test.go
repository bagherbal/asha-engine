package generation2bernoullipayoffnormalizationandzerocomplementsupportaudit

import (
	"math"
	"strings"
	"testing"
)

func TestInheritanceAndGeneralPayoffObservable(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.BernoulliObservableInherited || a.Inherited.Rho72Definition != "rho_72 = I_H72/72" || a.Inherited.EventProjector != "E_K7 = P_K7" || a.Inherited.ResponseOperator != "R_split = S_split P_K7" {
		t.Fatalf("bad Gate695 inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.EventWeight-7.0/72.0) > tolerance || math.Abs(a.Inherited.ComplementWeight-65.0/72.0) > tolerance || a.Inherited.H72Dimension != h72Dimension || a.Inherited.K7Dimension != k7Dimension || a.Inherited.ComplementDimension != k7ComplementDim {
		t.Fatalf("bad inherited dimensions/weights: %+v", a.Inherited)
	}
	if !a.Inherited.NoNativeHistoryResponse || !a.Inherited.NoNativeRho72Reason || !a.Inherited.NoNativePayoffReason || !a.Inherited.NoNativeSevenOver72 {
		t.Fatalf("Gate695 firewall not inherited: %+v", a.Inherited)
	}
	if a.General.Observable != "R_{a,b}=aP_K7+bP_perp" || a.General.ComplementProjector != "P_perp=I_H72-P_K7" || math.Abs(a.General.EventWeight-7.0/72.0) > tolerance || math.Abs(a.General.ComplementWeight-65.0/72.0) > tolerance {
		t.Fatalf("bad general payoff observable: %+v", a.General)
	}
	if math.Abs(a.General.ActiveEventPayoff-a.Inherited.SSplit) > tolerance || a.General.ActiveComplementPayoff != 0 || math.Abs(a.General.ActiveExpectation-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("bad active specialization: %+v", a.General)
	}
}

func TestAffineDegeneracyAndSupportLocality(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Degeneracy.ExpectationAloneDegenerate || !a.Degeneracy.AlternativeDifferentFromActive || math.Abs(a.Degeneracy.AlternativeExpectation-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("affine payoff degeneracy not certified: %+v", a.Degeneracy)
	}
	wantAltB := (7.0 / 65.0) * a.Inherited.SSplit
	if a.Degeneracy.AlternativeEventPayoff != 0 || math.Abs(a.Degeneracy.AlternativeComplementPayoff-wantAltB) > tolerance {
		t.Fatalf("bad explicit degenerate witness: %+v", a.Degeneracy)
	}
	if len(a.SupportLocality.Conditions) != 3 || !a.SupportLocality.ComplementPayoffForcedZero || !a.SupportLocality.EventPayoffUnfixedBySupportLocality {
		t.Fatalf("bad support locality audit: %+v", a.SupportLocality)
	}
	if !strings.Contains(a.SupportLocality.PPerpLeftAction, "bP_perp") || !strings.Contains(a.SupportLocality.PPerpRightAction, "bP_perp") || !strings.Contains(a.SupportLocality.PK7SandwichAction, "aP_K7") {
		t.Fatalf("bad support locality actions: %+v", a.SupportLocality)
	}
}

func TestActiveAssignmentAndAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Assignment.SupportLocalObservable != "R_a=aP_K7" || a.Assignment.BoundaryPayoffAssignment != "a=S_split" || math.Abs(a.Assignment.EventPayoff-a.Inherited.SSplit) > tolerance || a.Assignment.ComplementPayoff != 0 || a.Assignment.ReconstructedOperator != "R_split=S_split P_K7" || !a.Assignment.MatchesInheritedActive || !a.Assignment.DoesNotProvePayoffNatively {
		t.Fatalf("bad active payoff assignment: %+v", a.Assignment)
	}
	if len(a.Alternatives.Alternatives) != 5 || !a.Alternatives.FullPayoffRejected || !a.Alternatives.ComplementPayoffRejected || !a.Alternatives.CenteredRejected || !a.Alternatives.SignedHodgeRejected || !a.Alternatives.ActiveAccepted || !a.Alternatives.AllAudited {
		t.Fatalf("bad alternative payoff audit: %+v", a.Alternatives)
	}
	values := map[string]float64{}
	active := map[string]bool{}
	for _, x := range a.Alternatives.Alternatives {
		values[x.Name] = x.Expectation
		active[x.Name] = x.Active
	}
	if math.Abs(values["full payoff"]-a.Inherited.SSplit) > tolerance || math.Abs(values["complement payoff"]-(65.0/72.0)*a.Inherited.SSplit) > tolerance || math.Abs(values["centered observable"]) > tolerance || math.Abs(values["signed Hodge payoff"]-a.Inherited.SSplit/72.0) > tolerance || math.Abs(values["active support-local observable"]-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("unexpected alternative expectations: %+v", values)
	}
	if active["full payoff"] || active["complement payoff"] || active["centered observable"] || active["signed Hodge payoff"] || !active["active support-local observable"] {
		t.Fatalf("unexpected alternative active map: %+v", active)
	}
}

func TestSourceTypesMissingTheoremsAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.SourceTypes.PK7Role, "Boolean-octonionic") || !strings.Contains(a.SourceTypes.PPerpRole, "no-response") || !strings.Contains(a.SourceTypes.SSplitRole, "boundary") || !strings.Contains(a.SourceTypes.Rho72Role, "no-bias") || !strings.Contains(a.SourceTypes.RSplitRole, "support-local") {
		t.Fatalf("bad source-type audit: %+v", a.SourceTypes)
	}
	if len(a.Missing.Missing) != 4 || !strings.Contains(a.Missing.PreciseGap, "why physical history imposes K7 support-locality") || !strings.Contains(a.Missing.Verdict, StatusNoNativeReasonHistoryUsesSupportLocality) || !strings.Contains(a.Missing.Verdict, StatusNoNativeReasonK7EventReceivesSSplitPayoff) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Firewall.ClaimsHistoryUsesSupportLocalityNatively || a.Firewall.ClaimsK7PhysicalEventNatively || a.Firewall.ClaimsSSplitPayoffNatively || a.Firewall.ClaimsExpectationEqualsDBaseNatively || a.Firewall.ClaimsResidualExplained || a.Firewall.ClaimsNativeStateSelectionTheorem || a.Firewall.ClaimsNativePayoffTheorem || a.Firewall.ClaimsNativeSevenOver72Theorem || a.Firewall.ClaimsBoundaryStress || a.Firewall.ClaimsScalarRGMatching || a.Firewall.ClaimsHiggsMass || a.Firewall.ClaimsGaugeUnification || a.Firewall.ClaimsFlavorDerivation || a.Firewall.ClaimsCKMPMNS {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BernoulliPayoffNormalizationAndZeroComplementSupportAuditTheorem().Verify()
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
