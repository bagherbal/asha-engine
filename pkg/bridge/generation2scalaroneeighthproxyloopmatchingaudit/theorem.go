package generation2scalaroneeighthproxyloopmatchingaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarOneEighthProxyLoopMatchingAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 scalar one-eighth proxy and loop-matching correction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate622 scalar loop matching audit", Passed: false, Detail: err.Error()}}}
		}
		oneOver8Pi := 1 / (8 * math.Pi)
		oneOver64Pi := 1 / (64 * math.Pi)
		checks := []theorem.Check{
			{Name: "inherit Gate621 matching gap", Passed: a.Inherited.Verdict == StatusGate621Inherited && a.Inherited.DeltaLambdaMatch > 0 && a.Inherited.RelativeToProxy > 0.037, Detail: FormatInherited(a.Inherited)},
			{Name: "audit one-eighth proxy", Passed: a.OneEighth.Verdict == StatusOneEighthProxyAudited && math.Abs(a.OneEighth.ProxyMinusOneEighth) < 1.0e-4 && !a.OneEighth.NativeClaim, Detail: FormatOneEighth(a.OneEighth)},
			{Name: "compute relative loop correction", Passed: a.RelativeLoops.Verdict == StatusRelativeLoopComputed && math.Abs(a.RelativeLoops.RhoLambdaMatch-oneOver8Pi)/oneOver8Pi < 0.05, Detail: FormatRelativeLoops(a.RelativeLoops)},
			{Name: "compute absolute loop correction", Passed: a.AbsoluteLoops.Verdict == StatusAbsoluteLoopComputed && math.Abs(a.AbsoluteLoops.DeltaLambdaMatch-oneOver64Pi)/oneOver64Pi < 0.05, Detail: FormatAbsoluteLoops(a.AbsoluteLoops)},
			{Name: "require positive loop-sized correction", Passed: a.Sign.PositiveCorrection && a.Sign.LambdaProxyBelowRuntime && !a.Sign.NativeSourceCertified, Detail: FormatSign(a.Sign)},
			{Name: "compute refined loop proxy diagnostic", Passed: a.HiggsDiagnostic.Verdict == StatusRefinedLoopProxyComputed && a.HiggsDiagnostic.LambdaAnsatz > a.HiggsDiagnostic.LambdaProxy && math.Abs(a.HiggsDiagnostic.AnsatzMinusRuntime) < 3e-4 && !a.HiggsDiagnostic.ClaimsHiggsPrediction, Detail: FormatHiggsDiagnostic(a.HiggsDiagnostic)},
			{Name: "define scalar proxy-loop-runtime chain", Passed: a.RuntimeChain.Verdict == StatusRuntimeChainDefined && a.RuntimeChain.LambdaRuntimeL12 < 0 && !a.RuntimeChain.ChainNative, Detail: FormatRuntimeChain(a.RuntimeChain)},
			{Name: "record missing native scalar theorems", Passed: !a.NativeStatus.NativeBA2OneThirdTheorem && !a.NativeStatus.NativeCThreeEighthsTheorem && !a.NativeStatus.NativeOneOver8PiScalarMatching && !a.NativeStatus.NativeProxyRuntimeTheorem && !a.NativeStatus.NativeHiggsPoleTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate622 firewalls", Passed: !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsLambdaZeroBoundary && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeScalarTheorem, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Relative loop clue: rho_lambda_match is "+strings.TrimSpace(FormatRelativeLoops(a.RelativeLoops)))
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
