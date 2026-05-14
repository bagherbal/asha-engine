package shellfunctor

import "testing"

func TestFiniteShellFunctorDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.ModeCount != 14 {
		t.Fatalf("expected 14 finite modes from threshold carrier, got %d", a.ModeCount)
	}
	if a.ContinuumCount != 5 || a.OpenCount != 8 || a.VacuumCount != 1 {
		t.Fatalf("unexpected mode split: continuum=%d open=%d vacuum=%d", a.ContinuumCount, a.OpenCount, a.VacuumCount)
	}
	if !a.NestedProjectionFamilyConstructed || !a.CompositionClosed || !a.AssociativityVerified {
		t.Fatalf("nested projection family was not constructed as a closed associative family")
	}
	if !a.IdempotentSemilatticeDerived {
		t.Fatalf("expected idempotent semilattice composition to be derived")
	}
	if a.AdditiveSemigroupDerived {
		t.Fatalf("additive/logarithmic RG semigroup must remain unproved")
	}
	if !a.NontrivialAdditiveCounterexample {
		t.Fatalf("expected a nontrivial counterexample to additive composition")
	}
	if a.NativeFiniteRGFunctorDerived || a.CanonicalScaleLogParameterDerived || a.ThresholdActivationPredicateDerived {
		t.Fatalf("Gate 106 must not derive physical RG data")
	}
	if a.ResidualNullityAfter != a.ResidualNullityBefore || a.ResidualNullityAfter != 3 {
		t.Fatalf("residual nullity changed unexpectedly: before=%d after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)
	}
}

func TestProjectionCompositionIsMinNotAdditive(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	found := false
	for _, w := range a.CompositionTable {
		if w.Left == 1 && w.Right == 2 {
			found = true
			if !w.Closed || w.Composed != 1 {
				t.Fatalf("expected C1∘C2=C1 closed under min, got C%d closed=%t", w.Composed, w.Closed)
			}
			if w.AdditiveOK {
				t.Fatalf("C1∘C2 must not equal additive C3")
			}
		}
	}
	if !found {
		t.Fatalf("missing C1∘C2 witness")
	}
}
