package generation2higgsorientationsourcecandidateweaksocketselectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_HIGGS_ORIENTATION_SOURCE_CANDIDATE_WEAK_SOCKET_SELECTOR_AUDIT"
	theoremName = "Gate 892 — HiggsOrientation Source Candidate and WeakSocket Selector Audit"
)

func Generation2HiggsOrientationSourceCandidateWeakSocketSelectorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 891 descent obstruction inherited", Passed: a.WeakSelector.FullHMixesWeakSockets && !a.WeakSelector.NativeOrientationSource && containsAll(a.WeakSelector.Failures, []string{FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient}), Detail: FormatWeakSelector(a.WeakSelector)},
			{Name: "finite one-form and puncture/kernel pair are strongest candidates", Passed: !a.Candidates.AnyNativeSourceCertified && a.Candidates.RequiresOrientationSeal && containsAll(a.Candidates.Supports, []string{SupportFiniteOneFormStrongestCandidate, SupportPunctureKernelPointsToHPlus}) && containsAll(a.Candidates.Failures, []string{FailureNoNativeOneFormOrientationTheorem, FailurePunctureKernelNotNativeOrientation}), Detail: FormatCandidateAudit(a.Candidates)},
			{Name: "D_F support compatibility is circular if used as orientation source", Passed: a.EdgeOrientation.AssumesWeakFrame && !a.EdgeOrientation.DerivesWeakFrame && a.EdgeOrientation.CircularIfUsedAsSource && containsAll(a.EdgeOrientation.Failures, []string{FailureDFSupportRestatesOrientation}), Detail: FormatEdgeOrientation(a.EdgeOrientation)},
			{Name: "puncture/kernel pair points to h_+ but is not a native theorem", Passed: a.PunctureKernel.PointsToHPlus && !a.PunctureKernel.NativeSourceCertified && containsAll(a.PunctureKernel.Failures, []string{FailurePunctureKernelNotNativeOrientation}), Detail: FormatPunctureKernel(a.PunctureKernel)},
			{Name: "B-L imbalance compatible but not weak-frame selector", Passed: a.BMinusL.Compatible && !a.BMinusL.SelectsWeakFrame && a.BMinusL.ActiveTrace == 1 && a.BMinusL.PunctureTrace == -1 && a.BMinusL.FullRectangleTrace == 0 && containsAll(a.BMinusL.Failures, []string{FailureBLImbalanceDoesNotSelectWeakFrame}), Detail: FormatBMinusL(a.BMinusL)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4 physical generation flavor and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureNoNativeHiggsOrientationSource, FailureNoNativeOneFormOrientationTheorem, FailureNoNativeR3SectorLedger}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatWeakSelector(a.WeakSelector), FormatCandidateAudit(a.Candidates), FormatEdgeOrientation(a.EdgeOrientation), FormatPunctureKernel(a.PunctureKernel), FormatBMinusL(a.BMinusL), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeHiggsOrientationSource, FailureNoNativeOneFormOrientationTheorem, FailurePostOrientationNotFullAF, FailureNoNativeFiniteSectorProjectorTheorem, FailureNoNativeR3SectorLedger, FailureNoNativeIncidenceFunctor, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}
