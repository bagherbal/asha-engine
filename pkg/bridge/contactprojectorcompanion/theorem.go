package contactprojectorcompanion

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactProjectorActionQuarticCompanionModuleSemanticsAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-PROJECTOR-ACTION-QUARTIC-COMPANION-MODULE-SEMANTICS-AUDIT"
	const name = "Contact Projector Action / Quartic Companion Module Semantics Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 279 companion projector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "quartic companion matrix is constructed over Q", Passed: a.Companion.CharacteristicMatches && CompanionTraceOK(a), Detail: FormatCompanion(a.Companion)},
			{Name: "quartic and resolvent irreducibility are certified by modular witnesses", Passed: a.Irreducibility.QuarticPrimitiveOverZ && a.Irreducibility.QuarticIrreducibleOverModP && a.Irreducibility.QuarticIrreducibleOverQ && a.Irreducibility.ResolventIrreducibleOverModP && a.Irreducibility.ResolventIrreducibleOverQ, Detail: FormatIrreducibility(a.Irreducibility)},
			{Name: "centralizer is the quartic field and has no nontrivial rational idempotents", Passed: a.Centralizer.CompanionCyclic && a.Centralizer.CentralizerDimensionOverQ == 4 && a.Centralizer.CentralizerIsField && a.Centralizer.NontrivialIdempotentsOverQ == 0 && !a.Centralizer.IndividualRootProjectorsOverQ && !a.Centralizer.TwoPlusTwoProjectorsOverQ && !a.Centralizer.BlockDiagonalizes2x2OverQ, Detail: FormatCentralizer(a.Centralizer)},
			{Name: "native finite-geometry action candidates do not yield contact pair projectors", Passed: !a.NativeActions.AnyLegalAction && !a.NativeActions.AnyCommutingProjector && !a.NativeActions.AnyPairSelector && len(a.NativeActions.Candidates) >= 3, Detail: FormatNativeActions(a.NativeActions)},
			{Name: "2+2 pair projector requires a resolvent-root adjunction", Passed: a.Resolvent.PairProjectorRequiresResolventRoot && !a.Resolvent.ResolventRootAlreadySelected && a.Resolvent.AdjoiningResolventRootWouldSplit && !a.Resolvent.NativeAdjunctionDerived && a.Resolvent.BranchesAfterAdjunction == 3, Detail: FormatResolvent(a.Resolvent)},
			{Name: "Gate-277 sector pairing does not imply contact root or r-branch selection", Passed: a.Branch.SectorPairingSelected && a.Branch.SectorPairing == "{u,d}|{e,nu}" && !a.Branch.CompanionProjectorDerived && !a.Branch.ContactResolventRootSelected && !a.Branch.RootSectorBijectionDerived && !a.Branch.RBranchMapDerived && a.Branch.SelectedBranch == "", Detail: FormatBranch(a.Branch)},
			{Name: "firewalls reject numerical-ordering and arbitrary-resolvent promotion", Passed: a.Firewall.NoNumericalOrderingPromotion && a.Firewall.NoObservedMassesUsed && a.Firewall.NoEmpiricalYukawaInserted && a.Firewall.NoArbitraryResolventRoot && a.Firewall.NoAestheticRootPairing && a.Firewall.NoBGapToRootMagnitudeMap && a.Firewall.NoHiggsRatioClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future obligations are explicit", Passed: a.Future.NeedNativeOperatorOnContactModule && a.Future.NeedNontrivialIdempotent && a.Future.NeedResolventRootSelector && a.Future.NeedRootSectorBijection && a.Future.NeedRBranchMap && len(a.Future.Criteria) >= 5, Detail: FormatFuture(a.Future)},
			{Name: "summary records companion-module no-go without losing Gate-277 support", Passed: a.Summary.CompanionConstructed && a.Summary.IrreducibilityCertified && !a.Summary.NativeProjectorFound && !a.Summary.BlockDiagonalizedOverQ && !a.Summary.ResolventRootSelected && !a.Summary.RootSectorBijection && !a.Summary.AmplitudeBranchLocked && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 279 upgrades the Gate-278 projector firewall from root magnitudes to a companion-module theorem: over Q the contact module is irreducible and admits only trivial commuting idempotents.",
			"The topological {u,d}|{e,nu} sector split remains supported, but a contact resolvent root and the Gate-275 r branch remain unselected.",
		}}
	}}
}
