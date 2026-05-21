package generation2k7minustrialitytracebodycouplingpreconditionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE954-GENERATION2K7MINUSTRIALITYTRACEBODYCOUPLINGPRECONDITIONAUDIT"
	theoremName = "Gate 954: K7Minus Triality Tracebody Coupling Precondition Audit"
)

func Generation2K7MinusTrialityTracebodyCouplingPreconditionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 954 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits expected predecessor status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps R3 dual seal and blocks downstream overclaim", Passed: a.DualSealRequired && !a.NativeR3 && !a.OfficialLedgerUpdate && !a.GenerationCarrierCertified && !a.FlavorOrientationCertified && !a.IndividualYukawaCertified && !a.PhysicalAssignmentCertified, Detail: "dual seal required; no generation/flavor/Yukawa theorem granted"},
			{Name: "records compatible K7-minus and triality shapes", Passed: a.Coupling.CarrierDimension == 3 && len(a.Coupling.ActionLanes) == 3 && !a.Coupling.IntertwinerCertified, Detail: a.Coupling.Carrier + " / " + a.Coupling.ActionShape},
			{Name: "records required conditional supports", Passed: containsAll(allSupports, RequiredSupports()), Detail: stringsJoin(RequiredSupports())},
			{Name: "preserves required firewalls", Passed: containsAll(allFailures, RequiredFailures()), Detail: stringsJoin(RequiredFailures())},
			{Name: "matches verdict and classification", Passed: a.Verdict == Verdict && a.Classification == Classification && a.ShortStatus == ShortStatus, Detail: a.Verdict},
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
		notes = append(notes, ItemNotes(a.Items)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func RequiredSupports() []string {
	return []string{
		"CONDITIONAL_SUPPORT_K7_MINUS_AND_TRIALITY_HAVE_COMPATIBLE_NATIVE_THREEFOLD_SHAPES",
		"CONDITIONAL_SUPPORT_K7_MINUS_SUPPLIES_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_TRIALITY_SUPPLIES_NATIVE_THREEFOLD_ACTION_SHAPE",
		"CONDITIONAL_SUPPORT_R3_TRACEBODY_CAN_BE_USED_ONLY_AS_DUALSEALED_AGGREGATE_TARGET",
		"CONDITIONAL_SUPPORT_FLAVOR_FORMULAS_REMAIN_DOWNSTREAM_LEDGER_TARGETS_ONLY",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_ACTION_MAP",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_SHAPE_COMPATIBILITY_NOT_COUPLING_THEOREM",
		"FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}
