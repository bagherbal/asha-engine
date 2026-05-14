package innerfluctuationedgemeasure

import (
	"math"
	"testing"
)

func TestInnerFluctuationIsOneForm(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.InnerFluctuation.UsesCommutatorDF || !c.InnerFluctuation.IsFiniteOneForm {
		t.Fatalf("expected finite one-form: %+v", c.InnerFluctuation)
	}
}

func TestSupportProjectionSelectsEdges(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if len(c.Support.Edges) != int(JDoubledEdgeCount) {
		t.Fatalf("expected ten doubled edge slots: %+v", c.Support.Edges)
	}
	if !c.Support.AEqualsPEAPE || !c.Support.EdgeMeasureMandated || c.Support.NodeMeasureAdmissible {
		t.Fatalf("bad support selection: %+v", c.Support)
	}
}

func TestCCMEdgeMeasureSelectionTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Calculation.Theorem.Proven || !a.Calculation.EdgeMeasureSelected {
		t.Fatalf("expected edge measure theorem: %+v", a.Calculation.Theorem)
	}
}

func TestHiggsProxySealedButPoleOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	c := a.Calculation
	if !c.HiggsTreeProxySealed {
		t.Fatalf("expected tree proxy sealed: %+v", c.Higgs)
	}
	if math.Abs(c.Higgs.MassPfaffianGeV-HiggsTargetGeV) > 0.3 {
		t.Fatalf("bad Higgs proxy mass: %+v", c.Higgs)
	}
	if c.PhysicalPoleMassDerived || c.FullNumericalTOEClosed {
		t.Fatalf("must not claim pole/full closure: %+v", c)
	}
}

func TestNativeConstants(t *testing.T) {
	m := NativeConstants()
	if math.Abs(m["edge_count"]-10) > 1e-12 || math.Abs(m["node_count"]-7) > 1e-12 {
		t.Fatalf("bad counts: %+v", m)
	}
	if m["edge_measure_selected"] != 1 || m["higgs_tree_proxy_sealed"] != 1 {
		t.Fatalf("expected selected/sealed proxy flags: %+v", m)
	}
	if m["physical_pole_mass_derived"] != 0 || m["full_numerical_toe_closed"] != 0 {
		t.Fatalf("must preserve physical/full closure flags: %+v", m)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := InnerFluctuationOneFormSupportCCMEdgeMeasureSelectionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
