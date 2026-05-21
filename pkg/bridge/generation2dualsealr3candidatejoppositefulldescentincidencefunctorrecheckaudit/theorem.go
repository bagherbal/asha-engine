package generation2dualsealr3candidatejoppositefulldescentincidencefunctorrecheckaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_DUALSEAL_R3_CANDIDATE_J_OPPOSITE_FULL_DESCENT_INCIDENCE_FUNCTOR_RECHECK_AUDIT"
	theoremName = "Gate 890 — DualSeal R3-Candidate J/Opposite Extension, Full-Descent, and IncidenceFunctor Recheck Audit"
)

func Generation2DualSealR3CandidateJOppositeFullDescentIncidenceFunctorRecheckAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 889 dual-seal R3 candidate ledger", Passed: a.Ledger.R3DualSealCandidate && a.Ledger.CompleteOnHRMin && a.Ledger.StableUnderAFOrient && !a.Ledger.StableUnderFullAF && a.Ledger.EdgeCompatible && a.Ledger.ReadoutComplete && containsAll(a.Ledger.Supports, []string{SupportProjectorAndReadoutCoherent, SupportOperatorNEffReproduced}), Detail: FormatLedger(a.Ledger)},
			{Name: "J mirror of active oriented ledger exists but does not complete full H_F^min", Passed: jExtensionOK(a.JExtension) && containsAll(a.JExtension.Supports, []string{SupportJMirrorExistsAtSeal, SupportJExtensionPreservesRanks, SupportActiveRightLedgerFormalOppositeCopy}) && containsAll(a.JExtension.Failures, []string{FailureJExtensionNotFullHFMinLedger, FailureNoFullJOppositeActionTheorem}), Detail: FormatJExtension(a.JExtension)},
			{Name: "full A_F descent remains blocked by post-orientation weak frame", Passed: a.Descent.AFOrientLedgerStable && !a.Descent.FullAFLedgerStable && a.Descent.FullToOrientHiggsRestriction && !a.Descent.NativeDescentCertified && !a.Descent.WeakSocketFrameFullHInvariant && containsAll(a.Descent.Failures, []string{FailureAFOrientNotFullAF, FailureSocketProjectorsNotStableFullH, FailureNoNativeDescentFullToOrient}), Detail: FormatDescent(a.Descent)},
			{Name: "BoundaryExteriorIncidenceFlagFunctor recheck adds no new native source", Passed: a.Incidence.BoundaryAlphaSealCoherent && a.Incidence.AlphaReconstructedUnderSeal && !a.Incidence.NewNativeBoundarySourceFound && !a.Incidence.NativeFunctorCertified && !a.Incidence.CrossLaneExclusionCertified && a.Incidence.AlphaStillSealed && containsAll(a.Incidence.Failures, []string{FailureGate890AddsNoBoundarySource, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}), Detail: FormatIncidence(a.Incidence)},
			{Name: "operator diagnostics remain diagnostic and official ledger remains frozen", Passed: near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) && a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4 physical generation flavor and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && !hasPhysicalLeak(a) && containsAll(a.Ledger.Failures, []string{FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}) && containsAll(a.Incidence.Failures, []string{FailureNoNativeIncidenceFunctor, FailureAlphaStillSealed}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatJExtension(a.JExtension), FormatDescent(a.Descent), FormatIncidence(a.Incidence), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func hasPhysicalLeak(a Audit) bool {
	for _, atom := range a.Ledger.Atoms {
		if atom.PhysicalSector || atom.GenerationResolved || atom.FlavorResolved || atom.IndividualYukawaValue {
			return true
		}
	}
	for _, mirror := range a.JExtension.Mirrors {
		if mirror.PhysicalSector || mirror.YukawaMagnitudeSource {
			return true
		}
	}
	return false
}
