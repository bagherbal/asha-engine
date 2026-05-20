package generation2spectralquarticconventioncoefficientaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SpectralQuarticConventionCoefficientAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spectral quartic convention coefficient c_lambda audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate619 c_lambda audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate618 c_lambda blocker", Passed: a.Inherited.Verdict == StatusGate618Inherited && a.Inherited.LambdaRuntimeLambda12 < 0 && a.Inherited.FormalCandidate != "", Detail: FormatInherited(a.Inherited)},
			{Name: "classify convention family", Passed: len(a.Conventions) >= 6 && anyConventionCanChange(a.Conventions) && !allConventionsCertified(a.Conventions), Detail: FormatConventions(a.Conventions)},
			{Name: "define formal c_lambda target", Passed: len(a.CandidateFormulas) >= 1 && !a.CandidateFormulas[0].Certified && a.CandidateFormulas[0].Formula != "", Detail: FormatCandidates(a.CandidateFormulas)},
			{Name: "compute runtime b/a^2 diagnostics", Passed: len(a.Diagnostics) == 2 && a.Diagnostics[1].BOverA2 > 0 && a.Diagnostics[1].CLambdaRequiredRuntime < 0 && !a.Diagnostics[1].Complete, Detail: FormatDiagnostics(a.Diagnostics)},
			{Name: "audit b/a^2 sign", Passed: a.SignAudit.BOverA2NonNegative && a.SignAudit.LambdaRuntimeAtLambda12Negative && !a.SignAudit.DirectPositiveBoundaryPossible, Detail: FormatSignAudit(a.SignAudit)},
			{Name: "separate runtime transport from spectral quartic", Passed: a.RGSeparation.LambdaRuntimeRunUpLedger && !a.RGSeparation.LambdaRuntimeSpectralQuartic && a.RGSeparation.RequiresMatchingTheorem, Detail: FormatRGSeparation(a.RGSeparation)},
			{Name: "keep stress seal scalar side as runtime shadow", Passed: a.StressImpact.UsesLambdaRuntimeShadow && !a.StressImpact.CanUseLambdaCanon && !a.StressImpact.CanUseBA2Directly, Detail: FormatStressImpact(a.StressImpact)},
			{Name: "record missing native scalar convention theorem", Passed: !a.NativeStatus.NativeCLambda && !a.NativeStatus.NativeKPhi && !a.NativeStatus.NativeLambdaPhi && !a.NativeStatus.NativeBA2Theorem && !a.NativeStatus.NativeRuntimeMatch && !a.NativeStatus.NativeVEV, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve c_lambda firewalls", Passed: !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsLambdaZero && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsNativeStress && !a.Firewalls.ClaimsNativeCLambda, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func anyConventionCanChange(rows []ConventionFamilyRow) bool {
	for _, r := range rows {
		if r.CanChangeCLambda {
			return true
		}
	}
	return false
}
func allConventionsCertified(rows []ConventionFamilyRow) bool {
	for _, r := range rows {
		if !r.Certified {
			return false
		}
	}
	return true
}
