package o3quotient

import "testing"

func TestBuildDefaultO3GaugeQuotientAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.ProtectedDimension != 3 || a.O3Dimension != 3 {
		t.Fatalf("protected/O3 dimensions = %d/%d, want 3/3", a.ProtectedDimension, a.O3Dimension)
	}
	if !a.CurrentDataSupportsGaugeQuotient {
		t.Fatalf("expected current protected-contact data to support gauge quotient")
	}
	if a.PhysicalOrientationSelected {
		t.Fatalf("physical orientation should not be selected by current data")
	}
	if a.FullNoOrientationTheoremProven {
		t.Fatalf("full no-orientation theorem should remain open")
	}
}
