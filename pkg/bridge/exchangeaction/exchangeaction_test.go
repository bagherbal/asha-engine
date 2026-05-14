package exchangeaction

import (
	"math"
	"testing"
)

func TestFiniteExchangeActionPropagatorSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.UnitPropagatorBranchAvailable || !a.InverseKineticBranchAvailable || !a.KineticWeightedBranchAvailable {
		t.Fatalf("expected diagnostic branches to be available")
	}
	if math.Abs(a.UnitAttractiveKernel-2) > 1e-12 {
		t.Fatalf("expected unit diagnostic kernel 2, got %.12f", a.UnitAttractiveKernel)
	}
	if a.ExchangeActionSignDerived || a.PropagatorWeightsDerived || a.FiniteExchangeKernelDerived || a.CondensationClaimAllowed {
		t.Fatalf("Gate 63 must not claim exchange sign, propagators, native kernel, or condensation")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("Gate 63 must not use observed input")
	}
	if len(a.SectorDiagnostics) != 4 {
		t.Fatalf("expected four sector diagnostics, got %d", len(a.SectorDiagnostics))
	}
}
