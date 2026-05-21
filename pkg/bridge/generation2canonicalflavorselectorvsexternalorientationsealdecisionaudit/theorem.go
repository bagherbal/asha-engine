package generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE963-GENERATION2CANONICALFLAVORSELECTORVSEXTERNALORIENTATIONSEALDECISIONAUDIT"
	theoremName = "Gate 963: CanonicalFlavorSelector vs ExternalFlavorOrientationSeal Decision Audit"
)

func Generation2CanonicalFlavorSelectorVsExternalOrientationSealDecisionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 963 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, CandidateSupports(a.Candidates))
		allFailures := appendAll(a.Failures, CandidateFailures(a.Candidates))
		checks := []theorem.Check{
			{Name: "inherits Gate 962 U3 orbit obstruction", Passed: a.Inherited == InheritedStatus && a.Decision.InheritedU3OrbitOnly, Detail: a.Inherited},
			{Name: "no canonical flavor selector found", Passed: !a.Decision.CanonicalFlavorSelectorFound && !a.Decision.CanonicalRepresentativeSelected, Detail: a.Problem},
			{Name: "U3 family-gauge freedom remains unbroken", Passed: a.Decision.U3FamilyGaugeFreedomRetained && !a.Decision.CurrentASHADataBreaksU3Gauge, Detail: "C3 orbit has no selected representative"},
			{Name: "external flavor-orientation seal required but non-native", Passed: a.Decision.ExternalFlavorOrientationSealRequired && a.Decision.ExternalFlavorOrientationSealCanSelectRepresentative && !a.Decision.ExternalFlavorOrientationSealNative, Detail: a.RequiredSeal},
			{Name: "downstream flavor-ledger tests only under seal", Passed: a.Decision.DownstreamFlavorLedgerTestsAllowedUnderSeal, Detail: "sealed downstream diagnostics permitted; native claims blocked"},
			{Name: "does not derive flavor theorem, Yukawas, CKM/PMNS, particles, or ledger updates", Passed: !a.Decision.NativeFlavorTheoremCertified && !a.Decision.YukawaEigenvaluesDerived && !a.Decision.CKMPMNSDerived && !a.Decision.PhysicalParticlesAssigned && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "no downstream overclaim"},
			{Name: "preserves inherited R3DualSeal and ExternalC3 seal", Passed: a.Decision.R3DualSealPreserved && a.Decision.ExternalGenerationCarrierSealPreserved, Detail: "seals remain visible"},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Problem, a.RequiredSeal, a.Final, a.NextGate}
		notes = append(notes, CandidateNotes(a.Candidates)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
