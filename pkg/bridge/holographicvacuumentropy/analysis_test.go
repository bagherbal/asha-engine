package holographicvacuumentropy

import "testing"

func TestGate373HolographicAudit(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.NativeChargedModuli != 13 || a.Inheritance.ExternalLedger != 15 {
		t.Fatalf("unexpected inherited census: %+v", a.Inheritance)
	}
	if !a.Gravity.FixesAbsoluteScale || a.Gravity.FixesFlavorTexture || a.Gravity.NativeFlavorEquations != 0 {
		t.Fatalf("gravity boundary should fix scale but not texture: %+v", a.Gravity)
	}
	if a.VacuumEnergy.UniqueNativeFunctional || !a.VacuumEnergy.CountertermRequired || a.VacuumEnergy.IndependentFlavorEquations != 0 || a.VacuumEnergy.CKMTextureEquations != 0 {
		t.Fatalf("vacuum energy should not be a unique texture equation: %+v", a.VacuumEnergy)
	}
	if len(a.Holography.Lanes) != 7 {
		t.Fatalf("unexpected number of lanes: %d", len(a.Holography.Lanes))
	}
	if a.Holography.TotalIndependentFlavorEquations != 0 || a.Holography.AnyTextureConstraint || a.Holography.AnyVacuumSelection {
		t.Fatalf("holography unexpectedly reduced moduli: %+v", a.Holography)
	}
	if a.Information.NumberOperatorSelectedByGravity || a.Information.ThermalTimeActivated || a.Information.IndependentFlavorEquations != 0 {
		t.Fatalf("information horizon unexpectedly selected N: %+v", a.Information)
	}
	if a.Census.Reduction != 0 || a.Census.RemainingChargedModuli != 13 {
		t.Fatalf("unexpected census reduction: %+v", a.Census)
	}
	if !a.Firewall.NoSaturationAssumed || !a.Firewall.NoContinuumBetaFunctionsFitted || !a.Firewall.NoCosmologicalConstantImported {
		t.Fatalf("firewall breach: %+v", a.Firewall)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := HolographicVacuumEntropyGravitationalModuliConstraintSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
