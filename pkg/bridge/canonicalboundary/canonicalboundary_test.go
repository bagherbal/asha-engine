package canonicalboundary

import (
	"math"
	"testing"
)

func TestCanonicalFiniteRGBoundarySeed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.ActionSelectedDimensionlessBoundarySeed {
		t.Fatalf("dimensionless boundary seed was not selected")
	}
	if !a.SU2KineticIsotropic {
		t.Fatalf("SU(2) kinetic block should be isotropic, got %v", a.SU2KineticEntries)
	}
	if math.Abs(a.SU2KineticCoefficient-1) > 1e-8 {
		t.Fatalf("SU(2) coefficient=%g, want 1", a.SU2KineticCoefficient)
	}
	if math.Abs(a.ScalarContactU1KineticCoefficient-3) > 1e-8 {
		t.Fatalf("K(Y_phi)=%g, want 3", a.ScalarContactU1KineticCoefficient)
	}
	if math.Abs(a.ScalarContactBoundarySin2-0.25) > 1e-8 {
		t.Fatalf("contact diagnostic sin²=%g, want 1/4", a.ScalarContactBoundarySin2)
	}
	if math.Abs(a.MatterHyperchargeKY-5.0/3.0) > 1e-8 {
		t.Fatalf("matter kY=%g, want 5/3", a.MatterHyperchargeKY)
	}
	if math.Abs(a.MatterBoundarySin2-3.0/8.0) > 1e-8 {
		t.Fatalf("matter boundary sin²=%g, want 3/8", a.MatterBoundarySin2)
	}
	if math.Abs(a.RequiredEmbeddingScaleSq-5.0/9.0) > 1e-8 {
		t.Fatalf("required lambda²=%g, want 5/9", a.RequiredEmbeddingScaleSq)
	}
	if a.EmbeddingMapSelected || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived {
		t.Fatalf("physical embedding/weak angle/alpha/masses must remain unclaimed")
	}
}
