package contactzeta

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.BetaPermissionFirewallClosed {
		t.Fatal("expected beta firewall to remain closed")
	}
	if a.ContactZetaValues != 5 || a.PositiveNonzeroSpectrumRows != 7 || a.FiniteZetaPoleCount != 0 || a.AnalyticContinuationNeeded {
		t.Fatalf("unexpected zeta ledger: values=%d positive=%d poles=%d analytic=%t", a.ContactZetaValues, a.PositiveNonzeroSpectrumRows, a.FiniteZetaPoleCount, a.AnalyticContinuationNeeded)
	}
	if a.SpectralTripleComplete || a.CanonicalCutoffSelected || a.GaugeKineticMapRows != 0 {
		t.Fatalf("gate must not create missing spectral-action structure: triple=%t cutoff=%t gaugeMap=%d", a.SpectralTripleComplete, a.CanonicalCutoffSelected, a.GaugeKineticMapRows)
	}
	if a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.BoundaryConstraintsDerived != 0 {
		t.Fatalf("gate must not open beta or boundary claims: beta=%d zero=%d constraints=%d", a.ContactBetaRowsAllowed, a.ContactZeroRowsProved, a.BoundaryConstraintsDerived)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatal("gate used or derived forbidden physical data")
	}
}

func TestContactZetaValues(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	checks := map[int]Rational{
		0: NewRational(7, 1),
		1: NewRational(7993, 542),
		2: NewRational(10529233, 293764),
		3: NewRational(15529024549, 159220088),
		4: NewRational(24783201328945, 86297287696),
	}
	if len(a.ZetaValues) != len(checks) {
		t.Fatalf("zeta values = %d, want %d", len(a.ZetaValues), len(checks))
	}
	for _, z := range a.ZetaValues {
		want, ok := checks[z.S]
		if !ok {
			t.Fatalf("unexpected zeta index %d", z.S)
		}
		if !z.Full.Equal(want) {
			t.Fatalf("zeta(%d) = %s, want %s", z.S, z.Full.String(), want.String())
		}
		if !z.ExactOverQ || !z.GaloisInvariant || !z.BranchFree || !z.PoleFree || z.UsesBranchChoice || z.UsesObservedInput || z.RequiresRowSemantics {
			t.Fatalf("unclean zeta value: %s", FormatZetaValue(z))
		}
	}
}

func TestQuarticInverseMoments(t *testing.T) {
	prev, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	checks := map[int]Rational{
		0: NewRational(4, 1),
		1: NewRational(2235, 271),
		2: NewRational(1512333, 73441),
		3: NewRational(1177369209, 19902511),
		4: NewRational(998467775217, 5393580481),
	}
	for _, z := range prev.ZetaValues {
		want := checks[z.S]
		if !z.QuarticPart.Equal(want) {
			t.Fatalf("quartic inverse p%d = %s, want %s", z.S, z.QuarticPart.String(), want.String())
		}
	}
}

func TestActionCandidatesDoNotOpenBoundaryOrBeta(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	for _, c := range a.ActionCandidates {
		if c.UsesObservedInput || c.UsesBranchChoice || c.RequiresRowSemantics || !c.ExactOverQ || !c.GaloisInvariant || !c.BranchFree {
			t.Fatalf("candidate not clean: %+v", c)
		}
		if !c.RequiresSpectralTriple || !c.RequiresCutoffFunction || c.CoefficientCanonical {
			t.Fatalf("candidate should remain formal until Gate 163: %s", FormatActionCandidate(c))
		}
		if c.ConstrainsBoundary || c.MatchesKappaU1 || c.MatchesEmbeddedBoundary || c.MatchesContactWeakAngle || c.MatchesGeneratorWeakAngle {
			t.Fatalf("candidate unexpectedly constrains boundary: %s", FormatActionCandidate(c))
		}
		if c.BetaRowsAllowed != 0 || c.PhysicalConstantsDerived {
			t.Fatalf("candidate opened physics claim: %s", FormatActionCandidate(c))
		}
	}
}

func TestTheorem(t *testing.T) {
	res := FiniteContactSpectralZetaRegularizationSevenRootActionFunctionalAuditTheorem().Run()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s :: %s", check.Name, check.Detail)
		}
	}
}
