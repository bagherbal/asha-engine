package gaugeeating

import "testing"

func TestGaugeEatingDiagnostic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.BrokenImageRank != 3 {
		t.Fatalf("expected 3 broken image directions, got %d", a.BrokenImageRank)
	}
	if !a.GoldstoneImageTheoremDiagnostic {
		t.Fatalf("expected Goldstone image diagnostic to hold")
	}
	if a.EMNullNorm > 1e-9 {
		t.Fatalf("expected electromagnetic null generator, norm=%g", a.EMNullNorm)
	}
	if a.KineticNormalizationSelected {
		t.Fatalf("kinetic normalization should not be selected yet")
	}
	if a.FiniteGaugeEatingTheoremDerived {
		t.Fatalf("finite gauge-eating theorem should remain bridge-required")
	}
}
