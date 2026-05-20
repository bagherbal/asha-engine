package generation2k7overlambda4boundarypairprojectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate628Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate627Inherited || !a.Inherited.Gate627NumeratorIsNative || a.Inherited.Gate627Certified72Carrier || a.Inherited.Gate627ProjectionExists || a.Inherited.Gate627SourceTheorem {
		t.Fatalf("bad Gate627 inheritance: %+v", a.Inherited)
	}
	if a.Chamber.Lambda4Dimension != 70 || a.Chamber.BoundaryPairDimension != 2 || a.Chamber.AugmentedChamberDimension != 72 || !a.Chamber.EqualsTargetDenominator || a.Chamber.DirectSumCertifiedNative {
		t.Fatalf("bad chamber audit: %+v", a.Chamber)
	}
	if !a.DenominatorComparison.BestIs70Plus2 || a.DenominatorComparison.AnyNativeDenominator || !a.DenominatorComparison.AnyBridgeCandidate || a.DenominatorComparison.Rows[0].Expression != "70 + 2" || a.DenominatorComparison.Rows[0].UsesQuarantinedLedger {
		t.Fatalf("bad denominator comparison: %+v", a.DenominatorComparison)
	}
	if a.BoundaryPair.PairDimension != 2 || !a.BoundaryPair.PairIsGate613Boundary || !a.BoundaryPair.PairInheritedFromGate626 || a.BoundaryPair.PairNativeFiniteObject || !a.BoundaryPair.BridgeCoordinateOnly {
		t.Fatalf("bad boundary pair: %+v", a.BoundaryPair)
	}
	if a.K7Embedding.K7Dimension != 7 || a.K7Embedding.Lambda4Dimension != 70 || !a.K7Embedding.K7FitsInsideLambda4 || !a.K7Embedding.NativeCarrierCertified || a.K7Embedding.ProjectionToBoundaryFound {
		t.Fatalf("bad K7 embedding: %+v", a.K7Embedding)
	}
	if a.Complement.NonK7Lambda4ComplementDimension != 63 || a.Complement.AugmentedComplementDimension != 65 || !a.Complement.Equals65Over72 || !a.Complement.HasStructuredComplementReading || a.Complement.NativeComplementProjection {
		t.Fatalf("bad complement: %+v", a.Complement)
	}
	if math.Abs(a.ProjectionTrace.TraceFraction-7.0/72.0) > 1e-15 || !a.ProjectionTrace.TraceFractionMatches || a.ProjectionTrace.ProjectionOperatorExists || a.ProjectionTrace.TraceFunctionalCertified || a.ProjectionTrace.IntertwinerCertified {
		t.Fatalf("bad projection trace: %+v", a.ProjectionTrace)
	}
	if !a.WeightedClosure.ChamberRatioMatchesGate626 || math.Abs(a.WeightedClosure.WeightedClosureResidual-a.Inherited.Gate626WeightedResidual) > 1e-15 {
		t.Fatalf("bad weighted closure carry: %+v", a.WeightedClosure)
	}
	if !a.NativeStatus.Lambda4Native || !a.NativeStatus.K7Native || a.NativeStatus.BoundaryPairNativeFinite || a.NativeStatus.AugmentedChamberNative || a.NativeStatus.K7BoundaryPullProjectorNative || a.NativeStatus.TraceFractionTheoremNative {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7OverLambda4BoundaryPairProjectionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate627Inherited, Status70Plus2Identified, Status72Lambda4BoundaryPair, StatusBoundaryPairInherited, StatusK7InsideLambda4Audited, Status65ComplementCandidate, StatusProjectionTraceCandidate, StatusNoProductAirlock, StatusNoK7BoundaryPullProjector, StatusNoNativeAugmentedChamber, StatusGate628Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
