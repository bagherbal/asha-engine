package protectedconnection

import "testing"

func TestProtectedCarrierOperatorBFContactConnectionSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ProtectedDimension != 3 {
		t.Fatalf("protected dimension=%d, want 3", a.ProtectedDimension)
	}
	if !a.AbstractSO3ConnectionExists {
		t.Fatalf("expected abstract so(3) connection space to exist")
	}
	if !a.DiagonalSpurionAvailable {
		t.Fatalf("expected diagonal generation spurion to be available")
	}
	if !a.ContactCurvatureFlatOnProtected {
		t.Fatalf("expected protected carrier to remain curvature-flat under implemented contact curvature")
	}
	if !a.ActiveCurvatureNonzero {
		t.Fatalf("expected active sector curvature to remain nonzero")
	}
	if a.IntrinsicProtectedOperatorDerived {
		t.Fatalf("intrinsic protected operator should remain open")
	}
	if a.O3FreedomReduced || a.O3FreedomProvenGauge {
		t.Fatalf("O(3) freedom should remain unresolved")
	}
}
