package neutraleigenspacekernel

import "testing"

func TestGate249BlocksNeutralKernelWithoutQ8V(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.VectorCarrier.NativeCarrierKnown || a.VectorCarrier.Dimension != 8 {
		t.Fatalf("expected native 8_v carrier: %s", FormatVectorCarrier(a.VectorCarrier))
	}
	if a.EWAction.QMatrixOn8VDerived || a.EWAction.ZMatrixOn8VDerived || !a.EWAction.ManualRepresentationRejected {
		t.Fatalf("expected Q/Z action on 8_v to remain blocked: %s", FormatEWAction(a.EWAction))
	}
	if a.NeutralKernel.Computed || a.NeutralKernel.DimensionKnown || a.NeutralKernel.Dimension != -1 || a.NeutralKernel.ExactlyThreeDimensional {
		t.Fatalf("neutral kernel should not be computable: %s", FormatNeutralKernel(a.NeutralKernel))
	}
	if a.ScalarPlane.CanonicalIsomorphism || a.ScalarPlane.QZTYToNeutralBasisDerived || a.VTau.LawfulRepresentative || a.Triality.TrialityCanRun {
		t.Fatalf("scalar plane/v_tau/triality must remain blocked: %s / %s / %s", FormatScalarPlane(a.ScalarPlane), FormatVTau(a.VTau), FormatTriality(a.Triality))
	}
	if a.Firewall.InventedQActionOn8V || a.Firewall.ForcedNeutralKernelDim3 || a.Firewall.ConstructedVTauByHand || a.Firewall.InsertedYukawaTexture {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
