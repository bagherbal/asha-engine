package protectedintertwiner

import "testing"

func TestProtectedContactBrokenGeneratorIntertwinerSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.CountLevelResonance {
		t.Fatalf("expected protected/contact/broken 3D count resonance")
	}
	if !a.AbstractIsometryExists {
		t.Fatalf("expected an abstract 3D isometry family to exist")
	}
	if a.CanonicalIntertwinerDerived {
		t.Fatalf("canonical protected-to-broken intertwiner should remain open")
	}
	if a.AbstractIsometryFreedomDimension != 3 {
		t.Fatalf("expected O(3) freedom dimension 3, got %d", a.AbstractIsometryFreedomDimension)
	}
}
