package rawfinitetracerecomputation

import (
	"math"
	"testing"
)

func TestRawMeasureRatios(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.Symbolics.UniformMeasureLift {
		t.Fatalf("expected uniform measure lift: %+v", c.Symbolics)
	}
	if math.Abs(c.Symbolics.MeasureScale-10.0/7.0) > 1e-12 {
		t.Fatalf("bad measure scale: %+v", c.Symbolics)
	}
	if len(c.Measures) != 2 {
		t.Fatalf("expected node and edge measures: %+v", c.Measures)
	}
	want := TraceRatioNode * 7.0 / 10.0
	if math.Abs(c.Measures[1].RatioEOverA2-want) > 1e-12 {
		t.Fatalf("bad edge ratio: got %.18f want %.18f", c.Measures[1].RatioEOverA2, want)
	}
}

func TestEdgeMeasureNearClosureWithoutPostHocMultiplier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	lanes := map[string]HiggsLane{}
	for _, l := range a.Calculation.Lanes {
		lanes[l.Name] = l
	}
	near := lanes["edge-measure raw ratio with inherited node normalization"]
	if near.DoubleCounts {
		t.Fatalf("near lane must not double-count: %+v", near)
	}
	if math.Abs(near.MassPfaffianGeV-HiggsTargetGeV) > 0.3 {
		t.Fatalf("near lane should close: %+v", near)
	}
	if near.Sealed {
		t.Fatalf("near lane is conditional, not sealed: %+v", near)
	}
}

func TestLiteralF0UnitLaneStillFails(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	lanes := map[string]HiggsLane{}
	for _, l := range a.Calculation.Lanes {
		lanes[l.Name] = l
	}
	unit := lanes["literal CCM f0=1 with edge-measure ratio only"]
	if unit.MassPfaffianGeV < 320 {
		t.Fatalf("literal f0=1 edge-ratio lane should still overpredict: %+v", unit)
	}
}

func TestDoubleCountingRejected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	lanes := map[string]HiggsLane{}
	for _, l := range a.Calculation.Lanes {
		lanes[l.Name] = l
	}
	bad := lanes["double-counted edge ratio plus edge denominator"]
	if !bad.DoubleCounts {
		t.Fatalf("expected double-count flag: %+v", bad)
	}
	if bad.MassPfaffianGeV > HiggsTargetGeV-15 {
		t.Fatalf("double-count lane should substantially underpredict: %+v", bad)
	}
}

func TestClosureStillConditional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.EdgeMeasureTraceComputed || !c.TenOverSevenDerivedInsideRatio {
		t.Fatalf("expected trace recomputation support: %+v", c)
	}
	if c.EdgeMeasureSelectedNatively || c.HiggsMassSealed || c.FullNumericalTOEClosed {
		t.Fatalf("must remain conditional: %+v", c)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["edge_trace_ratio_e_over_a2"]-TraceRatioNode*0.7) > 1e-12 {
		t.Fatalf("bad constants: %+v", m)
	}
	if m["higgs_mass_sealed"] != 0 {
		t.Fatalf("must not seal: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := RawFiniteTraceRecomputationEdgeMeasureSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
