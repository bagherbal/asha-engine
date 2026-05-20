package generation2historywallbalancenormalvectorsourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HistoryWallBalanceNormalVectorSourceAndMinimalityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 671 — HistoryWallBalance Normal-Vector Source and Minimality Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate671 normal-vector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate670 HistoryWallBalanceSeal", Passed: a.Inherited.HistoryWallBalanceInherited && a.Inherited.SignedWallFormWritten && a.Inherited.FunctionalDefined && a.Inherited.NormalVectorAudited && a.Inherited.OrientationAudited && a.Inherited.CoordinateRolesClassified && a.Inherited.NoNativeWallAirlock && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoBoundaryStress && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define normal vector n_72", Passed: len(a.Normal.Coefficients) == 4 && math.Abs(a.Normal.Coefficients[0]-1) < 1e-15 && math.Abs(a.Normal.Coefficients[1]-1) < 1e-15 && math.Abs(a.Normal.Coefficients[2]-sixtyFiveOver72) < 1e-15 && math.Abs(a.Normal.Coefficients[3]+sevenOver72) < 1e-15 && a.Normal.HistorySideUnitWeights && a.Normal.BoundaryWeightsSumToOne && a.Normal.SignedAntiAlignment && strings.Contains(a.Normal.Verdict, StatusNormalVectorDefined), Detail: FormatNormal(a.Normal)},
			{Name: "audit normal-vector decomposition", Passed: a.Decomposition.HistoryBlockLabel == "(1,1)" && a.Decomposition.BoundaryBlockLabel == "(65/72,-7/72)" && strings.Contains(a.Decomposition.Verdict, StatusNormalVectorDecompositionAudited), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "compare typed alternative normals", Passed: len(a.Minimality.Alternatives) == 6 && a.Minimality.N72BestAmongTypedExact && a.Minimality.BestExactName == "seven over seventy two boundary pull" && math.Abs(a.Minimality.BestExactAbsResidual-8.52583441346e-10) < 1e-14 && strings.Contains(a.Minimality.Verdict, StatusTypedAlternativeNormalsCompared), Detail: FormatMinimality(a.Minimality)},
			{Name: "audit coordinate normalization", Passed: a.Coordinate.CoordinateSealed && a.Coordinate.PreservesOnlyGate669WallNormalization && strings.Contains(a.Coordinate.Verdict, StatusHistoryWallNormalCoordinateSealed), Detail: FormatCoordinate(a.Coordinate)},
			{Name: "audit exact versus orientation kappa_e", Passed: math.Abs(a.Orientation.ExactResidualN72-8.52583441346e-10) < 1e-14 && math.Abs(a.Orientation.OrientationResidualN72-2.77672572133e-6) < 1e-12 && a.Orientation.ResidualGrowth > 0 && strings.Contains(a.Orientation.Verdict, StatusExactVersusOrientationAudited), Detail: FormatOrientation(a.Orientation)},
			{Name: "audit Lambda_12 locality", Passed: a.ScaleLocal.Lambda12SelectedInGate662 && a.ScaleLocal.LocalGate662MinimumAtLambda12 && a.ScaleLocal.N72BestTypedNormalOnlyAtLambda12 && a.ScaleLocal.NearestLocalNonzeroResidual > a.ScaleLocal.N72AtLambda12Residual && strings.Contains(a.ScaleLocal.Verdict, StatusScaleLocalAuditComputed), Detail: FormatScaleLocal(a.ScaleLocal)},
			{Name: "audit source-type candidates", Passed: len(a.Source.Candidates) == 4 && a.Source.AugmentedTraceCandidate && a.Source.BoundaryInterpolationCandidate && a.Source.HistoryDeficitConservationCandidate && a.Source.CoordinateArtifactRisk && strings.Contains(a.Source.Verdict, StatusSourceTypeCandidatesAudited), Detail: FormatSource(a.Source)},
			{Name: "preserve firewalls", Passed: !a.Discipline.ClaimsNativeNormalVectorTheorem && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsWallDistanceAirlockTheorem && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsMassPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate671Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
