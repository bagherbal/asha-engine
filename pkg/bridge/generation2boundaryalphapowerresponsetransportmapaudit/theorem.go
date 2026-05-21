package generation2boundaryalphapowerresponsetransportmapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_ALPHA_POWER_RESPONSE_TRANSPORTMAP_AUDIT"
	theoremName = "Gate 867 — BoundaryAlpha Power-Response TransportMap Audit"
)

func Generation2BoundaryAlphaPowerResponseTransportMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 866 socket-rank alpha source typing", Passed: containsAll(a.Alpha.Supports, []string{StatusGate866Inherited, SupportSocketRankShapeInherited}) && a.R3.SocketMagnitudeTransferTyped, Detail: FormatAlpha(a.Alpha)},
			{Name: "audit linear dominant socket first-order lane", Passed: a.Alpha.Linear.Power == 1 && a.Alpha.Linear.Numerator == PiTopRank && a.Alpha.Linear.Denominator == LinearDenominator && a.Alpha.Linear.SocketRankSourced && !a.Alpha.Linear.TransportCertified && !a.Alpha.Linear.ResponseOrderDerived, Detail: FormatLane(a.Alpha.Linear)},
			{Name: "audit quadratic active right-domain second-order lane", Passed: a.Alpha.Quadratic.Power == 2 && a.Alpha.Quadratic.Numerator == HRminRank && a.Alpha.Quadratic.Denominator == QuadraticDenominator && a.Alpha.Quadratic.SocketRankSourced && !a.Alpha.Quadratic.TransportCertified && !a.Alpha.Quadratic.ResponseOrderDerived, Detail: FormatLane(a.Alpha.Quadratic)},
			{Name: "reconstruct alpha_B while blocking native transport", Passed: a.Alpha.ShapeCoherent && near(a.Alpha.ReconstructedAlpha, AlphaB) && !a.Alpha.Native && !a.Alpha.TransportMapCertified, Detail: FormatAlpha(a.Alpha)},
			{Name: "audit shared S_split transport requirement", Passed: a.SharedS.SameCoordinateUsed && a.SharedS.BothCodomainsTyped && !a.SharedS.TransportIntoBothCertified && !a.SharedS.PowerOrderDerived && containsAll(a.SharedS.Failures, []string{FailureSameSNotTransportedToBoth}), Detail: FormatSharedS(a.SharedS)},
			{Name: "audit boundary-augmented denominator typing without activation theorem", Passed: a.Denominators.TypedBoundaryAugmentedDomains && !a.Denominators.ActivationTheoremCertified && containsAll(a.Denominators.Failures, []string{FailureDenominatorTypingNotActivation}), Detail: FormatDenominators(a.Denominators)},
			{Name: "block power-response transport map and R3/R4 promotion", Passed: !a.Obstruction.TransportMapCertified && !a.Obstruction.PowerOrderDerived && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatObstruction(a.Obstruction) + " | " + FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 867 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatAlpha(a.Alpha), FormatSharedS(a.SharedS), FormatDenominators(a.Denominators), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
