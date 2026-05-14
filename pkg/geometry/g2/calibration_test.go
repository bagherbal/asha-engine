package g2

import (
	"math"
	"testing"
)

func TestBuildCalibrationSupport(t *testing.T) {
	support, err := BuildCalibrationSupport()
	if err != nil {
		t.Fatalf("BuildCalibrationSupport returned error: %v", err)
	}
	if support.MiddleDimension() != 70 {
		t.Fatalf("middle dimension = %d, want 70", support.MiddleDimension())
	}
	if support.SectorDimension() != 14 {
		t.Fatalf("sector dimension = %d, want 14", support.SectorDimension())
	}
	if rank := support.RankFromGram(1e-8); rank != 14 {
		t.Fatalf("rank = %d, want 14", rank)
	}
	iso, err := support.IsometryResidual()
	if err != nil {
		t.Fatal(err)
	}
	if iso > 1e-8 {
		t.Fatalf("isometry residual too high: %.3e", iso)
	}
	idem, err := support.Support.IdempotenceResidual()
	if err != nil {
		t.Fatal(err)
	}
	if idem > 1e-8 {
		t.Fatalf("projector idempotence residual too high: %.3e", idem)
	}
	trace, err := support.Support.Trace()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(trace-14) > 1e-8 {
		t.Fatalf("trace = %.12f, want 14", trace)
	}
}
