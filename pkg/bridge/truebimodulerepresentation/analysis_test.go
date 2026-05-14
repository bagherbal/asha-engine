package truebimodulerepresentation

import "testing"

func TestGate295LeftRightActionsAreNontrivial(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Left.ActsOnQuarks || !a.Left.ActsOnLeptons || a.Left.NonTrivialResidual <= 0.1 {
		t.Fatalf("weak left action not established: %+v", a.Left)
	}
	if !a.Right.ActsOnQuarks || a.Right.ActsOnLeptons || a.Right.NonTrivialResidual <= 0.1 {
		t.Fatalf("color right action not established: %+v", a.Right)
	}
}

func TestGate295BimoduleZeroOrderCommutes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Bimodule.ZeroOrderVerified || a.Bimodule.WeakColorCommutatorNorm > 1e-12 {
		t.Fatalf("expected weak-left/color-right commutation: %+v", a.Bimodule)
	}
	if a.Bimodule.NaiveLeftCrossTermNorm <= 0.1 {
		t.Fatalf("expected inherited naive left cross-term obstruction: %+v", a.Bimodule)
	}
	if !a.Bimodule.ResolvesGate294Paradox {
		t.Fatalf("true bimodule should resolve zero-order paradox: %+v", a.Bimodule)
	}
}

func TestGate295HyperchargeAndFirstOrderRemainFirewalled(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Hypercharge.DerivedByGate || a.Hypercharge.FractionalChargesGenerated {
		t.Fatalf("hypercharge was overpromoted: %+v", a.Hypercharge)
	}
	if a.OrderOne.FirstOrderVerified || a.OrderOne.CanonicalDiracAvailable || a.OrderOne.DiracConstraintsDerived {
		t.Fatalf("first-order/DF unexpectedly derived: %+v", a.OrderOne)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotInventHypercharge || !a.Firewalls.DoesNotInventDF || !a.Firewalls.DoesNotUnlockHiggs || !a.Firewalls.DoesNotUnlockBGap {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
