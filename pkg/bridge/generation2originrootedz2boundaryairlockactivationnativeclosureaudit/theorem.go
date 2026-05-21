package generation2originrootedz2boundaryairlockactivationnativeclosureaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE939-GENERATION2ORIGINROOTEDZ2BOUNDARYAIRLOCKACTIVATIONNATIVECLOSUREAUDIT"
	theoremName = "Gate 939: OriginRooted Z2 BoundaryAirlock Activation Native Closure Audit"
)

func Generation2OriginRootedZ2BoundaryAirlockActivationNativeClosureAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 939 origin-rooted native-closure audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate 938A native promotion blocker ledger", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "origin-rooted master functor is defined", Passed: a.MasterFunctor == MasterFunctorName && a.MasterFunctorID == MasterFunctorID, Detail: a.MasterFunctorID + " :: " + a.MasterFunctor},
			{Name: "four native gaps collapse into one certificate chain", Passed: len(a.Clauses) == 4 && allCollapsed(a.Clauses), Detail: FormatClauses(a.Clauses)},
			{Name: "full native/post-orientation R3 is not granted without all certificates", Passed: !a.FullPassEligible && !allNativeCertified(a.Clauses) && allBlockFullNative(a.Clauses), Detail: a.Truth},
			{Name: "clause failures are preserved", Passed: containsAll(append(a.Failures, clauseFailures(a.Clauses)...), clauseFailures(a.Clauses)), Detail: FormatClauseFailures(a.Clauses)},
			{Name: "R4 and official-ledger firewalls remain outside Gate 939", Passed: containsAll(a.Failures, r4Failures(a.R4Boundary)) && containsAll(a.Failures, Failures()), Detail: FormatR4(a.R4Boundary)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, MasterFunctorName, MasterFunctorID, MasterFunctorFlow, FormatDiagnostics(a.Diagnostics), FormatClauses(a.Clauses), FormatRetired(a.RetiredRoutes), FormatR4(a.R4Boundary), a.Final, NextGate}
		for _, c := range a.Clauses {
			notes = append(notes, "CLAUSE: "+c.Name, c.RequiredTheorem, c.CandidateReading)
			notes = append(notes, c.PassCondition...)
			notes = append(notes, c.Failure...)
		}
		for _, r := range a.RetiredRoutes {
			notes = append(notes, "RETIRED_FALSE_ROUTE: "+r.Name+" :: "+r.Reason)
		}
		for _, item := range a.R4Boundary {
			notes = append(notes, "R4_BOUNDARY: "+item.Name+" :: "+item.Failure)
		}
		notes = append(notes, clausePassConditions(a.Clauses)...)
		notes = append(notes, a.Statuses...)
		notes = append(notes, a.Supports...)
		notes = append(notes, a.Failures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
