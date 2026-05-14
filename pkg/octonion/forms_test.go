package octonion

import "testing"

func TestStandardForms(t *testing.T) {
	phi := StandardAssociativeForm()
	if phi.Dimension() != 7 {
		t.Fatalf("unexpected phi dimension: %d", phi.Dimension())
	}
	if got := phi.NonZeroCanonicalTerms(); got != 7 {
		t.Fatalf("expected 7 canonical associative terms, got %d", got)
	}
	if phi.Value(0, 1, 2) != 1 || phi.Value(1, 0, 2) != -1 {
		t.Fatalf("associative form antisymmetry failed")
	}
	if phi.Value(1, 4, 6) != -1 {
		t.Fatalf("Fano convention mismatch for 257 term")
	}

	psi := StandardCoassociativeForm(phi)
	if psi.Dimension() != 7 {
		t.Fatalf("unexpected psi dimension: %d", psi.Dimension())
	}
	if got := psi.NonZeroCanonicalTerms(); got != 7 {
		t.Fatalf("expected 7 canonical coassociative terms, got %d", got)
	}
	if psi.Value(3, 4, 5, 6) == 0 {
		t.Fatalf("expected nonzero coassociative form component")
	}
}
