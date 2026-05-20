package generation2scalaroneeighthproxyloopmatchingaudit

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
	return fmt.Sprintf("proxy=%s runtime=%s delta=%s relProxy=%s relRuntime=%s ba2=%s mProxy=%s mRuntime=%s gapVerdict=%q chainVerdict=%q verdict=%q", f64(h.LambdaProxyMZ), f64(h.LambdaRuntimeMZ), f64(h.DeltaLambdaMatch), f64(h.RelativeToProxy), f64(h.RelativeToRuntime), f64(h.BA2MZ), f64(h.MProxyGeV), f64(h.MRuntimeGeV), h.PreviousGapVerdict, h.PreviousChainVerdict, h.Verdict)
}
func FormatOneEighth(o OneEighthProxyAudit) string {
	return fmt.Sprintf("proxy=%s runtime=%s oneEighth=%s proxyMinus=%s runtimeMinus=%s proxyRel=%s runtimeRel=%s native=%t statement=%q verdict=%q", f64(o.LambdaProxy), f64(o.LambdaRuntime), f64(o.OneEighth), f64(o.ProxyMinusOneEighth), f64(o.RuntimeMinusOneEighth), f64(o.ProxyRelativeDeviation), f64(o.RuntimeRelativeDeviation), o.NativeClaim, o.ProxyFromBA2, o.Verdict)
}
func FormatLoopCandidate(c LoopCorrectionCandidate) string {
	return fmt.Sprintf("name=%q value=%s residual=%s relResidual=%s typed=%t native=%t comment=%q", c.Name, f64(c.Value), f64(c.Residual), f64(c.RelativeResidual), c.Typed, c.NativeCertified, c.Comment)
}
func FormatCandidates(rows []LoopCorrectionCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatLoopCandidate(r))
	}
	return strings.Join(parts, " | ")
}
func FormatRelativeLoops(r RelativeLoopCorrectionTable) string {
	return fmt.Sprintf("rho=%s closest=%q closestResidual=%s candidates=[%s] verdict=%q", f64(r.RhoLambdaMatch), r.ClosestName, f64(r.ClosestResidual), FormatCandidates(r.Candidates), r.Verdict)
}
func FormatAbsoluteLoops(a AbsoluteLoopCorrectionTable) string {
	return fmt.Sprintf("delta=%s closest=%q closestResidual=%s candidates=[%s] verdict=%q", f64(a.DeltaLambdaMatch), a.ClosestName, f64(a.ClosestResidual), FormatCandidates(a.Candidates), a.Verdict)
}
func FormatSign(s SignAudit) string {
	return fmt.Sprintf("proxyBelow=%t positive=%t slots=%q native=%t statement=%q verdict=%q", s.LambdaProxyBelowRuntime, s.PositiveCorrection, strings.Join(s.SignCompatibleSlots, ", "), s.NativeSourceCertified, s.Statement, s.Verdict)
}
func FormatHiggsDiagnostic(h HiggsProxyRefinementDiagnostic) string {
	return fmt.Sprintf("loop=%s proxy=%s ansatz=%s runtime=%s ansatzMinusRuntime=%s rel=%s massProxy=%s massAnsatz=%s massRuntime=%s deltaMass=%s claims=%t statement=%q verdict=%q", f64(h.LoopUnitOneOver8Pi), f64(h.LambdaProxy), f64(h.LambdaAnsatz), f64(h.LambdaRuntime), f64(h.AnsatzMinusRuntime), f64(h.RelativeAnsatzResidual), f64(h.MassProxyGeV), f64(h.MassAnsatzGeV), f64(h.MassRuntimeGeV), f64(h.DeltaMassAnsatzRuntimeGeV), h.ClaimsHiggsPrediction, h.Statement, h.Verdict)
}
func FormatRuntimeChain(r RuntimeTransportChain) string {
	return fmt.Sprintf("proxy=%s delta=%s runtimeMZ=%s runtimeL12=%s native=%t statement=%q verdict=%q", f64(r.LambdaProxy), f64(r.DeltaLambdaLoopSlot), f64(r.LambdaRuntimeMZ), f64(r.LambdaRuntimeL12), r.ChainNative, r.Statement, r.Verdict)
}
func FormatNativeStatus(n NativeStatus) string {
	return fmt.Sprintf("ba2=%t c3over8=%t one8pi=%t proxyRuntime=%t higgsPole=%t statement=%q verdict=%q", n.NativeBA2OneThirdTheorem, n.NativeCThreeEighthsTheorem, n.NativeOneOver8PiScalarMatching, n.NativeProxyRuntimeTheorem, n.NativeHiggsPoleTheorem, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("higgsMass=%t stability=%t lambdaZero=%t unification=%t nativeScalar=%t verdict=%q", f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsLambdaZeroBoundary, f.ClaimsGaugeUnification, f.ClaimsNativeScalarTheorem, f.Verdict)
}
