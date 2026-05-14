package chiralweakselector

import "testing"

func TestGate238BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Summary.GammaParityAvailable || !a.Summary.TemporalSpatialClasses {
		t.Fatalf("expected gamma preflight and temporal/spatial classes: %+v", a.Summary)
	}
	if a.Summary.GammaSelectsPlane || a.Summary.UniqueWeakPlaneDerived || a.Summary.GlobalHDerived || a.Summary.PhysicalLeftActionDerived {
		t.Fatalf("Gate 238 must not overclaim weak plane/chiral/global-H derivation: %+v", a.Summary)
	}
	if len(a.Planes) != 6 {
		t.Fatalf("expected six planes, got %d", len(a.Planes))
	}
	for _, p := range a.Planes {
		if p.DoubletEvenDimC != 4 || p.DoubletOddDimC != 4 || p.SingletEvenDimC != 4 || p.SingletOddDimC != 4 {
			t.Fatalf("unexpected parity split for %s: %+v", p.Plane, p)
		}
		if p.DoubletsUniformParity || p.SingletsUniformParity || p.SU2ActsOnlyOnOneParity == true {
			t.Fatalf("plane should not be chirally uniform: %+v", p)
		}
	}
}

func TestGate238Theorem(t *testing.T) {
	th := ChiralAlignmentWeakPlaneSelectorAuditTheorem()
	r := th.Run()
	if len(r.Checks) == 0 {
		t.Fatal("expected theorem checks")
	}
	for _, c := range r.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
