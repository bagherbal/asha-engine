package anomalykinetic

import "testing"

func TestAnomalyConstrainedU1KineticHessian(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.SymmetricHessianDimension != 6 || a.DiagonalSurvivingDimension != 3 || a.OffDiagonalDimension != 3 {
		t.Fatalf("unexpected Hessian dimensions: %+v", a)
	}
	if !a.DiagonalPositive {
		t.Fatalf("expected positive diagonal trace-Gram diagnostic")
	}
	if !a.AllKnownOffDiagonalSourcesCancel {
		t.Fatalf("expected known off-diagonal sources to cancel")
	}
	if a.NonzeroOffDiagonalSurvives {
		t.Fatalf("did not expect a nonzero off-diagonal source to survive")
	}
	if a.FullU1KineticHessianDerived || a.PhysicalU1CouplingDerived || a.FineStructureDerived {
		t.Fatalf("Gate 80 must not derive physical U(1) coupling or alpha")
	}
}
