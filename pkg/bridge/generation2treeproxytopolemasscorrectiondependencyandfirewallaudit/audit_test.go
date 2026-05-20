package generation2treeproxytopolemasscorrectiondependencyandfirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate742TreeProxyToPoleDependencies(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate741.Inherited || !a.Gate741.Level1B || !a.Gate741.NotPoleMass || !a.Gate741.NotIndependentPrediction {
		t.Fatalf("bad Gate741 inheritance: %+v", a.Gate741)
	}
	if math.Abs(a.Gate741.TreeProxyGeV-125.38000000298437) > 1e-9 || math.Abs(a.Gate741.LambdaRuntimeBridge-0.12965256505047373) > 1e-15 {
		t.Fatalf("bad inherited numbers: %+v", a.Gate741)
	}
	if a.Correction.Name != "Delta_pole" || a.Correction.ValueAssigned || !a.Correction.RequiresPoleConvention || !a.Correction.RequiresExternalCorrection {
		t.Fatalf("bad correction object: %+v", a.Correction)
	}
	if !a.Ingredients.AllListed || a.Ingredients.Count != 10 || !contains(a.Ingredients.Items, "top Yukawa / top mass input") || !contains(a.Ingredients.Items, "electroweak threshold corrections") {
		t.Fatalf("bad correction ingredients: %+v", a.Ingredients)
	}
}

func TestGate742FirewallsSealsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.TreeProxyEqualsPoleMass || a.Firewall.NearNumericalProximityIsPrediction || !a.Firewall.PoleObservableNeedsLoopThreshold || !a.Firewall.TreeProxyConventionLevel {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	if !a.Seals.Explicit || a.Seals.TotalCount != 17 || !a.Seals.IncludesRGScheme || !a.Seals.IncludesPoleMass || !a.Seals.IncludesThreshold || !a.Seals.IncludesTopYukawa || !a.Seals.IncludesGauge {
		t.Fatalf("bad seal inheritance: %+v", a.Seals)
	}
	if !a.Forecast.Level1BAllowed || !a.Forecast.Level1CAllowed || !a.Forecast.Level1CDiagnosticOnly || !a.Forecast.Level1CRequiresExternal || a.Forecast.Level2Allowed {
		t.Fatalf("bad forecast levels: %+v", a.Forecast)
	}
	res := Generation2TreeProxyToPoleMassCorrectionDependencyAndFirewallAuditTheorem().Verify()
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
