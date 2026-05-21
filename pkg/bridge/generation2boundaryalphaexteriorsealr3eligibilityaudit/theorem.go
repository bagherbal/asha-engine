package generation2boundaryalphaexteriorsealr3eligibilityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_BOUNDARY_ALPHA_EXTERIOR_SEAL_R3_ELIGIBILITY_AUDIT"
	theoremName = "Gate 873 — BoundaryAlpha ExteriorSeal and R3 Eligibility Audit"
)

func Generation2BoundaryAlphaExteriorSealR3EligibilityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "classify BoundaryAlpha exterior seal", Passed: a.AlphaSeal.ShapeTyped && a.AlphaSeal.RankSourcesTyped && !a.AlphaSeal.Native && containsAll(a.AlphaSeal.Supports, []string{SupportBoundaryAlphaExteriorSeal, SupportAlphaSocketRankAnatomy}), Detail: FormatAlphaSeal(a.AlphaSeal)},
			{Name: "reassemble conditional trace-magnitude chain", Passed: a.Chain.YDagYReadoutReady && a.Chain.SocketMagnitudesTypedGivenAlpha && a.Chain.HaggReconstructedGivenAlpha && a.Chain.CoherentGivenAlphaSeal && !a.Chain.AlphaNative, Detail: FormatChain(a.Chain)},
			{Name: "compute conditional operator N_eff and preserve frozen ledger", Passed: a.Readout.Conditional && a.Readout.NEffMatchesGate829 && !a.Readout.OfficialEqualsOperator, Detail: FormatReadout(a.Readout)},
			{Name: "audit R3 eligibility", Passed: a.R3.EligibleForConditionalR3Candidate && !a.R3.EligibleForOfficialR3 && !a.R3.EligibleForR4 && !a.R3.AlphaNative && !a.R3.TargetSelectionNative, Detail: FormatR3(a.R3)},
			{Name: "block official ledger updates and native promotion", Passed: !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 873 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatAlphaSeal(a.AlphaSeal), FormatChain(a.Chain), FormatReadout(a.Readout), FormatR3(a.R3), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
