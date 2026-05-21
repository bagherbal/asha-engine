package generation2k7minustrialityintertwinerconstructionr4stabilizationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE955-GENERATION2K7MINUSTRIALITYINTERTWINERCONSTRUCTIONR4STABILIZATIONAUDIT"
	theoremName = "Gate 955: K7Minus Triality Intertwiner Construction and R4 GenerationCarrier Stabilization Audit"
)

func Generation2K7MinusTrialityIntertwinerConstructionR4StabilizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 955 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits Gate 954 coupling-precondition wound", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "preserves R3 dual seal and blocks native/official overclaim", Passed: a.DualSealRequired && !a.NativeR3 && !a.OfficialLedgerUpdate && !a.GenerationCarrierCertified, Detail: "R3 remains dual-sealed; generation carrier not certified"},
			{Name: "abstract C3 action passes algebraic sanity tests", Passed: a.Action.OrderThree && a.Action.Nontrivial && nearly(a.Action.Trace, 0) && nearly(a.Action.Determinant, 1) && a.Action.MetricPreserving && a.Action.OrbitSpan == 3, Detail: "order3/nontrivial/trace0/det1/metric/orbit-span"},
			{Name: "native triality restriction remains uncertified", Passed: !a.Action.CanonicalTrialityInput && !a.Action.NativeRestriction, Detail: "abstract C3 model is not a native triality restriction"},
			{Name: "R3 tracebody intertwiner remains uncertified", Passed: !a.Intertwiner.Certified && a.Intertwiner.ArbitraryBasisChoiceRequired && a.Intertwiner.UsesR3RowsAsGenerationLabels, Detail: a.Intertwiner.CandidateName},
			{Name: "noncircularity firewall preserved", Passed: !a.Intertwiner.UsesFlavorBacksolve && !a.Intertwiner.UsesObservedYukawaOrMassData && !a.Intertwiner.UsesCKMPMNSInput, Detail: "no flavor/mass/CKM/PMNS input used"},
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
		"CONDITIONAL_SUPPORT_K7_MINUS_REMAINS_STRONGEST_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ACTION_HAS_ORDER_THREE_SHAPE",
		"CONDITIONAL_SUPPORT_ABSTRACT_ACTION_PRESERVES_K7_MINUS_BILINEAR_STRUCTURE",
		"CONDITIONAL_SUPPORT_ABSTRACT_K7_MINUS_C3_ORBITS_CAN_SPAN_THREE_GENERATION_SLOTS",
		"CONDITIONAL_SUPPORT_REQUIRED_OBJECT_IS_K7_MINUS_TRIALITY_TRACEBODY_INTERTWINER",
		"CONDITIONAL_SUPPORT_GENERATION_CARRIER_CONSTRUCTION_ATTEMPT_IS_NONEMPIRICAL_AND_NONCIRCULAR",
		"CONDITIONAL_SUPPORT_R4_GENERATION_CARRIER_ATTEMPT_PRESERVES_R3_DUALSEAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_TRIALITY_DOES_NOT_CANONICALLY_RESTRICT_TO_K7_MINUS_IN_CURRENT_CERTIFICATE",
		"FAILED_ROUTE_ABSTRACT_C3_ACTION_IS_NOT_NATIVE_TRIALITY_RESTRICTION",
		"FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER",
		"FAILED_ROUTE_INTERTWINER_ONLY_EXISTS_AFTER_ARBITRARY_BASIS_CHOICE",
		"FAILED_ROUTE_INTERTWINER_USES_R3_ROWS_AS_GENERATION_LABELS",
		"FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
	}
}

func nearly(a, b float64) bool { return a-b < Tol && b-a < Tol }
