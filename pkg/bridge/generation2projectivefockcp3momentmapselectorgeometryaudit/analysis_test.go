package generation2projectivefockcp3momentmapselectorgeometryaudit

import "testing"

func TestGate572ProjectiveFockCP3MomentMapSelectorGeometryAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Projective.ProjectiveQuotientCertified || a.Projective.BaseRealDimension != 6 || !a.Projective.FubiniStudyAvailable || a.Projective.PhysicalSpacetime {
		t.Fatalf("bad projective quotient: %s", FormatProjective(a.Projective))
	}
	if !a.Phase.ReebMatchesTotalNumber || !a.Phase.TrivialOnCP3 || !a.Phase.LawSpacePhaseOnly || a.Phase.PhysicalLorentzianTime || a.Phase.RGScale {
		t.Fatalf("bad phase quotient: %s", FormatPhase(a.Phase))
	}
	if !a.Selector.PhaseInvariant || !a.Selector.ComplexScaleInvariant || a.Selector.MaxInvarianceResidual > 1e-12 || !a.Selector.DefinesMomentFunctions || a.Selector.PhysicalHamiltonianFlow {
		t.Fatalf("bad selector moment: %s", FormatSelector(a.Selector))
	}
	if !a.BMinusL.CriticalStrataCertified || !a.BMinusL.ProjectiveOnePlusThree || a.BMinusL.WeakPlaneSelected || a.BMinusL.GenerationSelected {
		t.Fatalf("bad B-L moment: %s", FormatBMinusL(a.BMinusL))
	}
	if !a.Stabilizer.MatchesGate555Commutant || a.Stabilizer.StabilizerDimension != 10 || !a.Stabilizer.HomogeneousDimensionMatchesCP3 {
		t.Fatalf("bad stabilizer: %s", FormatStabilizer(a.Stabilizer))
	}
	if !a.SpatialBlock.NativeProjectiveRefinement || a.SpatialBlock.WeakPlaneSelected || !a.SpatialBlock.RequiresSecondSelector {
		t.Fatalf("bad spatial block: %s", FormatSpatialBlock(a.SpatialBlock))
	}
	if a.Second.CurrentNativeSecondSelector || a.Second.Gate555UniqueWeakPlane || a.Second.TauEtaPulledBackNative || a.Second.SpatialTwoPlusOneDerived {
		t.Fatalf("bad second selector obstruction: %s", FormatSecond(a.Second))
	}
	if !a.K7.Gate571BoundaryPreserved || a.K7.CP3ToK7FunctorFound || a.K7.HopfS7ToK7FunctorFound || a.K7.TotalPhaseToK7Action {
		t.Fatalf("bad K7 firewall: %s", FormatK7(a.K7))
	}
	if a.Time.MomentFlowPhysicalTime || a.Time.MomentFlowOSHilbert || a.Time.MomentFlowRGScale || a.Time.MomentFlowSpacetime || a.Time.MomentFlowObservedHistory {
		t.Fatalf("bad time firewall: %s", FormatTime(a.Time))
	}
	if a.FlavorEW.YukawaEigenvaluesDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.PhotonDynamicsDerived || a.FlavorEW.WZMassesDerived || a.FlavorEW.ObservedDataImported {
		t.Fatalf("bad flavor/electroweak firewall: %s", FormatFlavorEW(a.FlavorEW))
	}
	if !a.Final.CP3Certified || !a.Final.FubiniStudyAvailable || !a.Final.SelectorMomentFunctionsOnCP3 || !a.Final.BMinusLProjectiveCP0CP2Split || !a.Final.MatchesGate555Commutant || a.Final.NativeSecondSelectorOnCP2 || a.Final.K7RelationOrPhysicalTimeProven {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate572RayleighMomentRejectsZeroVector(t *testing.T) {
	if _, err := RayleighMoment([]float64{1, 2}, []complex128{0, 0}); err == nil {
		t.Fatal("expected zero vector rejection")
	}
}

func TestGate572Theorem(t *testing.T) {
	res := Generation2ProjectiveFockCP3MomentMapSelectorGeometryAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
