package generation2boundarypairexteriorjettruncationdegreetargetselectionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_PAIR_EXTERIOR_JET_TRUNCATION_DEGREE_TARGET_SELECTION_AUDIT"
	theoremName = "Gate 869 — BoundaryPair Exterior-Jet Truncation and Degree-Target Selection Audit"
)

func Generation2BoundaryPairExteriorJetTruncationDegreeTargetSelectionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 868 jet-response wound", Passed: containsAll(a.Candidate.Supports, []string{StatusGate868Inherited, SupportBoundaryJetSealSharpened}) && a.R3.BoundaryJetShapeTyped, Detail: FormatCandidate(a.Candidate)},
			{Name: "audit B2 exterior calculus and Lambda^3 truncation", Passed: a.BoundaryPair.Dimension == 2 && a.BoundaryPair.Lambda3Dim == 0 && a.BoundaryPair.TruncatesAfterSecondDegree && !a.BoundaryPair.TruncationDerivesAlphaResponse, Detail: FormatBoundaryPair(a.BoundaryPair)},
			{Name: "audit degree-one target to dominant socket", Passed: a.Targets.DegreeOne.Degree == 1 && a.Targets.DegreeOne.TargetRank == PiTopRank && a.Targets.DegreeOne.TargetChamberDim == H10Dim && !a.Targets.DegreeOne.TargetMapCertified, Detail: FormatDegree(a.Targets.DegreeOne)},
			{Name: "audit degree-two target to active right domain", Passed: a.Targets.DegreeTwo.Degree == 2 && a.Targets.DegreeTwo.TargetRank == HRminRank && a.Targets.DegreeTwo.TargetChamberDim == H72Dim && !a.Targets.DegreeTwo.TargetMapCertified, Detail: FormatDegree(a.Targets.DegreeTwo)},
			{Name: "reconstruct alpha_B only as exterior-degree candidate", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.Native && !a.Candidate.ExteriorFunctionalCertified, Detail: FormatCandidate(a.Candidate)},
			{Name: "audit zero-order and cross-lane exclusions", Passed: a.ZeroCross.ZeroOrderPresent && !a.ZeroCross.ZeroOrderContributes && a.ZeroCross.CubicAndHigherAbsent && a.ZeroCross.CubicStopDerivedByLambda3B2Zero && !a.ZeroCross.ZeroOrderSuppressionTheorem && !a.ZeroCross.CrossLaneExclusionTheorem, Detail: FormatZeroCross(a.ZeroCross)},
			{Name: "block degree-target/native alpha promotion", Passed: !a.Obstruction.NativeExteriorFunctionalCertified && !a.Obstruction.DegreeTargetSelectionCertified && !a.Obstruction.ZeroOrderSuppressionCertified && !a.Obstruction.AlphaNative, Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 869 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatBoundaryPair(a.BoundaryPair), FormatTargets(a.Targets), FormatZeroCross(a.ZeroCross), FormatCandidate(a.Candidate), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
