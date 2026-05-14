package spectralgraphtracenormalization

import (
	"math"
	"testing"
)

func TestNodeEdgeTraceDomains(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if len(c.Domains) != 2 {
		t.Fatalf("expected node and edge domains: %+v", c.Domains)
	}
	if math.Abs(c.Bridge.NodeCount-7) > 1e-12 || math.Abs(c.Bridge.EdgeCount-10) > 1e-12 {
		t.Fatalf("bad node/edge counts: %+v", c.Bridge)
	}
	if math.Abs(c.Bridge.NodeToEdgeRatio-10.0/7.0) > 1e-12 {
		t.Fatalf("bad bridge ratio: %+v", c.Bridge)
	}
}

func TestKineticTraceWitnessNotSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	k := a.Calculation.Kinetic
	if !k.UsesDFCommutator || !k.MandatesEdgeSupport {
		t.Fatalf("kinetic term should be edge-supported: %+v", k)
	}
	if k.MandatesEdgeDenominator {
		t.Fatalf("Gate 383 must not promote edge denominator to theorem: %+v", k)
	}
}

func TestHiggsLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	lanes := map[string]HiggsLane{}
	for _, l := range a.Calculation.Lanes {
		lanes[l.Name] = l
	}
	if lanes["contact-node denominator"].MassPfaffianGeV < 145 {
		t.Fatalf("contact-node lane should overpredict: %+v", lanes["contact-node denominator"])
	}
	if math.Abs(lanes["edge-trace denominator"].MassPfaffianGeV-HiggsTargetGeV) > 0.3 {
		t.Fatalf("edge trace lane should near-close: %+v", lanes["edge-trace denominator"])
	}
	if lanes["edge-trace denominator"].Sealed {
		t.Fatalf("edge trace lane must not be sealed yet: %+v", lanes["edge-trace denominator"])
	}
	if lanes["unit sharp-cutoff denominator"].MassPfaffianGeV < 390 {
		t.Fatalf("unit lane should badly overpredict: %+v", lanes["unit sharp-cutoff denominator"])
	}
}

func TestClosureStillOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if c.EdgeTraceDerived || c.TenOverSevenDerived || c.HiggsMassSealed || c.FullNumericalTOEClosed {
		t.Fatalf("closure must remain open: %+v", c)
	}
	if !c.Bridge.CombinatorialNative || c.Bridge.CCMNormalizationNative {
		t.Fatalf("bridge classification wrong: %+v", c.Bridge)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["contact_node_count"]-7) > 1e-12 || math.Abs(m["j_doubled_edge_count"]-10) > 1e-12 {
		t.Fatalf("bad constants: %+v", m)
	}
	if math.Abs(m["node_to_edge_ratio"]-10.0/7.0) > 1e-12 {
		t.Fatalf("bad bridge ratio: %+v", m)
	}
	if m["higgs_mass_sealed"] != 0 {
		t.Fatalf("must not seal Higgs mass: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SpectralGraphTraceNodeToEdgeKineticNormalizationSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
