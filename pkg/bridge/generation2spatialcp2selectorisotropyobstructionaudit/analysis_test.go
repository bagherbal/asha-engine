package generation2spatialcp2selectorisotropyobstructionaudit

import "testing"

func TestGate573SpatialCP2SelectorAndSU3IsotropyObstructionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Carrier.CertifiedAsSpatialBlock || !a.Carrier.Gate572CriticalStratumMatched || a.Carrier.RealDimension != 4 {
		t.Fatalf("bad carrier: %s", FormatCarrier(a.Carrier))
	}
	if !a.Symmetry.BMinusLScalarOnWSpatial || a.Symmetry.SuppliesFurtherSelector || a.Symmetry.PreferredSpatialDirection || a.Symmetry.SU3Dimension != 8 {
		t.Fatalf("bad symmetry: %s", FormatSymmetry(a.Symmetry))
	}
	if !a.Transit.ActsTransitively || a.Transit.QuotientRealDimension != 4 || a.Transit.InvariantPointSelected || a.Transit.InvariantRankOneProjector {
		t.Fatalf("bad transitivity: %s", FormatTransitivity(a.Transit))
	}
	if !a.Selector.ClassifiesCP1CP0Split || a.Selector.ProjectorIdempotentResidual > 1e-12 || a.Selector.ProjectorTrace != 1 || a.Selector.NativeWithoutU {
		t.Fatalf("bad selector classification: %s", FormatSelector(a.Selector))
	}
	if a.Search.CandidateCount != 9 || a.Search.NativeRankOneProjectorFound || a.Search.NativeProjectivePointFound || a.Search.NativeSecondSelectorFound {
		t.Fatalf("bad native search: %s", FormatSearch(a.Search))
	}
	for _, row := range a.Search.Candidates {
		if row.NativePUProvided || row.WouldSelectCP1CP0 {
			t.Fatalf("candidate improperly promoted: %s", FormatCandidate(row))
		}
	}
	if !a.Seal.SealedNotNative || a.Seal.CommutantDimension != 5 || !a.Seal.CommMatchesGate555Formula {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	if !a.WeakPlane.BasisDependent || a.WeakPlane.NativeDerived || a.WeakPlane.WeakIsospinIdentified || a.WeakPlane.GenerationHierarchy || a.WeakPlane.YukawaTextureDerived || a.WeakPlane.CKMPMNSDerived {
		t.Fatalf("bad weak-plane relation: %s", FormatWeakPlane(a.WeakPlane))
	}
	if a.Firewall.CP2ToK7FunctorOpened || a.Firewall.ProductTimeOpened || a.Firewall.PromotedToWeakIsospin || a.Firewall.GenerationHierarchyDerived || a.Firewall.YukawaTextureDerived || a.Firewall.CKMPMNSDerived || !a.Firewall.Gate564565BoundaryPreserved {
		t.Fatalf("bad firewall: %s", FormatFirewall(a.Firewall))
	}
	if !a.Final.SpatialCP2Certified || !a.Final.SU3Transitive || a.Final.SU3InvariantPointSelected || !a.Final.GeneralTwoPlusOneSelector || a.Final.NativeRankOnePU || a.Final.PhysicalWeakFlavorEWDerived || a.Final.K7OrProductTimeOpened {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate573RankOneProjectorUtility(t *testing.T) {
	p := rankOneProjectorReal([]float64{1, 1, 0})
	if projectorIdempotentResidual(p) > 1e-12 {
		t.Fatalf("projector not idempotent: %v", p)
	}
	if trace3(p) != 1 {
		t.Fatalf("expected rank-one trace 1, got %.12g", trace3(p))
	}
}

func TestGate573Theorem(t *testing.T) {
	res := Generation2SpatialCP2SelectorAndSU3IsotropyObstructionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
