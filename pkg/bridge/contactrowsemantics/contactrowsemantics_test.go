package contactrowsemantics

import "testing"

func TestContactRowSemanticsLocalVariableReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ReconstructionObstructionInherited || !a.QuotientForkObstructionInherited || a.Reconstruction.AnonymousInvariantLiftPossibleRows != 5040 {
		t.Fatalf("expected Gate 121 reconstruction obstruction: %s", FormatSummary(a.Summary))
	}
	if !a.IncidenceWeightedSpectrumSearchAttempted || !a.UniformFanoIncidenceDegreeAvailable || a.FanoPointDegree != 3 || a.FanoLineSize != 3 || !a.IncidenceWeightCanonical || !a.IncidenceWeightingPreservesRows || !a.IncidenceWeightedValuesDistinct || a.IncidenceWeightingAddsRowSemantics {
		t.Fatalf("uniform incidence weighting should be canonical but semantically inert: %s", FormatSummary(a.Summary))
	}
	if !a.SignedIncidenceAttempted || a.SignedIncidenceCanonical || !a.SignedIncidenceNeedsChoice || a.SignedIncidenceChoiceCount != 5040 || a.ContactFanoAssignmentDerived {
		t.Fatalf("signed incidence should need one of 7! contact-Fano assignments")
	}
	if !a.IncidenceMomentReconstructionAttempted || !a.IncidenceMomentsRecoverSpectrum || a.IncidenceMomentsRecoverRowIdentity || a.IncidenceMomentsRecoverFanoSemantics {
		t.Fatalf("incidence moments should recover spectrum but not row/Fano semantics")
	}
	if a.LocalVariablesDerived || a.ConstraintSemanticMapDerived || a.RepresentationRowRuleDerived || !a.RowSemanticsObstructionDerived || !a.IncidenceWeightedNoGoDerived || a.RowSemanticsDerived {
		t.Fatalf("expected row-semantics obstruction")
	}
	if a.ContactRows != 7 || a.OpenContactRowsAfter != 7 || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived || a.HiddenObservedInputUsed {
		t.Fatalf("physical bridge leaked")
	}
}
