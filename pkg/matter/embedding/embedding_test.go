package embedding

import "testing"

func TestCanonicalEmbeddingSearchFindsObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HasFourDimensionalActiveSector {
		t.Fatalf("expected four-dimensional active sector")
	}
	if !a.HasFockOnePlusThreeSplit {
		t.Fatalf("expected Fock 1+3 split")
	}
	if a.SpectrumDeterminesOnePlusThree {
		t.Fatalf("pair-degenerate spectrum should not determine a 1+3 embedding")
	}
	if a.CanonicalEmbeddingConstructed {
		t.Fatalf("canonical embedding should remain open without a charge-polarizing operator")
	}
	if a.DegeneracyFreedomDimension != 2 {
		t.Fatalf("expected O(2)xO(2) degeneracy freedom dimension 2, got %d", a.DegeneracyFreedomDimension)
	}
}
