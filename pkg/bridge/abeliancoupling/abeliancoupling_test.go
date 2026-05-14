package abeliancoupling

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if len(a.Fields) != 3 {
		t.Fatalf("expected 3 abelian diagnostics, got %d", len(a.Fields))
	}
	if !a.DiagonalTraceGramSelectedAsRepresentationMetric {
		t.Fatalf("expected representation metric diagnostic")
	}
	if a.DiagonalTraceGramSelectedAsGaugeKineticHessian {
		t.Fatalf("diagonal trace-Gram must not be promoted to physical Hessian by this gate")
	}
	if a.PhysicalGaugeCouplingsDerived || a.FineStructureDerived {
		t.Fatalf("physical couplings/alpha must remain open")
	}
	if a.Hypercharge.ChargeTableKY < 1.66 || a.Hypercharge.ChargeTableKY > 1.67 {
		t.Fatalf("expected kY near 5/3, got %.12f", a.Hypercharge.ChargeTableKY)
	}
}
