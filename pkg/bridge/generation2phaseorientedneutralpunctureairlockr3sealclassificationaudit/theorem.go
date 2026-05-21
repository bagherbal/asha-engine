package generation2phaseorientedneutralpunctureairlockr3sealclassificationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_PHASE_ORIENTED_NEUTRAL_PUNCTURE_AIRLOCK_R3_SEAL_CLASSIFICATION_AUDIT"
	theoremName = "Gate 900 — PhaseOriented NeutralPuncture Airlock R3-Seal Classification Audit"
)

func Generation2PhaseOrientedNeutralPunctureAirlockR3SealClassificationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "mature sealed R3-candidate chain assembled", Passed: a.MatureChain.NeutralAirlockFamily && a.MatureChain.RightPhaseSealRequired && a.MatureChain.BoundaryAlphaReconstructed && a.MatureChain.HiggsOrientationPunctureKernel && a.MatureChain.ProjectorLedger && a.MatureChain.PositiveReadoutRows && a.MatureChain.OperatorNEffReconstructed && containsAll(a.MatureChain.Supports, []string{SupportR3SealedCandidateComplete, SupportNeutralAirlockUnifiesWounds}), Detail: FormatMatureChain(a.MatureChain)},
			{Name: "ordered representative exists only under phase orientation seal", Passed: a.Representative.Puncture == PunctureOrderedPlus && a.Representative.PhaseOrder == RightPhaseOrder && a.Representative.HRMinComplete && !a.Representative.SelectedNatively && containsAll(a.Representative.Supports, []string{SupportPhaseOrientationSelectsAirlock, SupportOrderedRepresentative}) && containsAll(a.Representative.Failures, []string{FailureNoNativeSelectionSigmaPlus, FailureNoNativeRightPhaseOrientation}), Detail: FormatRepresentative(a.Representative)},
			{Name: "Y dagger Y positive readout reproduces operator diagnostics under seals", Passed: a.Readout.Positive && a.Readout.ReproducesNEff && near(a.Readout.OperatorNEff, OperatorNEffDiagnostic) && near(a.Readout.OperatorCYukawa, OperatorCYukawaDiagnostic) && containsAll(a.Readout.Supports, []string{SupportYDagYReproducesOperatorNEff, SupportLedgerRowsStable}), Detail: FormatReadout(a.Readout)},
			{Name: "native R3 promotion checklist fails exactly by seal dependence", Passed: !a.Promotion.AllowedNativeR3 && !a.Promotion.NativeAlphaSource && !a.Promotion.NativePhaseOrientation && !a.Promotion.FullAFDescent && !a.Promotion.NativeSectorLedger && !a.Promotion.PhysicalInterpretation && containsAll(a.Promotion.Failures, []string{FailureAlphaStillSealed, FailureNoNativeRightPhaseOrientation, FailureNoNativeDescentFullAF, FailureNoNativeR3SectorTraceLedger, FailureNoPhysicalParticleAssignment}), Detail: FormatPromotion(a.Promotion)},
			{Name: "official ledgers remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, alpha, phase, airlock, descent, physical-sector, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus, FailureNoR4NativeYukawaTheorem}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatMatureChain(a.MatureChain), FormatRepresentative(a.Representative), FormatReadout(a.Readout), FormatPromotion(a.Promotion), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
