package electroweakvevseal

import "testing"

func TestDimensionalOriginObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Origin.DimensionfulAnchors != 0 || !a.Origin.FiniteMatricesScaleInvariant || !a.Origin.ScalarRadiusDimensionless || !a.Origin.ScalarFundamentalClassDimensionless || !a.Origin.TopologicalTracesDimensionless {
		t.Fatalf("bad dimensional origin audit: %s", FormatOrigin(a.Origin))
	}
	if a.Origin.ElectroweakVEVDerived || a.Origin.UniqueMassUnitDerived || a.Origin.HiddenObservedScaleInserted || a.Origin.EightPiSquaredCarriesEnergyUnit {
		t.Fatalf("dimensionful scale leaked: %s", FormatOrigin(a.Origin))
	}
}

func TestEmpiricalVEVSealIsQuarantined(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Seal
	if !s.ExplicitBoundaryData || !s.Quarantined || !s.PositiveScaleRequired || !s.RequiredByDimensionalObstruction || !s.DownstreamMustDeclareSeal {
		t.Fatalf("VEV seal not properly recorded: %s", FormatVEVSeal(s))
	}
	if s.DerivedFromFiniteGeometry || s.NumericalValueSet || s.UsesObservedVEV || s.CarriesGaugeCoupling || s.CarriesTopologicalScale || s.CarriesBoundaryScale || s.UnlocksNumericalThresholds {
		t.Fatalf("VEV seal leaked forbidden data: %s", FormatVEVSeal(s))
	}
}

func TestFormalMassThresholdSymbolsOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MassLedger
	if m.FermionSectors != 4 || m.GenerationThresholdSymbols != 12 || !m.AllRequireTextureSeal || !m.AllRequireVEVSeal || !m.AllFormalThresholdsAvailable {
		t.Fatalf("formal mass threshold ledger missing: %s", FormatMassLedger(m))
	}
	if m.AnyNumericalThresholdKnown || m.AnyPhysicalMassDerivedFromFinite || m.GaugeBosonMassesAvailable || m.ScalarRadialMassNumerical {
		t.Fatalf("mass threshold firewall leaked: %s", FormatMassLedger(m))
	}
}

func TestPredicateConservativeAnswer(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.Predicate
	if !p.SharpStepAvailableConditionally || p.SharpStepDerivedNatively || !p.SmoothRegulatorSearched || p.SmoothRegulatorDerived || !p.FermionDecouplingSkeletonAvailable {
		t.Fatalf("bad threshold predicate audit: %s", FormatPredicate(p))
	}
	if p.MassOrderingKnown || p.MatchingScaleDerived || p.SchemeConventionDerived || p.ThresholdCorrectedRGDerived || p.NonUniversalDeltaBDerived {
		t.Fatalf("predicate overclaimed RG data: %s", FormatPredicate(p))
	}
}

func TestFirewallPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if !f.VEVSealInserted || !f.MassThresholdSymbolsAvailable || !f.StandardStepPredicateAdmittedAsConvention {
		t.Fatalf("conditional construction missing: %s", FormatFirewall(f))
	}
	if f.NumericalMassThresholdsAvailable || f.SmoothRegulatorNativeDerived || f.GaugeBosonThresholdsDerived || f.ThresholdBetaRowsDerived || f.ThresholdCorrectedRGFlowDerived || f.AbsoluteBoundaryScaleDerived || f.AbsoluteBoundaryCouplingDerived || f.GaugeCouplingsDerived || f.TopologicalEightPiSquaredImported || f.FiniteToContinuumScaleDerived || f.ObservedVEVImported || f.ObservedMassesImported {
		t.Fatalf("firewall leaked: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalVEVNullityBefore != 1 || f.ConditionalVEVNullityAfter != 0 || f.ConditionalThresholdNullityBefore != f.ConditionalThresholdNullityAfter {
		t.Fatalf("bad nullity ledger: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ElectroweakVEVScaleSealMassThresholdActivationFirewallTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
