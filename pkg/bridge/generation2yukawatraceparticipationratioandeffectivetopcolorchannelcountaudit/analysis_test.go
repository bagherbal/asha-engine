package generation2yukawatraceparticipationratioandeffectivetopcolorchannelcountaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate756InheritanceAtomsAndIPR(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate755.Inherited || !a.Gate755.AggregateTracePairAvailable || !a.Gate755.RequiresNoTopYukawaChoice || !a.Gate755.TopDominanceDeviationTyped || !a.Gate755.NumericalAlphaBetaBlocked || !a.Gate755.NativeYukawaAndScalarTheoremsBlocked {
		t.Fatalf("bad Gate755 inheritance: %+v", a.Gate755)
	}
	if math.Abs(a.Gate755.BOverA2Computed-bOverA2MZSeed) > 1e-15 || math.Abs(a.Gate755.DeltaRatio+0.0002583937062663466) > 1e-15 {
		t.Fatalf("bad inherited trace values: %+v", a.Gate755)
	}
	if !a.Atoms.ColorFactorExpandedAsRepeatedAtoms || !a.Atoms.AtomsPositive || !a.Atoms.RequiresDecomposedYukawaLedgerForAtoms || !a.Atoms.UsesOnlyAggregatePairForNEff {
		t.Fatalf("bad trace-atom expansion: %+v", a.Atoms)
	}
	if !a.IPR.RatioIsIPR || !a.IPR.BasisCleanAggregateDiagnostic || a.IPR.NativeYukawaTheorem || math.Abs(a.IPR.ComputedRatio-bOverA2MZSeed) > 1e-15 || math.Abs(a.IPR.SyntheticTopColorIPR-oneThird) > 1e-15 {
		t.Fatalf("bad inverse participation ratio typing: %+v", a.IPR)
	}
}

func TestGate756EffectiveCountAndProxyRewrite(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Effective.NativeGenerationTheorem || a.Effective.ChannelAssignmentWithoutLedger || !a.Effective.CurrentLedgerAboveThree || !a.Effective.NearThree || !a.Effective.InterpretedAsTinyTraceSpread {
		t.Fatalf("bad effective count typing: %+v", a.Effective)
	}
	if math.Abs(a.Effective.ComputedFromRatio-3.0023273474722147) > 1e-15 || math.Abs(a.Effective.DeviationFromThree-0.0023273474722147) > 1e-15 || math.Abs(a.Effective.RelativeDeviationFromThree-0.000775782490738249) > 1e-15 {
		t.Fatalf("bad effective count numerics: %+v", a.Effective)
	}
	if !a.Proxy.ThreeOverEightNEffIdentity || !a.Proxy.ProxyBelowOneEighth || a.Proxy.ScalarPotentialTheorem || a.Proxy.RuntimeLambdaTheorem || math.Abs(a.Proxy.LambdaProxyComputed-lambdaProxyMZ) > 1e-15 || math.Abs(a.Proxy.TopColorProxyLimit-oneEighth) > 1e-15 {
		t.Fatalf("bad proxy rewrite: %+v", a.Proxy)
	}
}

func TestGate756RelationToGate755AndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Relation.Gate755NeedsDecomposedLedger || !a.Relation.Gate756WorksFromAggregateTracePair || !a.Relation.CompatibleDiagnostics || math.Abs(a.Relation.ProbeCompatibilityResidual) > 1e-16 {
		t.Fatalf("bad Gate755/Gate756 relation: %+v", a.Relation)
	}
	if a.Yukawa.NEffIsNativeGenerationTheorem || a.Yukawa.NEffDerivesFlavorHierarchy || a.Yukawa.NEffMinusThreeAssignedToChannel || a.Yukawa.ClaimsYuDerived || a.Yukawa.ClaimsYdDerived || a.Yukawa.ClaimsYeDerived || a.Yukawa.ClaimsYnuDerived || a.Yukawa.ClaimsCKMPMNSDerived || a.Yukawa.ClaimsNativeFlavorTheorem || !a.Yukawa.SealedLedgerExplicit {
		t.Fatalf("bad Yukawa firewall: %+v", a.Yukawa)
	}
	if a.Runtime.LambdaProxyNearOneEighthIsScalarPotentialTheorem || a.Runtime.LambdaProxyEqualsRuntimeLambda || a.Runtime.RuntimeLambdaEqualsHiggsMass || a.Runtime.TreeProxyEqualsPoleMass || a.Runtime.ClaimsIndependentScalarRuntime || a.Runtime.ClaimsHiggsMassTheorem || a.Runtime.ClaimsPoleMassTheorem || !a.Runtime.RequiresHistoryLoopTransport || !a.Runtime.RequiresBoundaryHistoryResponse || !a.Runtime.RequiresKappaEReduction || !a.Runtime.RequiresScalarRuntimeBridge {
		t.Fatalf("bad runtime/Higgs firewall: %+v", a.Runtime)
	}
}

func TestGate756TheoremVerdictStatuses(t *testing.T) {
	res := Generation2YukawaTraceParticipationRatioAndEffectiveTopColorChannelCountAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
