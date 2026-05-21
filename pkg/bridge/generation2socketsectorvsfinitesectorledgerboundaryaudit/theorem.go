package generation2socketsectorvsfinitesectorledgerboundaryaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SOCKET_SECTOR_VS_FINITE_SECTOR_LEDGER_BOUNDARY_AUDIT"
	theoremName = "Gate 886 — SocketSector vs FiniteSector Ledger Boundary Audit"
)

func Generation2SocketSectorVsFiniteSectorLedgerBoundaryAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 885 socket trace atom ledger", Passed: len(a.Ledger.Atoms) == 3 && a.Ledger.ActiveRank == RankHRMin && containsAll(a.Ledger.Supports, []string{SupportSocketLedgerPostOrientationCandidate, SupportSocketAtomsEdgeAndReadoutStable}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify socket atoms are typed only in A_F^orient", Passed: a.Ledger.SocketSectorTyped && a.Ledger.StableInOrientLayer && !a.Ledger.StableInFullAF && a.Ledger.PostOrientationAlgebra == PostOrientationAlgebra && a.Ledger.FullUnbrokenAlgebra == FullUnbrokenAlgebra && containsAll(a.Ledger.Failures, []string{FailureSocketAtomsNotStableFullAF, FailureNoFullUnbrokenAFSectorLedger}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify no finite-sector lift is certified", Passed: !a.Lift.LiftCertified && !a.Lift.TargetCertified && !a.Lift.NativeR3 && a.Lift.MapName == MissingLiftMap && a.Lift.Sigma == MissingSigmaMap && a.Lift.TargetLedger == FiniteSectorLedger && containsAll(a.Lift.Failures, []string{FailureNoSocketToFiniteSectorMap, FailurePostOrientNotNativeFinite}), Detail: FormatLift(a.Lift)},
			{Name: "verify socket atoms are not physical/generation/flavor assignments", Passed: !a.Ledger.PhysicalSectors && !a.Lift.PhysicalAssignment && !a.Lift.GenerationCarrierPresent && !a.Lift.FlavorOrientationPresent && containsAll(a.Ledger.Failures, []string{FailureNoPhysicalSectorAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatLedger(a.Ledger)},
			{Name: "verify trace diagnostics remain inherited and diagnostic", Passed: near(a.Ledger.TraceTotal, 3+3*AlphaB) && near(a.Ledger.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatLedger(a.Ledger)},
			{Name: "verify classification and next branch", Passed: a.Ledger.Classification == Classification && a.Ledger.NextBranch == NextBranch && a.Ledger.R3CandidateUnderSeal && !a.Ledger.NativeR3 && containsAll(a.Ledger.Supports, []string{SupportR3FrontierRequiresLiftMap, SupportSocketSectorsTypedNotFinite, SupportBoundaryClarified}), Detail: FormatLedger(a.Ledger)},
			{Name: "preserve official ledger freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 886 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatLift(a.Lift), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
