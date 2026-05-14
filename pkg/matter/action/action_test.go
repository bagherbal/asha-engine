package action

import (
	"math"
	"testing"
)

func TestRepresentationActionBridge(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Operator.Rows() != 16 || a.Operator.Cols() != 16 {
		t.Fatalf("operator shape = %dx%d, expected 16x16", a.Operator.Rows(), a.Operator.Cols())
	}
	if math.Abs(a.VacuumResponse) > 1e-10 {
		t.Fatalf("vacuum response = %.12g, expected zero", a.VacuumResponse)
	}
	if !a.PairDegenerate {
		t.Fatalf("expected pair-degenerate one-particle response")
	}
	if a.TraceResidual > 1e-10 {
		t.Fatalf("trace residual = %.12g", a.TraceResidual)
	}
	if a.Rank != 15 {
		t.Fatalf("rank = %d, expected 15", a.Rank)
	}
	if a.ParityTraceResidual > 1e-10 {
		t.Fatalf("parity trace residual = %.12g", a.ParityTraceResidual)
	}
}
