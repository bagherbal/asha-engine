package generation2spatialprojectiveorientationsealminimalityconsequenceaudit

import "testing"

func TestGate574SpatialProjectiveOrientationSealMinimalityAndConsequenceAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.CP2SpatialBlockCertified || !a.Inherited.SU3ActsTransitively || !a.Inherited.NoSU3InvariantPoint || !a.Inherited.NoNativeRankOneProjector {
		t.Fatalf("bad inherited obstruction: %s", FormatInherited(a.Inherited))
	}
	if a.Seal.SealName != "SpatialProjectiveOrientationSeal" || a.Seal.Rank != 1 || a.Seal.Trace != 1 || a.Seal.IdempotentResidual > 1e-12 || !a.Seal.Hermitian || a.Seal.NativeDerived {
		t.Fatalf("bad seal: %s", FormatSealDefinition(a.Seal))
	}
	if !a.Selector.ConstructsCP1CP0Split || a.Selector.MultiplicityPattern != "2+1" || a.Selector.NativeWithoutSeal {
		t.Fatalf("bad selector: %s", FormatSealedSelector(a.Selector))
	}
	if a.Commutant.Commutant != "u(2)+u(1)" || a.Commutant.Dimension != 5 || !a.Commutant.MatchesGate555Formula || !a.Commutant.SealedSupportOnly {
		t.Fatalf("bad commutant: %s", FormatCommutant(a.Commutant))
	}
	if !a.Basis.BasisDependent || a.Basis.NativeBasisSelection || a.Basis.ConventionalPlaneName != "U_12" {
		t.Fatalf("bad basis example: %s", FormatBasis(a.Basis))
	}
	if a.WeakPlane.ComplementaryCP1CanBeCalledPhysicalWeakPlane || a.WeakPlane.CompatibilityProven || !a.WeakPlane.RequiresFirstOrderCompatibility {
		t.Fatalf("bad weak-plane firewall: %s", FormatWeakPlaneFirewall(a.WeakPlane))
	}
	if a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.PhysicalEWDynamicsDerived {
		t.Fatalf("bad flavor/electroweak firewall: %s", FormatFlavorGenerationFirewall(a.FlavorEW))
	}
	if !a.Boundaries.TauEtaTraceShadowOnly || !a.Boundaries.Q4ContactOnly || !a.Boundaries.PauliQuaternionicSocketOnly || !a.Boundaries.Gate564565BridgeSymbolic || !a.Boundaries.K7TimeRoutesSealed {
		t.Fatalf("bad prior boundaries: %s", FormatPreviousGateBoundaries(a.Boundaries))
	}
	if !a.Minimality.SealIsMinimal || !a.Minimality.PointProjectorEquivalence || !a.Minimality.Any2Plus1SelectorDeterminesPU {
		t.Fatalf("bad minimality: %s", FormatMinimality(a.Minimality))
	}
	if !a.Final.SealSufficient || !a.Final.SealMinimal || !a.Final.ReducesSymmetryToU2U1 || a.Final.DerivesPhysicalWeakFlavorElectroweakData || a.Final.K7OrProductTimeOpened {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate574RankOneProjectorUtility(t *testing.T) {
	p := rankOneProjectorReal([]float64{2, 0, 0})
	if projectorIdempotentResidual(p) > 1e-12 {
		t.Fatalf("projector not idempotent: %v", p)
	}
	if trace3(p) != 1 {
		t.Fatalf("expected rank-one trace 1, got %.12g", trace3(p))
	}
	if !symmetric3(p) {
		t.Fatalf("expected Hermitian/symmetric representative projector: %v", p)
	}
}

func TestGate574Theorem(t *testing.T) {
	res := Generation2SpatialProjectiveOrientationSealMinimalityAndConsequenceAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
