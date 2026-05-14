package ko6twistedrealstructure

import "testing"

func TestGate293EvenGradingTwistDoesNotFlipKOSign(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, c := range a.Twists.Candidates {
		if c.Name == "J0·gamma_F" {
			found = true
			if c.J2Sign != 1 || c.JGammaSign != 1 || c.KOSixLike {
				t.Fatalf("grading twist should remain KO0-like, got %+v", c)
			}
		}
	}
	if !found {
		t.Fatalf("missing grading twist candidate")
	}
}

func TestGate293OddTwistsExposeKO6SignsButAreDegenerate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Twists.KO6Candidates != 2 {
		t.Fatalf("expected exactly two odd KO6 sign candidates, got %+v", a.Twists)
	}
	if a.Twists.CanonicalKO6Found {
		t.Fatalf("odd twist must not be canonical without a selector")
	}
	for _, c := range a.Twists.Candidates {
		if c.KOSixLike && (c.J2Sign != 1 || c.JGammaSign != -1) {
			t.Fatalf("bad KO6 signs: %+v", c)
		}
	}
}

func TestGate293JDCommutationDoesNotSelectCanonicalDF(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.DiracSieve) != 2 {
		t.Fatalf("expected two odd twist Dirac sieves, got %d", len(a.DiracSieve))
	}
	for _, s := range a.DiracSieve {
		if s.GenericOddBlockParams != 4 || s.JDLinearConstraints != 1 || s.JDRealityFreeParams != 3 {
			t.Fatalf("unexpected JD sieve for %s: %+v", s.CandidateName, s)
		}
		if s.CanonicalDFSelected {
			t.Fatalf("JD sieve should not select canonical D_F: %+v", s)
		}
	}
}

func TestGate293OppositeActionAndDynamicsRemainFirewalled(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Opposite.OppositeActionConstructed || a.Opposite.PhysicalJAvailable {
		t.Fatalf("opposite action unexpectedly constructed: %+v", a.Opposite)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotPromoteOddTwistToPhysical || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
