package generation2scalarsourcesealclassificationr3bridgetheoremboundaryaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE947-GENERATION2SCALARSOURCESEALCLASSIFICATIONR3BRIDGETHEOREMBOUNDARYAUDIT"
	theoremName = "Gate 947: ScalarSourceSeal Classification and R3 Bridge-Theorem Boundary Audit"
)

func Generation2ScalarSourceSealClassificationR3BridgeTheoremBoundaryAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 947 scalar-source seal boundary audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits Gate 946 scalar-source seal", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps validated bridge diagnostics explicit", Passed: a.Ssplit == Ssplit && a.AlphaB == AlphaB && a.NEff == NEff && a.CYukawa == CYukawa && a.CHiggs == CHiggs, Detail: "validated alpha/N_eff/C diagnostics"},
			{Name: "classifies bridge theorem candidate as test-passed and closure-factored", Passed: a.Boundary.BridgeTestPassed && a.Boundary.ClosureFactored && a.Boundary.MeasureFactored, Detail: FormatBoundary(a.Boundary)},
			{Name: "keeps scalar source sealed", Passed: a.Boundary.ScalarSourceSealed && containsAll(allFailures, []string{"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT", "FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND"}), Detail: stringsJoin(allFailures)},
			{Name: "blocks native R3 and official ledger update", Passed: !a.Boundary.NativeR3 && !a.Boundary.OfficialLedgerUpdate && containsAll(allFailures, []string{"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED", "FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED"}), Detail: FormatBoundary(a.Boundary)},
			{Name: "keeps R4/yukawa spectrum out of scope", Passed: containsAll(allFailures, []string{"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES", "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP", "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM"}), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatBoundary(a.Boundary), a.Final, NextGateA, NextGateB}
		notes = append(notes, ItemNotes(a.Items)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
