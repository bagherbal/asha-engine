package electroweak

import "testing"

func TestOperatorSearchBuild(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TensorDimension != 64 {
		t.Fatalf("tensor dimension = %d, want 64", a.TensorDimension)
	}
	if a.GradingSquareResidual > 1e-12 {
		t.Fatalf("grading square residual too large: %.3e", a.GradingSquareResidual)
	}
	if a.PositiveGradingDimension != a.NegativeGradingDimension {
		t.Fatalf("grading imbalance: +%d -%d", a.PositiveGradingDimension, a.NegativeGradingDimension)
	}
	if a.NeutralChiralityFlippingDimension != 0 {
		t.Fatalf("neutral B-L chirality-flipping dimension = %d, want 0 before scalar hypercharge bridge", a.NeutralChiralityFlippingDimension)
	}
	if a.HyperchargeDerived {
		t.Fatalf("hypercharge must remain underived in this gate")
	}
}
