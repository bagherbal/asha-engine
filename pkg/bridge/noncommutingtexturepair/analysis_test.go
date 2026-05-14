package noncommutingtexturepair

import "testing"

func TestGate173OperatorInventory(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inventory.OperatorCount != 9 {
		t.Fatalf("expected nine operator candidates, got %d", a.Inventory.OperatorCount)
	}
	if a.Inventory.CanonicalOperators < 6 || a.Inventory.LinearTextureCandidates < 5 {
		t.Fatalf("inventory missed current finite sources: %+v", a.Inventory)
	}
	if a.Inventory.GenerationBreakingOperators != 1 {
		t.Fatalf("only the diagonal Higgs/contact spurion should split generations: %+v", a.Inventory)
	}
}

func TestRawNoncommutationIsNotTexturePair(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inventory.RawNonCommutingPairs == 0 {
		t.Fatalf("expected raw non-commuting triality/spurion maps")
	}
	if !a.NoGo.TrialityRawNoncommutationSeen || !a.NoGo.TrialityRawMapsAreSymmetries {
		t.Fatalf("raw triality noncommutation should be seen and quarantined: %+v", a.NoGo)
	}
	for _, p := range a.Pairs {
		if p.QualifiedNonCommutingPair {
			t.Fatalf("no pair should qualify as a Yukawa texture pair: %+v", p)
		}
	}
}

func TestNoQualifiedTexturePair(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inventory.QualifiedTextureOperators != 0 || a.Inventory.QualifiedNonCommutingPairs != 0 {
		t.Fatalf("qualified texture no-go failed: %+v", a.Inventory)
	}
	if a.Inventory.CanonicalBreakingTextures != 0 || a.Inventory.CanonicalMixingSources != 0 {
		t.Fatalf("canonical breaking/mixing source should not exist: %+v", a.Inventory)
	}
	if !a.NoGo.BFResidualZero || !a.NoGo.SourceTensorMinimumZero || !a.NoGo.DiagonalSpurionRequiresBridge {
		t.Fatalf("expected old Gates 29-36 obstructions to remain active: %+v", a.NoGo)
	}
}

func TestMassFirewallSealedAndGate174Recommended(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NoGo.MassGenerationSealedAtCurrentStage {
		t.Fatalf("mass-generation problem should be sealed at current stage: %+v", a.NoGo)
	}
	if !a.Firewall.GaugeRatioClosed || !a.Firewall.ScalarShapeTargetAvailable || !a.Firewall.MassProblemLocalizedToYukawaMatrix {
		t.Fatalf("upstream positive structure missing: %+v", a.Firewall)
	}
	if a.Firewall.NonCommutingTexturePairFound || a.Firewall.YukawaAmplitudesDerived || a.Firewall.FermionMassesDerived || a.Firewall.CKMPMNSDerived {
		t.Fatalf("mass/mixing should remain open: %+v", a.Firewall)
	}
	if a.Firewall.ResidualNullityBefore != 3 || a.Firewall.ResidualNullityAfter != 3 {
		t.Fatalf("Gate 173 should not change absolute-coupling nullity: %+v", a.Firewall)
	}
	if a.Firewall.RecommendedNextGate == "" {
		t.Fatalf("Gate 174 recommendation missing")
	}
}
