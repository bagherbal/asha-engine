package generation2chargedleptonrootextensionbranchchambermonodromyaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ChargedLeptonRootExtensionBranchChamberMonodromyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 charged-lepton root-extension branch and chamber monodromy audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate600 charged-lepton root-extension branch/chamber monodromy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate599 trace-ring algebraic root-chamber boundary", Passed: a.Inherited.TraceRingDefined && a.Inherited.CharacteristicPolynomial && a.Inherited.AlgebraicExtension && !a.Inherited.HEOneFourthNative && !a.Inherited.BFlavNative, Detail: FormatInherited(a.Inherited)},
			{Name: "type cubic splitting field over the trace ring", Passed: a.Splitting.Typed && !a.Splitting.TraceRingOrdersRoots && a.Splitting.SplittingField != "", Detail: FormatSplitting(a.Splitting)},
			{Name: "audit discriminant and monodromy data", Passed: a.Monodromy.Discriminant != "" && !a.Monodromy.NativeBranchSelector && !a.Monodromy.NativeOrdering, Detail: FormatMonodromy(a.Monodromy)},
			{Name: "audit fourth-root branch structure", Passed: a.FourthRoot.ComplexSheetsPerEigenvalue == 4 && a.FourthRoot.PositiveRealBranchUnique && a.FourthRoot.RequiresPositivity && !a.FourthRoot.PositivityNative && !a.FourthRoot.FourthRootNative, Detail: FormatFourthRoot(a.FourthRoot)},
			{Name: "audit Koide chamber ordering and wall selection", Passed: !a.Chamber.TraceRingSelectsWall && !a.Chamber.DiscriminantSelectsWall && !a.Chamber.MonodromySelectsOrder && !a.Chamber.NativeChamberSelector && a.Chamber.Wall != "", Detail: FormatChamber(a.Chamber)},
			{Name: "define minimal charged-lepton root branch chamber seal", Passed: a.BranchSeal.AlgebraicOverTrace && !a.BranchSeal.Native && a.BranchSeal.Environmental && len(a.BranchSeal.Components) >= 5, Detail: FormatBranchSeal(a.BranchSeal)},
			{Name: "update B_flav as branch-anchored environmental balance", Passed: a.BFlav.ChargedLeptonSideTraceRing && a.BFlav.ChargedLeptonSideSplittingField && a.BFlav.ChargedLeptonSideFourthRootBranch && a.BFlav.ChargedLeptonSideChamberSeal && !a.BFlav.ChargedLeptonSideNative && a.BFlav.EnvironmentalOnly, Detail: FormatBFlav(a.BFlav)},
			{Name: "preserve root/chamber/firewall boundaries", Passed: !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.DerivesPMNSCKMNeutrino && !a.Firewalls.PromotesHEOneFourthNative && !a.Firewalls.PromotesChamberNative && !a.Firewalls.PromotesBFlavZero && !a.Firewalls.AddsCarrier && !a.Firewalls.AddsSelector && !a.Firewalls.SearchesNewConstants && a.Firewalls.PreservesGate352 && a.Firewalls.PreservesGate596 && a.Firewalls.PreservesGate599, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "compile final branch/chamber monodromy verdict", Passed: a.Final.SplittingFieldTyped && !a.Final.TraceRingOrdersSpectrum && !a.Final.NativeEigenvalueBranch && !a.Final.NativePositiveFourthRoot && !a.Final.NativeChamberSelector && a.Final.EpsilonBranchAlgebraic && !a.Final.BFlavNative, Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Final.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
