package generation2spatialcp2selectorisotropyobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SpatialCP2SelectorAndSU3IsotropyObstructionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 spatial CP2 selector and SU(3) isotropy obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate573 spatial CP2 selector obstruction audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate572 CP3/B-L CP0|CP2 split and firewalls", Passed: a.Inherited.CP3ProjectiveLawSpace && a.Inherited.BMinusLProjectiveOnePlus3 && a.Inherited.SpatialCP2Block && a.Inherited.NoNativeSecondSelector && a.Inherited.NoCP3ToK7Functor && a.Inherited.NoPhysicalTimeOpened && a.Inherited.NoFlavorElectroweakData, Detail: FormatInherited(a.Inherited)},
			{Name: "certify CP2_sp as the B-L spatial projective block", Passed: a.Carrier.CertifiedAsSpatialBlock && a.Carrier.Gate572CriticalStratumMatched && a.Carrier.ComplexDimension == 2 && a.Carrier.RealDimension == 4, Detail: FormatCarrier(a.Carrier)},
			{Name: "verify U(3)/SU(3) symmetry and B-L scalar restriction", Passed: a.Symmetry.U3Dimension == 9 && a.Symmetry.SU3Dimension == 8 && a.Symmetry.BMinusLScalarOnWSpatial && !a.Symmetry.SuppliesFurtherSelector && !a.Symmetry.PreferredSpatialDirection, Detail: FormatSymmetry(a.Symmetry)},
			{Name: "verify SU(3) transitivity and point stabilizer on CP2", Passed: a.Transit.ActsTransitively && a.Transit.PointStabilizer == "S(U(1)xU(2))" && a.Transit.QuotientRealDimension == a.Transit.CP2RealDimension && !a.Transit.InvariantPointSelected && !a.Transit.InvariantRankOneProjector, Detail: FormatTransitivity(a.Transit)},
			{Name: "classify general Hermitian 2+1 selector and CP1|CP0 critical strata", Passed: a.Selector.ClassifiesCP1CP0Split && a.Selector.MultiplicityPattern == "2+1" && a.Selector.ProjectorIdempotentResidual < 1e-12 && a.Selector.ProjectorTrace == 1 && !a.Selector.NativeWithoutU, Detail: FormatSelector(a.Selector)},
			{Name: "search current ASHA data for native rank-one P_u", Passed: a.Search.CandidateCount >= 9 && !a.Search.NativeRankOneProjectorFound && !a.Search.NativeProjectivePointFound && !a.Search.NativeSecondSelectorFound, Detail: FormatSearch(a.Search)},
			{Name: "define minimal orientation seal and sealed commutant", Passed: a.Seal.SealName == "SpatialProjectiveOrientationSeal" && a.Seal.SealedNotNative && a.Seal.Commutant == "u(2)+u(1)" && a.Seal.CommutantDimension == 5 && a.Seal.CommMatchesGate555Formula, Detail: FormatSeal(a.Seal)},
			{Name: "mark U12 convention as basis-dependent sealed support only", Passed: a.WeakPlane.BasisDependent && !a.WeakPlane.NativeDerived && !a.WeakPlane.WeakIsospinIdentified && !a.WeakPlane.GenerationHierarchy && !a.WeakPlane.YukawaTextureDerived && !a.WeakPlane.CKMPMNSDerived, Detail: FormatWeakPlane(a.WeakPlane)},
			{Name: "preserve K7/product-time and flavor/electroweak firewalls", Passed: !a.Firewall.CP2ToK7FunctorOpened && !a.Firewall.ProductTimeOpened && !a.Firewall.PromotedToWeakIsospin && !a.Firewall.PromotedToPhysicalWeakPlane && !a.Firewall.GenerationHierarchyDerived && !a.Firewall.YukawaTextureDerived && !a.Firewall.CKMPMNSDerived && !a.Firewall.ObservedFlavorDataImported && !a.Firewall.PhysicalElectroweakDynamics && a.Firewall.Gate564565BoundaryPreserved, Detail: FormatFirewall(a.Firewall)},
			{Name: "return required A-G verdict", Passed: a.Final.SpatialCP2Certified && a.Final.SU3Transitive && !a.Final.SU3InvariantPointSelected && a.Final.GeneralTwoPlusOneSelector && !a.Final.NativeRankOnePU && a.Final.MinimalSeal != "" && !a.Final.PhysicalWeakFlavorEWDerived && !a.Final.K7OrProductTimeOpened, Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.MissingNextTheorem)
		for _, c := range a.Search.Candidates {
			notes = append(notes, FormatCandidate(c))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
