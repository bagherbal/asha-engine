package exchangekernel

import (
	"math"
	"testing"
)

func TestExchangeKernelAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SignedFierzAvailable || !a.ConditionalAttractiveBranchAvailable {
		t.Fatalf("expected signed Fierz and conditional attractive branch")
	}
	if math.Abs(a.UnitExchangePlusTotal+2) > 1e-12 {
		t.Fatalf("expected +J^2 branch total -2, got %.12f", a.UnitExchangePlusTotal)
	}
	if math.Abs(a.UnitExchangeMinusTotal-2) > 1e-12 {
		t.Fatalf("expected -J^2 branch total +2, got %.12f", a.UnitExchangeMinusTotal)
	}
	if a.AttractiveScalarChannelSignDerived || a.NativeFourFermionKernelDerived || a.CondensationClaimAllowed {
		t.Fatalf("Gate 62 must not claim derived attraction, native kernel, or condensation")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("Gate 62 must not use observed input")
	}
}
