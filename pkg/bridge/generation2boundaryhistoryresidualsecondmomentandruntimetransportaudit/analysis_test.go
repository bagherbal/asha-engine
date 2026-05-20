package generation2boundaryhistoryresidualsecondmomentandruntimetransportaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate729SecondMomentResidualCompression(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate728.Inherited || math.Abs(a.Gate728.P_K7-7.0/72.0) > 1e-18 || !a.Gate728.AssembledRuntimeNotIndependent || !a.Gate728.PremisesNotNative {
		t.Fatalf("bad Gate728 inheritance: %+v", a.Gate728)
	}
	if math.Abs(a.Uplift.LeadingExpectation-a.Gate728.P_K7*a.Gate728.SSplit) > 1e-18 || !a.Uplift.MatchesGate700Leading {
		t.Fatalf("bad uplift operator: %+v", a.Uplift)
	}
	expectedM2 := a.Gate728.P_K7 * a.Gate728.SSplit * a.Gate728.SSplit
	if math.Abs(a.Moment.M2Wall-expectedM2) > 1e-18 || math.Abs(a.Moment.C2Wall-a.Gate728.EWall/expectedM2) > 1e-18 || !a.Moment.SecondOrderSuppressed || !a.Moment.ResidualMuchSmallerThanMoment {
		t.Fatalf("bad second moment audit: %+v", a.Moment)
	}
	if !a.Coefficients.KappaEClosestSmall || a.Coefficients.ClosestName != "kappa_e" || !a.Coefficients.NotExact {
		t.Fatalf("bad coefficient audit: %+v", a.Coefficients)
	}
	if math.Abs(a.KappaECorr.KappaEM2-kappaE*a.Moment.M2Wall) > 1e-18 || !a.KappaECorr.ImprovesRawResidual || !a.KappaECorr.NotExact || !a.KappaECorr.NotIndependentlyCertified {
		t.Fatalf("bad kappa_e correction: %+v", a.KappaECorr)
	}
}

func TestGate729RuntimePropagationNoncircularityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Variance.VarianceWall-a.Gate728.P_K7*(1-a.Gate728.P_K7)*a.Gate728.SSplit*a.Gate728.SSplit) > 1e-18 || !a.Variance.RelevantTypedScale || a.Variance.SelectedActiveCorrection {
		t.Fatalf("bad variance audit: %+v", a.Variance)
	}
	if math.Abs(a.Runtime.RawRuntimeResidual-lambdaProxyMZ*a.Gate728.L*a.Gate728.EWall) > 1e-18 || math.Abs(a.Runtime.CorrectedRuntimeResidual-lambdaProxyMZ*a.Gate728.L*a.KappaECorr.ResidualAfterCorrection) > 1e-18 || !a.Runtime.CompressionFollowsWallResidual {
		t.Fatalf("bad runtime propagation: %+v", a.Runtime)
	}
	if !a.NonCircular.DBaseContainsKappaE || !a.NonCircular.KappaEUsedAsCoefficient || a.NonCircular.IndependentTheorem || !a.NonCircular.PartiallyDependent {
		t.Fatalf("bad noncircularity: %+v", a.NonCircular)
	}
	if a.Firewall.ClaimsNativeBoundaryHistory || a.Firewall.ClaimsNativeSecondOrderExpansion || a.Firewall.ClaimsNativeScalarRuntime || a.Firewall.ClaimsHiggsMassTheorem || a.Firewall.ClaimsYukawaOperatorTheorem || a.Firewall.ClaimsCKMPMNSTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2BoundaryHistoryResidualSecondMomentAndRuntimePropagationAuditTheorem().Verify()
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
