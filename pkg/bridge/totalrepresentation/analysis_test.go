package totalrepresentation

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.FiniteSpectralPreData || !a.BetaPermissionFirewallClosed {
		t.Fatal("expected finite spectral pre-data with closed beta firewall")
	}
	if a.FaithfulTotalRepresentations != 0 || a.CanonicalGlueMaps != 0 || a.SpectralTripleComplete || a.FiniteDiracSelected {
		t.Fatalf("representation gate should remain obstructed: faithfulTotal=%d glue=%d triple=%t D=%t", a.FaithfulTotalRepresentations, a.CanonicalGlueMaps, a.SpectralTripleComplete, a.FiniteDiracSelected)
	}
	if a.GaugeKineticMapRows != 0 || a.ThresholdBetaRows != 0 || a.BoundaryConstraintsDerived != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("gate must not open beta/boundary rows: gauge=%d beta=%d constraints=%d contactBeta=%d zero=%d", a.GaugeKineticMapRows, a.ThresholdBetaRows, a.BoundaryConstraintsDerived, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatal("gate used or derived forbidden physical data")
	}
}

func TestCarrierAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.CarrierAudit.CandidatesAudited != 6 || a.CarrierAudit.AvailableCandidates != 5 || a.CarrierAudit.CanonicalCandidates != 4 || a.CarrierAudit.TotalHilbertCandidates != 2 || a.CarrierAudit.CanonicalTotalHilberts != 0 {
		t.Fatalf("unexpected carrier audit: %s", FormatCarrierAudit(a.CarrierAudit))
	}
}

func TestActionAndGlueAudits(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.AlgebraActionAudit.CandidatesAudited != 8 || a.AlgebraActionAudit.AvailableCandidates != 7 || a.AlgebraActionAudit.CanonicalOwnCarrierActions != 6 || a.AlgebraActionAudit.FaithfulOwnCarrierActions != 4 {
		t.Fatalf("unexpected action audit: %s", FormatActionAudit(a.AlgebraActionAudit))
	}
	if a.AlgebraActionAudit.FaithfulTotalRepresentations != 0 || a.AlgebraActionAudit.NontrivialCrossSectorActions != 0 || a.AlgebraActionAudit.NonzeroOneFormActions != 0 {
		t.Fatalf("total representation should be blocked: %s", FormatActionAudit(a.AlgebraActionAudit))
	}
	if a.GlueAudit.CandidatesAudited != 5 || a.GlueAudit.AvailableCandidates != 4 || a.GlueAudit.CanonicalCandidates != 0 || a.GlueAudit.IntertwiningCandidates != 0 || a.GlueAudit.IsometricCandidates != 0 {
		t.Fatalf("unexpected glue audit: %s", FormatGlueAudit(a.GlueAudit))
	}
}

func TestAssembliesDoNotPromote(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.AssemblyAudit.CandidatesAudited != 6 || a.AssemblyAudit.AvailableCandidates != 4 || a.AssemblyAudit.CanonicalCandidates != 3 || a.AssemblyAudit.TotalCarrierComplete != 2 {
		t.Fatalf("unexpected assembly audit: %s", FormatAssemblyAudit(a.AssemblyAudit))
	}
	for _, c := range a.Assemblies {
		if c.FaithfulTotalRepresentation || c.NontrivialSectorMixing || c.NonzeroOneForms || c.PromotableToSpectralTriple {
			t.Fatalf("assembly should not be promotable: %s", FormatAssembly(c))
		}
		if c.GaugeFluctuationMapRows != 0 || c.GaugeKineticMapRows != 0 || c.ThresholdBetaRows != 0 || c.PhysicalConstantsDerived {
			t.Fatalf("assembly opened forbidden physics: %s", FormatAssembly(c))
		}
	}
}

func TestTheorem(t *testing.T) {
	res := FiniteAlgebraTotalHilbertRepresentationFaithfulActionObstructionAuditTheorem().Run()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s :: %s", check.Name, check.Detail)
		}
	}
}
