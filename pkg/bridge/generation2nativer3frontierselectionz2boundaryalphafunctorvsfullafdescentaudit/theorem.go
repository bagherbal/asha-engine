package generation2nativer3frontierselectionz2boundaryalphafunctorvsfullafdescentaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NATIVE_R3_FRONTIER_SELECTION_Z2_BOUNDARY_ALPHA_FUNCTOR_VS_FULL_AF_DESCENT_AUDIT"
	theoremName = "Gate 911 — Native R3 Frontier Selection: Z2 BoundaryAlpha Functor vs Full A_F Descent Audit"
)

func Generation2NativeR3FrontierSelectionZ2BoundaryAlphaFunctorVsFullAFDescentAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 910 sealed R3 plateau inherited without reopening phase, representative-alpha, or socket-order wounds", Passed: a.Inherited.R3TraceLedgerSealed && !a.Inherited.NativeR3 && !a.Inherited.LoopBackPhase && !a.Inherited.LoopBackRepAlpha && !a.Inherited.LoopBackSocketOrder && containsAll(a.Inherited.Supports, []string{StatusGate910Inherited, StatusNoLoopBack}) && containsAll(a.Inherited.Failures, []string{FailureNotNativeR3, FailureAlphaStillSealed}), Detail: FormatInherited(a.Inherited)},
			{Name: "sealed trace diagnostics remain diagnostic-only and tied to BoundaryAlpha class seal", Passed: a.Trace.DiagnosticsOnly && !a.Trace.OfficialUpdates && near(a.Trace.Alpha, AlphaB) && near(a.Trace.OperatorNEff, OperatorNEffDiagnostic) && near(a.Trace.OperatorCYukawa, OperatorCYukawaDiagnostic) && near(a.Trace.OperatorCHiggs, OperatorCHiggsDiagnostic) && containsAll(a.Trace.Supports, []string{SupportFrontierADirectTraceLedger, SupportFrontierAControlsAlphaWeights}) && containsAll(a.Trace.Failures, []string{FailureAlphaStillSealed, FailureNoOfficialNEffUpdate}), Detail: FormatTrace(a.Trace)},
			{Name: "Frontier A is selected first because it directly controls alpha_B and trace/operator diagnostics", Passed: a.FrontierA.SelectedFirst && a.FrontierA.RequiredForNativeR3 && a.FrontierA.DirectlyControlsAlpha && a.FrontierA.DirectlyControlsTraceWeights && a.FrontierA.DirectlyControlsNEff && a.FrontierA.DirectlyControlsCYukawa && a.FrontierA.DirectlyControlsCHiggs && !a.FrontierA.Deferred && !a.FrontierA.R4OrLater && containsAll(a.FrontierA.Supports, []string{SupportFrontierAZ2AlphaPrimary, SupportFrontierADirectTraceLedger, SupportFrontierAControlsAlphaWeights, SupportNextBranchZ2Alpha}) && containsAll(a.FrontierA.Failures, []string{FailureFrontierANotCertified, FailureNoNativeZ2BoundaryAlphaFunctor, FailureAlphaSealedUntilFunctor}), Detail: FormatFrontier(a.FrontierA)},
			{Name: "Frontier B is required but deferred until the BoundaryAlpha source is sharpened", Passed: !a.FrontierB.SelectedFirst && a.FrontierB.RequiredForNativeR3 && a.FrontierB.FullAFDescentProblem && a.FrontierB.Deferred && !a.FrontierB.DirectlyControlsAlpha && !a.FrontierB.R4OrLater && containsAll(a.FrontierB.Supports, []string{SupportFrontierBRequired, SupportFrontierBSecond, SupportFrontierBAfterAlpha}) && containsAll(a.FrontierB.Failures, []string{FailureFullAFDescentStillBlocked, FailureAFOrientNotFullAF, FailureHiggsOrientationSealed}), Detail: FormatFrontier(a.FrontierB)},
			{Name: "Frontier C remains R4-or-later and must not be entered from Gate 911", Passed: !a.FrontierC.SelectedFirst && !a.FrontierC.RequiredForNativeR3 && a.FrontierC.GenerationFlavorYukawaBranch && a.FrontierC.R4OrLater && a.FrontierC.Deferred && containsAll(a.FrontierC.Supports, []string{SupportFrontierCR4Later}) && containsAll(a.FrontierC.Failures, []string{FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}), Detail: FormatFrontier(a.FrontierC)},
			{Name: "frontier selection chooses Z2 BoundaryAlpha functor before full A_F descent and blocks loopback/offical update", Passed: a.Selection.AttackAFirst && !a.Selection.AttackBBeforeA && !a.Selection.EnterGenerationNow && !a.Selection.NativeR3 && !a.Selection.LoopBackToPhase && !a.Selection.LoopBackToRepAlpha && !a.Selection.LoopBackToSocket && !a.Selection.UpdateOfficialLedger && containsAll(a.Selection.Supports, []string{SupportFrontierAZ2AlphaPrimary, SupportFrontierBSecond, SupportFrontierCR4Later, SupportNoLoopBack}) && containsAll(a.Selection.Failures, []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked}), Detail: FormatSelection(a.Selection)},
			{Name: "all native/source, full descent, generation/flavor, Yukawa, and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureNoNativeZ2BoundaryAlphaFunctor, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatTrace(a.Trace), FormatFrontier(a.FrontierA), FormatFrontier(a.FrontierB), FormatFrontier(a.FrontierC), FormatSelection(a.Selection), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
