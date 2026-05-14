package wittso8coordinates

import (
	"math"
	"testing"
)

func TestGate253WittCoordinatesOpenButPhysicalEWStillBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Summary.WittPairingRetrieved || !a.Summary.NumberSO8Coordinates || !a.Summary.KnownFockLedgersCoordinateReady {
		t.Fatalf("expected Witt/number coordinate dictionary, got %+v", a.Summary)
	}
	if a.Summary.T3LYPhiSO8Coordinates || a.Summary.Q8vCConstructed || a.Summary.Neutral3PlaneDerived {
		t.Fatalf("unexpected physical electroweak/Q/kernel derivation: %+v", a.Summary)
	}
	if len(a.NumberOperators.Coordinates) != 4 {
		t.Fatalf("expected four Cartan coordinates, got %d", len(a.NumberOperators.Coordinates))
	}
	bl := a.FockLedgers.Ledgers[0]
	want := []float64{-0.5, 1.0 / 6.0, 1.0 / 6.0, 1.0 / 6.0}
	for i := range want {
		if math.Abs(bl.BivectorCoefficients[i]-want[i]) > 1e-12 {
			t.Fatalf("B-L bivector coeff[%d]=%g want %g", i, bl.BivectorCoefficients[i], want[i])
		}
	}
}

func TestGate253TrialityCandidatesNotSelectedByOutcome(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.Triality.CandidateCount != 2 {
		t.Fatalf("expected two audited Cartan triality candidates, got %d", a.Triality.CandidateCount)
	}
	for _, c := range a.Triality.Candidates {
		if !c.Orthogonal || !c.DetAbsOne || c.Selected {
			t.Fatalf("bad triality candidate audit: %+v", c)
		}
	}
	if a.Kernel.ThreePlaneDerived || a.Firewall.ForcedKernelDim3 || a.Firewall.SelectedTrialityByOutcome {
		t.Fatalf("kernel/triality firewall violated: kernel=%+v firewall=%+v", a.Kernel, a.Firewall)
	}
}
