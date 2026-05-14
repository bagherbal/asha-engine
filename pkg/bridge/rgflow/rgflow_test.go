package rgflow

import "testing"

func TestRGFlowAuditDoesNotDerivePhysicalCouplings(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.HyperchargeNormalizationKY <= 0 || a.BoundarySin2Candidate <= 0 {
		t.Fatalf("expected finite boundary normalization candidate: %+v", a)
	}
	if !a.BoundaryCouplingFree || !a.LogScaleIntervalFree {
		t.Fatalf("RG audit must expose free boundary coupling and free log interval")
	}
	if a.BetaCoefficientsDerived || a.ThresholdSpectrumDerived || a.GaugeKineticDerived || a.BoundaryScaleDerived {
		t.Fatalf("RG audit should not claim missing bridge data is derived")
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.UnitNoRunningDiagnosticPhysical {
		t.Fatalf("RG audit must not claim physical weak angle or alpha_em")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("RG audit inserted observed input")
	}
}
