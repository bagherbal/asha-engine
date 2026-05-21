package generation2r3dualsealbridgetheoremfinalboundaryr4preconditionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE949-GENERATION2R3DUALSEALBRIDGETHEOREFINALBOUNDARYR4PRECONDITIONAUDIT"
	theoremName = "Gate 949: R3 DualSeal Bridge-Theorem Final Boundary and R4 Precondition Audit"
)

func Generation2R3DualSealBridgeTheoremFinalBoundaryR4PreconditionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 949 dual-seal boundary audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits dual-sealed non-native tracebridge", Passed: a.Inherited == InheritedStatus && a.Boundary.TracebridgeTestPassed && a.Boundary.ScalarSourceSeal && a.Boundary.PostOrientationSeal, Detail: a.Inherited},
			{Name: "keeps validated diagnostics diagnostic only", Passed: a.Ssplit == Ssplit && a.AlphaB == AlphaB && a.NEff == NEff && a.CYukawa == CYukawa && a.CHiggs == CHiggs && !a.Boundary.OfficialLedgerUpdate, Detail: "diagnostic R3 bridge values retained without ledger update"},
			{Name: "freezes scalar-source seal", Passed: containsAll(allSupports, []string{"CONDITIONAL_SUPPORT_SCALAR_SOURCE_SEAL_FINALIZED"}) && containsAll(allFailures, []string{"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT", "FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND"}), Detail: stringsJoin(allFailures)},
			{Name: "freezes post-orientation seal", Passed: containsAll(allSupports, []string{"CONDITIONAL_SUPPORT_POST_ORIENTATION_SEAL_FINALIZED"}) && containsAll(allFailures, []string{"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED", "FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM"}), Detail: stringsJoin(allFailures)},
			{Name: "allows R4 only under explicit dual seal", Passed: a.Policy.MayProceedUnderSeal && a.Policy.RequiresScalarSourceSeal && a.Policy.RequiresPostOrientSeal && a.Policy.AllowsGenerationAudit && a.Policy.AllowsFlavorAudit && a.Policy.AllowsYukawaPrecondition && containsAll(allSupports, []string{"CONDITIONAL_SUPPORT_R4_WORK_MAY_PROCEED_ONLY_UNDER_EXPLICIT_DUAL_SEAL"}), Detail: FormatPolicy(a.Policy)},
			{Name: "blocks native R3, official ledgers, physical spectrum claims", Passed: !a.Boundary.NativeR3 && !a.Boundary.OfficialLedgerUpdate && !a.Boundary.PhysicalAssignment && !a.Boundary.GenerationCarrier && !a.Boundary.FlavorOrientation && !a.Boundary.IndividualYukawa && !a.Boundary.R4NativeSpectrum && containsAll(allFailures, []string{"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED", "FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED", "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT", "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM"}), Detail: FormatBoundary(a.Boundary)},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, FormatBoundary(a.Boundary), FormatPolicy(a.Policy), a.Final, NextGate}
		notes = append(notes, ItemNotes(a.Items)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
