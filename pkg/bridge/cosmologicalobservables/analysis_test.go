package cosmologicalobservables

import (
	"math"
	"testing"
)

func TestBuildDefaultCosmologicalObservables(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Inheritance.Executed || a.Inheritance.HighestInheritedGate != 374 || a.Inheritance.ChargedModuli != 13 || a.Inheritance.ExternalLedger != 15 {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Scales.Executed || a.Scales.PeVThresholdGeV <= 1e6 || a.Scales.SealedIntermediateGeV <= 1e11 || a.Scales.RequiredQuarticJump >= 0 {
		t.Fatalf("bad scale ledger: %+v", a.Scales)
	}
	if a.Scales.PeVThresholdDerivedAsMass || a.Scales.IntermediateScaleRelicTheorem || a.Scales.HeavyDarkStabilityDerived {
		t.Fatalf("scale ledger overpromoted relic theorem: %+v", a.Scales)
	}
	if a.DarkMatter.OmegaH2Derived || !math.IsNaN(a.DarkMatter.OmegaH2Prediction) || a.DarkMatter.BoltzmannKernelClosed || a.DarkMatter.StabilityDerived {
		t.Fatalf("dark matter should remain non-derived: %+v", a.DarkMatter)
	}
	if a.Lifetime.LifetimeYearsDerived || !math.IsNaN(a.Lifetime.LifetimeYears) || a.Lifetime.EuclideanBounceActionDerived || !math.IsNaN(a.Lifetime.EuclideanBounceAction) {
		t.Fatalf("lifetime should remain non-derived: %+v", a.Lifetime)
	}
	if a.DarkEnergy.CosmologicalConstantDerived || a.DarkEnergy.DarkEnergyDensityDerived || a.DarkEnergy.VacuumCountertermDerived || len(a.DarkEnergy.NativeSuppressionPowers) != 3 {
		t.Fatalf("dark energy should remain non-derived: %+v", a.DarkEnergy)
	}
	if a.Census.HardPredictionsDerived != 0 || a.Census.RemainingChargedModuli != 13 {
		t.Fatalf("unexpected observable census: %+v", a.Census)
	}
	if !a.Firewall.NoObservedOmegaDMFitted || !a.Firewall.NoObservedDarkEnergyFitted || !a.Firewall.NoRGTrajectoryFitted || !a.Firewall.NoClaimBeyondInputs {
		t.Fatalf("firewall breach: %+v", a.Firewall)
	}
}

func TestNativeCosmologicalScales(t *testing.T) {
	m := NativeCosmologicalScales()
	if m["lambdaH_over_gstar2"] <= 0.25 || m["lambdaH_over_gstar2"] >= 0.27 {
		t.Fatalf("unexpected lambda ratio: %.12g", m["lambdaH_over_gstar2"])
	}
	if m["PeV_threshold_GeV"] <= 1e6 || m["sealed_intermediate_GeV"] <= 1e11 {
		t.Fatalf("unexpected heavy scales: %+v", m)
	}
	if !(m["hierarchy_fourth_power"] > 0 && m["hierarchy_fourth_power"] < 1e-60) {
		t.Fatalf("unexpected hierarchy fourth power: %.12e", m["hierarchy_fourth_power"])
	}
}

func TestTheoremPasses(t *testing.T) {
	res := CosmologicalObservablesDarkSectorPredictionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
