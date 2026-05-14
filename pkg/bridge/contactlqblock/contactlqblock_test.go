package contactlqblock

import "testing"

func TestContactLeptoquarkSixBlockS6PermutationObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.SingletChoiceRequired || !a.Previous.PermutationRequired || a.ContactRows != 7 || a.LeptoquarkRows != 6 {
		t.Fatalf("expected Gate 130 inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.S6PermutationObstruction || a.SixPermutationOrder != 720 || a.AssignmentsPerBranch != 5040 || a.TotalCurrentAssignments != 10080 {
		t.Fatalf("expected S6 obstruction counts: %s", FormatSummary(a.Summary))
	}
	if !a.SixBlockExists || !a.AnonymousBlockCanonical || a.CurrentNaturalSixOrder || a.CanonicalCurrentAssignmentDerived {
		t.Fatalf("expected anonymous canonical block but no current-natural order: %s", FormatStrategies(a.Strategies))
	}
	if !a.SpectralOrderingAvailable || !a.SpectralOrientationAmbiguous {
		t.Fatalf("expected diagnostic spectral ordering ambiguity: %s", FormatBlocks(a.Blocks, 2))
	}
	if !a.FanoChoiceRequired || !a.ObservedInputRejected || a.HiddenObservedInputUsed {
		t.Fatalf("expected Fano/observed selectors rejected: %s", FormatStrategies(a.Strategies))
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
