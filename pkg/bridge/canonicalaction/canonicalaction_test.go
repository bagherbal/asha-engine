package canonicalaction

import (
	"math"
	"testing"
)

func TestCanonicalFiniteVariationalAction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.ScalarKineticSelected {
		t.Fatalf("scalar kinetic normalization was not selected")
	}
	if !a.BrokenSecondVariationSelected {
		t.Fatalf("broken second variation should select diag(1,1,4), got %v", a.BrokenSelectedDiagonal)
	}
	if math.Abs(a.KappaU1-6) > 1e-8 {
		t.Fatalf("kappa_U1=%g, want 6", a.KappaU1)
	}
	if !a.FullGaugeHessianSelected || !a.FullGaugeHessianPositive || a.FullGaugeHessianRank != 4 {
		t.Fatalf("full gauge Hessian not selected/positive/rank4")
	}
	if !a.GenerationSourceSelected || a.GenerationSource.ProducesMixing {
		t.Fatalf("generation source should be selected as diagonal non-mixing spectrum")
	}
	if a.ActiveToGenerationMixingSelected || !a.SourceAction.NaturalSelectsZero {
		t.Fatalf("active-to-generation mixing source must remain sealed/zero")
	}
	if !a.CanonicalActionSelected {
		t.Fatalf("canonical action should be selected")
	}
	if a.PhysicalCouplingsDerived || a.PhysicalMassesDerived || a.CKMPMNSDerived {
		t.Fatalf("physical couplings/masses/mixing must remain unclaimed")
	}
}
