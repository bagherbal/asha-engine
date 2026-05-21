package generation2sectortraceledgermapunderboundaryalphasealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SECTOR_TRACE_LEDGER_MAP_UNDER_BOUNDARY_ALPHA_SEAL_AUDIT"
	theoremName = "Gate 883 — SectorTraceLedgerMap Audit Under BoundaryAlpha Seal"
)

func Generation2SectorTraceLedgerMapUnderBoundaryAlphaSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "define active socket trace atoms", Passed: len(a.Ledger.Atoms) == 3 && a.Ledger.Atoms[0].Name == AtomPiPlus3 && a.Ledger.Atoms[1].Name == AtomPiMinus3 && a.Ledger.Atoms[2].Name == AtomPiMinus1 && containsAll(a.Ledger.Supports, []string{SupportSocketProjectorsTraceAtoms, SupportAggregateRefinedToSocketAtoms}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify orthogonal complete 3+3+1 ledger on H_R^min", Passed: a.Ledger.Orthogonal && a.Ledger.CompleteOnHRMin && a.Ledger.ActiveRank == RankHRMin && a.Ledger.ExpectedRank == RankHRMin && a.Ledger.Atoms[0].Rank == RankPiPlus3 && a.Ledger.Atoms[1].Rank == RankPiMinus3 && a.Ledger.Atoms[2].Rank == RankPiMinus1, Detail: FormatAtoms(a.Ledger.Atoms)},
			{Name: "verify positive trace weights under alpha seal", Passed: a.Ledger.PositiveWeights && a.Ledger.Atoms[0].Weight > 0 && a.Ledger.Atoms[1].Weight > 0 && a.Ledger.Atoms[2].Weight > 0 && containsAll(a.Ledger.Supports, []string{SupportYDaggerYPositiveWeights}), Detail: FormatLedger(a.Ledger)},
			{Name: "reproduce trace square trace and operator N_eff", Passed: near(a.Ledger.TraceTotal, 3+3*AlphaB) && near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatLedger(a.Ledger)},
			{Name: "classify as R3 candidate under alpha seal but not native R3", Passed: a.Ledger.Classification == Classification && a.Ledger.ConditionalCandidate && !a.Ledger.NativeR3 && !a.Ledger.OfficialUpdatesAllowed && containsAll(a.Ledger.Failures, []string{FailureAlphaStillSealed, FailureSocketLedgerNotNativeR3, FailureSocketProjectorsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve official freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 883 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
