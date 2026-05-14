package u1kinetic

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.ChargeLevelHyperchargeSelected {
		t.Fatalf("expected charge-level hypercharge selected")
	}
	if !a.MatterGramDerived || !a.CentralBMinusLOrthogonal {
		t.Fatalf("expected matter Gram block with central/B-L orthogonality")
	}
	if a.MatterGram.Rows() != 2 || a.MatterGram.Cols() != 2 {
		t.Fatalf("unexpected matter Gram size %dx%d", a.MatterGram.Rows(), a.MatterGram.Cols())
	}
	if a.BlockDiagonalDiagnostic.Rows() != 3 || a.BlockDiagonalDiagnostic.Cols() != 3 {
		t.Fatalf("unexpected block diagnostic size %dx%d", a.BlockDiagonalDiagnostic.Rows(), a.BlockDiagonalDiagnostic.Cols())
	}
	if a.FullU1KineticHessianDerived {
		t.Fatalf("full physical U(1) kinetic Hessian should remain open")
	}
	if a.PhysicalU1CouplingDerived || a.FineStructureDerived {
		t.Fatalf("physical U(1) coupling and alpha should remain open")
	}
}
