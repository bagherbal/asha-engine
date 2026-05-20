package generation2koidereactorckmresidualclosureuncertaintyaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideReactorCKMResidualClosureAndUncertaintyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide-reactor-CKM residual closure and uncertainty audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate591 residual audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate590 combined relation", Passed: a.Residual.ImprovementOverA > 10 && a.Residual.AbsDelta590 < 3e-6, Detail: FormatResidual(a.Residual)},
			{Name: "import CKM J uncertainty for uncertainty audit", Passed: a.Inputs.JCKMPlus > 0 && a.Inputs.JCKMMinus > 0 && a.Inputs.JCKMUncertaintyCtr > 3e-5, Detail: FormatInputs(a.Inputs)},
			{Name: "propagate theta13 and CKM uncertainties", Passed: a.Uncertainty.CoversKappa && a.Uncertainty.SigmaFractionPlus < 0.03 && a.Uncertainty.SigmaFractionMinus < 0.03, Detail: FormatUncertainty(a.Uncertainty)},
			{Name: "confirm theta13 dominates available uncertainty", Passed: a.Uncertainty.Theta13WidthPlus/a.Uncertainty.CKMWidthPlus > 100 && a.Uncertainty.Theta13WidthMinus/a.Uncertainty.CKMWidthMinus > 100, Detail: FormatUncertainty(a.Uncertainty)},
			{Name: "inverse theta13 prediction remains inside one sigma with CKM uncertainty", Passed: a.Inverse.WithinOneSigma && a.Inverse.ThetaPredLowDeg > a.Inverse.NuFITLowDeg && a.Inverse.ThetaPredHighDeg < a.Inverse.NuFITHighDeg, Detail: FormatInverse(a.Inverse)},
			{Name: "compare residual to Koide R and Q defects", Passed: a.Defects.DeltaSmallerThanR && a.Defects.DeltaSmallerThanQ && a.Defects.DeltaOverRDefect < 0.31 && a.Defects.DeltaOverAbsQ < 0.46, Detail: FormatDefects(a.Defects)},
			{Name: "test but reject R/Q corrected closure", Passed: !a.Corrections.AnyCertified && a.Corrections.BestCandidate.AbsResidual < a.Residual.AbsDelta590 && a.Corrections.RequiredRDefectCoefficient > 0.29 && a.Corrections.RequiredRDefectCoefficient < 0.31, Detail: FormatCorrections(a.Corrections)},
			{Name: "reject native cross-sector residual operator", Passed: !a.Lawfulness.CrossSectorOrientationIntertwinerPresent && !a.Lawfulness.RDefectToOrientationOperatorPresent && !a.Lawfulness.QDefectToOrientationOperatorPresent && !a.Lawfulness.NativeRootTraceOperatorPresent && !a.Lawfulness.DerivesDelta590 && !a.Lawfulness.DerivesKappa, Detail: FormatLawfulness(a.Lawfulness)},
			{Name: "preserve flavor and observed-data firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesTheta13 && !a.Firewalls.DerivesNeutrinoPhysics && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesFlavorHierarchy && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return residual as unresolved environmental seal", Passed: !a.Final.DeltaStatisticallyMeaningful && !a.Final.CorrectedFormulaCertified && !a.Final.CrossSectorBridgePresent && a.Final.KappaRemainsEnvironmental, Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
