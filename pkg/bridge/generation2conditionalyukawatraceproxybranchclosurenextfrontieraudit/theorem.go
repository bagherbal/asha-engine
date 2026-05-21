package generation2conditionalyukawatraceproxybranchclosurenextfrontieraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_CONDITIONAL_YUKAWA_TRACEPROXY_BRANCH_CLOSURE_NEXT_FRONTIER_AUDIT"
	theoremName = "Gate 881 — Conditional Yukawa TraceProxy Branch Closure and Next-Frontier Audit"
)

func Generation2ConditionalYukawaTraceProxyBranchClosureNextFrontierAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "close mature conditional trace proxy branch", Passed: a.Closure.BranchClosed && a.Closure.ConditionalProxyMature && a.Closure.Status == R2Status && !a.Closure.EligibleForNativeR3 && !a.Closure.EligibleForR4 && a.Closure.NextFrontier == RecommendedNextFrontier && containsAll(a.Closure.Supports, []string{SupportConditionalTraceProxyAchieved, SupportMatureR2Closure, SupportFrontierBSafest}), Detail: FormatClosure(a.Closure)},
			{Name: "record conditional diagnostic ledger and official freeze", Passed: a.Ledger.OfficialFrozen && a.Ledger.DiagnosticOnly && !a.Ledger.CanUpdate && near(a.Ledger.OperatorNEff, OperatorNEffDiagnostic) && near(a.Ledger.OfficialNEff, OfficialNEffFrozen) && containsAll(a.Ledger.Failures, []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}), Detail: FormatLedger(a.Ledger)},
			{Name: "inherit BoundaryAlpha incidence-flag seal", Passed: a.Alpha.Name == SealName && a.Alpha.FullName == FullSealName && a.Alpha.ReducedExteriorResponse && a.Alpha.IncidenceFlagSelector && !a.Alpha.NativeFunctor && near(a.Alpha.Alpha, AlphaB) && containsAll(a.Alpha.Failures, []string{FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion}), Detail: FormatAlpha(a.Alpha)},
			{Name: "reassemble closed conditional chain", Passed: a.Chain.CoherentGivenSeal && a.Chain.ReducedExteriorToAlpha && a.Chain.AlphaToSocketMagnitudes && a.Chain.SocketMagnitudesToYDagY && a.Chain.YDagYToHAgg && a.Chain.HAggToNEff && a.Chain.NEffToCYukawaProxy && containsAll(a.Chain.Failures, []string{FailureAlphaStillSealed, FailureConditionalProxyNotR3}), Detail: FormatChain(a.Chain)},
			{Name: "file exact native wall", Passed: a.Wall.Name == MissingNativeTheoremName && a.Wall.RequiredForR3 && a.Wall.BlocksOfficialLedger && !a.Wall.Native && containsAll(a.Wall.Failures, []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}), Detail: FormatWall(a.Wall)},
			{Name: "select next frontier without alpha loop or physical-yukawa jump", Passed: a.Decision.RecommendedFrontier == RecommendedNextFrontier && a.Decision.AvoidAlphaLoop && a.Decision.AvoidPhysicalYukawaJump && len(a.Decision.Frontiers) == 3 && a.Decision.Frontiers[1].Recommended && !a.Decision.Frontiers[0].Recommended && !a.Decision.Frontiers[2].Recommended, Detail: FormatDecision(a.Decision)},
			{Name: "preserve Gate 881 firewalls", Passed: firewallsOK(a.Firewalls), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, FormatLedger(a.Ledger), FormatAlpha(a.Alpha), FormatChain(a.Chain), FormatWall(a.Wall), FormatDecision(a.Decision), FormatClosure(a.Closure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
