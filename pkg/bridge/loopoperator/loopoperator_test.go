package loopoperator

import "testing"

func TestFiniteLoopOperatorConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.FiniteYukawaIncidenceOperatorDerived {
		t.Fatalf("expected finite Yukawa incidence operator")
	}
	if a.DomainDimension != 32 || a.RightDimension != 8 {
		t.Fatalf("unexpected dimensions Y:%dx%d", a.RightDimension, a.DomainDimension)
	}
	if a.AllowedFiberEntries != 16 {
		t.Fatalf("expected 16 allowed fiber entries, got %d", a.AllowedFiberEntries)
	}
	if a.Rank != 8 {
		t.Fatalf("expected rank 8, got %d", a.Rank)
	}
	if a.RightTrace != 16 || a.DomainTrace != 16 {
		t.Fatalf("expected trace 16, got right=%v domain=%v", a.RightTrace, a.DomainTrace)
	}
	if a.UnitTopLikeSkeleton != -6 {
		t.Fatalf("expected -6 top-like skeleton, got %v", a.UnitTopLikeSkeleton)
	}
	if a.TopDominanceSelected || a.MuSquaredSignDerived || a.NativeEffectivePotentialComputed {
		t.Fatalf("operator must not claim top dominance or native μ² sign")
	}
	if a.HiddenObservedCouplingsUsed {
		t.Fatalf("must not use observed couplings")
	}
}
