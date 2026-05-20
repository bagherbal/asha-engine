package generation2scalarwallairlockandquotientlinegluingaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarWallAirlockAndQuotientLineGluingAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 703 — Scalar-Wall Airlock and Quotient-Line Gluing Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 703 — Scalar-Wall Airlock and Quotient-Line Gluing Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate703ScalarWallAirlockBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate702 shared scalar-wall unit", Passed: a.Inherited.InheritedSharedScalarWallUnit && a.Inherited.SharedCoordinate == "lambda(Lambda_12)" && math.Abs(a.Inherited.EventProbability-eventProbK7) < tolerance && math.Abs(a.Inherited.ResponseCoefficient-eventProbK7) < tolerance && a.Inherited.CoefficientEqualsProbability && a.Inherited.NonTautologyInherited && a.Inherited.NoNativeSharedUnit && a.Inherited.NoNativeWallAlignment && a.Inherited.NoNativeBoundaryHistory && a.Inherited.NoNativeSevenOver72 && a.Inherited.Verdict == StatusGate702SharedScalarWallUnitInherited, Detail: FormatInheritance(a.Inherited)},
			{Name: "define scalar-wall line", Passed: a.ScalarWall.LineName == "L_lambda" && strings.Contains(a.ScalarWall.Definition, "span") && a.ScalarWall.Coordinate == "lambda(Lambda_12)" && a.ScalarWall.SignedOrientationPreserved && a.ScalarWall.BoundaryLambdaCoefficient == 1 && a.ScalarWall.HistoryLambdaCoefficient == 1 && a.ScalarWall.UnitScalarWallOnBothSides && containsAll(a.ScalarWall.Verdict, StatusScalarWallLineDefined, StatusSharedLambdaIsScalarWallAirlock), Detail: FormatScalarWall(a.ScalarWall)},
			{Name: "define quotient-line gluing diagram", Passed: strings.Contains(a.Diagram.Diagram, "Q_boundary") && strings.Contains(a.Diagram.Diagram, "L_lambda") && strings.Contains(a.Diagram.Diagram, "Q_history") && strings.Contains(a.Diagram.BoundaryCoordinate, "lambda") && strings.Contains(a.Diagram.HistoryCoordinate, "lambda") && a.Diagram.BoundaryMeasuredInLambda && a.Diagram.HistoryMeasuredInLambda && a.Diagram.CoordinatesAreAnchored && a.Diagram.Verdict == StatusQuotientLineGluingDiagramDefined, Detail: FormatDiagram(a.Diagram)},
			{Name: "audit unit lambda glue condition", Passed: a.UnitGlue.BoundaryLambdaCoefficient == 1 && a.UnitGlue.HistoryLambdaCoefficient == 1 && math.Abs(a.UnitGlue.Gamma-1) < tolerance && a.UnitGlue.UnitGlue && math.Abs(a.UnitGlue.EventProbability-eventProbK7) < tolerance && math.Abs(a.UnitGlue.ResponseCoefficient-eventProbK7) < tolerance && a.UnitGlue.CoefficientEqualsProbability && containsAll(a.UnitGlue.Verdict, StatusUnitLambdaGlueConditionAudited, StatusResponseCoefficientPreservationComputed, StatusResponseCoefficientEqualsProbabilityAfterUnitGlue), Detail: FormatUnitGlue(a.UnitGlue)},
			{Name: "compute response coefficient under rescaled gluing", Passed: math.Abs(a.Rescaled.Gamma-2) < tolerance && math.Abs(a.Rescaled.TransformedCoefficient-2*eventProbK7) < tolerance && !a.Rescaled.EqualsEventProbability && a.Rescaled.RequiresGammaOneForEquality && strings.Contains(a.Rescaled.Formula, "gamma p_K7") && a.Rescaled.Verdict == StatusResponseCoefficientPreservationComputed, Detail: FormatRescaled(a.Rescaled)},
			{Name: "audit alternative gluings", Passed: len(a.Alternatives.Examples) == 5 && a.Alternatives.BoundaryNormalizedRejected && a.Alternatives.HistoryNormalizedRejected && a.Alternatives.AbsoluteScalarRejected && a.Alternatives.HessianScalarRejected && a.Alternatives.SharedSignedLambdaAccepted && math.Abs(a.Alternatives.Examples[0].ResponseCoefficient-math.Sqrt2*eventProbK7) < tolerance && math.Abs(a.Alternatives.Examples[1].ResponseCoefficient-eventProbK7/math.Sqrt(3)) < tolerance && math.Abs(a.Alternatives.Examples[3].ResponseCoefficient-2*eventProbK7) < tolerance && a.Alternatives.Examples[4].AcceptedActiveAirlock && a.Alternatives.Verdict == StatusAlternativeGluingsAudited, Detail: FormatAlternatives(a.Alternatives)},
			{Name: "preserve non-tautology of shared lambda", Passed: a.NonTautology.SharedLambdaAirlockPresent && strings.Contains(a.NonTautology.RearrangedEquation, "65/72") && strings.Contains(a.NonTautology.RearrangedEquation, "R_3-1") && math.Abs(a.NonTautology.LambdaWeight+65.0/72.0) < tolerance && math.Abs(a.NonTautology.GaugeWeight-7.0/72.0) < tolerance && a.NonTautology.IndependentGaugeWoundPresent && a.NonTautology.LambdaIsAirlockNotProof && a.NonTautology.NonTautologicalRelationPreserved && a.NonTautology.Verdict == StatusNonTautologyOfSharedLambdaAudited, Detail: FormatNonTautology(a.NonTautology)},
			{Name: "classify source types", Passed: strings.Contains(a.Source.ScalarWallLineRole, "L_lambda") && strings.Contains(a.Source.BoundaryRole, "scalar-wall units") && strings.Contains(a.Source.HistoryRole, "scalar-wall units") && strings.Contains(a.Source.EventProbabilityRole, "p_K7") && strings.Contains(a.Source.ResponseCoefficientRole, "unit scalar-wall gluing") && containsAll(a.Source.Verdict, StatusSharedLambdaIsScalarWallAirlock, StatusResponseCoefficientEqualsProbabilityAfterUnitGlue, StatusGate700LawScalarWallGluedQuotientResponse), Detail: FormatSource(a.Source)},
			{Name: "preserve missing theorem boundary", Passed: len(a.Missing.Theorems) == 6 && containsAll(a.Missing.Verdict, StatusScalarWallGluingNotNative, StatusNoNativeScalarWallAirlockTheorem, StatusNoNativeBoundaryHistoryResponsePrinciple, StatusNoNativeSevenOver72Theorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve Gate703 scalar-wall airlock firewall", Passed: !a.Firewall.ClaimsScalarWallGluingNative && !a.Firewall.ClaimsNativeScalarWallAirlock && !a.Firewall.ClaimsNativeBoundaryHistory && !a.Firewall.ClaimsNativeSevenOver72Theorem && !a.Firewall.ClaimsBoundaryStressDerived && !a.Firewall.ClaimsScalarRGMatching && !a.Firewall.ClaimsHiggsMass && !a.Firewall.ClaimsGaugeUnification && !a.Firewall.ClaimsFlavorDerivation && !a.Firewall.ClaimsCKMPMNS && a.Firewall.Verdict == StatusGate703ScalarWallAirlockBoundary, Detail: FormatFirewall(a.Firewall)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 703 — Scalar-Wall Airlock and Quotient-Line Gluing Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
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
