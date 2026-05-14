package heavypropagatorselfquartic

import "testing"

func TestHeavySelfQuarticLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Quartic.Formalized || a.Quartic.RawSigmaQuartic <= 0.010 || a.Quartic.RawSigmaQuartic >= 0.011 || a.Quartic.CanonicalRankOneQuartic != 1 || a.Quartic.RawLanePhysical || !a.Quartic.CanonicalLanePhysical {
		t.Fatalf("bad quartic audit: %s", FormatQuartic(a.Quartic))
	}
}

func TestPropagatorNormalization(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Prop.Formalized || !a.Prop.CanonicalNormalization || a.Prop.HeavySupportRank != 1 || a.Prop.HeavyMetric != 1 || a.Prop.PropagatorAtThreshold != 1 || !a.Prop.RawTraceRequiresRescaling {
		t.Fatalf("bad propagator audit: %s", FormatPropagator(a.Prop))
	}
}

func TestThresholdSynthesis(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Threshold.Formalized || !a.Threshold.PortalWithinOnePercent || a.Threshold.CPortal < 0.391 || a.Threshold.CPortal > 0.392 {
		t.Fatalf("bad portal target: %s", FormatThreshold(a.Threshold))
	}
	if a.Threshold.RawSigmaLane.Viable || a.Threshold.RawSigmaLane.DeltaLambda > -9 {
		t.Fatalf("raw lane should be rejected and over-large: %s", FormatLane(a.Threshold.RawSigmaLane))
	}
	if !a.Threshold.CanonicalRankOneLane.Viable || a.Threshold.CanonicalRankOneLane.DeltaLambda < -0.098 || a.Threshold.CanonicalRankOneLane.DeltaLambda > -0.097 {
		t.Fatalf("canonical lane should match target: %s", FormatLane(a.Threshold.CanonicalRankOneLane))
	}
}

func TestTargetAlignment(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Alignment.Compared || !a.Alignment.ResolvesGate314Target || !a.Alignment.WithinOnePercent || !a.Alignment.StillConditional {
		t.Fatalf("bad alignment: %s", FormatAlignment(a.Alignment))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoFinalMassClaimed || !a.Firewalls.NoPoleMassClaimed || !a.Firewalls.NoFullSigmaPotentialClaim || !a.Firewalls.NoHeavyMassClaimed || !a.Firewalls.NoIndependentLambdaMixClaim || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := HeavyPropagatorSelfQuarticSieveThresholdNormalizationAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
