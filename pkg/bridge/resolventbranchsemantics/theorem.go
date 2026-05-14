package resolventbranchsemantics

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ResolventBranchSemanticsProjectorSectorOrientationSealAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-RESOLVENT-BRANCH-SEMANTICS-PROJECTOR-SECTOR-ORIENTATION-SEAL-AUDIT"
	const name = "Resolvent Branch Semantics / Projector-to-Sector Orientation Seal Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 281 resolvent branch semantics audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "trace/norm semantic audit covers all three resolvent branches", Passed: a.TraceNorm.BranchCount == 3 && a.TraceNorm.PossibleOrientations == 6 && len(a.TraceNorm.Branches) == 3, Detail: FormatTraceNorm(a.TraceNorm)},
			{Name: "Morita 1|3 multiplicity does not align with 2|2 contact projectors", Passed: !a.TraceNorm.AnyNativeOrientationPreferred && allBranchesRankTwo(a.TraceNorm), Detail: FormatTraceNorm(a.TraceNorm)},
			{Name: "ProjectorSectorOrientationSeal is conditional, not native", Passed: a.OrientationSeal.Active && a.OrientationSeal.OrientationIsSealedConditional && !a.OrientationSeal.OrientationIsNativeTheorem && a.OrientationSeal.GrantsProjectorSectorMap && !a.OrientationSeal.GrantsAmplitudeBranchMap, Detail: FormatSeal(a.OrientationSeal)},
			{Name: "sealed orientation does not derive r_+/r_- branch", Passed: a.RBranch.OrientationLocked && a.RBranch.ResolventBranchSelected && a.RBranch.ProjectorSectorMapAvailable && !a.RBranch.AlgebraicResolventToRMapDerived && !a.RBranch.UniqueAmplitudeBranch && a.RBranch.SelectedRBranch == "", Detail: FormatRBranch(a.RBranch)},
			{Name: "Seeley-de Witt preparation obligations remain explicit", Passed: !a.SeeleyPrep.HiggsRatioReady && !a.SeeleyPrep.RBranchLocked && !a.SeeleyPrep.PhysicalJDerived && !a.SeeleyPrep.HeatKernelProjectionDerived && len(a.SeeleyPrep.Criteria) >= 5, Detail: FormatSeeleyPrep(a.SeeleyPrep)},
			{Name: "firewalls preserve sealed-orientation status", Passed: a.Firewall.NoNativeOrientationOverclaim && a.Firewall.NoMultiplicityToOrientationOverclaim && a.Firewall.NoBasisDependentNormPromotion && a.Firewall.OrientationSealDoesNotRewriteNativeStatus && a.Firewall.NoRBranchOverclaim && a.Firewall.NoHiggsRatioClaimed && a.Firewall.NoObservedMassesUsed && a.Firewall.NoEmpiricalYukawaInserted && !a.Firewall.FiniteCorePolluted, Detail: FormatFirewall(a.Firewall)},
			{Name: "summary answers the multiplicity question without overclaim", Passed: a.Summary.TraceNormAuditComplete && !a.Summary.NativeOrientationPreferred && a.Summary.OrientationSealActivated && a.Summary.RepresentativeOrientationAssigned && a.Summary.ProjectorSectorMapConditional && !a.Summary.AmplitudeBranchLocked && !a.Summary.HiggsRatioDerived && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"The answer to the Gate-281 question is no: the 1⊕3 Morita trace multiplicities live in the finite Hilbert bimodule and do not natively orient rank-2 contact projectors.",
			"A ProjectorSectorOrientationSeal can quarantine a representative physical orientation for downstream stress tests, but the r_+/r_- branch and Higgs-ratio path remain bridge-required.",
		}}
	}}
}

func allBranchesRankTwo(t TraceNormSemanticAudit) bool {
	for _, b := range t.Branches {
		if !b.ProjectorTracesEqual || !b.ProjectorRanksEqual {
			return false
		}
		if b.ProjectorA.AlignsWithOnePlusThree || b.ProjectorB.AlignsWithOnePlusThree {
			return false
		}
		if b.ProjectorA.RankApprox < 1.999999 || b.ProjectorA.RankApprox > 2.000001 || b.ProjectorB.RankApprox < 1.999999 || b.ProjectorB.RankApprox > 2.000001 {
			return false
		}
	}
	return true
}
