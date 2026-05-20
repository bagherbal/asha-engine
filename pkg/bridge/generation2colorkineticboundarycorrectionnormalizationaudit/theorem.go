package generation2colorkineticboundarycorrectionnormalizationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ColorKineticBoundaryCorrectionNormalizationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 color kinetic boundary correction normalization audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate610 color kinetic boundary audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate609 sign audit", Passed: a.Inherited.Delta3Required > 0 && a.Inherited.DeltaB3Required < 0 && a.Inherited.ExtraColoredMatterWrongSign, Detail: FormatInherited(a.Inherited)},
			{Name: "define color boundary correction slot", Passed: containsBoundaryQuantity(a.BoundaryCorrections, "delta_3^color_boundary") && a.BoundaryCorrections[2].Value > 0, Detail: FormatBoundaryCorrections(a.BoundaryCorrections)},
			{Name: "compute fractional color kinetic shift", Passed: a.FractionalCorrection.EtaAgainstUStar > 0.09 && a.FractionalCorrection.EtaAgainstUStar < 0.10 && a.FractionalCorrection.EtaAgainstU3 > 0.10, Detail: FormatFractionalCorrection(a.FractionalCorrection)},
			{Name: "audit spectral-action gauge coefficient lane", Passed: a.GaugeCoefficientAudit.SignCompatible && !a.GaugeCoefficientAudit.Native && !a.GaugeCoefficientAudit.Certified, Detail: FormatGaugeCoefficientAudit(a.GaugeCoefficientAudit)},
			{Name: "compare trace normalizations", Passed: len(a.TraceNormalizations) >= 4 && containsTraceObject(a.TraceNormalizations, "k_Y") && containsTraceObject(a.TraceNormalizations, "SU(2)"), Detail: FormatTraceNormalizations(a.TraceNormalizations)},
			{Name: "audit finite spectral-action status", Passed: !a.FSAStatus.HasIndependentColorKineticCoefficient && !a.FSAStatus.HasSectorSplitF0Moment && !a.FSAStatus.HasSU3OnlyBoundaryCorrection, Detail: FormatFSAStatus(a.FSAStatus)},
			{Name: "classify localized threshold interpretation", Passed: a.ThresholdLocalized.SignCompatible && a.ThresholdLocalized.RequiredDeltaU > 0 && a.ThresholdLocalized.CleanerThanFullIntervalMatter, Detail: FormatThresholdLocalized(a.ThresholdLocalized)},
			{Name: "record two-loop and scheme caution", Passed: a.TwoLoopSchemeCaution.TwoLoopCouldShiftResidual && !a.TwoLoopSchemeCaution.ClosureCertified, Detail: FormatTwoLoopSchemeCaution(a.TwoLoopSchemeCaution)},
			{Name: "audit native ASHA status", Passed: !a.NativeStatus.ProvesColorKineticBoundaryCorrection && !a.NativeStatus.ProvesThresholdSpectrum && !a.NativeStatus.ProvesFullGaugeUnification && !a.NativeStatus.AltersAF && !a.NativeStatus.AddsColoredStates, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve firewalls", Passed: !a.Firewalls.ClaimsCorrectionExists && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.AltersFiniteAlgebra && !a.Firewalls.AddsNewColoredStates && !a.Firewalls.DerivesEndpoint, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
