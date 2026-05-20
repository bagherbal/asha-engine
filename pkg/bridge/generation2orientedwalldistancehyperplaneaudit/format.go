package generation2orientedwalldistancehyperplaneaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate669Inheritance) string {
	return fmt.Sprintf("wallInherited=%t formsEquivalent=%t scalarWall=%t gaugeWall=%t flavorWall=%t hessianSeparated=%t theoremNamed=%t noWall=%t no7=%t noBoundary=%t verdict=%q", x.WallCoordinatesInherited, x.PositiveAndSignedFormsEquivalent, x.ScalarWallDistance, x.GaugeWallDistance, x.FlavorWallDistance, x.HessianLayerSeparated, x.MissingWallTheoremNamed, x.NoNativeWallTheorem, x.NoSevenOver72, x.NoBoundaryStress, x.Verdict)
}

func FormatSigned(x SignedWallFormAudit) string {
	return fmt.Sprintf("positive=%q signed=%q posResidual=%.15g signedResidual=%.15g equivalent=%t lambdaNegative=%t verdict=%q", x.PositiveDistanceForm, x.SignedWallForm, x.PositiveResidual, x.SignedResidual, x.EquivalentBecauseLambdaNegative, x.LambdaIsNegative, x.Verdict)
}

func FormatRole(r WallCoordinateRole) string {
	return fmt.Sprintf("%s value=%.15g wall=%q sign=%q role=%q layer=%q", r.Coordinate, r.Value, r.Wall, r.Sign, r.Role, r.Layer)
}

func FormatRoles(x WallCoordinateRolesAudit) string {
	rows := make([]string, 0, len(x.Roles))
	for _, r := range x.Roles {
		rows = append(rows, FormatRole(r))
	}
	return fmt.Sprintf("roles=%d allClassified=%t verdict=%q ledger=[%s]", len(x.Roles), x.AllRolesClassified, x.Verdict, strings.Join(rows, "; "))
}

func FormatNormal(x NormalVectorAudit) string {
	coeffs := make([]string, 0, len(x.Coefficients))
	for i, c := range x.Coefficients {
		coeffs = append(coeffs, fmt.Sprintf("%s:%.15g", x.Coordinates[i], c))
	}
	return fmt.Sprintf("normal=%q coeffs=[%s] seven=%.15g sixtyFive=%.15g sum=%.15g wBest=%.15g delta=%.15g uniqueTyped=%t verdict=%q", x.NormalVectorLabel, strings.Join(coeffs, ", "), x.SevenOver72, x.SixtyFiveOver72, x.SumBoundaryWeights, x.BestWeight, x.BestWeightDeltaFromSevenOver72, x.TypedWeightUniqueInCurrentLedger, x.Verdict)
}

func FormatFunctional(x HistoryWallBalanceFunctional) string {
	return fmt.Sprintf("name=%q formula=%q value=%.15g abs=%.15g threshold=%.15g passes=%t seal=%q interpretation=%q verdict=%q", x.Name, x.Formula, x.Value, x.AbsoluteResidual, x.Threshold, x.PassesBridgeTolerance, x.SealName, x.Interpretation, x.Verdict)
}

func FormatOrientation(x OrientationApproximationAudit) string {
	return fmt.Sprintf("kappaE=%.15g kappaEOrient=%.15g exactResidual=%.15g orientResidual=%.15g growth=%.15g relBoundarySplit=%.15g verdict=%q", x.ExactKappaE, x.OrientationKappaE, x.ExactResidual, x.OrientationResidual, x.ResidualGrowth, x.RelativeToBoundarySplit, x.Verdict)
}

func FormatHessian(x HessianFirewallAudit) string {
	return fmt.Sprintf("quartic=%.15g hessian=%.15g separated=%t statement=%q verdict=%q", x.QuarticWallCoordinate, x.HessianCoordinate, x.KeepsHessianSeparate, x.Statement, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsWallAirlock=%t claims7=%t claimsBoundary=%t claimsScalarZero=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeWallDistanceAirlock, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsNativeScalarZeroBoundary, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
