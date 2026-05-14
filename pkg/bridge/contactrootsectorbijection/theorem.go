package contactrootsectorbijection

import "github.com/bagherbal/asha-engine/pkg/theorem"

func QuarticRootToYukawaSectorBijectionContactProjectorSemanticsAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-ROOT-YUKAWA-SECTOR-BIJECTION-CONTACT-PROJECTOR-SEMANTICS-AUDIT"
	const name = "Quartic Root-to-Yukawa Sector Bijection / Contact Projector Semantics Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 278 root-sector bijection audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "quartic contact roots are retrieved and magnitude-audited", Passed: RootResidualOK(a) && a.Magnitude.AllRootsO1 && !a.Magnitude.AnyNativeNullRoot && a.Magnitude.OrderingAvailable && !a.Magnitude.OrderingInvariant && !a.Magnitude.MagnitudeBijectionDerived && !a.Magnitude.BGapScaleComparableToRoots, Detail: FormatMagnitude(a.Magnitude)},
			{Name: "Morita, B_gap, and tau_eta constraints reach sectors but not roots", Passed: a.Constraints.MoritaMultiplicityAvailable && a.Constraints.BGapTagAvailable && a.Constraints.TauEtaPairingAvailable && a.Constraints.ConstraintsReachSectors && !a.Constraints.ConstraintsReachRoots && !a.Constraints.UsesObservedMasses, Detail: FormatConstraints(a.Constraints)},
			{Name: "contact projector semantics are still missing over the base field", Passed: a.Projectors.QuarticIrreducibleOverQ && a.Projectors.ResolventIrreducibleOverQ && !a.Projectors.IndividualRootProjectorsOverQ && !a.Projectors.TwoPlusTwoPairProjectorsOverQ && a.Projectors.RequiresSplittingField && a.Projectors.RequiresResolventRootAdjunction && !a.Projectors.RationalContactProjectorDerived, Detail: FormatProjectors(a.Projectors)},
			{Name: "all root pairings remain compatible absent a root-sector map", Passed: a.Pairings.TotalCandidates == 3 && a.Pairings.CompatibleWithSectorSplit == 3 && a.Pairings.SelectedPairings == 0 && !a.Pairings.UniqueRootPairing && a.Pairings.SelectedRootPairing == "", Detail: FormatPairings(a.Pairings)},
			{Name: "root-to-sector bijection is not uniquely derived", Passed: a.Bijection.TotalRootSectorBijections == 24 && a.Bijection.BijectionsAfterSectorPairing > 1 && !a.Bijection.UniqueBijection && len(a.Bijection.DerivedAssignment) == 0, Detail: FormatBijection(a.Bijection)},
			{Name: "Gate-275 amplitude branch remains unlocked", Passed: a.BranchProjection.RPlus > 0 && a.BranchProjection.RMinus > 0 && !a.BranchProjection.ResolventRootSelected && !a.BranchProjection.RootPairingToRBranchMap && !a.BranchProjection.UniqueAmplitudeBranch && a.BranchProjection.SelectedBranch == "", Detail: FormatBranchProjection(a.BranchProjection)},
			{Name: "firewalls prevent empirical or aesthetic assignment", Passed: a.Firewall.NoObservedMassesUsed && a.Firewall.NoCKMPMNSUsed && a.Firewall.NoEmpiricalYukawaInserted && a.Firewall.NoRootOrderingPromotion && a.Firewall.NoArbitraryRootSectorMap && a.Firewall.NoBGapScaleToRootMap && a.Firewall.NoMultiplicityToAmplitude && a.Firewall.NoHiggsRatioClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future obligations are explicit", Passed: a.Future.NeedContactProjectorAction && a.Future.NeedRootSectorBijection && a.Future.NeedResolventToRBranchMap && a.Future.NeedPhysicalJHypercharge && a.Future.NeedHeatKernelNormalization && len(a.Future.Criteria) >= 5, Detail: FormatFuture(a.Future)},
			{Name: "summary records no-go without losing prior support", Passed: a.Summary.RootsRetrieved && a.Summary.ConstraintsAudited && !a.Summary.ProjectorSemanticsFound && !a.Summary.UniqueRootPairing && !a.Summary.UniqueRootSectorMap && !a.Summary.AmplitudeBranchLocked && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 278 audits root magnitudes, Morita 1+3 multiplicity, B_gap semantics, and resolvent pairings, but refuses to turn labels into roots without contact projector semantics.",
			"The sector-level {u,d}|{e,nu} result from Gate 277 remains supported; the root-level resolvent selection and Gate-275 r branch remain blocked.",
		}}
	}}
}
