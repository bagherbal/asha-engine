package generation2scalarsu2provenance

import "testing"

func TestGate498ScalarSU2ComplexStructureProvenance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.BridgeGaugeQuotientClosed || !a.Inheritance.NativeGaugeOrbitStillOpen || !a.Inheritance.NativeDphiStillOpen {
		t.Fatalf("Gate498 must inherit Gate497 bridge quotient with native orbit/Dphi still open: %+v", a.Inheritance)
	}
	if !a.ComplexStructure.AbstractComplexDoubletSocket || !a.ComplexStructure.JCompatibleWithPairPlanes {
		t.Fatalf("expected compatible complex doublet socket: %+v", a.ComplexStructure)
	}
	if !nearlyZero(a.ComplexStructure.JSkewResidual) || !nearlyZero(a.ComplexStructure.JSquarePlusIResidual) || !nearlyZero(a.ComplexStructure.ScalarResponseCommJNorm) {
		t.Fatalf("bad complex structure residuals: %+v", a.ComplexStructure)
	}
	if a.ComplexStructure.ComplexStructureNativelyUnique {
		t.Fatalf("complex structure must not be promoted as natively unique")
	}
	if !a.SU2Action.AbstractDoubletRepresentation || !nearlyZero(a.SU2Action.SU2ClosureResidual) {
		t.Fatalf("expected abstract SU2 closure: %+v", a.SU2Action)
	}
	if !a.SU2Action.PairRotationU1Selected || a.SU2Action.FullSU2SelectedByScalarResponse || a.SU2Action.FullSU2ActionNativeSelected {
		t.Fatalf("scalar response must select pair U1 but not full SU2: %+v", a.SU2Action)
	}
	if !(a.SU2Action.CommT1Norm > eps && a.SU2Action.CommT2Norm > eps && a.SU2Action.CommT3Norm < eps) {
		t.Fatalf("unexpected scalar-response commutators: T1=%g T2=%g T3=%g", a.SU2Action.CommT1Norm, a.SU2Action.CommT2Norm, a.SU2Action.CommT3Norm)
	}
	if !a.GaugeOrbit.BridgeGoldstoneOrbitConsistent || a.GaugeOrbit.NativeGaugeOrbitSelected || a.GaugeOrbit.WZNativeMassPromotionAllowed {
		t.Fatalf("Goldstone orbit should remain bridge-consistent but not native: %+v", a.GaugeOrbit)
	}
	if a.Boundary.FullScalarSU2NativeSelected || a.Boundary.NativeDphiClosed || a.Boundary.NativeWZMassMatrixDerived {
		t.Fatalf("Gate498 must not close native SU2/Dphi/WZ: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.NativeScalarSU2Written || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
}
