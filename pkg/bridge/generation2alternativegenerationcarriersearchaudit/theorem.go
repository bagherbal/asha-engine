package generation2alternativegenerationcarriersearchaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GEN2-GATE958-GENERATION2ALTERNATIVEGENERATIONCARRIERSEARCHAUDIT"
	theoremName = "Gate 958: Alternative GenerationCarrier Search Audit"
)

func Generation2AlternativeGenerationCarrierSearchAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 958 audit", Passed: false, Detail: err.Error()}}, Notes: []string{Verdict, Classification, ShortStatus}}
		}
		allSupports := appendAll(a.Supports, CandidateSupports(a.Candidates))
		allFailures := appendAll(a.Failures, CandidateFailures(a.Candidates))
		checks := []theorem.Check{
			{Name: "inherits Gate 957 route bifurcation", Passed: a.Inherited == InheritedStatus, Detail: a.Inherited},
			{Name: "preserves R3 dual seal and avoids native R3 overclaim", Passed: a.R3DualSealRequired && !a.NativeR3, Detail: "dual seal preserved"},
			{Name: "does not certify generation/flavor/Yukawa/official status", Passed: !a.GenerationCarrierCertified && !a.Decision.NativeGenerationCarrier && !a.Decision.AlternativeCandidateFound && !a.Decision.FlavorOrientationCertified && !a.Decision.IndividualYukawaCertified && !a.Decision.PhysicalAssignmentCertified && !a.Decision.OfficialLedgerUpdateAllowed, Detail: "no generation carrier map certified"},
			{Name: "K7 Hodge polarity remains shape-only", Passed: a.Candidates[0].Status == CandidateShapeOnly && a.Candidates[0].CarrierShape && !a.Candidates[0].NativeActionOrSelector && !a.Candidates[0].TypedTracebodyMap, Detail: a.Candidates[0].Candidate},
			{Name: "Fock/projective selector rejected as generation carrier", Passed: a.Candidates[1].Status == CandidateRejected && !a.Candidates[1].TypedTracebodyMap, Detail: a.Candidates[1].Candidate},
			{Name: "B2/R3 trace response rejected as generation labeling", Passed: a.Candidates[2].Status == CandidateRejected && !a.Candidates[2].UsesR3RowsAsLabels, Detail: a.Candidates[2].Candidate},
			{Name: "Boolean-octonionic alternative search found no canonical three-carrier", Passed: a.Candidates[3].Status == CandidateNoThree && !a.Candidates[3].CarrierShape && !a.Candidates[3].TypedTracebodyMap, Detail: a.Candidates[3].Candidate},
			{Name: "external C3 is seal only", Passed: a.Candidates[4].Status == CandidateSealOnly && a.Candidates[4].AllowsOnlyAsSeal, Detail: a.Candidates[4].Candidate},
			{Name: "route decision requires external seal or new parent airlock", Passed: a.Decision.Decision == Verdict && a.Decision.ExternalSealAvailable && a.Decision.ParentAirlockRequired, Detail: a.Decision.Decision},
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
		notes = append(notes, CandidateNotes(a.Candidates)...)
		notes = append(notes, allSupports...)
		notes = append(notes, allFailures...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
