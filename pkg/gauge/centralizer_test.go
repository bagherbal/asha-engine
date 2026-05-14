package gauge

import "testing"

func TestContactCentralizer(t *testing.T) {
	c, err := BuildContactCentralizer()
	if err != nil {
		t.Fatalf("BuildContactCentralizer failed: %v", err)
	}
	if c.G2Rank != 14 {
		t.Fatalf("dim(g2)=%d, want 14", c.G2Rank)
	}
	if c.CentralizerDimension != 4 {
		t.Fatalf("dim(g2^R)=%d, want 4", c.CentralizerDimension)
	}
	if c.CenterDimension != 1 {
		t.Fatalf("center dimension=%d, want 1", c.CenterDimension)
	}
	if c.DerivedDimension != 3 {
		t.Fatalf("derived dimension=%d, want 3", c.DerivedDimension)
	}
	if residual := c.CentralizerResidual(); residual > 1e-8 {
		t.Fatalf("centralizer residual %.3e > tolerance", residual)
	}
	if residual := c.ClosureResidual(); residual > 1e-8 {
		t.Fatalf("closure residual %.3e > tolerance", residual)
	}
}
