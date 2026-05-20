package generation2boundaryweighteddeficitclosurestationarityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate663Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ScaleSweepInherited || !a.Inherited.Lambda12SelectedInGrid || !a.Inherited.Lambda12SelectedLocally || !a.Inherited.ExactWeightNearSevenOver72 || !a.Inherited.NoNativeScaleSelection || !a.Inherited.NoNativeSevenOver72 || !a.Inherited.NoFullUncertainty || !a.Inherited.NoNativeTransport || !a.Inherited.NoBoundaryStress {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Seed.Mu0GeV <= 0 || a.Seed.Lambda12GeV < 9e13 || len(a.Seed.InitialVector) != 13 {
		t.Fatalf("bad seed: %+v", a.Seed)
	}
	if a.Function.Lambda >= 0 || math.Abs(a.Function.AbsLambda-0.0497009420776833) > 5e-14 || math.Abs(a.Function.GaugeResidual-0.0509933868964996) > 5e-14 || math.Abs(a.Function.E72-8.525835107353608e-10) > 5e-16 {
		t.Fatalf("bad function: %+v", a.Function)
	}
	if a.Derivative.DE72DtAnalytic < 9e-4 || math.Abs(a.Derivative.DE72DtAnalytic-a.Derivative.DE72DtFiniteDifference) > 1e-10 || a.Derivative.Stationary || !a.Derivative.ZeroCrossingNotStationary {
		t.Fatalf("bad derivative: %+v", a.Derivative)
	}
	if a.BetaBalance.BalanceLeft > -9e-4 || a.BetaBalance.RequiredMinusActual < 0.009 || a.BetaBalance.StationarityWouldRequire {
		t.Fatalf("bad beta balance: %+v", a.BetaBalance)
	}
	if a.Curvature.SecondDerivative < 7e-5 || !strings.Contains(a.Curvature.LocalShape, "zero-crossing") || a.Curvature.ThresholdWidth1eMinus6 < 0.002 || a.Curvature.ThresholdWidth1eMinus4 < 0.2 {
		t.Fatalf("bad curvature: %+v", a.Curvature)
	}
	if math.Abs(a.ZeroScale.DeltaLogFromLambda12) > 1e-5 || math.Abs(a.ZeroScale.MuZeroOverLambda12-0.9999991071689) > 1e-9 || math.Abs(a.ZeroScale.E72AtZero) > 1e-12 || !a.ZeroScale.ClosureZeroAligned {
		t.Fatalf("bad zero scale: %+v", a.ZeroScale)
	}
	if len(a.WeightScale.Rows) != 3 || math.Abs(a.WeightScale.Rows[1].WBestMinus7Over72) > 1e-6 || !a.WeightScale.WeightIsSharpAtLambda12 || !a.WeightScale.CrossesSevenOver72NearLambda12 {
		t.Fatalf("bad weight-vs-scale: %+v", a.WeightScale)
	}
	if math.Abs(a.Orientation.OrientationE72AtLambda12-2.77672572136e-06) > 1e-16 || math.Abs(a.Orientation.OrientationZeroDeltaLog) < 0.002 || math.Abs(a.Orientation.OrientationWBestAtLambda12-0.09937065106106) > 5e-14 {
		t.Fatalf("bad orientation: %+v", a.Orientation)
	}
	if len(a.Source.Classification) != 4 || a.Discipline.ClaimsNativeScaleSelection || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsFullUncertaintyPropagation || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsNativeTransportTheorem || a.Discipline.ClaimsHiggsPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate663Boundary {
		t.Fatalf("firewall breach: %+v source=%+v", a.Discipline, a.Source)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BoundaryWeightedDeficitClosureStationarityAndBetaBalanceAuditTheorem().Verify()
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
