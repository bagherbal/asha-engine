package generation2boundarysplitjetresponsefunctionalaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_SPLIT_JET_RESPONSE_FUNCTIONAL_AUDIT"
	theoremName = "Gate 868 — BoundarySplit Jet-Response Functional Audit"
)

func Generation2BoundarySplitJetResponseFunctionalAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 867 power-response wound", Passed: containsAll(a.Functional.Supports, []string{StatusGate867Inherited, SupportWoundToJetFunctional}) && a.R3.SocketRankSourceTyped, Detail: FormatFunctional(a.Functional)},
			{Name: "define boundary-augmented H10 and H72 chambers", Passed: a.H10.Dimension == 10 && a.H10.RankUsed == PiTopRank && a.H10.Typed && a.H72.Dimension == 72 && a.H72.RankUsed == HRminRank && a.H72.Typed, Detail: FormatChamber(a.H10) + " | " + FormatChamber(a.H72)},
			{Name: "audit first jet dominant socket lane", Passed: a.Functional.First.Power == 1 && a.Functional.First.Rank == PiTopRank && a.Functional.First.ChamberDim == H10Dim && !a.Functional.First.OperatorTyped && !a.Functional.First.NativeDerived, Detail: FormatJet(a.Functional.First)},
			{Name: "audit second jet active right-domain lane", Passed: a.Functional.Second.Power == 2 && a.Functional.Second.Rank == HRminRank && a.Functional.Second.ChamberDim == H72Dim && !a.Functional.Second.OperatorTyped && !a.Functional.Second.NativeDerived, Detail: FormatJet(a.Functional.Second)},
			{Name: "reconstruct alpha_B formally without native jet theorem", Passed: a.Functional.ShapeCoherent && near(a.Functional.ReconstructedAlpha, AlphaB) && !a.Functional.Native && !a.Functional.FirstJetCertified && !a.Functional.SecondJetCertified, Detail: FormatFunctional(a.Functional)},
			{Name: "audit shared S_split jet coordinate", Passed: a.SharedS.FeedsFirstJet && a.SharedS.FeedsSecondJet && !a.SharedS.SharedJetFunctorCertified, Detail: FormatSharedS(a.SharedS)},
			{Name: "audit no-extra-terms and truncation requirement", Passed: a.Truncation.ConstantTermAbsent && a.Truncation.CubicAndHigherAbsent && !a.Truncation.TruncationTheoremCertified, Detail: FormatTruncation(a.Truncation)},
			{Name: "block R3/R4 and ledger updates", Passed: !a.R3.EligibleForR3 && !a.R3.EligibleForR4 && !a.Impact.CanUpdateNEff && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatR3(a.R3) + " | " + FormatImpact(a.Impact)},
			{Name: "preserve Gate 868 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatChamber(a.H10), FormatChamber(a.H72), FormatFunctional(a.Functional), FormatSharedS(a.SharedS), FormatTruncation(a.Truncation), FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
