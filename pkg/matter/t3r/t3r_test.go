package t3r

import "testing"

func TestMatterT3RSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.MatterSideOperatorFound {
		t.Fatalf("expected matter-side operator")
	}
	if a.Vectorlike.FlippingAvailable {
		t.Fatalf("vectorlike temporal polarization should not unlock chirality flipping")
	}
	if !a.ChiralEven.FlippingAvailable || !a.ChiralOdd.FlippingAvailable {
		t.Fatalf("expected both chiral restrictions to unlock flipping: even=%d odd=%d", a.ChiralEven.FlippingDim, a.ChiralOdd.FlippingDim)
	}
	if a.PhysicalOrientationSelected {
		t.Fatalf("orientation must remain open")
	}
}
