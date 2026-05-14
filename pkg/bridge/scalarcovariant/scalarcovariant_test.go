package scalarcovariant

import "testing"

func TestFiniteScalarCovariantDerivativeAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DimensionlessWZPhotonSignature {
		t.Fatalf("expected W/Z/photon signature")
	}
	if a.MassMatrixRank != 3 {
		t.Fatalf("expected rank 3 mass matrix, got %d", a.MassMatrixRank)
	}
	if a.EMAnnihilatesVacuumNorm > 1e-9 {
		t.Fatalf("EM generator should annihilate diagnostic vacuum: %.3e", a.EMAnnihilatesVacuumNorm)
	}
	if a.PhysicalMassesDerived {
		t.Fatalf("Gate 84 must not derive physical masses")
	}
}
