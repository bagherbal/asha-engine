package generation2neutralpunctureairlockunificationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NEUTRAL_PUNCTURE_AIRLOCK_UNIFICATION_AUDIT"
	theoremName = "Gate 895 — NeutralPuncture Airlock Unification Audit"
)

func Generation2NeutralPunctureAirlockUnificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 894 minimal null-edge orientation candidate inherited", Passed: containsAll(Statuses(), []string{StatusGate894Inherited}) && a.WeakKernel.QuotientIsKernel, Detail: FormatWeakKernel(a.WeakKernel)},
			{Name: "neutral puncture is defined before weak-socket orientation", Passed: a.PunctureIndependence.RequiresRightCharacterSplit && a.PunctureIndependence.RequiresLeptoColorSplit && !a.PunctureIndependence.RequiresWeakSocketFrame && a.PunctureIndependence.DefinedBeforeWeakOrientation, Detail: FormatPunctureIndependence(a.PunctureIndependence)},
			{Name: "puncture flag reconstructs alpha targets", Passed: a.AlphaFlag.F0SubsetF1SubsetF2 && a.AlphaFlag.RankQ1 == 3 && a.AlphaFlag.RankQ2 == 7 && a.AlphaFlag.ReconstructsAlpha && !a.AlphaFlag.NativeAlphaFunctor, Detail: FormatAlphaFlag(a.AlphaFlag)},
			{Name: "minimal image reconstructs h_+ tensor P_1 kernel candidate", Passed: a.WeakKernel.HLeftRank == 8 && a.WeakKernel.ImageRank == 7 && a.WeakKernel.KernelRank == 1 && a.WeakKernel.QuotientIsKernel && a.WeakKernel.CanReconstructWeakFrameCandidate && !a.WeakKernel.NativeMinimalImageRule && a.WeakKernel.DependsOnMinimalImageChoice, Detail: FormatWeakKernel(a.WeakKernel)},
			{Name: "alpha and orientation seals collapse to one puncture-airlock candidate, not theorem", Passed: a.Airlock.ControlsAlphaFlag && a.Airlock.ControlsWeakKernel && a.Airlock.TwoSealProblemReducesToOne && !a.Airlock.NativeAirlockFunctor && containsAll(a.Airlock.Failures, []string{FailureNoNeutralPunctureAirlockFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector}), Detail: FormatAirlock(a.Airlock)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, alpha, orientation, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNeutralPunctureAirlockFunctor}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatPunctureIndependence(a.PunctureIndependence), FormatAlphaFlag(a.AlphaFlag), FormatWeakKernel(a.WeakKernel), FormatAirlock(a.Airlock), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
