package generation2phaseanchoredneutralpunctureairlockfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_PHASE_ANCHORED_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR_AUDIT"
	theoremName = "Gate 901 — PhaseAnchored NeutralPuncture Airlock Functor Audit"
)

func Generation2PhaseAnchoredNeutralPunctureAirlockFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "phase anchor orders right-character pair only as bridge object", Passed: a.PhaseAnchor.OrdersPair && !a.PhaseAnchor.SelectsNatively && containsAll(a.PhaseAnchor.Supports, []string{SupportPhaseAnchorOrdersRightPair, SupportPhaseAnchorSelectsEPlus}) && containsAll(a.PhaseAnchor.Failures, []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambda}), Detail: FormatPhaseAnchor(a.PhaseAnchor)},
			{Name: "socket order and puncture are selected by phase anchor, not natively", Passed: a.SocketOrder.OrderedByPhaseAnchor && !a.SocketOrder.OrderedNatively && a.SocketOrder.Puncture == PPhase && containsAll(a.SocketOrder.Supports, []string{SupportPhaseAnchoredPuncture, SupportSocketOrderObstructionCollapsed}), Detail: FormatSocketOrder(a.SocketOrder)},
			{Name: "ordered edge table collapses to phase-anchored airlock rule", Passed: a.EdgeOrdering.GeneratedByPhaseAnchor && !a.EdgeOrdering.GeneratedNatively && a.EdgeOrdering.MissingEdge == "e_lambda tensor P_1 -> h_lambda tensor P_1 = 0" && containsAll(a.EdgeOrdering.Supports, []string{SupportPhaseAnchoredEdgeTable}), Detail: FormatEdgeOrdering(a.EdgeOrdering)},
			{Name: "weak socket frame is reconstructed as phase-indexed kernel complement under seal", Passed: a.WeakKernel.PhaseIndexedFrame && !a.WeakKernel.SelectorNative && a.WeakKernel.Kernel == "h_lambda tensor P_1" && containsAll(a.WeakKernel.Supports, []string{SupportPhaseAnchoredLeftKernel, SupportHiggsOrientationCollapsed}), Detail: FormatWeakKernel(a.WeakKernel)},
			{Name: "boundary alpha targets collapse to phase-anchored puncture flag", Passed: a.BoundaryAlpha.SelectedByPhasePuncture && !a.BoundaryAlpha.NativeAlphaFunctor && a.BoundaryAlpha.RankF1OverF0 == AlphaRankOneNumerator && a.BoundaryAlpha.RankF2OverF0 == AlphaRankTwoNumerator && near(a.BoundaryAlpha.Alpha, AlphaB) && containsAll(a.BoundaryAlpha.Supports, []string{SupportPhaseAnchoredAlphaTargets, SupportBoundaryIncidenceCollapsed}), Detail: FormatBoundaryAlpha(a.BoundaryAlpha)},
			{Name: "oriented stabilizer is phase-anchored layer, not full A_F descent", Passed: a.Stabilizer.PhaseAnchored && !a.Stabilizer.FullDescent && containsAll(a.Stabilizer.Supports, []string{SupportAFOrientPhaseAnchored}) && containsAll(a.Stabilizer.Failures, []string{FailureAFOrientNotFullAF, FailureNoNativeDescentFullAF}), Detail: FormatStabilizer(a.Stabilizer)},
			{Name: "multiple blockers reduce to one master phase-anchored airlock functor", Passed: a.MasterFunctor.UnifiesSocketOrder && a.MasterFunctor.UnifiesEdgeOrdering && a.MasterFunctor.UnifiesWeakKernel && a.MasterFunctor.UnifiesBoundaryAlpha && a.MasterFunctor.SingleMasterBlocker && !a.MasterFunctor.NativeFunctor && containsAll(a.MasterFunctor.Supports, []string{SupportAirlockUnifiesAlphaAndHiggs, SupportR3SealedReducesToSingleFunctor}), Detail: FormatMasterFunctor(a.MasterFunctor)},
			{Name: "operator diagnostics remain coherent and official ledgers frozen", Passed: a.Diagnostics.Coherent && a.Diagnostics.OfficialFrozen && !a.Diagnostics.CanUpdate && near(a.Diagnostics.Alpha, AlphaB) && near(a.Diagnostics.NEff, OperatorNEffDiagnostic) && !near(a.Diagnostics.NEff, a.Diagnostics.OfficialNEff), Detail: FormatDiagnostics(a.Diagnostics)},
			{Name: "native R3/R4, phase, alpha, Higgs orientation, descent, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativePhaseAnchoredAirlock, FailureNoNativeRightPhaseOrientation, FailureAlphaStillSealedWithoutPhaseFunctor, FailureHiggsStillSealedWithoutPhaseWeak, FailureNotNativeR3, FailureNoR4NativeYukawaTheorem}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatPhaseAnchor(a.PhaseAnchor), FormatSocketOrder(a.SocketOrder), FormatEdgeOrdering(a.EdgeOrdering), FormatWeakKernel(a.WeakKernel), FormatBoundaryAlpha(a.BoundaryAlpha), FormatStabilizer(a.Stabilizer), FormatMasterFunctor(a.MasterFunctor), FormatDiagnostics(a.Diagnostics), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
