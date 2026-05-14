package liecarrierprojection

import "testing"

func TestGate245DecompositionFailsCarrierProjection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.OperatorDecomposition.EWDecompositionTraced || !a.OperatorDecomposition.QZMixT3AndYPhi {
		t.Fatalf("expected EW decomposition through T3/Y_phi: %s", FormatOperatorDecomposition(a.OperatorDecomposition))
	}
	if a.OperatorDecomposition.SlotsAreThreeSU2BasisElements || !a.OperatorDecomposition.MissingT1T2SlotOrigins {
		t.Fatalf("tau_eta slots must not be promoted to su(2) basis: %s", FormatOperatorDecomposition(a.OperatorDecomposition))
	}
	if !a.DerivationBlade.CandidateSU2Capacity || a.DerivationBlade.OneToOneDerivationAxisMap {
		t.Fatalf("expected su2 capacity but no derived axis map: %s", FormatDerivationBlade(a.DerivationBlade))
	}
	if a.CarrierProjection.ChainedProjectionDerived || a.CarrierProjection.ExteriorRepresentativeConstructed || !a.CarrierProjection.HypotheticalProjectionRejected {
		t.Fatalf("projection should remain rejected: %s", FormatCarrierProjection(a.CarrierProjection))
	}
	if a.Summary.WeakPlaneDerived || a.Summary.GlobalHDerived || a.Summary.GenerationTextureDerived {
		t.Fatalf("Gate 245 must not derive weak plane, global H, or generation texture: %s", FormatSummary(a.Summary))
	}
}
