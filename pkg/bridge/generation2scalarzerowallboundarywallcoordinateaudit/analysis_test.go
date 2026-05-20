package generation2scalarzerowallboundarywallcoordinateaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate669Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ScalarCoordinateInherited || !a.Inherited.HessianLayerSeparated || !a.Inherited.NoScalarAirlock || !a.Inherited.NoSevenOver72 || !a.Inherited.NoBoundaryStress {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Scalar.IsBelowWall || !a.Scalar.AbsoluteValueTyped || math.Abs(a.Scalar.DistanceBelowWall-absLambda12) > 1e-15 {
		t.Fatalf("bad scalar wall: %+v", a.Scalar)
	}
	if !a.Gauge.IsAboveWall || math.Abs(a.Gauge.GaugeResidual-r3Minus1) > 1e-15 {
		t.Fatalf("bad gauge wall: %+v", a.Gauge)
	}
	if !a.Boundary.EquivalentFormsAgree || math.Abs(a.Boundary.ClosureResidualPositiveForm-a.Boundary.ClosureResidualSignedForm) > 1e-18 || math.Abs(a.Boundary.XiBoundary-xiBoundary) > 1e-15 {
		t.Fatalf("bad boundary rewrite: %+v", a.Boundary)
	}
	if len(a.Flavor.Rows) != 3 || !a.Flavor.FlavorWallSupported || !a.Flavor.ScalarWallSupported || !a.Flavor.GaugeWallSupported {
		t.Fatalf("bad flavor analogy: %+v", a.Flavor)
	}
	if !a.Hessian.LayersSeparated || math.Abs(a.Hessian.HessianCoordinate-2*absLambda12) > 1e-15 {
		t.Fatalf("bad hessian separation: %+v", a.Hessian)
	}
	if a.Target.PrimaryName != "BoundaryWallCoordinateAirlockTheorem" || len(a.Target.RequiredObjects) != 4 || !strings.Contains(a.Target.Verdict, StatusNoNativeWallDistanceAirlockTheorem) {
		t.Fatalf("bad target: %+v", a.Target)
	}
	if a.Discipline.ClaimsNativeWallDistanceAirlock || a.Discipline.ClaimsNativeScalarZeroBoundary || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsHiggsMassPrediction || a.Discipline.ClaimsScalarStability || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNSDerivation || a.Discipline.Verdict != StatusGate669Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarZeroWallDistanceAndBoundaryWallCoordinateAuditTheorem().Verify()
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

func TestSignedFormMatchesPositiveDistanceForm(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Boundary.ClosureResidualPositiveForm-a.Boundary.ClosureResidualSignedForm) > 1e-18 {
		t.Fatalf("forms disagree: %+v", a.Boundary)
	}
}
