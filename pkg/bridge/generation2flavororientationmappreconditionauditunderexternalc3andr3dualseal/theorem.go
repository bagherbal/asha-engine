package generation2flavororientationmappreconditionauditunderexternalc3andr3dualseal

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE961-GENERATION2FLAVORORIENTATIONMAPPRECONDITIONAUDITUNDEREXTERNALC3ANDR3DUALSEAL"
	theoremName = "Gate 961: FlavorOrientationMap Precondition Audit Under ExternalC3 and R3 DualSeal"
)

func Generation2FlavorOrientationMapPreconditionAuditUnderExternalC3AndR3DualSealTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 961 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, AuditSupports(a.Audits))
		allFailures := appendAll(a.Failures, AuditFailures(a.Audits))
		checks := []theorem.Check{
			{Name: "inherits Gate 960 sealed generation-carrier status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "types external C3 as sealed domain only", Passed: a.Decision.ExternalC3DomainAvailable && !a.Decision.ExternalC3NativeGeneration && a.Audits[0].DomainTyped, Detail: a.Domain},
			{Name: "types ledger codomain without physical assignments", Passed: a.Decision.LedgerCodomainTyped && a.Audits[1].CodomainTyped && !a.Decision.PhysicalAssignmentAllowed, Detail: a.Codomain},
			{Name: "identifies Phi_flav as required but not certified", Passed: a.Decision.FlavorOrientationMapRequired && !a.Decision.FlavorOrientationMapCertified && a.Audits[2].MapRequired && !a.Audits[2].MapCertified, Detail: a.RequiredMap},
			{Name: "forbids observed flavor and formula backsolve sources", Passed: !a.Decision.ObservedFlavorInputAllowed && !a.Decision.FlavorFormulaBacksolveAllowed && !a.Audits[3].UsesObservedFlavorData && !a.Audits[3].UsesFlavorFormulaAsSource, Detail: "flavor data are downstream targets only"},
			{Name: "preserves R3 dual seal and external C3 seal", Passed: a.Decision.R3DualSealPreserved && a.Decision.ExternalC3SealPreserved && a.Audits[4].InheritedR3DualSeal && a.Audits[4].InheritedExternalC3Seal, Detail: a.Audits[4].Object},
			{Name: "does not allow downstream spectrum claims or ledger update", Passed: !a.Decision.IndividualYukawaAllowed && !a.Decision.CKMPMNSAllowed && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "no Yukawa/CKM/PMNS/official update"},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Domain, a.Codomain, a.RequiredMap, a.Final, a.NextGate}
		notes = append(notes, AuditNotes(a.Audits)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
