package lietrialitypullback

import "testing"

func TestGate252InfinitesimalTrialityCapacityButNoPullback(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Summary.InfinitesimalTrialityCapacity || !a.Summary.SpinorEWBridgeKnown {
		t.Fatalf("expected infinitesimal triality capacity and bridge EW inputs, got %+v", a.Summary)
	}
	if a.Summary.SpinorSO8Coordinates || a.Summary.ExplicitLieTrialityMap || a.Summary.VectorEWMatriciesDerived {
		t.Fatalf("unexpected derived domain/map/vector matrices: %+v", a.Summary)
	}
	if a.Summary.Q8vCConstructed || a.Summary.Neutral3PlaneDerived {
		t.Fatalf("unexpected Q8vC/neutral kernel derivation: %+v", a.Summary)
	}
	if a.Summary.JCompatibleTransport || a.Summary.VTauConstructed || a.Summary.TrialityUnblocked || a.Summary.YukawaTextureDerived {
		t.Fatalf("unexpected downstream triality/Yukawa derivation: %+v", a.Summary)
	}
}
