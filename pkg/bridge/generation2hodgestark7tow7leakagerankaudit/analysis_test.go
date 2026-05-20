package generation2hodgestark7tow7leakagerankaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate632Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.PairingProblemSharpened || !a.Inherited.HodgeRankTestRequired || !a.Inherited.Gate631FirewallPreserved {
		t.Fatalf("bad Gate631 inheritance: %+v", a.Inherited)
	}
	if !a.HodgeStar.TypedOnLambda4R8 || a.HodgeStar.MatrixDimension != 70 || a.HodgeStar.StarSquaredResidual > 1e-12 || a.HodgeStar.SelfDualDimension != 35 || a.HodgeStar.AntiSelfDualDimension != 35 {
		t.Fatalf("bad Hodge-star audit: %+v", a.HodgeStar)
	}
	if a.Basis.QKCols != 7 || a.Basis.QWCols != 7 || a.Basis.SpanCols != 63 || !a.Basis.QWOrthogonalToUAndV || !a.Basis.K7ContainedInUAndV {
		t.Fatalf("bad basis certificate: %+v", a.Basis)
	}
	if a.Basis.QKOrthonormalResidual > 1e-8 || a.Basis.QWOrthonormalResidual > 1e-8 || a.Basis.PBQKMinusQKResidual > 1e-8 || a.Basis.PGQKMinusQKResidual > 1e-8 || a.Basis.PBQWResidual > 1e-8 || a.Basis.PGQWResidual > 1e-8 {
		t.Fatalf("basis residuals too high: %+v", a.Basis)
	}
	if a.Leakage.Rank != 0 || a.Leakage.FrobeniusNorm > 1e-10 || math.Abs(a.Leakage.Determinant) > 1e-14 || len(a.Leakage.SingularValues) != 7 {
		t.Fatalf("Hodge leakage should fail at rank zero: %+v", a.Leakage)
	}
	if !a.ImageContainment.StarK7ContainedInUPlusV || a.ImageContainment.TransverseComponentDetected || a.ImageContainment.PWStarK7FrobeniusNorm > 1e-10 || a.ImageContainment.PUVStarK7FrobeniusNorm < 2.6 {
		t.Fatalf("bad containment audit: %+v", a.ImageContainment)
	}
	if !a.PairingMetric.Degenerate || a.PairingMetric.RankFull || a.PairingMetric.Trace > 1e-20 {
		t.Fatalf("bad pairing metric audit: %+v", a.PairingMetric)
	}
	if a.Orientation.NonzeroDeterminant || a.Orientation.Sign != 0 || a.Orientation.PhysicalOrientationCertified {
		t.Fatalf("bad orientation audit: %+v", a.Orientation)
	}
	if a.AlternativeComposites.AnyHigherRankThanPhiStar || a.AlternativeComposites.AnyNondegenerate || len(a.AlternativeComposites.Rows) != 5 {
		t.Fatalf("bad alternative composite audit: %+v", a.AlternativeComposites)
	}
	for _, row := range a.AlternativeComposites.Rows {
		if row.Rank != 0 || row.FrobeniusNorm > 1e-8 {
			t.Fatalf("alternative composite should not produce leakage: %+v", row)
		}
	}
	if a.BoundaryReadiness.HodgePairingCertified || a.BoundaryReadiness.K7ToW7PairingFound || a.BoundaryReadiness.BoundaryAssignmentCertified {
		t.Fatalf("bad boundary readiness audit: %+v", a.BoundaryReadiness)
	}
	if a.NativeStatus.HodgePairingNondegenerate || a.NativeStatus.CanonicalK7ToW7PairingFound || a.NativeStatus.BoundaryStressAssignmentNative {
		t.Fatalf("bad native status: %+v", a.NativeStatus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HodgeStarK7ToW7LeakageRankAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate631Inherited, StatusHodgeStarTyped, StatusK7AndW7BasesCertified, StatusHodgeLeakageMatrixComputed, StatusHodgeStarDoesNotPairK7ToW7, StatusNoCanonicalK7W7PairingFound, StatusNoBoundaryStressAssignment, StatusGate632Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
