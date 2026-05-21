package generation2trialitygenerationcarriercandidateauditunderr3dualseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE952-GENERATION2TRIALITYGENERATIONCARRIERCANDIDATEAUDITUNDERR3DUALSEAL"
	theoremName = "Gate 952: Triality GenerationCarrier Candidate Audit Under R3 DualSeal"
)

func Generation2TrialityGenerationCarrierCandidateAuditUnderR3DualSealTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 952 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
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
		"CONDITIONAL_SUPPORT_TRIALITY_IS_DEEP_NATIVE_THREEFOLD_SOURCE_CANDIDATE",
		"CONDITIONAL_SUPPORT_CL17_ROOT_MAKES_TRIALITY_A_LAWFUL_R4_SOURCE_TO_AUDIT",
		"CONDITIONAL_SUPPORT_VECTOR_SPINOR_SPINOR_THREEFOLD_IS_STRUCTURAL_NOT_EMPIRICAL",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_TRIALITY_THREEFOLD_SHAPE_NOT_GENERATION_THEOREM",
		"FAILED_ROUTE_NO_TRIALITY_TO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_TRIALITY_ACTION_ON_R3_TRACEBODY",
		"FAILED_ROUTE_NO_TRIALITY_TO_YUKAWA_SOCKET_LEDGER_FUNCTOR",
	}
}
