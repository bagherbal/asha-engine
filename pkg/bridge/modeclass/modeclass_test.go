package modeclass

import "testing"

func TestModeClassClassifierKeepsThresholdsSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ResidualNullityAfter != 3 || a.ResidualNullityAfter != a.ResidualNullityBefore {
		t.Fatalf("unexpected residual nullity: before=%d after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)
	}
	if !a.BGapClassifiedAsConstrainedFinite || !a.BGapExcludedFromThresholdBeta {
		t.Fatalf("B-sector gap was not safely classified/excluded")
	}
	if a.ContactOpenRows != 7 || a.ContactOverlapClassDerived {
		t.Fatalf("contact overlap modes should remain seven open modes, got open=%d derived=%t", a.ContactOpenRows, a.ContactOverlapClassDerived)
	}
	if a.PhysicalHeavyThresholdRows != 0 || a.BetaCorrectionRowsAllowed != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("threshold beta physics leaked through: physical=%d betaAllowed=%d thresholdBeta=%t", a.PhysicalHeavyThresholdRows, a.BetaCorrectionRowsAllowed, a.ThresholdCorrectedBetaDerived)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical prediction or hidden input leaked through")
	}
}

func TestAmbiguityWitnessesContainBetaChangingBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	foundChanging := false
	for _, w := range a.AmbiguityWitnesses {
		if w.WouldCorrectBeta && w.CompatibleWithFiniteData {
			foundChanging = true
		}
	}
	if !foundChanging {
		t.Fatalf("expected a compatible beta-changing ambiguity witness")
	}
}
