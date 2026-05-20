package generation2effectiveparticipationscalarproxynormalformandruntimepropagationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate757InheritanceAndSubstitution(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate756.Inherited || !a.Gate756.EffectiveParticipationTyped || !a.Gate756.YukawaTraceLedgerSealed {
		t.Fatalf("bad Gate756 inheritance: %+v", a.Gate756)
	}
	if math.Abs(a.Gate756.NEff-nEffMZ) > 1e-15 || math.Abs(a.Gate756.InverseNEff-bOverA2MZ) > 1e-15 || math.Abs(a.Gate756.LambdaProxyFromNEff-lambdaProxyMZ) > 1e-15 {
		t.Fatalf("bad Gate756 numerics: %+v", a.Gate756)
	}
	if !a.Gate752.Inherited || !a.Gate752.ScalarMapTyped || a.Gate752.NativeScalarRuntime || math.Abs(a.Gate752.FWall3Red-0.00012565521035654) > 1e-16 || math.Abs(a.Gate752.RuntimeBracket-1.038025177923625) > 1e-12 {
		t.Fatalf("bad Gate752 inheritance: %+v", a.Gate752)
	}
	if !a.NormalForm.ProxySubstituted || !a.NormalForm.NormalFormWritten || !a.NormalForm.EquivalentExpandedTraceFormWritten || a.NormalForm.IndependentRuntimePrediction || a.NormalForm.NativeScalarProxyTheorem {
		t.Fatalf("bad normal form: %+v", a.NormalForm)
	}
	if math.Abs(a.NormalForm.LambdaRuntimeEff-(a.Gate756.LambdaProxyFromNEff*a.Gate752.RuntimeBracket)) > 1e-16 {
		t.Fatalf("bad runtime propagation: %+v", a.NormalForm)
	}
}

func TestGate757TopShadowRuntimeAndTreeProxyPropagation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TopShadow.ProxyBelowTopShadow || !a.TopShadow.RuntimeLoweredByParticipation || !a.TopShadow.TreeProxyDiagnosticOnly || a.TopShadow.HiggsPoleMassPrediction {
		t.Fatalf("bad top-shadow typing: %+v", a.TopShadow)
	}
	if math.Abs(a.TopShadow.LambdaProxyTopShadow-oneEighth) > 1e-15 || math.Abs(a.TopShadow.ProxyShift+9.689763984987998e-05) > 1e-15 {
		t.Fatalf("bad proxy shadow values: %+v", a.TopShadow)
	}
	if math.Abs(a.TopShadow.RuntimeShift+0.00010058218984558) > 1e-12 || math.Abs(a.TopShadow.TreeProxyShiftGeV+0.04862437568908) > 1e-10 {
		t.Fatalf("bad runtime/tree proxy propagation: %+v", a.TopShadow)
	}
}

func TestGate757InterpretationLayeringAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Interpretation.NEffGreaterThanThree || !a.Interpretation.LambdaProxyBelowOneEighth || !a.Interpretation.TraceLedgerMoreSpread || !a.Interpretation.NonTopChannelsDiluteIPR || !a.Interpretation.NonTopChannelsLowerProxy || a.Interpretation.AssignedSectorCorrection {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if !a.Relation.UsesAggregateTracePair || !a.Relation.RequiresNoTopYukawaChoice || !a.Relation.RequiresGate752TransportBracket || !a.Relation.Compatible {
		t.Fatalf("bad Gate756 relation: %+v", a.Relation)
	}
	if a.Layering.NEffIsNativeGenerationTheorem || a.Layering.NEffIsYukawaEigenvalueTheorem || a.Layering.NEffIsScalarPotentialTheorem || a.Layering.NEffIsRuntimeLambdaTheorem || a.Layering.NEffIsHiggsMassTheorem || !a.Layering.RuntimeBracketSeparateTransport {
		t.Fatalf("bad layer separation: %+v", a.Layering)
	}
	if a.Firewalls.NEffNativeGenerationTheorem || a.Firewalls.NEffMinusThreeAssignedToSector || a.Firewalls.LambdaProxyScalarPotentialTheorem || a.Firewalls.LambdaRuntimeIndependentPrediction || a.Firewalls.TreeProxyHiggsPoleMassPrediction || a.Firewalls.ClaimsYukawaEigenvaluesDerived || a.Firewalls.ClaimsFlavorHierarchyDerived || a.Firewalls.ClaimsCKMPMNSDerived || a.Firewalls.ClaimsHiggsMassTheorem || a.Firewalls.ClaimsPoleMassTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate757TheoremVerdictStatuses(t *testing.T) {
	res := Generation2EffectiveParticipationScalarProxyNormalFormAndRuntimePropagationAuditTheorem().Verify()
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
