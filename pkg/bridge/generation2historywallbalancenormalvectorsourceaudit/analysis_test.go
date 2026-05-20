package generation2historywallbalancenormalvectorsourceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate671Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.HistoryWallBalanceInherited || !a.Inherited.SignedWallFormWritten || !a.Inherited.FunctionalDefined || !a.Inherited.NoNativeWallAirlock || !a.Inherited.NoNativeSevenOver72 || !a.Inherited.NoBoundaryStress || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if len(a.Normal.Coefficients) != 4 || math.Abs(a.Normal.Coefficients[2]-sixtyFiveOver72) > 1e-15 || math.Abs(a.Normal.Coefficients[3]+sevenOver72) > 1e-15 || !a.Normal.HistorySideUnitWeights || !a.Normal.BoundaryWeightsSumToOne || !a.Normal.SignedAntiAlignment {
		t.Fatalf("bad normal vector: %+v", a.Normal)
	}
	if a.Decomposition.HistoryBlockLabel != "(1,1)" || a.Decomposition.BoundaryBlockLabel != "(65/72,-7/72)" {
		t.Fatalf("bad decomposition: %+v", a.Decomposition)
	}
	if len(a.Minimality.Alternatives) != 6 || !a.Minimality.N72BestAmongTypedExact || a.Minimality.BestExactName != "seven over seventy two boundary pull" || math.Abs(a.Minimality.BestExactAbsResidual-8.52583441346e-10) > 1e-14 {
		t.Fatalf("bad minimality: %+v", a.Minimality)
	}
	if !a.Coordinate.CoordinateSealed || !a.Coordinate.PreservesOnlyGate669WallNormalization {
		t.Fatalf("bad coordinate audit: %+v", a.Coordinate)
	}
	if math.Abs(a.Orientation.ExactResidualN72-8.52583441346e-10) > 1e-14 || math.Abs(a.Orientation.OrientationResidualN72-2.77672572133e-6) > 1e-12 || a.Orientation.ResidualGrowth <= 0 {
		t.Fatalf("bad orientation audit: %+v", a.Orientation)
	}
	if !a.ScaleLocal.Lambda12SelectedInGate662 || !a.ScaleLocal.LocalGate662MinimumAtLambda12 || !a.ScaleLocal.N72BestTypedNormalOnlyAtLambda12 || a.ScaleLocal.NearestLocalNonzeroResidual <= a.ScaleLocal.N72AtLambda12Residual {
		t.Fatalf("bad scale-local audit: %+v", a.ScaleLocal)
	}
	if len(a.Source.Candidates) != 4 || !a.Source.AugmentedTraceCandidate || !a.Source.BoundaryInterpolationCandidate || !a.Source.HistoryDeficitConservationCandidate || !a.Source.CoordinateArtifactRisk {
		t.Fatalf("bad source audit: %+v", a.Source)
	}
	if a.Discipline.ClaimsNativeNormalVectorTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsWallDistanceAirlockTheorem || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate671Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestAlternativeResidualOrdering(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	var n72, oneTenth AlternativeNormal
	for _, alt := range a.Minimality.Alternatives {
		switch alt.Name {
		case "seven over seventy two boundary pull":
			n72 = alt
		case "one tenth boundary pull":
			oneTenth = alt
		}
	}
	if !(n72.AbsExact < oneTenth.AbsExact) {
		t.Fatalf("expected 7/72 exact residual to beat 1/10: n72=%+v oneTenth=%+v", n72, oneTenth)
	}
	if !(oneTenth.AbsOrient < n72.AbsOrient) {
		t.Fatalf("expected orientation approximation to make nearby 1/10 competitive: n72=%+v oneTenth=%+v", n72, oneTenth)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HistoryWallBalanceNormalVectorSourceAndMinimalityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
