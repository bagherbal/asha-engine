package resolventcubictagselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ResolventCubicSelectorBGapTauEtaSymmetryBreakingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-RESOLVENT-CUBIC-SELECTOR-BGAP-TAUETA-SYMMETRY-BREAKING-AUDIT"
	const name = "Resolvent Cubic Selector / B-Gap and Tau-Eta Symmetry Breaking Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 277 resolvent selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "resolvent cubic and quartic roots are retrieved", Passed: a.Resolvent.Retrieved && a.Resolvent.EncodesTwoPlusTwo && a.Resolvent.IrreducibleOverQ && len(a.Resolvent.QuarticRoots) == 4 && len(a.Resolvent.Branches) == 3 && !a.Resolvent.CanonicalRootPreviouslySelected, Detail: FormatResolvent(a.Resolvent)},
			{Name: "tau_eta and B_gap tags are applied only at sector level", Passed: a.Tags.TauEtaBindsUD && a.Tags.BGapTagsNeutrino && a.Tags.TagsReachSectorLabels && !a.Tags.TagsReachQuarticRoots && !a.Tags.UsesObservedMasses, Detail: FormatTags(a.Tags)},
			{Name: "Galois sieve selects unique sector pairing", Passed: a.Sieve.TotalCandidates == 3 && a.Sieve.SurvivingSectorPairings == 1 && a.Sieve.SelectedSectorPairing == "{u,d}|{e,nu}" && a.Sieve.UniqueSectorPairing && !a.Sieve.UniqueContactRoot, Detail: FormatSieve(a.Sieve)},
			{Name: "Gate 275 r branches are inherited but not selected", Passed: len(a.BranchProjection.BranchesInherited) == 2 && BranchResidualOK(a.BranchProjection.BranchesInherited) && !a.BranchProjection.ResolventRootSelected && !a.BranchProjection.ResolventRootToRBranchMap && !a.BranchProjection.UniqueRBranchSelected, Detail: FormatProjection(a.BranchProjection)},
			{Name: "firewalls prevent overpromotion", Passed: a.Firewall.NoObservedMassesUsed && a.Firewall.NoCKMPMNSUsed && a.Firewall.NoEmpiricalYukawaInserted && a.Firewall.NoArbitraryRootSectorMap && a.Firewall.SectorPairingNotOverpromoted && a.Firewall.ContactRootNotOverpromoted && a.Firewall.RBranchNotOverpromoted && a.Firewall.HiggsRatioNotClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future obligations require root-sector and branch maps", Passed: a.Future.NeedRootSectorMap && a.Future.NeedContactProjectors && a.Future.NeedResolventToRMap && a.Future.NeedBranchSelector && a.Future.NeedHeatKernelMap && len(a.Future.Criteria) >= 5, Detail: FormatFuture(a.Future)},
			{Name: "summary records support plus no-go", Passed: a.Summary.ResolventRetrieved && a.Summary.TagsApplied && a.Summary.UniqueSectorPairing && !a.Summary.UniqueContactRoot && !a.Summary.UniqueAmplitudeBranch && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 277 supports the tau_eta+B_gap sector-pairing selector but refuses to identify it with a contact resolvent root without a quartic-root/Yukawa-sector theorem.",
			"The Gate-275 r_+/r_- amplitude branch and the Seeley-de Witt Higgs ratio remain unselected.",
		}}
	}}
}
