package finitetraceedgemultiplicity

import (
	"math"
	"testing"
)

func TestMomentAndEdgeMultiplicity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if math.Abs(c.Moment.LockedValue-1) > 1e-12 {
		t.Fatalf("f0 moment must be locked to one: %+v", c.Moment)
	}
	if c.EdgeMultiplicity.JDoubledEdgeCount != 10 || c.EdgeMultiplicity.CanReplaceF0 {
		t.Fatalf("bad edge multiplicity firewall: %+v", c.EdgeMultiplicity)
	}
}

func TestCoefficientLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	lanes := map[string]CoefficientLane{}
	for _, l := range a.Calculation.Lanes {
		lanes[l.Name] = l
	}
	if lanes["unit f₀, no extra multiplicity"].MassPfaffianGeV < 390 {
		t.Fatalf("unit f0 should overpredict strongly: %+v", lanes["unit f₀, no extra multiplicity"])
	}
	if lanes["wrong numerator multiplication"].MassPfaffianGeV < 1200 {
		t.Fatalf("numerator multiplication should be catastrophic: %+v", lanes["wrong numerator multiplication"])
	}
	if math.Abs(lanes["denominator edge normalization witness"].MassPfaffianGeV-HiggsBoundaryGeV) > 0.3 {
		t.Fatalf("denominator witness should near-close: %+v", lanes["denominator edge normalization witness"])
	}
	if lanes["denominator edge normalization witness"].Native || !lanes["denominator edge normalization witness"].CircularRisk {
		t.Fatalf("denominator witness must not be promoted to native theorem: %+v", lanes["denominator edge normalization witness"])
	}
}

func TestTraceRatioNoExtraTen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	d := a.Calculation.TraceDecomposition
	if !d.IsAlreadyTraceRatio || d.CanPullExtraTenFromRatio {
		t.Fatalf("trace ratio must not pull an extra ten: %+v", d)
	}
	if math.Abs(d.FiniteRatioValue-1197.0/4624.0) > 1e-15 {
		t.Fatalf("bad finite ratio: %.18g", d.FiniteRatioValue)
	}
}

func TestTenOverSevenGapOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	g := a.Calculation.Gap
	if math.Abs(g.RatioTenOverSeven-10.0/7.0) > 1e-12 || g.RecognizedNative {
		t.Fatalf("bad ten-over-seven gap audit: %+v", g)
	}
	if a.Calculation.HiggsMassSealed || a.Calculation.FullNumericalTOEClosure {
		t.Fatalf("must not seal closure: %+v", a.Calculation)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["unit_ccm_f0"]-1) > 1e-12 || math.Abs(m["j_doubled_edge_count"]-10) > 1e-12 {
		t.Fatalf("bad native constants: %+v", m)
	}
	if math.Abs(m["ten_over_seven_gap"]-10.0/7.0) > 1e-12 {
		t.Fatalf("bad gap: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := FiniteTraceEdgeMultiplicityEffectiveCoefficientSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
