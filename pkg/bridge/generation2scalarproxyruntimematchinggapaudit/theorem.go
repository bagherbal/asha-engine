package generation2scalarproxyruntimematchinggapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ScalarTreeProxyToRuntimeMatchingGapAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 scalar tree-proxy to runtime matching gap audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate621 scalar matching audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate620 proxy/runtime separation", Passed: a.Inherited.Verdict == StatusGate620Inherited && a.Inherited.PreviousSeparation != "" && a.Inherited.LambdaProxyMZ > 0 && a.Inherited.LambdaRuntimeLambda12 < 0, Detail: FormatInherited(a.Inherited)},
			{Name: "compute low-scale matching gap", Passed: a.MatchingGap.Verdict == StatusMatchingGapComputed && a.MatchingGap.DeltaLambdaMatch > 0 && a.MatchingGap.RelativeToProxy < 0.04, Detail: FormatMatchingGap(a.MatchingGap)},
			{Name: "compute effective c_lambda correction", Passed: a.EffectiveCLambda.Verdict == StatusEffectiveCLambdaComputed && a.EffectiveCLambda.CNeededMZ > a.EffectiveCLambda.CProxy && a.EffectiveCLambda.RelativeDeltaC < 0.04, Detail: FormatEffectiveCLambda(a.EffectiveCLambda)},
			{Name: "compute Higgs proxy gap diagnostic without pole claim", Passed: a.HiggsProxyGap.Verdict == StatusHiggsProxyGapComputed && a.HiggsProxyGap.DeltaMassRuntimeMinusProxyGeV > 0 && !a.HiggsProxyGap.ClaimsHiggsDerivation && !a.HiggsProxyGap.ClaimsPoleMassTheorem, Detail: FormatHiggsProxyGap(a.HiggsProxyGap)},
			{Name: "classify typed scalar matching sources", Passed: len(a.SourceCandidates) >= 6, Detail: FormatSourceCandidates(a.SourceCandidates)},
			{Name: "positive low-scale matching correction required", Passed: a.Sign.PositiveDeltaLambda && a.Sign.ProxyBelowRuntime && !a.Sign.NativePositiveCorrection, Detail: FormatSignAudit(a.Sign)},
			{Name: "audit missing neutrino trace completion", Passed: !a.NeutrinoTrace.VisibleV1IncludesNeutrinoYukawa && a.NeutrinoTrace.CouldShiftAAndB && !a.NeutrinoTrace.ValuesInserted && !a.NeutrinoTrace.NativeCompletionTheorem, Detail: FormatNeutrinoTrace(a.NeutrinoTrace)},
			{Name: "define proxy-to-runtime-to-RG chain", Passed: a.RuntimeChain.DeltaLambdaMatch > 0 && a.RuntimeChain.LambdaRuntimeL12 < 0 && !a.RuntimeChain.ChainClosedByNativeTheorem, Detail: FormatRuntimeChain(a.RuntimeChain)},
			{Name: "keep stress seal on high-scale runtime lambda", Passed: a.StressImpact.ImprovesScalarLaneArchitecture && a.StressImpact.StressStillUsesLambdaRuntimeL12 && !a.StressImpact.CanReplaceStressLambdaWithProxy, Detail: FormatStressImpact(a.StressImpact)},
			{Name: "record missing native theorems", Passed: !a.NativeStatus.NativeCThreeEighthsTheorem && !a.NativeStatus.NativeMatchingCorrection && !a.NativeStatus.NativeHiggsPoleTheorem && !a.NativeStatus.NativeProxyRuntimeTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate621 firewalls", Passed: !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsLambdaZeroBoundary && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeScalarTheorem && !a.Firewalls.ClaimsNativeMatching, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
