package characteristicpullback

import "testing"

func TestGate244TracesTauEtaOriginButRejectsModeMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Origin.ExactOperatorOriginsRecovered || !a.Origin.StableNativeDegrees {
		t.Fatalf("origin should be traced exactly: %s", FormatOrigin(a.Origin))
	}
	if a.SpatialAlignment.NativeOperatorToModeMapDerived || !a.SpatialAlignment.ManualMapRejected {
		t.Fatalf("operator-to-mode map should be rejected: %s", FormatSpatialAlignment(a.SpatialAlignment))
	}
	if a.CharacteristicRep.RepresentativeConstructed || !a.CharacteristicRep.HypotheticalRepresentativeRejected {
		t.Fatalf("exterior representative should not be constructed: %s", FormatCharacteristic(a.CharacteristicRep))
	}
}

func TestGate244LeavesWeakAndGenerationConditional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.WeakPlane.PhysicalWeakPlaneDerived || a.WeakPlane.S3DegeneracyBroken || a.WeakPlane.GlobalHSummandDerived {
		t.Fatalf("weak plane should remain un-derived: %s", FormatWeak(a.WeakPlane))
	}
	if !a.Generation.DistinctEigenvalueCapacity || a.Generation.GenerationOperatorDerived || a.Generation.GenerationTextureDerived {
		t.Fatalf("generation should remain capacity-only: %s", FormatGeneration(a.Generation))
	}
	if a.Firewall.ForcedExteriorRepresentative || a.Firewall.PromotedScalarTraceToMatrix || a.Firewall.FiniteCorePolluted {
		t.Fatalf("firewall leaked: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := CharacteristicClassOperatorToModePullbackAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
