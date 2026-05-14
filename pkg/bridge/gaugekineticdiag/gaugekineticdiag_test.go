package gaugekineticdiag

import "testing"

func TestGaugeKineticDiagWhitensBrokenMetric(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CandidatePositive {
		t.Fatalf("candidate Hessian must be positive")
	}
	if !a.WhitenedExact {
		t.Fatalf("expected diag(1,1,4) to whiten raw metric, condition=%g", a.WhitenedCondition)
	}
	if a.SelectedByFiniteAction {
		t.Fatalf("Gate 94 must not claim finite action selection")
	}
	if a.PhysicalCouplingsDerived || a.PhysicalMassesDerived {
		t.Fatalf("Gate 94 must not claim physical couplings or masses")
	}
}
