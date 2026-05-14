package contactdualcurrenttarget

import "testing"

func TestContactDualCurrentTargetEnlargementSevenRowCarrierSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Previous.ContactDualPairingNoGoDerived || a.ContactRows != 7 || a.OpenContactRowsAfter != 7 {
		t.Fatalf("expected Gate 124 dual-pairing no-go with seven open rows")
	}
	if a.UniformTargetDimension != 1 || !a.UniformTargetRowBlind || a.ContactEWTargetDimension != 4 || !a.ContactEWTargetDerived || a.ContactEWTargetSevenRows {
		t.Fatalf("existing uniform/contact EW targets should not be seven-row contact targets: %s", FormatSummary(a.Summary))
	}
	if a.PatiSalamTargetDimension != 16 || !a.PatiSalamTargetDerived || a.PatiSalamTargetSevenRows || a.LeptoquarkTargetDimension != 6 || a.LeptoquarkTargetSevenRows {
		t.Fatalf("u(4)/leptoquark targets should not select seven contact rows")
	}
	if !a.SpectralSevenTargetConstructed || !a.SpectralSevenTargetCanonical || a.SpectralSevenTargetRowsDistinguished != 7 || a.SpectralSevenTargetCurrentDerived || a.SpectralSevenTargetSemantic {
		t.Fatalf("spectral seven target should be diagnostic only: %s", FormatRows(a.Rows, 7))
	}
	if !a.FanoSevenTargetConstructed || a.FanoSevenTargetCanonical || !a.FanoSevenTargetRequiresChoice || a.FanoSevenTargetHiddenChoices != 5040 || a.FanoSevenTargetCurrentDerived {
		t.Fatalf("Fano seven target should require hidden 7! choice")
	}
	if !a.AnonymousSevenTargetConstructed || !a.AnonymousSevenTargetCanonical || a.AnonymousSevenTargetRowSemantic {
		t.Fatalf("anonymous seven target should preserve cardinality only")
	}
	if a.DualCurrentTargetDerived || !a.SevenRowTargetNoGoDerived || a.NaturalSevenRowLabelsDerived || a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
