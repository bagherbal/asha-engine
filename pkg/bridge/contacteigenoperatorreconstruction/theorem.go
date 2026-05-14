package contacteigenoperatorreconstruction

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactEigenoperatorInternalReconstructionQ4ContactOnlyTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Contact-eigenoperator internal reconstruction / q4 contact-only classification"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate406 audit", Passed: false, Detail: err.Error()}}}
		}
		internal := findRoute(a.Classification.Routes, "internal contact companion eigenoperator")
		rational := findRoute(a.Classification.Routes, "rational contact centralizer/idempotent route")
		resolvent := findRoute(a.Classification.Routes, "resolvent-field contact split")
		hphi := findRoute(a.Classification.Routes, "H_phi scalar identity selector")
		edge := findRoute(a.Classification.Routes, "one-form edge pullback / edge-weight selector")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits contact-internal and external-obstruction ledgers", Passed: a.Inheritance.Executed && a.Inheritance.Gate148Q4CandidateRows && a.Inheritance.Gate279CompanionConstructed && a.Inheritance.Gate279IrreducibleOverQ && a.Inheritance.Gate279NoNontrivialIdempotentQ && a.Inheritance.Gate398NoQuarticBundleFunctor && a.Inheritance.Gate399QuaternionicPolynomialNo && a.Inheritance.Gate400NoMixedEdgeQ4 && a.Inheritance.Gate401ChargeWeightsDisjoint && a.Inheritance.Gate402GraphNoQ4 && a.Inheritance.Gate403OrientationNoQ4 && a.Inheritance.Gate404QuotientNoQ4 && a.Inheritance.Gate405NoContactEdgePullback && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 reconstructs internally as contact companion/eigenoperator", Passed: a.ContactQ4.ReconstructedInternally && a.ContactQ4.Degree == Q4Degree && a.ContactQ4.Dimension == ContactPrimaryDim && a.ContactQ4.CharacteristicMatchesQ4 && a.ContactQ4.MinimalMatchesQ4 && a.ContactQ4.IrreducibleOverQ && a.ContactQ4.CompanionCyclic && !a.ContactQ4.UsesHphiBasis && !a.ContactQ4.UsesEdgeBasis && !a.ContactQ4.UsesObservedInput, Detail: FormatContactQ4(a.ContactQ4)},
			{Name: "contact centralizer is field with no native split", Passed: a.ContactAlgebra.CentralizerDimensionOverQ == 4 && a.ContactAlgebra.CentralizerIsField && a.ContactAlgebra.NontrivialIdempotentsOverQ == 0 && !a.ContactAlgebra.TwoByTwoBlockSplitOverQ && !a.ContactAlgebra.IndividualRootProjectorsOverQ && a.ContactAlgebra.ResolventIrreducibleOverQ && !a.ContactAlgebra.ResolventRootSelectedNatively && !a.ContactAlgebra.NativeRootSectorSemantics, Detail: FormatContactAlgebra(a.ContactAlgebra)},
			{Name: "internal route is promoted only inside contact sector", Passed: internal.Native && internal.ContactInternal && internal.PreservesQ4Internally && !internal.PromotableToScalarBundle && !internal.PromotableToYukawaTexture && internal.Verdict == StatusContactQ4Reconstructed, Detail: FormatRoute(internal)},
			{Name: "rational contact split fails over Q", Passed: rational.Native && rational.ContactInternal && !rational.PromotableToScalarBundle && rational.Verdict == StatusFailedNoNative2x2ContactSplit, Detail: FormatRoute(rational)},
			{Name: "resolvent split remains sealed", Passed: !resolvent.Native && resolvent.ContactInternal && resolvent.RequiresResolventAdjunction && resolvent.RequiresRootOrdering && resolvent.Verdict == StatusResolventObligationExplicit, Detail: FormatRoute(resolvent)},
			{Name: "H_phi selector route remains blocked", Passed: !hphi.Native && !hphi.HphiSelector && hphi.RequiresManualBasis && !hphi.PromotableToScalarBundle && hphi.Verdict == StatusFailedQ4NotHphiSelector, Detail: FormatRoute(hphi)},
			{Name: "edge pullback route remains blocked", Passed: !edge.Native && !edge.EdgeSelector && edge.RequiresManualBasis && !edge.PromotableToScalarBundle && edge.Verdict == StatusFailedNoContactEdgePullback, Detail: FormatRoute(edge)},
			{Name: "classification is contact-only under current functors", Passed: a.Classification.Executed && a.Classification.NativeInternalRoutes >= 1 && a.Classification.NativeHphiSelectorRoutes == 0 && a.Classification.NativeEdgePullbackRoutes == 0 && a.Classification.NativeYukawaReductionRoutes == 0 && a.Classification.SealedResolventRoutes >= 1 && a.Classification.ContactOnly && a.Classification.HphiIdentityStillOpen, Detail: FormatClassification(a.Classification)},
			{Name: "impact preserves scalar lane and flavor firewall", Passed: a.Impact.Q4InternalContactInvariant && !a.Impact.Q4ScalarBundleIdentifier && !a.Impact.Q4EdgeWeightOrPullback && !a.Impact.ContactProjectorOrSplitDerived && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.ScalarHphiLanePreserved && a.Impact.ContactLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and cross-sector firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoManualRootOrderingPromoted && a.Firewall.NoResolventRootPromoted && a.Firewall.NoArbitraryBasisPromoted && a.Firewall.NoCompanionOperatorCrossSector && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches Hphi-native algebra, not q4 forcing", Passed: a.Next.Gate == 407 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}
