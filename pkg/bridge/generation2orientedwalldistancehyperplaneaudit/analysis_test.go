package generation2orientedwalldistancehyperplaneaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate670Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.WallCoordinatesInherited || !a.Inherited.PositiveAndSignedFormsEquivalent || !a.Inherited.HessianLayerSeparated || !a.Inherited.NoNativeWallTheorem || !a.Inherited.NoSevenOver72 || !a.Inherited.NoBoundaryStress {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Signed.LambdaIsNegative || !a.Signed.EquivalentBecauseLambdaNegative || math.Abs(a.Signed.PositiveResidual-a.Signed.SignedResidual) > 1e-18 {
		t.Fatalf("bad signed form: %+v", a.Signed)
	}
	if len(a.Roles.Roles) != 4 || !a.Roles.AllRolesClassified {
		t.Fatalf("bad roles: %+v", a.Roles)
	}
	if len(a.Normal.Coefficients) != 4 || math.Abs(a.Normal.Coefficients[2]-sixtyFiveOver72) > 1e-15 || math.Abs(a.Normal.Coefficients[3]+sevenOver72) > 1e-15 || math.Abs(a.Normal.SumBoundaryWeights-1) > 1e-15 || !a.Normal.TypedWeightUniqueInCurrentLedger {
		t.Fatalf("bad normal: %+v", a.Normal)
	}
	if a.Functional.Name != "HistoryWallBalanceSeal" || !a.Functional.PassesBridgeTolerance || math.Abs(a.Functional.Value-8.52583441346e-10) > 1e-14 {
		t.Fatalf("bad functional: %+v", a.Functional)
	}
	if math.Abs(a.Orientation.OrientationResidual-2.77672572133e-6) > 1e-12 || a.Orientation.ResidualGrowth <= 0 {
		t.Fatalf("bad orientation approximation: %+v", a.Orientation)
	}
	if !a.Hessian.KeepsHessianSeparate || math.Abs(a.Hessian.HessianCoordinate-2*absLambda12) > 1e-15 {
		t.Fatalf("bad hessian firewall: %+v", a.Hessian)
	}
	if a.Discipline.ClaimsNativeWallDistanceAirlock || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsNativeScalarZeroBoundary || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate670Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestSignedWallFunctionalMatchesPositiveDistanceForm(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Signed.PositiveResidual-a.Functional.Value) > 1e-18 || math.Abs(a.Signed.SignedResidual-a.Functional.Value) > 1e-18 {
		t.Fatalf("wall forms mismatch: signed=%+v functional=%+v", a.Signed, a.Functional)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2OrientedWallDistanceHyperplaneAuditTheorem().Verify()
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
