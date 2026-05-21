package generation2sectortracemagnitudereadoutmapunderboundaryalphasealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SECTOR_TRACE_MAGNITUDE_READOUT_MAP_UNDER_BOUNDARY_ALPHA_SEAL_AUDIT"
	theoremName = "Gate 884 — SectorTraceMagnitude ReadoutMap Under BoundaryAlpha Seal Audit"
)

func Generation2SectorTraceMagnitudeReadoutMapUnderBoundaryAlphaSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 883 socket atoms and define readout rows", Passed: len(a.Readout.Rows) == 3 && a.Readout.Rows[0].Atom == AtomPiPlus3 && a.Readout.Rows[1].Atom == AtomPiMinus3 && a.Readout.Rows[2].Atom == AtomPiMinus1 && containsAll(a.Readout.Supports, []string{SupportYDaggerYTraceMagnitudeReadout, SupportR3PreparationAdvancesToReadout}), Detail: FormatReadout(a.Readout)},
			{Name: "verify orthogonal complete positive socket ledger", Passed: a.Readout.Orthogonal && a.Readout.CompleteOnHRMin && a.Readout.Positive && a.Readout.ActiveRank == RankHRMin && a.Readout.ExpectedRank == RankHRMin, Detail: FormatReadout(a.Readout)},
			{Name: "verify row contribution formulas", Passed: near(a.Readout.Rows[0].TraceContribution, 3) && near(a.Readout.Rows[0].SquareContribution, 3) && near(a.Readout.Rows[1].TraceContribution, 3*AlphaB*(1-AlphaB)) && near(a.Readout.Rows[1].SquareContribution, 3*AlphaB*AlphaB*(1-AlphaB)*(1-AlphaB)) && near(a.Readout.Rows[2].TraceContribution, 3*AlphaB*AlphaB) && near(a.Readout.Rows[2].SquareContribution, 9*AlphaB*AlphaB*AlphaB*AlphaB), Detail: FormatRows(a.Readout.Rows)},
			{Name: "reconstruct trace square trace and operator N_eff from rows", Passed: near(a.Readout.TraceTotal, 3+3*AlphaB) && near(a.Readout.SquareTrace, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) && near(a.Readout.OperatorNEff, OperatorNEffDiagnostic) && near(a.Readout.OperatorCYukawa, OperatorCYukawaDiagnostic), Detail: FormatReadout(a.Readout)},
			{Name: "classify readout under alpha seal but not native R3", Passed: a.Readout.Classification == Classification && a.Readout.ConditionalReadout && !a.Readout.NativeR3 && !a.Readout.OfficialUpdatesAllowed && containsAll(a.Readout.Failures, []string{FailureAlphaStillSealed, FailureReadoutUnderSealNotNative, FailureSocketTraceNotNativeR3, FailureSocketAtomsNotPhysical, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues}), Detail: FormatReadout(a.Readout)},
			{Name: "preserve official freeze", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) && !near(a.Freeze.OperatorCYukawa, a.Freeze.OfficialCYukawa), Detail: FormatFreeze(a.Freeze)},
			{Name: "preserve Gate 884 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatReadout(a.Readout), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
