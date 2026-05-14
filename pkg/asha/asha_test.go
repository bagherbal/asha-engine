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
	if e.Metadata.LatestGate != 425 {
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
	if e.Family.NativeChargedFlavorDim != 13 || e.Family.KXYChargedCoeffDim != 9 {
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
