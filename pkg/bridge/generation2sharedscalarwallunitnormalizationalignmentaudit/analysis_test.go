package generation2sharedscalarwallunitnormalizationalignmentaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestGate701InheritanceAndSharedLambda(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	inh := a.Inherited
	if !inh.InheritedQuotientNormalization || !nearly(inh.EventProbability, eventProbK7) || !nearly(inh.CanonicalCoefficient, eventProbK7) || !strings.Contains(inh.CoefficientCovarianceFormula, "beta/alpha") || !inh.Gate700CoordinateSealed || !inh.NoNativeWallAlignment || !inh.NoNativeBoundaryHistory || !inh.NoNativeSevenOver72 || inh.Verdict != StatusGate701QuotientNormalizationInherited {
		t.Fatalf("bad inheritance: %+v", inh)
	}
	shared := a.SharedLambda
	if !strings.Contains(shared.BoundaryCoordinate, "lambda") || !strings.Contains(shared.HistoryCoordinate, "lambda") || shared.SharedCoordinate != "lambda(Lambda_12)" || shared.BoundaryLambdaCoefficient != 1 || shared.HistoryLambdaCoefficient != 1 || !shared.BoundaryContainsSharedLambda || !shared.HistoryContainsSharedLambda || !shared.SameSignedScalarZeroWall || !shared.UnitCoefficientAlignment {
		t.Fatalf("bad shared lambda audit: %+v", shared)
	}
	for _, want := range []string{StatusSharedLambdaCoordinateIdentified, StatusLambdaUnitCoefficientAlignmentAudited, StatusSharedScalarWallUnitAnchorsQuotientNormalization} {
		if !strings.Contains(shared.Verdict, want) {
			t.Fatalf("missing %s in %q", want, shared.Verdict)
		}
	}
}

func TestAnchorKeepsCoefficientEqualToEventProbability(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	anchor := a.Anchor
	if !strings.Contains(anchor.Rule, "coefficient(lambda") || anchor.BoundaryLambdaCoefficient != 1 || anchor.HistoryLambdaCoefficient != 1 || anchor.Alpha != 1 || anchor.Beta != 1 || anchor.BetaOverAlpha != 1 || !nearly(anchor.ResponseCoefficient, eventProbK7) || !nearly(anchor.EventProbability, eventProbK7) || !anchor.CoefficientEqualsProbability {
		t.Fatalf("bad anchor audit: %+v", anchor)
	}
	for _, want := range []string{StatusLambdaUnitCoefficientAlignmentAudited, StatusResponseCoefficientRemainsEventProbabilityUnderSharedUnit, StatusResponseCoefficientEqualsEventProbabilitySharedLambda} {
		if !strings.Contains(anchor.Verdict, want) {
			t.Fatalf("missing %s in %q", want, anchor.Verdict)
		}
	}
}

func TestAlternativeNormalizations(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	alts := a.Alternatives
	if len(alts.Examples) != 5 || !alts.EuclideanBoundaryRejected || !alts.HistoryNormRejected || !alts.GaugeAnchorConditioned || !alts.AbsoluteFormRejected || !alts.SharedUnitAccepted || alts.Verdict != StatusAlternativeNormalizationsAudited {
		t.Fatalf("bad alternatives: %+v", alts)
	}
	if !nearly(alts.Examples[0].TransformedCoefficient, math.Sqrt2*eventProbK7) || alts.Examples[0].EqualsEventProbability {
		t.Fatalf("bad Euclidean boundary example: %+v", alts.Examples[0])
	}
	if !nearly(alts.Examples[1].TransformedCoefficient, eventProbK7/math.Sqrt(3)) || alts.Examples[1].EqualsEventProbability {
		t.Fatalf("bad history-normalized example: %+v", alts.Examples[1])
	}
	if !alts.Examples[2].EqualsEventProbability || !alts.Examples[2].PreservesLambdaUnitAlignment || alts.Examples[2].AcceptedActiveAlignment {
		t.Fatalf("bad gauge-anchor conditioned example: %+v", alts.Examples[2])
	}
	if alts.Examples[3].AcceptedActiveAlignment || !strings.Contains(alts.Examples[3].Reason, "erases signed wall orientation") {
		t.Fatalf("bad absolute-form example: %+v", alts.Examples[3])
	}
	if !alts.Examples[4].EqualsEventProbability || !alts.Examples[4].PreservesLambdaUnitAlignment || !alts.Examples[4].AcceptedActiveAlignment {
		t.Fatalf("bad shared-unit example: %+v", alts.Examples[4])
	}
}

func TestNonTautologySourceAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NonTautology
	if !n.SharedLambdaPresent || !strings.Contains(n.RearrangedEquation, "65/72") || !strings.Contains(n.RearrangedEquation, "R_3-1") || !nearly(n.LambdaWeight, -65.0/72.0) || !nearly(n.GaugeWeight, 7.0/72.0) || !n.IndependentGaugeWoundPresent || !n.CoefficientsDiffer || !n.LambdaIsAlignmentAnchorNotProof || !n.NonTautologicalRelationPreserved || n.Verdict != StatusNonTautologyWithSharedLambdaAudited {
		t.Fatalf("bad non-tautology audit: %+v", n)
	}
	s := a.Source
	if !strings.Contains(s.EventProbabilityRole, "p_K7") || !strings.Contains(s.LambdaUnitRole, "unit coefficient") || !strings.Contains(s.BoundaryRole, "signed scalar-wall units") || !strings.Contains(s.HistoryRole, "same signed scalar-wall units") || !strings.Contains(s.Conclusion, "same scalar-wall unit") {
		t.Fatalf("bad source classification: %+v", s)
	}
	for _, want := range []string{StatusSharedScalarWallUnitAnchorsQuotientNormalization, StatusResponseCoefficientEqualsEventProbabilitySharedLambda, StatusGate700LawScalarWallUnitSealed} {
		if !strings.Contains(s.Verdict, want) {
			t.Fatalf("missing %s in source verdict %q", want, s.Verdict)
		}
	}
	m := a.Missing
	if len(m.Theorems) != 6 || !strings.Contains(m.Verdict, StatusSharedLambdaUnitAlignmentNotNative) || !strings.Contains(m.Verdict, StatusNoNativeWallCoordinateNormalizationAlignmentTheorem) || !strings.Contains(m.Verdict, StatusNoNativeBoundaryHistoryResponsePrinciple) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing theorem audit: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsSharedLambdaAlignmentNative || f.ClaimsNativeWallNormalizationAlignment || f.ClaimsNativeBoundaryHistoryPrinciple || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate702SharedScalarWallUnitBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
}

func TestTheoremNotesContainExpectedVerdicts(t *testing.T) {
	res := Generation2SharedScalarWallUnitNormalizationAlignmentAuditTheorem().Verify()
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
