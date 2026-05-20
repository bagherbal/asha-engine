package generation2k7boundaryprojectionweightaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate627Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate626Inherited || !a.Inherited.Gate626ClosureIsBridgeOnly || a.Inherited.Gate626NativeWeightSource {
		t.Fatalf("bad Gate626 inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Weight.Value-7.0/72.0) > 1e-15 || a.Weight.Numerator != 7 || a.Weight.Denominator != 72 || a.Weight.ComplementNumerator != 65 {
		t.Fatalf("bad weight identification: %+v", a.Weight)
	}
	if !a.Numerator.MatchesDimK7 || !a.Numerator.K7NativeCarrierCertified || a.Numerator.ProjectionToBoundaryFound {
		t.Fatalf("bad K7 numerator audit: %+v", a.Numerator)
	}
	if a.Numerator.K7Dimension != 7 || a.Numerator.RankPB != 56 || a.Numerator.RankPG != 14 {
		t.Fatalf("bad K7 geometry: %+v", a.Numerator)
	}
	if !a.Denominator.AnyExistingTypedCandidate || a.Denominator.CertifiedBoundaryCarrier || len(a.Denominator.Rows) < 4 {
		t.Fatalf("bad denominator audit: %+v", a.Denominator)
	}
	if a.Denominator.Rows[0].Value != 72 || a.Denominator.Rows[0].CertifiedAsDenom {
		t.Fatalf("bad 8x9 candidate row: %+v", a.Denominator.Rows[0])
	}
	if !a.Complement.ComplementEquals65Over72 || !a.Complement.ArithmeticComplementOnly || a.Complement.NativeComplementCarrier {
		t.Fatalf("bad complement audit: %+v", a.Complement)
	}
	if math.Abs(a.Midpoint.MidpointPullWeight-7.0/36.0) > 1e-15 || math.Abs(a.Midpoint.RewriteResidual) > 1e-15 || a.Midpoint.NativeMidpointProjection {
		t.Fatalf("bad midpoint rewrite: %+v", a.Midpoint)
	}
	if !a.Projection.WeightEqualsDimRatio || a.Projection.ProjectionOperatorExists || a.Projection.TraceCertified || a.Projection.CandidateChamberDimension != 72 {
		t.Fatalf("bad projection audit: %+v", a.Projection)
	}
	if !a.NativeStatus.NumeratorK7Native || a.NativeStatus.Denominator72BoundaryCarrierNative || a.NativeStatus.SevenOverSeventyTwoSourceTheorem {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2K7BoundaryProjectionWeightAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed theorem: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate626Inherited, StatusWeightIdentified, StatusNumeratorK7Candidate, StatusDenominator72Candidate, StatusComplementAudited, StatusMidpointRewriteAudited, StatusProjectionMissing, StatusNoCertified72Carrier, StatusNoNativeWeightTheorem, StatusGate627Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
