package orientationtruechirality

import "testing"

func TestGate239OrientationTrueChiralityAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Orientation.ActsOnSC || !a.Orientation.EquivalentToGamma || a.Orientation.DistinctFromGamma {
		t.Fatalf("volume orientation should be available but equivalent to gamma: %+v", a.Orientation)
	}
	if a.TauEta.CanonicalPullbackDerived || a.TauEta.CanActAsChiralityOperator {
		t.Fatalf("tau_eta must not be promoted to S_C chirality operator: %+v", a.TauEta)
	}
	if len(a.Planes) != 6 {
		t.Fatalf("expected six candidate planes, got %d", len(a.Planes))
	}
	if a.Sieve.UniformDoubletPlanes != 0 || a.Sieve.ChiBreaksDegeneracy {
		t.Fatalf("chi should not select a weak plane: %+v", a.Sieve)
	}
	for _, p := range a.Planes {
		if p.DoubletPlusDimC != 4 || p.DoubletMinusDimC != 4 || p.SingletPlusDimC != 4 || p.SingletMinusDimC != 4 {
			t.Fatalf("unexpected chi split for plane %+v", p)
		}
	}
	if a.Summary.PhysicalChiralityDerived || a.Summary.GlobalHDerived {
		t.Fatalf("summary overclaims physical chirality/H: %+v", a.Summary)
	}
}
