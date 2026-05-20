package generation2k7hodge43higgsflavorshadowfirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func closeEnough(a, b float64) bool { return math.Abs(a-b) < tolerance }

func TestK7Hodge43ShadowAndFanoFrame(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.InheritedGate707.CentralBaselineGaugeInherited || !a.InheritedGate707.K7LocalUpliftReference || !a.InheritedGate707.BaselineChoiceNotNative || a.InheritedGate707.Verdict != StatusGate707CentralBaselineGaugeInherited {
		t.Fatalf("bad Gate707 inheritance: %+v", a.InheritedGate707)
	}
	if a.Hodge.K7Dimension != 7 || a.Hodge.PlusDimension != 4 || a.Hodge.MinusDimension != 3 || !closeEnough(a.Hodge.Trace, 1) || !closeEnough(a.Hodge.Determinant, -1) || !a.Hodge.HodgeStable || !a.Hodge.MixedHodgePolarity || !a.Hodge.Gate634FirewallPreserved {
		t.Fatalf("bad hodge inheritance: %+v", a.Hodge)
	}
	if !a.Shadow.HiggsDimensionMatches || !a.Shadow.FlavorTripletMatches || a.Shadow.PhysicalMapCertified || !a.Shadow.OnlyDimensionShadow || !strings.Contains(a.Shadow.Verdict, StatusK7Hodge43MatchesShadow) || !strings.Contains(a.Shadow.Verdict, StatusSevenNumeratorInternal43EventRankShadow) {
		t.Fatalf("bad shadow audit: %+v", a.Shadow)
	}
	if !a.FanoHitchin.CouplingFrameCandidate || a.FanoHitchin.YukawaMapCertified || a.FanoHitchin.YukawaEigenvaluesCertified || a.FanoHitchin.FlavorHierarchyCertified || a.FanoHitchin.CouplingFrameSize != 12 || !strings.Contains(a.FanoHitchin.Verdict, StatusFanoHitchinCouplingFrameCandidate) {
		t.Fatalf("bad Fano-Hitchin audit: %+v", a.FanoHitchin)
	}
}

func TestObstructionNumbersFirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	o := a.Obstructions
	if !closeEnough(o.CosTheta, 13/math.Sqrt(217)) || !closeEnough(o.RhoSquared, 48.0/217.0) || o.Root13 != 13 || o.Numerator48 != 48 || o.Denominator217 != 217 || !o.InternalOnly || !o.NotSMFlavorParams || o.Verdict != StatusInternalObstructionNumbersRecorded {
		t.Fatalf("bad obstruction audit: %+v", o)
	}
	f := a.Firewalls
	if f.ClaimsK7PlusIsPhysicalHiggsDoublet || f.ClaimsK7MinusIsGenerationSpace || f.ClaimsFanoTripletFlavorTheorem || f.ClaimsOmegaIsYukawaMatrix || f.ClaimsSevenDerivesHiggsFlavor || f.ClaimsHiggsMassDerivation || f.ClaimsYukawaEigenvalueTheorem || f.ClaimsFlavorHierarchyTheorem || f.ClaimsCKMPMNSTheorem || f.ClaimsInternal13AsSMFlavorDerivation || f.Verdict != StatusGate708K7HodgeHiggsFlavorShadowBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	m := a.Missing
	if len(m.Missing) != 8 || !strings.Contains(m.Verdict, StatusNoTypedK7PlusToHiggsDoublet) || !strings.Contains(m.Verdict, StatusNoTypedK7MinusToGenerationSpace) || !strings.Contains(m.Verdict, StatusNoNativeYukawaEigenvalueTheorem) || !strings.Contains(m.Verdict, StatusInternal13NotSMFlavorParameterDerivation) || !strings.Contains(m.Verdict, StatusNoNativeHiggsFlavorRepresentationMap) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing boundary: %+v", m)
	}
	res := Generation2K7Hodge43HiggsFlavorShadowFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
