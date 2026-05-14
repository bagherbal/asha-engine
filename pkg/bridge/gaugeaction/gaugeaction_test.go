package gaugeaction

import (
	"math"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TraceGramAsRepresentationMetric {
		t.Fatalf("expected trace-Gram representation metric")
	}
	if a.GaugeKineticActionSelected || a.BoundaryCouplingFixed || a.PhysicalU1CouplingDerived || a.FineStructureDerived {
		t.Fatalf("Gate 82 must not claim selected kinetic action or physical coupling")
	}
	if SelectedActionCount(a.CandidateActions) != 0 {
		t.Fatalf("no candidate action should be selected by Gate 82")
	}
	if math.Abs(a.ChargeTableKY-5.0/3.0) > 1e-12 {
		t.Fatalf("expected inherited kY=5/3, got %.12f", a.ChargeTableKY)
	}
	if a.TwoCarrierHyperchargeNormDiagnostic <= 0 || a.TwoCarrierInverseNormDiagnostic <= 0 {
		t.Fatalf("expected positive diagnostic norm family")
	}
}
