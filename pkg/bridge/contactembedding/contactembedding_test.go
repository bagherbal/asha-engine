package contactembedding

import (
	"math"
	"testing"
)

func TestContactMatterHyperchargeEmbedding(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.ContactToMatterEmbeddingSelected {
		t.Fatalf("contact-to-matter hypercharge embedding was not selected")
	}
	if math.Abs(a.ContactU1KineticCoefficient-3) > 1e-8 {
		t.Fatalf("contact K=%g, want 3", a.ContactU1KineticCoefficient)
	}
	if math.Abs(a.MatterHyperchargeKY-5.0/3.0) > 1e-8 {
		t.Fatalf("matter kY=%g, want 5/3", a.MatterHyperchargeKY)
	}
	if math.Abs(a.EmbeddingScaleSq-5.0/9.0) > 1e-8 {
		t.Fatalf("lambda²=%g, want 5/9", a.EmbeddingScaleSq)
	}
	if math.Abs(a.EmbeddingScale-math.Sqrt(5.0)/3.0) > 1e-8 {
		t.Fatalf("lambda=%g, want sqrt(5)/3", a.EmbeddingScale)
	}
	if !a.OrientationSelected {
		t.Fatalf("orientation-preserving branch should be selected")
	}
	if !a.EmbeddedMatterHessianSelected || math.Abs(a.EmbeddedMatterHessian.At(3, 3)-5.0/3.0) > 1e-8 {
		t.Fatalf("embedded matter Hessian not selected: %s", FormatMatrix(a.EmbeddedMatterHessian))
	}
	if math.Abs(a.EmbeddedMatterBoundarySin2-3.0/8.0) > 1e-8 {
		t.Fatalf("embedded boundary sin²=%g, want 3/8", a.EmbeddedMatterBoundarySin2)
	}
	if a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalCouplingsDerived || a.PhysicalMassesDerived {
		t.Fatalf("physical couplings/masses must remain sealed")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("no observed input should be used")
	}
}
