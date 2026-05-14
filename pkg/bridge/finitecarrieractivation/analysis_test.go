package finitecarrieractivation

import "testing"

func TestGate204Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.PreviousGate204
	if !p.Gate204Inherited || !p.Gate204ConditionalSupportPreserved || !p.RepresentationLatticeConstructed || !p.Gate201ShapesOnLattice || !p.ContactMapFailed || !p.UniversalBetaSourceStillExternal || !p.UniversalFitAvoided || !p.NoPhysicalPredictionClaim || len(p.TargetShapes) != 2 {
		t.Fatalf("bad Gate 204 inheritance: %s", FormatGate204(p))
	}
}

func TestContactModesRemainFiniteAnchorsOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.ContactModes) != 7 || !allFinitePositive(a.ContactModes) || !noContactModePromoted(a.ContactModes) {
		t.Fatalf("contact modes over/under claimed: %s", FormatContactModes(a.ContactModes, 7))
	}
}

func TestGaugeChargeSemanticsBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	g := a.GaugeCharge
	if g.ContactModesAudited != 7 || g.TargetShapesAudited != 2 || !g.FiniteOverlapCarrierAvailable || g.NativeSU3DynkinIndicesDerived || g.NativeSU2DynkinIndicesDerived || g.NativeHyperchargeDerived || g.CanonicalGaugeRepInheritance || g.CanFormDiracVectorlikeDoublet || g.CanFormWeylSU2Adjoint || g.CandidateRowsAssigned != 0 || g.GaugeChargeSemanticsComplete {
		t.Fatalf("gauge charge firewall leak: %s", FormatGaugeCharge(g))
	}
}

func TestSpinStatisticsSemanticsBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SpinStatistics
	if s.ContactModesAudited != 7 || s.LocalContinuumFieldClassDerived || s.LorentzKineticOperatorDerived || s.WeylCoefficientDerived || s.DiracCoefficientDerived || s.ScalarCoefficientDerived || s.SpinStatisticsAssigned || s.StandardBetaCoefficientSelected {
		t.Fatalf("spin-statistics firewall leak: %s", FormatSpinStatistics(s))
	}
}

func TestMassActivationSemanticsBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MassActivation
	if m.ContactModesAudited != 7 || !m.DimensionlessSpectralValuesAvailable || m.CanonicalPhysicalMassUnitDerived || m.VEVIndependentActivationDerived || m.DecouplingScaleDerived || m.ActivationPredicateDerived || m.MatchingSchemeDerived || m.ThresholdCorrectedBetaRowsAllowed {
		t.Fatalf("mass activation firewall leak: %s", FormatMassActivation(m))
	}
}

func TestCarrierActivationClassifiedAsObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Classification
	if c.RequiredPillars != 3 || c.CompletePillars != 0 || len(c.MissingPillars) != 3 || c.CarrierActivationDerived || c.ContactModesCanBeHeavyRows || c.ContactModesCanBeTargetShapes || c.Verdict != "FAILED_ROUTE" {
		t.Fatalf("carrier activation not sealed: %s", FormatClassification(c))
	}
}

func TestFirewallSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if !f.Gate204Inherited || !f.Gate204ConditionalSupportPreserved || !f.RepresentationLatticeConstructed || !f.Gate201ShapesRemainConditional || f.ContactModesPromotedToBetaRows || f.ContactModesAssignedToGate201Shapes || f.ArbitraryChargeAssignmentInserted || f.ArbitrarySpinStatisticInserted || f.ArbitraryMassScaleInserted || f.PhenomenologicalVEVUsedForActivation || f.UniversalBetaFitAttempted || f.ContinuousScalesSolved || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || f.AbsoluteMassPredicted || f.FiniteMatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != f.StrictNullityAfter || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("nullity changed unexpectedly: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := FiniteCarrierActivationContactToRowSemanticsObstructionAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
