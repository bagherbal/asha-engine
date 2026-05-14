package cosmologicalobservablesdarksector

import (
	"math"
	"testing"
)

func TestGate385HiggsInheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Higgs.TreeProxySealed || !a.Higgs.UsesEdgeMeasure || a.Higgs.PoleMassDerived {
		t.Fatalf("bad Higgs inheritance: %+v", a.Higgs)
	}
	if a.Higgs.LambdaEW < 0.12 || a.Higgs.LambdaEW > 0.13 || a.Higgs.HiggsMassProxyGeV < 124 || a.Higgs.HiggsMassProxyGeV > 126 {
		t.Fatalf("unexpected Higgs values: %+v", a.Higgs)
	}
}

func TestDarkMatterNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.DarkMatter.NativeMassScaleAvailable || a.DarkMatter.StableCandidateDerived || a.DarkMatter.BoltzmannKernelClosed || a.DarkMatter.OmegaH2Derived || !math.IsNaN(a.DarkMatter.OmegaH2) {
		t.Fatalf("dark matter should remain non-derived: %+v", a.DarkMatter)
	}
	if len(a.DarkMatter.Inputs) < 6 {
		t.Fatalf("expected missing Boltzmann inputs: %+v", a.DarkMatter.Inputs)
	}
}

func TestVacuumFateNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Vacuum.FullRGTrajectoryDerived || a.Vacuum.ThresholdMatchingDerived || a.Vacuum.LambdaMinimumDerived || a.Vacuum.LifetimeDerived {
		t.Fatalf("vacuum fate should remain non-derived: %+v", a.Vacuum)
	}
	if !math.IsNaN(a.Vacuum.EuclideanBounceAction) || !math.IsNaN(a.Vacuum.LifetimeYears) {
		t.Fatalf("bounce/lifetime must not be numeric: %+v", a.Vacuum)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeCosmologyConstants()
	if m["lambda_EW_edge"] < 0.12 || m["lambda_EW_edge"] > 0.13 {
		t.Fatalf("bad lambda: %+v", m)
	}
	if m["omega_dm_h2_derived"] != 0 || m["vacuum_lifetime_derived"] != 0 || m["hard_predictions_derived"] != 0 {
		t.Fatalf("must not claim hard observables: %+v", m)
	}
	if m["conditional_targets_opened"] != 2 {
		t.Fatalf("expected two opened targets: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := CosmologicalObservablesDarkSectorPredictionAfterHiggsSealSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
