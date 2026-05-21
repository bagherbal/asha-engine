package generation2r4generationcarriercomparisonflavorformulafirewallaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE953-GENERATION2R4GENERATIONCARRIERCOMPARISONFLAVORFORMULAFIREWALLAUDIT"
	theoremName = "Gate 953: R4 GenerationCarrier Candidate Comparison and FlavorFormula Firewall Audit"
)

func Generation2R4GenerationCarrierComparisonFlavorFormulaFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 953 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
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
		"CONDITIONAL_SUPPORT_K7_MINUS_IS_STRONGEST_FINITE_DIMENSION_THREE_CARRIER_SHAPE",
		"CONDITIONAL_SUPPORT_TRIALITY_IS_STRONGEST_ORIGIN_THREEFOLD_ACTION_CANDIDATE",
		"CONDITIONAL_SUPPORT_K7_MINUS_TRIALITY_SYNTHESIS_IS_NEXT_PRECONDITION",
	}
}

func RequiredFailures() []string {
	return []string{
		"FAILED_ROUTE_NO_GENERATION_CARRIER_MAP",
		"FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP",
		"FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES",
		"FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT",
	}
}
