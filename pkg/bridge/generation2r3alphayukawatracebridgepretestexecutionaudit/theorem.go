package generation2r3alphayukawatracebridgepretestexecutionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE937-GENERATION2R3ALPHAYUKAWATRACEBRIDGEPRETESTEXECUTIONAUDIT"
	theoremName = "Gate 937: R3 Alpha/Yukawa TraceBridge Pre-Test Execution Audit"
)

func Generation2R3AlphaYukawaTraceBridgePreTestExecutionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 937 pre-test execution audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "inherited status and classification", Passed: a.Inherited == InheritedStatus && a.Classification == Classification && a.ShortStatus == ShortStatus && a.Truth == FinalTruth, Detail: a.Inherited + " | " + a.Classification + " | " + a.ShortStatus},
			{Name: "all positive pre-test execution checks pass", Passed: positiveOK(a.Positive), Detail: FormatPositive(a.Positive[0])},
			{Name: "all negative pre-test routes are rejected", Passed: negativeOK(a.Negative), Detail: FormatNegative(a.Negative[0])},
			{Name: "numeric alpha and operator trace diagnostics reconstruct", Passed: numericOK(a.Numeric), Detail: FormatNumeric(a.Numeric)},
			{Name: "support and firewall markers are preserved", Passed: containsAll(a.Supports, Supports()) && containsAll(a.Failures, Failures()), Detail: FormatFirewalls(a.Failures)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, ReducedB2Response, AdmissibleSupportChain, ClosureOperator, ThetaFunctor, BoundaryActivationMeasure, AlphaFormula, TraceRows, NEffFormula, FormatNumeric(a.Numeric), a.Final, NextGate}
		for _, c := range a.Positive {
			notes = append(notes, FormatPositive(c), c.Marker)
		}
		for _, c := range a.Negative {
			notes = append(notes, FormatNegative(c), c.Marker)
		}
		notes = append(notes, a.Statuses...)
		notes = append(notes, a.Supports...)
		notes = append(notes, a.Failures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
