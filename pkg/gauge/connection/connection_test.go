package connection

import "testing"

func TestProjectedConnectionIdentity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if a.MaxBlockReconstructionResidual > 1e-8 {
		t.Fatalf("block reconstruction residual too large: %.3e", a.MaxBlockReconstructionResidual)
	}
	if a.MaxOffDiagonalNorm <= 1e-8 {
		t.Fatalf("expected nonzero off-diagonal sector")
	}
	if a.MaxProjectionDefectNorm <= 1e-8 {
		t.Fatalf("expected nonzero projection defect")
	}
	if a.MaxSecondFundamentalNorm <= 1e-8 {
		t.Fatalf("expected nonzero second fundamental curvature")
	}
	if a.MaxCurvatureIdentityResidual > 1e-8 {
		t.Fatalf("curvature identity residual too large: %.3e", a.MaxCurvatureIdentityResidual)
	}
	if a.MaxCurvatureIdentityRelative > 1e-8 {
		t.Fatalf("curvature identity relative residual too large: %.3e", a.MaxCurvatureIdentityRelative)
	}
}
