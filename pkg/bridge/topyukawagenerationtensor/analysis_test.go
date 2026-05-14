package topyukawagenerationtensor

import "testing"

func TestGenerationTraceAndTopology(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Trace.Formalized || !a.Trace.TreatsRPlusAsTrace || a.Trace.TreatsRPlusAsSingleTop || a.Trace.Generations != 3 || a.Trace.RPlusDecimal < 1.6 || a.Trace.NumericalYukawasInserted {
		t.Fatalf("bad generation trace: %s", FormatTrace(a.Trace))
	}
	if len(a.Topology.TauEta) != 3 || a.Topology.TauEta[0] != 2 || a.Topology.TauEta[1] != -2 || a.Topology.TauEta[2] != 1 || !a.Topology.BreaksAllThreeCapacity || !a.Topology.ScalarTraceFunctionalOnly || a.Topology.TauEtaToGenerationPullback || a.Topology.CanonicalTopEigenvectorDerived {
		t.Fatalf("bad topology: %s", FormatTopology(a.Topology))
	}
}

func TestFractionLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Lanes) != 5 {
		t.Fatalf("expected five lanes, got %d", len(a.Lanes))
	}
	if a.Lanes[0].TopFraction != 1 || a.Lanes[1].TopFraction < 0.33 || a.Lanes[1].TopFraction > 0.34 || a.Lanes[2].TopFraction < 0.11 || a.Lanes[2].TopFraction > 0.112 || a.Lanes[3].TopFraction < 0.44 || a.Lanes[3].TopFraction > 0.445 || a.Lanes[4].TopFraction != 0 {
		t.Fatalf("unexpected lane fractions: %s || %s || %s || %s || %s", FormatLane(a.Lanes[0]), FormatLane(a.Lanes[1]), FormatLane(a.Lanes[2]), FormatLane(a.Lanes[3]), FormatLane(a.Lanes[4]))
	}
	if a.Lanes[2].Canonical || !a.Lanes[2].DerivedFromTauEta || !a.Lanes[2].Ambiguous {
		t.Fatalf("tau_eta witness must remain conditional/ambiguous: %s", FormatLane(a.Lanes[2]))
	}
}

func TestRGSlopeReevaluation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Results) != 5 {
		t.Fatalf("expected five results, got %d", len(a.Results))
	}
	for _, r := range a.Results {
		if !r.Computed || !r.Perturbative || !r.LambdaPositive {
			t.Fatalf("bad RG result: %s", FormatResult(r))
		}
	}
	if !(a.Results[0].HiggsMassGeV > a.Results[3].HiggsMassGeV && a.Results[3].HiggsMassGeV > a.Results[1].HiggsMassGeV && a.Results[1].HiggsMassGeV > a.Results[2].HiggsMassGeV && a.Results[2].HiggsMassGeV > a.Results[4].HiggsMassGeV) {
		t.Fatalf("RG masses are not monotone with top fraction: %s || %s || %s || %s || %s", FormatResult(a.Results[0]), FormatResult(a.Results[1]), FormatResult(a.Results[2]), FormatResult(a.Results[3]), FormatResult(a.Results[4]))
	}
}

func TestCapacityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Capacity.Formalized || !a.Capacity.FractionalizationFlattens || a.Capacity.FractionalizationCanResolve || a.Capacity.CanonicalFractionDerived || a.Capacity.LegacyMassGeV < 331 || a.Capacity.BestFractionalMassGeV < 260 || a.Capacity.GaugeOnlyMassGeV < 150 || !a.Capacity.MinimumPossibleMassAbove125 {
		t.Fatalf("capacity should show partial flattening but unresolved tension: %s", FormatCapacity(a.Capacity))
	}
	if !a.Firewalls.NoObservedMassFitInserted || !a.Firewalls.NoObservedTopMassInserted || !a.Firewalls.NoCKMImported || !a.Firewalls.NoGenerationTextureInvented || !a.Firewalls.TauEtaNotPromotedToOperator || !a.Firewalls.NoThresholdJumpInserted || !a.Firewalls.NoTwoLoopRGExecuted || !a.Firewalls.NoPoleMassConversionInserted || !a.Firewalls.NoFinalMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.GenerationTraceFormalized || !a.Summary.TauEtaTopologyRetrieved || !a.Summary.FractionalizationAudited || !a.Summary.RGSlopeReevaluated || a.Summary.CanonicalTopFractionDerived || a.Summary.FractionalizationResolvesTension || !a.Summary.FirewallPreserved || a.Summary.FinalMassClaimed {
		t.Fatalf("summary failed: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := TopYukawaGenerationTensorSieveAmplitudeFractionalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
