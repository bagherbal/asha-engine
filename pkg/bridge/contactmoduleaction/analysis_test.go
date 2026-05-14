package contactmoduleaction

import "testing"

func TestCliffordSpinorPreactionButNoContactAlgebraModule(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.CliffordSpinor
	if !c.CliffordBookkeepingAvailable || !c.K7VectorActionCanonical || !c.ActionOnSpinorsCanonical || !c.LinearK7ToEndFockMapDerived {
		t.Fatalf("expected canonical Clifford-spinor preaction: %s", FormatCliffordSpinor(c))
	}
	if c.MultiplicativeContactAlgebraHom || c.CommutativeSpectralIdempotentAction || c.OmegaIntertwiningLawDerived || c.InducesFockProjectiveModule || c.InducesPhysicalSpinorBundle {
		t.Fatalf("Clifford vector action must not be promoted to C[Ω] module: %s", FormatCliffordSpinor(c))
	}
}

func TestQuarticScalarIdealAbstractOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.QuarticScalar
	if q.QuarticPrimaryDim != 4 || q.ScalarCarrierDim != 4 || !q.GaloisSafePrimaryIdeal || !q.AbstractRankOneModuleOverQuartic || !q.CompanionRepresentationAvailable || !q.BranchFreeQuarticBlock {
		t.Fatalf("expected abstract quartic module predata: %s", FormatQuarticScalar(q))
	}
	if q.ScalarOperatorWithQuarticMinimal || q.CanonicalHphiIdentification || q.ProjectiveScalarModuleDerived || q.PhysicalScalarBundleDerived {
		t.Fatalf("quartic ideal should not be identified with H_phi yet: %s", FormatQuarticScalar(q))
	}
}

func TestConnectionInducedActionDoesNotClose(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Connection
	if !c.ProjectedConnectionAvailable || !c.OffDiagonalBlockConnectionAvailable || !c.SecondFundamentalCurvatureAvailable || !c.CompressedConnectionCanonical || !c.AdjointActionCandidate || !c.CommutatorActionCandidate {
		t.Fatalf("expected connection-induced predata: %s", FormatConnectionAction(c))
	}
	if c.ClosesOnContactSpectralAlgebra || c.PullbackToFockDerived || c.PullbackToScalarDerived || c.FockDiracCommutatorClosed || c.GaugeCovariantModuleActionDerived {
		t.Fatalf("connection route should not close as C[Ω] action: %s", FormatConnectionAction(c))
	}
}

func TestSummaryExcludesArbitraryMapsAndNoPhysicalBundle(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Summary
	if s.CandidatesAudited != 5 || s.CanonicalPreactions < 5 || s.ArbitraryMapsUsed != 0 {
		t.Fatalf("unexpected candidate summary: %s", FormatSummary(s))
	}
	if s.PhysicalFockActions != 0 || s.PhysicalScalarActions != 0 || s.CompletePhysicalBundleMaps != 0 || s.ChernWeilReadyActions != 0 {
		t.Fatalf("physical bundle should remain open: %s", FormatSummary(s))
	}
}

func TestFirewallNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.ArbitraryLinearMapUsed || !f.ContactBaseInherited || !f.ContactRegularModuleInherited || !f.CliffordSpinorPreactionDerived || !f.QuarticAbstractScalarModuleDerived || !f.ConnectionPreactionAudited {
		t.Fatalf("unexpected firewall predata flags: %s", FormatFirewall(f))
	}
	if f.CanonicalFockActionDerived || f.CanonicalScalarActionDerived || f.PhysicalBundleMapDerived || f.ChernWeilCarrierDerived || f.HeatKernelMatchingDerived || f.ThresholdCorrectedBetaDerived || f.AbsoluteCouplingPromoted || f.PhysicalConstantsDerived {
		t.Fatalf("firewall should remain closed: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
