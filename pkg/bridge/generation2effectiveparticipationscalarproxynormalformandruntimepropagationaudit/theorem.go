package generation2effectiveparticipationscalarproxynormalformandruntimepropagationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2EffectiveParticipationScalarProxyNormalFormAndRuntimePropagationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 757 — Effective-Participation Scalar Proxy Normal Form and Runtime Propagation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate757 effective-participation scalar-proxy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate756 Yukawa trace participation", Passed: a.Gate756.Inherited && a.Gate756.EffectiveParticipationTyped && a.Gate756.YukawaTraceLedgerSealed && math.Abs(a.Gate756.NEff-nEffMZ) < 1e-15 && math.Abs(a.Gate756.InverseNEff-bOverA2MZ) < 1e-15 && math.Abs(a.Gate756.LambdaProxyFromNEff-lambdaProxyMZ) < 1e-15 && strings.Contains(a.Gate756.LambdaProxyFormula, "3/(8N_eff)"), Detail: FormatGate756(a.Gate756)},
			{Name: "inherit Gate752 flavor-reduced scalar-Higgs normal form", Passed: a.Gate752.Inherited && a.Gate752.ScalarMapTyped && !a.Gate752.NativeScalarRuntime && math.Abs(a.Gate752.P-pK7) < 1e-15 && math.Abs(a.Gate752.KappaERed-kappaERed) < 1e-18 && math.Abs(a.Gate752.FWall3Red-0.00012565521035654) < 1e-16 && math.Abs(a.Gate752.RuntimeBracket-1.038025177923625) < 1e-12 && strings.Contains(a.Gate752.RuntimeBridgeFormula, "lambda_proxy"), Detail: FormatGate752(a.Gate752)},
			{Name: "substitute effective participation proxy into scalar-Higgs normal form", Passed: a.NormalForm.ProxySubstituted && a.NormalForm.NormalFormWritten && a.NormalForm.EquivalentExpandedTraceFormWritten && !a.NormalForm.IndependentRuntimePrediction && !a.NormalForm.NativeScalarProxyTheorem && math.Abs(a.NormalForm.LambdaRuntimeEff-(a.Gate756.LambdaProxyFromNEff*a.Gate752.RuntimeBracket)) < 1e-16 && strings.Contains(a.NormalForm.RuntimeFormula, "3/(8N_eff)"), Detail: FormatNormalForm(a.NormalForm)},
			{Name: "audit top-color shadow comparison and runtime propagation", Passed: a.TopShadow.ProxyBelowTopShadow && a.TopShadow.RuntimeLoweredByParticipation && a.TopShadow.TreeProxyDiagnosticOnly && !a.TopShadow.HiggsPoleMassPrediction && math.Abs(a.TopShadow.NEffTop-3.0) < 1e-15 && math.Abs(a.TopShadow.LambdaProxyTopShadow-oneEighth) < 1e-15 && math.Abs(a.TopShadow.ProxyShift+9.689763984987998e-05) < 1e-15 && math.Abs(a.TopShadow.RuntimeShift+0.00010058218984558) < 1e-12 && math.Abs(a.TopShadow.TreeProxyShiftGeV+0.04862437568908) < 1e-10, Detail: FormatTopShadow(a.TopShadow)},
			{Name: "record participation interpretation", Passed: a.Interpretation.NEffGreaterThanThree && a.Interpretation.LambdaProxyBelowOneEighth && a.Interpretation.TraceLedgerMoreSpread && a.Interpretation.NonTopChannelsDiluteIPR && a.Interpretation.NonTopChannelsLowerProxy && !a.Interpretation.AssignedSectorCorrection, Detail: FormatInterpretation(a.Interpretation)},
			{Name: "record relation to Gate756 participation form", Passed: a.Relation.UsesAggregateTracePair && a.Relation.RequiresNoTopYukawaChoice && a.Relation.RequiresGate752TransportBracket && a.Relation.Compatible && strings.Contains(a.Relation.Gate757NormalFormFormula, "transport_bracket_red"), Detail: FormatRelation(a.Relation)},
			{Name: "enforce layer separation", Passed: !a.Layering.NEffIsNativeGenerationTheorem && !a.Layering.NEffIsYukawaEigenvalueTheorem && !a.Layering.NEffIsScalarPotentialTheorem && !a.Layering.NEffIsRuntimeLambdaTheorem && !a.Layering.NEffIsHiggsMassTheorem && a.Layering.RuntimeBracketSeparateTransport, Detail: FormatLayering(a.Layering)},
			{Name: "enforce Yukawa, scalar-runtime, Higgs, and pole firewalls", Passed: !a.Firewalls.NEffNativeGenerationTheorem && !a.Firewalls.NEffMinusThreeAssignedToSector && !a.Firewalls.LambdaProxyScalarPotentialTheorem && !a.Firewalls.LambdaRuntimeIndependentPrediction && !a.Firewalls.TreeProxyHiggsPoleMassPrediction && !a.Firewalls.ClaimsYukawaEigenvaluesDerived && !a.Firewalls.ClaimsFlavorHierarchyDerived && !a.Firewalls.ClaimsCKMPMNSDerived && !a.Firewalls.ClaimsHiggsMassTheorem && !a.Firewalls.ClaimsPoleMassTheorem, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
