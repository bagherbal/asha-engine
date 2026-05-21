package generation2reducedboundarypairexteriorresponsedegreetargetselectionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_REDUCED_BOUNDARY_PAIR_EXTERIOR_RESPONSE_DEGREE_TARGET_SELECTION_AUDIT"
	theoremName = "Gate 870 — Reduced BoundaryPair Exterior Response and Degree-Target Selection Audit"
)

func Generation2ReducedBoundaryPairExteriorResponseDegreeTargetSelectionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 869 exterior truncation wound", Passed: containsAll(a.Candidate.Supports, []string{StatusGate869Inherited, SupportAlphaShapeSharpened}) && a.R3.SocketRankAlphaSourceTyped, Detail: FormatCandidate(a.Candidate)},
			{Name: "define reduced B2 exterior response", Passed: a.Response.Lambda0Removed && a.Response.DegreeOnePresent && a.Response.DegreeTwoPresent && a.Response.CubicAndHigherVanish && a.Response.ExactShapeCandidate && !a.Response.NativeFunctionalCertified, Detail: FormatResponse(a.Response)},
			{Name: "suppress zero-order and cubic-plus terms at shape level", Passed: a.CrossLane.ZeroOrderSuppressed && a.CrossLane.CubicAndHigherSuppressed && !a.CrossLane.NativeReducedFunctional, Detail: FormatCrossLane(a.CrossLane)},
			{Name: "audit degree-one target to Pi_top", Passed: a.Targets.DegreeOne.Degree == 1 && a.Targets.DegreeOne.TargetRank == PiTopRank && a.Targets.DegreeOne.ChamberDim == H10Dim && !a.Targets.DegreeOne.TargetMapCertified, Detail: FormatDegree(a.Targets.DegreeOne)},
			{Name: "audit degree-two target to H_R^min", Passed: a.Targets.DegreeTwo.Degree == 2 && a.Targets.DegreeTwo.TargetRank == HRminRank && a.Targets.DegreeTwo.ChamberDim == H72Dim && !a.Targets.DegreeTwo.TargetMapCertified, Detail: FormatDegree(a.Targets.DegreeTwo)},
			{Name: "formally reconstruct alpha_B from reduced response candidate", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.Native, Detail: FormatCandidate(a.Candidate)},
			{Name: "block target-selection and native response promotion", Passed: !a.Obstruction.NativeReducedFunctionalCertified && !a.Obstruction.DegreeTargetSelectionCertified && !a.Obstruction.AlphaNative, Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 870 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatResponse(a.Response), FormatTargets(a.Targets), FormatCrossLane(a.CrossLane), FormatCandidate(a.Candidate), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
