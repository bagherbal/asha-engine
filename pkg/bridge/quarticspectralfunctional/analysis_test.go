package quarticspectralfunctional

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.BetaPermissionFirewallClosed {
		t.Fatal("expected beta firewall to remain closed")
	}
	if a.QuarticCollectiveBlocks != 1 || a.QuarticSpectralMoments != 10 {
		t.Fatalf("unexpected collective ledger: blocks=%d moments=%d", a.QuarticCollectiveBlocks, a.QuarticSpectralMoments)
	}
	if a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.BoundaryConstraintsDerived != 0 {
		t.Fatalf("gate must not open beta or boundary claims: beta=%d zero=%d constraints=%d", a.ContactBetaRowsAllowed, a.ContactZeroRowsProved, a.BoundaryConstraintsDerived)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatal("gate used or derived forbidden physical data")
	}
}

func TestQuarticMoments(t *testing.T) {
	p := quarticPolynomial()
	m := quarticPowerMoments(p, 4)
	checks := map[int]Rational{
		1: NewRational(71, 30),
		2: NewRational(1471, 900),
		3: NewRational(33581, 27000),
		4: NewRational(809891, 810000),
	}
	for k, want := range checks {
		if !m[k].Equal(want) {
			t.Fatalf("moment p%d = %s, want %s", k, m[k].String(), want.String())
		}
	}
	inverseTrace := p.TripleSum.Div(p.Product)
	if !inverseTrace.Equal(NewRational(2235, 271)) {
		t.Fatalf("inverse trace = %s", inverseTrace.String())
	}
}

func TestFunctionalCandidatesDoNotOpenBoundary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	for _, c := range a.FunctionalCandidates {
		if c.UsesObservedInput || !c.ExactOverQ || !c.GaloisInvariant || !c.BranchFree {
			t.Fatalf("candidate not clean: %+v", c)
		}
		if c.ConstrainsBoundary || c.MatchesKappaU1 || c.MatchesEmbeddedBoundary || c.MatchesContactWeakAngle || c.MatchesGeneratorWeakAngle {
			t.Fatalf("candidate unexpectedly constrains boundary: %s", FormatCandidate(c))
		}
		if c.BetaRowsAllowed != 0 || c.PhysicalConstantsDerived {
			t.Fatalf("candidate opened physics claim: %s", FormatCandidate(c))
		}
	}
}

func TestTheorem(t *testing.T) {
	res := CollectiveQuarticSpectralFunctionalActionLevelContributionTheorem().Run()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s :: %s", check.Name, check.Detail)
		}
	}
}
