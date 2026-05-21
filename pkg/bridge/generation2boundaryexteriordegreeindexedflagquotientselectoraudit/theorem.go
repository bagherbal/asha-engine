package generation2boundaryexteriordegreeindexedflagquotientselectoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_EXTERIOR_DEGREE_INDEXED_FLAG_QUOTIENT_SELECTOR_AUDIT"
	theoremName = "Gate 878 — BoundaryExterior Degree-Indexed FlagQuotient Selector Audit"
)

func Generation2BoundaryExteriorDegreeIndexedFlagQuotientSelectorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 877 plateau and official freeze", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdate && !a.Impact.CanUpdateNEff && a.Impact.DegreeIndexedSelectorCandidate, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "validate puncture flag", Passed: a.Flag.ValidNestedFlag && a.Flag.RankF0 == 1 && a.Flag.RankF1 == 4 && a.Flag.RankF2 == 8, Detail: FormatFlag(a.Flag)},
			{Name: "dimension mismatch forces selector typing", Passed: a.Mismatch.SelectorNotLinearSurjection && a.Mismatch.Lambda1Dim == 2 && a.Mismatch.F1OverF0Rank == 3 && a.Mismatch.Lambda2Dim == 1 && a.Mismatch.F2OverF0Rank == 7 && containsAll(a.Mismatch.Failures, []string{FailureNotLinearMap}), Detail: FormatMismatch(a.Mismatch)},
			{Name: "degree one selects F1/F0", Passed: a.DegreeOne.SelectorMode && !a.DegreeOne.LinearSurjection && !a.DegreeOne.NativeSelector && a.DegreeOne.SelectedQuotient == "F_1/F_0" && a.DegreeOne.Target == "Pi_top" && a.DegreeOne.QuotientRank == 3 && containsAll(a.DegreeOne.Supports, []string{SupportLambda1SelectsF1OverF0, SupportDimensionMismatchTypeCorrection}), Detail: FormatSelector(a.DegreeOne)},
			{Name: "degree two selects cumulative F2/F0", Passed: a.DegreeTwo.SelectorMode && !a.DegreeTwo.LinearSurjection && !a.DegreeTwo.NativeSelector && a.DegreeTwo.SelectedQuotient == "F_2/F_0" && a.DegreeTwo.Target == "H_R^min" && a.DegreeTwo.QuotientRank == 7 && containsAll(a.DegreeTwo.Supports, []string{SupportLambda2SelectsF2OverF0, SupportDegreeTwoCumulativeEnclosure}), Detail: FormatSelector(a.DegreeTwo)},
			{Name: "reject F2/F1 pure associated graded slice", Passed: a.WrongSlice.Rejected && a.WrongSlice.RejectedRank == 4 && a.WrongSlice.RequiredRank == 7 && containsAll(a.WrongSlice.Failures, []string{FailureDegreeTwoNotF2OverF1}), Detail: FormatRejected(a.WrongSlice)},
			{Name: "cross-lane exclusion remains conditional", Passed: a.CrossLane.ExcludedIfSelector && !a.CrossLane.NativeExclusion && containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeDegreeIndexedSelector}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "alpha reconstructed from degree-indexed selectors but remains sealed", Passed: a.Alpha.ReconstructedFromSelectors && near(a.Alpha.ReconstructedAlpha, AlphaB) && !a.Alpha.NativeSelectorFunctor && containsAll(a.Alpha.Failures, []string{FailureNoNativeDegreeIndexedSelector, FailureAlphaStillSealed}), Detail: FormatAlpha(a.Alpha)},
			{Name: "block R3/R4 and ledger updates", Passed: a.R3.ConditionalTraceProxyMature && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanPromoteToR3 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 878 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatFlag(a.Flag), FormatMismatch(a.Mismatch), FormatSelector(a.DegreeOne), FormatSelector(a.DegreeTwo), FormatRejected(a.WrongSlice), FormatCrossLane(a.CrossLane), FormatAlpha(a.Alpha), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
