package generation2k7kernelcokernelindexzeroaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate630Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate629DualCandidate || !a.Inherited.Gate629IsomorphismMissing || !a.Inherited.Gate629BoundaryAssignmentMissing || !a.Inherited.Gate629FirewallPreserved {
		t.Fatalf("bad Gate629 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.UDimension != 56 || a.Inherited.VDimension != 14 || a.Inherited.DirectSumDimension != 70 || a.Inherited.Lambda4Dimension != 70 || a.Inherited.IntersectionDimension != 7 || a.Inherited.SpanDimension != 63 || a.Inherited.CokernelDimension != 7 || a.Inherited.BoundaryPairDimension != 2 || a.Inherited.AugmentedChamberDimension != 72 {
		t.Fatalf("bad inherited dimensions: %+v", a.Inherited)
	}
	if !a.AdditionMap.SquareOperator || a.AdditionMap.DomainDimension != 70 || a.AdditionMap.CodomainDimension != 70 || !strings.Contains(a.AdditionMap.Formula, "u+v") {
		t.Fatalf("bad addition map: %+v", a.AdditionMap)
	}
	if !a.AdditionMap.KernelIsK7 || a.AdditionMap.KernelDimension != 7 || !strings.Contains(a.AdditionMap.KernelExpression, "K_7") {
		t.Fatalf("bad kernel audit: %+v", a.AdditionMap)
	}
	if !a.AdditionMap.ImageIsSpan || a.AdditionMap.ImageDimension != 63 || !a.AdditionMap.CokernelMatchesK7 || a.AdditionMap.CokernelDimension != 7 || a.AdditionMap.RankDefect != 7 || !a.AdditionMap.IndexZero {
		t.Fatalf("bad image/cokernel/index audit: %+v", a.AdditionMap)
	}
	if !a.Defect.DefectsBalanced || a.Defect.Index != 0 || !a.Defect.CandidateDefectPair || !a.Defect.FredholmAnalogyOnly || !strings.Contains(a.Defect.MissingPairing, "ker(A)") {
		t.Fatalf("bad defect audit: %+v", a.Defect)
	}
	if !a.BlockCompression.CompressionExact || !a.BlockCompression.DefectBlockCandidate || a.BlockCompression.PBBlocks != 8 || a.BlockCompression.PGBlocks != 2 || a.BlockCompression.SpanBlocks != 9 || a.BlockCompression.Lambda4Blocks != 10 || a.BlockCompression.BoundaryCoordinates != 2 || math.Abs(a.BlockCompression.BoundaryWeight-7.0/72.0) > 1e-15 {
		t.Fatalf("bad K7 block compression: %+v", a.BlockCompression)
	}
	if len(a.Pairing.Candidates) != 4 || a.Pairing.CanonicalPairingFound || a.Pairing.MetricPairingCertified || a.Pairing.HodgeStarPairingCertified || a.Pairing.EtaPairingCertified || a.Pairing.ProjectorPairingCertified || !strings.Contains(a.Pairing.MissingObject, "ker(A)") {
		t.Fatalf("bad pairing audit: %+v", a.Pairing)
	}
	if !a.BoundaryAssignment.DefectBlockCanSupplySeven || a.BoundaryAssignment.AssignmentCertified || a.BoundaryAssignment.NativeTransportTheorem || math.Abs(a.BoundaryAssignment.BoundaryWeight-7.0/72.0) > 1e-15 {
		t.Fatalf("bad boundary assignment: %+v", a.BoundaryAssignment)
	}
	if !a.NativeStatus.Lambda4Native || !a.NativeStatus.UImageRankNative || !a.NativeStatus.VImageRankNative || !a.NativeStatus.K7IntersectionNative || !a.NativeStatus.AdditionMapTyped || !a.NativeStatus.KernelDimensionTyped || !a.NativeStatus.CokernelDimensionTyped || !a.NativeStatus.IndexZeroTyped || a.NativeStatus.CanonicalKernelCokernelPairing || a.NativeStatus.BoundaryStressAssignmentNative || a.NativeStatus.K7DefectBoundaryTraceTheorem {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7KernelCokernelIndexZeroAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate629Inherited, StatusAdditionMapDefined, StatusKernelAIsK7, StatusCokernelADim7, StatusIndexZeroComputed, StatusK7BlockCompressionComputed, StatusK7DefectBlockCandidate, StatusNoCanonicalKerCokerPairing, StatusNoNativeBoundaryStressFromDefect, StatusGate630Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
