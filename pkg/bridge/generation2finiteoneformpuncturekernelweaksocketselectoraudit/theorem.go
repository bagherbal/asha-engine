package generation2finiteoneformpuncturekernelweaksocketselectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_FINITE_ONE_FORM_PUNCTURE_KERNEL_WEAK_SOCKET_SELECTOR_AUDIT"
	theoremName = "Gate 893 — FiniteOneForm / PunctureKernel WeakSocket Selector Audit"
)

func Generation2FiniteOneFormPunctureKernelWeakSocketSelectorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 892 orientation obstruction inherited", Passed: a.WeakFrame.FullHMixesFrame && !a.WeakFrame.NativeOrientationSource && containsAll(a.WeakFrame.Failures, []string{FailureFullHMixesWeakSockets, FailureNoNativeHiggsOrientationSource}), Detail: FormatWeakFrame(a.WeakFrame)},
			{Name: "neutral null-edge route points to h_+ but is not native", Passed: a.NullEdge.YPlus1Zero && a.NullEdge.SelectsHPlusCandidate && !a.NullEdge.NativeOrientationTheorem && containsAll(a.NullEdge.Supports, []string{SupportNeutralNullEdgeSelectsHPlusCandidate}) && containsAll(a.NullEdge.Failures, []string{FailureNullEdgeNotNativeOrientation}), Detail: FormatNullEdge(a.NullEdge)},
			{Name: "finite one-form route is compatible but circular as source", Passed: a.OneForm.UsesWeakFrame && a.OneForm.CompatibleWithFrame && !a.OneForm.ForcesFrame && a.OneForm.CircularIfUsedAsSource && containsAll(a.OneForm.Failures, []string{FailureDFPatternRestatesOrientation}), Detail: FormatFiniteOneForm(a.OneForm)},
			{Name: "noncircularity requires independent weak socket selector", Passed: a.NonCircularity.OrientationToDFAllowed && !a.NonCircularity.DFToOrientationCertified && a.NonCircularity.RequiresIndependentSelector && !a.NonCircularity.NativeSelectorFunctional && containsAll(a.NonCircularity.Failures, []string{FailureNoNonCircularWeakSocketSelector}), Detail: FormatNonCircularity(a.NonCircularity)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4 physical generation flavor and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureNoNativeHiggsOrientationSource, FailureNoNonCircularWeakSocketSelector, FailureNoNativeDescentFullToOrient}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatWeakFrame(a.WeakFrame), FormatNullEdge(a.NullEdge), FormatFiniteOneForm(a.OneForm), FormatNonCircularity(a.NonCircularity), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
