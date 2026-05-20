package generation2level1bhiggstreeproxyestimateandvevconventionfirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate741Level1BTreeProxyEstimate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate740.Inherited || !a.Gate740.Level1BAllowed || !a.Gate740.RuntimeLambdaNotPoleMass || !a.Gate740.TreeProxyNotPoleMassTheorem {
		t.Fatalf("bad Gate740 inheritance: %+v", a.Gate740)
	}
	if math.Abs(a.Runtime.LambdaRuntimeBridge-0.12965256505047373) > 1e-15 || !a.Runtime.ClassifiedAsRuntimeQuartic || !a.Runtime.NotIndependentlyDerived || !a.Runtime.NotPoleMass {
		t.Fatalf("bad runtime value: %+v", a.Runtime)
	}
	if !a.VEV.SuppliedInput || a.VEV.NativeDerivation || !a.VEV.Convention || math.Abs(a.VEV.VGeV-246.2196508) > 1e-12 {
		t.Fatalf("bad VEV seal: %+v", a.VEV)
	}
	wantProxy := math.Sqrt(2*a.Runtime.LambdaRuntimeBridge) * a.VEV.VGeV
	if math.Abs(a.Proxy.TreeProxyGeV-wantProxy) > 1e-12 || math.Abs(a.Proxy.TreeProxyGeV-125.38000000298437) > 1e-9 || a.Proxy.PoleMassPrediction {
		t.Fatalf("bad tree proxy: %+v want %.17g", a.Proxy, wantProxy)
	}
}

func TestGate741FirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sensitivity.LinearInV || !a.Sensitivity.HalfPowerInLambda || a.Sensitivity.DeltaMOverMFromDeltaVOverV != 1 || a.Sensitivity.DeltaMOverMFromDeltaLambdaOverLambda != 0.5 {
		t.Fatalf("bad sensitivity: %+v", a.Sensitivity)
	}
	if !a.Seals.Explicit || len(a.Seals.Labels) != 12 || !a.Seals.IncludesVEVSeal || !a.Seals.IncludesConvention || !a.Seals.ProxyRemainsSealed {
		t.Fatalf("bad seal carry forward: %+v", a.Seals)
	}
	if a.Firewall.TreeProxyEqualsPoleMass || a.Firewall.RuntimeLambdaIndependentlyDerived || a.Firewall.HasPoleCorrectionTheorem || a.Firewall.HasHiggsMassTheorem || a.Firewall.Level2PredictionAllowed {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	if !a.Level.Level1BAllowed || a.Level.Level2Allowed || !a.Level.ExplicitSeals {
		t.Fatalf("bad level classification: %+v", a.Level)
	}
	res := Generation2Level1BHiggsTreeProxyEstimateAndVEVConventionFirewallAuditTheorem().Verify()
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
