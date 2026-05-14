package gaugehessian

import "testing"

func TestGate83LeavesHessianOpen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.SymmetricHessianDim != 6 || a.DiagonalNoMixingDim != 3 || a.OffDiagonalDim != 3 {
		t.Fatalf("unexpected Hessian dimensions: symmetric=%d diagonal=%d off=%d", a.SymmetricHessianDim, a.DiagonalNoMixingDim, a.OffDiagonalDim)
	}
	if DerivedActionSlotCount(a.ActionSlots) != 3 {
		t.Fatalf("expected exactly 3 derived action slots from current data, got %d", DerivedActionSlotCount(a.ActionSlots))
	}
	if SelectedCandidateCount(a.Candidates) != 0 {
		t.Fatalf("no Hessian candidate should be derived from a second variation yet")
	}
	if a.HessianSelected || a.BoundaryCouplingFixed || a.PhysicalAlphaDerived {
		t.Fatalf("Gate 83 must not derive Hessian, boundary coupling, or alpha")
	}
}
