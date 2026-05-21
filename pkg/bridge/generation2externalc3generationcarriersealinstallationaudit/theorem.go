package generation2externalc3generationcarriersealinstallationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE960-GENERATION2EXTERNALC3GENERATIONCARRIERSEALINSTALLATIONAUDIT"
	theoremName = "Gate 960: ExternalC3 GenerationCarrier Seal Installation Audit"
)

func Generation2ExternalC3GenerationCarrierSealInstallationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 960 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, AuditSupports(a.Audits))
		allFailures := appendAll(a.Failures, AuditFailures(a.Audits))
		checks := []theorem.Check{
			{Name: "inherits Gate 959 fork decision", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "external C3 seal installed", Passed: a.Decision.ExternalC3SealInstalled && a.GenerationCarrierAvailable && a.Audits[0].Installed, Detail: a.ExternalCarrier},
			{Name: "denies native generation multiplicity theorem", Passed: !a.NativeMultiplicityTheorem && !a.Decision.NativeGenerationCarrier && !a.Audits[0].NativeGenerationTheorem, Detail: "external C3 is a seal only"},
			{Name: "preserves R3 dual seal", Passed: a.Decision.R3DualSealPreserved && a.Audits[1].CompatibleWithR3DualSeal, Detail: a.Audits[1].Object},
			{Name: "does not provide flavor orientation or individual spectrum", Passed: !a.Decision.FlavorOrientationMapAvailable && !a.Decision.IndividualYukawaAllowed && !a.Decision.PhysicalAssignmentAllowed && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "downstream R4 flavor remains blocked"},
			{Name: "parent airlock remains separate native route", Passed: a.Decision.ParentAirlockStillOpen && a.Audits[3].Status == SealForbiddenPromotion, Detail: a.Audits[3].Object},
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
		notes = append(notes, AuditNotes(a.Audits)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
