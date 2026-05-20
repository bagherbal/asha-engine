package generation2activesevenoverseventytwoboundaryweightsourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate660Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.BoundaryWeightedClosureInherited || !a.Inherited.ActiveTransportLane || !a.Inherited.FanoHitchinRouteSealed || math.Abs(a.Inherited.W72-0.04982659643506822) > 5e-15 || math.Abs(a.Inherited.WeightedResidual-8.525834413464217e-10) > 5e-18 || !a.Inherited.NoNativeSevenOver72Theorem || !a.Inherited.NoNativeKappaClosureTheorem || !a.Inherited.NoNativeTransportTheorem || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if len(a.Numerator.Candidates) != 4 || a.Numerator.K7CarrierDimension != 7 || !a.Numerator.FanoHitchinStrengthens7 || !a.Numerator.IntersectionDefect7 || !a.Numerator.CokernelDefect7 || a.Numerator.BoundaryMapConstructed {
		t.Fatalf("bad numerator audit: %+v", a.Numerator)
	}
	if len(a.Denominator.Candidates) != 4 || a.Denominator.Lambda4PlusBoundaryPairDim != 72 || a.Denominator.PreferredCandidate != "70+2" || !a.Denominator.AugmentedChamberTyped || !a.Denominator.BoundaryPairEnvironmental || a.Denominator.NativeTraceTheorem {
		t.Fatalf("bad denominator audit: %+v", a.Denominator)
	}
	if math.Abs(a.Interpolation.Weight-sevenOver72) > 1e-15 || math.Abs(a.Interpolation.ComplementWeight-65.0/72.0) > 1e-15 || math.Abs(a.Interpolation.W72-0.04982659643506822) > 5e-15 || math.Abs(a.Interpolation.Residual-8.525834413464217e-10) > 5e-18 || !a.Interpolation.ActiveTransport || a.Interpolation.FanoHitchinRoute {
		t.Fatalf("bad interpolation: %+v", a.Interpolation)
	}
	if math.Abs(a.FormulaLift.KappaLambdaFromW72Exact-0.044323042243493656) > 5e-15 || math.Abs(a.FormulaLift.KappaLambdaResidual+8.525834413464217e-10) > 5e-18 || math.Abs(a.FormulaLift.RuntimePredictionExactKappaE-0.12965256505471276) > 5e-15 || math.Abs(a.FormulaLift.RuntimeResidualExactKappaE-4.2369718844526005e-12) > 5e-17 || math.Abs(a.FormulaLift.RuntimeResidualOrientKappaE-1.3799595133257014e-08) > 5e-15 || !a.FormulaLift.BridgeLayerOnly {
		t.Fatalf("bad formula lift: %+v", a.FormulaLift)
	}
	if a.Residuals.RawToW72Improvement <= 100000 || a.Residuals.WeightedToBoundarySplit >= 1e-6 || a.Residuals.ExactRuntimeRelative >= 4e-11 || math.Abs(a.Residuals.RuntimeResidualExactKappaE-4.2369718844526005e-12) > 5e-17 {
		t.Fatalf("bad residual hierarchy: %+v", a.Residuals)
	}
	if !a.Classification.SevenOver72ActingAsK7TraceWeight || !a.Classification.SevenOver72ActingAsAugmentedDimension || !a.Classification.SevenOver72ActingAsBoundaryWeight || a.Classification.SevenOver72ActingAsTransportArtifact || a.Classification.SevenOver72UnsourcedEnvironmentalCoeff || a.Classification.FanoHitchinBoundaryMapConstructed || a.Classification.RandomConstantSearch {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if a.Firewalls.ClaimsNativeSevenOver72Theorem || a.Firewalls.ClaimsNativeK7BoundaryMap || a.Firewalls.ClaimsNativeScalarFlavorTransport || a.Firewalls.ClaimsBoundaryStressDerivation || a.Firewalls.ClaimsHiggsPrediction || a.Firewalls.ClaimsScalarStability || a.Firewalls.ClaimsFlavorDerivation || a.Firewalls.ClaimsCKMPMNSDerivation || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsFanoHitchinBoundaryMap || a.Firewalls.ClaimsPhysicalSpacetime || a.Firewalls.Verdict != StatusGate660Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ActiveSevenOverSeventyTwoBoundaryWeightSourceTypeAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate659WeightedClosureInherited, StatusActiveW72InterpolationDefined, StatusNumeratorSevenSourceAudited, StatusDenominator72SourceAudited, StatusBoundaryInterpolationRoleAudited, StatusFormulaLiftComputed, StatusResidualHierarchyAudited, StatusSevenOver72ActiveWeightSupport, StatusNumeratorSevenK7Candidates, StatusDenominator72AugmentedChamberCandidate, StatusW72ScalarRuntimeFormulaSupport, StatusNoNativeSevenOver72SourceTheorem, StatusNoNativeK7BoundaryMap, StatusNoNativeScalarFlavorBoundaryTheorem, StatusNoFanoHitchinBoundaryRevival, StatusNoBoundaryStressDerivation, StatusNoHiggsFlavorGaugeClaim, StatusGate660Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
