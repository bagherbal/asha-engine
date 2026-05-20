package generation2spectralactionabf0canonicalscalarquarticairlockaudit

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
	return fmt.Sprintf("Lambda12=%s lambdaRuntime=%s R3MinusOne=%s xi=%s verdict=%q", f64(h.Lambda12GeV), f64(h.LambdaRuntime), f64(h.R3MinusOne), f64(h.XiBoundary), h.Verdict)
}
func FormatABTrace(r ABTraceStatus) string {
	return fmt.Sprintf("symbol=%q def=%q nativeForm=%t observedValues=%t bridgeSealed=%t role=%q obstruction=%q verdict=%q", r.Symbol, r.FormalDefinition, r.NativeForm, r.ObservedValues, r.BridgeSealed, r.Role, r.Obstruction, r.Verdict)
}
func FormatABTraces(rows []ABTraceStatus) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatABTrace(r))
	}
	return strings.Join(parts, " | ")
}
func FormatKineticAudit(a ScalarKineticAudit) string {
	return fmt.Sprintf("coefficient=%q dependency=%q KphiNative=%t f0Value=%t aNative=%t formulaCertified=%t statement=%q verdict=%q", a.Coefficient, a.CandidateDependency, a.KPhiNative, a.F0NativeValue, a.ATraceNativeForm, a.CertifiedFormula, a.Statement, a.Verdict)
}
func FormatQuarticAudit(a ScalarQuarticAudit) string {
	return fmt.Sprintf("coefficient=%q dependency=%q LambdaPhiNative=%t bNative=%t f0Value=%t formulaCertified=%t statement=%q verdict=%q", a.Coefficient, a.CandidateDependency, a.LambdaPhiNative, a.BTraceNativeForm, a.F0NativeValue, a.CertifiedFormula, a.Statement, a.Verdict)
}
func FormatRatioAudit(a CanonicalRatioAudit) string {
	return fmt.Sprintf("candidate=%q cCertified=%t requiresKPhi=%t requiresLambdaPhi=%t requiresConvention=%t data=%q statement=%q verdict=%q", a.CandidateFormula, a.CLambdaCertified, a.RequiresKPhi, a.RequiresLambdaPhi, a.RequiresConvention, strings.Join(a.RequiredData, ", "), a.Statement, a.Verdict)
}
func FormatConvention(c ConventionDependency) string {
	return fmt.Sprintf("name=%q required=%t certified=%t impact=%q verdict=%q", c.Name, c.Required, c.Certified, c.Impact, c.Verdict)
}
func FormatConventions(rows []ConventionDependency) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatConvention(r))
	}
	return strings.Join(parts, " | ")
}
func FormatRuntimeConnection(r RuntimeTransportConnection) string {
	return fmt.Sprintf("lambdaMZCanonical=%t lambda12V1=%t matching=%t equalCanon=%t top=%t alphaS=%t thresholds=%t loop=%t statement=%q verdict=%q", r.LambdaMZCanonical, r.LambdaLambda12V1, r.MatchingTheorem, r.EquivalentToLambdaCanon, r.TopMassSensitive, r.AlphaSSensitive, r.ThresholdSensitive, r.LoopOrderSensitive, r.Statement, r.Verdict)
}
func FormatStressImpact(s StressSealImpact) string {
	return fmt.Sprintf("lambdaRuntime=%s xi=%s liftCanon=%t fixCLambda=%t status=%q statement=%q verdict=%q", f64(s.LambdaRuntime), f64(s.XiBoundary), s.CanLiftToLambdaCanon, s.CanNumericallyFixCLambda, s.ScalarSideStatus, s.Statement, s.Verdict)
}
func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("Kphi=%t LambdaPhi=%t cLambda=%t abf0=%t runtimeMatch=%t VEV=%t stress=%t statement=%q verdict=%q", n.NativeKPhi, n.NativeLambdaPhi, n.NativeCLambda, n.NativeABF0ToLambda, n.NativeRuntimeMatch, n.NativeVEV, n.NativeStress, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t stability=%t lambdaZero=%t unification=%t stress=%t verdict=%q", f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsLambdaZero, f.ClaimsGaugeUnification, f.ClaimsNativeStress, f.Verdict)
}
