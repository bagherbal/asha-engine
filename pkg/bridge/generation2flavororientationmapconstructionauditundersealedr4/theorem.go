package generation2flavororientationmapconstructionauditundersealedr4

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE962-GENERATION2FLAVORORIENTATIONMAPCONSTRUCTIONAUDITUNDERSEALEDR4"
	theoremName = "Gate 962: FlavorOrientationMap Construction Audit Under Sealed R4"
)

func Generation2FlavorOrientationMapConstructionAuditUnderSealedR4Theorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 962 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, CandidateSupports(a.Candidates))
		allFailures := appendAll(a.Failures, CandidateFailures(a.Candidates))
		checks := []theorem.Check{
			{Name: "inherits Gate 961 missing flavor-orientation status", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "exposes U3 family-gauge orbit without canonical basis", Passed: a.Decision.U3FamilyGaugeFreedomDetected && a.Decision.U3OrbitClassAvailable && !a.Decision.CanonicalFlavorBasisCertified && !a.Decision.CanonicalRepresentativeSelected, Detail: a.Domain},
			{Name: "keeps Phi_flav construction uncertified", Passed: !a.Decision.FlavorOrientationMapCertified && a.MissingObject != "", Detail: a.MissingObject},
			{Name: "allows A_F^orient as interface but not family selector", Passed: a.Decision.AFOrientIsValidInterfaceTarget && !a.Decision.AFOrientSuppliesFamilySelector, Detail: a.Target},
			{Name: "keeps R3 tracebody aggregate-only", Passed: a.Decision.R3TracebodyAggregateTargetValid && !a.Decision.R3RowsUsedAsGenerationLabels, Detail: "R3 trace rows are not generation labels"},
			{Name: "rejects flavor formulas and observed data as orientation sources", Passed: !a.Decision.FlavorFormulaBacksolveAllowed && !a.Decision.ObservedFlavorInputAllowed, Detail: "flavor data remain downstream targets only"},
			{Name: "does not assign particles, Yukawas, CKM/PMNS, or ledger updates", Passed: !a.Decision.PhysicalAssignmentAllowed && !a.Decision.IndividualYukawaAllowed && !a.Decision.CKMPMNSAllowed && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "no downstream flavor claims"},
			{Name: "preserves R3 dual seal and ExternalC3 seal", Passed: a.Decision.R3DualSealPreserved && a.Decision.ExternalC3SealPreserved, Detail: "seals remain inherited"},
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
		notes := []string{a.Verdict, a.Classification, a.ShortStatus, a.Inherited, a.Domain, a.Target, a.MissingObject, a.Final, a.NextGate}
		notes = append(notes, CandidateNotes(a.Candidates)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
