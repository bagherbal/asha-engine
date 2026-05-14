package contactquotientsemantics

import "testing"

func TestCurrentSideSectorQuotientSemanticsContactRowEquivalenceSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.KernelProjectionNoGoDerived || !a.Previous.CurrentSideQuotientOnly || a.U4Dimension != 16 || a.ContactRows != 7 {
		t.Fatalf("expected Gate 127 quotient-only inheritance: %s", FormatSummary(a.Summary))
	}
	if a.NaturalCurrentQuotients != 2 || !a.CurrentSideQuotientSemanticsFound || !a.NaturalPatternsConflict || a.CurrentToContactRelationDerived {
		t.Fatalf("expected two conflicting current-side quotient semantics and no contact relation: %s", FormatSectorQuotients(a.SectorQuotients))
	}
	if !a.ContactRowEquivalenceFound || !a.RowPreservingRelationDerived || a.CanonicalSemanticRelationDerived || a.Summary.CurrentContactRelations != 0 {
		t.Fatalf("contact singleton relation should remain diagnostic only: %s", FormatContactRelations(a.ContactRelations))
	}
	if !a.HiddenAssignmentRequired || !a.ArbitraryCutoffRequired {
		t.Fatalf("non-diagnostic contact relations should require hidden assignment or cutoff: %s", FormatContactRelations(a.ContactRelations))
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
