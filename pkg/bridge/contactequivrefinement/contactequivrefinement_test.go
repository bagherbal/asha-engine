package contactequivrefinement

import "testing"

func TestContactRowEquivalenceRefinementSectorPatternMismatchObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.CurrentSideQuotientSemanticsFound || !a.Previous.NaturalPatternsConflict || a.U4Dimension != 16 || a.ContactRows != 7 {
		t.Fatalf("expected Gate 128 quotient-semantics inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.SectorPatternMismatch || !a.CurrentPatternStable || !a.ContactSingletonsStable || a.Summary.CurrentSectorPattern != "1+6" || a.Summary.ContactSingletonPattern != "1+1+1+1+1+1+1" {
		t.Fatalf("expected sector-pattern mismatch: %s", FormatSummary(a.Summary))
	}
	if a.CanonicalRefinementDerived || a.CurrentDerivedRefinement || !a.HiddenAssignmentRequired || a.Summary.MinimalHiddenChoicesPerBranch != 5040 || a.Summary.TotalHiddenBranchChoices < 10080 {
		t.Fatalf("expected hidden-assignment obstruction: %s", FormatCandidates(a.Candidates))
	}
	if !a.FanoChoiceRequired || !a.ObservedInputRejected || !a.ArbitraryCutoffRejected || a.HiddenObservedInputUsed {
		t.Fatalf("expected non-native refinements to be rejected: %s", FormatCandidates(a.Candidates))
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
