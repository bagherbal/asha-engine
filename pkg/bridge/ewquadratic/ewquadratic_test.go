package ewquadratic

import (
	"math"
	"testing"
)

func TestFullElectroweakQuadraticAbelianCompletion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.FullQuadraticActionFamilyTyped || !a.PositiveQuadraticFamilyExists {
		t.Fatalf("expected positive full EW quadratic family to be typed")
	}
	if a.SemisimpleRank != 3 {
		t.Fatalf("expected semisimple rank 3, got %d", a.SemisimpleRank)
	}
	if !a.Diag114ReachableInFamily || math.Abs(a.Diag114Kappa-6) > 1e-12 {
		t.Fatalf("expected diag(1,1,4) reachable at kappa=6, got %.12f", a.Diag114Kappa)
	}
	if a.Diag114SelectedByAction || a.AbelianCoefficientSelected || a.GaugeKineticHessianSelected || a.PhysicalCouplingsOrMasses {
		t.Fatalf("Gate 98 must not claim action-selected Hessian, U(1) coefficient, or physical couplings")
	}
}
