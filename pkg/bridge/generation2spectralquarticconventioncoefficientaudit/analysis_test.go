package generation2spectralquarticconventioncoefficientaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate619Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate618Inherited {
		t.Fatalf("bad inherited verdict %s", a.Inherited.Verdict)
	}
	if len(a.Conventions) < 6 {
		t.Fatalf("expected convention family")
	}
	if len(a.Diagnostics) != 2 {
		t.Fatalf("expected MZ and Lambda12 diagnostics")
	}
	l12 := a.Diagnostics[1]
	if math.Abs(l12.BOverA2-0.3330764110541872) > 1e-15 {
		t.Fatalf("bad b/a^2 %.18g", l12.BOverA2)
	}
	if !(l12.CLambdaRequiredRuntime < 0) {
		t.Fatalf("expected negative required c_lambda for negative runtime lambda")
	}
	if !a.SignAudit.BOverA2NonNegative || a.SignAudit.DirectPositiveBoundaryPossible {
		t.Fatalf("bad sign audit %+v", a.SignAudit)
	}
	if !a.StressImpact.UsesLambdaRuntimeShadow || a.StressImpact.CanUseLambdaCanon {
		t.Fatalf("bad stress impact %+v", a.StressImpact)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2SpectralQuarticConventionCoefficientAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusNoCLambdaValue, StatusNegativeRuntimeNotDirect, StatusStressRuntimeShadow, StatusGate619Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
