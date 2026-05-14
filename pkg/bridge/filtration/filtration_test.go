package filtration

import "testing"

func TestFiniteFiltrationDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.ModeCount != 14 || a.ContinuumCount != 5 || a.OpenCount != 8 || a.VacuumCount != 1 {
		t.Fatalf("unexpected mode split: modes=%d continuum=%d open=%d vacuum=%d", a.ModeCount, a.ContinuumCount, a.OpenCount, a.VacuumCount)
	}
	if !a.StatusPreorderConstructed || len(a.Antichains) == 0 {
		t.Fatalf("expected status preorder with threshold-open antichain")
	}
	if !a.SpectralValueOrdersConstructed || !a.ReverseOrderEquallyCompatible || !a.NonUniqueFiltrationWitnessed {
		t.Fatalf("expected equally compatible ascending/descending finite filtrations")
	}
	if a.CanonicalTotalOrderDerived || a.CanonicalOrientationDerived || a.CanonicalCutoffDerived {
		t.Fatalf("Gate 107 must not derive canonical total order, orientation, or cutoff")
	}
	if !a.MonotonePredicateFamilyConstructed || a.DerivedActivationPredicate {
		t.Fatalf("expected monotone predicate family but no physical activation predicate")
	}
	if a.ThresholdCorrectedBetaDerived || a.DerivedDecouplingMatchingRule || a.NativeFiniteRGFunctorDerived {
		t.Fatalf("Gate 107 must not derive threshold-corrected beta flow")
	}
	if a.ResidualNullityAfter != a.ResidualNullityBefore || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual nullity changed unexpectedly: before=%d after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)
	}
}

func TestAscendingAndDescendingSelectDifferentFirstOpenModes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.FirstAscendingOpen == "" || a.FirstDescendingOpen == "" {
		t.Fatalf("missing first-open witnesses")
	}
	if a.FirstAscendingOpen == a.FirstDescendingOpen {
		t.Fatalf("ascending and descending orders should expose different first threshold-open modes")
	}
	if a.FirstAscendingOpen != "B-sector first spectral gap" {
		t.Fatalf("expected ascending order to start with B-gap, got %q", a.FirstAscendingOpen)
	}
	if a.FirstDescendingOpen != "contact partial-overlap mode 1" {
		t.Fatalf("expected descending order to start with largest contact overlap, got %q", a.FirstDescendingOpen)
	}
}
