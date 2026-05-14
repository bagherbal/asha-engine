package scalarbundlemap

import "testing"

func TestBuildDefaultScalarBundleMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.TargetCarrier.PhysicalProjectorPairExists || !a.TargetCarrier.DimensionallyCompatible {
		t.Fatalf("expected physical high/low 2+2 projectors: %s", FormatTargetCarrier(a.TargetCarrier))
	}
	if !a.Intertwiner.BranchwiseIntertwinersExist || a.Intertwiner.CanonicalAssignmentDerived || a.Intertwiner.PhysicalScalarBundleMapDerived {
		t.Fatalf("expected only conditional intertwiners: %s", FormatIntertwiner(a.Intertwiner))
	}
	if a.Sources.EtaOddSourceFound || a.Sources.CanonicalOrientationDerived || a.Sources.PhysicalHighLowPullbackFound {
		t.Fatalf("unexpected eta-odd orientation source: %s", FormatSources(a.Sources))
	}
	if a.Trivialization.UniqueIntertwiner || a.Trivialization.CanonicalChangeOfBasisDerived {
		t.Fatalf("unexpected unique trivialization: %s", FormatTrivialization(a.Trivialization))
	}
	if a.Firewall.PhysicalScalarBundleDerived || a.Firewall.ChernWeilCarrierDerived || a.Firewall.HeatKernelMatchingDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestPhysicalProjectorLaws(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.TargetCarrier.HighProjector.Trace != 2 || a.TargetCarrier.LowProjector.Trace != 2 {
		t.Fatalf("expected trace-2 projectors: %s", FormatTargetCarrier(a.TargetCarrier))
	}
	if a.TargetCarrier.HighProjector.IdempotenceResidual != 0 || a.TargetCarrier.LowProjector.IdempotenceResidual != 0 {
		t.Fatalf("expected exact idempotents: %s", FormatTargetCarrier(a.TargetCarrier))
	}
	if !a.TargetCarrier.ProjectorsOrthogonal || !a.TargetCarrier.ProjectorsComplete {
		t.Fatalf("expected orthogonal complete target projectors: %s", FormatTargetCarrier(a.TargetCarrier))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ScalarBundleMapHphiProjectorIdentificationAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
