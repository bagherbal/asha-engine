package adjointbivectoraction

import "testing"

func TestGate250CliffordAdjointButNoQ8V(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if !a.Carrier.CommutatorActionTyped || !a.SimpleBlade.SkewSymmetric || a.SimpleBlade.Rank != 2 || a.SimpleBlade.KernelDimension != 6 {
		t.Fatalf("expected computable simple bivector adjoint matrix: %s", FormatSimpleBlade(a.SimpleBlade))
	}
	if a.KernelParity.Exact3DKernelPossible {
		t.Fatalf("real bivector route must not allow exact 3D kernel: %s", FormatKernelParity(a.KernelParity))
	}
	if a.EWBivectors.T3Grade2BladeDerived || a.EWBivectors.YPhiGrade2BladeDerived || a.Matrices.Q8VConstructed || a.Matrices.Neutral3PlaneDerived {
		t.Fatalf("EW Q8v route should remain blocked: %s | %s", FormatEWBivectors(a.EWBivectors), FormatMatrices(a.Matrices))
	}
	if a.ScalarPlane.VTauConstructed || a.Summary.TrialityUnblocked || a.Summary.YukawaTextureDerived {
		t.Fatalf("must not construct vtau/triality/yukawa: %s", FormatSummary(a.Summary))
	}
}
