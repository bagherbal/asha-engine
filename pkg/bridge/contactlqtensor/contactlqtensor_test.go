package contactlqtensor

import "testing"

func TestContactLeptoquarkSlotRepresentationTensorColorDoubletObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.S6ObstructionInherited || a.LeptoquarkRows != 6 || a.Previous.SixPermutationOrder != 720 || a.Previous.AssignmentsPerBranch != 5040 {
		t.Fatalf("expected Gate 131 S6 obstruction inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.CurrentRealTensorDerived || a.CurrentLQSlots != 6 || a.ColorSlots != 3 || a.RealOrientationSlots != 2 || !a.ColorTripletSemantics || !a.RealOrientationSemantics {
		t.Fatalf("expected derived current leptoquark real tensor: %s", FormatSlots(a.Slots))
	}
	if !a.ColorWeakCountMatch || !a.ColorDoubletCountTrap || !a.SemanticBridgeMissing || a.WeakDoubletSemanticsDerived || a.HyperchargeSemanticsDerived || a.LocalFieldSemanticsDerived {
		t.Fatalf("expected color-doublet semantic obstruction: %s", FormatCandidates(a.Candidates))
	}
	if a.CurrentNaturalRepresentation || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 {
		t.Fatalf("contact representation/beta firewall should remain closed")
	}
	if a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived || a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
