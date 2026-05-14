package faithfuloppositeactionrep

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FaithfulOppositeActionRepresentationNonVacuousOneFormCalculusAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FAITHFUL-OPPOSITE-ACTION-REPRESENTATION-NONVACUOUS-ONE-FORM-CALCULUS-AUDIT"
	const name = "Faithful Opposite-Action Representation / Non-Vacuous One-Form Calculus Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 270 faithful-opposite-action audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 269 order-one sieve is inherited as predecessor", Passed: a.Inheritance.OrderOneDefined && a.Inheritance.ModeLevelSieveReduced && a.Inheritance.SurvivingFamilyDimC == 2 && !a.Inheritance.CanonicalDFDerived && !a.Inheritance.HiggsRatioDerived && a.Inheritance.FirewallPreserved, Detail: FormatInheritance(a.Inheritance)},
			{Name: "faithful full doubled-S_C representation remains missing", Passed: a.Lift.FaithfulLiftAudited() && !a.Lift.FullSCRepresentationDerived && a.Lift.ModePreflightComplexDimension == 8 && a.Lift.TargetComplexDimension == 32 && !a.Lift.ImportedConnesRepresentation, Detail: FormatLift(a.Lift)},
			{Name: "physical opposite action through J is not derived", Passed: a.Opposite.CandidateSwapActionAvailable && !a.Opposite.AntiLinearJDerived && !a.Opposite.OppositeActionDerived && !a.Opposite.CandidateIsPhysical, Detail: FormatOpposite(a.Opposite)},
			{Name: "chiral bimodule diagnostic is constructed without promotion", Passed: a.Chiral.LeftRightActionsDiffer && a.Chiral.NonVacuousPossible && !a.Chiral.FullSCPhysical, Detail: FormatChiral(a.Chiral)},
			{Name: "candidate one-form commutator is non-vacuous on diagnostic probe", Passed: a.OneForm.NonZero && a.OneForm.CentralProbeVanishes && a.OneForm.FrobeniusNormSq == 2 && !a.OneForm.PhysicalOneForm, Detail: FormatOneForm(a.OneForm)},
			{Name: "candidate non-vacuous action fails full order-one residual", Passed: !a.Residual.CandidatePasses && !a.Residual.FullOrderOneProved && a.Residual.FrobeniusNormSq == 1, Detail: FormatResidual(a.Residual)},
			{Name: "spectral moment ratio remains x:y dependent", Passed: !a.Ratio.XToYSelected && !a.Ratio.RatioStableAcrossFamily && a.Ratio.DependsOnXY && !a.Ratio.GaugeProjectionDerived && !a.Ratio.ScalarFluctuationMapDerived && !a.Ratio.HiggsRatioDerived, Detail: FormatRatio(a.Ratio)},
			{Name: "firewall keeps diagnostic representation from becoming a theorem", Passed: a.Firewall.EmpiricalYukawaSealPreserved && a.Firewall.SpontaneousCarrierSealPreserved && a.Firewall.NoObservedMassInserted && a.Firewall.NoVEVInserted && a.Firewall.NoCutoffScaleInserted && a.Firewall.NoConnesRepresentationImported && a.Firewall.CandidateNotPromoted && a.Firewall.NoHiggsPredictionClaim && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future map records the exact missing spectral-triple ingredients", Passed: len(a.Future.Obligations) >= 6 && a.Future.NeedFullSCRepresentation && a.Future.NeedPhysicalJ && a.Future.NeedOrderOnePassingBimod && a.Future.NeedNonVacuousCalculus && a.Future.NeedCanonicalXYSelector && a.Future.NeedHeatKernelProjection, Detail: FormatFuture(a.Future)},
			{Name: "summary records non-vacuous diagnostic progress and failed invariant-Higgs route", Passed: a.Summary.Gate269Inherited && a.Summary.FaithfulLiftAudited && !a.Summary.FullSCRepresentation && !a.Summary.PhysicalOppositeAction && a.Summary.CandidateOneFormsNonzero && !a.Summary.CandidateOrderOnePasses && !a.Summary.CanonicalDFDerived && !a.Summary.RatioStable && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 270 shows that a chiral mismatch can make [D_F,a] nonzero, but the tested mismatch is only a mode-level diagnostic and fails the full order-one residual.",
			"A physical spectral-action theorem still requires a faithful doubled-S_C representation, a derived anti-linear J/opposite action, and nonzero one-forms satisfying order-one simultaneously.",
		}}
	}}
}

func (a FaithfulLiftAudit) FaithfulLiftAudited() bool {
	return a.TargetComplexDimension == 32 && a.ModePreflightComplexDimension == 8 && a.ChiralGradingRespected
}
