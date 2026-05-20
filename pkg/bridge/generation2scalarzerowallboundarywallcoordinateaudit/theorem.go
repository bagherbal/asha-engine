package generation2scalarzerowallboundarywallcoordinateaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2ScalarZeroWallDistanceAndBoundaryWallCoordinateAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 669 — Scalar Zero-Wall Distance and Boundary Wall-Coordinate Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate669 wall coordinate audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate668 scalar coordinate audit", Passed: a.Inherited.ScalarCoordinateInherited && a.Inherited.HessianLayerSeparated && a.Inherited.NoScalarAirlock && a.Inherited.NoSevenOver72 && a.Inherited.NoBoundaryStress && a.Inherited.NoTransport, Detail: FormatInherited(a.Inherited)},
			{Name: "define scalar zero-wall distance", Passed: a.Scalar.IsBelowWall && a.Scalar.AbsoluteValueTyped && math.Abs(a.Scalar.DistanceBelowWall-absLambda12) < 1e-15 && strings.Contains(a.Scalar.Verdict, StatusAbsLambdaScalarZeroWallDistance), Detail: FormatScalar(a.Scalar)},
			{Name: "define gauge meeting-wall distance", Passed: a.Gauge.IsAboveWall && math.Abs(a.Gauge.GaugeResidual-r3Minus1) < 1e-15 && strings.Contains(a.Gauge.Verdict, StatusR3GaugeMeetingWallDistance), Detail: FormatGauge(a.Gauge)},
			{Name: "rewrite signed boundary stress form", Passed: a.Boundary.EquivalentFormsAgree && math.Abs(a.Boundary.XiBoundary-xiBoundary) < 1e-15 && math.Abs(a.Boundary.ClosureResidualPositiveForm-a.Boundary.ClosureResidualSignedForm) < 1e-18 && math.Abs(a.Boundary.ClosureResidualPositiveForm-8.52583441346e-10) < 1e-14, Detail: FormatBoundary(a.Boundary)},
			{Name: "audit flavor wall analogy", Passed: len(a.Flavor.Rows) == 3 && a.Flavor.FlavorWallSupported && a.Flavor.ScalarWallSupported && a.Flavor.GaugeWallSupported && strings.Contains(a.Flavor.Verdict, StatusEpsilonEFlavorWallDistance), Detail: FormatFlavor(a.Flavor)},
			{Name: "preserve Hessian layer separation", Passed: a.Hessian.LayersSeparated && math.Abs(a.Hessian.QuarticWallCoordinate-absLambda12) < 1e-15 && math.Abs(a.Hessian.HessianCoordinate-2*absLambda12) < 1e-15, Detail: FormatHessian(a.Hessian)},
			{Name: "name missing wall-distance theorem target", Passed: a.Target.PrimaryName == "BoundaryWallCoordinateAirlockTheorem" && len(a.Target.RequiredObjects) == 4 && strings.Contains(a.Target.Verdict, StatusNoNativeWallDistanceAirlockTheorem), Detail: FormatTarget(a.Target)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeWallDistanceAirlock && !a.Discipline.ClaimsNativeScalarZeroBoundary && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate669Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
