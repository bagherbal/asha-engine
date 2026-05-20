package generation2gaugescalarboundarypairingrobustnessaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2GaugeScalarBoundaryPairingRobustnessAndScaleDependenceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 gauge-scalar boundary pairing robustness and scale-dependence audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate612 robustness audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate611 pairing", Passed: a.InheritedPairing.R3MinusOne > 0 && a.InheritedPairing.AbsLambda12 > 0 && a.InheritedPairing.Verdict == StatusGate611PairingInherited, Detail: FormatInheritedPairing(a.InheritedPairing)},
			{Name: "enumerate candidate gauge boundary scales", Passed: len(a.CandidateScales) == 4 && hasScale(a.CandidateScales, "Lambda_12") && hasScale(a.CandidateScales, "Lambda_geom"), Detail: FormatCandidateScales(a.CandidateScales)},
			{Name: "compute scale-dependent gauge residuals", Passed: len(a.GaugeResiduals) == 4 && residualFor(a.GaugeResiduals, "Lambda_12").GaugeRelativeResidual > 0.05, Detail: FormatGaugeResiduals(a.GaugeResiduals)},
			{Name: "compute scalar transport at candidate scales", Passed: len(a.ScalarValues) == 4 && scalarFor(a.ScalarValues, "Lambda_12").Lambda < 0, Detail: FormatScalarValues(a.ScalarValues)},
			{Name: "compute pairing ratios by scale", Passed: len(a.PairingRatios) == 4 && pairingFor(a.PairingRatios, "Lambda_12").RatioGaugeToAbsLambda > 1.0, Detail: FormatPairings(a.PairingRatios)},
			{Name: "audit Lambda12 uniqueness", Passed: a.UniquenessAudit.Verdict == StatusLambda12SharpensPairing && a.UniquenessAudit.Lambda12UniqueBest, Detail: FormatUniqueness(a.UniquenessAudit)},
			{Name: "estimate local scalar sensitivity", Passed: a.LocalSensitivity.BetaLambdaLambda12 != 0 && a.LocalSensitivity.Verdict == StatusLocalSensitivityComputed, Detail: FormatLocalSensitivity(a.LocalSensitivity)},
			{Name: "record loop and matching sensitivity", Passed: a.SensitivityLedger.ScalarSideFragile && !a.SensitivityLedger.ClosureCertified, Detail: FormatSensitivity(a.SensitivityLedger)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.ProvidesNativeJointCorrectionTheorem && !a.NativeStatus.ProvidesNativeScalarBoundaryCondition && !a.NativeStatus.ClaimsHiggsPrediction, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsLambdaZeroBoundary && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsThresholdExistence && !a.Firewalls.DerivesEndpoint, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasScale(rows []CandidateScale, name string) bool {
	for _, r := range rows {
		if r.Name == name {
			return true
		}
	}
	return false
}
