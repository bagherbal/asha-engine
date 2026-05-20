package generation2activesevenoverseventytwoboundaryweightsourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ActiveSevenOverSeventyTwoBoundaryWeightSourceTypeAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 660 — Active Seven-Over-Seventy-Two Boundary Weight Source-Type Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate660 active 7/72 source-type audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate659 boundary-weighted closure", Passed: a.Inherited.BoundaryWeightedClosureInherited && a.Inherited.ActiveTransportLane && a.Inherited.FanoHitchinRouteSealed && math.Abs(a.Inherited.W72-0.04982659643506822) < 5e-15 && math.Abs(a.Inherited.WeightedResidual-8.525834413464217e-10) < 5e-18 && a.Inherited.NoNativeSevenOver72Theorem && a.Inherited.NoNativeKappaClosureTheorem && a.Inherited.NoNativeTransportTheorem && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit numerator seven source candidates", Passed: len(a.Numerator.Candidates) == 4 && a.Numerator.K7CarrierDimension == 7 && a.Numerator.FanoHitchinStrengthens7 && a.Numerator.IntersectionDefect7 && a.Numerator.CokernelDefect7 && !a.Numerator.BoundaryMapConstructed, Detail: FormatNumerator(a.Numerator)},
			{Name: "audit denominator seventy-two source candidates", Passed: len(a.Denominator.Candidates) == 4 && a.Denominator.Lambda4PlusBoundaryPairDim == 72 && a.Denominator.PreferredCandidate == "70+2" && a.Denominator.AugmentedChamberTyped && a.Denominator.BoundaryPairEnvironmental && !a.Denominator.NativeTraceTheorem, Detail: FormatDenominator(a.Denominator)},
			{Name: "define active W72 interpolation", Passed: math.Abs(a.Interpolation.Weight-sevenOver72) < 1e-15 && math.Abs(a.Interpolation.ComplementWeight-65.0/72.0) < 1e-15 && math.Abs(a.Interpolation.W72-0.04982659643506822) < 5e-15 && math.Abs(a.Interpolation.Residual-8.525834413464217e-10) < 5e-18 && a.Interpolation.ActiveTransport && !a.Interpolation.FanoHitchinRoute, Detail: FormatInterpolation(a.Interpolation)},
			{Name: "lift W72 into scalar runtime matching formula", Passed: math.Abs(a.FormulaLift.KappaLambdaFromW72Exact-0.044323042243493656) < 5e-15 && math.Abs(a.FormulaLift.KappaLambdaResidual+8.525834413464217e-10) < 5e-18 && math.Abs(a.FormulaLift.RuntimePredictionExactKappaE-0.12965256505471276) < 5e-15 && math.Abs(a.FormulaLift.RuntimeResidualExactKappaE-4.2369718844526005e-12) < 5e-17 && math.Abs(a.FormulaLift.RuntimeResidualOrientKappaE-1.3799595133257014e-08) < 5e-15 && a.FormulaLift.BridgeLayerOnly, Detail: FormatFormulaLift(a.FormulaLift)},
			{Name: "audit residual hierarchy", Passed: a.Residuals.RawToW72Improvement > 100000 && a.Residuals.WeightedToBoundarySplit < 1e-6 && a.Residuals.ExactRuntimeRelative < 4e-11 && math.Abs(a.Residuals.RuntimeResidualExactKappaE-4.2369718844526005e-12) < 5e-17, Detail: FormatResiduals(a.Residuals)},
			{Name: "classify active source type without Fano revival", Passed: a.Classification.SevenOver72ActingAsK7TraceWeight && a.Classification.SevenOver72ActingAsAugmentedDimension && a.Classification.SevenOver72ActingAsBoundaryWeight && !a.Classification.SevenOver72ActingAsTransportArtifact && !a.Classification.SevenOver72UnsourcedEnvironmentalCoeff && !a.Classification.FanoHitchinBoundaryMapConstructed && !a.Classification.RandomConstantSearch, Detail: FormatClassification(a.Classification)},
			{Name: "preserve active 7/72 source-type firewalls", Passed: !a.Firewalls.ClaimsNativeSevenOver72Theorem && !a.Firewalls.ClaimsNativeK7BoundaryMap && !a.Firewalls.ClaimsNativeScalarFlavorTransport && !a.Firewalls.ClaimsBoundaryStressDerivation && !a.Firewalls.ClaimsHiggsPrediction && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsFlavorDerivation && !a.Firewalls.ClaimsCKMPMNSDerivation && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsFanoHitchinBoundaryMap && !a.Firewalls.ClaimsPhysicalSpacetime && a.Firewalls.Verdict == StatusGate660Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func notesContain(notes []string, want string) bool {
	return strings.Contains(strings.Join(notes, "\n"), want)
}
