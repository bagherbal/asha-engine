package couplingnorm

import (
	"math"
	"testing"
)

func TestBuildDefaultCouplingNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.ContactIndex <= 0 || a.TopologicalActionSeal <= 0 {
		t.Fatalf("invalid action anchors: I=%v S=%v", a.ContactIndex, a.TopologicalActionSeal)
	}
	if math.Abs(a.TopologicalActionSeal-8*math.Pi*math.Pi) > 1e-8 {
		t.Fatalf("unexpected action seal: got %.12f", a.TopologicalActionSeal)
	}
	if math.Abs(a.UnitTraceGaugeCouplingSq-1) > 1e-12 {
		t.Fatalf("unit-trace diagnostic should give g^2=1, got %.12f", a.UnitTraceGaugeCouplingSq)
	}
	if math.Abs(a.UnitTraceInverseAlpha-4*math.Pi) > 1e-12 {
		t.Fatalf("unit-trace inverse alpha should be 4π, got %.12f", a.UnitTraceInverseAlpha)
	}
	if a.GaugeCouplingDerived || a.FineStructureDerived || a.DimensionfulScaleDerived {
		t.Fatalf("bridge should not claim derived physical coupling or scale")
	}
	if a.HiddenObservedCouplingUsed {
		t.Fatalf("observed coupling was unexpectedly marked as used")
	}
}
