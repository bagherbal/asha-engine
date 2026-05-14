package ashafinalarchitectureledger

import "testing"

func TestFinalArchitectureBuilds(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Final.ProjectSealed {
		t.Fatalf("final ledger should be sealed as architecture: %+v", a.Final)
	}
	if a.Final.EnvironmentalModuliCount != 13 {
		t.Fatalf("expected 13 environmental moduli, got %+v", a.Final)
	}
	if a.Final.CosmologicalHardCount != 0 {
		t.Fatalf("must not claim hard cosmological predictions: %+v", a.Final)
	}
}

func TestAbsoluteLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Absolute.GaugeGroup != "SU(3) × SU(2) × U(1)" || a.Absolute.Generations != 3 || a.Absolute.HiggsDoublets != 1 || a.Absolute.GaugeBosons != 12 {
		t.Fatalf("bad absolute ledger: %+v", a.Absolute)
	}
	if a.Absolute.ParameterFreeDerivations < 10 {
		t.Fatalf("expected at least ten catalogued derivations: %+v", a.Absolute.Predictions)
	}
}

func TestHiggsTreeProxyButNotPoleMass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Higgs.TreeProxySealed || !a.Higgs.EdgeMeasureSelected || a.Higgs.PoleMassDerived {
		t.Fatalf("bad Higgs boundary: %+v", a.Higgs)
	}
	if a.Higgs.LambdaEW < 0.12 || a.Higgs.LambdaEW > 0.13 || a.Higgs.MassPfaffianGeV < 124 || a.Higgs.MassPfaffianGeV > 126 {
		t.Fatalf("unexpected Higgs values: %+v", a.Higgs)
	}
}

func TestCosmologyBoundary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Cosmology.HardPredictionsDerived != 0 || a.Cosmology.ConditionalTargetsOpened != 2 {
		t.Fatalf("bad cosmology census: %+v", a.Cosmology)
	}
	if len(a.Cosmology.Structures) != 3 {
		t.Fatalf("expected three cosmological structures: %+v", a.Cosmology.Structures)
	}
}

func TestModuliQuarantine(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Moduli.MinimalChargedFiniteDiracDim != 13 || a.Moduli.ExternalMinimalLedger != 15 || a.Moduli.HiddenFlavorConstraints != 0 || a.Moduli.NativeReductionBelow13 {
		t.Fatalf("bad moduli quarantine: %+v", a.Moduli)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := AshaFrameworkFinalArchitectureLedgerEpistemologicalSealTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
