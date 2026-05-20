package generation2topdominancetracedeviationexpansionandnontopyukawafirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate755InheritanceSplitAndRestVariables(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate754.Inherited || !a.Gate754.TopColorDominanceLimitDefined || !a.Gate754.OneThirdTraceShadowConditional || !a.Gate754.NativeBA2OneThirdTheoremBlocked || !a.Gate754.NativeDeltaDecompositionBlocked || !a.Gate754.RuntimeAndHiggsTheoremsBlocked {
		t.Fatalf("bad Gate754 inheritance: %+v", a.Gate754)
	}
	if math.Abs(a.Gate754.BOverA2Computed-bOverA2MZSeed) > 1e-15 || math.Abs(a.Gate754.DeltaRatio+0.0002583937062663466) > 1e-15 {
		t.Fatalf("bad inherited ratio/deviation: %+v", a.Gate754)
	}
	if !a.Split.SplitIsAlgebraicIdentity || a.Split.TopYukawaValueDerived || math.Abs(a.Split.TopLimitRatio-oneThird) > 1e-15 {
		t.Fatalf("bad top-color split: %+v", a.Split)
	}
	if !a.Rest.AlphaRequiresTopValue || !a.Rest.BetaRequiresTopValue || !a.Rest.RequiresDecomposedYukawaLedger || a.Rest.NumericalAlphaBetaAvailable {
		t.Fatalf("bad normalized rest variables: %+v", a.Rest)
	}
}

func TestGate755DeviationFormulaAndProxyTransport(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Deviation.NativeDeltaRatioTheorem || !a.Deviation.CurrentRatioBelowOneThird || !a.Deviation.FirstOrderExplainsSign || math.Abs(a.Deviation.ProbeFormulaResidual) > 1e-16 || math.Abs(a.Deviation.ProbeRatioDirect-a.Deviation.ProbeRatioByFormula) > 1e-16 {
		t.Fatalf("bad trace deviation formula: %+v", a.Deviation)
	}
	if !a.Proxy.DeviationTransported || a.Proxy.ScalarPotentialTheorem || a.Proxy.RuntimeLambdaTheorem || math.Abs(a.Proxy.LambdaProxyComputed-lambdaProxyMZ) > 1e-13 || math.Abs(a.Proxy.TransportIdentityResidual) > 1e-16 || math.Abs(a.Proxy.ProxyMinusOneEighth-a.Proxy.ThreeEighthsTimesDelta) > 1e-16 {
		t.Fatalf("bad one-eighth proxy transport: %+v", a.Proxy)
	}
}

func TestGate755RequiredDataAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.RequiredData.RequiredItems) != 5 || a.RequiredData.CanComputeAlphaBeta || a.RequiredData.CanAssignBottomTauCharm || a.RequiredData.CanAssignNeutrinoConvention || a.RequiredData.CanAssignScaleDependence || a.RequiredData.CanAssignFiniteTraceResidual || a.RequiredData.TypedTopLikeTAvailable || a.RequiredData.DecomposedYukawaLedger {
		t.Fatalf("bad required decomposition data firewall: %+v", a.RequiredData)
	}
	if a.Yukawa.DeltaRatioIsNativeYukawaTheorem || a.Yukawa.TopDominanceDerivesTopYukawa || a.Yukawa.AlphaBetaDerivesHierarchy || a.Yukawa.ClaimsYuDerived || a.Yukawa.ClaimsYdDerived || a.Yukawa.ClaimsYeDerived || a.Yukawa.ClaimsYnuDerived || a.Yukawa.ClaimsCKMPMNSDerived || a.Yukawa.ClaimsGenerationCarrier || a.Yukawa.ClaimsFlavorTheorem || !a.Yukawa.SealedLedgerExplicit {
		t.Fatalf("bad Yukawa firewall: %+v", a.Yukawa)
	}
	if a.Runtime.LambdaProxyNearOneEighthIsScalarPotentialTheorem || a.Runtime.LambdaProxyEqualsRuntimeLambda || a.Runtime.RuntimeLambdaEqualsHiggsMass || a.Runtime.TreeProxyEqualsPoleMass || a.Runtime.ClaimsIndependentScalarRuntime || a.Runtime.ClaimsHiggsMassTheorem || a.Runtime.ClaimsPoleMassTheorem || !a.Runtime.RequiresHistoryLoopTransport || !a.Runtime.RequiresBoundaryHistoryResponse || !a.Runtime.RequiresKappaEReduction || !a.Runtime.RequiresScalarRuntimeBridge {
		t.Fatalf("bad runtime/Higgs firewall: %+v", a.Runtime)
	}
}

func TestGate755TheoremVerdictStatuses(t *testing.T) {
	res := Generation2TopDominanceTraceDeviationExpansionAndNonTopYukawaFirewallAuditTheorem().Verify()
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
