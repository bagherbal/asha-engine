package pevobservabilityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func PeVThresholdIndirectSignatureObservabilityAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-PEV-THRESHOLD-INDIRECT-SIGNATURE-OBSERVABILITY"
	const name = "PeV-threshold indirect-signature / experimental observability audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 220 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 219 sealed PeV precision result is inherited", Passed: a.Gate219.Gate219Inherited && a.Gate219.ThresholdSpectrumSealActive && a.Gate219.MatchingCorrectionSealActive && a.Gate219.MatchingPlausible && a.Gate219.BottomTauComplete, Detail: FormatGate219(a.Gate219)},
			{Name: "sealed PeV spectrum is audited without finite-core promotion", Passed: a.Spectrum.ConditionalOnly && !a.Spectrum.DecayOperatorsDerived && !a.Spectrum.MassSplittingsDerived && a.Spectrum.SingleScaleMBGeV > 1e6, Detail: FormatSpectrum(a.Spectrum)},
			{Name: "direct collider production is parametrically out of reach", Passed: a.DirectReach.DirectProductionSafe && a.DirectReach.MassOverReach > 10, Detail: FormatDirectReach(a.DirectReach)},
			{Name: "electroweak precision oblique proxy is PeV-suppressed", Passed: a.EWPO.ObliqueSafe && !a.EWPO.TreeLevelViolationDerived && !a.EWPO.HeavyYukawaCouplingDerived, Detail: FormatEWPO(a.EWPO)},
			{Name: "Higgs-loop imprints decouple without derived heavy Higgs mass coupling", Passed: a.HiggsLoops.HiggsLoopSafeUnderDecoupling && !a.HiggsLoops.HeavyYukawaCouplingDerived && !a.HiggsLoops.NonDecouplingMassFromHiggs, Detail: FormatHiggs(a.HiggsLoops)},
			{Name: "cosmological stable-relic warning is explicitly logged", Passed: a.Cosmology.StableNeutralRelicWarning && a.Cosmology.StableChargedRelicWarning && a.Cosmology.StableColoredRelicWarning && !a.Cosmology.DarkMatterCandidateClaimed && !a.Cosmology.OverclosureComputed, Detail: FormatCosmology(a.Cosmology)},
			{Name: "observability summary separates safety from relic warning", Passed: a.Summary.DirectReachSafe && a.Summary.EWPOSafe && a.Summary.HiggsLoopSafe && a.Summary.CosmologyWarning && !a.Summary.FatalObservableFailure, Detail: FormatSummary(a.Summary)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate219Inherited && a.Firewall.ThresholdSpectrumSealActive && a.Firewall.MatchingCorrectionSealActive && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && !a.Firewall.PeVMassFiniteDerived && !a.Firewall.DecayOperatorInvented && !a.Firewall.HeavyHiggsYukawaInvented && !a.Firewall.MassSplittingInvented && !a.Firewall.DarkMatterClaimed && !a.Firewall.OverclosureComputed && !a.Firewall.PhysicalObservationClaimed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "Gate 220 is conditional phenomenology: it audits indirect observability of the sealed PeV spectrum and does not derive carrier decays, heavy Yukawa couplings, dark matter stability, or physical observables from the finite core."}}
	}}
}
