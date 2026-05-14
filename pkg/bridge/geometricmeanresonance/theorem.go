package geometricmeanresonance

import "github.com/bagherbal/asha-engine/pkg/theorem"

func GeometricMeanIntermediateResonanceAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-GEOMETRIC-MEAN-INTERMEDIATE-RESONANCE-AUDIT"
	const name = "Geometric-mean intermediate scale resonance / sealed hierarchy audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 227 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 226 sealed axion phenomenology is inherited", Passed: a.Gate226.Gate226Inherited && a.Gate226.AxionSealActive && a.Gate226.NativeAxionNotDerived && a.Gate226.FAGeV > 0 && a.Gate226.HeavyDMAbsenceBinding, Detail: FormatGate226(a.Gate226)},
			{Name: "Gate 223 relic-decay EFT scale is inherited", Passed: a.Gate223.Gate223Inherited && a.Gate223.RelicDecaySealGranted && a.Gate223.OctetPortalFound && a.Gate223.LambdaEFTMaxGeV > 0 && a.Gate223.LeptoquarkDynamicsSealAlive && !a.Gate223.WilsonCoefficientDerived, Detail: FormatGate223(a.Gate223)},
			{Name: "sealed hierarchy values are used without new fitting", Passed: a.Hierarchy.ValuesInheritedOnly && !a.Hierarchy.NativeScaleDerived && a.Hierarchy.MBGeV > 0 && a.Hierarchy.MStarGeV > a.Hierarchy.MBGeV && a.Hierarchy.FARequirementGeV > 0 && a.Hierarchy.LambdaEFTMaxGeV > 0, Detail: FormatHierarchy(a.Hierarchy)},
			{Name: "geometric mean brackets the two independent intermediate scales", Passed: a.Geometric.SymmetricSeesaw && a.Geometric.FAMatch && a.Geometric.LambdaMatch && a.Geometric.BothTargetsBracketed && a.Geometric.NullHypothesisRejected, Detail: FormatGeometric(a.Geometric)},
			{Name: "two-step seesaw pattern is suggested but not natively derived", Passed: a.Seesaw.TwoStepPatternSuggested && a.Seesaw.AxionScaleCanLiveThere && a.Seesaw.RelicDecayScaleCanLiveThere && !a.Seesaw.NativeBreakingPotentialDerived && !a.Seesaw.NativeOrderParameterDerived && !a.Seesaw.NativeIntermediateScaleDerived, Detail: FormatSeesaw(a.Seesaw)},
			{Name: "Pati-Salam/u4 route remains quarantined by leptoquark dynamics seal", Passed: a.PatiSalam.DormantU4SlotsPresent && a.PatiSalam.LeptoquarkDynamicsSealActive && a.PatiSalam.IntermediateU4BreakingAudited && !a.PatiSalam.NativeU4GaugeConnectionDerived && !a.PatiSalam.NativeLeptoquarkCurvature && !a.PatiSalam.ProtonDecayChannelReopened && a.PatiSalam.ConsistentOnlyWhileSealed && !a.PatiSalam.LifetimeComputed, Detail: FormatPatiSalam(a.PatiSalam)},
			{Name: "null hypothesis is explicitly tested", Passed: a.Null.Tested && !a.Null.PassedNoResonance && a.Null.WorstGapDecades < 1.0, Detail: FormatNull(a.Null)},
			{Name: "firewalls remain closed", Passed: a.Firewall.Gate226Inherited && a.Firewall.Gate223Inherited && a.Firewall.UsedOnlySealedValues && !a.Firewall.IntermediateScaleFiniteDerived && !a.Firewall.AxionNativeDerived && !a.Firewall.EFTMediatorDerived && !a.Firewall.PatiSalamImportedAsTheorem && !a.Firewall.LeptoquarkSealViolated && !a.Firewall.ProtonLifetimeComputed && !a.Firewall.BGapPromotedWithoutSeal && !a.Firewall.NewPhenomenologicalFitAdded && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 227 finds a strong geometric-mean resonance among sealed scales, but it does not derive the intermediate breaking mechanism, axion, EFT mediator, or Pati-Salam dynamics from the finite core."}}
	}}
}
