package generation2runtimequartictohiggsmasstranslationfirewallandrequiredinputsaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate740RuntimeQuarticTranslationFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate739.Inherited || !a.Gate739.Level1Allowed || !a.Gate739.NotIndependentPrediction || !a.Gate739.RuntimeBridgeNotHiggsMass {
		t.Fatalf("bad Gate739 inheritance: %+v", a.Gate739)
	}
	if math.Abs(a.Quartic.LambdaRuntimeBridge-0.12965256505047373) > 1e-15 || !a.Quartic.ClassifiedAsRuntimeQuartic || !a.Quartic.BridgeLayer || a.Quartic.PhysicalPoleMass {
		t.Fatalf("bad quartic classification: %+v", a.Quartic)
	}
	if !a.Proxy.RequiresV || !a.Proxy.RequiresConvention || !a.Proxy.ConventionDependent || a.Proxy.PoleMassTheorem {
		t.Fatalf("bad tree proxy relation: %+v", a.Proxy)
	}
	if math.Abs(a.Proxy.SqrtTwoLambdaFactor-math.Sqrt(2*a.Gate739.LambdaRuntimeBridge)) > 1e-16 {
		t.Fatalf("bad proxy factor: %.17g", a.Proxy.SqrtTwoLambdaFactor)
	}
	if !a.Required.AllListed || len(a.Required.Inputs) != 7 || !a.Required.HasVEV || !a.Required.HasConvention || !a.Required.HasScaleMatching || !a.Required.HasRGTransport || !a.Required.HasThresholdLoop || !a.Required.HasGaugeYukawaTop || !a.Required.HasUncertainty {
		t.Fatalf("bad required inputs: %+v", a.Required)
	}
}

func TestGate740ForecastAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.RuntimeLambdaEqualsPoleMass || a.Firewall.TreeProxyEqualsPoleMass || a.Firewall.NearAgreementIsIndependentPrediction || !a.Firewall.RuntimeLambdaNotPoleMass || !a.Firewall.TreeProxyNotPoleMassTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	if !a.Seals.Explicit || len(a.Seals.Seals) != 10 || !a.Seals.TreeProxyWouldRemainLevel1 || !a.Seals.NoSealReduction {
		t.Fatalf("bad seal carryover: %+v", a.Seals)
	}
	if !a.Forecast.Level1AAllowed || !a.Forecast.Level1BAllowed || a.Forecast.Level2Allowed {
		t.Fatalf("bad forecast levels: %+v", a.Forecast)
	}
	res := Generation2RuntimeQuarticToHiggsMassTranslationFirewallAndRequiredInputsAuditTheorem().Verify()
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
