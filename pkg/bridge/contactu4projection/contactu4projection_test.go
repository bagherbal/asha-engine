package contactu4projection

import "testing"

func TestContactSevenRowTargetProjectionU4QuotientObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.SevenRowTargetNoGoDerived || a.ContactRows != 7 || a.OpenContactRowsAfter != 7 {
		t.Fatalf("expected Gate 125 seven-row no-go with seven open rows")
	}
	if a.U4CurrentDimension != 16 || !a.U4DecompositionCanonical || len(a.U4SectorDimensions) != 4 {
		t.Fatalf("bad u4 decomposition: %s", FormatSectors(a.Sectors))
	}
	if !a.RankSevenLinearMapsExist || a.GenericKernelDimension != 9 || !a.ContinuousProjectionFamily || a.ContinuousProjectionFreeParameters != 63 || a.CanonicalProjectionCount != 0 || a.U4ToContactProjectionDerived {
		t.Fatalf("generic rank-seven projections should exist abstractly but not be selected: %s", FormatSummary(a.Summary))
	}
	if a.CentralPlusLeptoDimension != 7 || a.BLPlusLeptoDimension != 7 || a.DimensionSevenSectorSums != 2 || a.DimensionSevenSectorSumsCanonical || !a.DimensionSevenSectorSumsWrongSemantics {
		t.Fatalf("dimension-seven sector sums should be noncanonical/wrong semantics")
	}
	if a.ColorEightToSevenQuotientDerived || a.ContactEWFourPlusThreeDerived || a.SpectralSevenIsU4Quotient || a.FanoSevenIsU4Quotient || !a.ObservedProjectionRejected {
		t.Fatalf("shortcut quotient should remain blocked")
	}
	if a.U4ToContactQuotientDerived || a.NaturalSevenRowProjectionDerived || !a.SevenRowProjectionNoGoDerived || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
