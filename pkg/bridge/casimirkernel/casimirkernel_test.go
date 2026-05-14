package casimirkernel

import "testing"

func TestCurrentSectorCasimirDiagnostics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AllCasimirDiagnosticsBuilt || len(a.Diagnostics) != 4 {
		t.Fatalf("expected four diagnostics, got %d", len(a.Diagnostics))
	}
	if !a.ColorSectorZeroMode {
		t.Fatalf("expected color sector to expose the lepton zero mode")
	}
	if a.PropagatorDenominatorsDerived || a.ExchangeKernelUpdated || a.CondensationClaimAllowed {
		t.Fatalf("gate must not derive propagators, exchange kernels, or condensation")
	}
	if a.DominantDirectSector == "" || a.DominantInverseSector == "" {
		t.Fatalf("expected direct and inverse diagnostic dominance to be exposed")
	}
}
