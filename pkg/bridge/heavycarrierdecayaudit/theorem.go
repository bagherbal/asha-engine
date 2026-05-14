package heavycarrierdecayaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func HeavyCarrierDecayRelicSafetyAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-HEAVY-CARRIER-DECAY-RELIC-SAFETY"
	const name = "Heavy-carrier decay and mass-splitting / cosmological-relic safety audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 221 audit", Passed: false, Detail: err.Error()}}}
		}
		carrierDetails := ""
		for _, c := range a.Carriers {
			carrierDetails += FormatCarrier(c) + "\n"
		}
		checks := []theorem.Check{
			{Name: "Gate 220 stable-relic warning is inherited", Passed: a.Gate220.Gate220Inherited && a.Gate220.CosmologyWarning && a.Gate220.StableNeutralRelicWarning && a.Gate220.StableChargedRelicWarning && a.Gate220.StableColoredRelicWarning, Detail: FormatGate220(a.Gate220)},
			{Name: "sealed PeV carriers are identified without finite-core promotion", Passed: len(a.Carriers) == 2 && a.Carriers[0].ConditionalOnly && a.Carriers[1].ConditionalOnly && a.Carriers[0].HasNeutralComponent && a.Carriers[1].HasColoredComponent, Detail: carrierDetails},
			{Name: "operator basis search finds no native decay portal", Passed: a.Operators.NativePortalSearchFailed && a.Operators.DecayOperatorsDerived == 0 && !a.Operators.AnyDecayWidthComputable && !a.Operators.PortalOperatorFound, Detail: FormatOperatorAudit(a.Operators)},
			{Name: "mass splitting is not derived", Passed: a.MassSplitting.TreeDegenerateBySeal && !a.MassSplitting.ElectroweakLoopSplittingDerived && !a.MassSplitting.VEVCouplingSplittingDerived && !a.MassSplitting.ChargedToNeutralCascadeDerived && a.MassSplitting.StableChargedRisk && a.MassSplitting.StableColoredRisk, Detail: FormatMassSplitting(a.MassSplitting)},
			{Name: "BBN lifetime check fails by operator absence", Passed: !a.Lifetime.DecayWidthDerived && !a.Lifetime.PassesBBN && a.Lifetime.FailsBBNByOperatorAbsence && a.Lifetime.RequiredWidthGeV > 0, Detail: FormatLifetime(a.Lifetime)},
			{Name: "RelicDecaySeal is required but not granted", Passed: a.RelicSeal.SealRequired && !a.RelicSeal.SealGranted && a.RelicSeal.OperationalStatus == StatusRelicDecaySealRequired, Detail: FormatRelicSeal(a.RelicSeal)},
			{Name: "cosmological safety is not cleared", Passed: a.Summary.FatalPathology && !a.Summary.CosmologyCleared && a.Summary.Status == StatusFailedCosmologicalPathology, Detail: FormatSummary(a.Summary)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate220Inherited && a.Firewall.ThresholdSpectrumSealActive && a.Firewall.MatchingCorrectionSealActive && a.Firewall.EmpiricalCarrierSealActive && a.Firewall.LeptoquarkDynamicsSealActive && !a.Firewall.DecayOperatorInvented && !a.Firewall.MassSplittingInvented && !a.Firewall.LifetimeComputedFromAbsentOp && !a.Firewall.RelicAbundanceComputed && !a.Firewall.DarkMatterClaimed && !a.Firewall.ArbitraryCouplingIntroduced && a.Firewall.BBNUsedAsFilterOnly && !a.Firewall.PeVMassFiniteDerived, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.TruthStatement, "Gate 221 is a failed-route cosmology theorem: the PeV spectrum remains precision-safe, but it is not cosmologically safe until a finite or explicitly sealed decay/splitting sector supplies a BBN-safe lifetime."}}
	}}
}
