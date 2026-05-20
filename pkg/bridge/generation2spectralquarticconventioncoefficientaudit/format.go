package generation2spectralquarticconventioncoefficientaudit

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
	return fmt.Sprintf("lambda12=%s R3MinusOne=%s xi=%s formal=%q blocker=%q verdict=%q", f64(h.LambdaRuntimeLambda12), f64(h.R3MinusOne), f64(h.XiBoundary), h.FormalCandidate, h.Blocker, h.Verdict)
}
func FormatConvention(c ConventionFamilyRow) string {
	return fmt.Sprintf("factor=%q role=%q canChange=%t certified=%t impact=%q verdict=%q", c.Factor, c.Role, c.CanChangeCLambda, c.Certified, c.Impact, c.Verdict)
}
func FormatConventions(rows []ConventionFamilyRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatConvention(r))
	}
	return strings.Join(parts, " | ")
}
func FormatCandidate(c CandidateFormula) string {
	return fmt.Sprintf("name=%q formula=%q native=%t certified=%t data=%q statement=%q verdict=%q", c.Name, c.Formula, c.Native, c.Certified, strings.Join(c.RequiredData, ", "), c.Statement, c.Verdict)
}
func FormatCandidates(rows []CandidateFormula) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCandidate(r))
	}
	return strings.Join(parts, " | ")
}
func FormatDiagnostic(d BA2Diagnostic) string {
	return fmt.Sprintf("scale=%q a=%s b=%s b/a^2=%s lambdaRuntime=%s cReqRuntime=%s cReqZero=%s cReqNegXi=%s neutrinoIncluded=%t complete=%t statement=%q verdict=%q", d.Scale, f64(d.ATrace), f64(d.BTrace), f64(d.BOverA2), f64(d.LambdaRuntime), f64(d.CLambdaRequiredRuntime), f64(d.CLambdaRequiredZero), f64(d.CLambdaRequiredNegXi), d.NeutrinoIncluded, d.Complete, d.Statement, d.Verdict)
}
func FormatDiagnostics(rows []BA2Diagnostic) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDiagnostic(r))
	}
	return strings.Join(parts, " | ")
}
func FormatSignAudit(s SignAudit) string {
	return fmt.Sprintf("bOverA2Nonnegative=%t positiveCPositive=%t lambda12Negative=%t directPositiveBoundary=%t statement=%q verdict=%q", s.BOverA2NonNegative, s.PositiveCLambdaGivesPositive, s.LambdaRuntimeAtLambda12Negative, s.DirectPositiveBoundaryPossible, s.Statement, s.Verdict)
}
func FormatRGSeparation(r RGTransportSeparation) string {
	return fmt.Sprintf("boundaryInitial=%t runUpLedger=%t spectralQuartic=%t matching=%t loopThreshold=%t statement=%q verdict=%q", r.LambdaRuntimeBoundaryInitial, r.LambdaRuntimeRunUpLedger, r.LambdaRuntimeSpectralQuartic, r.RequiresMatchingTheorem, r.RequiresLoopThresholdControl, r.Statement, r.Verdict)
}
func FormatStressImpact(s StressSealImpact) string {
	return fmt.Sprintf("runtimeShadow=%t canUseCanon=%t canUseBA2=%t lambda=%s R3MinusOne=%s xi=%s statement=%q verdict=%q", s.UsesLambdaRuntimeShadow, s.CanUseLambdaCanon, s.CanUseBA2Directly, f64(s.LambdaRuntime), f64(s.R3MinusOne), f64(s.XiBoundary), s.Statement, s.Verdict)
}
func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("cLambda=%t Kphi=%t LambdaPhi=%t bOverA2=%t runtimeMatch=%t sign=%t VEV=%t statement=%q verdict=%q", n.NativeCLambda, n.NativeKPhi, n.NativeLambdaPhi, n.NativeBA2Theorem, n.NativeRuntimeMatch, n.NativeSignConvention, n.NativeVEV, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t stability=%t lambdaZero=%t unification=%t nativeStress=%t nativeCLambda=%t verdict=%q", f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsLambdaZero, f.ClaimsGaugeUnification, f.ClaimsNativeStress, f.ClaimsNativeCLambda, f.Verdict)
}
