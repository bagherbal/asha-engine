package dualcarrier

import "testing"

func TestDualCarrierArchitecture(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DualCarrierSplitDefined || !a.ForcedEmbeddingRejected {
		t.Fatalf("expected dual-carrier split with rejected forced embedding")
	}
	if !a.PatiSalamCarrierPreservesU4 || a.PatiSalamCarrier.Dimension != 16 {
		t.Fatalf("bad Pati-Salam carrier: %+v", a.PatiSalamCarrier)
	}
	if !a.ContactCarrierPreservesEWSeed || a.ContactCarrier.Dimension != 4 {
		t.Fatalf("bad contact carrier: %+v", a.ContactCarrier)
	}
	if a.CouplingTensorDimension != 64 {
		t.Fatalf("unexpected coupling tensor dimension: %d", a.CouplingTensorDimension)
	}
	if a.CouplingTensorSelected || a.CouplingActionDerived || a.CurrentHessianComputable || a.CondensationClaimAllowed {
		t.Fatalf("bridge claims should remain open: %+v", a)
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("no observed physical input should be used")
	}
}
