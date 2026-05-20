package generation2scalarproxyruntimematchinggapaudit

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
	return fmt.Sprintf("proxyMZ=%s lambdaMZ=%s proxyL12=%s lambdaL12=%s ba2MZ=%s ba2L12=%s separation=%q stress=%q proxyClose=%q highScale=%q verdict=%q", f64(h.LambdaProxyMZ), f64(h.LambdaRuntimeMZ), f64(h.LambdaProxyLambda12), f64(h.LambdaRuntimeLambda12), f64(h.BA2MZ), f64(h.BA2Lambda12), h.PreviousSeparation, h.PreviousStressVerdict, h.PreviousProxyClose, h.PreviousHighScaleBlock, h.Verdict)
}
func FormatMatchingGap(m MatchingGapTable) string {
	return fmt.Sprintf("scale=%q proxy=%s runtime=%s delta=%s relProxy=%s relRuntime=%s positive=%t claimsCorrection=%t verdict=%q", m.Scale, f64(m.LambdaProxy), f64(m.LambdaRuntime), f64(m.DeltaLambdaMatch), f64(m.RelativeToProxy), f64(m.RelativeToRuntime), m.PositiveCorrectionRequired, m.BridgeMatchingCorrectionClaim, m.Verdict)
}
func FormatEffectiveCLambda(c EffectiveCLambdaCorrection) string {
	return fmt.Sprintf("ba2=%s cProxy=%s cNeeded=%s deltaC=%s relDeltaC=%s statement=%q verdict=%q", f64(c.BOverA2), f64(c.CProxy), f64(c.CNeededMZ), f64(c.DeltaC), f64(c.RelativeDeltaC), c.Statement, c.Verdict)
}
func FormatHiggsProxyGap(h HiggsProxyGapDiagnostic) string {
	return fmt.Sprintf("v=%s lambdaProxy=%s lambdaRuntime=%s massProxy=%s massRuntime=%s deltaMass=%s relMass=%s claimsHiggs=%t claimsPole=%t statement=%q verdict=%q", f64(h.VRuntimeGeV), f64(h.LambdaProxyMZ), f64(h.LambdaRuntimeMZ), f64(h.MassProxyGeV), f64(h.MassRuntimeGeV), f64(h.DeltaMassRuntimeMinusProxyGeV), f64(h.RelativeMassGap), h.ClaimsHiggsDerivation, h.ClaimsPoleMassTheorem, h.Statement, h.Verdict)
}
func FormatSourceCandidate(s TypedSourceCandidate) string {
	return fmt.Sprintf("name=%q positive=%t observed=%t native=%t comment=%q verdict=%q", s.Name, s.CanHavePositiveSign, s.RequiresObservedInput, s.NativeCertified, s.Comment, s.Verdict)
}
func FormatSourceCandidates(rows []TypedSourceCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSourceCandidate(r))
	}
	return strings.Join(parts, " | ")
}
func FormatSignAudit(s SignAudit) string {
	return fmt.Sprintf("proxyBelow=%t positiveDelta=%t positiveSlots=%t nativePositive=%t statement=%q verdict=%q", s.ProxyBelowRuntime, s.PositiveDeltaLambda, s.PositiveSourcesExistAsSlots, s.NativePositiveCorrection, s.Statement, s.Verdict)
}
func FormatNeutrinoTrace(n NeutrinoTraceCompletionAudit) string {
	return fmt.Sprintf("includesNu=%t couldShift=%t inserted=%t native=%t statement=%q verdict=%q", n.VisibleV1IncludesNeutrinoYukawa, n.CouldShiftAAndB, n.ValuesInserted, n.NativeCompletionTheorem, n.Statement, n.Verdict)
}
func FormatRuntimeChain(r RuntimeTransportChain) string {
	return fmt.Sprintf("proxyMZ=%s deltaMatch=%s runtimeMZ=%s runtimeL12=%s native=%t statement=%q verdict=%q", f64(r.LambdaProxyMZ), f64(r.DeltaLambdaMatch), f64(r.LambdaRuntimeMZ), f64(r.LambdaRuntimeL12), r.ChainClosedByNativeTheorem, r.Statement, r.Verdict)
}
func FormatStressImpact(s StressSealImpact) string {
	return fmt.Sprintf("improves=%t stressRuntimeL12=%t replace=%t proxyL12=%s runtimeL12=%s statement=%q verdict=%q", s.ImprovesScalarLaneArchitecture, s.StressStillUsesLambdaRuntimeL12, s.CanReplaceStressLambdaWithProxy, f64(s.LambdaProxyL12), f64(s.LambdaRuntimeL12), s.Statement, s.Verdict)
}
func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("c3over8=%t matching=%t poleMSbar=%t neutrinoTrace=%t higgsPole=%t proxyRuntime=%t statement=%q verdict=%q", n.NativeCThreeEighthsTheorem, n.NativeMatchingCorrection, n.NativePoleMSbarConversion, n.NativeNeutrinoTraceCompletion, n.NativeHiggsPoleTheorem, n.NativeProxyRuntimeTheorem, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t stability=%t lambdaZero=%t unification=%t nativeScalar=%t nativeMatching=%t verdict=%q", f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsLambdaZeroBoundary, f.ClaimsGaugeUnification, f.ClaimsNativeScalarTheorem, f.ClaimsNativeMatching, f.Verdict)
}
