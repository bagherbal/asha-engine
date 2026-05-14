package scalarsu2

import "testing"

func TestScalarContactSU2ActionSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.AbstractDoubletRepresentation {
		t.Fatalf("expected abstract SU(2) doublet representation")
	}
	if a.FullSU2SelectedByScalarData {
		t.Fatalf("full SU(2) should not be selected by anisotropic scalar response")
	}
	if !a.U1PairRotationSelected {
		t.Fatalf("expected T3/U(1) pair-rotation subgroup to commute with scalar response")
	}
	if a.ScalarResponseCommT3Norm > 1e-9 {
		t.Fatalf("expected T3 to commute with scalar response, got %g", a.ScalarResponseCommT3Norm)
	}
	if a.ScalarResponseCommT1Norm <= 1e-6 || a.ScalarResponseCommT2Norm <= 1e-6 {
		t.Fatalf("expected T1/T2 to expose anisotropic noncommutation, got %g/%g", a.ScalarResponseCommT1Norm, a.ScalarResponseCommT2Norm)
	}
}
