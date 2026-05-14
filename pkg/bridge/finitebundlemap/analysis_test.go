package finitebundlemap

import "testing"

func TestContactSpectralBaseDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.SpectralBase
	if !b.Commutative || !b.SemisimpleAfterComplexify || !b.SevenPointBaseDerived || b.ComplexMaximalIdeals != 7 || b.RationalPrimaryBlocks != 4 || b.BranchChoicesUsed != 0 {
		t.Fatalf("unexpected spectral base audit: %s", FormatSpectralBase(b))
	}
}

func TestProjectiveModuleRouteIsPartialSuccess(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.ModuleRoute
	if !m.FiniteLocalFieldDerived || m.CanonicalContactModules != 1 || m.ProjectiveModules != 1 || m.FreeModules != 1 || m.ContactLocalFieldAlgebras != 1 {
		t.Fatalf("contact module route should succeed exactly once: %s", FormatModuleRoute(m))
	}
	if m.PhysicalFockModules != 0 || m.PhysicalScalarModules != 0 || m.CanonicalFockScalarBundleMaps != 0 || m.PhysicalLocalBundleDerived || m.GaugeLocalFieldMaps != 0 || m.ChernWeilReadyModules != 0 {
		t.Fatalf("physical bundle should not be derived: %s", FormatModuleRoute(m))
	}
}

func TestHomologyRouteDoesNotPromoteFourCycle(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	h := a.HomologyRoute
	if !h.BooleanIncidenceComplexAvailable || !h.FanoIncidenceComplexAvailable || h.BoundaryOperatorsAvailable != 1 {
		t.Fatalf("expected finite incidence predata: %s", FormatHomologyRoute(h))
	}
	if h.ClosedFourChainsFound != 0 || h.NontrivialH4ClassesDerived != 0 || h.FiniteFundamentalClasses != 0 || h.IntegerTopologicalChargeMaps != 0 || h.HomologicalFourCycleDerived {
		t.Fatalf("homology route should not derive a fundamental cycle: %s", FormatHomologyRoute(h))
	}
}

func TestMatrixRouteHasTraceButNoQuantizedChernCharacter(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.MatrixRoute
	if m.MatrixAlgebrasAudited < 4 || m.FiniteTracesAvailable < 3 || m.CommutatorPolynomialCandidates < 3 {
		t.Fatalf("expected matrix/fuzzy predata: %s", FormatMatrixRoute(m))
	}
	if m.IntegerValuedTracePolynomials != 0 || m.TopologicallyQuantizedTraceMaps != 0 || m.ChernCharacterCandidates != 0 || m.FuzzyFourGeometryDerived || m.ChernWeilCarrierDerived {
		t.Fatalf("matrix route should not derive Chern-Weil carrier: %s", FormatMatrixRoute(m))
	}
}

func TestFirewallNarrowedButNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.ContinuousBaseRequired || !f.SevenPointContactBaseDerived || !f.ContactProjectiveModuleDerived {
		t.Fatalf("unexpected finite-locality flags: %s", FormatFirewall(f))
	}
	if f.PhysicalFockBundleDerived || f.PhysicalScalarBundleDerived || f.ChernWeilCarrierDerived || f.HeatKernelMatchingDerived || f.ThresholdCorrectedBetaDerived || f.AbsoluteCouplingPromoted || f.PhysicalConstantsDerived {
		t.Fatalf("firewall should remain closed: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
