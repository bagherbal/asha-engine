package finitethresholdoperator

import "testing"

func TestBuildDefaultFiniteThresholdOperator(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Candidates) < 8 {
		t.Fatalf("expected broad threshold candidate inventory, got %d", len(a.Candidates))
	}
	if !a.PreviousGate177.Firewall.NonUniversalThresholdCanFitByConstruction || a.PreviousGate177.Firewall.NonUniversalThresholdDerived {
		t.Fatalf("Gate 177 non-universal branch should remain fit-only and underived")
	}
}

func TestNoCandidateHasCompleteThresholdChain(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Requirements
	if !r.NoCandidateHasAllPieces || r.CompleteThresholdOperators != 0 || r.FiniteDerivedThresholdOps != 0 {
		t.Fatalf("expected no complete finite threshold operator: %s", FormatRequirements(r))
	}
	if r.WithFiniteSpectrum < 4 || r.WithGaugeRepresentation < 3 {
		t.Fatalf("audit should see partial data but no full chain: %s", FormatRequirements(r))
	}
	if r.WithPhysicalMassUnit != 0 || r.WithActivationPredicate != 0 || r.WithDecouplingRule != 0 {
		t.Fatalf("mass/activation/decoupling should remain missing: %s", FormatRequirements(r))
	}
}

func TestBaselineRowNotHeavyThreshold(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BetaMatching.ScalarSectorRowConstructed || !a.BetaMatching.ScalarSectorMatchesBaseline {
		t.Fatalf("expected scalar baseline row from beta matching")
	}
	if a.BetaMatching.ScalarSectorIsThresholdCorrection || a.BetaMatching.BetaCorrectionRowsAllowed != 0 {
		t.Fatalf("baseline row must not become a heavy threshold correction")
	}
	if a.Requirements.BaselineRowsAlreadyCounted != 1 {
		t.Fatalf("expected exactly one baseline row in candidate audit: %s", FormatRequirements(a.Requirements))
	}
}

func TestFiniteSpectraRemainOpenAnchors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Requirements.OpenFiniteSpectrumAnchors < 4 {
		t.Fatalf("expected multiple open finite spectral anchors: %s", FormatRequirements(a.Requirements))
	}
	if a.BetaMatching.BGapRepresentationCompleted || a.BetaMatching.ContactOverlapRepresentationCompleted {
		t.Fatalf("B/contact representation rows should remain incomplete")
	}
	if a.BetaMatching.ActivationRuleDerived || a.BetaMatching.DecouplingMatchingRuleDerived {
		t.Fatalf("activation/decoupling must remain underived")
	}
}

func TestCombinationAttemptsRejected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Combinations) < 5 {
		t.Fatalf("expected at least five combination attempts")
	}
	if !noStrictCombination(a.Combinations) {
		t.Fatalf("no combination should be admissible as a strict theorem: %s", FormatCombinations(a.Combinations))
	}
	if !hasObservedCombination(a.Combinations) {
		t.Fatalf("expected one observed-fit combination to be explicitly quarantined")
	}
}

func TestDeltaBWitnessQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	d := a.DeltaBWitness
	if !d.Gate177NonUniversalFitExists || d.Gate177FiniteThresholdDerived {
		t.Fatalf("expected Gate 177 non-universal fit-only branch: %s", FormatDeltaBWitness(d))
	}
	if !d.MinimumNormWitnessUsesExternalFit || d.CanBePromotedToFiniteOperator {
		t.Fatalf("Δb witness must remain external and non-promotable: %s", FormatDeltaBWitness(d))
	}
}

func TestFirewallNoThresholdNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.ThresholdOperatorDerived || f.NonUniversalDeltaBDerived || f.ThresholdCorrectedBetaDerived || f.Gate177RepairPromoted {
		t.Fatalf("threshold firewall violation: %s", FormatFirewall(f))
	}
	if f.UsesObservedInputForDerivation || f.PhysicalConstantsDerived || f.BoundaryScaleDerived {
		t.Fatalf("hidden physical input/constant derivation violation: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
