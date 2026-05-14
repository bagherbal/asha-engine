package sectoroperators

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Generators) != 16 {
		t.Fatalf("expected 16 generators, got %d", len(a.Generators))
	}
	if len(a.Operators) != 4 {
		t.Fatalf("expected four sector operators, got %d", len(a.Operators))
	}
	if !a.RepresentationLevelMapsDerived {
		t.Fatalf("expected sector Casimir maps to be constructed")
	}
	if !a.OperatorsPositive {
		t.Fatalf("expected positive sector operators")
	}
	if a.PropagatorDenominatorsDerived {
		t.Fatalf("propagator denominators must remain open")
	}
	if a.CondensationClaimAllowed {
		t.Fatalf("condensation claim must remain forbidden")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("must not use observed inputs")
	}
}

func TestSectorCasimirValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ColorCasimirValue < 5.3333 || a.ColorCasimirValue > 5.3334 {
		t.Fatalf("unexpected color Casimir value %.12f", a.ColorCasimirValue)
	}
	if a.LeptoquarkLeptonValue != 6 || a.LeptoquarkColorValue != 2 {
		t.Fatalf("unexpected leptoquark spectrum lepton=%.12f color=%.12f", a.LeptoquarkLeptonValue, a.LeptoquarkColorValue)
	}
	if a.BLLeptonValue != 1 {
		t.Fatalf("unexpected B-L lepton square %.12f", a.BLLeptonValue)
	}
}
