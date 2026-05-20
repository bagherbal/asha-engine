package generation2orientedwalldistancehyperplaneaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2OrientedWallDistanceHyperplaneAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 670 — Oriented Wall-Distance Hyperplane Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate670 oriented wall hyperplane audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate669 wall-coordinate audit", Passed: a.Inherited.WallCoordinatesInherited && a.Inherited.PositiveAndSignedFormsEquivalent && a.Inherited.HessianLayerSeparated && a.Inherited.MissingWallTheoremNamed && a.Inherited.NoNativeWallTheorem && a.Inherited.NoSevenOver72 && a.Inherited.NoBoundaryStress, Detail: FormatInherited(a.Inherited)},
			{Name: "write signed wall form", Passed: a.Signed.LambdaIsNegative && a.Signed.EquivalentBecauseLambdaNegative && math.Abs(a.Signed.PositiveResidual-a.Signed.SignedResidual) < 1e-18 && strings.Contains(a.Signed.Verdict, StatusSignedWallFormWritten), Detail: FormatSigned(a.Signed)},
			{Name: "classify wall-coordinate roles", Passed: len(a.Roles.Roles) == 4 && a.Roles.AllRolesClassified && strings.Contains(a.Roles.Verdict, StatusWallCoordinateRolesClassified), Detail: FormatRoles(a.Roles)},
			{Name: "audit hyperplane normal vector and 7/72 weight", Passed: len(a.Normal.Coefficients) == 4 && math.Abs(a.Normal.Coefficients[2]-sixtyFiveOver72) < 1e-15 && math.Abs(a.Normal.Coefficients[3]+sevenOver72) < 1e-15 && math.Abs(a.Normal.SumBoundaryWeights-1) < 1e-15 && a.Normal.TypedWeightUniqueInCurrentLedger && strings.Contains(a.Normal.Verdict, StatusSevenOver72TypedNormalWeight), Detail: FormatNormal(a.Normal)},
			{Name: "define HistoryWallBalanceSeal functional", Passed: a.Functional.Name == "HistoryWallBalanceSeal" && a.Functional.PassesBridgeTolerance && math.Abs(a.Functional.Value-8.52583441346e-10) < 1e-14 && strings.Contains(a.Functional.Verdict, StatusHistoryWallBalanceSealDefined), Detail: FormatFunctional(a.Functional)},
			{Name: "audit orientation approximation", Passed: math.Abs(a.Orientation.OrientationResidual-2.77672572133e-6) < 1e-12 && a.Orientation.ResidualGrowth > 0 && a.Orientation.RelativeToBoundarySplit > 0.002 && strings.Contains(a.Orientation.Verdict, StatusOrientationApproximationAudited), Detail: FormatOrientation(a.Orientation)},
			{Name: "preserve Hessian-layer firewall", Passed: a.Hessian.KeepsHessianSeparate && math.Abs(a.Hessian.QuarticWallCoordinate-absLambda12) < 1e-15 && math.Abs(a.Hessian.HessianCoordinate-2*absLambda12) < 1e-15, Detail: FormatHessian(a.Hessian)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeWallDistanceAirlock && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsNativeScalarZeroBoundary && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate670Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
