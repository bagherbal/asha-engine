package bfsource

import "testing"

func TestBFActionSourceTextureSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.FullCurvatureRank == 0 || a.FullCurvatureNorm <= 1e-8 {
		t.Fatalf("expected nonzero finite curvature input")
	}
	if a.ProtectedBFResponseRank != 0 || a.ProtectedBFMaxNorm > 1e-8 {
		t.Fatalf("expected protected BF response to vanish, rank=%d norm=%g", a.ProtectedBFResponseRank, a.ProtectedBFMaxNorm)
	}
	if a.MixedBFResponseRank != 0 || a.MixedBFMaxNorm > 1e-8 {
		t.Fatalf("expected mixed BF response to vanish, rank=%d norm=%g", a.MixedBFResponseRank, a.MixedBFMaxNorm)
	}
	if a.ActiveBFResponseRank == 0 || a.ActiveBFMaxNorm <= 1e-8 {
		t.Fatalf("expected active-only BF response to be nonzero")
	}
	if a.ProtectedQuadratic.Rank != 0 || a.MixedQuadratic.Rank != 0 {
		t.Fatalf("generation quadratics should vanish: protected=%d mixed=%d", a.ProtectedQuadratic.Rank, a.MixedQuadratic.Rank)
	}
	if a.ActiveQuadratic.Rank == 0 {
		t.Fatalf("expected active scalar quadratic to be nonzero")
	}
	if a.CanonicalTextureFound {
		t.Fatalf("current finite data must not claim a canonical generation texture")
	}
}
