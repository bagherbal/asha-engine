package generation2boundaryweighteddeficitclosureaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(i Gate625Inheritance) string {
	return fmt.Sprintf("L=%s kappaE=%s kappaLambda=%s kappaSum=%s absLambda12=%s R3MinusOne=%s proxy=%s runtime=%s gate625ClosureResidual=%s gate625ScalarResidual=%s closureSeal=%t nativeClosure=%t nativeScalarRG=%t nativeFlavor=%t verdict=%q", f64(i.LoopUnit), f64(i.KappaE), f64(i.KappaLambda), f64(i.KappaSum), f64(i.AbsLambda12), f64(i.R3MinusOne), f64(i.LambdaProxyMZ), f64(i.LambdaRuntimeMZ), f64(i.Gate625ClosureResidual), f64(i.Gate625ScalarPredictionResidual), i.Gate625ClosureSealDefined, i.Gate625NativeClosureTheorem, i.Gate625NativeScalarRGMatching, i.Gate625NativeFlavorOrientation, i.Verdict)
}

func FormatBoundarySplit(s BoundarySplitAudit) string {
	return fmt.Sprintf("absLambda12=%s R3MinusOne=%s split=%s residual=%s ratio=%s splitPositive=%t insideSplit=%t boundaryLane=%t verdict=%q", f64(s.AbsLambda12), f64(s.R3MinusOne), f64(s.BoundarySplit), f64(s.KappaSumMinusAbsLambda12), f64(s.ClosureResidualOverSplit), s.SplitPositive, s.ResidualInsideBoundarySplit, s.BoundaryStressLaneInherited, s.Verdict)
}

func FormatWeightCandidate(c BoundaryWeightCandidate) string {
	return fmt.Sprintf("expression=%q value=%s observedRatio=%s ratioResidual=%s absRatioResidual=%s weightedClosure=%s kappaSum=%s closureResidual=%s absClosureResidual=%s relClosureResidual=%s typed=%t native=%t interpretation=%q verdict=%q", c.Expression, f64(c.Value), f64(c.ObservedRatio), f64(c.RatioResidual), f64(c.AbsoluteRatioResidual), f64(c.WeightedClosure), f64(c.KappaSum), f64(c.WeightedClosureResidual), f64(c.AbsoluteClosureResidual), f64(c.RelativeClosureResidual), c.TypedOperands, c.NativeSourceCertified, c.CandidateInterpretation, c.Verdict)
}

func FormatWeightedClosure(c BoundaryWeightedClosureAudit) string {
	return fmt.Sprintf("kappaSum=%s absLambda12=%s R3MinusOne=%s weightBoundary=%s weightScalar=%s split=%s mixture=%s residual=%s absResidual=%s relResidual=%s improvesGate625=%t improvement=%s equation=%q equivalent=%q bridgeOnly=%t verdict=%q", f64(c.KappaSum), f64(c.AbsLambda12), f64(c.R3MinusOne), f64(c.BoundaryWeight), f64(c.ScalarWeight), f64(c.BoundarySplit), f64(c.WeightedMixture), f64(c.Residual), f64(c.AbsoluteResidual), f64(c.RelativeResidual), c.ImprovesGate625, f64(c.ImprovementFactor), c.ClosureEquation, c.EquivalentEquation, c.BridgeOnly, c.Verdict)
}

func FormatScalarFormula(s WeightedScalarFormula) string {
	return fmt.Sprintf("actualKappaLambda=%s kappaE=%s kappaEOrient=%s mixture=%s predExact=%s predOrient=%s residualExact=%s residualOrient=%s exact=%q orient=%q scalarFormula=%q native=%t verdict=%q", f64(s.KappaLambdaActual), f64(s.KappaEExact), f64(s.KappaEOrient), f64(s.WeightedBoundaryMixture), f64(s.PredictedKappaLambdaExact), f64(s.PredictedKappaLambdaOrient), f64(s.KappaLambdaResidualExact), f64(s.KappaLambdaResidualOrient), s.ExactFormula, s.OrientationSubstitutedFormula, s.CombinedScalarFormula, s.NativeScalarFormulaClaimed, s.Verdict)
}

