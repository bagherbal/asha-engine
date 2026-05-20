package generation2historyloopdeficitclosuretriangleaudit

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

func FormatInherited(i Gate624Inheritance) string {
	return fmt.Sprintf("L=%s kappaE=%s kappaLambda=%s proxy=%s runtime=%s orientKappa=%s flavorResidual=%s quarterPhase=%t nativeHistoryLoop=%t nativeHopfScalar=%t nativeHopfFlavor=%t verdict=%q", f64(i.LoopUnit), f64(i.KappaE), f64(i.KappaLambda), f64(i.LambdaProxyMZ), f64(i.LambdaRuntimeMZ), f64(i.FlavorOrientationKappa), f64(i.FlavorOrientationResidual), i.Gate624QuarterPhase, i.NativeHistoryLoopUnit, i.NativeHopfToScalar, i.NativeHopfToFlavor, i.Verdict)
}

func FormatKappas(k KappaDefinitions) string {
	return fmt.Sprintf("kappaE=%s kappaLambda=%s L=%s bothPositive=%t scalarDeficitLarger=%t verdict=%q", f64(k.KappaE), f64(k.KappaLambda), f64(k.LoopUnit), k.BothPositive, k.ScalarDeficitLarger, k.Verdict)
}

func FormatClosureRow(r DeficitClosureRow) string {
	return fmt.Sprintf("target=%q targetValue=%s kappaSum=%s residual=%s absResidual=%s relResidual=%s typed=%t native=%t comment=%q", r.Target, f64(r.TargetValue), f64(r.KappaSum), f64(r.Residual), f64(r.AbsoluteResidual), f64(r.RelativeResidual), r.Typed, r.NativeCertified, r.Comment)
}

func FormatClosureRows(rows []DeficitClosureRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatClosureRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatClosureTable(c DeficitClosureTable) string {
	return fmt.Sprintf("kappaE=%s kappaLambda=%s kappaSum=%s closest=%q closestResidual=%s closestRelative=%s closesOnAbsLambda=%t rows=[%s] verdict=%q", f64(c.KappaE), f64(c.KappaLambda), f64(c.KappaSum), c.ClosestTarget, f64(c.ClosestResidual), f64(c.ClosestRelative), c.ClosesOnAbsLambda, FormatClosureRows(c.Rows), c.Verdict)
}

func FormatScalarFormula(s ScalarDeficitFormula) string {
	return fmt.Sprintf("actualKappaLambda=%s absLambda12=%s kappaE=%s kappaEOrient=%s predExact=%s predOrient=%s residualExact=%s residualOrient=%s exact=%q orient=%q bridgeOnly=%t verdict=%q", f64(s.KappaLambdaActual), f64(s.AbsLambda12), f64(s.KappaEExact), f64(s.KappaEOrient), f64(s.PredictedKappaLambdaExact), f64(s.PredictedKappaLambdaOrient), f64(s.ResidualExact), f64(s.ResidualOrient), s.ExactFormula, s.OrientationSubstitutedForm, s.BridgeOnly, s.Verdict)
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
	return fmt.Sprintf("proxy=%s runtime=%s L=%s absLambda12=%s bestResidual=%s improvesRaw=%t diagnosticOnly=%t rows=[%s] verdict=%q", f64(p.LambdaProxyMZ), f64(p.LambdaRuntimeMZ), f64(p.LoopUnit), f64(p.AbsLambda12), f64(p.BestResidual), p.ImprovesGate623RawLAnsatz, p.DiagnosticOnly, FormatPredictionRows(p.Rows), p.Verdict)
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
	return fmt.Sprintf("rawScalarResidual=%s closureScalarResidual=%s scalarImprovement=%s sharper=%t closureDimensionless=%s rows=[%s] verdict=%q", f64(r.Gate623ScalarAnsatzResidual), f64(r.ClosureScalarResidual), f64(r.ScalarImprovementFactor), r.ClosureSharperThanRawScalarAnsatz, f64(r.ClosureResidualDimensionless), FormatResidualScaleRows(r.Rows), r.Verdict)
}

func FormatSignAndRole(s SignAndRoleAudit) string {
	return fmt.Sprintf("kappaE=%s kappaLambda=%s absLambda12=%s flavorRole=%q scalarRole=%q highScaleRole=%q equation=%q opposedRGWound=%t native=%t verdict=%q", f64(s.KappaE), f64(s.KappaLambda), f64(s.AbsLambda12), s.FlavorDeficitRole, s.ScalarMatchingDeficitRole, s.HighScaleScalarWoundRole, s.StructuralEquation, s.OpposedRGWoundSign, s.NativeTheoremClaimed, s.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("nativeKappaClosure=%t nativeScalarRGMatching=%t nativeFlavorOrientation=%t nativeLowScaleToHighScale=%t nativeHistoryLoopDeficitClosure=%t statement=%q verdict=%q", n.NativeKappaClosureTheorem, n.NativeScalarRGMatchingTheorem, n.NativeFlavorOrientationTheorem, n.NativeLowScaleMatchingToHighScaleWoundLaw, n.NativeHistoryLoopDeficitClosureTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t scalarStability=%t koide=%t pmnsCkm=%t gaugeUnification=%t nativeClosure=%t nativeHistoryLoopUnit=%t verdict=%q", f.ClaimsHiggsMassDerived, f.ClaimsScalarStability, f.ClaimsKoideDerived, f.ClaimsPMNSCKMDerived, f.ClaimsGaugeUnification, f.ClaimsNativeASHAClosure, f.ClaimsNativeHistoryLoopUnit, f.Verdict)
}
