package canonicalfinitediracselector

import "testing"

func TestBuildDefaultOrderOneSieveReducesButDoesNotSelect(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Summary.OrderOneDefined || !a.Summary.OrderOneSieveReduced {
		t.Fatalf("expected order-one sieve progress: %s", FormatSummary(a.Summary))
	}
	if a.Summary.CanonicalDFDerived || a.Canonical.UniqueDFSelected || a.Sieve.CanonicalBlockSelected {
		t.Fatalf("Gate 269 must not select canonical D_F: %s", FormatCanonical(a.Canonical))
	}
	if a.Sieve.AllowedComplexParameters != 2 || a.Sieve.InitialComplexParameters != 16 {
		t.Fatalf("unexpected parameter reduction: %s", FormatSieve(a.Sieve))
	}
}

func TestModeLevelConstraints(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Constraints) != 4 {
		t.Fatalf("expected four symbolic constraints, got %d", len(a.Constraints))
	}
	for _, r := range a.Constraints {
		if !r.Satisfied {
			t.Fatalf("constraint not satisfied: %s", FormatConstraint(r))
		}
	}
	if !a.Sieve.TemporalSpatialLeakageRemoved || !a.Sieve.ColorAnisotropyRemoved || !a.Sieve.OneFormsVanishForAllowedFamily {
		t.Fatalf("sieve did not record expected structural effects: %s", FormatSieve(a.Sieve))
	}
}

func TestOrderOneAllowedMomentRatioStillVaries(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Moments.Rows) != 3 {
		t.Fatalf("expected three diagnostic moment rows, got %d", len(a.Moments.Rows))
	}
	unit := a.Moments.Rows[0]
	lep := a.Moments.Rows[1]
	col := a.Moments.Rows[2]
	if unit.TraceD2 != 8 || unit.TraceD4 != 8 || unit.RawRatio != 1 {
		t.Fatalf("unexpected unit row: %s", FormatMomentRow(unit))
	}
	if lep.RawRatio == unit.RawRatio || col.RawRatio == unit.RawRatio || lep.RawRatio == col.RawRatio {
		t.Fatalf("expected distinct allowed raw ratios: %s", FormatMoments(a.Moments))
	}
	if a.Moments.RawRatioStableAcrossAllowedDF || !a.Moments.DependsOnSurvivingAmplitudes {
		t.Fatalf("moment audit should expose surviving amplitude dependence: %s", FormatMoments(a.Moments))
	}
}

func TestFirewallAndFutureObligations(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewall.ToySieveNotPromoted || a.Firewall.FiniteCorePolluted || !a.Firewall.NoConnesAlgebraImported {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
	missing := 0
	for _, o := range a.Future.Obligations {
		if o.Required && !o.Satisfied {
			missing++
		}
	}
	if missing < 7 {
		t.Fatalf("expected all future obligations missing, got %d: %s", missing, FormatFuture(a.Future))
	}
}

func TestTheoremPassesWithBridgeRequiredStatus(t *testing.T) {
	res := CanonicalFiniteDiracSelectorOrderOneSpectralTripleCompletionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
