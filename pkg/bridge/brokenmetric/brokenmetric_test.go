package brokenmetric

import (
	"math"
	"testing"
)

func TestBrokenImageMetricNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.RawAnisotropic {
		t.Fatalf("expected raw broken-image metric to be anisotropic")
	}
	if math.Abs(a.NeutralToChargedRatio-4) > 1e-8 {
		t.Fatalf("expected neutral/charged ratio 4, got %.12f", a.NeutralToChargedRatio)
	}
	if math.Abs(a.NeutralNormFactor-0.5) > 1e-8 {
		t.Fatalf("expected neutral normalization factor 1/2, got %.12f", a.NeutralNormFactor)
	}
	if !a.IsotropizationExact {
		t.Fatalf("expected exact isotropization diagnostic, condition %.12f", a.IsotropizedCondition)
	}
	if a.PhysicalAnisotropyDerived || a.GaugeEatingTheoremCompleted {
		t.Fatalf("gate must not claim physical anisotropy or completed gauge eating")
	}
}
