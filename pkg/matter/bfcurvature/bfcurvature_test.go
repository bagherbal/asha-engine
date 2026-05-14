package bfcurvature

import "testing"

func TestFiniteMaurerCartanCurvatureBuilds(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GeneratorCount == 0 || a.SeedSpanRank == 0 {
		t.Fatalf("invalid seed: %+v", a)
	}
	if a.FullDimension != 56 {
		t.Fatalf("expected 56D Boolean support, got %d", a.FullDimension)
	}
	if a.ProtectedDimension != 3 {
		t.Fatalf("expected 3D protected carrier, got %d", a.ProtectedDimension)
	}
	if a.ActiveDimension != 4 {
		t.Fatalf("expected 4D active carrier, got %d", a.ActiveDimension)
	}
}
