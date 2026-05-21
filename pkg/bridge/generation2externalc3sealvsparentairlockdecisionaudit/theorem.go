package generation2externalc3sealvsparentairlockdecisionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE959-GENERATION2EXTERNALC3SEALVSPARENTAIRLOCKDECISIONAUDIT"
	theoremName = "Gate 959: ExternalC3 Seal vs ParentAirlock Decision Audit"
)

func Generation2ExternalC3SealVsParentAirlockDecisionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 959 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, RouteSupports(a.Routes))
		allFailures := appendAll(a.Failures, RouteFailures(a.Routes))
		checks := []theorem.Check{
			{Name: "inherits Gate 958 negative closure", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "active board exhausted under current certificate", Passed: a.Decision.ActiveBoardExhausted && a.Routes[0].Status == RouteExhausted, Detail: a.Routes[0].Route},
			{Name: "external C3 allowed only as explicit seal", Passed: a.Decision.ExternalC3SealAllowed && a.Routes[1].Status == RouteSealAvailable && a.Routes[1].ExternalSeal && !a.Routes[1].NativeCarrier, Detail: a.Routes[1].Route},
			{Name: "parent airlock is only remaining native route", Passed: a.Decision.ParentAirlockOnlyNativeRoute && a.Routes[2].Status == RouteNativeOpenPending && a.Routes[2].ParentAirlock, Detail: a.Routes[2].Route},
			{Name: "flavor branch forbidden as source", Passed: a.Routes[3].Status == RouteForbiddenAsSource && !a.Routes[3].AllowedForFutureWork, Detail: a.Routes[3].Route},
			{Name: "does not certify native generation carrier or native R3", Passed: !a.NativeR3 && !a.GenerationCarrierCertified && !a.Decision.NativeGenerationCarrierCertified, Detail: "no native generation carrier map certified"},
			{Name: "does not allow downstream flavor/Yukawa/physical/official claims", Passed: !a.Decision.FlavorDerivationAllowed && !a.Decision.IndividualYukawaAllowed && !a.Decision.PhysicalAssignmentAllowed && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "downstream claims blocked"},
			{Name: "records required supports", Passed: containsAll(allSupports, RequiredSupports()), Detail: stringsJoin(RequiredSupports())},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Final, a.NextGate}
		notes = append(notes, RouteNotes(a.Routes)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
