package generation2sockettraceatomsectortypingr3eligibilityfirewallaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SOCKET_TRACE_ATOM_SECTOR_TYPING_R3_ELIGIBILITY_FIREWALL_AUDIT"
	theoremName = "Gate 885 — SocketTraceAtom SectorTyping and R3 Eligibility Firewall Audit"
)

func Generation2SocketTraceAtomSectorTypingR3EligibilityFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 884 readout rows and define socket trace atoms", Passed: len(a.Ledger.Atoms) == 3 && a.Ledger.Atoms[0].Atom == AtomPiPlus3 && a.Ledger.Atoms[1].Atom == AtomPiMinus3 && a.Ledger.Atoms[2].Atom == AtomPiMinus1 && containsAll(a.Ledger.Supports, []string{SupportPositiveReadoutRowsInherited, SupportSocketTraceLedgerR3Candidate}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify orthogonal complete active 3+3+1 ledger", Passed: a.Ledger.Orthogonal && a.Ledger.CompleteOnHRMin && a.Ledger.ActiveRank == RankHRMin && a.Ledger.ExpectedRank == RankHRMin && a.Ledger.Atoms[0].Rank == 3 && a.Ledger.Atoms[1].Rank == 3 && a.Ledger.Atoms[2].Rank == 1, Detail: FormatLedger(a.Ledger)},
			{Name: "verify post-orientation stability but not full unbroken A_F stability", Passed: a.Ledger.StableInOrientLayer && !a.Ledger.StableInFullAF && containsAll(a.Ledger.Failures, []string{FailureSocketAtomsNotStableFullAF, FailureNoFullUnbrokenAFSectorLedger}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify symbolic edge support typing", Passed: a.Ledger.EdgeTyped && a.Ledger.Atoms[0].EdgeSupport == EdgePiPlus3 && a.Ledger.Atoms[1].EdgeSupport == EdgePiMinus3 && a.Ledger.Atoms[2].EdgeSupport == EdgePiMinus1 && containsAll(a.Ledger.Supports, []string{SupportSocketAtomsEdgeSupportAtoms, SupportNotRandomProjectors}), Detail: FormatAtoms(a.Ledger.Atoms)},
			{Name: "verify trace magnitude rows still reproduce operator diagnostics", Passed: near(a.Ledger.TraceTotal, 3+3*AlphaB) && near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatLedger(a.Ledger)},
			{Name: "classify R3 candidate under alpha seal but not native R3", Passed: a.Ledger.Classification == Classification && a.Ledger.R3CandidateUnderSeal && !a.Ledger.NativeR3 && !a.Ledger.PhysicalSectors && containsAll(a.Ledger.Failures, []string{FailureAlphaStillSealed, FailureSocketAtomsNotNativeR3Sectors, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve official freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 885 R3 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
