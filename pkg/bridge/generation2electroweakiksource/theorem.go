package generation2electroweakiksource

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ElectroweakKOverlapSourceSearchTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 electroweak K-overlap source search audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate474 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate473 missing-I_K boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate473MassClosureFailed && a.Inheritance.MissingIK && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "audits Higgs, gauge, and PMNS-facing channels", Passed: a.Sieve.Executed && len(a.Sieve.Candidates) == 3 && a.Sieve.NativeSelectors == 0 && !a.Sieve.IKHalfDerived && a.Sieve.Verdict == StatusFailedNoNativeIKSource, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects Higgs VEV and gauge couplings as family K-overlap selectors", Passed: !a.Sieve.Candidates[0].SuppliesIK && a.Sieve.Candidates[0].GenerationBlind && !a.Sieve.Candidates[1].SuppliesIK && a.Sieve.Candidates[1].GenerationBlind, Detail: FormatCandidate(a.Sieve.Candidates[0]) + "\n" + FormatCandidate(a.Sieve.Candidates[1])},
			{Name: "quarantines PMNS/lepton sector as bridge comparator only", Passed: a.Sieve.Candidates[2].FamilySensitive && a.Sieve.Candidates[2].RequiresEmpiricalAirlock && !a.Sieve.Candidates[2].SuppliesIK && a.Frontier.CanUsePMNSAsBridgeComparator && !a.Frontier.CanUseAnyAsNativeIKSelector, Detail: FormatCandidate(a.Sieve.Candidates[2]) + "\n" + FormatFrontier(a.Frontier)},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.HiggsVEVNativeIK && !a.Firewall.GaugeCouplingsNativeIK && !a.Firewall.PMNSNativeIK && !a.Firewall.IKHalfNative && !a.Firewall.CKMNativePrediction && !a.Firewall.PMNSNativePrediction && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusAuditExecuted, StatusFailedHiggsGenerationBlind, StatusFailedGaugeGenerationBlind, StatusFailedPMNSNeedsAirlock, StatusFailedNoNativeIKSource, StatusFailedIKHalfNotDerived, StatusFrontierDefined, StatusFirewallPreserved, a.Truth}}
	}}
}
