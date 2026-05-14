package texture

import (
	"math"
	"testing"
)

func TestTrialityInvariantEigenvaluesHaveDoublet(t *testing.T) {
	eig := TrialityInvariantEigenvalues(1, 0.25)
	if len(eig) != 3 {
		t.Fatalf("expected 3 eigenvalues")
	}
	if math.Abs(eig[1]-eig[2]) > 1e-12 {
		t.Fatalf("expected exact doublet degeneracy, got %v", eig)
	}
}

func TestTextureSearchDoesNotSelectCouplings(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GenerationBreakingOperatorFound {
		t.Fatalf("current gate must not pretend a generation-breaking operator was found")
	}
	if a.CouplingsDerived || a.CKMDerived || a.PMNSDerived {
		t.Fatalf("Gate 28 must not derive phenomenological matrices")
	}
	if a.TrialityInvariantTextureDim != 2 || a.GeneralTextureDim != 9 || a.SymmetricTextureDim != 6 {
		t.Fatalf("unexpected texture dimensions: %+v", a)
	}
	if len(a.KindSummaries) != 4 {
		t.Fatalf("expected four fermion-kind texture summaries, got %d", len(a.KindSummaries))
	}
}
