package gencurvature

import "testing"

func TestCurvatureOnGenerationCarrier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.CarrierDimension != 3 {
		t.Fatalf("carrier dimension = %d, want 3", a.CarrierDimension)
	}
	if a.ActiveDimension != 4 {
		t.Fatalf("active dimension = %d, want 4", a.ActiveDimension)
	}
	if len(a.Operators) != 6 {
		t.Fatalf("protected operators = %d, want 6", len(a.Operators))
	}
	if len(a.ActiveOperators) != 6 {
		t.Fatalf("active operators = %d, want 6", len(a.ActiveOperators))
	}
	if a.MaxCurvatureNorm > 1e-8 {
		t.Fatalf("protected carrier should be curvature-flat at this stage, got %.3e", a.MaxCurvatureNorm)
	}
	if a.ActiveMaxCurvatureNorm <= 1e-8 {
		t.Fatalf("expected nonzero curvature on active Higgs/contact carrier")
	}
}
