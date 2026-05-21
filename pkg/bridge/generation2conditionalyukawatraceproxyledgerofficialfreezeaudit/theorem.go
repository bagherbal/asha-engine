package generation2conditionalyukawatraceproxyledgerofficialfreezeaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_CONDITIONAL_YUKAWA_TRACEPROXY_LEDGER_OFFICIAL_FREEZE_AUDIT"
	theoremName = "Gate 874 — Conditional Yukawa TraceProxy Ledger and Official-Freeze Audit"
)

func Generation2ConditionalYukawaTraceProxyLedgerOfficialFreezeAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "record conditional Yukawa trace proxy chain", Passed: a.Chain.CoherentGivenSeal && a.Chain.ConditionalYukawaLike && !a.Chain.AlphaNative && containsAll(a.Chain.Supports, []string{SupportConditionalYukawaTraceProxy, SupportFullTraceMagnitudeChain}), Detail: FormatChain(a.Chain)},
			{Name: "compute operator diagnostic ledger and separate official frozen values", Passed: a.Ledger.CYukawaMatchesThreeOverNEff && a.Ledger.OfficialFrozen && !a.Ledger.OperatorEqualsOfficialNEff && !a.Ledger.OperatorEqualsOfficialCYukawa && !a.Ledger.OperatorEqualsOfficialCHiggs, Detail: FormatLedger(a.Ledger)},
			{Name: "record R3 promotion requirements", Passed: a.PromotionRequirements.NeedLambda1ToPiTop && a.PromotionRequirements.NeedLambda2ToHRMin && a.PromotionRequirements.NeedNoLambda1ToHRMin && a.PromotionRequirements.NeedNoLambda2ToPiTop && !a.PromotionRequirements.AllRequirementsMet && !a.PromotionRequirements.EligibleForR3, Detail: FormatRequirements(a.PromotionRequirements)},
			{Name: "freeze official N_eff/C_Yukawa/C_Higgs ledger", Passed: !a.Impact.CanUpdateNEff && !a.Impact.CanUpdateCYukawa && !a.Impact.CanUpdateCHiggs && !a.Impact.CanPromoteToR3 && !a.Impact.CanPromoteToR4, Detail: FormatImpact(a.Impact)},
			{Name: "preserve Gate 874 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatChain(a.Chain), FormatLedger(a.Ledger), FormatRequirements(a.PromotionRequirements), FormatImpact(a.Impact), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
