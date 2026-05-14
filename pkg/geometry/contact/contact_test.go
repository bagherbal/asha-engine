package contact

import (
	"math"
	"testing"
)

func TestBuildDefaultContactSpace(t *testing.T) {
	space, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if got, want := space.Dimension(), space.ExpectedContactDenominator(); got != want {
		t.Fatalf("contact dimension = %d, want %d", got, want)
	}
	if got := space.ContactIndex(); math.Abs(got-1) > 1e-8 {
		t.Fatalf("contact index = %.16g, want 1", got)
	}

	frameResidual, err := space.FrameIsometryResidual()
	if err != nil {
		t.Fatal(err)
	}
	if frameResidual > 1e-8 {
		t.Fatalf("contact frame residual = %.3e", frameResidual)
	}

	projectorResidual, err := space.ContactProjector.IdempotenceResidual()
	if err != nil {
		t.Fatal(err)
	}
	if projectorResidual > 1e-8 {
		t.Fatalf("contact projector residual = %.3e", projectorResidual)
	}

	booleanResidual, err := space.BooleanContainmentResidual()
	if err != nil {
		t.Fatal(err)
	}
	if booleanResidual > 1e-8 {
		t.Fatalf("Boolean containment residual = %.3e", booleanResidual)
	}

	g2Residual, err := space.G2ContainmentResidual()
	if err != nil {
		t.Fatal(err)
	}
	if g2Residual > 1e-8 {
		t.Fatalf("G₂ containment residual = %.3e", g2Residual)
	}

	if leakage := space.BareLeakageNorm(); leakage <= 0 || math.IsNaN(leakage) {
		t.Fatalf("bare leakage should be positive and finite, got %.16g", leakage)
	}
}
