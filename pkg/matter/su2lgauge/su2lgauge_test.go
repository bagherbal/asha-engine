package su2lgauge

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Dimension != 8 {
		t.Fatalf("dimension=%d", a.Dimension)
	}
	if !a.NonabelianSU2LGeneratorsDerived {
		t.Fatalf("expected SU(2)L generators")
	}
	if a.CommutatorT3TPlusNorm > 1e-10 {
		t.Fatalf("[T3,T+] residual=%g", a.CommutatorT3TPlusNorm)
	}
	if a.CommutatorT3TMinusNorm > 1e-10 {
		t.Fatalf("[T3,T-] residual=%g", a.CommutatorT3TMinusNorm)
	}
	if a.CommutatorTPlusTMinusNorm > 1e-10 {
		t.Fatalf("[T+,T-] residual=%g", a.CommutatorTPlusTMinusNorm)
	}
	if a.GellMannNishijimaNorm > 1e-10 {
		t.Fatalf("Q identity residual=%g", a.GellMannNishijimaNorm)
	}
}
