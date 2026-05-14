package protectedmetric

import "testing"

func TestProtectedMetricAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.ProtectedDimension != 3 || a.BrokenImageDimension != 3 {
		t.Fatalf("expected 3D protected and broken carriers, got %d and %d", a.ProtectedDimension, a.BrokenImageDimension)
	}
	if !a.AbstractEuclideanMetricAvailable {
		t.Fatal("expected abstract Euclidean metric diagnostic")
	}
	if a.FiniteProtectedMetricDerived || a.FiniteProtectedConnectionDerived || a.O3FreedomReduced {
		t.Fatal("Gate 88 must not claim a derived metric/connection or reduced O(3) freedom")
	}
	if !a.PullbackCircularityDetected {
		t.Fatal("expected circularity detection for the broken metric pullback route")
	}
}
