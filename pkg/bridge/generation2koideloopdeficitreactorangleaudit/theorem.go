package generation2koideloopdeficitreactorangleaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideLoopDeficitReactorAngleAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide loop-deficit reactor-angle audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate588 reactor-angle audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate587 kappa, loop unit, and prior PMNS/CKM clues", Passed: a.Runtime.Kappa > 0.0054 && a.Runtime.Kappa < 0.0056 && a.Runtime.PriorPMNSAssistedRel < 0 && a.Runtime.CKMAlpha2MidpointRel < 0, Detail: FormatRuntime(a.Runtime)},
			{Name: "inherit NuFIT 6.0 reactor-angle input with uncertainty", Passed: a.Input.Sin2Theta13 == 0.02215 && a.Input.Sin2Theta13Plus == 0.00056 && a.Input.Sin2Theta13Minus == 0.00058 && a.Input.Theta13Deg > 8.55, Detail: FormatInput(a.Input)},
			{Name: "compute sin^2(theta13)/4 candidate", Passed: nearly(a.Candidate.Value, 0.0055375, 1e-16) && a.Candidate.Near && !a.Candidate.Certified, Detail: FormatCandidate(a.Candidate)},
			{Name: "verify kappa_e lies within theta13 one-sigma reactor-quarter range", Passed: a.Candidate.CoversKappa && a.Candidate.Min1Sigma < a.Runtime.Kappa && a.Candidate.Max1Sigma > a.Runtime.Kappa, Detail: FormatCandidate(a.Candidate)},
			{Name: "invert relation to theta13 prediction inside NuFIT one-sigma", Passed: a.Inverse.WithinSin2OneSigma && a.Inverse.WithinThetaOneSigma && a.Inverse.Sin2Theta13Pred > 0.0220 && a.Inverse.Theta13PredDeg > 8.53, Detail: FormatInverse(a.Inverse)},
			{Name: "compute full epsilon prediction", Passed: a.Epsilon.CoversTargetEpsilon && a.Epsilon.SignedResidualRad < 0 && abs(a.Epsilon.SignedResidualDeg) < 1e-4, Detail: FormatEpsilon(a.Epsilon)},
			{Name: "show reactor-quarter beats previous PMNS-assisted and sqrt(J_CKM), but not CKM/alpha midpoint", Passed: a.Comparison.BeatsPriorPMNSAssisted && a.Comparison.BeatsSqrtJCKM && !a.Comparison.BeatsCKMAlpha2Midpoint && a.Comparison.CKMMidpointStillClosest, Detail: FormatComparison(a.Comparison)},
			{Name: "keep factor one-quarter as weak-normalization clue only", Passed: a.Operator.FactorOneQuarterInterpretedAsWeakNormalizationClue && !a.Operator.NativeLeptonOrientationOperatorPresent && !a.Operator.NativeWeakDoubletOperatorPresent && !a.Operator.NativeRootTraceOperatorPresent && !a.Operator.DerivesTheta13 && !a.Operator.DerivesKappa, Detail: FormatOperator(a.Operator)},
			{Name: "preserve Koide, PMNS, neutrino, flavor, and root-trace firewalls", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesNeutrinoParameters && !a.Firewalls.DerivesTheta13 && !a.Firewalls.DerivesFlavorTexture && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return kappa_e as environmental seal", Passed: a.Final.KappaWithinTheta13OneSigma && a.Final.Theta13PredWithinOneSigma && a.Final.BetterThanPriorPMNS && !a.Final.AnyNativeOperator && a.Final.KappaRemainsSeal, Detail: FormatFinal(a.Final)},
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
