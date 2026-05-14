package heatkerneldynamicspreflight

import "testing"

func TestGate299HeatKernelFormalized(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.HeatKernel.Formalized || a.HeatKernel.AsymptoticExpansion == "" || !a.HeatKernel.RequiresAlmostCommutativeProduct {
		t.Fatalf("heat-kernel expansion not formalized: %+v", a.HeatKernel)
	}
}

func TestGate299CoefficientMap(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CoefficientMap.HiggsQuadraticMapped || !a.CoefficientMap.GaugeKineticMapped || !a.CoefficientMap.HiggsQuarticMapped {
		t.Fatalf("coefficient map incomplete: %+v", a.CoefficientMap)
	}
	if !a.CoefficientMap.OnlyFormalProjection {
		t.Fatalf("Gate 299 must remain formal only: %+v", a.CoefficientMap)
	}
}

func TestGate299NormalizationSieve(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Normalization.AllCataloged || !a.Normalization.AnyMissing || len(a.Normalization.Requirements) < 6 {
		t.Fatalf("normalization obligations not cataloged: %+v", a.Normalization)
	}
	blocking := 0
	for _, r := range a.Normalization.Requirements {
		if r.BlocksPrediction {
			blocking++
		}
	}
	if blocking != len(a.Normalization.Requirements) {
		t.Fatalf("expected all obligations to block prediction until derived/sealed: %+v", a.Normalization.Requirements)
	}
}

func TestGate299BGapPreflight(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.BGap.MajoranaEdgeDerived || a.BGap.InstantonActionDerived || a.BGap.InverseCouplingGenerated {
		t.Fatalf("B-gap dynamics must remain firewalled: %+v", a.BGap)
	}
}

func TestGate299Firewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotClaimHiggsPotential || !a.Firewalls.DoesNotClaimHiggsMassRatio || !a.Firewalls.DoesNotClaimBGapInstanton || !a.Firewalls.DoesNotInventYukawaMatrices {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
	if a.Summary.ActualDynamicsDerived {
		t.Fatalf("Gate 299 must not derive dynamics: %+v", a.Summary)
	}
}
