package hierarchyscalingaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func GaugeHierarchyScalingAuditPlanckFactorSieveTheorem() theorem.Theorem {
	const id = "BRIDGE-GAUGE-HIERARCHY-SCALING-AUDIT-PLANCK-FACTOR-SIEVE"
	const name = "Gauge Hierarchy Scaling Audit / Planck Factor Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 339 hierarchy scaling audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 338 precision audit inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.SInst > 12 && a.Inputs.STop > 78, Detail: FormatInputs(a.Inputs)},
			{Name: "hierarchy ratios formalized", Passed: a.Targets.RhoUnreduced > 1e-17 && a.Targets.RhoUnreduced < 3e-17 && a.Targets.RhoReduced > 9e-17 && a.Targets.RhoReduced < 2e-16, Detail: FormatTargets(a.Targets)},
			{Name: "topological candidates audited", Passed: len(a.Candidates.Candidates) >= 8 && a.Candidates.BestUnreduced.Name != "", Detail: FormatCandidates(a.Candidates)},
			{Name: "scale synthesis rejects arbitrary exponent fitting", Passed: !a.Synthesis.NativeDerived && hasFitLaneRejected(a.Synthesis), Detail: FormatSynthesis(a.Synthesis)},
			{Name: "rank-56 near miss not promoted", Passed: hasCandidate(a.Candidates, "rank-56 Boolean near miss") && !a.Candidates.BestUnreduced.Promotable, Detail: FormatCandidate(a.Candidates.BestUnreduced)},
			{Name: "hierarchy firewalls preserved", Passed: a.Firewalls.NoHierarchyScalingFactor && a.Firewalls.F2MomentUnlocked && a.Firewalls.PlanckScaleNotNative && a.Firewalls.ArbitraryExponentFittingRejected, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}

func hasCandidate(l CandidateLedger, name string) bool {
	for _, c := range l.Candidates {
		if c.Name == name {
			return true
		}
	}
	return false
}

func hasFitLaneRejected(s SynthesisSieve) bool {
	for _, l := range s.Lanes {
		if l.Name == "fit exponent lane" && l.RequiresFreeExponent && l.RequiresUnprovedScaleLaw {
			return true
		}
	}
	return false
}
