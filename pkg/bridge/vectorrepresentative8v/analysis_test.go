package vectorrepresentative8v

import "testing"

func TestGate248RejectsManualScalarTo8VRepresentative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.VectorBasis.NativeCarrierKnown || a.VectorBasis.Dimension != 8 {
		t.Fatalf("expected native 8_v carrier: %s", FormatVectorBasis(a.VectorBasis))
	}
	if !a.ScalarBundle.TraceOriginKnown || !a.ScalarBundle.DimensionallyEmbeddable {
		t.Fatalf("expected scalar trace origin and dimensional embeddability: %s", FormatScalarBundle(a.ScalarBundle))
	}
	if a.ScalarVectorMap.NativeMapDerived || a.ScalarVectorMap.QZTYToBasisDerived || !a.ScalarVectorMap.ManualAssignmentRejected {
		t.Fatalf("expected scalar-to-8v map to remain blocked: %s", FormatScalarVectorMap(a.ScalarVectorMap))
	}
	if a.VTau.Constructed || a.VTau.LawfulRepresentative || a.Triality.VTauAvailable || a.Triality.SpinorTextureConstructed {
		t.Fatalf("v_tau/triality must remain unconstructed: %s / %s", FormatVTau(a.VTau), FormatTriality(a.Triality))
	}
	if a.Firewall.ForcedHphiTo8VMap || a.Firewall.ConstructedVTauByHand || a.Firewall.InsertedYukawaTexture || a.Firewall.ClaimedFiniteFlavorTheorem {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
