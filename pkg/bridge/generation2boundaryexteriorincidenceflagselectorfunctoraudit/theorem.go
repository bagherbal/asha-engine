package generation2boundaryexteriorincidenceflagselectorfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_SELECTOR_FUNCTOR_AUDIT"
	theoremName = "Gate 879 — BoundaryExterior IncidenceFlag Selector Functor Audit"
)

func Generation2BoundaryExteriorIncidenceFlagSelectorFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 878 and official freeze", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdate && !a.Impact.CanUpdateNEff && a.Impact.IncidenceSelectorCandidate, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "audit source incidence degrees", Passed: a.Source.DegreeOneBeforeDegreeTwo && a.Source.Lambda3Zero && a.Source.SelectorIndexNotGenerator && containsAll(a.Source.Supports, []string{SupportExteriorDegreeIncidenceShape, SupportSourcePosetOneLessTwo}) && containsAll(a.Source.Failures, []string{FailureNotLinearMap, FailureNoNativeIncidenceFunctor}), Detail: FormatSource(a.Source)},
			{Name: "validate puncture-complement target flag", Passed: a.Flag.Nested && a.Flag.QuotientsValid && a.Flag.RankF0 == 1 && a.Flag.RankF1 == 4 && a.Flag.RankF2 == 8 && a.Flag.F1OverF0Rank == 3 && a.Flag.F2OverF0Rank == 7, Detail: FormatFlag(a.Flag)},
			{Name: "degree one incidence selects F1/F0", Passed: a.DegreeOne.SelectorMode && !a.DegreeOne.LinearSurjection && !a.DegreeOne.NativeFunctor && a.DegreeOne.SelectedQuotient == "F_1/F_0" && a.DegreeOne.Target == "Pi_top" && a.DegreeOne.QuotientRank == 3 && containsAll(a.DegreeOne.Failures, []string{FailureNoNativeDegreeOneIncidence, FailureNotLinearMap}), Detail: FormatSelector(a.DegreeOne)},
			{Name: "degree two incidence selects F2/F0", Passed: a.DegreeTwo.SelectorMode && !a.DegreeTwo.LinearSurjection && !a.DegreeTwo.NativeFunctor && a.DegreeTwo.SelectedQuotient == "F_2/F_0" && a.DegreeTwo.Target == "H_R^min" && a.DegreeTwo.QuotientRank == 7 && containsAll(a.DegreeTwo.Failures, []string{FailureNoNativeDegreeTwoIncidence, FailureNotLinearMap}), Detail: FormatSelector(a.DegreeTwo)},
			{Name: "cross-lane exclusion remains conditional", Passed: a.CrossLane.ExcludedIfFunctor && !a.CrossLane.NativeExclusion && containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeIncidenceFunctor}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "alpha reconstructed by incidence selector but remains sealed", Passed: a.Alpha.ReconstructedByIncidence && near(a.Alpha.ReconstructedAlpha, AlphaB) && !a.Alpha.NativeIncidenceFunctor && containsAll(a.Alpha.Failures, []string{FailureNoNativeIncidenceFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}), Detail: FormatAlpha(a.Alpha)},
			{Name: "block R3/R4 and ledger updates", Passed: a.R3.ConditionalTraceProxyMature && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanPromoteToR3 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 879 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatSource(a.Source), FormatFlag(a.Flag), FormatSelector(a.DegreeOne), FormatSelector(a.DegreeTwo), FormatCrossLane(a.CrossLane), FormatAlpha(a.Alpha), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
