package nonuniversalrgfit

import "testing"

func TestPiSeparationExactNoGo(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.PiNoGo
	if !p.LedgerCouplingsRationalDecimals || !p.BoundaryContainsExactPi || !p.DeterminantBOneANonZero || !p.ExactClosureRequiresDeltaOnSMBetaRay || !p.SMBetaRayHasNegativeComponents || !p.RationalLatticeNonnegativeSemigroup || !p.ExactClosureImpossible {
		t.Fatalf("bad pi no-go: %s", FormatPiNoGo(p))
	}
	if p.DeterminantBOneA != "-7165690553429/176850000000" {
		t.Fatalf("unexpected determinant: %s", p.DeterminantBOneA)
	}
}

func TestSearchBasisFilteredFromGate204(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.GeneratorAudit.SourceUniqueRows != 158 || a.GeneratorAudit.SafeGenerators == 0 || a.GeneratorAudit.SafeGenerators != len(a.Generators) || !a.GeneratorAudit.NoUniversalBetaRowInserted || !a.GeneratorAudit.NoContinuousRowCoefficients {
		t.Fatalf("bad generator audit: %s", FormatGeneratorAudit(a.GeneratorAudit))
	}
	for _, g := range a.Generators {
		if !g.AnomalyCompatible || !g.LeptoquarkSealCompatible || g.DeltaB.B1.Num < 0 || g.DeltaB.B2.Num < 0 || g.DeltaB.B3.Num < 0 {
			t.Fatalf("unsafe generator admitted: %+v", g)
		}
	}
}

func TestBoundedSearchFindsNoExactSafeSolution(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Search.CombinationsAudited <= 0 || a.Search.OrderedScaleCandidates <= 0 || a.Search.ExactClosureCandidates != 0 || a.Search.ExactAsymptoticallySafeCandidates != 0 {
		t.Fatalf("unexpected exact candidate: %s", FormatSearch(a.Search))
	}
	if !a.Search.BestSafeCandidate.NoSubPlanckLandauPole || a.Search.BestSafeCandidate.ExactClosure || a.Search.BestSafeCandidate.MaxClosureResidualS <= 0 {
		t.Fatalf("bad best safe near miss: %s", FormatCandidate(a.Search.BestSafeCandidate))
	}
}

func TestFirewallsPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UniversalBetaRowInserted || f.ArbitraryRealRowCoefficientInserted || f.ExactClosureClaimed || f.ConditionalPredictionEmitted || f.ObservedLedgerUsedForFiniteCore || f.ProtonDecaySealViolated || f.ProtonLifetimeComputed || f.AbsoluteMassPredicted || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || f.FiniteMatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := NonUniversalRationalLatticeRGFitAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
