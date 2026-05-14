package u1source

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FactorizedTraceSourceDerived {
		t.Fatal("expected factorized trace source diagnostic")
	}
	if a.CrossCarrierSourceDerived {
		t.Fatal("cross-carrier source should remain open")
	}
	if a.FactorizedBLContact != 0 {
		t.Fatalf("expected B-L/contact factorized trace to vanish, got %g", a.FactorizedBLContact)
	}
	if a.FactorizedCentralContact != 0 {
		t.Fatalf("expected central/contact factorized trace to vanish, got %g", a.FactorizedCentralContact)
	}
}
