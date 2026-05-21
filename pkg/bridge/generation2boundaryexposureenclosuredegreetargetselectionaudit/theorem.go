package generation2boundaryexposureenclosuredegreetargetselectionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_EXPOSURE_ENCLOSURE_DEGREE_TARGET_SELECTION_AUDIT"
	theoremName = "Gate 872 — Boundary Exposure/Enclosure Degree-Target Selection Audit"
)

func Generation2BoundaryExposureEnclosureDegreeTargetSelectionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit reduced boundary response shape", Passed: a.Response.ZeroOrderSuppressed && a.Response.HigherTruncated && !a.Response.NativeFunctionalCertified, Detail: FormatResponse(a.Response)},
			{Name: "audit degree-one exposure to Pi_top", Passed: a.Exposure.Degree == 1 && a.Exposure.Type == "single-boundary exposure" && a.Exposure.TargetRank == PiTopRank && a.Exposure.ChamberDim == H10Dim && near(a.Exposure.Contribution, float64(PiTopRank)/float64(H10Dim)*SBoundary) && !a.Exposure.MapCertified, Detail: FormatExposureEnclosure(a.Exposure)},
			{Name: "audit degree-two enclosure to H_R^min", Passed: a.Enclosure.Degree == 2 && a.Enclosure.Type == "full boundary-pair enclosure" && a.Enclosure.TargetRank == HRminRank && a.Enclosure.ChamberDim == H72Dim && near(a.Enclosure.Contribution, float64(HRminRank)/float64(H72Dim)*SBoundary*SBoundary) && !a.Enclosure.MapCertified, Detail: FormatExposureEnclosure(a.Enclosure)},
			{Name: "audit cross-lane exclusion by exposure/enclosure type", Passed: a.CrossLane.ExposureToHRMinExcludedCandidate && a.CrossLane.EnclosureToPiTopExcludedCandidate && !a.CrossLane.CrossLaneExclusionTheorem && containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "audit puncture role in enclosed active rank seven", Passed: a.Puncture.PunctureRequiredForRankSeven && a.Puncture.AmbientRightRank-a.Puncture.PunctureRank == a.Puncture.ActiveRightRank && !a.Puncture.PuncturedEnclosureTheorem, Detail: FormatPuncture(a.Puncture)},
			{Name: "reconstruct alpha_B from exposure/enclosure candidate", Passed: a.Candidate.ShapeCoherent && near(a.Candidate.ReconstructedAlpha, AlphaB) && !a.Candidate.Native, Detail: FormatCandidate(a.Candidate)},
			{Name: "block native target selection and alpha promotion", Passed: a.Obstruction.ExposureEnclosureTyped && !a.Obstruction.TargetSelectionCertified && !a.Obstruction.CrossLaneExclusionCertified && !a.Obstruction.AlphaNative, Detail: FormatObstruction(a.Obstruction)},
			{Name: "block R3/R4 and ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 872 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatResponse(a.Response), FormatExposureEnclosure(a.Exposure), FormatExposureEnclosure(a.Enclosure), FormatCrossLane(a.CrossLane), FormatPuncture(a.Puncture), FormatCandidate(a.Candidate), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
