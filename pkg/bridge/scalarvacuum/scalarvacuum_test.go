package scalarvacuum

import "testing"

func TestScalarVacuumOrientationFiniteMinimizerSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.LowPairSelected {
		t.Fatalf("expected finite scalar response to select the low active pair")
	}
	if !a.DiagnosticVacuumIsMinimizer {
		t.Fatalf("expected Gate 84 diagnostic vacuum to be a constrained minimizer")
	}
	if a.FiniteVacuumOrientationDerived {
		t.Fatalf("exact vector orientation should remain open")
	}
	if a.ResidualPhaseFreedomDimension != 1 {
		t.Fatalf("expected residual S^1 phase freedom, got dimension %d", a.ResidualPhaseFreedomDimension)
	}
}
