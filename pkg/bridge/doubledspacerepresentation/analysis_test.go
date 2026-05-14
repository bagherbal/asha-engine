package doubledspacerepresentation

import "testing"

func TestGate294JSwapKOSigns(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.JSwap.KOSixLike || a.JSwap.J2Sign != 1 || a.JSwap.JGammaSign != -1 {
		t.Fatalf("bad J_swap signs: %+v", a.JSwap)
	}
	if a.JSwap.ResidualJ2 > 1e-12 || a.JSwap.ResidualGamma > 1e-12 {
		t.Fatalf("unexpected J_swap residuals: %+v", a.JSwap)
	}
}

func TestGate294NaiveWeakColorActionFailsDirectSumRepresentation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.NaiveDiagnostic.MultiplicativityResidual <= 0.1 {
		t.Fatalf("expected nonzero cross-term residual: %+v", a.NaiveDiagnostic)
	}
	if a.Representations[0].Associative {
		t.Fatalf("naive tensor action should not be a direct-sum representation: %+v", a.Representations[0])
	}
}

func TestGate294BlockSeparatedActionIsNotPhysicalSMBimodule(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	block := a.Representations[1]
	if !block.Associative || !block.Unital {
		t.Fatalf("block separated action should be associative/unital: %+v", block)
	}
	if block.PhysicalSMBimodule {
		t.Fatalf("block separated action must not be promoted to physical SM bimodule: %+v", block)
	}
}

func TestGate294OppositeAndOrderOneRemainBlocked(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Opposite.ConstructedForPhysicalHF || a.Opposite.ZeroOrderVerified {
		t.Fatalf("opposite action unexpectedly constructed: %+v", a.Opposite)
	}
	if a.OrderOne.OrderOneVerified || a.OrderOne.DiracConstraintsDerived {
		t.Fatalf("order-one unexpectedly verified: %+v", a.OrderOne)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotInventHF || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
