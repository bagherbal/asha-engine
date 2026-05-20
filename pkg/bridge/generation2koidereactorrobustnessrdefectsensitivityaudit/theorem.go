package generation2koidereactorrobustnessrdefectsensitivityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideReactorRobustnessRDefectSensitivityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide-reactor robustness and R-defect sensitivity audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate589 robustness audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit observed and exact-R=1 wall coordinates", Passed: a.Runtime.RObs > 0.9999 && a.Runtime.RDefect > 0 && a.Runtime.EpsilonR1Rad > a.Runtime.EpsilonObsRad && a.Runtime.KappaObs > a.Runtime.KappaR1, Detail: FormatRuntime(a.Runtime)},
			{Name: "inherit NuFIT 6.0 reactor-quarter candidate", Passed: nearly(a.Reactor.Candidate, 0.0055375, 1e-16) && a.Reactor.CandidateMin < a.Runtime.KappaObs && a.Reactor.CandidateMax > a.Runtime.KappaObs, Detail: FormatReactor(a.Reactor)},
			{Name: "verify reactor-quarter matches observed epsilon better than exact-R=1 epsilon", Passed: a.Robustness.ObservedBetter && a.Robustness.ExactR1WeakerFactor > 6.0 && abs(a.Robustness.Observed.RelativeResidual) < abs(a.Robustness.ExactR1.RelativeResidual), Detail: FormatRobustness(a.Robustness)},
			{Name: "verify observed inverse theta13 prediction lies inside one sigma", Passed: a.Robustness.Observed.WithinThetaOneSigma && a.Robustness.Observed.CoveredByOneSigma && a.Robustness.Observed.Sin2Theta13Pred > a.Reactor.Sin2Theta13-a.Reactor.Sin2Theta13Minus, Detail: FormatKappaComparison(a.Robustness.Observed)},
			{Name: "verify exact-R=1 inverse theta13 prediction falls outside one sigma", Passed: !a.Robustness.ExactR1.WithinThetaOneSigma && !a.Robustness.ExactR1.CoveredByOneSigma && a.Robustness.ExactR1.Sin2Theta13Pred < a.Reactor.Sin2Theta13-a.Reactor.Sin2Theta13Minus, Detail: FormatKappaComparison(a.Robustness.ExactR1)},
			{Name: "compute R-defect correction coefficient and reject typed simple candidates", Passed: a.RDefect.RequiredC > 20.0 && a.RDefect.RequiredC < 21.0 && a.RDefect.BestCandidate.Name == "8*pi" && !a.RDefect.BestCandidateCertified, Detail: FormatRDefect(a.RDefect)},
			{Name: "show kappa shift is exactly epsilon projection shift but not typed R/Q residual", Passed: a.Shift.ControlledByEpsilonShift && !a.Shift.ControlledByRDefectTyped && !a.Shift.ControlledByQResidualTyped && a.Shift.RatioToDROneMinusR == a.RDefect.RequiredC, Detail: FormatShift(a.Shift)},
			{Name: "preserve operator obstruction", Passed: !a.Operator.NativeKoideReactorOperatorPresent && !a.Operator.NativeRDefectCorrectionOperatorPresent && !a.Operator.NativeRootTraceOperatorPresent && !a.Operator.DerivesTheta13 && !a.Operator.DerivesKappa && !a.Operator.DerivesEpsilon, Detail: FormatOperator(a.Operator)},
			{Name: "preserve flavor, PMNS, observed-data, and Gate352 firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesTheta13 && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesNeutrinoPhysics && !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesFlavorLaw && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return relation as environmental seal", Passed: a.Final.ReactorMatchesObservedBetter && a.Final.ObservedInsideOneSigma && !a.Final.ExactR1InsideOneSigma && a.Final.RDefectRequiredForBestMatch && !a.Final.TypedRDefectCorrectionPresent && !a.Final.NativeOperatorPresent && a.Final.RelationRemainsEnvironmental, Detail: FormatFinal(a.Final)},
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
