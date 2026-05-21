package generation2r3sectorledgerrequirementsunderboundaryalphasealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_R3_SECTOR_LEDGER_REQUIREMENTS_UNDER_BOUNDARY_ALPHA_SEAL_AUDIT"
	theoremName = "Gate 882 — R3 SectorLedger Requirements Under BoundaryAlpha Seal Audit"
)

func Generation2R3SectorLedgerRequirementsUnderBoundaryAlphaSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 881 conditional supplies", Passed: a.Supplies.BoundaryAlphaSeal && a.Supplies.PostOrientationFiniteTripleSeal && a.Supplies.SymbolicDFEdgeMatrix && a.Supplies.YDaggerYPositiveReadout && a.Supplies.AggregateHAgg && a.Supplies.DiagnosticOnly && containsAll(a.Supplies.Supports, []string{SupportR3PreparationUnderSeal, SupportAggregateTraceProxyInput, SupportYDaggerYPositiveFiniteBody}), Detail: FormatSupplies(a.Supplies)},
			{Name: "audit R3 requirements", Passed: !a.Requirements.NativeR3Ready && !a.Requirements.TypedSectorProjectors && !a.Requirements.SectorTraceAtoms && !a.Requirements.PositiveReadoutMap && !a.Requirements.SectorLedgerConsistency && !a.Requirements.NonCircularAlphaSource && a.Requirements.GenerationFlavorFirewall && containsAll(a.Requirements.Failures, []string{FailureAlphaStillSealed, FailureNoSectorTraceLedgerMap, FailureNoSectorTraceMagnitudeMap, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap}), Detail: FormatRequirements(a.Requirements)},
			{Name: "rank blockers", Passed: len(a.Blockers) == 5 && a.Blockers[0].Name == BlockerBoundaryIncidenceFunctor && a.Blockers[0].Priority == 1 && a.Blockers[1].Name == BlockerSectorTraceLedgerMap && a.Blockers[1].Priority == 2 && a.Blockers[0].BlocksR3 && a.Blockers[1].BlocksR3, Detail: FormatBlockers(a.Blockers)},
			{Name: "classify under-seal R3 candidate but not R3", Passed: a.Eligibility.Classification == R2Status && a.Eligibility.ConditionalCandidate && !a.Eligibility.NativeR3 && !a.Eligibility.NativeR4 && !a.Eligibility.OfficialUpdatesAllowed && a.Eligibility.NextBranch == NextRecommendedBranch && containsAll(a.Eligibility.Failures, []string{FailureAlphaStillSealed, FailureNoNativeR3SectorLedger, FailureAggregateProxyNotSectorLedger}), Detail: FormatEligibility(a.Eligibility)},
			{Name: "preserve diagnostic official freeze", Passed: a.Ledger.Frozen && a.Ledger.DiagnosticOnly && !a.Ledger.CanUpdate && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic) && near(a.Ledger.OperatorCHiggs, OperatorCHiggsDiagnostic) && !near(a.Ledger.OperatorNEff, a.Ledger.OfficialNEff), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve Gate 882 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatSupplies(a.Supplies), FormatRequirements(a.Requirements), FormatBlockers(a.Blockers), FormatEligibility(a.Eligibility), FormatLedger(a.Ledger), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
