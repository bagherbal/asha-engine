package generation2gaugescalarboundarystresssourcetypeaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2GaugeScalarBoundaryStressSourceTypeAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 gauge-scalar boundary stress source-type and spectral-action lane audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate614 stress source-type audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate613 stress seal", Passed: a.Inherited.XiBoundary > 0 && a.Inherited.Verdict == StatusGate613Inherited, Detail: FormatInherited(a.Inherited)},
			{Name: "classify source types", Passed: len(a.SourceTypes) == 4 && hasSourceVerdict(a.SourceTypes, StatusSpectralActionSlotRelevant) && hasSourceVerdict(a.SourceTypes, StatusNoNativeXi), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "audit spectral-action lanes", Passed: len(a.SpectralActionLanes) >= 5 && hasLane("f0_cutoff_moment", a.SpectralActionLanes) && hasLane("gauge_kinetic_C_i_Tr_F_i2", a.SpectralActionLanes), Detail: FormatSpectralActionLanes(a.SpectralActionLanes)},
			{Name: "audit kinetic/quartic pairing lane", Passed: a.KineticQuarticPairing.SymbolicPairingSlot && !a.KineticQuarticPairing.NativeCoefficientLaw && a.KineticQuarticPairing.SameF0Dependence, Detail: FormatKineticQuarticPairing(a.KineticQuarticPairing)},
			{Name: "define approximate boundary stress equation", Passed: a.BoundaryEquation.AbsResidualOverXi < 0.03 && a.BoundaryEquation.HalfResidualOverXi < 0.02, Detail: FormatBoundaryEquation(a.BoundaryEquation)},
			{Name: "audit eta relation", Passed: a.EtaRelation.EtaOverTwoXi > 0.93 && a.EtaRelation.EtaOverTwoXi < 0.96, Detail: FormatEtaRelation(a.EtaRelation)},
			{Name: "record v1 sensitivity", Passed: a.SensitivityLedger.ScalarV1Sensitive && a.SensitivityLedger.TopMassSensitive && a.SensitivityLedger.ThresholdSensitive && a.SensitivityLedger.MatchingSensitive, Detail: FormatSensitivity(a.SensitivityLedger)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.NativeXiBoundary && !a.NativeStatus.NativeColorKineticCorrection && !a.NativeStatus.NativeF0SectorSplit && !a.NativeStatus.NativeGaugeScalarCoefficientLaw && !a.NativeStatus.NativeThresholdSpectrum, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsThresholdExists && !a.Firewalls.ClaimsHiggsStability && !a.Firewalls.ClaimsLambdaZeroBoundary && !a.Firewalls.ClaimsHiggsMassPrediction && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsWZHiggsPrediction && !a.Firewalls.ClaimsNativeCorrection, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func hasSourceVerdict(rows []SourceTypeClassification, verdict string) bool {
	for _, r := range rows {
		if r.Verdict == verdict {
			return true
		}
	}
	return false
}

func hasLane(name string, rows []SpectralActionLane) bool {
	for _, r := range rows {
		if r.Lane == name {
			return true
		}
	}
	return false
}
