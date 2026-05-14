package ewcurvature

import "testing"

func TestFullElectroweakCurvatureAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Closed {
		t.Fatalf("full electroweak connection should close")
	}
	if a.AdjointRank != 3 {
		t.Fatalf("expected rank-3 adjoint metric, got %d", a.AdjointRank)
	}
	if a.Diag114SelectedByCurvature {
		t.Fatalf("diag(1,1,4) must not be selected by curvature alone")
	}
	if a.U1KineticSelected {
		t.Fatalf("U(1) kinetic normalization must remain open")
	}
}
