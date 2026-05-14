package chiraltrace

import "testing"

func TestFiniteChiralBilinearMetric(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.BilinearMetricConstructed || !a.ScalarLRProjectorConstructed {
		t.Fatalf("expected finite scalar LR bilinear metric and projector to be constructed")
	}
	if a.ProjectorRank != a.RightDimension {
		t.Fatalf("rank=%d right=%d", a.ProjectorRank, a.RightDimension)
	}
	if a.DomainComplementDimension != 24 {
		t.Fatalf("expected 24 unused domain directions, got %d", a.DomainComplementDimension)
	}
	if a.FullCliffordTraceRulesDerived {
		t.Fatalf("full Clifford trace rules must remain open in Gate 58")
	}
	if a.CurrentScalarProjectionCoefficientsKnown || a.AttractiveSignDerived {
		t.Fatalf("Gate 58 must not claim Fierz coefficients or attraction")
	}
}
