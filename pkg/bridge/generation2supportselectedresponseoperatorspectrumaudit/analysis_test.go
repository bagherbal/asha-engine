package generation2supportselectedresponseoperatorspectrumaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate688Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.FactorizationFirewallInherited || a.Inherited.SelectedProjector != "P_K7" || !a.Inherited.PriorCentralScalarFirewall {
		t.Fatalf("bad Gate687 inheritance: %+v", a.Inherited)
	}
	if a.Response.Operator != "R_split = S_split P_K7" || !a.Response.ResponseInEndH72 || !a.Response.SupportSelectedBeforeTracing {
		t.Fatalf("bad response operator definition: %+v", a.Response)
	}
	if a.Discipline.ClaimsProjectorActivation || a.Discipline.ClaimsNativeSevenOver72 {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestOperatorSpectrum(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Spectrum.EigenvalueOnK7-auditedSSplit) > tolerance || a.Spectrum.K7Multiplicity != 7 || a.Spectrum.ZeroMultiplicity != 65 {
		t.Fatalf("bad spectrum: %+v", a.Spectrum)
	}
	if a.Spectrum.RankIfSSplitNonzero != 7 || !a.Spectrum.SSplitNonzero || a.Spectrum.SpectrumDimensionSum != 72 {
		t.Fatalf("bad spectral rank/dimension: %+v", a.Spectrum)
	}
	if !strings.Contains(a.Spectrum.Verdict, StatusSSplitEigenvalueOnK7Support) {
		t.Fatalf("missing eigenvalue verdict: %+v", a.Spectrum)
	}
}

func TestTracePowerCable(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range a.TraceCable.Powers {
		expected := float64(7) * math.Pow(auditedSSplit, float64(p.Power))
		if math.Abs(p.Trace-expected) > tolerance {
			t.Fatalf("bad trace power n=%d: got %.18g want %.18g", p.Power, p.Trace, expected)
		}
		if math.Abs(p.NormalizedTrace-expected/72.0) > tolerance {
			t.Fatalf("bad normalized trace power n=%d: %+v", p.Power, p)
		}
	}
	if math.Abs(a.TraceCable.FirstTraceResidual-activeTraceResidual) > tolerance {
		t.Fatalf("bad first trace residual: %+v", a.TraceCable)
	}
}

func TestFirstTraceSelectionAndSupportInvariance(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LinearResponse.UsesFirstOrdinaryTrace || a.LinearResponse.UsesSecondTrace || a.LinearResponse.UsesFrobeniusNorm || a.LinearResponse.UsesHodgeSignedTrace {
		t.Fatalf("active response should be first ordinary trace only: %+v", a.LinearResponse)
	}
	if !a.Support.PBRSplitEqualsRSplit || !a.Support.PGRSplitEqualsRSplit || !a.Support.ImageInIntersectionCarrier {
		t.Fatalf("response operator should remain support-invariant: %+v", a.Support)
	}
	if !a.Support.SelectedIndependentlyOfTrace {
		t.Fatalf("support selection should remain separate from trace scalarization: %+v", a.Support)
	}
}

func TestSpectralDegeneracyRequiresSupport(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Degeneracy.Candidates) != 3 || !a.Degeneracy.AllShareSpectrumAndTrace {
		t.Fatalf("expected three rank-seven responses sharing spectrum and trace: %+v", a.Degeneracy)
	}
	if a.Degeneracy.SpectrumSelectsK7 || a.Degeneracy.TraceSelectsK7 || !a.Degeneracy.SupportSelectsK7 {
		t.Fatalf("spectrum/trace alone must not select K7; support must: %+v", a.Degeneracy)
	}
	if !strings.Contains(a.Degeneracy.Verdict, StatusSpectrumTraceAloneDoNotSelectK7) {
		t.Fatalf("missing spectral degeneracy firewall: %+v", a.Degeneracy)
	}
}

func TestHodgePolarityComparison(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Hodge.K7PlusDimension != 4 || a.Hodge.K7MinusDimension != 3 || a.Hodge.OrdinaryTraceMultiplicity != 7 || a.Hodge.HodgeSignedMultiplicity != 1 {
		t.Fatalf("bad Hodge polarity comparison: %+v", a.Hodge)
	}
	if !a.Hodge.ActiveUsesOrdinaryTrace || a.Hodge.ActiveUsesSignedTrace {
		t.Fatalf("active bridge should use ordinary total support trace, not Hodge-signed trace: %+v", a.Hodge)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2SupportSelectedResponseOperatorSpectrumAuditTheorem().Verify()
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
