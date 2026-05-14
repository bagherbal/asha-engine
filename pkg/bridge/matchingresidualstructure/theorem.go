package matchingresidualstructure

import "github.com/bagherbal/asha-engine/pkg/theorem"

func MatchingResidualStructureSpectralHeatKernelSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-MATCHING-RESIDUAL-STRUCTURE-SPECTRAL-HEAT-KERNEL-SEARCH"
	const name = "matching-residual structure audit / spectral heat-kernel coefficient search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 216 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 215 single-scale matching target is inherited", Passed: a.Gate215.Gate215Inherited && a.Gate215.ThresholdSpectrumSealInherited && a.Gate215.MatchingEnvelopeInherited && a.Gate215.SingleScaleCandidateUnique && !a.Gate215.FiniteMatchingDerived, Detail: FormatGate215(a.Gate215)},
			{Name: "finite spectral data are audited without physical activation", Passed: a.SpectralData.BGapAvailable && a.SpectralData.ContactPartialModeCount == 7 && a.SpectralData.ContactZetaLedgerAvailable && a.SpectralData.FiniteFundamentalClassAvailable && a.SpectralData.AllSpectralScalarsDimensionless && !a.SpectralData.FiniteMatchingRowsDerived, Detail: FormatSpectralData(a.SpectralData)},
			{Name: "Gate-215 residual is an alternating target, not a derived trace", Passed: a.ResidualTarget.AlternatingPattern && a.ResidualTarget.NearlyEqualMagnitudes && !a.ResidualTarget.TraceTargetInterpretedAsDerived, Detail: FormatResidualTarget(a.ResidualTarget)},
			{Name: "spectral trace candidates do not provide a canonical full match", Passed: len(a.TraceCandidates) >= 6 && a.HeatKernelMap.SignOnlyResonances >= 1 && a.HeatKernelMap.FullStructuralMatches == 0 && a.CoefficientSearch.ExactMagnitudeMatches == 0, Detail: FormatCandidates(a.TraceCandidates)},
			{Name: "heat-kernel matching map remains missing", Passed: a.HeatKernelMap.A2A4LanguageAudited && !a.HeatKernelMap.FiniteDiracOperatorDerived && !a.HeatKernelMap.SpectralTripleComplete && a.HeatKernelMap.GaugeCurvatureProjectionRows == 0 && !a.HeatKernelMap.GaugeKineticTraceMapDerived && !a.HeatKernelMap.CanonicalCutoffMomentsDerived && !a.HeatKernelMap.ThresholdSubtractionSchemeDerived && a.HeatKernelMap.DeltaMatchRowsDerived == 0, Detail: FormatHeatKernel(a.HeatKernelMap)},
			{Name: "coefficient search rejects near-misses and fitted normalizations", Passed: a.CoefficientSearch.CanonicalScalarsAudited > 0 && a.CoefficientSearch.CanonicalLoopScaledCandidates > 0 && !a.CoefficientSearch.ClosestCandidateAccepted && a.CoefficientSearch.FittedCoefficientNeeded && !a.CoefficientSearch.ArbitraryNormalizationInserted, Detail: FormatCoefficient(a.CoefficientSearch)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate215Inherited && a.Firewall.ThresholdSpectrumSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.MatchingResidualPromoted && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.SpectralCoefficientTuned && !a.Firewall.HeatKernelMapImported && !a.Firewall.MSbarSchemeImported && !a.Firewall.PhysicalMassPredictionClaimed && !a.Firewall.PhysicalUnificationClaimed && !a.Firewall.ContactModesPromotedToParticles && !a.Firewall.BGapPromotedToMass && !a.Firewall.ProtonLifetimeComputed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{a.TruthStatement, "FAILED_ROUTE is the correct result: Gate 216 finds diagnostics and a sign-only eta resonance, but no canonical spectral heat-kernel matching correction."}}
	}}
}
