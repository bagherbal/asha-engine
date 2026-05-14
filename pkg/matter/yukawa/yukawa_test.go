package yukawa

import "testing"

func TestIntertwinerSelection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TensorDimension != 64 {
		t.Fatalf("tensor dimension=%d, want 64", a.TensorDimension)
	}
	if a.ChargePreservingDimension != 672 {
		t.Fatalf("charge-preserving dimension=%d, want 672", a.ChargePreservingDimension)
	}
	if a.ChargeChangingDimension != 3424 {
		t.Fatalf("charge-changing dimension=%d, want 3424", a.ChargeChangingDimension)
	}
	if a.OneParticleChargePreservingDimension != 160 {
		t.Fatalf("one-particle preserving dimension=%d, want 160", a.OneParticleChargePreservingDimension)
	}
	if a.ChargeRuleResidual > 1e-8 {
		t.Fatalf("neutral witness commutator residual=%g", a.ChargeRuleResidual)
	}
	if a.PhysicalYukawaTextureDerived || a.MassMatrixDerived {
		t.Fatalf("gate must not claim physical Yukawa texture or mass matrix")
	}
}
