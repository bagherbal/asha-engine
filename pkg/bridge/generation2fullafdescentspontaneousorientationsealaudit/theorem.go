package generation2fullafdescentspontaneousorientationsealaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE948-GENERATION2FULLAFDESCENTSPONTANEOUSORIENTATIONSEALAUDIT"
	theoremName = "Gate 948: Full A_F Descent vs SpontaneousOrientation Seal Audit"
)

func Generation2FullAFDescentSpontaneousOrientationSealAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 948 full-AF orientation seal audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits scalar-source-sealed tracebridge", Passed: a.Inherited == InheritedStatus && a.Boundary.ScalarSourceSealed, Detail: a.Inherited},
			{Name: "keeps validated diagnostics diagnostic only", Passed: a.Ssplit == Ssplit && a.AlphaB == AlphaB && a.NEff == NEff && a.CYukawa == CYukawa && a.CHiggs == CHiggs, Detail: "validated bridge diagnostics retained"},
			{Name: "blocks full A_F descent", Passed: a.Boundary.FullAFDescentBlocked && containsAll(allFailures, []string{"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED", "FAILED_ROUTE_SOCKET_PROJECTORS_NOT_STABLE_UNDER_FULL_H_ACTION"}), Detail: stringsJoin(allFailures)},
			{Name: "supports post-orientation stabilizer layer", Passed: a.Boundary.StableInAFOrient && containsAll(allSupports, []string{"CONDITIONAL_SUPPORT_TRACEBRIDGE_STABLE_IN_A_F_ORIENT_LAYER", "CONDITIONAL_SUPPORT_A_F_ORIENT_IS_POST_ORIENTATION_STABILIZER"}), Detail: stringsJoin(allSupports)},
			{Name: "keeps spontaneous orientation sealed rather than native", Passed: a.Boundary.SpontaneousOrientSeal && !a.Boundary.NativeOrientation && containsAll(allFailures, []string{"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM", "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"}), Detail: FormatBoundary(a.Boundary)},
			{Name: "blocks native R3 and downstream claims", Passed: !a.Boundary.NativeR3 && !a.Boundary.OfficialLedgerUpdate && !a.Boundary.PhysicalAssignment && !a.Boundary.GenerationFlavorClaims && containsAll(allFailures, []string{"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED", "FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED", "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM"}), Detail: stringsJoin(allFailures)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatBoundary(a.Boundary), a.Final, NextGate}
		notes = append(notes, ItemNotes(a.Items)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
