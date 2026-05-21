package generation2k7minusgenerationcarriercandidateauditunderr3dualseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE951-GENERATION2K7MINUSGENERATIONCARRIERCANDIDATEAUDITUNDERR3DUALSEAL"
	theoremName = "Gate 951: K7Minus GenerationCarrier Candidate Audit Under R3 DualSeal"
)

func Generation2K7MinusGenerationCarrierCandidateAuditUnderR3DualSealTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 951 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, ItemSupports(a.Items))
		allFailures := appendAll(a.Failures, ItemFailures(a.Items))
		checks := []theorem.Check{
			{Name: "inherits expected predecessor status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "keeps R3 dual seal and blocks native/physical overclaim", Passed: a.DualSealRequired && !a.NativeR3 && !a.OfficialLedgerUpdate && !a.GenerationCarrierCertified && !a.FlavorOrientationCertified && !a.IndividualYukawaCertified && !a.PhysicalAssignmentCertified, Detail: "dual seal required; no downstream theorem granted"},
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
		"CONDITIONAL_SUPPORT_K7_MINUS_HAS_NATIVE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_K7_HODGE_POLARITY_INHERITS_NATIVE_SPLIT_SIGNATURE_CARRIER",
		"CONDITIONAL_SUPPORT_K7_CANDIDATE_IS_NONEMPIRICAL_AND_FINITE",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_DIM_K7_MINUS_EQUALS_THREE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_K7_POLARITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_TYPED_MAP_FROM_K7_MINUS_TO_R3_TRACEBODY",
		"FAILED_ROUTE_NO_K7_MINUS_ACTION_ON_Y_DAGGER_Y_ROWS",
	}
}
