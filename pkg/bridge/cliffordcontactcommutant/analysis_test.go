package cliffordcontactcommutant

import "testing"

func TestFockRankObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.FockRank
	if f.ContactPointCount != 7 || f.FockDimension != 16 || f.RemainderModuloPoints != 2 || f.UniformRankInteger {
		t.Fatalf("expected 16 mod 7 rank obstruction: %s", FormatFockRank(f))
	}
	if !f.NonUniformRanksExist || !f.NonUniformRanksRequireSelector || f.CanonicalContactPointSelectorAvailable || f.FaithfulMultiplicativeFockActionDerived {
		t.Fatalf("non-uniform ranks should require forbidden selector: %s", FormatFockRank(f))
	}
	if !f.CliffordVectorActionAvailable || f.CliffordActionMultiplicativeForC7 {
		t.Fatalf("Clifford action should remain vector preaction only: %s", FormatFockRank(f))
	}
}

func TestCartanCommutantObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Cartan
	if c.PrimitiveCartanIdempotents != 8 || c.ContactPointCount != 7 || !c.DimensionalEmbeddingPossible {
		t.Fatalf("unexpected Cartan dimensional audit: %s", FormatCartan(c))
	}
	if !c.RequiresChoiceOfCartan || c.CanonicalCartanSelectorDerived || !c.RequiresDeleteOrMergeOneCartanPoint || c.EmbeddingIntoCommutantDerived {
		t.Fatalf("Cartan embedding should be selector/gauge obstructed: %s", FormatCartan(c))
	}
}

func TestQuarticScalarEscapeIsAbstractOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.QuarticScalar
	if q.QuarticPrimaryDim != 4 || q.ScalarCarrierDim != 4 || !q.RankOneDimensionMatch || q.IntegerRankObstruction {
		t.Fatalf("expected 4->4 rank-one scalar escape hatch: %s", FormatQuarticScalar(q))
	}
	if !q.GaloisSafePrimaryIdeal || !q.AbstractRegularModuleDerived || !q.CompanionAlgebraActionAvailable {
		t.Fatalf("expected abstract quartic module data: %s", FormatQuarticScalar(q))
	}
	if q.ActsOnPhysicalHphi || q.CanonicalHphiBasisOrOperatorDerived || q.ScalarOperatorHasQuarticMinimalPolynomial || q.PhysicalScalarBundleDerived || q.ChernWeilReady {
		t.Fatalf("quartic module must not yet be promoted to physical H_phi: %s", FormatQuarticScalar(q))
	}
}

func TestCandidateSummary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.Summary
	if s.CandidatesAudited != 5 || s.DimensionCompatibleCandidates != 4 || s.SelectorBlockedCandidates != 3 || s.AlgebraHomomorphisms != 1 || s.AbstractQuarticModules != 1 {
		t.Fatalf("unexpected summary: %s", FormatSummary(s))
	}
	if s.PhysicalFockActions != 0 || s.PhysicalScalarActions != 0 || s.CompletePhysicalBundleMaps != 0 {
		t.Fatalf("no physical bundle map should be derived: %s", FormatSummary(s))
	}
}

func TestFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.ArbitraryLinearMapUsed || !f.ContactBaseInherited || !f.CliffordPreactionInherited || !f.FockRankObstructionProved || !f.CartanCommutantObstructionProved || !f.QuarticScalarAbstractModuleDerived {
		t.Fatalf("unexpected firewall predata flags: %s", FormatFirewall(f))
	}
	if f.CanonicalFockActionDerived || f.CanonicalScalarActionDerived || f.PhysicalBundleMapDerived || f.ChernWeilCarrierDerived || f.HeatKernelMatchingDerived || f.ThresholdCorrectedBetaDerived || f.AbsoluteCouplingPromoted || f.PhysicalConstantsDerived {
		t.Fatalf("firewall should remain closed: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
