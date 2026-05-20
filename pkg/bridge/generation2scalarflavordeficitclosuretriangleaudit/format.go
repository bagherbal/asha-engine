package generation2scalarflavordeficitclosuretriangleaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate658Inheritance) string {
	return fmt.Sprintf("spine=%t active=%t matching=%t boundary=%t kappa=%t noProxy=%t noRG=%t noBoundary=%t noHiggs=%t firewall=%t verdict=%q", x.ScalarTransportSpineInherited, x.ScalarBoundarySpineActive, x.LowScaleMatchingActive, x.BoundaryStressTransportActive, x.KappaLambdaDefined, x.NoNativeProxyRuntimeTheorem, x.NoNativeRGThresholdTheorem, x.NoNativeBoundaryStressTheorem, x.NoHiggsMassOrStabilityClaim, x.FirewallPreserved, x.Verdict)
}

func FormatFlavor(x FlavorKappaESeal) string {
	return fmt.Sprintf("kappa_e=%.15g source=%q orientation=%q environmental=%t native=%t verdict=%q", x.KappaE, x.Source, x.OrientationBalanceExpression, x.EnvironmentalSeal, x.NativeFlavorTheorem, x.Verdict)
}

func FormatClosure(x ClosureTriangle) string {
	return fmt.Sprintf("kappa_lambda=%.15g kappa_e=%.15g K_sum=%.15g absLambda12=%.15g delta=%.15g relAbs=%.15g relK=%.15g relXi=%.15g closes=%t small=%t target=%t verdict=%q", x.KappaLambda, x.KappaE, x.KSum, x.AbsLambdaLambda12, x.DeltaClosure, x.RelativeToAbsLambda, x.RelativeToKSum, x.RelativeToXiBoundary, x.ClosesOnScalarWound, x.RawClosureResidualSmall, x.HighScaleScalarWoundTarget, x.Verdict)
}

func FormatBoundaryWeight(x BoundaryWeightAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.15g diff=%.15g abs=%.15g status=%q best=%t certByFit=%t", c.Name, c.Value, c.Difference, c.AbsDifference, c.TypedStatus, c.BestAmongTypedSet, c.CanBeCertifiedByFit))
	}
	return fmt.Sprintf("split=%.15g delta=%.15g observedW=%.15g closest=%q closestDiff=%.15g sevenClosest=%t typedOnly=%t verdict=%q candidates=[%s]", x.BoundarySplit, x.DeltaClosure, x.ObservedWeight, x.ClosestCandidate, x.ClosestDifference, x.SevenOver72Closest, x.UsedTypedSetOnly, x.Verdict, strings.Join(parts, "; "))
}

func FormatInterpolation(x SevenOver72Interpolation) string {
	return fmt.Sprintf("absLambda=%.15g r3=%.15g weight=%.15g complement=%.15g target=%.15g Ksum=%.15g rawResidual=%.15g weightedResidual=%.15g improvement=%.15g residualRelK=%.15g formula=%q bridgeOnly=%t verdict=%q", x.AbsLambdaLambda12, x.R3Minus1, x.Weight, x.ComplementWeight, x.WeightedTarget, x.KSum, x.RawClosureResidual, x.WeightedResidual, x.ImprovementFactor, x.ResidualRelativeToKSum, x.Formula, x.BridgeLayerOnly, x.Verdict)
}

func FormatSources(x SourceTypeAudit) string {
	parts := make([]string, 0, len(x.Objects))
	for _, o := range x.Objects {
		parts = append(parts, fmt.Sprintf("%s=%.15g role=%q status=%q native=%t", o.Name, o.Value, o.Role, o.TypedStatus, o.NativeSource))
	}
	return fmt.Sprintf("fano7=%t transport7=%t fanoMap=%t random=%t typedOnly=%t verdict=%q objects=[%s]", x.SevenOver72InFanoLane, x.SevenOver72InTransportLane, x.FanoBoundaryMapConstructed, x.RandomConstantsSearched, x.TypedCandidatesOnly, x.Verdict, strings.Join(parts, "; "))
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("flavor=%t scalar=%t seven72=%t boundary=%t higgs=%t gauge=%t ckm=%t spacetime=%t closure=%t verdict=%q", x.ClaimsNativeFlavorTheorem, x.ClaimsNativeScalarTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.ClaimsPhysicalSpacetime, x.ClaimsNativeClosureTheorem, x.Verdict)
}
