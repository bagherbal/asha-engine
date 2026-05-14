package complexweightspacekernel

import "testing"

func TestGate251ComplexCapacityButNoNativeHermitianMatrices(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Summary.Complex8VKnown || !a.Summary.HermitianWeightCapacity || !a.Summary.OddComplexKernelCapacity {
		t.Fatalf("expected complex/Hermitian/odd-kernel capacity, got %+v", a.Summary)
	}
	if a.Summary.NativeHermitianMatrices || a.Summary.ComplexNeutralKernelDerived || a.Summary.NeutralKernelDim3 {
		t.Fatalf("unexpected derived Q8vC/kernel: %+v", a.Summary)
	}
	if !a.Summary.ComplexTrialityArena || a.Summary.CanonicalTrialityMap || a.Summary.RealStructureCompatible {
		t.Fatalf("expected triality arena but no canonical J-compatible map: %+v", a.Summary)
	}
	if a.Summary.VTauConstructed || a.Summary.TrialityUnblocked || a.Summary.YukawaTextureDerived {
		t.Fatalf("unexpected downstream derivation: %+v", a.Summary)
	}
}
