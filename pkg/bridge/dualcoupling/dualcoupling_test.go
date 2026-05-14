package dualcoupling

import "testing"

func TestDualCarrierCouplingSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.NaiveTensorDimension != 64 {
		t.Fatalf("expected naive 64D tensor, got %d", a.NaiveTensorDimension)
	}
	if !a.Direct64TensorRejected {
		t.Fatalf("raw tensor should be rejected")
	}
	if !a.ColorContactCouplingRejected || !a.LeptoquarkContactCouplingRejected {
		t.Fatalf("color/leptoquark direct contact couplings should be rejected")
	}
	if !a.AbelianBridgeDomainExposed || a.AbelianBridgeDimension != 2 {
		t.Fatalf("expected 2D abelian bridge domain, got exposed=%v dim=%d", a.AbelianBridgeDomainExposed, a.AbelianBridgeDimension)
	}
	if a.CouplingTensorSelected || a.CouplingActionDerived || a.DualCarrierHessianComputable || a.CondensationClaimAllowed {
		t.Fatalf("bridge claims should remain open: %+v", a)
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("no observed physical input should be used")
	}
}
