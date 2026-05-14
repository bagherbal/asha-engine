package nativemodulispacecensus

import "testing"

func TestGate372NativeDiracCensus(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Parameterization.MinimalChargedRawDim != 54 || a.Parameterization.DiracNeutrinoRawDim != 72 || a.Parameterization.MajoranaExtendedRawDim != 84 {
		t.Fatalf("unexpected raw dims: %+v", a.Parameterization)
	}
	if a.Axioms.AdditionalGenerationConstraints != 0 {
		t.Fatalf("unexpected hidden generation constraints: %d", a.Axioms.AdditionalGenerationConstraints)
	}
	quark := FindScenario(a.Quotient.Scenarios, "quark Yukawa sector")
	if quark.PhysicalDim != 10 {
		t.Fatalf("quark quotient dimension = %d, want 10", quark.PhysicalDim)
	}
	charged := FindScenario(a.Quotient.Scenarios, "charged-lepton-only sector")
	if charged.PhysicalDim != 3 {
		t.Fatalf("charged lepton dimension = %d, want 3", charged.PhysicalDim)
	}
	minimal := FindScenario(a.Quotient.Scenarios, "minimal charged finite-Dirac flavor sector")
	if minimal.PhysicalDim != 13 {
		t.Fatalf("minimal charged finite-Dirac dimension = %d, want 13", minimal.PhysicalDim)
	}
	diracExtended := FindScenario(a.Quotient.Scenarios, "quark plus Dirac-neutrino finite-Dirac sector")
	if diracExtended.PhysicalDim != 20 {
		t.Fatalf("Dirac-neutrino finite-Dirac dimension = %d, want 20", diracExtended.PhysicalDim)
	}
	majoranaExtended := FindScenario(a.Quotient.Scenarios, "quark plus Majorana finite-Dirac sector")
	if majoranaExtended.PhysicalDim != 31 {
		t.Fatalf("Majorana finite-Dirac dimension = %d, want 31", majoranaExtended.PhysicalDim)
	}
	if a.Native.NativeReductionBelow15 || a.Native.HiddenCrossSectorConstraints != 0 {
		t.Fatalf("unexpected native reduction or hidden constraints: %+v", a.Native)
	}
	if !a.Firewall.NoGaugeFlavorConflation || !a.Firewall.NoMajoranaMinimalConflation {
		t.Fatalf("category firewalls not preserved: %+v", a.Firewall)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := NativeModuliSpaceDimensionExactDiracParameterCensusSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
