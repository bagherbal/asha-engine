package sectorspectrum

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Sectors) != 4 {
		t.Fatalf("expected four current sectors, got %d", len(a.Sectors))
	}
	if a.AssignedSectorCount != 0 {
		t.Fatalf("expected no assigned current sectors yet")
	}
	if a.PropagatorDenominatorsDerived {
		t.Fatalf("propagator denominators must remain open")
	}
	if a.CondensationClaimAllowed {
		t.Fatalf("condensation claim must not be allowed")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("must not use observed input")
	}
}
