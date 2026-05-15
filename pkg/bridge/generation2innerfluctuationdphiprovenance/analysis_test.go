package generation2innerfluctuationdphiprovenance

import "testing"

func TestGate499InnerFluctuationDphiProvenance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Inheritance.FullScalarSU2NotSelectedByResponse || !a.Inheritance.NativeDphiStillOpen {
		t.Fatalf("Gate499 must inherit Gate498 SU2/Dphi obstruction: %+v", a.Inheritance)
	}
	if !a.InnerFluctuation.GaugeBosonContentRecovered || a.InnerFluctuation.GaugeBosonDimension != 12 || !a.InnerFluctuation.StructuralFieldContent {
		t.Fatalf("expected structural gauge content from inner fluctuations: %+v", a.InnerFluctuation)
	}
	if !a.InnerFluctuation.HiggsDoubletRecovered || a.InnerFluctuation.ComplexDoublets != 1 || a.InnerFluctuation.RealScalarDimension != 4 || a.InnerFluctuation.FiniteOneFormEdges != 4 {
		t.Fatalf("expected exactly one complex finite-one-form Higgs doublet: %+v", a.InnerFluctuation)
	}
	if !a.Dphi.StructuralDphiSocketFound || !a.Dphi.ScalarSU2RepresentationProvenanceClosed {
		t.Fatalf("expected structural Dphi/representation socket: %+v", a.Dphi)
	}
	if a.Dphi.NativeDphiActionDerived || a.Dphi.ProductGeometryKineticProjectionDerived || a.Dphi.PhysicalMassMatrixDerived {
		t.Fatalf("Gate499 must not promote Dphi action or W/Z masses: %+v", a.Dphi)
	}
	if !a.Reconciliation.NoContradiction || !a.Reconciliation.RepresentationVsResponseSeparated || !a.Reconciliation.NativeGaugeEatingStillBlocked {
		t.Fatalf("expected response/representation reconciliation with native gauge eating blocked: %+v", a.Reconciliation)
	}
	if !a.Boundary.StructuralScalarDoubletProvenancePromoted || !a.Boundary.StructuralDphiTransformationSocketPromoted {
		t.Fatalf("expected structural provenance promotion: %+v", a.Boundary)
	}
	if a.Boundary.NativeFullDphiActionClosed || a.Boundary.NativeKappaSelected || a.Boundary.NativeWZMassMatrixDerived {
		t.Fatalf("boundary over-promoted native electroweak data: %+v", a.Boundary)
	}
	if a.Firewall.ObservedWMassImported || a.Firewall.WeakAngleImported || a.Firewall.HiggsVEVImported || a.Firewall.NativeWZMassWritten {
		t.Fatalf("firewall violation: %+v", a.Firewall)
	}
	if a.Next.Gate != 500 {
		t.Fatalf("expected Gate500 redirect, got %+v", a.Next)
	}
}

func TestGate499TheoremPasses(t *testing.T) {
	res := Generation2InnerFluctuationDphiProvenanceAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem failed:\n%s", res.Details())
	}
}
