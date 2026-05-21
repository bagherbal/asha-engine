package generation2boundaryexteriordegreetargetselectionmapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_MAP_AUDIT"
	theoremName = "Gate 871 — BoundaryExterior Degree-Target Selection Map Audit"
)

func Generation2BoundaryExteriorDegreeTargetSelectionMapAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 870 reduced B2 response shape", Passed: a.Response.ZeroOrderSuppressed && a.Response.HigherTruncated && !a.Response.NativeFunctionalCertified && containsAll(a.Response.Supports, []string{StatusGate870Inherited, SupportReducedB2ShapeInherited}), Detail: FormatResponse(a.Response)},
			{Name: "audit degree-one target to Pi_top", Passed: a.DegreeOne.Degree == 1 && a.DegreeOne.TargetRank == PiTopRank && a.DegreeOne.ChamberDim == H10Dim && near(a.DegreeOne.Contribution, float64(PiTopRank)/float64(H10Dim)*SBoundary) && !a.DegreeOne.MapCertified, Detail: FormatTarget(a.DegreeOne)},
			{Name: "audit degree-two target to H_R^min", Passed: a.DegreeTwo.Degree == 2 && a.DegreeTwo.TargetRank == HRminRank && a.DegreeTwo.ChamberDim == H72Dim && near(a.DegreeTwo.Contribution, float64(HRminRank)/float64(H72Dim)*SBoundary*SBoundary) && !a.DegreeTwo.MapCertified, Detail: FormatTarget(a.DegreeTwo)},
			{Name: "audit cross-lane exclusion wound", Passed: a.CrossLane.LinearToHRMinExcludedCandidate && a.CrossLane.QuadraticToPiTopExcludedCandidate && !a.CrossLane.CrossLaneExclusionTheorem && containsAll(a.CrossLane.Failures, []string{FailureNoCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "audit response chamber typing without target theorem", Passed: a.Chambers.H10Typed && a.Chambers.H72Typed && !a.Chambers.ResponseChamberTypingTheorem, Detail: FormatChambers(a.Chambers)},
			{Name: "reconstruct alpha_B from candidate degree targets", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.Native, Detail: FormatCandidate(a.Candidate)},
			{Name: "block native degree-target and alpha promotion", Passed: !a.Obstruction.DegreeTargetSelectionCertified && !a.Obstruction.CrossLaneExclusionCertified && !a.Obstruction.AlphaNative, Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 871 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatResponse(a.Response), FormatTarget(a.DegreeOne), FormatTarget(a.DegreeTwo), FormatCrossLane(a.CrossLane), FormatChambers(a.Chambers), FormatCandidate(a.Candidate), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
