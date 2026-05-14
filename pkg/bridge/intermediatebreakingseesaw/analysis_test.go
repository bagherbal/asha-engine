package intermediatebreakingseesaw

import "testing"

func TestGate231BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Seal.Active || a.Seal.FiniteDerived {
		t.Fatalf("expected phenomenological IntermediateBreakingSeal activation, got %+v", a.Seal)
	}
	if a.Compute.OrderOneInPlausibleWindow || !a.Compute.OrderOneAboveCosmologyBound {
		t.Fatalf("order-one seesaw should fail, got %+v", a.Compute)
	}
	if a.Compute.OrderOneMassEV < 80 || a.Compute.OrderOneMassEV > 100 {
		t.Fatalf("expected order-one mass near 91 eV, got %.12g", a.Compute.OrderOneMassEV)
	}
	if a.Compute.YukawaForAtmosphericScale < 0.02 || a.Compute.YukawaForAtmosphericScale > 0.03 {
		t.Fatalf("expected small Yukawa near 0.023, got %.12g", a.Compute.YukawaForAtmosphericScale)
	}
	if a.Matrix.LightNeutrinoMatrixDerived {
		t.Fatalf("neutrino matrix must not be derived")
	}
	if a.Firewall.FiniteCorePolluted || a.Firewall.UsesObservedNeutrinoMassAsInput {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}
