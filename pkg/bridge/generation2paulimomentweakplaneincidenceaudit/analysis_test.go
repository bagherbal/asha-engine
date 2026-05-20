package generation2paulimomentweakplaneincidenceaudit

import "testing"

func TestGate561PauliMomentWeakPlaneIncidenceAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate560PauliTriplet || !a.Inherited.Gate560HopfIdentity || !a.Inherited.Gate560NoTransferFunctor {
		t.Fatalf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.SpatialLabels.InsideBLSpatialEigenspace || !a.SpatialLabels.BasisConventionOnly || a.SpatialLabels.NativeOrientedMetricThreeSpace || a.SpatialLabels.MetricCertificateAvailable || a.SpatialLabels.OrientationCertificateAvailable {
		t.Fatalf("spatial label audit failed: %s", FormatSpatialLabels(a.SpatialLabels))
	}
	if !a.Incidence.CoordinateBivectorsAvailable || a.Incidence.NativeIncidenceSelector || !a.Incidence.NotationalOnly || len(a.Incidence.PlaneToBivector) != 3 {
		t.Fatalf("incidence audit failed: %s", FormatIncidence(a.Incidence))
	}
	if !a.Hodge.FormalNormalSelectsPlane || a.Hodge.NativeHodgeStarConstructed || a.Hodge.MetricAvailableNatively || a.Hodge.OrientationAvailableNatively {
		t.Fatalf("hodge audit failed: %s", FormatHodge(a.Hodge))
	}
	if a.Intertwiner.MapToSpatialFound || a.Intertwiner.MapToIncidenceFound || a.Intertwiner.BasisIndependent || a.Intertwiner.UnitMetricCompatible {
		t.Fatalf("intertwiner unexpectedly available: %s", FormatIntertwiner(a.Intertwiner))
	}
	if a.Plane.CanonicalU12 || a.Plane.CanonicalU13 || a.Plane.CanonicalU23 || !a.Plane.OnlyBasisDependentPlane {
		t.Fatalf("canonical plane audit failed: %s", FormatPlane(a.Plane))
	}
	if !a.BL.FormalSelectionCommutesWithBL || a.BL.CompatibilityNontrivial || a.BL.BLSuppliesPlaneLabels || a.BL.MixesLeptonSlot {
		t.Fatalf("B-L audit failed: %s", FormatBL(a.BL))
	}
	if a.Spectral.IncidenceFunctorFound || a.Spectral.CompatibilityPassed || a.Spectral.FiniteOneFormRelationFound {
		t.Fatalf("spectral compatibility unexpectedly available: %s", FormatSpectral(a.Spectral))
	}
	if !a.Firewall.Preserved || a.Firewall.WeakIsospinIdentified || a.Firewall.GaugeBosonsIdentified || a.Firewall.GenerationHierarchyIdentified || a.Firewall.YukawaTextureDerived || a.Firewall.CKMPMNSDerived || a.Firewall.ObservedFlavorImported || a.Firewall.HiggsLanePromoted {
		t.Fatalf("firewall failed: %s", FormatFirewall(a.Firewall))
	}
	if a.Final.SSpatialNativeOrientedMetric || a.Final.WeakPlanesNativeBivectors || a.Final.HodgeStarNative || a.Final.PauliToIncidenceIntertwiner || a.Final.ScalarMomentSelectsWeakPlane || a.Final.LawfulTransferAvailable {
		t.Fatalf("final verdict failed: %s", FormatFinal(a.Final))
	}
}

func TestGate561Theorem(t *testing.T) {
	res := Generation2PauliMomentWeakPlaneIncidenceIntertwinerAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
