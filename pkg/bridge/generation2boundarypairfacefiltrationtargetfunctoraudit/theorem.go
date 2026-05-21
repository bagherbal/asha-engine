package generation2boundarypairfacefiltrationtargetfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_PAIR_FACE_FILTRATION_TARGET_FUNCTOR_AUDIT"
	theoremName = "Gate 877 — BoundaryPair Face-Filtration TargetFunctor Audit"
)

func Generation2BoundaryPairFaceFiltrationTargetFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 876 plateau and ledger freeze", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdate && !a.Impact.CanUpdateNEff && a.Impact.DegreeToFlagFunctorCandidate, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "validate puncture-complement flag", Passed: a.Flag.ValidNestedFlag && a.Flag.RankP == 1 && a.Flag.RankF1 == 4 && a.Flag.RankF2 == 8 && containsAll(a.Flag.Failures, []string{FailureNoNativeFlagFunctor, FailureNoNativeDegreeToFlagFunctor}), Detail: FormatFlag(a.Flag)},
			{Name: "degree one targets F1/p = Pi_top", Passed: a.DegreeOne.Degree == "Lambda^1 B_2" && a.DegreeOne.Quotient == "F_1/p" && a.DegreeOne.QuotientRank == PiTopRank && a.DegreeOne.MatchesTarget && !a.DegreeOne.NativeMap && containsAll(a.DegreeOne.Supports, []string{SupportPiTopEqualsF1OverP, SupportDegreeTargetFlagFunctorCandidate}), Detail: FormatQuotient(a.DegreeOne)},
			{Name: "degree two targets F2/p = H_R^min", Passed: a.DegreeTwo.Degree == "Lambda^2 B_2" && a.DegreeTwo.Quotient == "F_2/p" && a.DegreeTwo.QuotientRank == HRMinRank && a.DegreeTwo.MatchesTarget && !a.DegreeTwo.NativeMap && containsAll(a.DegreeTwo.Supports, []string{SupportHRMinEqualsF2OverP, SupportDegreeTargetFlagFunctorCandidate}), Detail: FormatQuotient(a.DegreeTwo)},
			{Name: "cross-lane exclusion remains conditional on functor", Passed: a.CrossLane.ExcludedIfFunctor && !a.CrossLane.NativeExclusion && containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeDegreeToFlagFunctor}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "alpha reconstructed from flag quotients but remains sealed", Passed: a.Alpha.ReconstructedFromFlagQuotients && near(a.Alpha.ReconstructedAlpha, AlphaB) && !a.Alpha.NativeFunctor && containsAll(a.Alpha.Failures, []string{FailureNoNativeDegreeToFlagFunctor, FailureAlphaStillSealed}), Detail: FormatAlpha(a.Alpha)},
			{Name: "block R3/R4 promotion", Passed: a.R3.ConditionalTraceProxyMature && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 877 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatFlag(a.Flag), FormatQuotient(a.DegreeOne), FormatQuotient(a.DegreeTwo), FormatCrossLane(a.CrossLane), FormatAlpha(a.Alpha), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
