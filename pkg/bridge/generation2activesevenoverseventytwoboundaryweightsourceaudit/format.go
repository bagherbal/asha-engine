package generation2activesevenoverseventytwoboundaryweightsourceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate659Inheritance) string {
	return fmt.Sprintf("weightedInherited=%t activeTransport=%t fanoSealed=%t Ksum=%.15g W72=%.15g weightedResidual=%.15g rawResidual=%.15g split=%.15g no7=%t noKappa=%t noTransport=%t firewall=%t verdict=%q", x.BoundaryWeightedClosureInherited, x.ActiveTransportLane, x.FanoHitchinRouteSealed, x.KSum, x.W72, x.WeightedResidual, x.RawClosureResidual, x.BoundarySplit, x.NoNativeSevenOver72Theorem, x.NoNativeKappaClosureTheorem, x.NoNativeTransportTheorem, x.FirewallPreserved, x.Verdict)
}

func FormatCandidate(x SourceCandidate) string {
	return fmt.Sprintf("%s=%.15g role=%q status=%q native=%t active=%t", x.Name, x.Value, x.Role, x.Status, x.NativeTheorem, x.ActiveInTransport)
}

func FormatNumerator(x NumeratorSevenAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("k7=%d fanoStrengthens=%t ker7=%t coker7=%t boundaryMap=%t verdict=%q candidates=[%s]", x.K7CarrierDimension, x.FanoHitchinStrengthens7, x.IntersectionDefect7, x.CokernelDefect7, x.BoundaryMapConstructed, x.Verdict, strings.Join(parts, "; "))
}

func FormatDenominator(x Denominator72Audit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("lambda4BoundaryDim=%d preferred=%q augmentedTyped=%t boundaryEnv=%t nativeTrace=%t verdict=%q candidates=[%s]", x.Lambda4PlusBoundaryPairDim, x.PreferredCandidate, x.AugmentedChamberTyped, x.BoundaryPairEnvironmental, x.NativeTraceTheorem, x.Verdict, strings.Join(parts, "; "))
}

func FormatInterpolation(x ActiveW72Interpolation) string {
	return fmt.Sprintf("absLambda=%.15g r3=%.15g split=%.15g weight=%.15g complement=%.15g W72=%.15g Ksum=%.15g residual=%.15g formula=%q active=%t fano=%t verdict=%q", x.AbsLambdaLambda12, x.R3Minus1, x.BoundarySplit, x.Weight, x.ComplementWeight, x.W72, x.KSum, x.Residual, x.Formula, x.ActiveTransport, x.FanoHitchinRoute, x.Verdict)
}

func FormatFormulaLift(x ScalarRuntimeFormulaLift) string {
	return fmt.Sprintf("proxy=%.15g runtime=%.15g L=%.15g W72=%.15g kappaE=%.15g kappaEOrient=%.15g kappaFromW=%.15g kappaActual=%.15g kappaResidual=%.15g predExact=%.15g residualExact=%.15g predOrient=%.15g residualOrient=%.15g formulaExact=%q formulaOrient=%q bridgeOnly=%t verdict=%q", x.LambdaProxyMZ, x.LambdaRuntimeMZ, x.L, x.W72, x.KappaEExact, x.KappaEOrientation, x.KappaLambdaFromW72Exact, x.KappaLambdaActual, x.KappaLambdaResidual, x.RuntimePredictionExactKappaE, x.RuntimeResidualExactKappaE, x.RuntimePredictionOrientKappaE, x.RuntimeResidualOrientKappaE, x.FormulaExact, x.FormulaOrientation, x.BridgeLayerOnly, x.Verdict)
}

func FormatResiduals(x ResidualHierarchy) string {
	return fmt.Sprintf("rawClosure=%.15g weighted=%.15g split=%.15g runtimeExact=%.15g runtimeOrient=%.15g improvement=%.15g weighted/split=%.15g exactRuntimeRel=%.15g verdict=%q", x.RawKappaClosureResidual, x.W72WeightedResidual, x.BoundarySplit, x.RuntimeResidualExactKappaE, x.RuntimeResidualOrientKappaE, x.RawToW72Improvement, x.WeightedToBoundarySplit, x.ExactRuntimeRelative, x.Verdict)
}

func FormatClassification(x SourceTypeClassification) string {
	return fmt.Sprintf("traceWeight=%t augmented=%t boundaryWeight=%t artifact=%t unsourced=%t fanoMap=%t random=%t verdict=%q", x.SevenOver72ActingAsK7TraceWeight, x.SevenOver72ActingAsAugmentedDimension, x.SevenOver72ActingAsBoundaryWeight, x.SevenOver72ActingAsTransportArtifact, x.SevenOver72UnsourcedEnvironmentalCoeff, x.FanoHitchinBoundaryMapConstructed, x.RandomConstantSearch, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("seven72=%t k7Map=%t transport=%t boundary=%t higgs=%t stability=%t flavor=%t ckm=%t gauge=%t fano=%t spacetime=%t verdict=%q", x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativeK7BoundaryMap, x.ClaimsNativeScalarFlavorTransport, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.ClaimsGaugeUnification, x.ClaimsFanoHitchinBoundaryMap, x.ClaimsPhysicalSpacetime, x.Verdict)
}
