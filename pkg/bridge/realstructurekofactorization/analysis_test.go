package realstructurekofactorization

import "testing"

func TestGate292JFactorizesAcrossSpacetimeFiberSplit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Factor.FullComplementMatchesTensor || a.Factor.Residual > 1e-12 {
		t.Fatalf("expected exact factorization, got %+v", a.Factor)
	}
	if a.Factor.FullDimension != 16 || a.Factor.SpacetimeDimension != 4 || a.Factor.FiberDimension != 4 {
		t.Fatalf("bad dimensions: %+v", a.Factor)
	}
}

func TestGate292FiberComplementIsNotKO6(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.KO.J2Sign != 1 || a.KO.JGammaSign != 1 {
		t.Fatalf("expected fiber complement to commute with parity and square to +1, got %+v", a.KO)
	}
	if a.KO.KOSixLike {
		t.Fatalf("fiber occupation complement must not be promoted to KO6: %+v", a.KO)
	}
}

func TestGate292DRealityDoesNotSelectCanonicalDF(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.DReality.GenericOddBlockParams != 4 || a.DReality.JRealityFreeParams != 2 {
		t.Fatalf("unexpected J-reality sieve dimensions: %+v", a.DReality)
	}
	if a.DReality.CanonicalDFSelected {
		t.Fatalf("J-reality sieve should not select canonical D_F")
	}
}

func TestGate292FirewallsRemainActive(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Opposite.OppositeActionConstructed || a.Opposite.HeatKernelUnblocked || a.Opposite.BGapInstantonUnblocked {
		t.Fatalf("unexpected dynamics unlock: %+v", a.Opposite)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotClaimKO6 || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
