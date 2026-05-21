package generation2nativer3promotiongapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE938A-GENERATION2NATIVER3PROMOTIONGAPAUDIT"
	theoremName = "Gate 938A: Native R3 Promotion Gap Audit"
)

func Generation2NativeR3PromotionGapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 938A native promotion gap audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate 937 bridge pre-test pass", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "classification prevents false native promotion", Passed: a.Classification == Classification && a.ShortStatus == ShortStatus && a.Truth == FinalTruth, Detail: a.Classification + " | " + a.ShortStatus},
			{Name: "four native R3 blockers are explicit", Passed: len(a.Blockers) == 4 && allPrimary(a.Blockers), Detail: FormatBlockers(a.Blockers)},
			{Name: "former phase/representative false routes are retired from primary status", Passed: len(a.RetiredWounds) >= 8, Detail: FormatRetired(a.RetiredWounds)},
			{Name: "R4 generation/flavor boundary remains separate", Passed: len(a.R4Boundary) >= 5 && containsAll(a.Failures, r4Failures(a.R4Boundary)), Detail: FormatR4(a.R4Boundary)},
			{Name: "support and firewall markers are preserved", Passed: containsAll(a.Supports, Supports()) && containsAll(a.Failures, Failures()) && containsAll(a.Failures, blockerFailures(a.Blockers)), Detail: FormatFirewalls(a.Failures)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, BridgeFormula, TraceRows, TraceFormula, FormatDiagnostics(a.Diagnostics), FormatBlockers(a.Blockers), FormatRetired(a.RetiredWounds), FormatR4(a.R4Boundary), a.Final, NextGate}
		for _, b := range a.Blockers {
			notes = append(notes, b.Name, b.RequiredGate, b.CurrentSupport, b.Failure)
		}
		for _, r := range a.RetiredWounds {
			notes = append(notes, "RETIRED_PRIMARY_BLOCKER: "+r.Name+" :: "+r.Reason)
		}
		for _, item := range a.R4Boundary {
			notes = append(notes, "R4_BOUNDARY: "+item.Name+" :: "+item.Failure)
		}
		notes = append(notes, a.Statuses...)
		notes = append(notes, a.Supports...)
		notes = append(notes, a.Failures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
