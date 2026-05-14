package propagatorspectrum

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SpectralDenominatorsAvailable {
		t.Fatalf("expected spectral denominators to be available")
	}
	if len(a.Families) == 0 {
		t.Fatalf("expected denominator families")
	}
	if a.PropagatorDenominatorsDerived {
		t.Fatalf("propagator denominators must remain open")
	}
	if a.CondensationClaimAllowed {
		t.Fatalf("condensation claim must not be allowed")
	}
	if a.StrongestDiagnosticKernel <= 0 {
		t.Fatalf("expected positive diagnostic kernel")
	}
}
