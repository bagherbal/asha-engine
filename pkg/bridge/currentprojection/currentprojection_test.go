package currentprojection

import "testing"

func TestCurrentActionScalarProjection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.CurrentActionConstructed {
		t.Fatalf("current action not constructed")
	}
	if len(a.Generators) != 16 {
		t.Fatalf("expected 16 current generators, got %d", len(a.Generators))
	}
	if !a.UnsignedScalarProjectionCoefficientsKnown {
		t.Fatalf("expected finite overlap diagnostics")
	}
	if a.SignedScalarProjectionCoefficientsKnown {
		t.Fatalf("signed coefficients should remain open")
	}
	if a.UpDownSplittingDerived {
		t.Fatalf("up/down splitting should remain open")
	}
}
