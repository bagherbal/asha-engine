package brokengeneratornorm

import (
	"math"
	"testing"
)

func TestNormalizedBrokenGeneratorBasis(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if math.Abs(a.NeutralNormalization-0.5) > 1e-8 {
		t.Fatalf("expected neutral normalization factor 1/2, got %.12f", a.NeutralNormalization)
	}
	if !a.NormalizationExact || math.Abs(a.NormalizedCondition-1) > 1e-8 {
		t.Fatalf("expected exact normalized condition 1, got %.12f", a.NormalizedCondition)
	}
	if math.Abs(a.RawCoordinateKineticCandidate[2]-4) > 1e-8 {
		t.Fatalf("expected raw-coordinate kinetic candidate neutral entry 4, got %.12f", a.RawCoordinateKineticCandidate[2])
	}
	if a.FiniteActionSelectsBasis || a.GaugeKineticHessianSelected || a.PhysicalMassesDerived {
		t.Fatalf("Gate 93 must not claim action-selected Hessian or physical masses")
	}
}
