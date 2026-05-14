package flavoralignmentdmabsence

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FlavorAlignmentSafetyDarkMatterAbsenceTheorem() theorem.Theorem {
	const id = "BRIDGE-FLAVOR-ALIGNMENT-DARK-MATTER-ABSENCE"
	const name = "Flavor alignment safety audit / Dark Matter absence theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 224 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 223 RelicDecaySeal is inherited", Passed: a.Gate223.Gate223Inherited && a.Gate223.RelicDecaySealGranted && a.Gate223.TripletPortalActive && a.Gate223.OctetPortalActive, Detail: a.Gate223.TruthStatement},
			{Name: "generic flavor tensors are rejected as unsafe", Passed: !a.Flavor.GenericFlavorSafe && !a.Flavor.ArbitraryFirstSecondAllowed && a.Flavor.FlavorChangingRiskLogged && len(a.Flavor.PortalsAudited) == 3, Detail: FormatFlavor(a.Flavor)},
			{Name: "FlavorAlignmentSeal is granted only as a third-generation phenomenological seal", Passed: a.Seal.SealGranted && a.Seal.StillPhenomenological && !a.Seal.NativeFlavorTheoremDerived && len(a.Seal.QuarantinedInputs) > 0, Detail: FormatSeal(a.Seal)},
			{Name: "RelicDecaySeal remains valid only under flavor alignment", Passed: a.DarkMatter.RelicDecaySealActive && a.DarkMatter.FlavorAlignmentSealActive && a.DarkMatter.TripletDecaysBeforeBBN && a.DarkMatter.OctetDecaysBeforeBBN, Detail: FormatDarkMatter(a.DarkMatter)},
			{Name: "heavy PeV threshold sector is not dark matter", Passed: a.DarkMatter.OmegaHeavySectorH2 == 0 && a.DarkMatter.PresentDayStableFraction == 0 && !a.DarkMatter.HeavySectorDMCandidate && len(a.DarkMatter.DarkMatterDeferredTo) > 0, Detail: FormatDarkMatter(a.DarkMatter)},
			{Name: "future dark matter and flavor precision routes are explicitly deferred", Passed: len(a.Future.OpenDMInventory) > 0 && len(a.Future.RequiredNextObjects) > 0 && len(a.Future.OpenFlavorInventory) > 0, Detail: a.Future.Verdict},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate223Inherited && a.Firewall.RelicDecaySealActive && a.Firewall.ThresholdSpectrumSealActive && a.Firewall.MatchingCorrectionSealActive && a.Firewall.EmpiricalCarrierSealActive && a.Firewall.LeptoquarkDynamicsSealActive && !a.Firewall.NativeFlavorClaimed && !a.Firewall.ExactFCNCRatesClaimed && !a.Firewall.FlavorInputsTuned && !a.Firewall.WilsonCoefficientsDerived && !a.Firewall.RelicAbundanceThermalClaimed && !a.Firewall.HeavySectorDMClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "Gate 224 keeps the decay portals alive only after a FlavorAlignmentSeal. Since both sealed heavy carriers decay before BBN, the PeV threshold sector contributes no present-day dark matter. Dark matter must be searched for in a different finite or sealed sector."}}
	}}
}
