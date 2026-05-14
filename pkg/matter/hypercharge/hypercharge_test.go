package hypercharge

import "testing"

func TestScalarHyperchargeBridge(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ScalarDoubletCandidate {
		t.Fatalf("expected 2+2 scalar doublet candidate")
	}
	if a.PairMultiplicity != 2 {
		t.Fatalf("pair multiplicity=%d, want 2", a.PairMultiplicity)
	}
	if a.ScalarDimension != 4 || a.TensorDimension != 64 {
		t.Fatalf("dims scalar=%d tensor=%d", a.ScalarDimension, a.TensorDimension)
	}
	if a.CommutesWithScalarResponseNorm > 1e-8 {
		t.Fatalf("T_phi should commute with pair-degenerate scalar response: %g", a.CommutesWithScalarResponseNorm)
	}
	if a.PatiSalamFlippingDim != 0 {
		t.Fatalf("Pati-Salam diagnostic should expose Γ_F obstruction, flipping=%d", a.PatiSalamFlippingDim)
	}
	if a.RawFlippingDim == 0 {
		t.Fatalf("raw compensator should show scalar charge can create grading-flipping channels")
	}
	if a.StandardModelHyperchargeDerived || a.ElectroweakYukawaDerived {
		t.Fatalf("gate must not claim full SM hypercharge or EW Yukawa")
	}
}
