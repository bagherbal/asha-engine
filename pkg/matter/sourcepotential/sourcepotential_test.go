package sourcepotential

import "testing"

func TestSymmetryBreakingSourceActionDoesNotDeriveNonzeroM(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.TensorDimension != 12 || a.GenerationDimension != 3 || a.ActiveDimension != 4 {
		t.Fatalf("unexpected source tensor dimensions: gen=%d active=%d tensor=%d", a.GenerationDimension, a.ActiveDimension, a.TensorDimension)
	}
	if !a.ScalarInvariantsFound {
		t.Fatalf("expected finite scalar invariants to exist")
	}
	if !a.PositiveQuarticFound {
		t.Fatalf("expected positive Higgs/contact quartic data")
	}
	if a.TachyonicSignDerived {
		t.Fatalf("tachyonic source-tensor sign should not be derived yet")
	}
	if a.NonzeroRadiusDerived || a.TensorOrientationFound {
		t.Fatalf("nonzero source tensor radius/orientation should not be derived yet")
	}
	if !a.StableZeroPersists {
		t.Fatalf("stable zero-map result from Gate 35 should persist")
	}
}
