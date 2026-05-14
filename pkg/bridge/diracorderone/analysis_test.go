package diracorderone

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.FiniteSpectralPreData || a.ContactRows != 7 || a.ContactZetaValues != 5 {
		t.Fatalf("expected inherited spectral pre-data, got contactRows=%d zeta=%d preData=%t", a.ContactRows, a.ContactZetaValues, a.FiniteSpectralPreData)
	}
	if a.SpectralTripleComplete || a.SpectralActionPrincipleReady || a.FiniteDiracSelected || a.RealStructureSelected || a.GradingSelected || a.OrderOneCalculusVerified || a.GaugeFluctuationMapDerived {
		t.Fatal("gate must not complete the spectral triple or Dirac chain")
	}
	if !a.BetaPermissionFirewallClosed || a.ThresholdBetaRows != 0 || a.BoundaryConstraintsDerived != 0 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("gate opened forbidden bridge rows: firewall=%t beta=%d constraints=%d contactBeta=%d zero=%d", a.BetaPermissionFirewallClosed, a.ThresholdBetaRows, a.BoundaryConstraintsDerived, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatal("gate used or derived forbidden physical data")
	}
}

func TestRepresentationAndStructureAudits(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.RepresentationAudit.CandidatesAudited != 5 || a.RepresentationAudit.AvailableCandidates != 4 || a.RepresentationAudit.CanonicalCandidates != 2 || a.RepresentationAudit.FaithfulTotalRepresentations != 0 || a.RepresentationAudit.NonzeroOneFormCandidates != 0 {
		t.Fatalf("unexpected representation audit: %s", FormatRepresentationAudit(a.RepresentationAudit))
	}
	if a.RealStructureAudit.CandidatesAudited != 4 || a.RealStructureAudit.CanonicalTotalCandidates != 0 || a.RealStructureAudit.KOCompatibleCandidates != 0 || a.RealStructureAudit.BranchChoicesUsed != 1 {
		t.Fatalf("unexpected real-structure audit: %s", FormatRealStructureAudit(a.RealStructureAudit))
	}
	if a.GradingAudit.CandidatesAudited != 5 || a.GradingAudit.CanonicalTotalCandidates != 0 || a.GradingAudit.OddDiracCompatible != 0 || a.GradingAudit.BranchChoicesUsed != 1 {
		t.Fatalf("unexpected grading audit: %s", FormatGradingAudit(a.GradingAudit))
	}
}

func TestOrderOneCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.OrderOneAudit.DiracCandidatesAudited != 7 || a.OrderOneAudit.OrderOneTestable != 4 || a.OrderOneAudit.OrderOneVerified != 4 || a.OrderOneAudit.OrderOneVacuous != 5 {
		t.Fatalf("unexpected order-one counts: %s", FormatOrderOneAudit(a.OrderOneAudit))
	}
	if a.OrderOneAudit.NontrivialCommutatorCandidates != 2 || a.OrderOneAudit.NontrivialOrderOneVerified != 0 || a.OrderOneAudit.GaugeFluctuationNontrivial != 0 || a.OrderOneAudit.PromotableFiniteDirac != 0 {
		t.Fatalf("nontrivial order-one path should remain blocked: %s", FormatOrderOneAudit(a.OrderOneAudit))
	}
	for _, c := range a.DiracCandidates {
		if c.UsesObservedInput || c.PhysicalConstantsDerived || c.GaugeKineticMapRows != 0 || c.ThresholdBetaRows != 0 || c.BoundaryConstraintsDerived != 0 {
			t.Fatalf("candidate made forbidden claim: %s", FormatDirac(c))
		}
		if c.PromotableToFiniteDirac {
			t.Fatalf("candidate should not be promotable: %s", FormatDirac(c))
		}
	}
}

func TestContactOrderOneIsVacuous(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	contactCount := 0
	for _, c := range a.DiracCandidates {
		if c.OrderOneVerified {
			contactCount++
			if !c.OrderOneVacuous || c.NontrivialCommutators || !c.GaugeFluctuationTrivial || c.GaugeFluctuationNontrivial {
				t.Fatalf("verified order-one candidate must be vacuous and gauge-trivial: %s", FormatDirac(c))
			}
		}
	}
	if contactCount != 4 {
		t.Fatalf("expected 4 contact spectral order-one diagnostics, got %d", contactCount)
	}
}

func TestTheorem(t *testing.T) {
	res := FiniteDiracCandidateOrderOneAxiomObstructionAuditTheorem().Run()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s :: %s", check.Name, check.Detail)
		}
	}
}
