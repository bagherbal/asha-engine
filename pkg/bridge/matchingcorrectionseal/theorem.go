package matchingcorrectionseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

func MatchingCorrectionSealFullSMYukawaTwoLoopIntegrationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-MATCHING-CORRECTION-SEAL-FULL-SM-YUKAWA-TWO-LOOP"
	const name = "MatchingCorrectionSeal / full SM Yukawa 2-loop integration audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 218 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 217 spectral-action obstruction is inherited", Passed: a.Gate217.Gate217Inherited && a.Gate217.FiniteSpectralTripleFailed && a.Gate217.MatchingResidualTargetInherited && a.Gate217.ThresholdSpectrumSealInherited && a.Gate217.MissingFiniteDiracOperator && a.Gate217.MissingGaugeCurvatureProjection && a.Gate217.MissingCutoffSubtractionScheme, Detail: FormatGate217(a.Gate217)},
			{Name: "MatchingCorrectionSeal quarantines the required residual", Passed: a.Seal.Active && a.Seal.ID == MatchingCorrectionSealID && !a.Seal.ResidualPromotedAsDerived && !a.Seal.FiniteSpectralTripleDerived, Detail: FormatSeal(a.Seal)},
			{Name: "empirical top and Higgs seeds are explicit phenomenological inputs", Passed: !a.Inputs.FiniteCoreDerived && a.Inputs.TopPoleMassGeV > 170 && a.Inputs.HiggsMassGeV > 124 && a.Inputs.InitialYTop > 0.9 && a.Inputs.InitialLambda > 0.12, Detail: FormatInputs(a.Inputs)},
			{Name: "full SM top-Yukawa gauge terms and Higgs-quartic running are included", Passed: a.Config.YukawaTermsIncluded && a.Config.LambdaRunningIncluded && a.FullSM.YTopBetaIncluded && a.FullSM.LambdaBetaIncluded && !a.FullSM.SMYukawaMatricesDerived && !a.FullSM.HeavySectorYukawasAdded, Detail: FormatConfig(a.Config) + " :: " + FormatFullSM(a.FullSM)},
			{Name: "sealed single-scale heavy spectrum remains the Gate-215 survivor", Passed: a.Spectrum.ConditionalOnly && a.Spectrum.Row1Rep == "(1,3,Y=1)" && a.Spectrum.Row2Rep == "(8,2,Y=1/2)" && a.Spectrum.TotalDeltaB.U1GUT > 5 && a.Spectrum.TotalDeltaB.SU3C == 8, Detail: FormatSpectrum(a.Spectrum)},
			{Name: "forced full-SM-Yukawa degenerate fit converges and remains inside the matching envelope", Passed: a.Fit.Converged && a.Fit.ScaleOrdered && a.Fit.SubPlanck && a.Fit.PositiveToBoundary && a.Fit.NoLandauBelowPlanck && a.Fit.MatchingPlausible && a.Fit.ResidualOverEpsilon < 1.0, Detail: FormatFit(a.Fit)},
			{Name: "Yukawa/scalar sector shifts the target but does not derive it", Passed: a.Comparison.Gate218ResidualOverEpsilon > 0 && a.Comparison.PlausibilityPreserved && a.Comparison.ResidualShiftMaxAbs > 0 && !a.Summary.MatchingCorrectionsDerived, Detail: FormatComparison(a.Comparison)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate217Inherited && a.Firewall.MatchingCorrectionSealActive && a.Firewall.RequiredResidualQuarantined && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.MatchingResidualPromoted && a.Firewall.ThresholdSpectrumSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.TopMassFiniteDerived && !a.Firewall.HiggsMassFiniteDerived && !a.Firewall.SMYukawaMatricesFiniteDerived && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.ProtonLifetimeComputed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 218 seals δ_match and includes top-Yukawa/Higgs-quartic running as empirical SM input; it does not derive matching corrections, masses, or Yukawa textures from the finite core."}}
	}}
}
