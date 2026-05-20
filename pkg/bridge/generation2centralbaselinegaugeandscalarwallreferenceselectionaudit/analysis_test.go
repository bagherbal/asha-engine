package generation2centralbaselinegaugeandscalarwallreferenceselectionaudit

import (
	"math"
	"strings"
	"testing"
)

func nearly(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func TestCentralBaselineGaugeFamilyAndActiveChoice(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.CentralBaselineUpliftInherited || !a.Inherited.CentralBaselineIdentityShift || !a.Inherited.NontrivialContentK7Uplift || !a.Inherited.DBaseBaselineSubtracted || !a.Inherited.BaselineDoesNotSelectK7OrRho72 || !a.Inherited.NoNativeScalarBaseline || a.Inherited.Verdict != StatusGate706CentralBaselineUpliftInherited {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	g := a.GaugeFamily
	if !strings.Contains(g.GeneralForm, "(R-c)P_K7") || !strings.Contains(g.GeneralForm, "(|lambda|-c)P_perp") || !nearly(g.ExpectationAtC, g.RawExpectation) || !g.GaugeInvariant || !strings.Contains(g.Verdict, StatusTotalExpectationBaselineGaugeInvariant) {
		t.Fatalf("bad gauge family: %+v", g)
	}
	active := a.ActiveChoice
	if !nearly(active.C, math.Abs(lambdaLambda12)) || !nearly(active.K7Uplift, lambdaLambda12+r3Minus1) || !nearly(active.ComplementUplift, 0) || !active.ComplementZero || !active.K7LocalResponse || !active.UniqueComplementZeroChoice || !strings.Contains(active.Verdict, StatusAbsLambdaUniqueBaselineForK7LocalUplift) {
		t.Fatalf("bad active choice: %+v", active)
	}
}

func TestAlternativesSupportAirlockAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Alternatives) != 5 || alternativeAcceptedCount(a.Alternatives) != 1 || !hasAlternative(a.Alternatives, "scalar zero-wall baseline", true) || !hasAlternative(a.Alternatives, "gauge baseline", false) || !allAlternativeExpectationsInvariant(a.Alternatives) {
		t.Fatalf("bad alternatives: %+v", a.Alternatives)
	}
	s := a.SupportLocality
	if !s.ComplementZeroForAbs || !s.ForcesCAbsLambda || !s.Gate696Consistent || !strings.Contains(s.Verdict, StatusScalarBaselineSelectionSupportLocalGauge) || !strings.Contains(s.Verdict, StatusK7UpliftFormSharperThanRawTwoPayoff) {
		t.Fatalf("bad support locality: %+v", s)
	}
	air := a.Airlock
	if !air.UsesScalarZeroWallDepth || !air.GaugeBaselineAlgebraic || !air.GaugeBaselineReversesSector || !air.CompatibleWithAirlock || air.Verdict != StatusScalarWallAirlockCompatibilityAudited {
		t.Fatalf("bad airlock: %+v", air)
	}
	m := a.Missing
	if len(m.Missing) != 5 || !strings.Contains(m.Verdict, StatusBaselineChoiceNotNativeYet) || !strings.Contains(m.Verdict, StatusNoNativeScalarBaselineReferenceSelection) || !strings.Contains(m.Verdict, StatusNoNativeK7RatherThanComplementCarriesUplift) || !strings.Contains(m.Verdict, StatusNoNativeSevenOver72Theorem) {
		t.Fatalf("bad missing: %+v", m)
	}
	f := a.Firewall
	if f.ClaimsBaselineChoiceNative || f.ClaimsNativeScalarBaselineReferenceSelection || f.ClaimsNativeK7RatherThanComplementUplift || f.ClaimsNativeBoundaryWoundUpliftTheorem || f.ClaimsNativeSevenOver72Theorem || f.ClaimsBoundaryStressDerived || f.ClaimsScalarRGMatching || f.ClaimsHiggsMass || f.ClaimsGaugeUnification || f.ClaimsFlavorDerivation || f.ClaimsCKMPMNS || f.Verdict != StatusGate707CentralBaselineGaugeBoundary {
		t.Fatalf("firewall violated: %+v", f)
	}
	res := Generation2CentralBaselineGaugeAndScalarWallReferenceSelectionAuditTheorem().Verify()
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
