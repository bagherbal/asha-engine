package spectralaction

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.FiniteSpectralPreData || a.ContactRows != 7 || a.ContactZetaValues != 5 {
		t.Fatalf("expected exact spectral pre-data, got contactRows=%d zeta=%d preData=%t", a.ContactRows, a.ContactZetaValues, a.FiniteSpectralPreData)
	}
	if a.SpectralTripleComplete || a.SpectralActionPrincipleReady || a.FiniteDiracSelected || a.RealStructureSelected || a.GradingSelected || a.CanonicalCutoffSelected || a.GaugeFluctuationMapDerived {
		t.Fatal("gate must not complete spectral-action structure")
	}
	if !a.BetaPermissionFirewallClosed || a.ThresholdBetaRows != 0 || a.BoundaryConstraintsDerived != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("gate opened forbidden bridge rows: firewall=%t beta=%d constraints=%d contactBeta=%d zero=%d", a.BetaPermissionFirewallClosed, a.ThresholdBetaRows, a.BoundaryConstraintsDerived, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatal("gate used or derived forbidden physical data")
	}
}

func TestAxiomAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.AxiomAudit.IngredientsAudited != 11 || a.AxiomAudit.AvailableIngredients != 4 || a.AxiomAudit.CanonicalIngredients != 3 || a.AxiomAudit.MissingRequiredCanonical != 8 {
		t.Fatalf("unexpected axiom audit: %s", FormatAxiomAudit(a.AxiomAudit))
	}
	if !a.AxiomAudit.FiniteAlgebraAvailable || !a.AxiomAudit.FiniteHilbertCandidate || !a.AxiomAudit.ContactZetaLedgerAvailable || !a.AxiomAudit.ContactOverlapAvailable {
		t.Fatalf("expected available pre-data: %s", FormatAxiomAudit(a.AxiomAudit))
	}
	if a.AxiomAudit.AlgebraRepresentationReady || a.AxiomAudit.FiniteDiracSelected || a.AxiomAudit.RealStructureSelected || a.AxiomAudit.GradingSelected || a.AxiomAudit.SpectralTripleComplete {
		t.Fatalf("axiom audit should remain incomplete: %s", FormatAxiomAudit(a.AxiomAudit))
	}
}

func TestDiracCandidatesRemainDiagnostic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.DiracAudit.CandidatesAudited != 5 || a.DiracAudit.PromotableCandidates != 0 || a.DiracAudit.BetaRowsAllowed != 0 || a.DiracAudit.GaugeKineticMapsDerived != 0 || a.DiracAudit.CanonicalActionCoefficients != 0 {
		t.Fatalf("unexpected Dirac audit: %s", FormatDiracAudit(a.DiracAudit))
	}
	for _, c := range a.DiracCandidates {
		if !c.ExactOverQ || !c.SelfAdjoint || !c.FiniteSpectrum || !c.GaloisInvariant || !c.BranchFree {
			t.Fatalf("candidate lost exact spectral cleanliness: %s", FormatDiracCandidate(c))
		}
		if c.UsesObservedInput || c.UsesBranchChoice || c.PromotableToSpectralTriple || c.GaugeKineticMapDerived || c.ActionCoefficientCanonical || c.BetaRowsAllowed != 0 || c.PhysicalConstantsDerived {
			t.Fatalf("candidate made forbidden claim: %s", FormatDiracCandidate(c))
		}
	}
}

func TestActionAnsatzesRemainFormal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.ActionAudit.AnsatzesAudited != 5 || a.ActionAudit.CanonicalCoefficients != 0 || a.ActionAudit.GaugeKineticRows != 0 || a.ActionAudit.BoundaryConstraintsDerived != 0 || a.ActionAudit.ThresholdBetaRows != 0 || a.ActionAudit.PhysicalConstantsDerived {
		t.Fatalf("unexpected action audit: %s", FormatActionAudit(a.ActionAudit))
	}
	for _, c := range a.ActionAnsatzes {
		if !c.UsesZetaLedger || !c.ExactOverQ || !c.GaloisInvariant || !c.BranchFree {
			t.Fatalf("ansatz lost exact formal status: %s", FormatActionAnsatz(c))
		}
		if c.UsesObservedInput || c.UsesBranchChoice || c.CoefficientsCanonical || c.GaugeKineticRows != 0 || c.BoundaryConstraintsDerived != 0 || c.ThresholdBetaRows != 0 || c.PhysicalConstantsDerived {
			t.Fatalf("ansatz made forbidden claim: %s", FormatActionAnsatz(c))
		}
	}
}

func TestTheorem(t *testing.T) {
	res := FiniteSpectralActionPrincipleSpectralTripleConstructionAuditTheorem().Run()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s :: %s", check.Name, check.Detail)
		}
	}
}
