package sourcemap

import "testing"

func TestSourceTensorSelectionTruth(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GenerationDimension != 3 || a.ActiveDimension != 4 {
		t.Fatalf("unexpected source tensor domain: %d x %d", a.GenerationDimension, a.ActiveDimension)
	}
	if a.MapSpaceDimension != 12 {
		t.Fatalf("expected Hom dimension 12, got %d", a.MapSpaceDimension)
	}
	if a.CanonicalSourceTensorFound {
		t.Fatal("current finite data should not yet select a canonical active-to-generation source tensor")
	}
	if !a.ArbitraryMapsExist {
		t.Fatal("the abstract 3x4 map space should exist even though it is not canonical")
	}
}
