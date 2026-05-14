package continuumdecouplingbridge

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContinuumDecouplingBridgeAxiomInventoryHeatKernelPreflightTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTINUUM-DECOUPLING-BRIDGE-AXIOM-INVENTORY-HEAT-KERNEL-PREFLIGHT"
	const name = "continuum decoupling bridge axiom inventory / finite heat-kernel matching preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build continuum decoupling bridge preflight", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 179 dichotomy is consumed without promoting either origin branch", Passed: a.PreviousGate179.Dichotomy.DichotomyCompleteAtCurrentStage && !a.PreviousGate179.Dichotomy.ThresholdOriginDerived && !a.PreviousGate179.Dichotomy.Gate177RepairPromoted && a.PreviousGate179.Firewall.StrictNullityAfter == 3 && a.PreviousGate179.Firewall.ConditionalNullityAfter == 2, Detail: a.PreviousGate179.Firewall.Verdict},
			{Name: "continuum/heat-kernel axiom inventory is explicit and blocking", Passed: a.Preflight.AxiomsAudited >= 13 && a.Preflight.RequiredHeatKernelAxioms >= 8 && a.Preflight.RequiredThresholdAxioms >= 11 && a.Preflight.CanonicalBridgeAxioms == 0 && a.Preflight.MissingHeatKernelAxioms == a.Preflight.RequiredHeatKernelAxioms && a.Preflight.MissingThresholdAxioms == a.Preflight.RequiredThresholdAxioms, Detail: FormatPreflight(a.Preflight)},
			{Name: "exact finite anchors remain predata, not heat-kernel coefficients", Passed: a.Preflight.ExactFiniteAnchors >= 4 && a.Preflight.PromotableHeatKernelAnchors == 0 && a.Preflight.PromotableThresholdAnchors == 0 && noAnchorPromoted(a.Anchors), Detail: FormatAnchors(a.Anchors)},
			{Name: "Chern-Weil and trace-normalization bridge is not derived", Passed: a.ChernWeilTrace.TopologicalSealAvailable && a.ChernWeilTrace.RepresentationTraceRatioClosed && !a.ChernWeilTrace.ChernWeilFormDerived && !a.ChernWeilTrace.OrientedFourCycleDerived && !a.ChernWeilTrace.PrincipalBundleDerived && !a.ChernWeilTrace.TraceNormalizationDerived && !a.ChernWeilTrace.InstantonNumberMapDerived && !a.ChernWeilTrace.AbsoluteCouplingPromoted, Detail: FormatChernWeilTrace(a.ChernWeilTrace)},
			{Name: "decoupling and matching law is absent", Passed: !a.DecouplingLaw.MassUnitDerived && !a.DecouplingLaw.ActivationPredicateDerived && !a.DecouplingLaw.HeavyLightSplitDerived && !a.DecouplingLaw.MatchingScaleDerived && !a.DecouplingLaw.ThresholdLogLawDerived && !a.DecouplingLaw.NonUniversalDeltaBDerived, Detail: FormatDecouplingLaw(a.DecouplingLaw)},
			{Name: "heat-kernel coefficients and threshold beta rows remain sealed", Passed: !a.Preflight.A0CoefficientDerived && !a.Preflight.A2CoefficientDerived && !a.Preflight.A4GaugeCoefficientDerived && !a.Preflight.FiniteHeatKernelMatchingDerived && !a.Preflight.ContinuumDecouplingBridgeDerived && !a.Preflight.Gate177RepairPromoted, Detail: FormatPreflight(a.Preflight)},
			{Name: "nullity firewall remains unchanged", Passed: !a.Firewall.UsesObservedInputForDerivation && !a.Firewall.ContinuumBridgeDerived && !a.Firewall.HeatKernelMatchingDerived && !a.Firewall.ThresholdCorrectedBetaDerived && !a.Firewall.NonUniversalDeltaBDerived && !a.Firewall.AbsoluteCouplingDerived && !a.Firewall.PhysicalConstantsDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && a.Firewall.ConditionalNullityBefore == 2 && a.Firewall.ConditionalNullityAfter == 2, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"Gate 180 does not reject future heat-kernel matching. It proves the current finite spectra lack the geometric/analytic bridge required to serve as heat-kernel or decoupling data.",
			"The next minimal construction target is the oriented four-cycle / Chern-Weil carrier, because without it neither S_top nor finite spectra can be integrated as continuum gauge action data.",
		}}
	}}
}

func noAnchorPromoted(xs []AnchorHeatKernelAudit) bool {
	for _, x := range xs {
		if x.CanContributeToHeatKernel || x.CanGenerateNonUniversalDeltaB || x.SeeleyDeWittCoefficients || x.DecouplingMatchingLaw {
			return false
		}
	}
	return true
}
