package abelianmixing

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChargeLevelHyperchargeBridge {
		t.Fatalf("expected charge-level hypercharge bridge")
	}
	if a.Central.Value != 0 {
		t.Fatalf("expected central coefficient 0, got %v", a.Central.Value)
	}
	if a.BMinusL.Value != 0.5 {
		t.Fatalf("expected B-L coefficient 1/2, got %v", a.BMinusL.Value)
	}
	if a.GaugeKineticNormalizationDerived {
		t.Fatalf("kinetic normalization must remain open")
	}
}
