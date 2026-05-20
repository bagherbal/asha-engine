package generation2historywallbalancenormalvectorsourceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate670Inheritance) string {
	return fmt.Sprintf("historyWall=%t signed=%t functional=%t normal=%t orientation=%t roles=%t noWall=%t no7=%t noBoundary=%t firewall=%t residual=%.15g verdict=%q", x.HistoryWallBalanceInherited, x.SignedWallFormWritten, x.FunctionalDefined, x.NormalVectorAudited, x.OrientationAudited, x.CoordinateRolesClassified, x.NoNativeWallAirlock, x.NoNativeSevenOver72, x.NoBoundaryStress, x.FirewallPreserved, x.InheritedResidual, x.Verdict)
}

func FormatNormal(x NormalVectorAudit) string {
	pairs := make([]string, 0, len(x.Coefficients))
	for i, c := range x.Coefficients {
		pairs = append(pairs, fmt.Sprintf("%s:%.15g", x.Coordinates[i], c))
	}
	return fmt.Sprintf("label=%q coeffs=[%s] historyUnit=%t boundarySumOne=%t scalarDominant=%t gaugeWeight=%.15g gaugeLabel=%q signedAntiAlignment=%t verdict=%q", x.NormalLabel, strings.Join(pairs, ", "), x.HistorySideUnitWeights, x.BoundaryWeightsSumToOne, x.ScalarBoundaryDominant, x.GaugeBoundaryWeight, x.GaugeBoundaryWeightLabel, x.SignedAntiAlignment, x.Verdict)
}

func FormatDecomposition(x NormalVectorDecompositionAudit) string {
	return fmt.Sprintf("history=%q boundary=%q historyMeaning=%q boundaryMeaning=%q equation=%q verdict=%q", x.HistoryBlockLabel, x.BoundaryBlockLabel, x.HistoryMeaning, x.BoundaryMeaning, x.SplitEquation, x.Verdict)
}

func FormatAlternative(x AlternativeNormal) string {
	return fmt.Sprintf("%s vector=%q scalar=%.15g gauge=%.15g weight=%.15g typed=%q exact=%.15g absExact=%.15g orient=%.15g absOrient=%.15g constraints=%t", x.Name, x.VectorLabel, x.ScalarCoeff, x.GaugeCoeff, x.Weight, x.WeightTyped, x.ResidualExact, x.AbsExact, x.ResidualOrient, x.AbsOrient, x.MeetsConstraints)
}

func FormatMinimality(x MinimalityAudit) string {
	rows := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		rows = append(rows, FormatAlternative(a))
	}
	return fmt.Sprintf("alternatives=%d bestExact=%q %.15g bestOrient=%q %.15g n72BestExact=%t n72BestOrient=%t constraints=[%s] verdict=%q ledger=[%s]", len(x.Alternatives), x.BestExactName, x.BestExactAbsResidual, x.BestOrientationName, x.BestOrientationAbsResidual, x.N72BestAmongTypedExact, x.N72BestAmongTypedOrient, strings.Join(x.TypedConstraints, "; "), x.Verdict, strings.Join(rows, " | "))
}

func FormatCoordinate(x CoordinateNormalizationAudit) string {
	return fmt.Sprintf("coordinates=[%s] sealed=%t warning=%q onlyGate669=%t verdict=%q", strings.Join(x.CanonicalCoordinates, ", "), x.CoordinateSealed, x.RescalingWarning, x.PreservesOnlyGate669WallNormalization, x.Verdict)
}

func FormatOrientation(x ExactVersusOrientationAudit) string {
	return fmt.Sprintf("exactKappaE=%.15g orientKappaE=%.15g exactResidual=%.15g orientResidual=%.15g growth=%.15g orientBest=%q orientBestResidual=%.15g interpretation=%q verdict=%q", x.ExactKappaE, x.OrientationKappaE, x.ExactResidualN72, x.OrientationResidualN72, x.ResidualGrowth, x.OrientationBestTypedName, x.OrientationBestTypedResidual, x.Interpretation, x.Verdict)
}

func FormatScaleLocal(x ScaleLocalAudit) string {
	return fmt.Sprintf("lambda12Global=%t localMin=%t residualAtLambda12=%.15g nearestDelta=%.15g nearestResidual=%.15g n72Local=%t statement=%q verdict=%q", x.Lambda12SelectedInGate662, x.LocalGate662MinimumAtLambda12, x.N72AtLambda12Residual, x.NearestLocalNonzeroDeltaLog, x.NearestLocalNonzeroResidual, x.N72BestTypedNormalOnlyAtLambda12, x.Statement, x.Verdict)
}

func FormatSourceCandidate(x SourceTypeCandidate) string {
	return fmt.Sprintf("%s candidate=%q support=%q firewall=%q", x.Name, x.Candidate, x.Support, x.Firewall)
}

func FormatSource(x SourceTypeAudit) string {
	rows := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		rows = append(rows, FormatSourceCandidate(c))
	}
	return fmt.Sprintf("augmentedTrace=%t boundaryInterpolation=%t historyDeficit=%t coordinateRisk=%t verdict=%q candidates=[%s]", x.AugmentedTraceCandidate, x.BoundaryInterpolationCandidate, x.HistoryDeficitConservationCandidate, x.CoordinateArtifactRisk, x.Verdict, strings.Join(rows, "; "))
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsNormal=%t claims7=%t claimsWallAirlock=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeNormalVectorTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsWallDistanceAirlockTheorem, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
