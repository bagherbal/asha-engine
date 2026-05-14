package boundary

import "testing"

func TestBoundaryFixedClosureBuilds(t *testing.T) {
	c, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if c.SeedDimension == 0 {
		t.Fatalf("expected nonzero seed dimension")
	}
	if c.ClosureDimension < c.SeedDimension {
		t.Fatalf("closure dimension %d smaller than seed dimension %d", c.ClosureDimension, c.SeedDimension)
	}
	if c.MaxBoundaryLeakage > 1e-8 {
		t.Fatalf("boundary-fixed closure leaked across contact boundary: %.3e", c.MaxBoundaryLeakage)
	}
}
