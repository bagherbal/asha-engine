package contactsource

import "testing"

func TestContactSemanticSourceCouplingObservableSelectorSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Semantics.RowSemanticsObstructionDerived || !a.Semantics.IncidenceWeightedNoGoDerived || a.Semantics.SignedIncidenceChoiceCount != 5040 {
		t.Fatalf("expected Gate 122 row-semantics obstruction")
	}
	if !a.UniformActionSourceAvailable || !a.UniformActionSourceCanonical || !a.UniformActionSourceRowBlind || a.UniformActionSourceSelectsRows {
		t.Fatalf("uniform source should be canonical but row-blind: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralObservableConstructed || !a.SpectralObservableCanonical || a.SpectralObservableRowsDistinguished != 7 || a.SpectralObservableAddsSemantics || !a.SpectralObservableOnlyDiagnostic {
		t.Fatalf("spectral observable should be diagnostic only: %s", FormatRows(a.Rows, 7))
	}
	if !a.CurrentSourceAttempted || !a.CurrentSourceObstructionInherited || a.CurrentToContactMapDerived || a.CurrentSourceRowsDerived != 0 || a.ActionCouplingSelectorDerived || a.SemanticSourceSelectorDerived {
		t.Fatalf("source selector should remain unselected")
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
