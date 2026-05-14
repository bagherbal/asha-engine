package fourcyclechernweil

import "testing"

func TestBuildDefaultFourCycleChernWeilSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.CandidatesAudited < 8 || a.Search.ExactFiniteCandidates < 7 {
		t.Fatalf("expected broad finite candidate audit: %s", FormatSearch(a.Search))
	}
	if a.PreviousGate180.Firewall.StrictNullityAfter != 3 || a.PreviousGate180.Firewall.ConditionalNullityAfter != 2 {
		t.Fatalf("unexpected Gate 180 nullity input: %s", a.PreviousGate180.Firewall.Verdict)
	}
}

func TestNoBoundarylessFundamentalFourCycleDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.BoundarylessCycleCandidates != 0 || a.Search.CanonicalFundamentalClasses != 0 || a.Search.IntegrationFunctionals != 0 || a.Search.FiniteFourCycleDerived {
		t.Fatalf("four-cycle should not be derived: %s", FormatSearch(a.Search))
	}
	if !a.Homology.GradeFourDataAvailable || !a.Homology.DimensionFourVectorSpacesExist || !a.Homology.BoundaryOperatorDerived {
		t.Fatalf("expected grade-four/4D/boundary predata to be recorded: %s", FormatHomology(a.Homology))
	}
	if a.Homology.NonzeroClosedFourCycleDerived || a.Homology.CanonicalRepresentativeDerived || a.Homology.HochschildCycleRealizesGamma {
		t.Fatalf("homology orientability should remain unpromoted: %s", FormatHomology(a.Homology))
	}
}

func TestNoChernWeilCarrierComplete(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.GaugeBundleMaps != 0 || a.Search.CurvaturePairings != 0 || a.Search.TraceNormalizations != 0 || a.Search.IntegerChargeMaps != 0 || a.Search.CompleteChernWeilCarriers != 0 {
		t.Fatalf("Chern-Weil carrier should not be complete: %s", FormatSearch(a.Search))
	}
	if !allRequirementsBlocking(a.Requirements) {
		t.Fatalf("all Chern-Weil/four-cycle requirements should block promotion: %s", FormatRequirements(a.Requirements))
	}
}

func TestTopologicalSealNotPromotedToInstantonNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.ChernWeil
	if !c.GaugeAlgebraClosed || !c.RepresentationTraceRatioClosed || !c.TopologicalSealAvailable {
		t.Fatalf("expected gauge/topological predata: %s", FormatChernWeil(c))
	}
	if c.PrincipalBundleDerived || c.ConnectionOnFourCarrierDerived || c.CurvatureTwoFormDerived || c.TracePairingDerived || c.IntegralOfTrFedgeFDerived || c.IntegerInstantonNumberDerived || c.ContinuumNormalizationPromoted {
		t.Fatalf("Chern-Weil bridge should not be promoted: %s", FormatChernWeil(c))
	}
}

func TestFirewallNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.FourCycleDerived || f.ChernWeilCarrierDerived || f.InstantonTraceBridgeDerived || f.AbsoluteCouplingPromoted || f.HeatKernelMatchingDerived || f.ThresholdCorrectedBetaDerived || f.PhysicalConstantsDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
