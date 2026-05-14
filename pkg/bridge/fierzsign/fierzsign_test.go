package fierzsign

import (
	"math"
	"testing"
)

func TestFierzSignConstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.SigmaBarSigmaIdentityResidual > 1e-12 {
		t.Fatalf("sigma identity residual too large: %.3e", a.SigmaBarSigmaIdentityResidual)
	}
	if math.Abs(a.UniversalLRScalarFierzCoefficient+2) > 1e-12 {
		t.Fatalf("expected LR scalar Fierz coefficient -2, got %.12f", a.UniversalLRScalarFierzCoefficient)
	}
	if !a.SignedCliffordFierzCoefficientsDerived {
		t.Fatalf("expected signed Fierz coefficients to be derived")
	}
	if a.AttractiveScalarChannelSignDerived || a.NativeFourFermionKernelDerived {
		t.Fatalf("Gate 61 must not claim attraction or native kernel")
	}
}
