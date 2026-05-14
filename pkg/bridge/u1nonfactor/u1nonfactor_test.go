package u1nonfactor

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NonFactorizedSupportDerived {
		t.Fatalf("expected non-factorized support")
	}
	if !a.LocalNonzeroCorrelation {
		t.Fatalf("expected local nonzero correlation")
	}
	if !a.TotalSignedCancellation {
		t.Fatalf("expected signed cancellation")
	}
	if a.CrossCarrierSourceDerived {
		t.Fatalf("net cross-carrier source must remain open")
	}
}
