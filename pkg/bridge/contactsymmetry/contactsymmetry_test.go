package contactsymmetry

import "testing"

func TestContactSymmetryBreakingSelectorStabilizerSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NaturalityObstructionInherited || !a.FanoAutomorphismGroupDerived || a.FanoAutomorphismGroupOrder != 168 {
		t.Fatalf("expected Gate 117 Fano naturality obstruction with |Aut|=168")
	}
	if !a.StabilizerArithmeticDerived || a.PointStabilizerOrder != 24 || a.LineStabilizerOrder != 24 || a.IncidentFlagStabilizerOrder != 8 {
		t.Fatalf("unexpected stabilizer arithmetic: point=%d line=%d flag=%d summary=%s", a.PointStabilizerOrder, a.LineStabilizerOrder, a.IncidentFlagStabilizerOrder, FormatStabilizerSummary(a.StabilizerSummary))
	}
	if !a.StabilizerReductionPossibleAfterChoice {
		t.Fatalf("expected stabilizer reduction to be possible after an external choice")
	}
	if a.CanonicalSymmetryBreakingObjectDerived || a.CanonicalFanoPointSelected || a.CanonicalFanoLineSelected || a.CanonicalFanoFlagSelected || a.CanonicalContactFanoAssignmentDerived {
		t.Fatalf("unexpected canonical selector derived")
	}
	if !a.SpectralOrderingAvailable || a.SpectralOrderingCanonicalForFano || a.SignedFanoOrientationBreaksSymmetry || a.ContactAutomorphismActionDerived || a.NaturalitySquareFormulable {
		t.Fatalf("spectral/signed diagnostics should not be promoted to a natural selector")
	}
	if a.ContactRows != 7 || a.OpenContactRowsAfter != 7 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
