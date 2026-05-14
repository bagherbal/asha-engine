package resolventfieldadjunction

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ResolventFieldAdjunctionContactProjectorConstructionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-RESOLVENT-FIELD-ADJUNCTION-CONTACT-PROJECTOR-CONSTRUCTION-AUDIT"
	const name = "Resolvent Field Adjunction / Contact Projector Construction Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 280 resolvent adjunction audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "ResolventAdjunctionSeal is active but not a native branch selector", Passed: a.Seal.Active && a.Seal.GrantsConditionalProjectors && !a.Seal.GrantsNativeBranchSelection, Detail: FormatSeal(a.Seal)},
			{Name: "three conditional resolvent branches are constructed", Passed: a.BranchSpace.BranchCount == 3 && a.BranchSpace.ConditionalBranchCount == 3 && a.BranchSpace.AllBranchesProjectorsValid, Detail: FormatBranchSpace(a.BranchSpace)},
			{Name: "each conditional branch has idempotent commuting projectors", Passed: allProjectorResidualsOK(a.BranchSpace), Detail: FormatBranchSpace(a.BranchSpace)},
			{Name: "Gate-277 sector pairing is not overpromoted to projector-sector bijection", Passed: a.SectorBijection.SectorPairingSelected && a.SectorBijection.ConditionalProjectorsExist && a.SectorBijection.RequiresMappingProjectorToSector && !a.SectorBijection.MappingDerived && !a.SectorBijection.UsesNumericalRootOrdering, Detail: FormatSector(a.SectorBijection)},
			{Name: "Gate-275 r branch remains unmapped", Passed: !a.RBranch.ResolventToRMapDerived && !a.RBranch.UniqueAmplitudeBranch && a.RBranch.SelectedBranch == "", Detail: FormatRBranch(a.RBranch)},
			{Name: "firewalls preserve conditional-adjunction status", Passed: a.Firewall.NoArbitraryResolventRootPromoted && a.Firewall.NoNumericalOrderingPromotion && a.Firewall.NoEmpiricalYukawaInserted && a.Firewall.NoObservedMassesUsed && a.Firewall.ConditionalAdjunctionNotNativeTheorem && a.Firewall.NoProjectorSectorOverclaim && a.Firewall.NoHiggsRatioClaimed && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "future obligations are explicit", Passed: a.Future.NeedNativeResolventSelector && a.Future.NeedProjectorSectorSemantics && a.Future.NeedResolventToRBranchMap && a.Future.NeedHeatKernelProjection && len(a.Future.Criteria) >= 4, Detail: FormatFuture(a.Future)},
			{Name: "summary records projectors without branch lock", Passed: a.Summary.SealActivated && a.Summary.ConditionalProjectorsConstructed && a.Summary.AllBranchProjectorsValid && !a.Summary.NativeResolventRootSelected && !a.Summary.SectorBijectionDerived && !a.Summary.AmplitudeBranchLocked && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 280 shows the mathematical effect of adjoining a resolvent root: 2+2 projectors become constructible on each sealed branch.",
			"The seal does not select a branch, assign projectors to {u,d}|{e,nu}, or map a branch to r_+/r_-; those remain future obligations.",
		}}
	}}
}

func allProjectorResidualsOK(bs BranchSpaceAudit) bool {
	if bs.BranchCount != 3 || !bs.AllBranchesProjectorsValid {
		return false
	}
	for _, b := range bs.Branches {
		if !b.ProjectorsValid || b.ProjectorA.IdempotentResidual > 1e-8 || b.ProjectorB.IdempotentResidual > 1e-8 || b.ProjectorA.CommutesWithCompanionResidual > 1e-8 || b.ProjectorB.CommutesWithCompanionResidual > 1e-8 || b.SumToIdentityResidual > 1e-8 || b.OrthogonalityResidual > 1e-8 || b.ResolventResidualAbs > 1e-6 || b.FactorizationResidualAbs > 1e-10 {
			return false
		}
	}
	return true
}
