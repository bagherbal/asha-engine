package asha

import (
	"bytes"
	"strings"
	"testing"
)

func TestRuntimeReportAllPasses(t *testing.T) {
	e := New()
	r, err := e.Report(ScenarioAll)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range r.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
	if e.Metadata.LatestGate != 551 {
		t.Fatalf("unexpected latest gate %d", e.Metadata.LatestGate)
	}
}

func TestHiggsProxy(t *testing.T) {
	e := New()
	if !nearly(e.Bridge.HiggsTreeGeV, 124.925370288, 2e-3) {
		t.Fatalf("unexpected Higgs proxy %.12f", e.Bridge.HiggsTreeGeV)
	}
	if !nearly(e.Bridge.LambdaH, 0.12774563655, 1e-10) {
		t.Fatalf("unexpected lambda %.12f", e.Bridge.LambdaH)
	}
}

func TestFamilyAxiomBoundary(t *testing.T) {
	e := New(WithBeta(1))
	if e.Family.NativeChargedFlavorDim != 13 || e.Family.KXYChargedCoeffDim != 9 || !e.Family.KGenGeometricallyForced || !e.Family.Generation2BareZero {
		t.Fatalf("bad family dims: %+v", e.Family)
	}
	if e.Family.CommKXNorm <= 0 || e.Family.CPWitness == 0 {
		t.Fatalf("expected conditional mixing/CP capacity: %+v", e.Family)
	}
}

func TestFormatters(t *testing.T) {
	e := New()
	r, err := e.Report(ScenarioCI)
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	if err := WriteReport(&b, r, FormatJSON); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(b.String(), "dim M_charged^native") {
		t.Fatal("json report missing family firewall quantity")
	}
	md := Markdown(r)
	for _, want := range []string{"ASHA Runtime Board", "m_H^tree", "dim M_charged^native"} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q", want)
		}
	}
}

func TestEnvironmentalScenarioComputesNumbers(t *testing.T) {
	e := New()
	dm := ComputeDarkMatterConditional(e.Bridge.HeavyBGapMajoranaGeV)
	if !nearly(dm.OverclosureFactor, 1.309310776326e13, 1e8) {
		t.Fatalf("unexpected dark overclosure %.12e", dm.OverclosureFactor)
	}
	if !(dm.RequiredFractionOfThermal > 7e-14 && dm.RequiredFractionOfThermal < 8e-14) {
		t.Fatalf("unexpected required fraction %.12e", dm.RequiredFractionOfThermal)
	}
	cc := ComputeCosmologyConditional(e.Bridge.PlanckMassGeV, e.Bridge.VPfGeV)
	if !nearly(cc.DigitsOfCancellation, 120.686941491987, 1e-9) {
		t.Fatalf("unexpected cancellation digits %.12f", cc.DigitsOfCancellation)
	}
	if !nearly(cc.EWVacuumOverGate344Target, 1.679361894449e55, 1e49) {
		t.Fatalf("unexpected EW/Gate344 target ratio %.12e", cc.EWVacuumOverGate344Target)
	}
	r, err := e.Report(ScenarioEnvironment)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Sections) != 2 {
		t.Fatalf("expected cosmology + vacuum fate sections, got %d", len(r.Sections))
	}
}
