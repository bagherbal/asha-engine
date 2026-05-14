package spinctwistedchirality

import "testing"

func TestGate240SpinCTwistedChiralityAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.U1.ActsOnSC || a.U1.ImportedSMHypercharge || len(a.U1.ModeWeights) != 4 {
		t.Fatalf("native u1 should be retrieved without SM import: %+v", a.U1)
	}
	if !a.Twist.DistinctFromGamma || a.Twist.IsInvolution || a.Twist.PhysicalChiralityDerived {
		t.Fatalf("twist should be diagnostic, distinct, and non-involutive: %+v", a.Twist)
	}
	if len(a.Planes) != 6 {
		t.Fatalf("expected 6 planes, got %d", len(a.Planes))
	}
	if len(a.Sieve.U1PreservingPlanes) != 3 || len(a.Sieve.U1RejectedPlanes) != 3 || !a.Sieve.TemporalSpatialRejected {
		t.Fatalf("u1 sieve should leave exactly three pure-spatial planes: %+v", a.Sieve)
	}
	if len(a.Sieve.UniformDoubletPlanes) != 0 || len(a.Sieve.SelectedPlanes) != 0 || a.Sieve.TwistBreaksDegeneracy {
		t.Fatalf("twist must not select a weak plane: %+v", a.Sieve)
	}
	for _, p := range a.Planes {
		if p.PlaneClass == "pure-spatial" && !p.SU2PreservesY {
			t.Fatalf("pure-spatial plane should preserve native u1: %+v", p)
		}
		if p.PlaneClass == "temporal-spatial" && p.SU2PreservesY {
			t.Fatalf("temporal-spatial plane should fail native u1 commutation: %+v", p)
		}
		if p.DoubletsUniformTwist {
			t.Fatalf("no plane should have uniform twisted doublets: %+v", p)
		}
	}
	if a.Summary.PhysicalChiralityDerived || a.Summary.GlobalHDerived || a.Summary.UniqueWeakPlaneDerived {
		t.Fatalf("summary overclaims chirality/H/plane: %+v", a.Summary)
	}
}
