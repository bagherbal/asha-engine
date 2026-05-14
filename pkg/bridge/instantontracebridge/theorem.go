package instantontracebridge

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FiniteToContinuumInstantonTraceNormalizationBridgeTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-CONTINUUM-INSTANTON-TRACE-NORMALIZATION"
	const name = "finite-to-continuum instanton trace-normalization bridge"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build finite-to-continuum instanton bridge audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "Gate 174 conditional branch is inherited without observed input", Passed: a.Input.Gate174ConditionalBranchAvailable && !a.Input.Gate174StrictAbsoluteUDerived && a.Input.RelativeGaugeRatioClosed && a.Input.WeakAngleSeedClosed && !a.Input.UsesObservedInput, Detail: FormatInput(a.Input)},
			{Name: "continuum index bridge requirements are not met", Passed: a.Firewall.ContinuumIndexRequirements == 5 && a.Firewall.ContinuumIndexRequirementsMet < a.Firewall.ContinuumIndexRequirements && !a.Firewall.StrictContinuumIndexBridgeDerived, Detail: FormatRequirements(a.IndexNeeds)},
			{Name: "trace/kinetic normalization requirements are not met", Passed: a.Firewall.TraceRequirements == 5 && a.Firewall.TraceRequirementsMet < a.Firewall.TraceRequirements && !a.Firewall.StrictTraceKineticBridgeDerived, Detail: FormatRequirements(a.TraceNeeds)},
			{Name: "all shortcut routes are quarantined", Passed: a.Firewall.CandidateRoutesAudited == 5 && noStrictRoute(a.Routes) && hasConditionalRoute(a.Routes) && hasObservedForbiddenRoute(a.Routes), Detail: FormatRoutes(a.Routes)},
			{Name: "representation trace remains relative, not absolute", Passed: a.TraceAudit.RepresentationTraceRatioClosed && a.TraceAudit.GeneratorNormalizationRelative && !a.TraceAudit.AbsoluteTraceScaleDerived && a.TraceAudit.FieldRescalingFreedomOpen && a.TraceAudit.F0ConventionDependenceOpen && !a.TraceAudit.KineticIntegralNormalization, Detail: FormatTraceAudit(a.TraceAudit)},
			{Name: "conditional u=1 branch is preserved", Passed: a.Firewall.ConditionalAbsoluteUPreserved && a.Firewall.ConditionalNullityAfter == 2, Detail: FormatFirewall(a.Firewall)},
			{Name: "strict absolute coupling remains sealed", Passed: !a.Firewall.StrictAbsoluteUDerived && a.Firewall.StrictNullityBefore == 3 && a.Firewall.StrictNullityAfter == 3 && !a.Firewall.PhysicalCouplingsDerived && !a.Firewall.FineStructureDerived && !a.Firewall.HiddenObservedInputUsed, Detail: FormatFirewall(a.Firewall)},
		}, Notes: []string{
			a.TruthStatement,
			"The theorem preserves the useful conditional branch u=1, but refuses to promote it to a derived physical coupling without a continuum Chern--Weil and kinetic-trace bridge.",
			"Next gate may study RG solvability under the quarantined conditional branch, clearly separated from strict theorem status.",
		}}
	}}
}

func noStrictRoute(xs []CandidateRoute) bool {
	for _, x := range xs {
		if x.AdmissibleAsStrictTheorem {
			return false
		}
	}
	return true
}

func hasConditionalRoute(xs []CandidateRoute) bool {
	for _, x := range xs {
		if x.AdmissibleAsConditionalBranch {
			return true
		}
	}
	return false
}

func hasObservedForbiddenRoute(xs []CandidateRoute) bool {
	for _, x := range xs {
		if x.UsesObservedInput && !x.AdmissibleAsStrictTheorem && !x.AdmissibleAsConditionalBranch {
			return true
		}
	}
	return false
}
