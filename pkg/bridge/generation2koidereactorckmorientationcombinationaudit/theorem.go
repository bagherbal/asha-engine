package generation2koidereactorckmorientationcombinationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideReactorCKMOrientationCombinationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide-reactor-CKM orientation combination audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate590 combination audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit observed near-Koide wall coordinate", Passed: a.Runtime.KappaObs > 0 && a.Runtime.RObs > 0.9999 && a.Runtime.EpsilonObsRad > 0, Detail: FormatRuntime(a.Runtime)},
			{Name: "inherit NuFIT reactor angle and runtime CKM Jarlskog", Passed: nearly(a.Reactor.Sin2Theta13, 0.02215, 1e-14) && a.CKM.JCKM > 3.0e-5 && a.CKM.JCKM < 3.2e-5, Detail: FormatCKM(a.CKM)},
			{Name: "compute reactor-quarter candidate A and combined candidate B", Passed: nearly(a.Combination.AReactorQuarter.Value, 0.0055375, 1e-16) && a.Combination.BReactorMinusCKM.Value < a.Combination.AReactorQuarter.Value, Detail: FormatCombination(a.Combination)},
			{Name: "verify B outperforms A by more than one order of magnitude", Passed: a.Combination.BOutperformsA && a.Combination.BImprovementFactor > 10.0 && abs(a.Combination.BReactorMinusCKM.RelativeResidual) < 6.0e-4, Detail: FormatCombination(a.Combination)},
			{Name: "compute epsilon prediction and show combined residual improvement", Passed: abs(a.Epsilon.ResidualB_rad) < abs(a.Epsilon.ResidualA_rad) && a.Epsilon.ImprovementFactor > 10.0 && abs(a.Epsilon.ResidualB_deg) < 1e-5, Detail: FormatEpsilon(a.Epsilon)},
			{Name: "compute inverse theta13 prediction inside NuFIT one sigma", Passed: a.Inverse.WithinSin2OneSigma && a.Inverse.WithinThetaOneSigma && abs(a.Inverse.ThetaResidualDeg) < 0.003, Detail: FormatInverse(a.Inverse)},
			{Name: "propagate available theta13 uncertainty and mark CKM uncertainty missing", Passed: a.Uncertainty.CoversKappaWithTheta13 && !a.Uncertainty.CKMUncertaintyPresent && !a.Uncertainty.FullUncertaintyCertified, Detail: FormatUncertainty(a.Uncertainty)},
			{Name: "reject native cross-sector orientation bridge", Passed: !a.Lawfulness.CrossSectorOrientationIntertwinerPresent && !a.Lawfulness.CKMToChargedLeptonWallOperatorPresent && !a.Lawfulness.NativeRootTraceOperatorPresent && !a.Lawfulness.DerivesKappa && !a.Lawfulness.DerivesTheta13 && !a.Lawfulness.DerivesJCKM, Detail: FormatLawfulness(a.Lawfulness)},
			{Name: "show remaining residual is not typed R-defect or Q-residual correction", Passed: !a.Residual.TypedCoefficientPresent && abs(a.Residual.CombinedResidual) < abs(a.Runtime.RDefect) && abs(a.Residual.CombinedResidual) < abs(a.Runtime.QResidual), Detail: FormatResidual(a.Residual)},
			{Name: "preserve flavor and observed-data firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesTheta13 && !a.Firewalls.DerivesNeutrinoPhysics && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesFlavorHierarchy && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return relation as environmental seal", Passed: a.Final.BOutperformsA && a.Final.InverseTheta13WithinOneSigma && !a.Final.CrossSectorBridgePresent && a.Final.KappaRemainsEnvironmental && abs(a.Final.RemainingRelativeResidual) < 6e-4, Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func nearly(a, b, tol float64) bool { return abs(a-b) <= tol }
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