func FormatPredictionRow(r ScalarPredictionRow) string {
	return fmt.Sprintf("name=%q kappaEUsed=%s predicted=%s runtime=%s residual=%s relResidual=%s formula=%q", r.Name, f64(r.KappaEUsed), f64(r.PredictedLambda), f64(r.RuntimeLambda), f64(r.Residual), f64(r.RelativeResidual), r.Formula)
}

func FormatPredictionRows(rows []ScalarPredictionRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPredictionRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatScalarPrediction(p FullScalarPredictionAudit) string {
	return fmt.Sprintf("proxy=%s runtime=%s L=%s mixture=%s bestResidual=%s gate625Residual=%s improvesGate625=%t improvement=%s diagnosticOnly=%t rows=[%s] verdict=%q", f64(p.LambdaProxyMZ), f64(p.LambdaRuntimeMZ), f64(p.LoopUnit), f64(p.WeightedBoundaryMixture), f64(p.BestResidual), f64(p.Gate625PredictionResidual), p.ImprovesGate625Prediction, f64(p.ImprovementFactor), p.DiagnosticOnly, FormatPredictionRows(p.Rows), p.Verdict)
}

func FormatResidualScaleRow(r ResidualScaleRow) string {
	return fmt.Sprintf("name=%q residual=%s relResidual=%s scale=%q meaning=%q", r.Name, f64(r.Residual), f64(r.RelativeResidual), r.Scale, r.Meaning)
}

func FormatResidualScaleRows(rows []ResidualScaleRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatResidualScaleRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatResidualScales(r ResidualScaleComparison) string {
	return fmt.Sprintf("gate625Closure=%s weightedClosure=%s gate625Scalar=%s weightedScalar=%s closureImprovement=%s scalarImprovement=%s rows=[%s] verdict=%q", f64(r.Gate625ClosureResidual), f64(r.BoundaryWeightedResidual), f64(r.Gate625ScalarResidual), f64(r.BoundaryWeightedScalarResidual), f64(r.ClosureImprovementFactor), f64(r.ScalarImprovementFactor), FormatResidualScaleRows(r.Rows), r.Verdict)
}

func FormatSignAndRole(s SignAndRoleAudit) string {
	return fmt.Sprintf("kappaE=%s kappaLambda=%s absLambda12=%s R3MinusOne=%s boundaryWeight=%s scalarWeight=%s flavorRole=%q scalarRole=%q scalarBoundary=%q gaugeBoundary=%q equation=%q native=%t verdict=%q", f64(s.KappaE), f64(s.KappaLambda), f64(s.AbsLambda12), f64(s.R3MinusOne), f64(s.BoundaryWeight), f64(s.ScalarWeight), s.FlavorRole, s.ScalarMatchingRole, s.ScalarBoundaryWoundRole, s.GaugeBoundaryWoundRole, s.StructuralEquation, s.NativeTheoremClaimed, s.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("native7over72=%t nativeTransport=%t nativeWeightedClosure=%t nativeScalarRG=%t nativeFlavor=%t statement=%q verdict=%q", n.NativeSevenOverSeventyTwoSource, n.NativeGaugeScalarFlavorDeficitTransport, n.NativeBoundaryWeightedClosureTheorem, n.NativeScalarRGMatchingTheorem, n.NativeFlavorOrientationTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t scalarStability=%t koide=%t pmnsCkm=%t gaugeUnification=%t nativeWeight=%t nativeTransport=%t endpoint=%t verdict=%q", f.ClaimsHiggsMassDerived, f.ClaimsScalarStability, f.ClaimsKoideDerived, f.ClaimsPMNSCKMDerived, f.ClaimsGaugeUnification, f.ClaimsNativeWeightTheorem, f.ClaimsNativeTransportTheorem, f.ClaimsEndpointDerivation, f.Verdict)
}
