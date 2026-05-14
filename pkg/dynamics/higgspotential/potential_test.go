package higgspotential

import "testing"

func TestPotentialCandidateKinematics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() failed: %v", err)
	}
	if a.ActiveContactDimension != 4 {
		t.Fatalf("active contact dimension = %d, want 4", a.ActiveContactDimension)
	}
	if a.ProtectedContactDimension != 3 {
		t.Fatalf("protected contact dimension = %d, want 3", a.ProtectedContactDimension)
	}
	if !a.PairDegenerateSpectrum {
		t.Fatalf("expected pair-degenerate contact spectrum, residual %.3e", a.PairDegeneracyResidual)
	}
	if a.OrderParameterNormSquared <= 0 {
		t.Fatalf("order parameter trace should be positive")
	}
	if a.QuarticTrace <= 0 || a.NormalizedQuarticShape <= 0 {
		t.Fatalf("quartic invariants should be positive")
	}
	if !a.MexicanHatKinematics {
		t.Fatalf("expected finite Mexican-hat kinematic ingredients")
	}
}
