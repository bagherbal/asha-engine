package generation2boundaryalphasocketranksourcereentryaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_ALPHA_SOCKET_RANK_SOURCE_REENTRY_AUDIT"
	theoremName = "Gate 866 — BoundaryAlpha SocketRank Source Re-entry Audit"
)

func Generation2BoundaryAlphaSocketRankSourceReEntryAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 865 alpha wound", Passed: containsAll(a.Alpha.Supports, []string{StatusGate865Inherited}) && a.R3.SocketMagnitudePressureReducedToAlpha, Detail: FormatR3(a.R3)},
			{Name: "source-type 3/10 by dominant socket rank over 8+2", Passed: a.Alpha.LinearLane.Numerator == PiTopRank && a.Alpha.LinearLane.Denominator == LinearDenominator && near(a.Alpha.LinearLane.Coefficient, 0.3) && containsAll(a.Alpha.LinearLane.Supports, []string{SupportThreeSourceIsPiTopRank, SupportLinearLaneSocketRank}), Detail: FormatRankSource(a.Alpha.LinearLane)},
			{Name: "source-type 7/72 by H_R^min rank over H72", Passed: a.Alpha.QuadraticLane.Numerator == HRminRank && a.Alpha.QuadraticLane.Denominator == QuadraticDenominator && near(a.Alpha.QuadraticLane.Coefficient, 7.0/72.0) && containsAll(a.Alpha.QuadraticLane.Supports, []string{SupportSevenSourceIsHRMinRank, SupportQuadraticLaneSocketRank}), Detail: FormatRankSource(a.Alpha.QuadraticLane)},
			{Name: "reconstruct alpha_B from socket-rank source candidates", Passed: a.Alpha.ReconstructsAlpha && near(a.Alpha.ReconstructedAlpha, AlphaB) && !a.Alpha.Native && !a.Alpha.TransportMapCertified, Detail: FormatAlpha(a.Alpha)},
			{Name: "preserve dual-seven firewall", Passed: a.DualSeven.SameInteger && !a.DualSeven.Identified && !a.DualSeven.TypedMapCertified && containsAll(a.DualSeven.Failures, []string{FailureNoHRMinToK7Map, FailureDualSevenNotIdentified}), Detail: FormatDualSeven(a.DualSeven)},
			{Name: "block activation-map and R3/R4 promotion", Passed: !a.Obstruction.ActivationMapCertified && !a.Obstruction.AlphaNative && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatObstruction(a.Obstruction) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 866 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatAlpha(a.Alpha), FormatDualSeven(a.DualSeven), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
