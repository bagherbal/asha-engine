package generation2sharedscalarwallunitnormalizationalignmentaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2SharedScalarWallUnitNormalizationAlignmentAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 702 — Shared Scalar-Wall Unit Normalization Alignment Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 702 — Shared Scalar-Wall Unit Normalization Alignment Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate702SharedScalarWallUnitBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate701 quotient normalization", Passed: a.Inherited.InheritedQuotientNormalization && math.Abs(a.Inherited.EventProbability-eventProbK7) < tolerance && math.Abs(a.Inherited.CanonicalCoefficient-eventProbK7) < tolerance && strings.Contains(a.Inherited.CoefficientCovarianceFormula, "beta/alpha") && a.Inherited.Gate700CoordinateSealed && a.Inherited.NoNativeWallAlignment && a.Inherited.NoNativeBoundaryHistory && a.Inherited.NoNativeSevenOver72 && a.Inherited.Verdict == StatusGate701QuotientNormalizationInherited, Detail: FormatInheritance(a.Inherited)},
			{Name: "identify shared signed scalar-wall coordinate", Passed: strings.Contains(a.SharedLambda.BoundaryCoordinate, "lambda") && strings.Contains(a.SharedLambda.HistoryCoordinate, "lambda") && a.SharedLambda.SharedCoordinate == "lambda(Lambda_12)" && a.SharedLambda.BoundaryLambdaCoefficient == 1 && a.SharedLambda.HistoryLambdaCoefficient == 1 && a.SharedLambda.BoundaryContainsSharedLambda && a.SharedLambda.HistoryContainsSharedLambda && a.SharedLambda.SameSignedScalarZeroWall && a.SharedLambda.UnitCoefficientAlignment && strings.Contains(a.SharedLambda.Verdict, StatusSharedLambdaCoordinateIdentified) && strings.Contains(a.SharedLambda.Verdict, StatusLambdaUnitCoefficientAlignmentAudited), Detail: FormatSharedLambda(a.SharedLambda)},
			{Name: "audit lambda unit coefficient alignment", Passed: strings.Contains(a.Anchor.Rule, "coefficient(lambda") && a.Anchor.BoundaryLambdaCoefficient == 1 && a.Anchor.HistoryLambdaCoefficient == 1 && a.Anchor.Alpha == 1 && a.Anchor.Beta == 1 && a.Anchor.BetaOverAlpha == 1 && math.Abs(a.Anchor.ResponseCoefficient-eventProbK7) < tolerance && math.Abs(a.Anchor.EventProbability-eventProbK7) < tolerance && a.Anchor.CoefficientEqualsProbability && strings.Contains(a.Anchor.Verdict, StatusResponseCoefficientRemainsEventProbabilityUnderSharedUnit) && strings.Contains(a.Anchor.Verdict, StatusResponseCoefficientEqualsEventProbabilitySharedLambda), Detail: FormatAnchor(a.Anchor)},
			{Name: "audit alternative normalizations", Passed: len(a.Alternatives.Examples) == 5 && a.Alternatives.EuclideanBoundaryRejected && a.Alternatives.HistoryNormRejected && a.Alternatives.GaugeAnchorConditioned && a.Alternatives.AbsoluteFormRejected && a.Alternatives.SharedUnitAccepted && math.Abs(a.Alternatives.Examples[0].TransformedCoefficient-math.Sqrt2*eventProbK7) < tolerance && math.Abs(a.Alternatives.Examples[1].TransformedCoefficient-eventProbK7/math.Sqrt(3)) < tolerance && a.Alternatives.Examples[4].AcceptedActiveAlignment && a.Alternatives.Verdict == StatusAlternativeNormalizationsAudited, Detail: FormatAlternatives(a.Alternatives)},
			{Name: "preserve non-tautology with shared lambda", Passed: a.NonTautology.SharedLambdaPresent && strings.Contains(a.NonTautology.RearrangedEquation, "65/72") && strings.Contains(a.NonTautology.RearrangedEquation, "R_3-1") && math.Abs(a.NonTautology.LambdaWeight+65.0/72.0) < tolerance && math.Abs(a.NonTautology.GaugeWeight-7.0/72.0) < tolerance && a.NonTautology.IndependentGaugeWoundPresent && a.NonTautology.CoefficientsDiffer && a.NonTautology.LambdaIsAlignmentAnchorNotProof && a.NonTautology.NonTautologicalRelationPreserved && a.NonTautology.Verdict == StatusNonTautologyWithSharedLambdaAudited, Detail: FormatNonTautology(a.NonTautology)},
			{Name: "classify source types", Passed: strings.Contains(a.Source.EventProbabilityRole, "p_K7") && strings.Contains(a.Source.LambdaUnitRole, "unit coefficient") && strings.Contains(a.Source.BoundaryRole, "signed scalar-wall units") && strings.Contains(a.Source.HistoryRole, "same signed scalar-wall units") && strings.Contains(a.Source.Conclusion, "same scalar-wall unit") && strings.Contains(a.Source.Verdict, StatusSharedScalarWallUnitAnchorsQuotientNormalization) && strings.Contains(a.Source.Verdict, StatusGate700LawScalarWallUnitSealed), Detail: FormatSource(a.Source)},
			{Name: "preserve missing theorem boundary", Passed: len(a.Missing.Theorems) == 6 && containsAll(a.Missing.Verdict, StatusSharedLambdaUnitAlignmentNotNative, StatusNoNativeWallCoordinateNormalizationAlignmentTheorem, StatusNoNativeBoundaryHistoryResponsePrinciple, StatusNoNativeSevenOver72Theorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve Gate702 shared scalar-wall unit firewall", Passed: !a.Firewall.ClaimsSharedLambdaAlignmentNative && !a.Firewall.ClaimsNativeWallNormalizationAlignment && !a.Firewall.ClaimsNativeBoundaryHistoryPrinciple && !a.Firewall.ClaimsNativeSevenOver72Theorem && !a.Firewall.ClaimsBoundaryStressDerived && !a.Firewall.ClaimsScalarRGMatching && !a.Firewall.ClaimsHiggsMass && !a.Firewall.ClaimsGaugeUnification && !a.Firewall.ClaimsFlavorDerivation && !a.Firewall.ClaimsCKMPMNS && a.Firewall.Verdict == StatusGate702SharedScalarWallUnitBoundary, Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 702 — Shared Scalar-Wall Unit Normalization Alignment Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func containsAll(s string, xs ...string) bool {
	for _, x := range xs {
		if !strings.Contains(s, x) {
			return false
		}
	}
	return true
}
