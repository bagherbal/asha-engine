package scalarorientationsource

import "testing"

func TestBuildDefaultEtaOddScalarOrientationSource(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.WeakGauge.T3CommutesWithHighLowProjectors || !a.WeakGauge.YCommutesWithHighLowProjectors || !a.WeakGauge.WeylReflectionExchangesPlanes {
		t.Fatalf("weak gauge audit did not identify preserve/exchange structure: %s", FormatWeakGauge(a.WeakGauge))
	}
	if a.WeakGauge.EtaOddGaugeInvariantSource || a.WeakGauge.GaugeActionSelectsOrientation {
		t.Fatalf("unexpected weak eta selector: %s", FormatWeakGauge(a.WeakGauge))
	}
	if !a.Conjugation.MirrorsEtaInvolution || a.Conjugation.SelectsEtaOrientation {
		t.Fatalf("charge conjugation should exchange but not select: %s", FormatConjugation(a.Conjugation))
	}
	if a.SourceSearch.EtaOddSourceFound || a.SourceSearch.GaugeInvariantSourceFound || a.SourceSearch.CanonicalOrientationDerived {
		t.Fatalf("unexpected eta-odd source: %s", FormatSources(a.SourceSearch))
	}
	if !a.Spontaneous.OrientationInsertionPointIsolated || !a.Firewall.EtaOrientationClassifiedSpontaneous {
		t.Fatalf("expected spontaneous insertion point isolation: %s", FormatSpontaneous(a.Spontaneous))
	}
	if a.Firewall.PhysicalScalarBundleDerived || a.Firewall.ChernWeilCarrierDerived || a.Firewall.HeatKernelMatchingDerived || a.Firewall.PhysicalConstantsDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := EtaOddScalarOrientationSourceMatterPullbackSearchTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
