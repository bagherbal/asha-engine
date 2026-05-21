package generation2nestedpuncturecomplementtargetselectionfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NESTED_PUNCTURE_COMPLEMENT_TARGET_SELECTION_FUNCTOR_AUDIT"
	theoremName = "Gate 876 — Nested Puncture-Complement TargetSelection Functor Audit"
)

func Generation2NestedPunctureComplementTargetSelectionFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 875 wound and ledger freeze", Passed: a.Ledger.OfficialFrozen && !a.Ledger.CanUpdateOfficial && a.Impact.NestedComplementStrongest, Detail: FormatLedger(a.Ledger) + " | " + FormatImpact(a.Impact)},
			{Name: "define neutral puncture", Passed: a.Puncture.Rank == 1 && a.Puncture.Name == "p=e_+ tensor P_1", Detail: FormatPuncture(a.Puncture)},
			{Name: "degree-one exposed-face complement reconstructs Pi_top", Passed: a.DegreeOne.Degree == "Lambda^1 B_2" && a.DegreeOne.AmbientRank == 4 && a.DegreeOne.ComplementRank == PiTopRank && a.DegreeOne.MatchesTarget && !a.DegreeOne.NativeMap && containsAll(a.DegreeOne.Supports, []string{SupportLambda1ExposedFaceComplement, SupportPiTopRankFromComplement}), Detail: FormatComplement(a.DegreeOne)},
			{Name: "degree-two full-rectangle complement reconstructs H_R^min", Passed: a.DegreeTwo.Degree == "Lambda^2 B_2" && a.DegreeTwo.AmbientRank == 8 && a.DegreeTwo.ComplementRank == HRMinRank && a.DegreeTwo.MatchesTarget && !a.DegreeTwo.NativeMap && containsAll(a.DegreeTwo.Supports, []string{SupportLambda2FullRectangleComplement, SupportHRMinRankFromComplement}), Detail: FormatComplement(a.DegreeTwo)},
			{Name: "face/enclosure cross-lane exclusion remains candidate only", Passed: a.CrossLane.TypeCandidate && !a.CrossLane.NativeExclusion && containsAll(a.CrossLane.Failures, []string{FailureNoFaceVsEnclosureDegreeTheorem, FailureNoNativeCrossLaneExclusionTheorem}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "alpha reconstructed but remains sealed", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.NativeFunctor && containsAll(a.Candidate.Failures, []string{FailureNestedComplementNotNativeFunctor, FailureAlphaStillSealed}), Detail: FormatCandidate(a.Candidate)},
			{Name: "block R3/R4 and official updates", Passed: a.R3.ConditionalTraceProxyMature && !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanPromoteToR3 && !a.Impact.CanUpdateNEff, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 876 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatPuncture(a.Puncture), FormatComplement(a.DegreeOne), FormatComplement(a.DegreeTwo), FormatCrossLane(a.CrossLane), FormatCandidate(a.Candidate), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
