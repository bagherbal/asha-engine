package generation2closurefactoredboundaryactivationmeasureconsolidationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE934-GENERATION2CLOSUREFACTOREDBOUNDARYACTIVATIONMEASURECONSOLIDATIONAUDIT"
	theoremName = "Gate 934: ClosureFactored BoundaryActivationMeasure Consolidation Audit"
)

func Generation2ClosureFactoredBoundaryActivationMeasureConsolidationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 934 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "inherited status and classification", Passed: a.Inherited == InheritedStatus && a.Classification == Classification && a.ShortStatus == ShortStatus && a.Truth == FinalTruth, Detail: a.Inherited + " | " + a.Classification + " | " + a.ShortStatus},
			{Name: "gate components pass with support markers", Passed: componentsOK(a.Components) && containsAll(a.Supports, Supports()), Detail: FormatComponent(a.Components[0])},
			{Name: "numeric alpha and trace diagnostics are stable", Passed: numericOK(a.Numeric), Detail: FormatNumeric(a.Numeric)},
			{Name: "native firewalls are preserved", Passed: containsAll(a.Failures, Failures()), Detail: FormatFirewalls(a.Failures)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, F0, F1, F2, AdmissibleSupportChain, Z2SupportLattice, ClosureOperator, ThetaFunctor, ReducedB2Response, BoundaryActivationMeasure, AlphaFormula, TraceRows, NEffFormula, FormatNumeric(a.Numeric), a.Final, NextGate}
		for _, c := range a.Components {
			notes = append(notes, FormatComponent(c), c.Marker)
		}
		notes = append(notes, a.Statuses...)
		notes = append(notes, a.Supports...)
		notes = append(notes, a.Failures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
