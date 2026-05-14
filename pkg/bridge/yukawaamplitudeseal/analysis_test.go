package yukawaamplitudeseal

import "testing"

func TestBuildDefaultEmpiricalYukawaSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Seal.ExplicitBoundaryData || !a.Seal.Quarantined || a.Seal.MatrixCount != 4 {
		t.Fatalf("seal not properly recorded: %s", FormatSeal(a.Seal))
	}
	if a.Seal.ComplexEntriesTotal != 36 || a.Seal.RawRealParametersTotal != 72 {
		t.Fatalf("unexpected texture parameter count: %s", FormatSeal(a.Seal))
	}
	if a.Seal.DerivedFromFiniteGeometry || a.Seal.UsesObservedMassTargets || a.Seal.CarriesGaugeCoupling || a.Seal.CarriesHiggsVEVAmplitude || a.Seal.CarriesPhysicalMassScale {
		t.Fatalf("seal leaked forbidden data: %s", FormatSeal(a.Seal))
	}
}

func TestFormalSVDDoesNotDeriveMasses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SVD.AllFourSVDsExist || !a.SVD.AllConditionalOnSeal || !a.SVD.WeakToMassBasisFormalized {
		t.Fatalf("SVD not formalized: %s", FormatSVD(a.SVD))
	}
	if a.SVD.AnyNumericalSVDRun || a.SVD.AnySingularValueDerived || a.SVD.AnyMassDerived || !a.SVD.VEVRequiredButNotDerived {
		t.Fatalf("SVD firewall leaked: %s", FormatSVD(a.SVD))
	}
	for _, ch := range a.SVD.Channels {
		if !ch.ExistsForAnyComplexMatrix || !ch.SingularValuesNonNegative || !ch.ZeroSingularValuesAllowed || !ch.NonUniqueUnderDegeneracy || ch.MassesDerived {
			t.Fatalf("bad SVD channel: %s", FormatSVDChannel(ch))
		}
	}
}

func TestCKMPMNSAreFormalMisalignments(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Mixing) != 2 {
		t.Fatalf("expected CKM and PMNS audits, got %d", len(a.Mixing))
	}
	for _, m := range a.Mixing {
		if !m.UnitaryByConstruction || !m.RotatesChargedCurrent || !m.RequiresEmpiricalTexture {
			t.Fatalf("mixing matrix not formalized: %s", FormatMixing(a.Mixing))
		}
		if m.AnglesDerived || m.PhasesDerived || m.NumericalEntriesDerived {
			t.Fatalf("mixing numeric data leaked: %s", FormatMixing(a.Mixing))
		}
	}
	if !a.ChargedCurrent.GenerationMixingAppearsOnlyInChargedCurrent || !a.ChargedCurrent.NeutralCurrentsRemainGenerationDiagonal || a.ChargedCurrent.MixingCoefficientsNumericallyDerived {
		t.Fatalf("charged-current audit failed: %s", FormatChargedCurrent(a.ChargedCurrent))
	}
}

func TestFirewallPreserved(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	fw := a.Firewall
	if !fw.EmpiricalTextureSealInserted || !fw.YukawaMatricesAvailableConditionally || !fw.SVDMassBasisAvailableConditionally || !fw.CKMPMNSAvailableConditionally {
		t.Fatalf("conditional construction missing: %s", FormatFirewall(fw))
	}
	if fw.PhysicalMassesDerived || fw.HiggsVEVAmplitudeDerived || fw.ObservedMassRatiosImported || fw.CabibboAngleImported || fw.GenerationTextureDerivedFromFiniteData || fw.ThresholdBetaRowsDerived || fw.ThresholdMassesAvailable || fw.GaugeCouplingsDerived || fw.AbsoluteBoundaryScaleDerived || fw.TopologicalEightPiSquaredImported || fw.FiniteToContinuumScaleDerived {
		t.Fatalf("firewall leaked: %s", FormatFirewall(fw))
	}
	if fw.StrictNullityBefore != 3 || fw.StrictNullityAfter != 3 || fw.ConditionalTextureNullityBefore != 1 || fw.ConditionalTextureNullityAfter != 0 {
		t.Fatalf("unexpected nullity accounting: %s", FormatFirewall(fw))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := SpontaneousYukawaAmplitudeSealEmpiricalTextureAxiomFirewallTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
