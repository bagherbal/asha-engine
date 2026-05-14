package grandnormalizationsieve

import "testing"

func TestGate300Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Input.HeatKernelExpansionFormalized || !a.Input.A2ScalarQuadraticChannel || !a.Input.A4ScalarKineticChannel || !a.Input.A4GaugeKineticChannel || !a.Input.A4ScalarQuarticChannel {
		t.Fatalf("bad Gate 299 inheritance: %s", FormatInput(a.Input))
	}
	if a.Input.RawTraceRatioNumerator != 1197 || a.Input.RawTraceRatioDenominator != 4624 || a.Input.PhysicalDynamicsDerived {
		t.Fatalf("raw ratio inheritance overpromoted: %s", FormatInput(a.Input))
	}
}

func TestKineticIsolationClassifiers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Kinetic.SeparatesKineticPotential || !a.Kinetic.RejectsVacuumTerms || !a.Kinetic.RejectsBGapMassInsertion {
		t.Fatalf("kinetic isolation failed: %s", FormatKinetic(a.Kinetic))
	}
	var scalarKinetic, gaugeKinetic, quartic bool
	for _, c := range a.Kinetic.ClassifierRules {
		if c.AcceptedForZH && c.DerivativeOrder == 2 && c.ScalarPower == 2 && c.GaugeCurvaturePower == 0 {
			scalarKinetic = true
		}
		if c.AcceptedForGaugeNorm && c.GaugeCurvaturePower == 2 && c.ScalarPower == 0 {
			gaugeKinetic = true
		}
		if c.AcceptedForPotential && c.ScalarPower == 4 && c.DerivativeOrder == 0 {
			quartic = true
		}
	}
	if !scalarKinetic || !gaugeKinetic || !quartic {
		t.Fatalf("missing classifiers: %s", FormatKinetic(a.Kinetic))
	}
}

func TestZHAlgorithmFormalizedButNotNumerical(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.ZH.AlgorithmValid || a.ZH.Rescaling != "H_raw = H_phys / sqrt(Z_H)" {
		t.Fatalf("ZH algorithm not formalized: %s", FormatZH(a.ZH))
	}
	if a.ZH.NumericalZHComputed || a.ZH.PositiveZHProved {
		t.Fatalf("Gate 300 must not claim numerical or positive ZH: %s", FormatZH(a.ZH))
	}
}

func TestMassQuarticMapDoesNotPromote1197Over4624(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Rescaling.AlgorithmFormalized || len(a.Rescaling.Coefficients) != 2 {
		t.Fatalf("rescaling map incomplete: %s", FormatRescaling(a.Rescaling))
	}
	if a.Rescaling.RawRatioPromoted || a.Rescaling.PhysicalMassDerived || a.Rescaling.PhysicalQuarticDerived {
		t.Fatalf("raw trace ratio overpromoted: %s", FormatRescaling(a.Rescaling))
	}
}

func TestGaugeNormalizationMapFormalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Gauge.AlgorithmFormalized || !a.Gauge.RelativeNormalizationAudited || len(a.Gauge.GaugeGroups) != 3 {
		t.Fatalf("gauge normalization algorithm incomplete: %s", FormatGauge(a.Gauge))
	}
	if a.Gauge.AbsoluteCouplingsDerived {
		t.Fatalf("absolute couplings must remain firewalled: %s", FormatGauge(a.Gauge))
	}
}

func TestGate300Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoCutoffMomentsInserted || !a.Firewalls.NoSubtractionSchemeInvented || !a.Firewalls.NoYukawaNumbersInserted || !a.Firewalls.NoObservedMassesInserted || !a.Firewalls.NoBGapInstantonClaimed || !a.Firewalls.NoRawRatioPromotion || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failure: %s", FormatFirewalls(a.Firewalls))
	}
	if a.Summary.PhysicalDynamicsDerived {
		t.Fatalf("Gate 300 must not derive physical dynamics: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := GrandNormalizationSieveWaveFunctionRenormalizationExtractionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
