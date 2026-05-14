package threshold

import "testing"

func TestThresholdAuditDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.ScalarActiveSpectrum) != 4 {
		t.Fatalf("expected 4 active scalar modes, got %d", len(a.ScalarActiveSpectrum))
	}
	if len(a.ContactPartialOverlap) != 7 {
		t.Fatalf("expected 7 contact partial modes, got %d", len(a.ContactPartialOverlap))
	}
	if a.BGap <= 0 {
		t.Fatalf("expected positive B gap")
	}
	if a.PhysicalMassUnitDerived {
		t.Fatalf("physical mass unit must not be marked derived")
	}
	if a.ThresholdCorrectedBetaDerived {
		t.Fatalf("threshold-corrected beta coefficients must remain open")
	}
	if a.HiddenThresholdScaleInserted || a.ObservedMassesUsed {
		t.Fatalf("threshold audit must not insert observed scales")
	}
}
