package generation2scalarwallairlockandquotientlinegluingaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate702InheritanceScalarWallAndDiagram(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	inh := a.Inherited
	if !inh.InheritedSharedScalarWallUnit || inh.SharedCoordinate != "lambda(Lambda_12)" || !nearly(inh.EventProbability, eventProbK7) || !nearly(inh.ResponseCoefficient, eventProbK7) || !inh.CoefficientEqualsProbability || !inh.NonTautologyInherited || !inh.NoNativeSharedUnit || !inh.NoNativeWallAlignment || !inh.NoNativeBoundaryHistory || !inh.NoNativeSevenOver72 || inh.Verdict != StatusGate702SharedScalarWallUnitInherited {
		t.Fatalf("bad inheritance: %+v", inh)
	}
	w := a.ScalarWall
	if w.LineName != "L_lambda" || !strings.Contains(w.Definition, "span") || w.Coordinate != "lambda(Lambda_12)" || !w.SignedOrientationPreserved || w.BoundaryLambdaCoefficient != 1 || w.HistoryLambdaCoefficient != 1 || !w.UnitScalarWallOnBothSides || !strings.Contains(w.Verdict, StatusScalarWallLineDefined) || !strings.Contains(w.Verdict, StatusSharedLambdaIsScalarWallAirlock) {
		t.Fatalf("bad scalar wall: %+v", w)
	}
	d := a.Diagram
	if !strings.Contains(d.Diagram, "Q_boundary") || !strings.Contains(d.Diagram, "L_lambda") || !strings.Contains(d.Diagram, "Q_history") || !d.BoundaryMeasuredInLambda || !d.HistoryMeasuredInLambda || !d.CoordinatesAreAnchored || d.Verdict != StatusQuotientLineGluingDiagramDefined {
		t.Fatalf("bad diagram: %+v", d)
	}
}

func TestUnitGlueAndRescaledCoefficient(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	u := a.UnitGlue
	if u.BoundaryLambdaCoefficient != 1 || u.HistoryLambdaCoefficient != 1 || !nearly(u.Gamma, 1) || !u.UnitGlue || !nearly(u.EventProbability, eventProbK7) || !nearly(u.ResponseCoefficient, eventProbK7) || !u.CoefficientEqualsProbability || !strings.Contains(u.Verdict, StatusResponseCoefficientEqualsProbabilityAfterUnitGlue) {
		t.Fatalf("bad unit glue: %+v", u)
	}
	r := a.Rescaled
	if !nearly(r.Gamma, 2) || !nearly(r.TransformedCoefficient, 2*eventProbK7) || r.EqualsEventProbability || !r.RequiresGammaOneForEquality || !strings.Contains(r.Formula, "gamma p_K7") || r.Verdict != StatusResponseCoefficientPreservationComputed {
		t.Fatalf("bad rescaled glue: %+v", r)
	}
}

func TestAlternativeGluings(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	alts := a.Alternatives
	if len(alts.Examples) != 5 || !alts.BoundaryNormalizedRejected || !alts.HistoryNormalizedRejected || !alts.AbsoluteScalarRejected || !alts.HessianScalarRejected || !alts.SharedSignedLambdaAccepted || alts.Verdict != StatusAlternativeGluingsAudited {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	if !nearly(alts.Examples[0].ResponseCoefficient, math.Sqrt2*eventProbK7) || !alts.Examples[0].Rejected {
		t.Fatalf("bad boundary-normalized example: %+v", alts.Examples[0])
	}
	if !nearly(alts.Examples[1].ResponseCoefficient, eventProbK7/math.Sqrt(3)) || !alts.Examples[1].Rejected {
		t.Fatalf("bad history-normalized example: %+v", alts.Examples[1])
	}
	if !strings.Contains(alts.Examples[2].Reason, "erases scalar-wall orientation") || !alts.Examples[2].Rejected {
		t.Fatalf("bad absolute scalar example: %+v", alts.Examples[2])
	}
	if !nearly(alts.Examples[3].ResponseCoefficient, 2*eventProbK7) || !strings.Contains(alts.Examples[3].Reason, "Hessian") || !alts.Examples[3].Rejected {
		t.Fatalf("bad Hessian scalar example: %+v", alts.Examples[3])
	}
	if !alts.Examples[4].AcceptedActiveAirlock || alts.Examples[4].Rejected || !nearly(alts.Examples[4].ResponseCoefficient, eventProbK7) {
		t.Fatalf("bad shared signed lambda example: %+v", alts.Examples[4])
	}
}

func TestNonTautologySourceMissingAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonTautology
	if !n.SharedLambdaAirlockPresent || !strings.Contains(n.RearrangedEquation, "65/72") || !strings.Contains(n.RearrangedEquation, "R_3-1") || !nearly(n.LambdaWeight, -65.0/72.0) || !nearly(n.GaugeWeight, 7.0/72.0) || !n.IndependentGaugeWoundPresent || !n.LambdaIsAirlockNotProof || !n.NonTautologicalRelationPreserved || n.Verdict != StatusNonTautologyOfSharedLambdaAudited {
		t.Fatalf("bad non-tautology audit: %+v", n)
	}
	s := a.Source
	if !strings.Contains(s.ScalarWallLineRole, "L_lambda") || !strings.Contains(s.BoundaryRole, "scalar-wall units") || !strings.Contains(s.HistoryRole, "scalar-wall units") || !strings.Contains(s.EventProbabilityRole, "p_K7") || !strings.Contains(s.ResponseCoefficientRole, "unit scalar-wall gluing") || !strings.Contains(s.Conclusion, "scalar-wall glued quotient response") {
		t.Fatalf("bad source classification: %+v", s)
	}
	m := a.Missing
	if len(m.Theorems) != 6 || !strings.Contains(m.Verdict, StatusScalarWallGluingNotNative) || !strings.Contains(m.Verdict, StatusNoNativeScalarWallAirlockTheorem) || !strings.Contains(m.Verdict, StatusNoNativeBoundaryHistoryResponsePrinciple) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsScalarWallGluingNative || f.ClaimsNativeScalarWallAirlock || f.ClaimsNativeBoundaryHistory || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate703ScalarWallAirlockBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
}

func TestTheoremNotesContainExpectedVerdicts(t *testing.T) {
	res := Generation2ScalarWallAirlockAndQuotientLineGluingAuditTheorem().Verify()
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
