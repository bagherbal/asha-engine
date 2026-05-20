package generation2ba2onethirdrigidityspectralquarticproxyaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2BA2OneThirdRigiditySpectralQuarticProxyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 b/a squared one-third rigidity and spectral quartic proxy audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate620 b/a2 rigidity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate619 b/a2 diagnostic", Passed: a.Inherited.Verdict == StatusGate619Inherited && a.Inherited.LambdaRuntimeLambda12 < 0 && a.Inherited.MZBA2 > 0 && a.Inherited.Lambda12BA2 > 0, Detail: FormatInherited(a.Inherited)},
			{Name: "b/a2 near one-third at MZ and Lambda12", Passed: len(a.RigidityRows) == 2 && abs(a.RigidityRows[0].DeltaFromOneThird) < 3e-4 && abs(a.RigidityRows[1].DeltaFromOneThird) < 3e-4 && a.RigiditySummary.NearlyInvariant, Detail: FormatRigidityRows(a.RigidityRows) + " || " + FormatRigiditySummary(a.RigiditySummary)},
			{Name: "audit top/color dominance explanation", Passed: len(a.TopDominance) == 2 && close(a.TopDominance[0].ApproxBA2, oneThird, 1e-15) && close(a.TopDominance[1].ApproxBA2, oneThird, 1e-15), Detail: FormatTopDominance(a.TopDominance)},
			{Name: "compute c_lambda=3/8 proxy", Passed: len(a.ProxyRows) == 2 && close(a.ProxyRows[0].CLambdaCandidate, sin2Theta, 1e-15) && a.ProxyRows[0].LambdaProxy > 0 && a.ProxyRows[1].LambdaProxy > 0, Detail: FormatProxyRows(a.ProxyRows)},
			{Name: "lambda proxy close to low-scale runtime lambda", Passed: a.ProxyRows[0].AbsResidual < 0.006 && a.ProxyRows[0].RelativeToRuntimeAbs < 0.04, Detail: FormatProxyRow(a.ProxyRows[0])},
			{Name: "lambda proxy does not equal negative high-scale runtime lambda", Passed: !a.ProxyRows[1].SignCompatible && a.ProxyRows[1].LambdaRuntime < 0 && a.ProxyRows[1].LambdaProxy > 0, Detail: FormatProxyRow(a.ProxyRows[1])},
			{Name: "compute low-scale Higgs proxy diagnostic without derivation", Passed: a.HiggsProxy.MassProxyGeV > 120 && a.HiggsProxy.MassProxyGeV < 126 && !a.HiggsProxy.ClaimsHiggsDerivation, Detail: FormatHiggsProxy(a.HiggsProxy)},
			{Name: "separate spectral tree proxy from runtime RG quartic", Passed: a.Separation.SpectralTreeProxyPositive && a.Separation.RuntimeQuarticRGTransported && !a.Separation.ProxyEqualsRuntimeAtL12 && a.Separation.RequiresMatchingTheorem, Detail: FormatSeparation(a.Separation)},
			{Name: "keep stress seal on runtime lambda", Passed: a.StressImpact.StressUsesRuntimeLambda && a.StressImpact.SpectralLaneUsesProxy && !a.StressImpact.CanReplaceRuntimeWithProxy, Detail: FormatStressImpact(a.StressImpact)},
			{Name: "record missing native theorems", Passed: !a.NativeStatus.NativeTopDominanceTheorem && !a.NativeStatus.NativeBA2OneThirdTheorem && !a.NativeStatus.NativeThreeEighthsTheorem && !a.NativeStatus.NativeProxyRuntimeMatching, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve Gate620 firewalls", Passed: !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsLambdaZero && !a.Firewalls.ClaimsNativeQuartic && !a.Firewalls.ClaimsProxyRuntimeMatch && !a.Firewalls.ClaimsNativeBA2, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
func close(a, b, tol float64) bool { return abs(a-b) <= tol }
