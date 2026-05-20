package generation2finiteyukawatraceratioonethirdshadowandtopcolordominanceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate754TopColorDominanceTraceRatio(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate753.Inherited || !a.Gate753.BA2OneThirdTheoremBlocked || !a.Gate753.ScalarProxyDerivationBlocked || !a.Gate753.ProxyToRuntimeMatchingBlocked || !a.Gate753.HiggsMassOrPoleMassTheoremBlocked {
		t.Fatalf("bad Gate753 inheritance: %+v", a.Gate753)
	}
	if !a.Inputs.NativeTraceShapes || !a.Inputs.SealedYukawaLedger || a.Inputs.NativeRatioTheorem || math.Abs(a.Inputs.BOverA2Computed-bOverA2MZSeed) > 1e-15 {
		t.Fatalf("bad trace ratio input: %+v", a.Inputs)
	}
	if !a.TopLimit.SingleDominantChannel || !a.TopLimit.ColoredChannel || !a.TopLimit.ExactLimitDerived || a.TopLimit.NativeYukawaTheorem || math.Abs(a.TopLimit.RatioAtUnitY-oneThird) > 1e-15 {
		t.Fatalf("bad top-color dominance limit: %+v", a.TopLimit)
	}
}

func TestGate754DeviationShadowAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Deviation.DeltaRatio+0.0002583937062663466) > 1e-15 || !a.Deviation.MeasuresNonTopDominanceCorrection || a.Deviation.DeltaSourceAssigned || a.Deviation.NativeDeltaDecomposition {
		t.Fatalf("bad deviation audit: %+v", a.Deviation)
	}
	if !a.OneEighth.CloseToOneEighth || a.OneEighth.ScalarPotentialTheorem || math.Abs(a.OneEighth.IdealProxy-oneEighth) > 1e-15 || math.Abs(a.OneEighth.ShadowIdentityResidual) > 1e-16 {
		t.Fatalf("bad one-eighth shadow: %+v", a.OneEighth)
	}
	if !a.SourceLayers.AllLayersSeparated || a.SourceLayers.ClaimsThreeEighthsScalarLaw || a.SourceLayers.ClaimsOneThirdYukawaLaw || a.SourceLayers.ClaimsOneEighthScalarLaw {
		t.Fatalf("bad source-layer firewall: %+v", a.SourceLayers)
	}
	if a.Yukawa.ClaimsYuDerived || a.Yukawa.ClaimsYdDerived || a.Yukawa.ClaimsYeDerived || a.Yukawa.ClaimsYnuDerived || a.Yukawa.ClaimsTopYukawaDerived || a.Yukawa.ClaimsYukawaHierarchyDerived || a.Yukawa.ClaimsCKMPMNSDerived || a.Yukawa.ClaimsGenerationCarrier || !a.Yukawa.SealedLedgerExplicit {
		t.Fatalf("bad Yukawa firewall: %+v", a.Yukawa)
	}
	if a.Runtime.ClaimsRuntimeLambdaTheorem || a.Runtime.ClaimsIndependentScalarRuntime || a.Runtime.ClaimsProxyToRuntimeMatching || a.Runtime.ClaimsHiggsMassTheorem || a.Runtime.ClaimsPoleMassTheorem || !a.Runtime.RequiresHistoryLoopTransport || !a.Runtime.RequiresBoundaryHistoryResponse || !a.Runtime.RequiresKappaEReduction || !a.Runtime.RequiresScalarRuntimeBridge {
		t.Fatalf("bad runtime firewall: %+v", a.Runtime)
	}
}

func TestGate754TheoremVerdictStatuses(t *testing.T) {
	res := Generation2FiniteYukawaTraceRatioOneThirdShadowAndTopColorDominanceAuditTheorem().Verify()
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
