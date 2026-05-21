package generation2noncircularssplitreplacementfinitescalarproxyaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE946-GENERATION2NONCIRCULARSSPLITREPLACEMENTFINITESCALARPROXYAUDIT"
	theoremName = "Gate 946: NonCircular S_split Replacement and FiniteScalar Proxy Audit"
)

func Generation2NonCircularSSplitReplacementFiniteScalarProxyAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 946 proxy audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		candidateNotes := []string{}
		for _, c := range a.Candidates {
			candidateNotes = append(candidateNotes, FormatCandidate(c))
		}
		criterionNotes := []string{}
		for _, c := range a.Criteria {
			criterionNotes = append(criterionNotes, FormatCriterion(c))
		}
		allFailures := appendAll(a.Failures, CandidateFailures(a.Candidates))
		allSupports := appendAll(a.Supports, CandidateSupports(a.Candidates))
		checks := []theorem.Check{
			{Name: "inherits Gate 945 addend-origin and circularity block", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps S_split target magnitude explicit", Passed: a.TargetValue == Ssplit, Detail: "target S_split"},
			{Name: "audits seven replacement candidates", Passed: len(a.Candidates) == 7, Detail: stringsJoin(CandidateNames(a.Candidates))},
			{Name: "finds no valid noncircular finite replacement", Passed: !HasSuccessfulReplacement(a.Candidates), Detail: stringsJoin(CandidateVerdicts(a.Candidates))},
			{Name: "rejects D_base as reparameterization", Passed: containsAll(allFailures, []string{"FAILED_ROUTE_D_BASE_REPLACEMENT_IS_REPARAMETERIZATION_OF_S_SPLIT"}), Detail: stringsJoin(allFailures)},
			{Name: "rejects rank ratio as coefficient not scalar", Passed: containsAll(allFailures, []string{"FAILED_ROUTE_7_OVER_72_IS_NORMALIZATION_COEFFICIENT_NOT_S_SPLIT_SCALAR"}), Detail: stringsJoin(allFailures)},
			{Name: "rejects output fixed point/inversion as circular", Passed: containsAll(allFailures, []string{"FAILED_ROUTE_FIXED_POINT_SCALAR_RECOVERY_IS_CIRCULAR_WITH_TRACEBRIDGE_OUTPUT"}), Detail: stringsJoin(allFailures)},
			{Name: "confirms scalar source seal remains", Passed: containsAll(allFailures, []string{"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT", "FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED"}), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Final, NextGate}
		notes = append(notes, criterionNotes...)
		notes = append(notes, candidateNotes...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
