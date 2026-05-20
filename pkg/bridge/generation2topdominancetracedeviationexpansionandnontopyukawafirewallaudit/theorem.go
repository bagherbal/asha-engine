package generation2topdominancetracedeviationexpansionandnontopyukawafirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2TopDominanceTraceDeviationExpansionAndNonTopYukawaFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 755 — Top-Dominance Trace-Deviation Expansion and Non-Top Yukawa Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate755 top-dominance trace-deviation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate754 one-third trace shadow", Passed: a.Gate754.Inherited && a.Gate754.TopColorDominanceLimitDefined && a.Gate754.OneThirdTraceShadowConditional && a.Gate754.NativeBA2OneThirdTheoremBlocked && a.Gate754.NativeDeltaDecompositionBlocked && a.Gate754.RuntimeAndHiggsTheoremsBlocked && math.Abs(a.Gate754.BOverA2Computed-bOverA2MZSeed) < 1e-15 && math.Abs(a.Gate754.DeltaRatio+0.0002583937062663466) < 1e-15, Detail: FormatGate754(a.Gate754)},
			{Name: "define top-color dominant split", Passed: a.Split.SplitIsAlgebraicIdentity && !a.Split.TopYukawaValueDerived && strings.Contains(a.Split.DominantSquaredSingularValueSymbol, "T") && strings.Contains(a.Split.ATopFormula, "3T") && strings.Contains(a.Split.BTopFormula, "3T^2") && strings.Contains(a.Split.TopLimitRatioFormula, "1/3") && math.Abs(a.Split.TopLimitRatio-oneThird) < 1e-15, Detail: FormatSplit(a.Split)},
			{Name: "define normalized rest variables without numerical assignment", Passed: a.Rest.AlphaRequiresTopValue && a.Rest.BetaRequiresTopValue && a.Rest.RequiresDecomposedYukawaLedger && !a.Rest.NumericalAlphaBetaAvailable && strings.Contains(a.Rest.AlphaDefinition, "a_rest/(3T)") && strings.Contains(a.Rest.BetaDefinition, "b_rest/(3T^2)") && strings.Contains(a.Rest.RatioRewritten, "(1/3)"), Detail: FormatRest(a.Rest)},
			{Name: "derive exact trace-deviation formula", Passed: !a.Deviation.NativeDeltaRatioTheorem && a.Deviation.CurrentRatioBelowOneThird && a.Deviation.AssumptionAlphaPositive && a.Deviation.AssumptionBetaMuchLessAlpha && a.Deviation.FirstOrderExplainsSign && math.Abs(a.Deviation.ProbeFormulaResidual) < 1e-16 && math.Abs(a.Deviation.ProbeRatioDirect-a.Deviation.ProbeRatioByFormula) < 1e-16 && strings.Contains(a.Deviation.ExactFormula, "beta-2alpha-alpha^2"), Detail: FormatDeviation(a.Deviation)},
			{Name: "rewrite one-eighth proxy deviation", Passed: a.Proxy.DeviationTransported && !a.Proxy.ScalarPotentialTheorem && !a.Proxy.RuntimeLambdaTheorem && math.Abs(a.Proxy.LambdaProxyComputed-lambdaProxyMZ) < 1e-13 && math.Abs(a.Proxy.ProxyMinusOneEighth-a.Proxy.ThreeEighthsTimesDelta) < 1e-16 && math.Abs(a.Proxy.TransportIdentityResidual) < 1e-16 && strings.Contains(a.Proxy.FormulaFromTraceDeviation, "delta_ratio"), Detail: FormatProxy(a.Proxy)},
			{Name: "list required Yukawa decomposition data", Passed: len(a.RequiredData.RequiredItems) == 5 && !a.RequiredData.CanComputeAlphaBeta && !a.RequiredData.CanAssignBottomTauCharm && !a.RequiredData.CanAssignNeutrinoConvention && !a.RequiredData.CanAssignScaleDependence && !a.RequiredData.CanAssignFiniteTraceResidual && !a.RequiredData.TypedTopLikeTAvailable && !a.RequiredData.DecomposedYukawaLedger, Detail: FormatRequiredData(a.RequiredData)},
			{Name: "enforce non-top Yukawa theorem firewall", Passed: !a.Yukawa.DeltaRatioIsNativeYukawaTheorem && !a.Yukawa.TopDominanceDerivesTopYukawa && !a.Yukawa.AlphaBetaDerivesHierarchy && !a.Yukawa.ClaimsYuDerived && !a.Yukawa.ClaimsYdDerived && !a.Yukawa.ClaimsYeDerived && !a.Yukawa.ClaimsYnuDerived && !a.Yukawa.ClaimsCKMPMNSDerived && !a.Yukawa.ClaimsGenerationCarrier && !a.Yukawa.ClaimsFlavorTheorem && a.Yukawa.SealedLedgerExplicit, Detail: FormatYukawaFirewall(a.Yukawa)},
			{Name: "enforce runtime and Higgs firewalls", Passed: !a.Runtime.LambdaProxyNearOneEighthIsScalarPotentialTheorem && !a.Runtime.LambdaProxyEqualsRuntimeLambda && !a.Runtime.RuntimeLambdaEqualsHiggsMass && !a.Runtime.TreeProxyEqualsPoleMass && !a.Runtime.ClaimsIndependentScalarRuntime && !a.Runtime.ClaimsHiggsMassTheorem && !a.Runtime.ClaimsPoleMassTheorem && a.Runtime.RequiresHistoryLoopTransport && a.Runtime.RequiresBoundaryHistoryResponse && a.Runtime.RequiresKappaEReduction && a.Runtime.RequiresScalarRuntimeBridge, Detail: FormatRuntimeFirewall(a.Runtime)},
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
