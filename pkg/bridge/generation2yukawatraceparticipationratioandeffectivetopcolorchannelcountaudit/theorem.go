package generation2yukawatraceparticipationratioandeffectivetopcolorchannelcountaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2YukawaTraceParticipationRatioAndEffectiveTopColorChannelCountAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 756 — Yukawa Trace Participation Ratio and Effective Top-Color Channel Count Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate756 Yukawa trace participation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate755 top-dominance trace-deviation boundary", Passed: a.Gate755.Inherited && a.Gate755.AggregateTracePairAvailable && a.Gate755.RequiresNoTopYukawaChoice && a.Gate755.TopDominanceDeviationTyped && a.Gate755.NumericalAlphaBetaBlocked && a.Gate755.NativeYukawaAndScalarTheoremsBlocked && math.Abs(a.Gate755.BOverA2Computed-bOverA2MZSeed) < 1e-15 && math.Abs(a.Gate755.DeltaRatio+0.0002583937062663466) < 1e-15, Detail: FormatGate755(a.Gate755)},
			{Name: "define positive trace-atom expansion", Passed: a.Atoms.ColorFactorExpandedAsRepeatedAtoms && a.Atoms.AtomsPositive && a.Atoms.RequiresDecomposedYukawaLedgerForAtoms && a.Atoms.UsesOnlyAggregatePairForNEff && strings.Contains(a.Atoms.AFormula, "sum_i x_i") && strings.Contains(a.Atoms.BFormula, "sum_i x_i^2") && strings.Contains(a.Atoms.IPRFormula, "sum_i w_i^2"), Detail: FormatAtoms(a.Atoms)},
			{Name: "type b/a^2 as inverse participation ratio", Passed: a.IPR.RatioIsIPR && a.IPR.BasisCleanAggregateDiagnostic && !a.IPR.NativeYukawaTheorem && math.Abs(a.IPR.ComputedRatio-bOverA2MZSeed) < 1e-15 && math.Abs(a.IPR.SeedResidual) < 1e-15 && math.Abs(a.IPR.SyntheticTopColorIPR-oneThird) < 1e-15 && len(a.IPR.SyntheticTopColorWeights) == 3, Detail: FormatIPR(a.IPR)},
			{Name: "compute effective top-color channel count", Passed: !a.Effective.NativeGenerationTheorem && !a.Effective.ChannelAssignmentWithoutLedger && a.Effective.CurrentLedgerAboveThree && a.Effective.NearThree && a.Effective.InterpretedAsTinyTraceSpread && math.Abs(a.Effective.ComputedFromRatio-3.0023273474722147) < 1e-15 && math.Abs(a.Effective.ComputedFromTracePair-3.0023273474722143) < 1e-15 && math.Abs(a.Effective.DeviationFromThree-0.0023273474722147) < 1e-15 && math.Abs(a.Effective.RelativeDeviationFromThree-0.000775782490738249) < 1e-15, Detail: FormatEffective(a.Effective)},
			{Name: "rewrite one-eighth proxy through N_eff", Passed: a.Proxy.ThreeOverEightNEffIdentity && a.Proxy.ProxyBelowOneEighth && !a.Proxy.ScalarPotentialTheorem && !a.Proxy.RuntimeLambdaTheorem && math.Abs(a.Proxy.LambdaProxyComputed-lambdaProxyMZ) < 1e-15 && math.Abs(a.Proxy.TopColorProxyLimit-oneEighth) < 1e-15 && strings.Contains(a.Proxy.ParticipationFormula, "3/(8N_eff)") && strings.Contains(a.Proxy.EquivalentFormula, "3/N_eff"), Detail: FormatProxy(a.Proxy)},
			{Name: "record compatibility with Gate755 alpha-beta form", Passed: a.Relation.Gate755NeedsDecomposedLedger && a.Relation.Gate756WorksFromAggregateTracePair && a.Relation.CompatibleDiagnostics && math.Abs(a.Relation.ProbeCompatibilityResidual) < 1e-16 && math.Abs(a.Relation.ProbeRatioGate755-a.Relation.ProbeInverseNEff) < 1e-16 && strings.Contains(a.Relation.NEffAlphaBetaFormula, "3(1+alpha)^2/(1+beta)"), Detail: FormatRelation(a.Relation)},
			{Name: "enforce Yukawa and flavor firewall", Passed: !a.Yukawa.NEffIsNativeGenerationTheorem && !a.Yukawa.NEffDerivesFlavorHierarchy && !a.Yukawa.NEffMinusThreeAssignedToChannel && !a.Yukawa.ClaimsYuDerived && !a.Yukawa.ClaimsYdDerived && !a.Yukawa.ClaimsYeDerived && !a.Yukawa.ClaimsYnuDerived && !a.Yukawa.ClaimsCKMPMNSDerived && !a.Yukawa.ClaimsNativeFlavorTheorem && a.Yukawa.SealedLedgerExplicit, Detail: FormatYukawaFirewall(a.Yukawa)},
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
