package fourfermion

import "testing"

func TestNativeFourFermionKernelAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if a.U4Dimension != 16 {
		t.Fatalf("expected dim u(4)=16, got %d", a.U4Dimension)
	}
	if !a.DecompositionComplete {
		t.Fatalf("expected complete u(4) decomposition")
	}
	if !a.CurrentAlgebraAvailable || !a.CurrentCurrentTemplateAvailable || !a.ScalarLRChannelAvailable {
		t.Fatalf("expected current algebra/template/scalar channel to be available")
	}
	if a.FierzProjectionDerived || a.AttractiveChannelSignDerived || a.FourFermionStrengthDerived || a.NativeNJLKernelDerived {
		t.Fatalf("Gate 56 must not claim Fierz projection, attraction, or G_hat")
	}
	if a.UpDownSplittingDerived || a.RegulatorDerived || a.CriticalityClosed {
		t.Fatalf("Gate 56 must not claim top selection or criticality")
	}
	if a.HiddenObservedCouplingsUsed || a.HiddenMassScaleUsed {
		t.Fatalf("must not use observed couplings or mass scales")
	}
}
