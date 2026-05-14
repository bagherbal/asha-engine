package u1completion

import "testing"

func TestAbelianCoefficientSelectionSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.TargetKappa != 6 {
		t.Fatalf("expected whitening target kappa=6, got %.12f", a.TargetKappa)
	}
	if a.CandidateHitCount < 2 {
		t.Fatalf("expected multiple count resonances to hit 6, got %d", a.CandidateHitCount)
	}
	if a.ActionSelectsKappa || a.FiniteSecondVariation || a.KappaPhysical || a.PhysicalCouplingsOrMasses {
		t.Fatalf("Gate 99 must not claim action-selected kappa or physical couplings")
	}
}
