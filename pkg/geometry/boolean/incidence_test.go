package boolean

import (
	"math"
	"testing"
)

func TestBuildIncidenceSupportR8(t *testing.T) {
	support, err := BuildIncidenceSupport(8, 3, 4)
	if err != nil {
		t.Fatalf("BuildIncidenceSupport failed: %v", err)
	}
	if support.LowerDimension() != 56 {
		t.Fatalf("lower dimension = %d, want 56", support.LowerDimension())
	}
	if support.UpperDimension() != 70 {
		t.Fatalf("upper dimension = %d, want 70", support.UpperDimension())
	}
	if support.RankFromGram(1e-8) != 56 {
		t.Fatalf("rank = %d, want 56", support.RankFromGram(1e-8))
	}
	residual, err := support.IsometryResidual()
	if err != nil {
		t.Fatalf("IsometryResidual failed: %v", err)
	}
	if residual > 1e-8 {
		t.Fatalf("isometry residual too large: %.3e", residual)
	}
	trace, err := support.Support.Trace()
	if err != nil {
		t.Fatalf("trace failed: %v", err)
	}
	if math.Abs(trace-56) > 1e-8 {
		t.Fatalf("Tr(P_B)=%.12f, want 56", trace)
	}
	idem, err := support.Support.IdempotenceResidual()
	if err != nil {
		t.Fatalf("idempotence failed: %v", err)
	}
	if idem > 1e-8 {
		t.Fatalf("projector residual too large: %.3e", idem)
	}
}
