package kineticnorm

import "testing"

func TestGeneratorKineticNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.KineticTraceNormalizationDerived {
		t.Fatalf("expected finite kinetic trace normalization")
	}
	if !a.UnsignedUnitProjectionCoefficientsDerived {
		t.Fatalf("expected unsigned unit projection coefficients")
	}
	if a.SignedCliffordFierzCoefficientsDerived {
		t.Fatalf("signed Fierz coefficients must remain open")
	}
	if a.NativeFourFermionKernelDerived {
		t.Fatalf("native four-fermion kernel must remain open")
	}
	if len(a.SectorNormalizations) != 4 {
		t.Fatalf("expected 4 current sectors, got %d", len(a.SectorNormalizations))
	}
	if a.TotalKineticTrace <= 0 {
		t.Fatalf("expected positive total kinetic trace")
	}
}
