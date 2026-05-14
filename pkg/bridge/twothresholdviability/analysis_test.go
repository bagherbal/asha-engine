package twothresholdviability

import "testing"

func TestGeneratorBasisMatchesGate210(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.GeneratorAudit.InheritsGate210FilteredBasis || a.GeneratorAudit.SafeGenerators != 108 || a.GeneratorAudit.SafeGenerators != len(a.Generators) {
		t.Fatalf("bad generator audit: %s", FormatGeneratorAudit(a.GeneratorAudit))
	}
	for _, g := range a.Generators {
		if !g.AnomalyCompatible || !g.LeptoquarkSealCompatible || g.DeltaB.B1.Num < 0 || g.DeltaB.B2.Num < 0 || g.DeltaB.B3.Num < 0 {
			t.Fatalf("unsafe generator admitted: %+v", g)
		}
	}
}

func TestTwoBoundaryTargetsAudited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.TargetAudits) != 2 {
		t.Fatalf("expected two targets, got %d", len(a.TargetAudits))
	}
	top := firstTarget(a, "u_topological")
	centroid := firstTarget(a, "u_centroid")
	if top.OrderedPairsAudited != 108*107 || centroid.OrderedPairsAudited != 108*107 {
		t.Fatalf("bad pair counts: top=%d centroid=%d", top.OrderedPairsAudited, centroid.OrderedPairsAudited)
	}
	if top.InvertibleSystems == 0 || centroid.InvertibleSystems == 0 {
		t.Fatalf("expected invertible systems: top=%s centroid=%s", FormatTargetAudit(top), FormatTargetAudit(centroid))
	}
}

func TestTopologicalBranchHasViableWitnesses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	top := firstTarget(a, "u_topological")
	if top.ViablePairs == 0 || len(top.AllViableSolutions) != top.ViablePairs {
		t.Fatalf("expected viable topological witnesses: %s", FormatTargetAudit(top))
	}
	for _, s := range top.AllViableSolutions {
		if !s.Viable || !s.ExactLinearClosure || !s.ScaleOrdered || !s.DistinctThresholds || !s.SubPlanck || !s.PositiveCouplingsToMStar || !s.NoSubPlanckLandauPole || !s.AnomalyCompatible || !s.LeptoquarkSealCompatible {
			t.Fatalf("bad viable solution: %s", FormatSolution(s))
		}
		if s.MatchesBGapOrContactData {
			t.Fatalf("contact data promoted unexpectedly: %s", FormatSolution(s))
		}
	}
}

func TestCentroidBranchRemainsFiltered(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	centroid := firstTarget(a, "u_centroid")
	if centroid.OrderedPairsAudited == 0 {
		t.Fatalf("centroid target was not audited")
	}
	if centroid.ViablePairs != 0 {
		t.Fatalf("unexpected centroid viable pair: %s", FormatTargetAudit(centroid))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if !f.Gate210Inherited || !f.LeptoquarkDynamicsSealInherited || !f.EmpiricalLedgerQuarantined || f.ObservedLedgerUsedForFiniteCore || f.UniversalBetaRowInserted || f.ArbitraryRealRowCoefficientInserted || f.PhysicalPredictionClaimed || f.AbsoluteMassDerivedFromFiniteCore || f.ProtonDecaySealViolated || f.ProtonLifetimeComputed || f.MatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := TwoThresholdRationalLatticeViabilityFilterTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
