package contactu4kernel

import "testing"

func TestU4ProjectionKernelCanonicalQuotientRelationSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.SevenRowProjectionNoGoDerived || a.U4Dimension != 16 || a.TargetContactRows != 7 || a.RequiredKernelDimension != 9 {
		t.Fatalf("expected Gate 126 u4->contact projection no-go inheritance: %s", FormatSummary(a.Summary))
	}
	if !a.GenericKernelsExist || a.GrassmannKernelDimension != 63 || a.CanonicalKernelDerived {
		t.Fatalf("generic kernels should exist but no canonical kernel should be selected: %s", FormatSummary(a.Summary))
	}
	if a.NaturalNineDimensionalKernels != 2 || !a.SectorKernelAmbiguity || a.ColorBLKernelDimension != 9 || a.CentralColorKernelDimension != 9 || !a.SectorKernelsWrongSemantics || !a.CurrentSideQuotientOnly {
		t.Fatalf("expected two ambiguous current-side sector kernels: %s", FormatKernels(a.KernelCandidates))
	}
	if a.CanonicalQuotientRelation || a.ContactSemanticKernelDerived || !a.KernelProjectionNoGoDerived || a.Summary.QuotientRelationsDerived != 0 {
		t.Fatalf("no contact-semantic quotient relation should be derived: %s", FormatQuotients(a.QuotientCandidates))
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived || a.FullBetaMatchingTensorDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
