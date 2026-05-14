package twothresholdminimality

import "testing"

func TestGate212BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.OrderedViablePairs != 44 {
		t.Fatalf("expected 44 ordered Gate-211 witnesses, got %d", a.Summary.OrderedViablePairs)
	}
	if a.Summary.UnorderedPairClasses != 22 {
		t.Fatalf("expected 22 unordered pair classes, got %d", a.Summary.UnorderedPairClasses)
	}
	if a.Summary.CanonicalUniquePairFound {
		t.Fatalf("Gate 212 must not claim canonical uniqueness")
	}
	if !a.Summary.ThresholdSpectrumSealRequired || a.Summary.Status != StatusFailedCanonicalUniqueness {
		t.Fatalf("expected threshold-spectrum seal requirement and failed uniqueness status: %s", FormatSummary(a.Summary))
	}
}

func TestFiniteOriginObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.FiniteOrigin.CanonicalFiniteOriginMatches != 0 {
		t.Fatalf("contact/B-sector data must not canonically select a pair: %s", FormatFiniteOrigin(a.FiniteOrigin))
	}
	if a.FiniteOrigin.PairDimensionHitsSeven != 0 || a.FiniteOrigin.PairWeylDimensionHitsSeven != 0 {
		t.Fatalf("unexpected exact seven-mode dimension match: %s", FormatFiniteOrigin(a.FiniteOrigin))
	}
	if !a.FiniteOrigin.CarrierActivationSealIntact {
		t.Fatalf("carrier seal was not preserved")
	}
}

func TestParentageObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Parentage.CompleteParentageDerived != 0 || a.Parentage.UnifiedParentGaugeImported || a.Parentage.ThresholdSplittingRuleDerived {
		t.Fatalf("parentage firewall leak: %s", FormatParentage(a.Parentage))
	}
	if a.Parentage.PairClassesAudited != 22 {
		t.Fatalf("expected 22 pair classes in parentage audit: %s", FormatParentage(a.Parentage))
	}
}

func TestFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.ContactModesPromotedToCarriers || f.BGapPromotedToMass || f.SU5OrPatiSalamGaugeImported || f.MatchingCorrectionsDerived || f.PhysicalPredictionClaimed || f.ProtonLifetimeComputed || f.UniqueThresholdSpectrumClaimed {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := TwoThresholdSolutionMinimalityFiniteOriginParentageAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
