package contactassignment

import "testing"

func TestContactSingletLeptoquarkAssignmentNaturalityPermutationObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.SectorPatternMismatch || !a.Previous.HiddenAssignmentRequired || a.ContactRows != 7 || a.CurrentPattern != "1+6" {
		t.Fatalf("expected Gate 129 inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralRowsDistinct || a.ContactDiagnosticSelectors < 3 || a.CurrentNaturalSelector {
		t.Fatalf("expected diagnostic-only spectral selectors: %s", FormatExtrema(a.MinSpectralRow, a.MaxSpectralRow, a.MedianSpectralRow))
	}
	if !a.SingletChoiceRequired || !a.PermutationRequired || a.CanonicalAssignmentDerived || a.CurrentDerivedAssignment {
		t.Fatalf("expected hidden singlet/permutation obstruction: %s", FormatCandidates(a.Candidates))
	}
	if a.Summary.MinimalSingletChoices != 7 || a.Summary.MinimalPermutationChoices != 720 || a.Summary.TotalHiddenBranchChoices < 10080 {
		t.Fatalf("expected 7 and 6! choice counts: %s", FormatSummary(a.Summary))
	}
	if !a.FanoChoiceRequired || !a.ObservedInputRejected || !a.ArbitraryOrientationRejected || a.HiddenObservedInputUsed {
		t.Fatalf("expected non-native selectors rejected: %s", FormatCandidates(a.Candidates))
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
