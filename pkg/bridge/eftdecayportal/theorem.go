package eftdecayportal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func EFTDecayPortalRelicDecaySealActivationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-EFT-DECAY-PORTAL-RELIC-SEAL-ACTIVATION"
	const name = "EFT decay portal construction / RelicDecaySeal activation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 222 audit", Passed: false, Detail: err.Error()}}}
		}
		candidateDetails := ""
		for _, c := range a.Operators.Candidates {
			candidateDetails += FormatCandidate(c) + "\n"
		}
		checks := []theorem.Check{
			{Name: "Gate 221 cosmological pathology is inherited", Passed: a.Gate221.Gate221Inherited && a.Gate221.FatalCosmologicalPathology && a.Gate221.RelicDecaySealRequired && !a.Gate221.RelicDecaySealGranted, Detail: FormatGate221(a.Gate221)},
			{Name: "SM field inventory is explicit", Passed: len(a.Operators.SMFields) >= 7, Detail: FormatSMFields(a.Operators.SMFields)},
			{Name: "triplet lepton-Higgs portal is identified as sealed EFT only", Passed: a.Operators.TripletPortalFound && a.Kinematics.TripletCanPassBBNUnderSeal && a.Kinematics.TripletYukawaMin > 0 && a.Kinematics.TripletYukawaMin < 1e-12, Detail: candidateDetails + FormatKinematics(a.Kinematics)},
			{Name: "false octet-Q mass-mixing claim is rejected", Passed: a.Operators.OctetMassMixingRejected && !a.Firewall.OctetQMixingClaimed, Detail: FormatOperatorAudit(a.Operators)},
			{Name: "colored octet remains without certified pure-SM decay portal", Passed: a.Operators.OctetPortalObstructed && !a.Operators.OctetPureSMPortalFound && !a.Kinematics.OctetCanPassBBNUnderSeal, Detail: FormatOperatorAudit(a.Operators)},
			{Name: "full RelicDecaySeal is not granted", Passed: a.RelicSeal.SealRequested && !a.RelicSeal.SealGranted && a.RelicSeal.PartialTripletSubseal && a.RelicSeal.OperationalStatus == StatusFullRelicSealNotGranted, Detail: FormatRelicSeal(a.RelicSeal)},
			{Name: "cosmological pathology is not fully cleared", Passed: a.Summary.TripletRescuedByEFTPortal && !a.Summary.ColoredOctetRescued && !a.Summary.FullSpectrumRelicSafe && !a.Summary.FatalPathologyCleared, Detail: FormatSummary(a.Summary)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate221Inherited && a.Firewall.ThresholdSpectrumSealActive && a.Firewall.MatchingCorrectionSealActive && a.Firewall.EmpiricalCarrierSealActive && a.Firewall.LeptoquarkDynamicsSealActive && !a.Firewall.NativeOperatorClaimed && !a.Firewall.OctetQMixingClaimed && !a.Firewall.ArbitraryCouplingFixed && !a.Firewall.RelicAbundanceComputed && !a.Firewall.ProtonSealViolated && a.Firewall.BBNUsedAsFilterOnly && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.TruthStatement, "Gate 222 corrects a tempting but false shortcut: the colored octet is not SM-Q-like. The triplet can be rescued by a quarantined Yukawa portal, but the full PeV spectrum still lacks a legal colored decay channel."}}
	}}
}
