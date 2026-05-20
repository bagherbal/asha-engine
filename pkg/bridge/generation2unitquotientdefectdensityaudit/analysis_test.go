package generation2unitquotientdefectdensityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate681Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.GlobalTraceInherited || a.Inherited.H72Dimension != 72 || a.Inherited.QuotientDimension != 1 || a.Inherited.K7Rank != 7 || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Middle.Dimension != 70 || a.Augmentation.TotalDimension != 72 || a.Quotient.QuotientDimension != 1 {
		t.Fatalf("bad ladder dimensions: middle=%+v aug=%+v quotient=%+v", a.Middle, a.Augmentation, a.Quotient)
	}
	if a.Defect.IntersectionRank != 7 || a.Polarity.PositiveDim != 4 || a.Polarity.NegativeDim != 3 {
		t.Fatalf("bad K7/polarity: defect=%+v polarity=%+v", a.Defect, a.Polarity)
	}
	if !a.Density.MatchesActiveTau || math.Abs(a.Density.Density-7.0/72.0) > 1e-15 || math.Abs(a.Density.Residual-a.Inherited.ResidualGlobal) > 1e-18 {
		t.Fatalf("bad density: %+v inherited residual=%g", a.Density, a.Inherited.ResidualGlobal)
	}
	if a.Discipline.ClaimsPrimitiveDensityTheorem || a.Discipline.ClaimsTraceQuotientTheorem || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.Verdict != StatusGate681Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestPrimitiveDensityBeatsAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Alternatives) != 4 {
		t.Fatalf("expected four alternatives: %+v", a.Alternatives)
	}
	active := a.Alternatives[2]
	if active.Name != "global_defect_quotient" {
		t.Fatalf("unexpected active alternative: %+v", active)
	}
	for _, alt := range []DenominatorAlternative{a.Alternatives[0], a.Alternatives[1], a.Alternatives[3]} {
		if !(active.AbsResidual < alt.AbsResidual) {
			t.Fatalf("active residual should beat %s: active=%g alt=%g", alt.Name, active.AbsResidual, alt.AbsResidual)
		}
	}
}

func TestSacredGeometryFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.SacredFirewall.ClaimsPentagonalTheorem || a.SacredFirewall.ClaimsGoldenRatioTheorem || !a.SacredFirewall.RequiresFivefoldCarrier {
		t.Fatalf("bad sacred firewall: %+v", a.SacredFirewall)
	}
	if !strings.Contains(a.Missing.Verdict, StatusNoNativeFivefoldGoldenRatioCarrier) {
		t.Fatal("missing fivefold firewall")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2UnitQuotientDefectDensityAndPrimitiveObjectLadderAuditTheorem().Verify()
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
