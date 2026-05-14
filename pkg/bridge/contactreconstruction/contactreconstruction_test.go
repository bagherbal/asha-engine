package contactreconstruction

import "testing"

func TestContactSpectralReconstructionInvariantToRowLiftingObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.QuotientForkObstructionInherited || !a.OrbitCollapseObstructionInherited || !a.SpectralInformationLossInherited || a.Quotient.Summary.CompatibleBijectionCount != 5040 {
		t.Fatalf("expected Gate 120 quotient fork obstruction: %s", FormatSummary(a.Summary))
	}
	if !a.WeightedSingletonLiftDerived || !a.WeightedSingletonLiftCanonical || !a.WeightedSingletonLiftPreservesRows || a.WeightedSingletonLiftFanoLike || a.WeightedSingletonLiftRepUsable {
		t.Fatalf("weighted singleton lift should preserve rows but not derive Fano-like representation data: %s", FormatSummary(a.Summary))
	}
	if !a.AnonymousInvariantLiftAttempted || !a.AnonymousInvariantLiftConstructed || a.AnonymousInvariantLiftCanonical || !a.AnonymousInvariantLiftNeedsChoice || a.AnonymousInvariantLiftPossibleRows != 5040 || a.AnonymousInvariantLiftCanonicalRows != 0 {
		t.Fatalf("anonymous lift should have 7! noncanonical row reconstructions: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralMultisetRecovered || !a.SpectralMultisetRecoversValues || a.SpectralMultisetRecoversRows || a.SpectralMultisetRecoversFanoRows {
		t.Fatalf("spectral multiset should recover values but not row/Fano semantics")
	}
	if !a.ReconstructionObstructionDerived || !a.RowLiftingAmbiguityDerived || !a.InformationChoiceNoGoDerived || a.NoLossNoChoiceLiftExists || a.InvariantToRowReconstructionDerived || a.FanoEquivariantRowLiftDerived {
		t.Fatalf("expected invariant-to-row lifting obstruction")
	}
	if a.ContactRows != 7 || a.OpenContactRowsAfter != 7 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
