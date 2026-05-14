package inputsensitivityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func InputSensitivityBottomTauYukawaCompletenessAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-INPUT-SENSITIVITY-BOTTOM-TAU-YUKAWA-COMPLETENESS"
	const name = "input-sensitivity and bottom-tau-Yukawa completeness audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 219 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 218 sealed full-SM-Yukawa audit is inherited", Passed: a.Gate218.Gate218Inherited && a.Gate218.MatchingCorrectionSealActive && a.Gate218.ThresholdSpectrumSealActive && a.Gate218.MatchingPlausible && a.Gate218.BottomTauOmitted, Detail: FormatGate218(a.Gate218)},
			{Name: "bottom and tau Yukawa completeness is added without finite-core promotion", Passed: a.Completeness.TopYukawaIncluded && a.Completeness.BottomYukawaIncluded && a.Completeness.TauYukawaIncluded && a.Completeness.HiggsQuarticIncluded && !a.Completeness.FullYukawaMatricesDerived && !a.Completeness.HeavyYukawaCouplingsAdded, Detail: FormatCompleteness(a.Completeness)},
			{Name: "empirical input ledger defines bounded 1σ scan without tuning", Passed: len(a.Inputs) >= 5 && !a.CentralInputs.FiniteCoreDerived && a.CentralInputs.UsesTreeLevelSeeds, Detail: FormatInputs(a.Inputs) + " :: " + FormatPhenomenologicalInputs(a.CentralInputs)},
			{Name: "central bottom/tau-complete single-scale fit remains inside matching envelope", Passed: a.CentralFit.Converged && a.CentralFit.MatchingPlausible && a.CentralFit.ResidualOverEpsilon < 1 && a.CentralFit.PositiveToBoundary && a.CentralFit.NoLandauBelowPlanck, Detail: FormatFit(a.CentralFit)},
			{Name: "1σ sensitivity scan preserves plausibility envelope", Passed: a.Sensitivity.CasesAudited >= 11 && a.Sensitivity.ConvergedCases == a.Sensitivity.CasesAudited && a.Sensitivity.BrokenEnvelopeCases == 0 && a.Sensitivity.WorstResidualOverEpsilon < 1, Detail: FormatSensitivity(a.Sensitivity) + "\n" + FormatScanTop(a.ScanCases, 4)},
			{Name: "MatchingCorrectionSeal remains an uncertainty seal, not a derived correction", Passed: a.Seal.MatchingCorrectionSealInherited && a.Seal.RequiredResidualQuarantined && !a.Seal.MatchingCorrectionsDerived && !a.Seal.MatchingResidualPromoted && a.Seal.EpsilonU > 0, Detail: FormatSeal(a.Seal)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate218Inherited && a.Firewall.MatchingCorrectionSealActive && a.Firewall.ThresholdSpectrumSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.InputUncertaintiesFiniteDerived && !a.Firewall.InputsTunedToForceZeroResidual && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.MatchingResidualPromoted && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.ProtonLifetimeComputed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "Gate 219 is conditional phenomenology: it propagates empirical input uncertainty through a sealed spectrum/matching framework and does not derive SM masses or matching corrections from the finite core."}}
	}}
}
