package generationbreak

import "testing"

func TestFiniteGenerationBreakingSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.GenerationCarrierDimension != 3 || a.ProtectedContactDimension != 3 {
		t.Fatalf("expected 3D generation carrier/protected contact sector, got %d/%d", a.GenerationCarrierDimension, a.ProtectedContactDimension)
	}
	if !a.DiagonalSpurionFound {
		t.Fatalf("expected Higgs/contact anisotropy to expose a diagonal spurion")
	}
	if a.BestCandidate.ProducesMixing {
		t.Fatalf("diagonal spurion should not claim mixing")
	}
	if a.CanonicalOperatorFound {
		t.Fatalf("canonical operator should remain open")
	}
	if len(a.PartialOverlapSpectrum) <= a.GenerationCarrierDimension {
		t.Fatalf("expected overcomplete contact leakage spectrum, got %d modes", len(a.PartialOverlapSpectrum))
	}
}
