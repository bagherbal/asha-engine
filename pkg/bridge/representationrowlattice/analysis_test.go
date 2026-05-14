package representationrowlattice

import "testing"

func TestGate203Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PreviousGate203.Gate203Inherited || !a.PreviousGate203.Gate203FailedRoutePreserved || !a.PreviousGate203.UniversalBetaSourceStillExternal || len(a.PreviousGate203.Gate201NonUniversalShapeRequirements) != 2 {
		t.Fatalf("bad Gate 203 inheritance: %s", FormatGate203(a.PreviousGate203))
	}
}

func TestExactRowFormulaMatchesKnownGate201Shapes(t *testing.T) {
	diracQ := betaRow(Statistic{Name: "Dirac fermion", Symbol: "Dirac", Coeff: R(4, 3)}, GroupRep{Symbol: "3", Dim: 3, DynkinT: R(1, 2)}, GroupRep{Symbol: "2", Dim: 2, DynkinT: R(1, 2)}, R(1, 6))
	if !diracQ.Equal(RT(R(2, 15), R(2, 1), R(4, 3))) {
		t.Fatalf("bad Dirac quark doublet row: %s", diracQ)
	}
	weylAdj := betaRow(Statistic{Name: "Weyl fermion", Symbol: "Weyl", Coeff: R(2, 3)}, GroupRep{Symbol: "1", Dim: 1, DynkinT: R(0, 1)}, GroupRep{Symbol: "3", Dim: 3, DynkinT: R(2, 1)}, R(0, 1))
	if !weylAdj.Equal(RT(R(0, 1), R(4, 3), R(0, 1))) {
		t.Fatalf("bad Weyl SU2 adjoint row: %s", weylAdj)
	}
}

func TestGrammarAndLatticeAreFiniteExactRational(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.GrammarAudit.FiniteAlphabetDeclared || !a.GrammarAudit.UnboundedEnumerationAvoided || a.GrammarAudit.CandidateRowsGenerated != len(a.Rows) || a.GrammarAudit.UniqueRows != len(a.UniqueRows) || a.GrammarAudit.ExactRationalRows != len(a.Rows) || a.GrammarAudit.CommonDenominatorLCM <= 0 {
		t.Fatalf("bad grammar: %s", FormatGrammar(a.GrammarAudit))
	}
	if !a.LatticeAudit.IntegerGridEmbedded || !a.LatticeAudit.ContainsZeroRow || !a.LatticeAudit.SemigroupOnly || !a.LatticeAudit.NoContinuousScales || !a.LatticeAudit.NoUniversalFit {
		t.Fatalf("bad lattice: %s", FormatLattice(a.LatticeAudit))
	}
}

func TestGate201ShapesHaveConditionalSupportOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.MembershipAudit.ShapesAudited != 2 || !a.MembershipAudit.AllGate201ShapesSupported || a.MembershipAudit.DirectGeneratorMatches != 2 || a.MembershipAudit.ConditionalSupportCount != 2 || !a.MembershipAudit.UniversalCompletionIgnored {
		t.Fatalf("bad membership audit: %s :: %s", FormatMembershipAudit(a.MembershipAudit), FormatMemberships(a.Memberships))
	}
	for _, m := range a.Memberships {
		if !m.Found || !m.DirectGenerator || !m.ConditionalSupport || m.FiniteDerived || !m.UniversalCompletionIgnored {
			t.Fatalf("membership over/under-claim: %s", FormatMembership(m))
		}
	}
}

func TestContactInventoryNotPromotedToRows(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.ContactInventory
	if c.ContactPartialOverlapModes != 7 || c.ContactModesHaveChargeLabels || c.ContactModesHaveGaugeRepSemantics || c.ContactModesHaveDynkinIndices || c.ContactModesHaveSpinStatistics || c.ContactModesHaveMassActivation || c.ContactModesHaveDecouplingLaw || c.CanonicalMapToRowBasisFound || c.CandidateRowsAssigned != 0 || c.FiniteHeavySectorBasisDerived {
		t.Fatalf("contact inventory overclaimed: %s", FormatContactInventory(c))
	}
}

func TestFirewallSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if !f.UniversalBetaSourceStillExternal || f.Gate201ShapesPromotedToFinitePrediction || f.UniversalBetaFitAttempted || f.ContinuousScalesSolved || f.ObservedInputsUsedForFiniteDerivation || f.ContactModesPromotedToBetaRows || f.FockGenerationPromotedToNewThreshold || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || f.AbsoluteMassPredicted || f.FiniteMatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != f.StrictNullityAfter || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("nullity changed unexpectedly: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := RepresentationRowLatticeCompletionAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
