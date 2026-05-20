package generation2ba2onethirdrigidityspectralquarticproxyaudit

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
	return fmt.Sprintf("MZ b/a2=%s L12 b/a2=%s lambdaMZ=%s lambdaL12=%s prevSign=%q prevStress=%q verdict=%q", f64(h.MZBA2), f64(h.Lambda12BA2), f64(h.LambdaRuntimeMZ), f64(h.LambdaRuntimeLambda12), h.PreviousSignVerdict, h.PreviousStressVerdict, h.Verdict)
}
func FormatRigidityRow(r BA2RigidityRow) string {
	return fmt.Sprintf("scale=%q a=%s b=%s b/a2=%s delta1/3=%s rel1/3=%s lambdaRuntime=%s verdict=%q", r.Scale, f64(r.ATrace), f64(r.BTrace), f64(r.BOverA2), f64(r.DeltaFromOneThird), f64(r.RelativeToOneThird), f64(r.LambdaRuntime), r.Verdict)
}
func FormatRigidityRows(rows []BA2RigidityRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatRigidityRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatRigiditySummary(s BA2RigiditySummary) string {
	return fmt.Sprintf("drift=%s absDrift=%s relDrift=%s nearlyInvariant=%t statement=%q verdict=%q", f64(s.BA2Drift), f64(s.AbsBA2Drift), f64(s.RelativeDriftToMZ), s.NearlyInvariant, s.Statement, s.Verdict)
}
func FormatTopDominanceRow(r TopDominanceRow) string {
	return fmt.Sprintf("scale=%q yt=%s approxA=%s approxB=%s approxBA2=%s runtimeA=%s runtimeB=%s runtimeBA2=%s relA=%s relB=%s statement=%q verdict=%q", r.Scale, f64(r.YTop), f64(r.ApproxA), f64(r.ApproxB), f64(r.ApproxBA2), f64(r.RuntimeA), f64(r.RuntimeB), f64(r.RuntimeBA2), f64(r.ADeltaRelative), f64(r.BDeltaRelative), r.Statement, r.Verdict)
}
func FormatTopDominance(rows []TopDominanceRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatTopDominanceRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatProxyRow(r CLambdaProxyRow) string {
	return fmt.Sprintf("scale=%q c=%s b/a2=%s lambdaProxy=%s lambdaRuntime=%s proxyMinusRuntime=%s abs=%s relAbs=%s signCompatible=%t verdict=%q", r.Scale, f64(r.CLambdaCandidate), f64(r.BOverA2), f64(r.LambdaProxy), f64(r.LambdaRuntime), f64(r.ProxyMinusRuntime), f64(r.AbsResidual), f64(r.RelativeToRuntimeAbs), r.SignCompatible, r.Verdict)
}
func FormatProxyRows(rows []CLambdaProxyRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatProxyRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatHiggsProxy(h HiggsProxyDiagnostic) string {
	return fmt.Sprintf("v=%s lambdaProxyMZ=%s massProxy=%s lambdaRuntimeMZ=%s runtimeMass=%s massResidual=%s claimsHiggs=%t statement=%q verdict=%q", f64(h.VRuntimeGeV), f64(h.LambdaProxyMZ), f64(h.MassProxyGeV), f64(h.LambdaRuntimeMZ), f64(h.RuntimeMassGeV), f64(h.MassResidualGeV), h.ClaimsHiggsDerivation, h.Statement, h.Verdict)
}
func FormatSeparation(s RuntimeTransportSeparation) string {
	return fmt.Sprintf("proxyPositive=%t runtimeRG=%t proxyEqualsMZ=%t proxyEqualsL12=%t lambdaL12Negative=%t matching=%t statement=%q verdict=%q", s.SpectralTreeProxyPositive, s.RuntimeQuarticRGTransported, s.ProxyEqualsRuntimeAtMZ, s.ProxyEqualsRuntimeAtL12, s.LambdaL12Negative, s.RequiresMatchingTheorem, s.Statement, s.Verdict)
}
func FormatStressImpact(s StressSealImpact) string {
	return fmt.Sprintf("stressRuntime=%t spectralProxy=%t canReplace=%t lambdaRuntimeL12=%s proxyL12=%s statement=%q verdict=%q", s.StressUsesRuntimeLambda, s.SpectralLaneUsesProxy, s.CanReplaceRuntimeWithProxy, f64(s.RuntimeLambdaL12), f64(s.ProxyLambdaL12), s.Statement, s.Verdict)
}
func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("topTheorem=%t ba2Theorem=%t c3over8Theorem=%t proxyRuntime=%t higgsMass=%t stability=%t statement=%q verdict=%q", n.NativeTopDominanceTheorem, n.NativeBA2OneThirdTheorem, n.NativeThreeEighthsTheorem, n.NativeProxyRuntimeMatching, n.NativeHiggsMassTheorem, n.NativeScalarStabilityTheorem, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t stability=%t lambdaZero=%t nativeQuartic=%t proxyRuntime=%t nativeBA2=%t verdict=%q", f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsLambdaZero, f.ClaimsNativeQuartic, f.ClaimsProxyRuntimeMatch, f.ClaimsNativeBA2, f.Verdict)
}
