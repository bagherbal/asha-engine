package tensor

import (
	"math"
	"testing"
)

func TestTensorFactorBridgeSeparatesChargeAndScalar(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TensorDimension != 64 {
		t.Fatalf("expected 16x4 tensor dimension 64, got %d", a.TensorDimension)
	}
	if a.ChargeScalarCommutatorNorm > 1e-10 {
		t.Fatalf("expected charge and scalar factors to commute, got %g", a.ChargeScalarCommutatorNorm)
	}
	if a.TensorScalarTraceResidual > 1e-10 {
		t.Fatalf("tensor scalar trace identity failed: %g", a.TensorScalarTraceResidual)
	}
	if math.Abs(a.VacuumFiberCharge) > 1e-10 || a.VacuumScalarFiberDimension != 4 {
		t.Fatalf("unexpected neutral vacuum scalar fiber: charge=%g dim=%d", a.VacuumFiberCharge, a.VacuumScalarFiberDimension)
	}
	if a.YukawaIntertwinerConstructed {
		t.Fatalf("tensor bridge must not pretend to construct a Yukawa intertwiner")
	}
}
