package generation2orthogonalcokernelk7pairingaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate631Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate630IndexZero || !a.Inherited.Gate630PairingMissing || !a.Inherited.Gate630BoundaryMissing || !a.Inherited.Gate630FirewallPreserved {
		t.Fatalf("bad Gate630 inheritance: %+v", a.Inherited)
	}
	if a.Inherited.HDimension != 70 || a.Inherited.UDimension != 56 || a.Inherited.VDimension != 14 || a.Inherited.DirectSumDimension != 70 || a.Inherited.K7Dimension != 7 || a.Inherited.SpanDimension != 63 || a.Inherited.CokernelDimension != 7 || a.Inherited.Index != 0 || math.Abs(a.Inherited.BoundaryWeight-7.0/72.0) > 1e-15 {
		t.Fatalf("bad inherited dimensions: %+v", a.Inherited)
	}
	if !a.OrthogonalW7.RepresentsCokernel || a.OrthogonalW7.WDimension != 7 || !a.OrthogonalW7.WOrthogonalToU || !a.OrthogonalW7.WOrthogonalToV || !a.OrthogonalW7.DirectSumCertified || !strings.Contains(a.OrthogonalW7.WDefinition, "perp") {
		t.Fatalf("bad orthogonal W7 audit: %+v", a.OrthogonalW7)
	}
	if !a.ExactSequence.ExactAtK7 || !a.ExactSequence.ExactAtDirectSum || !a.ExactSequence.ExactAtH || !a.ExactSequence.ExactAtW7 || a.ExactSequence.DimensionAlternatingSum != 0 || !a.ExactSequence.ExactByRankNullity {
		t.Fatalf("bad exact sequence: %+v", a.ExactSequence)
	}
	if len(a.CandidatePairings.Candidates) != 5 || a.CandidatePairings.CanonicalPairingFound || a.CandidatePairings.NondegeneratePairingFound || !a.CandidatePairings.PairingProblemSharpened || !strings.Contains(a.CandidatePairings.MissingObject, "W_7") {
		t.Fatalf("bad candidate pairings: %+v", a.CandidatePairings)
	}
	if !a.HodgeStar.HodgeStarTypedOnLambda4 || !a.HodgeStar.MapsLambda4ToLambda4 || !a.HodgeStar.RequiresOrientationChoice || a.HodgeStar.RankTestImplemented || a.HodgeStar.NondegenerateCertified {
		t.Fatalf("bad Hodge-star audit: %+v", a.HodgeStar)
	}
	if len(a.ProjectorAlgebra.Rows) != 5 || !a.ProjectorAlgebra.K7FixedByPB || !a.ProjectorAlgebra.K7FixedByPG || !a.ProjectorAlgebra.PWKillsUPlusV || a.ProjectorAlgebra.AnyPairingCertified {
		t.Fatalf("bad projector audit: %+v", a.ProjectorAlgebra)
	}
	if a.Eta.TypedEtaOnLambda4Available || a.Eta.RankTestImplemented || a.Eta.PairingCertified || a.Eta.CompatibilityCertified {
		t.Fatalf("bad eta audit: %+v", a.Eta)
	}
	if !a.DeterminantLine.CanonicalLineRelation || a.DeterminantLine.PointwiseIsomorphism || !a.DeterminantLine.OrientationDependent || !a.DeterminantLine.CanSupportVolumeBookkeeping || a.DeterminantLine.CanSupportNormalizedTraceByItself {
		t.Fatalf("bad determinant line audit: %+v", a.DeterminantLine)
	}
	if a.BoundaryReadiness.K7ToW7PairingCertified || !a.BoundaryReadiness.DeterminantLineRelationAvailable || !a.BoundaryReadiness.StillRequiresW7ToBoundary || !a.BoundaryReadiness.StillRequiresDefectTraceToBoundary || a.BoundaryReadiness.BoundaryAssignmentCertified {
		t.Fatalf("bad boundary readiness audit: %+v", a.BoundaryReadiness)
	}
	if !a.NativeStatus.OrthogonalRepresentativeTyped || !a.NativeStatus.ExactDefectSequenceTyped || a.NativeStatus.CanonicalK7ToW7Pairing || a.NativeStatus.BoundaryStressAssignmentNative {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2OrthogonalCokernelK7PairingAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate630Inherited, StatusOrthogonalCokernelRepresentativeDefined, StatusExactDefectSequenceWritten, StatusCokernelRepresentedByW7, StatusK7W7PairingProblemSharpened, StatusNoCanonicalK7ToW7Pairing, StatusProjectorAlgebraFails, StatusHodgeStarRequiresExplicitRankTest, StatusNoBoundaryStressAssignment, StatusGate631Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
