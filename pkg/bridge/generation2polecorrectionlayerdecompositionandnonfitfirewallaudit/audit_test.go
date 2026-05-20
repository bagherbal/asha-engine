package generation2polecorrectionlayerdecompositionandnonfitfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate744LayerDecompositionAndSymbolicDelta(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate743.Inherited || a.Gate743.DeltaPoleObject != "Delta_pole" || a.Gate743.DeltaPoleValueAssigned || !a.Gate743.FullCorrectionPackageDefined || !a.Gate743.Level1CAllowed || !a.Gate743.Level1CDiagnosticOnly || a.Gate743.Level2Allowed {
		t.Fatalf("bad Gate743 inheritance: %+v", a.Gate743)
	}
	if !NearlyEqual(a.Gate743.TreeProxyGeV, 125.38000000298437, 1e-9) {
		t.Fatalf("bad inherited tree proxy: %.17g", a.Gate743.TreeProxyGeV)
	}
	if a.DeltaPole.Name != "Delta_pole" || a.DeltaPole.Expression != "m_H_pole - m_H_tree_proxy" || a.DeltaPole.ValueAssigned || a.DeltaPole.ExternalObservableSupplied || a.DeltaPole.CorrectionPackageSupplied {
		t.Fatalf("Delta_pole should remain symbolic: %+v", a.DeltaPole)
	}
	if a.Decomposition.Count != 6 || !a.Decomposition.AllRequired || a.Decomposition.AnyNativeDerived || a.Decomposition.CompressibleToFit {
		t.Fatalf("bad decomposition: %+v", a.Decomposition)
	}
	for _, want := range CorrectionLayerNames {
		found := false
		for _, layer := range a.Decomposition.Layers {
			if layer.Name == want {
				found = true
				if !layer.Required || layer.NativeDerived || layer.Role == "" {
					t.Fatalf("bad layer %+v", layer)
				}
			}
		}
		if !found {
			t.Fatalf("missing layer %s", want)
		}
	}
}

func TestGate744MinimalityForecastAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Minimality.Minimal || !a.Minimality.AllRequired || a.Minimality.Count != 6 {
		t.Fatalf("bad minimality: %+v", a.Minimality)
	}
	for _, item := range a.Minimality.Items {
		if !item.Required || item.Layer == "" || item.RemovalEffect == "" {
			t.Fatalf("bad minimality item: %+v", item)
		}
	}
	if a.NonFit.ObservedMinusProxyIsDerivedTheorem || !a.NonFit.ExternalDiagnosticAllowed || a.NonFit.ExternalDiagnosticIsPrediction || !a.NonFit.DeltaPoleKeptLayered || !a.NonFit.SingleFittedNumberLosesTypeInfo {
		t.Fatalf("bad non-fit firewall: %+v", a.NonFit)
	}
	if !a.Classification.Level1BAllowed || !a.Classification.Level1CAllowed || a.Classification.Level2Allowed || !strings.Contains(a.Classification.DeltaPoleStatus, "multi-layer") {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	res := Generation2PoleCorrectionLayerDecompositionAndNonFitFirewallAuditTheorem().Verify()
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
			t.Fatalf("missing note %s", want)
		}
	}
}
