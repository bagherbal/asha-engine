package generation2scalarcanonicalnormalizationspectralquarticairlockaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ScalarCanonicalNormalizationSpectralQuarticAirlockAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 scalar canonical normalization and spectral quartic airlock audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate617 scalar normalization airlock audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate616 canonical-normalization blocker", Passed: a.Inherited.Verdict == StatusGate616Inherited && a.Inherited.XiBoundary > 0 && a.Inherited.LambdaLambda12 < 0, Detail: FormatInherited(a.Inherited)},
			{Name: "classify scalar coefficient types", Passed: len(a.ScalarCoefficients) >= 8 && hasCoeff(a.ScalarCoefficients, "K_phi") && hasCoeff(a.ScalarCoefficients, "Lambda_phi") && hasCoeff(a.ScalarCoefficients, "lambda_runtime") && hasCoeff(a.ScalarCoefficients, "a") && hasCoeff(a.ScalarCoefficients, "b"), Detail: FormatScalarCoefficients(a.ScalarCoefficients)},
			{Name: "audit runtime lambda convention", Passed: a.RuntimeConvention.CanonicalSMConvention && a.RuntimeConvention.BridgeRuntime && a.RuntimeConvention.RequiresMatching, Detail: FormatRuntimeConvention(a.RuntimeConvention)},
			{Name: "write symbolic canonical normalization map", Passed: !a.CanonicalMap.ExactFormulaNative && a.CanonicalMap.CanonicalField != "" && a.CanonicalMap.CanonicalQuartic != "", Detail: FormatCanonicalMap(a.CanonicalMap)},
			{Name: "audit a,b,f0 spectral-action scalar lane", Passed: a.ABF0Audit.AAvailableNative && a.ABF0Audit.BAvailableNative && !a.ABF0Audit.FormulaCertified && !a.ABF0Audit.KPhiAvailableNative, Detail: FormatABF0Audit(a.ABF0Audit)},
			{Name: "record runtime-to-boundary airlock status", Passed: a.RuntimeAirlock.IsCanonicalRuntimeLedger && a.RuntimeAirlock.IsV1Transported && a.RuntimeAirlock.TopMassSensitive && !a.RuntimeAirlock.EquivalentToPreCanonical, Detail: FormatRuntimeAirlock(a.RuntimeAirlock)},
			{Name: "assess stress seal scalar-side impact", Passed: !a.StressSealImpact.CanReplaceByLambdaCanon && !a.StressSealImpact.CanReplaceByLambdaPhi && a.StressSealImpact.RuntimeStressResidual > 0, Detail: FormatStressImpact(a.StressSealImpact)},
			{Name: "record native scalar airlock status", Passed: !a.NativeStatus.NativeKPhi && !a.NativeStatus.NativeScalarMetric && !a.NativeStatus.NativeLambdaPhi && !a.NativeStatus.NativeABF0ToLambda && !a.NativeStatus.NativeVEV && !a.NativeStatus.NativeMatching && !a.NativeStatus.NativeStressTheorem, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve scalar/gauge firewalls", Passed: !a.Firewalls.ClaimsLambdaZero && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsHiggsPoleMass && !a.Firewalls.ClaimsNativeQuartic && !a.Firewalls.ClaimsNativeStressSeal && !a.Firewalls.ClaimsGaugeUnification, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasCoeff(rows []ScalarCoefficient, symbol string) bool {
	for _, r := range rows {
		if r.Symbol == symbol {
			return true
		}
	}
	return false
}
