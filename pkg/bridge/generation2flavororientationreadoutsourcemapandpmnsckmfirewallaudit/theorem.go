package generation2flavororientationreadoutsourcemapandpmnsckmfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FlavorOrientationReadoutSourceMapAndPMNSCKMFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 788 — Flavor Orientation Readout Source and PMNS-CKM Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate788}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate787 flavor-boundary readout factorization", Passed: a.Gate787.Inherited && strings.Contains(a.Gate787.CompositeSeal, "FlavorBoundaryReadoutSeal") && a.Gate787.KappaOrientIsFocus, Detail: a.Gate787.CompositeSeal},
			{Name: "decompose kappa_orient into PMNS reactor and CKM orientation terms", Passed: a.Decomposition.Recorded && a.Decomposition.Formula == "sin^2(theta13)/4 - J_CKM" && strings.Contains(a.Decomposition.PMNSTerm, "PMNS") && strings.Contains(a.Decomposition.CKMTerm, "J_CKM") && strings.Contains(a.Decomposition.QuarterFactor, "1/4") && a.Decomposition.ShapeTyped && !a.Decomposition.Native, Detail: a.Decomposition.Formula},
			{Name: "audit PMNS reactor leakage term", Passed: a.PMNS.Audited && containsAll(a.PMNS.Theta13SourceCandidates, []string{"generation carrier theorem", "Yukawa operator theorem", "PMNS mixing theorem"}) && !a.PMNS.Theta13Native && a.PMNS.QuarterResonance && containsAll(a.PMNS.QuarterResonanceSources, []string{"rho_plus=I_K7+/4", "Tr(rho_plus P_rad)=1/4"}) && !a.PMNS.TypedMapFromK7Quarter && a.PMNS.RemainsFlavorSealInput, Detail: a.PMNS.Term},
			{Name: "audit CKM Jarlskog orientation term", Passed: a.CKM.Audited && strings.Contains(a.CKM.SourceType, "orientation area") && containsAll(a.CKM.SourceCandidates, []string{"Yukawa operator theorem", "CKM mixing theorem", "generation orientation theorem"}) && !a.CKM.JCKMNative && a.CKM.NegativeSignCandidate && !a.CKM.NativeSignTheorem && a.CKM.RemainsFlavorSeal, Detail: a.CKM.Term},
			{Name: "audit boundary-only replacement", Passed: a.BoundaryOnly.Audited && closeRel(a.BoundaryOnly.KappaOrient, kappaOrientSnapshot, 2e-15) && closeRel(a.BoundaryOnly.KappaBoundary, kappaBoundarySnapshot, 2e-15) && a.BoundaryOnly.AbsRatioBoundaryToOrient < 1e-3 && a.BoundaryOnly.BoundaryPartSmallCorrection && !a.BoundaryOnly.BoundaryReplacesOrient, Detail: FormatBoundaryOnly(a.BoundaryOnly)},
			{Name: "audit existing ASHA geometry source candidates", Passed: a.Geometry.Audited && len(a.Geometry.Candidates) >= 6 && !a.Geometry.K7HodgeDerivesPMNSCKM && !a.Geometry.HiggsRadialDerivesTheta13 && !a.Geometry.BoundaryPairDerivesFlavorMixing && !a.Geometry.NEffDerivesMixingAngles && !a.Geometry.GenerationMixingOperatorFound, Detail: a.Geometry.Verdict},
			{Name: "refine FlavorBoundaryReadoutSeal", Passed: a.SealRefinement.Recorded && a.SealRefinement.OriginalSeal == "FlavorBoundaryReadoutSeal" && containsAll(a.SealRefinement.RefinedSeals, []string{"FlavorOrientationReadoutSeal", "BoundaryGaugeCorrectionSeal"}) && closeRel(a.SealRefinement.KappaOrient, kappaOrientSnapshot, 2e-15) && closeRel(a.SealRefinement.KappaBoundary, kappaBoundarySnapshot, 2e-15) && a.SealRefinement.BoundaryGaugeStronglyTyped && a.SealRefinement.OrientationTrueObstruction && !a.SealRefinement.OrientationSealNative, Detail: FormatSealRefinement(a.SealRefinement)},
			{Name: "audit kappa_orient runtime target absence", Passed: a.Runtime.Audited && containsAll(a.Runtime.ForbiddenDirectVariables, []string{"lambda_runtime", "m_H_tree", "C_Higgs", "G_F", "v"}) && !a.Runtime.ContainsForbidden && a.Runtime.FormulaLevelIndependent && !a.Runtime.TheoremLevelIndependent, Detail: a.Runtime.Verdict},
			{Name: "record status propagation", Passed: a.Propagation.Recorded && strings.Contains(a.Propagation.KappaOrient, "FlavorOrientationReadoutSeal") && strings.Contains(a.Propagation.KappaERed, "mixed flavor-boundary") && strings.Contains(a.Propagation.FWall3, "Level B+") && strings.Contains(a.Propagation.CHistory, "Level B") && strings.Contains(a.Propagation.CHiggs, "not Level C"), Detail: a.Propagation.Verdict},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.KappaOrientNativeFlavorTheorem && !a.Firewalls.Theta13DerivedPMNSTheorem && !a.Firewalls.JCKMDerivedCKMTheorem && !a.Firewalls.QuarterResonanceProof && !a.Firewalls.K7QuarterTheta13SourceTheorem && !a.Firewalls.NEffMixingAngleTheorem && !a.Firewalls.BoundaryPairFlavorMixingTheorem && !a.Firewalls.KappaBoundaryFullKappaETheorem && !a.Firewalls.FWallNativeBoundaryTheorem && !a.Firewalls.CHistoryFullPrediction && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate788, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth, FormatKappaSplit(a), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
