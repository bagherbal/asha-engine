package generation2weakdoubletcarrieridentityspatialcp1nonidentificationaudit

import "testing"

func TestGate576FiniteWeakDoubletCarrierIdentityAndSpatialCP1NonidentificationAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SealedCP1SplitExists || !a.Inherited.CommutesWithBMinusL || a.Inherited.CarriesImHAction || a.Inherited.PartOfFiniteWeakCarrier || a.Inherited.CanBePhysicalWeakPlane {
		t.Fatalf("bad inherited Gate575 boundary: %s", FormatInherited(a.Inherited))
	}
	if a.FiniteAlgebra.Algebra == "" || !a.FiniteAlgebra.QuaternionicWeakSocket || a.FiniteAlgebra.ImHLieAlgebra != "Im(H) ≅ su(2)_L structurally" || a.FiniteAlgebra.AbsoluteDynamicsDerived {
		t.Fatalf("bad finite algebra audit: %s", FormatFiniteAlgebra(a.FiniteAlgebra))
	}
	if !a.WeakFermions.LLPresent || !a.WeakFermions.QLPresent || !a.WeakFermions.HActsOnLL || !a.WeakFermions.HActsOnQL || a.WeakFermions.QLColorMultiplicity != 3 || a.WeakFermions.ColorIsWeakStructure || a.WeakFermions.SealedSpatialCP1Used {
		t.Fatalf("bad weak fermion inventory: %s", FormatWeakFermions(a.WeakFermions))
	}
	if a.ScalarDoublet.CarrierName != "H_phi" || a.ScalarDoublet.ComplexDimension != 2 || a.ScalarDoublet.RealDimension != 4 || !a.ScalarDoublet.HActionStructural || !a.ScalarDoublet.SeparateFromWSpatial || !a.ScalarDoublet.SeparateFromUperp || a.ScalarDoublet.SealedSpatialCP1Used {
		t.Fatalf("bad scalar doublet inventory: %s", FormatScalarDoublet(a.ScalarDoublet))
	}
	if !a.SealedCompare.CP1SplitExistsAlgebraically || a.SealedCompare.IsFiniteWeakCarrier || a.SealedCompare.AppearsInAFRepresentation || a.SealedCompare.AppearsInDFEdges || a.SealedCompare.AppearsInFirstOrder || a.SealedCompare.AppearsInOneFormHiggsLane {
		t.Fatalf("sealed CP1 was incorrectly promoted: %s", FormatSealedCompare(a.SealedCompare))
	}
	if a.WeakCount.TotalWeakDoublets != 4 || a.WeakCount.LeptonWeakDoublets != 1 || a.WeakCount.QuarkWeakDoublets != 3 || !a.WeakCount.ComesFromColorMultiplicity || a.WeakCount.ComesFromSpatialCP1Selection {
		t.Fatalf("bad weak doublet count: %s", FormatWeakCount(a.WeakCount))
	}
	if !a.EdgeLane.CanonicalEdgesReconfirmed || !a.EdgeLane.FirstOrderCompatible || !a.EdgeLane.UsesHPhiScalarLane || a.EdgeLane.UsesSealedSpatialSelector || a.EdgeLane.UsesUperpCarrier {
		t.Fatalf("bad edge lane: %s", FormatEdgeLane(a.EdgeLane))
	}
	if !a.NonIdentity.Certified || a.NonIdentity.UperpEqualsHPhi || a.NonIdentity.UperpEqualsLL || a.NonIdentity.UperpEqualsQL || a.NonIdentity.UperpEqualsImH {
		t.Fatalf("bad nonidentity theorem: %s", FormatNonIdentity(a.NonIdentity))
	}
	if a.Firewalls.PhysicalWeakPlaneDerived || a.Firewalls.WeakIsospinDerivedFromCP1 || a.Firewalls.WZPhotonDynamicsDerived || a.Firewalls.MassesDerived || a.Firewalls.GenerationHierarchy || a.Firewalls.YukawaTexture || a.Firewalls.CKMPMNS || a.Firewalls.ObservedFlavorData {
		t.Fatalf("bad firewall: %s", FormatFirewalls(a.Firewalls))
	}
	if a.Final.SealedSpatialCP1IsWeakCarrier || !a.Final.HPhiIsScalarWeakDoublet || !a.Final.WeakDoubletOnePlusThreeFromColor || a.Final.DerivesPhysicalWeakFlavorEWData || a.Final.AdditionalTheoremRequired == "" {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate576ContainsAllEdges(t *testing.T) {
	edges := []string{"Q_L↔u_R", "Q_L↔d_R", "L_L↔e_R", "L_L↔ν_R"}
	if !containsAllEdges(edges) {
		t.Fatalf("expected canonical edge set")
	}
	if containsAllEdges(edges[:3]) {
		t.Fatalf("incomplete edge set should fail")
	}
}

func TestGate576Theorem(t *testing.T) {
	res := Generation2FiniteWeakDoubletCarrierIdentityAndSpatialCP1NonidentificationAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
