package contactdualpairing

import "testing"

func TestContactSourceCurrentDualPairingNaturalityObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Source.ContactSourceSelectorNoGoDerived || a.ContactRows != 7 || a.OpenContactRowsAfter != 7 {
		t.Fatalf("expected Gate 123 source-selector no-go with seven open rows")
	}
	if !a.UniformPairingConstructed || !a.UniformPairingCanonical || a.UniformPairingRank != 1 || a.UniformPairingRowsDistinguished != 0 || !a.UniformPairingRowBlind {
		t.Fatalf("uniform pairing should be canonical but row-blind: %s", FormatSummary(a.Summary))
	}
	if !a.SpectralPairingConstructed || !a.SpectralPairingCanonical || !a.SpectralPairingNonDegenerate || a.SpectralPairingRowsDistinguished != 7 || a.SpectralPairingAddsSemantics || !a.SpectralPairingDiagnosticOnly {
		t.Fatalf("spectral pairing should be nondegenerate diagnostic only: %s", FormatRows(a.Rows, 7))
	}
	if !a.CurrentDualPairingAttempted || !a.CurrentDualObstructionInherited || a.CurrentToContactMapDerived || a.CurrentFunctionalDerived || a.SourceFunctionalDerived || a.CurrentDualRowsDerived != 0 || a.CurrentDualPairingDerived {
		t.Fatalf("current dual pairing should remain unselected")
	}
	if !a.RequiresHiddenFanoChoice || a.HiddenFanoChoices != 5040 || a.NaturalRowLabelDerived || !a.ContactDualPairingNoGoDerived {
		t.Fatalf("Fano-labelled pairing should require hidden 7! choice")
	}
	if a.RepresentationCompleteRows != 0 || a.RepresentationOpenRows != 7 || a.ContactBetaRowsAllowed != 0 || a.ContactZeroRowsProved != 0 || a.ThresholdCorrectedBetaDerived {
		t.Fatalf("contact beta firewall should remain closed")
	}
	if a.ResidualNullityAfter != 3 || a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived || a.PhysicalScaleDerived {
		t.Fatalf("physical bridge leaked")
	}
}
