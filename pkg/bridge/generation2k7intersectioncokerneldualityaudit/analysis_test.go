package generation2k7intersectioncokerneldualityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate629Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate628ChamberCandidate || !a.Inherited.Gate628ProjectionMissing || !a.Inherited.Gate628ProductAirlockMissing || !a.Inherited.Gate628FirewallPreserved {
		t.Fatalf("bad Gate628 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.Lambda4Dimension != 70 || a.Inherited.BoundaryPairDimension != 2 || a.Inherited.AugmentedChamberDimension != 72 || a.Inherited.K7Dimension != 7 || a.Inherited.NonK7Lambda4Complement != 63 || a.Inherited.AugmentedComplementDimension != 65 {
		t.Fatalf("bad inherited dimensions: %+v", a.Inherited)
	}
	if a.Span.RankPB != 56 || a.Span.RankPG != 14 || a.Span.IntersectionDimension != 7 || a.Span.SpanDimension != 63 || a.Span.Lambda4Dimension != 70 || a.Span.CokernelDimension != 7 || !a.Span.SpanMatchesExpected || !a.Span.CokernelMatchesK7Dimension || !a.Span.SpanDimensionCertifiedByRank {
		t.Fatalf("bad span audit: %+v", a.Span)
	}
	if !a.Duality.DimensionsEqual || !a.Duality.EqualityIsOnlyDimensional || !a.Duality.DualityCandidate || a.Duality.CanonicalIsomorphismFound || a.Duality.CanonicalPairingFound || !strings.Contains(a.Duality.MissingMap, "Phi") {
		t.Fatalf("bad duality audit: %+v", a.Duality)
	}
	if !a.ChamberSplit.SplitMatches72 || a.ChamberSplit.IntersectionOrGapDimension != 7 || a.ChamberSplit.SpanDimension != 63 || a.ChamberSplit.BoundaryPairDimension != 2 || a.ChamberSplit.AugmentedChamberDimension != 72 || !a.ChamberSplit.SharperThan70Plus2 {
		t.Fatalf("bad chamber split: %+v", a.ChamberSplit)
	}
	if a.ComplementRole.SpanBoundaryComplement != 65 || !a.ComplementRole.Equals65Over72 || math.Abs(a.ComplementRole.ComplementWeight-65.0/72.0) > 1e-15 || a.ComplementRole.NativeRoleTheoremFound {
		t.Fatalf("bad complement role: %+v", a.ComplementRole)
	}
	if len(a.BoundaryPullAssignment.Candidates) != 3 || !a.BoundaryPullAssignment.Candidates[0].CanSupplySeven || !a.BoundaryPullAssignment.Candidates[1].CanSupplySeven || !a.BoundaryPullAssignment.Candidates[2].CanSupplySeven || a.BoundaryPullAssignment.AssignmentCertified {
		t.Fatalf("bad boundary-pull candidates: %+v", a.BoundaryPullAssignment)
	}
	if !a.WeightedMixture.ScalarWeightAs63Plus2 || !a.WeightedMixture.BoundaryWeightAsSeven || math.Abs(a.WeightedMixture.BoundaryWeight-7.0/72.0) > 1e-15 || math.Abs(a.WeightedMixture.ScalarWeight-65.0/72.0) > 1e-15 {
		t.Fatalf("bad mixture reinterpretation: %+v", a.WeightedMixture)
	}
	if !a.NativeStatus.Lambda4Native || !a.NativeStatus.PBImageRankNative || !a.NativeStatus.PGImageRankNative || !a.NativeStatus.K7IntersectionNative || !a.NativeStatus.BooleanOctonionicSpanDimensionTyped || !a.NativeStatus.Lambda4CokernelDimensionTyped || a.NativeStatus.IntersectionCokernelIsomorphism || a.NativeStatus.BoundaryPullAssignmentNative || a.NativeStatus.DualBoundaryProjectorNative {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7IntersectionCokernelDualityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate628Inherited, StatusSpanDimensionComputed, StatusCokernelDimensionComputed, Status72SplitAudited, Status63SpanCandidate, StatusIntersectionCokernelCandidate, Status65SpanBoundaryComplement, StatusNoNativeIntersectionCokernelIso, StatusNoBoundaryPullAssignment, StatusNoNativeDualBoundaryProjector, StatusGate629Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
