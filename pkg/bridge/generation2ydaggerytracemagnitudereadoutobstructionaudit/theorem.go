package generation2ydaggerytracemagnitudereadoutobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_Y_DAGGER_Y_TRACE_MAGNITUDE_READOUT_OBSTRUCTION_AUDIT"
	theoremName = "Gate 864 — Y^dagger Y TraceMagnitude Readout Obstruction Audit"
)

func Generation2YDaggerYTraceMagnitudeReadoutObstructionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 863 post-orientation finite-triple seal", Passed: a.R3.PostOrientationFiniteTripleSeal && containsAll(a.Readout.Supports, []string{StatusYdaggerYConstructed}), Detail: FormatR3(a.R3)},
			{Name: "construct Y^dagger Y positive right-module readout candidate", Passed: a.Readout.Positive && a.Readout.RightModuleReadout && a.Readout.CorrectActiveSupport && a.Readout.PunctureAbsent && a.Readout.LeftKernelExcluded && containsAll(a.Readout.Supports, []string{SupportYdaggerYCandidate, SupportCarrierWiseYes}), Detail: FormatReadout(a.Readout)},
			{Name: "compute required socket magnitudes for aggregate table match", Passed: weightsOK(a.Readout.Weights) && containsAll(a.Readout.Supports, []string{StatusRequiredWeights, SupportRequiredSocketValues}), Detail: FormatWeights(a.Readout.Weights)},
			{Name: "reconstruct aggregate trace only conditionally if socket values are inserted", Passed: near(a.Readout.TraceIfMatched, TraceIfMatched(AlphaB)) && near(a.Readout.SquareTraceIfMatched, SquareTraceIfMatched(AlphaB)) && near(a.Readout.OperatorNEffIfMatched, OperatorNEffIfMatched(AlphaB)) && !a.Readout.MagnitudeWiseMatch && !a.Readout.SocketMagnitudesDerived && a.Readout.RequiresInsertedSocketValues && containsAll(a.Readout.Failures, []string{FailureOnlyIfValuesInserted, FailureSocketValuesRestateTable}), Detail: FormatReadout(a.Readout)},
			{Name: "reduce R3 pressure to missing socket magnitude source", Passed: a.Obstruction.YSocketMagnitudeSourceMissing && a.Obstruction.AlphaSourceMissing && !a.Obstruction.TraceReadoutNative && !a.Obstruction.NonCircularMagnitudes && a.Obstruction.NextMissingObject != "" && containsAll(a.Obstruction.Supports, []string{SupportR3PressureSocketSource}) && containsAll(a.Obstruction.Failures, []string{FailureYMagnitudesNotDerived, FailureReadoutNotNative, FailureNoSectorTraceMagnitude}), Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and official ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && a.Ledger.OfficialFrozen, Detail: FormatLedger(a.Ledger) + " | " + FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 864 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatReadout(a.Readout), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
