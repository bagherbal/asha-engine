package generation2scalarcanonicalnormalizationspectralquarticairlockaudit

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

func FormatInherited(h Inherited) string {
	return fmt.Sprintf("Lambda12=%s lambda12=%s xi=%s R3MinusOne=%s deltaLambda=%s verdict=%q", f64(h.Lambda12GeV), f64(h.LambdaLambda12), f64(h.XiBoundary), f64(h.R3MinusOne), f64(h.DeltaLambdaBridge), h.Verdict)
}

func FormatScalarCoefficient(c ScalarCoefficient) string {
	return fmt.Sprintf("symbol=%q meaning=%q layer=%q status=%q native=%t bridge=%t observed=%t missing=%q verdict=%q", c.Symbol, c.Meaning, c.Layer, c.Status, c.Native, c.Bridge, c.Observed, c.MissingData, c.Verdict)
}
func FormatScalarCoefficients(rows []ScalarCoefficient) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatScalarCoefficient(r))
	}
	return strings.Join(parts, " | ")
}

func FormatRuntimeConvention(r RuntimeLambdaConventionAudit) string {
	return fmt.Sprintf("formula=%q convention=%q canonicalSM=%t bridgeRuntime=%t requiresMatching=%t statement=%q verdict=%q", r.RuntimeFormula, r.PotentialConvention, r.CanonicalSMConvention, r.BridgeRuntime, r.RequiresMatching, r.Statement, r.Verdict)
}

func FormatCanonicalMap(m CanonicalNormalizationMap) string {
	return fmt.Sprintf("kinetic=%q quartic=%q field=%q lambda=%q native=%t dependency=%q statement=%q verdict=%q", m.PreCanonicalKinetic, m.PreCanonicalQuartic, m.CanonicalField, m.CanonicalQuartic, m.ExactFormulaNative, m.Dependency, m.Statement, m.Verdict)
}

func FormatABF0Audit(a SpectralActionABF0Audit) string {
	return fmt.Sprintf("shape=%q aNative=%t bNative=%t f0Native=%t KphiNative=%t formulaCertified=%t conventionCertified=%t statement=%q verdict=%q", a.CandidateShape, a.AAvailableNative, a.BAvailableNative, a.F0AvailableNative, a.KPhiAvailableNative, a.FormulaCertified, a.ConventionCertified, a.Statement, a.Verdict)
}

func FormatRuntimeAirlock(a RuntimeToBoundaryAirlockStatus) string {
	return fmt.Sprintf("lambda12=%s canonicalRuntime=%t v1=%t top=%t alphaS=%t thresholds=%t loopOrder=%t preCanonical=%t statement=%q verdict=%q", f64(a.LambdaRuntimeValue), a.IsCanonicalRuntimeLedger, a.IsV1Transported, a.TopMassSensitive, a.AlphaSSensitive, a.ThresholdSensitive, a.LoopOrderSensitive, a.EquivalentToPreCanonical, a.Statement, a.Verdict)
}

func FormatStressImpact(s StressSealImpactAssessment) string {
	return fmt.Sprintf("scalarShadow=%s xi=%s stressResidual=%s canCanon=%t canPreCanon=%t status=%q statement=%q verdict=%q", f64(s.OriginalScalarShadow), f64(s.XiBoundary), f64(s.RuntimeStressResidual), s.CanReplaceByLambdaCanon, s.CanReplaceByLambdaPhi, s.ScalarSideStatus, s.Statement, s.Verdict)
}

func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("Kphi=%t metric=%t LambdaPhi=%t abf0=%t VEV=%t matching=%t lambdaZero=%t stress=%t statement=%q verdict=%q", n.NativeKPhi, n.NativeScalarMetric, n.NativeLambdaPhi, n.NativeABF0ToLambda, n.NativeVEV, n.NativeMatching, n.NativeLambdaZero, n.NativeStressTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("lambdaZero=%t stability=%t poleMass=%t quartic=%t stress=%t unification=%t verdict=%q", f.ClaimsLambdaZero, f.ClaimsHiggsStability, f.ClaimsHiggsPoleMass, f.ClaimsNativeQuartic, f.ClaimsNativeStressSeal, f.ClaimsGaugeUnification, f.Verdict)
}
