package continuumdecouplingbridge

import "testing"

func TestBuildDefaultContinuumDecouplingBridgePreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PreviousGate179.Dichotomy.DichotomyCompleteAtCurrentStage || a.PreviousGate179.Dichotomy.ThresholdOriginDerived {
		t.Fatalf("Gate 179 input not in expected open dichotomy state: %s", a.PreviousGate179.Dichotomy.Verdict)
	}
	if a.Preflight.AxiomsAudited < 13 || a.Preflight.AnchorsAudited < 5 {
		t.Fatalf("expected substantial axiom/anchor inventory: %s", FormatPreflight(a.Preflight))
	}
}

func TestAxiomInventoryBlocksPromotion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Preflight
	if p.RequiredHeatKernelAxioms < 8 || p.RequiredThresholdAxioms < 11 {
		t.Fatalf("required axiom counts too small: %s", FormatPreflight(p))
	}
	if p.CanonicalBridgeAxioms != 0 {
		t.Fatalf("no canonical bridge axioms should be derived yet: %s", FormatPreflight(p))
	}
	if p.MissingHeatKernelAxioms != p.RequiredHeatKernelAxioms || p.MissingThresholdAxioms != p.RequiredThresholdAxioms {
		t.Fatalf("all required bridge axioms should still be missing canonically: %s", FormatPreflight(p))
	}
}

func TestFiniteAnchorsAreNotHeatKernelContributions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Preflight.ExactFiniteAnchors < 4 {
		t.Fatalf("expected finite spectral anchors from Gate 179: %s", FormatAnchors(a.Anchors))
	}
	if a.Preflight.PromotableHeatKernelAnchors != 0 || a.Preflight.PromotableThresholdAnchors != 0 || !noAnchorPromoted(a.Anchors) {
		t.Fatalf("anchors should remain predata only: %s", FormatAnchors(a.Anchors))
	}
}

func TestChernWeilTraceAndDecouplingRemainMissing(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.ChernWeilTrace
	if !c.TopologicalSealAvailable || !c.RepresentationTraceRatioClosed {
		t.Fatalf("expected topological/trace predata: %s", FormatChernWeilTrace(c))
	}
	if c.ChernWeilFormDerived || c.OrientedFourCycleDerived || c.PrincipalBundleDerived || c.TraceNormalizationDerived || c.InstantonNumberMapDerived || c.AbsoluteCouplingPromoted {
		t.Fatalf("Chern-Weil bridge should not be promoted: %s", FormatChernWeilTrace(c))
	}
	d := a.DecouplingLaw
	if d.MassUnitDerived || d.ActivationPredicateDerived || d.HeavyLightSplitDerived || d.MatchingScaleDerived || d.ThresholdLogLawDerived || d.NonUniversalDeltaBDerived {
		t.Fatalf("decoupling law should remain absent: %s", FormatDecouplingLaw(d))
	}
}

func TestFirewallNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.ContinuumBridgeDerived || f.HeatKernelMatchingDerived || f.ThresholdCorrectedBetaDerived || f.NonUniversalDeltaBDerived || f.AbsoluteCouplingDerived || f.PhysicalConstantsDerived || f.BoundaryScaleDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
