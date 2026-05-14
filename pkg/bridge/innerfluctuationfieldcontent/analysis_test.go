package innerfluctuationfieldcontent

import "testing"

func TestGate298GaugeContent(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gauge.GaugeContentRecovered || a.Gauge.TotalDimension != 12 {
		t.Fatalf("expected 12 gauge boson directions: %+v", a.Gauge)
	}
	if a.Gauge.UnimodularGaugeGroup != "U(1)_Y×SU(2)_L×SU(3)_C" {
		t.Fatalf("unexpected gauge group: %+v", a.Gauge)
	}
}

func TestGate298TraceNormalizationThirdSin2Path(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Trace.SU2SU3Equal || !a.Trace.ReproducesSin2 {
		t.Fatalf("trace normalization failed: %+v", a.Trace)
	}
	if RatString(a.Trace.KY) != "5/3" || RatString(a.Trace.Sin2Theta) != "3/8" {
		t.Fatalf("expected kY=5/3 and sin2=3/8: %+v", a.Trace)
	}
}

func TestGate298HiggsDoubletContent(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Higgs.SingleDoubletRecovered || a.Higgs.ComplexDoublets != 1 || a.Higgs.RealScalarDimension != 4 {
		t.Fatalf("expected one complex Higgs doublet: %+v", a.Higgs)
	}
	if len(a.Higgs.Edges) != 4 {
		t.Fatalf("expected four Yukawa edge legs supported by one doublet/conjugate: %+v", a.Higgs.Edges)
	}
}

func TestGate298Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotInventHyperchargeNormalization || !a.Firewalls.DoesNotInventYukawaMatrices || !a.Firewalls.DoesNotClaimHiggsPotential || !a.Firewalls.DoesNotClaimHeatKernel || !a.Firewalls.DoesNotActivateBGapMajorana || !a.Firewalls.DoesNotPredictMasses {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
