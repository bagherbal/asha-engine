package goldstone

import "testing"

func TestGoldstoneAuditCountLevelResonance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ActiveRealDirections != 4 || a.RadialDirections != 1 || a.ScalarAngularDirections != 3 {
		t.Fatalf("unexpected scalar split: active=%d radial=%d angular=%d", a.ActiveRealDirections, a.RadialDirections, a.ScalarAngularDirections)
	}
	if !a.GoldstoneCountResonance {
		t.Fatal("expected count-level Goldstone/electroweak resonance")
	}
	if a.GaugeEatingTheoremDerived {
		t.Fatal("gate must not claim the full gauge-eating theorem")
	}
	if a.CanonicalProtectedToBrokenMapDerived {
		t.Fatal("gate must not claim a canonical protected-to-broken map")
	}
}
