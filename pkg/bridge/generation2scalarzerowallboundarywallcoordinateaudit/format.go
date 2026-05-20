package generation2scalarzerowallboundarywallcoordinateaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate668Inheritance) string {
	return fmt.Sprintf("inherited=%t class=%q activePair=%q hessianSeparated=%t noScalarAirlock=%t no7=%t noBoundary=%t noTransport=%t verdict=%q", x.ScalarCoordinateInherited, x.Classification, x.ActivePair, x.HessianLayerSeparated, x.NoScalarAirlock, x.NoSevenOver72, x.NoBoundaryStress, x.NoTransport, x.Verdict)
}

func FormatScalar(x ScalarZeroWallAudit) string {
	return fmt.Sprintf("wall=%q lambda=%.15g distance=%.15g below=%t absTyped=%t layer=%q verdict=%q", x.WallEquation, x.SignedLambda, x.DistanceBelowWall, x.IsBelowWall, x.AbsoluteValueTyped, x.CoordinateLayer, x.Verdict)
}

func FormatGauge(x GaugeMeetingWallAudit) string {
	return fmt.Sprintf("wall=%q residual=%.15g above=%t layer=%q coordinate=%q verdict=%q", x.WallEquation, x.GaugeResidual, x.IsAboveWall, x.CoordinateLayer, x.CanonicalCoordinate, x.Verdict)
}

func FormatBoundary(x SignedBoundaryStressAudit) string {
	return fmt.Sprintf("gauge=%.15g lambda=%.15g depth=%.15g xi=%.15g W72=%.15g posResidual=%.15g signedResidual=%.15g equivalent=%t posForm=%q signedForm=%q verdict=%q", x.GaugeExcess, x.ScalarSignedWound, x.ScalarDepth, x.XiBoundary, x.W72, x.ClosureResidualPositiveForm, x.ClosureResidualSignedForm, x.EquivalentFormsAgree, x.PositiveDistanceForm, x.SignedStressForm, x.Verdict)
}

func FormatWallRow(r WallCoordinateRow) string {
	return fmt.Sprintf("%s wall=%q coordinate=%q value=%.15g type=%q role=%q verdict=%q", r.Name, r.Wall, r.Coordinate, r.Value, r.DistanceType, r.Role, r.Verdict)
}

func FormatFlavor(x FlavorWallAnalogyAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatWallRow(r))
	}
	return fmt.Sprintf("rows=%d flavor=%t scalar=%t gauge=%t pattern=%q verdict=%q ledger=[%s]", len(x.Rows), x.FlavorWallSupported, x.ScalarWallSupported, x.GaugeWallSupported, x.RecurringPattern, x.Verdict, strings.Join(rows, "; "))
}

func FormatHessian(x HessianLayerSeparation) string {
	return fmt.Sprintf("quartic=%.15g hessian=%.15g separated=%t relation=%q verdict=%q", x.QuarticWallCoordinate, x.HessianCoordinate, x.LayersSeparated, x.HessianRelation, x.Verdict)
}

func FormatTarget(x MissingTheoremTarget) string {
	return fmt.Sprintf("primary=%q alternate=%q required=[%s] statements=[%s] verdict=%q", x.PrimaryName, x.AlternateName, strings.Join(x.RequiredObjects, "; "), strings.Join(x.Statements, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsWallAirlock=%t claimsScalarZero=%t claims7=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeWallDistanceAirlock, x.ClaimsNativeScalarZeroBoundary, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
