package generation2gaugescalarboundarystresssealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2GaugeScalarBoundaryStressSealAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 joint gauge-scalar boundary stress seal audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate613 boundary stress audit", Passed: false, Detail: err.Error()}}}
		}
		mean := findCandidate(a.CompressionCandidates, "xi_mean")
		checks := []theorem.Check{
			{Name: "inherit Gate612 pairing robustness", Passed: a.Inherited.PairingSharpensAtLambda12 && a.Inherited.PairingIsV1Sensitive && a.Inherited.Verdict == StatusGate612Inherited, Detail: FormatInherited(a.Inherited)},
			{Name: "test one-parameter boundary-stress compression", Passed: len(a.CompressionCandidates) == 4 && mean.Xi > 0 && mean.MaxAbsNormalizedResidual < 0.02, Detail: FormatCompressionCandidates(a.CompressionCandidates)},
			{Name: "define signed stress vector", Passed: a.SignedStressVector.SPlus > 0 && a.SignedStressVector.SMinus > 0 && a.SignedStressVector.Verdict == StatusSignedStressVectorDefined, Detail: FormatSignedStress(a.SignedStressVector)},
			{Name: "audit anti-alignment", Passed: a.AntiAlignment.AntiAligned && a.AntiAlignment.RelativeAntiAlignment < 0.03, Detail: FormatAntiAlignment(a.AntiAlignment)},
			{Name: "compare eta3 to twice stress scale", Passed: len(a.EtaComparisons) >= 2 && etaFor(a.EtaComparisons, "xi_mean").EtaOverTwoXi > 0.9 && etaFor(a.EtaComparisons, "xi_mean").EtaOverTwoXi < 1.0, Detail: FormatEtaComparisons(a.EtaComparisons)},
			{Name: "define gauge-scalar boundary stress seal", Passed: !a.StressSeal.NativeCorrectionTheorem && a.StressSeal.XiBoundary > 0 && a.StressSeal.Verdict == StatusBoundaryStressSealDefined, Detail: FormatStressSeal(a.StressSeal)},
			{Name: "inherit robustness and sensitivity cautions", Passed: a.Robustness.PairingSharpensAtLambda12 && a.Robustness.ScalarV1Sensitive && a.Robustness.ThresholdSensitive, Detail: FormatRobustness(a.Robustness)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.ProvidesNativeXiBoundary && !a.NativeStatus.ProvidesNativeGaugeScalarEquation && !a.NativeStatus.ClaimsHiggsPrediction, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsLambdaZeroBoundary && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsScalarStability && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsThresholdExistence && !a.Firewalls.DerivesEndpoint, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func etaFor(rows []EtaComparison, name string) EtaComparison {
	for _, r := range rows {
		if r.XiName == name {
			return r
		}
	}
	return EtaComparison{XiName: name}
}
