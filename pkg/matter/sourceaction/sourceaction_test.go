package sourceaction

import "testing"

func TestSourceTensorActionSelectsZeroWithoutGeometricSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.GenerationDimension != 3 || a.ActiveDimension != 4 || a.TensorDimension != 12 {
		t.Fatalf("unexpected tensor dimensions: gen=%d active=%d dim=%d", a.GenerationDimension, a.ActiveDimension, a.TensorDimension)
	}
	if !a.NaturalHessianPositive || !a.NaturalUniqueMinimum {
		t.Fatalf("minimal source action should be stable and uniquely minimized")
	}
	if !a.NaturalSelectsZero {
		t.Fatalf("expected current finite data to select zero source tensor, sourceNorm=%g", a.NaturalSourceNorm)
	}
	if a.NonzeroStationaryFound {
		t.Fatalf("nonzero stationary source tensor should not be derived yet")
	}
	if !a.ArbitraryMapRejected {
		t.Fatalf("arbitrary map space must be rejected as a derivation")
	}
}
