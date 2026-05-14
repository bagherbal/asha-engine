package ccmpfaffianf0closure

import (
	"math"
	"testing"
)

func TestEffectiveF0Targets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.Executed || !c.ConditionalNearClosure || c.NativeHiggsMassClosed || c.FullNumericalTOEClosure {
		t.Fatalf("bad closure flags: %+v", c)
	}
	if math.Abs(c.F0Targets.UsingStandardEWVEV-9.89710333989887) > 1e-9 {
		t.Fatalf("bad standard f0 target: %.18g", c.F0Targets.UsingStandardEWVEV)
	}
	if math.Abs(c.F0Targets.UsingPfaffianVEV-9.972101066696943) > 1e-9 {
		t.Fatalf("bad pfaffian f0 target: %.18g", c.F0Targets.UsingPfaffianVEV)
	}
}

func TestF0Candidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	std := a.Calculation.StandardVEVPredictions
	pf := a.Calculation.PfaffianVEVPredictions
	if len(std) != 3 || len(pf) != 3 {
		t.Fatalf("missing predictions")
	}
	if std[0].PredictedMass <= 140 || std[1].PredictedMass < 124 || std[1].PredictedMass > 125 || std[2].PredictedMass >= 110 {
		t.Fatalf("unexpected standard predictions: %+v", std)
	}
	if pf[0].PredictedMass <= 140 || math.Abs(pf[1].PredictedMass-HiggsMassBoundaryGeV) > 0.3 || pf[2].PredictedMass >= 110 {
		t.Fatalf("unexpected pfaffian predictions: %+v", pf)
	}
	if math.Abs(lambdaForF0(10)-0.12774563654956705) > 1e-15 {
		t.Fatalf("bad lambda f0=10: %.18g", lambdaForF0(10))
	}
}

func TestEdgeSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	e := a.Calculation.EdgeSieve
	if e.FundamentalEdgeCount != 5 || e.JDoubledEdgeCount != 10 || !e.MatchesF0Target {
		t.Fatalf("bad edge sieve: %+v", e)
	}
	if e.IsSpectralMomentProof {
		t.Fatalf("edge count must not be promoted to f0 moment theorem")
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["f0_eff_pfaffian_unreduced_planck"]-9.972101066696943) > 1e-9 {
		t.Fatalf("bad constants: %+v", m)
	}
	if math.Abs(m["j_doubled_edge_count"]-10) > 1e-12 {
		t.Fatalf("bad edge count: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SelfConsistentCCMPfaffianCoefficientClosureF0SieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
