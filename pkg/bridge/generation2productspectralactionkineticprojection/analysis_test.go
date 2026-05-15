package generation2productspectralactionkineticprojection

import "testing"

func TestGate500ProductSpectralActionKineticProjection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.StructuralDphiSocketFound || !a.Inheritance.ProductKineticProjectionWasOpen || !a.Inheritance.HeatKernelScalarCoefficientWasOpen {
		t.Fatalf("expected Gate499 structural Dphi socket with open kinetic gap: %+v", a.Inheritance)
	}
	if !a.ProductAction.CCMFormulaInstalled || !a.ProductAction.ProductGeometryRecognized || !a.ProductAction.StructuralClosure || a.ProductAction.FullNumericalClosure {
		t.Fatalf("expected CCM product action ledger without numerical closure: %+v", a.ProductAction)
	}
	if !a.KineticProjection.SymbolicKineticProjectionReadOff || !a.KineticProjection.ProductActionContainsDphiSquared || !a.KineticProjection.CanonicalScalarRescalingReadOff {
		t.Fatalf("expected symbolic scalar kinetic projection: %+v", a.KineticProjection)
	}
	if !a.KineticProjection.CoefficientDependsOnYukawaTraceA || a.KineticProjection.YukawaTraceANativelyNumeric || a.KineticProjection.HeatKernelCoefficientNumeric || a.KineticProjection.CanonicalI4MetricSelected || a.KineticProjection.NativeKineticProjectionClosed {
		t.Fatalf("Gate500 must not promote scalar kinetic normalization: %+v", a.KineticProjection)
	}
	if !a.YukawaAirlock.TraceASealedByFirewall || a.YukawaAirlock.YukawaNativeSelectors != 0 || a.YukawaAirlock.YukawaRankThreeDerived {
		t.Fatalf("expected Yukawa trace a to be sealed by Gate489: %+v", a.YukawaAirlock)
	}
	if !a.Boundary.SymbolicProductKineticProjectionAccept || a.Boundary.NativeScalarKineticCoefficientDerived || a.Boundary.NativeWZMassMatrixDerived || a.Boundary.NativeKappaSelected {
		t.Fatalf("boundary over-promoted symbolic kinetic projection: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.HiggsVEVImported || a.Firewall.YukawaImported || a.Firewall.NativeKineticWritten || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 501 {
		t.Fatalf("expected Gate501 redirect, got %+v", a.Next)
	}
}

func TestGate500TheoremPasses(t *testing.T) {
	res := Generation2ProductSpectralActionScalarKineticProjectionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
