package generation2minimalrightmoduleweakdoubletsocketedgeoperatoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_MINIMAL_RIGHTMODULE_WEAKDOUBLET_SOCKET_EDGE_OPERATOR_AUDIT"
	theoremName = "Gate 847 — Minimal RightModule / WeakDoublet Socket Edge-Operator Audit"
)

func Generation2MinimalRightModuleWeakDoubletSocketEdgeOperatorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		checks := []theorem.Check{}
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks = append(checks,
			theorem.Check{Name: "inherit minimal right edge domain and Gate 846 response cells", Passed: a.Edges.DomainRank == HRMinRank && a.Impact.MinimalRightDomainInherited && a.Edges.ReconstructsGate846Cells && containsAll(a.Edges.Supports, []string{SupportResponseTableEdgeGenerator}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "audit weak doublet socket split as orientation seal", Passed: a.Weak.Complete && a.Weak.Orthogonal && a.Weak.RankHPlus == 1 && a.Weak.RankHMinus == 1 && a.Weak.WeakDim == 2 && a.Weak.OrientationSeal && !a.Weak.NativeSplit && !a.Weak.HiggsOrientationCertified && containsAll(a.Weak.Supports, []string{SupportWeakSocketPairSeal}) && containsAll(a.Weak.Failures, []string{FailureWeakSplitNotNative}), Detail: FormatWeak(a.Weak)},
			theorem.Check{Name: "construct three active symbolic socket edges", Passed: a.Edges.ActiveEdges == 3 && a.Edges.ExpectedActiveEdges == 3 && len(a.Edges.Edges) == 3 && allEdgesPresentNoMagnitude(a.Edges.Edges) && containsAll(a.Edges.Supports, []string{SupportThreeActiveSocketEdges, SupportRightDomainToWeakTargets}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "preserve puncture edge absence", Passed: a.Edges.MissingEdge.Puncture && !a.Edges.MissingEdge.Present && a.Edges.PunctureAbsent && containsAll(a.Edges.Supports, []string{SupportNeutralSingletonNullEdge}) && containsAll(a.Edges.Failures, []string{FailurePunctureNullEdgeOnlySeal, FailureNoNativeNullEdgeTheorem}), Detail: FormatEdge(a.Edges.MissingEdge)},
			theorem.Check{Name: "verify lepto-color preserving support", Passed: a.Edges.PreservesLeptoColor && edgesPreserveLeptoColor(a.Edges.Edges) && a.Edges.MissingEdge.LeptoColor == "P_1 -> P_1" && containsAll(a.Edges.Supports, []string{SupportLeptoColorPreservingEdges}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "classify edge operator as support seal, not D_F matrix", Passed: a.Edges.SupportOnly && !a.Edges.ExplicitDFMatrix && !a.Edges.NativeDFMatrix && !a.Edges.FirstOrderCertified && !a.Edges.BimoduleCommutantProof && containsAll(a.Edges.Failures, []string{FailureEdgeOperatorSealNotNative, FailureNoExplicitDFMatrix, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof}), Detail: FormatEdges(a.Edges)},
			theorem.Check{Name: "preserve magnitude, alpha, and particle firewalls", Passed: !a.Edges.Magnitudes && !a.Edges.AlphaDerived && !a.Edges.ParticleAssignment && a.Impact.AlphaStillSealed && a.Impact.MagnitudesStillMissing && containsAll(a.Edges.Failures, []string{FailureEdgeSupportNotYukawa, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoPhysicalParticleAssign, FailureNoRightNeutrinoTheorem, FailureNoThreeGenerationTheorem}), Detail: FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve official ledger freeze and no promotions", Passed: a.Ledger.OfficialFrozen && a.Ledger.R2PlusPlusPlusPlusEdgeStage && !a.Ledger.R3 && !a.Ledger.R4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			theorem.Check{Name: "preserve Gate 847 firewalls", Passed: a.Firewalls.Enforced && a.Firewalls.WeakSplitNotNative && a.Firewalls.EdgeOperatorSealNotNative && a.Firewalls.NoExplicitDFMatrix && a.Firewalls.NoFirstOrderProof && a.Firewalls.NoBimoduleCommutantProof && a.Firewalls.EdgeSupportNotYukawa && a.Firewalls.NoNumericalYukawaValues && a.Firewalls.AlphaStillSealed && a.Firewalls.NoTraceMagnitudeReadout && a.Firewalls.PunctureNullEdgeOnlySeal && a.Firewalls.NoNativeNullEdgeTheorem && a.Firewalls.NoPhysicalParticleAssignment && a.Firewalls.NoRightNeutrinoTheorem && a.Firewalls.NoThreeGenerationTheorem && a.Firewalls.NoNEffUpdate && a.Firewalls.NoCYukawaUpdate && a.Firewalls.NotR3 && a.Firewalls.NotR4 && a.Firewalls.Verdict == StatusFirewallGate847, Detail: FormatFirewalls(a.Firewalls)},
		)
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatWeak(a.Weak), FormatEdges(a.Edges), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func allEdgesPresentNoMagnitude(edges []Edge) bool {
	for _, e := range edges {
		if !e.Present || e.Puncture || e.HasMagnitude || e.DomainRank <= 0 || e.TargetRank <= 0 {
			return false
		}
	}
	return true
}

func edgesPreserveLeptoColor(edges []Edge) bool {
	for _, e := range edges {
		if e.LeptoColor != "P_3 -> P_3" && e.LeptoColor != "P_1 -> P_1" {
			return false
		}
	}
	return true
}
