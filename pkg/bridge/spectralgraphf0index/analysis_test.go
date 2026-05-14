package spectralgraphf0index

import (
	"math"
	"strings"
	"testing"
)

func TestEdgeProjectionTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	p := a.Calculation.EdgeProjection
	if p.FundamentalEdgeCount != 5 || p.JDoubledEdgeSlotCount != 10 || math.Abs(p.ProjectionTraceOnEdges-10) > 1e-12 {
		t.Fatalf("bad edge projection: %+v", p)
	}
	if p.IsTraceOverHFWellTyped {
		t.Fatalf("edge projection must not be promoted to Tr_HF without theorem")
	}
}

func TestCCMF0DefinitionFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	m := a.Calculation.Moment
	if m.SameMathematicalObject {
		t.Fatalf("CCM f0 must not be identified with edge projection trace")
	}
	if m.SharpCutoffValue != 1 {
		t.Fatalf("unexpected sharp cutoff f0 value: %.18g", m.SharpCutoffValue)
	}
	if !strings.Contains(m.WouldRequireTheorem, "SpectralGraphMomentTheorem") {
		t.Fatalf("missing required theorem statement: %+v", m)
	}
}

func TestIndexTheoremSieve(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	i := a.Calculation.Index
	if i.KernelIndexDerived || i.AllEdgeCountIndex || i.CanIdentifyF0WithIndex {
		t.Fatalf("index theorem was overpromoted: %+v", i)
	}
}

func TestHiggsNearClosureInheritedButNotSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	h := a.Calculation.Higgs
	if math.Abs(h.F0Candidate-10) > 1e-12 || math.Abs(h.MassPfaffianVEVGeV-HiggsBoundaryGeV) > 0.3 {
		t.Fatalf("bad f0=10 inherited near-closure: %+v", h)
	}
	if h.GeometricallySealed || a.Calculation.HiggsMassSealed || a.Calculation.F0MomentIndexDerived {
		t.Fatalf("near-closure must not be sealed without f0 theorem: %+v", a.Calculation)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["edge_projection_trace"]-10) > 1e-12 {
		t.Fatalf("bad constants: %+v", m)
	}
	if math.Abs(m["ccm_sharp_cutoff_f0_value"]-1) > 1e-12 {
		t.Fatalf("bad ccm f0 value: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SpectralGraphProjectionF0IndexTheoremSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
