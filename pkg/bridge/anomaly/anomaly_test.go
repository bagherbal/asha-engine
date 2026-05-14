package anomaly

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault error: %v", err)
	}
	if len(a.States) != 16 {
		t.Fatalf("expected 16 Weyl states, got %d", len(a.States))
	}
	if !a.YAnomalyCancels || !a.BMinusLAnomalyCancels || !a.MixedAbelianCancels {
		t.Fatalf("expected abelian anomaly ledgers to cancel")
	}
	if !a.AnomalyShadowSupported {
		t.Fatalf("expected anomaly-shadow interpretation to be supported")
	}
	if a.KineticMixingDerived {
		t.Fatalf("kinetic mixing must not be marked derived")
	}
}
