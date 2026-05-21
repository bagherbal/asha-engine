package generation2socketmagnitudesourcebernoullibminusltransferaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_SOCKET_MAGNITUDE_SOURCE_BERNOULLI_BMINUSL_TRANSFER_AUDIT"
	theoremName = "Gate 865 — SocketMagnitude Source and Bernoulli/B-L Transfer Audit"
)

func Generation2SocketMagnitudeSourceBernoulliBMinusLTransferAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 864 Y^dagger Y readout obstruction", Passed: a.R3.YdaggerYReadoutCarrier && containsAll(Statuses(), []string{StatusGate864Inherited}), Detail: FormatR3(a.R3)},
			{Name: "source-type required socket magnitudes given sealed alpha_B", Passed: magnitudesOK(a.Transfer.Magnitudes) && a.Transfer.DominantNormalization && a.Transfer.RestBMinusLTransfer && containsAll(a.Transfer.Supports, []string{SupportGivenAlphaSocketMagnitudes, SupportRestColorBernoulli, SupportRestLeptonTripletQuadratic}), Detail: FormatTransfer(a.Transfer)},
			{Name: "recover B-L rest transfer and Bernoulli/quadratic socket pattern", Passed: a.Transfer.BernoulliComplement && a.Transfer.TripletQuadraticTransfer && a.Transfer.TraceZeroRedistribution && a.Transfer.YdaggerYEqualsHAggGivenMagnitudes && !a.Transfer.Native && !a.Transfer.NonCircular, Detail: FormatTransfer(a.Transfer)},
			{Name: "reproduce trace and square-trace diagnostics given transfer magnitudes", Passed: a.Trace.TracePreserved && a.Trace.SquareTraceMatches && near(a.Trace.RestTrace, 3*AlphaB) && near(a.Trace.RestSquareTrace, 3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB), Detail: FormatTrace(a.Trace)},
			{Name: "enforce noncircularity and reduce wound to alpha_B source", Passed: a.Obstruction.LayerA_GivenAlphaSourceTyped && !a.Obstruction.LayerB_AlphaDerived && !a.Obstruction.SocketMagnitudeNative && !a.Obstruction.NonCircularSource && containsAll(a.Obstruction.Failures, []string{FailureAlphaStillSealed, FailureSocketNotNativeNoAlpha, FailureNoNativeTransferTheorem}), Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and official ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4 && a.Ledger.OfficialFrozen, Detail: FormatLedger(a.Ledger) + " | " + FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 865 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatTransfer(a.Transfer), FormatTrace(a.Trace), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
