package su2spinorlift

import "testing"

func TestGate237BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Summary.CandidateSU2Lifts || !a.Summary.DoubletDimensionalSupport || !a.Summary.PseudoRealLocalHSupport {
		t.Fatalf("expected candidate lift/doublet/local-H support: %+v", a.Summary)
	}
	if a.Summary.GlobalHDerived || a.Summary.ExactSMAlgebraDerived || a.Summary.CanonicalWeakPlane {
		t.Fatalf("Gate 237 must not overclaim H/SM algebra/weak plane derivation: %+v", a.Summary)
	}
	if len(a.Planes) != 6 {
		t.Fatalf("expected six two-mode planes, got %d", len(a.Planes))
	}
	for _, p := range a.Planes {
		if p.DoubletStateDimC != 8 || p.SingletStateDimC != 8 || !p.LocalHModule {
			t.Fatalf("bad plane audit: %+v", p)
		}
	}
}

func TestGate237Theorem(t *testing.T) {
	th := SU2SpinorLiftQuaternionicClosureAuditTheorem()
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
