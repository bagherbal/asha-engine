package generation2dualsealr3candidateclassificationpromotionfirewallaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_DUALSEAL_R3_CANDIDATE_CLASSIFICATION_PROMOTION_FIREWALL_AUDIT"
	theoremName = "Gate 889 — DualSeal R3-Candidate Classification and Promotion Firewall Audit"
)

func Generation2DualSealR3CandidateClassificationPromotionFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit complete dual-seal projector/readout ledger", Passed: a.Ledger.R3SealedCandidate && a.Ledger.CompleteOnHRMin && a.Ledger.ReadoutComplete && a.Ledger.EdgeCompatible && containsAll(a.Ledger.Supports, []string{SupportR3CandidateUnderDualSeal, SupportProjectorsAndReadoutComplete}), Detail: FormatLedger(a.Ledger)},
			{Name: "operator diagnostic values reproduced but official ledger frozen", Passed: near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) && a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "classify as R3 candidate under dual seal", Passed: a.Seals.BoundaryAlphaSealSuppliesWeights && a.Seals.PostOrientationSealSuppliesProjectors && a.Seals.ProjectorLedgerCompleteUnderSeals && a.Seals.TraceReadoutCompleteUnderSeals && !a.Seals.NativeR3 && !a.Seals.R4NativeYukawa && containsAll(a.Seals.Supports, []string{SupportR3CandidateUnderDualSeal, SupportBoundaryAlphaSuppliesWeights, SupportPostOrientationSuppliesProjector}), Detail: FormatSeals(a.Seals)},
			{Name: "native R3 blockers reduced to seal removal", Passed: !a.Requirements.Satisfied && a.Requirements.NeedsNativeIncidenceFunctor && a.Requirements.NeedsNativeCrossLaneExclusion && a.Requirements.NeedsFullUnbrokenAFProjectors && a.Requirements.NeedsSealFreeTraceMagnitudeReadout && containsAll(a.Requirements.Failures, []string{FailureNoNativeIncidenceFunctor, FailurePostOrientationNotFullAF, FailureNoNativeR3SectorLedger}), Detail: FormatRequirements(a.Requirements)},
			{Name: "physical particle generation flavor and individual-yukawa claims blocked", Passed: !hasPhysicalLeak(a) && containsAll(a.Ledger.Failures, []string{FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve native R3/R4 promotion firewalls", Passed: firewallsOK(a.Firewalls) && containsAll(a.Seals.Failures, []string{FailureNotNativeR3, FailureNoR4, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSeals(a.Seals), FormatRequirements(a.Requirements), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
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
	return false
}
