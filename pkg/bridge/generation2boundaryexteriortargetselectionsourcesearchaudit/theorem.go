package generation2boundaryexteriortargetselectionsourcesearchaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_EXTERIOR_TARGET_SELECTION_SOURCE_SEARCH_AUDIT"
	theoremName = "Gate 875 — BoundaryExterior Target-Selection Source Search Audit"
)

func Generation2BoundaryExteriorTargetSelectionSourceSearchAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 874 conditional proxy plateau and exact wound", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdateOfficial && a.Wound.NeedLambda1ToPiTop && a.Wound.NeedLambda2ToHRMin && a.Wound.NeedNoLambda1ToHRMin && a.Wound.NeedNoLambda2ToPiTop && a.Wound.MissingObject == "BoundaryExteriorTargetSelectionFunctor", Detail: FormatLedger(a.Ledger) + " | " + FormatWound(a.Wound)},
			{Name: "audit puncture/complement route as strongest source candidate", Passed: a.PunctureRoute.Name == "puncture/complement route" && a.PunctureRoute.Strength == "strongest internal source candidate" && !a.PunctureRoute.NativeFunctor && containsAll(a.PunctureRoute.Supports, []string{SupportPunctureComplementStrongestRoute, SupportExposureVisibleComplementPiTop, SupportEnclosurePuncturedActiveDomain}), Detail: FormatRoute(a.PunctureRoute)},
			{Name: "audit boundary degree/support codimension route", Passed: a.CodimRoute.Name == "boundary degree / support codimension route" && !a.CodimRoute.NativeFunctor && containsAll(a.CodimRoute.Failures, []string{FailureCodimRouteNotFunctor, FailureNoNativeTargetSelectionFunctor}), Detail: FormatRoute(a.CodimRoute)},
			{Name: "audit trace-normalization chamber route", Passed: a.ChamberRoute.Name == "trace-normalization chamber route" && !a.ChamberRoute.NativeFunctor && containsAll(a.ChamberRoute.Failures, []string{FailureChamberRouteNotFunctor, FailureResponseChamberTypingNotTheorem}), Detail: FormatRoute(a.ChamberRoute)},
			{Name: "reconstruct alpha while preserving target-selection obstruction", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.TargetSelectionNative && containsAll(a.Candidate.Failures, []string{FailureNoNativeTargetSelectionFunctor, FailureAlphaStillSealed}), Detail: FormatCandidate(a.Candidate)},
			{Name: "block R3/R4 and official ledger updates", Passed: a.R3.ConditionalTraceProxyMature && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 875 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatWound(a.Wound), FormatRoute(a.PunctureRoute), FormatRoute(a.CodimRoute), FormatRoute(a.ChamberRoute), FormatCandidate(a.Candidate), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
