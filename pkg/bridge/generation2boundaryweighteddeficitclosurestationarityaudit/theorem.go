package generation2boundaryweighteddeficitclosurestationarityaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryWeightedDeficitClosureStationarityAndBetaBalanceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 663 — BoundaryWeightedDeficitClosure Stationarity and Beta-Balance Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate663 stationarity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate662 scale-selected closure", Passed: a.Inherited.ScaleSweepInherited && a.Inherited.Lambda12SelectedInGrid && a.Inherited.Lambda12SelectedLocally && a.Inherited.ExactWeightNearSevenOver72 && a.Inherited.NoNativeScaleSelection && a.Inherited.NoNativeSevenOver72 && a.Inherited.NoFullUncertainty && a.Inherited.NoNativeTransport && a.Inherited.NoBoundaryStress && math.Abs(a.Inherited.E72AtLambda12-8.525834413464217e-10) < 5e-17, Detail: FormatInherited(a.Inherited)},
			{Name: "construct v1 transport seed", Passed: a.Seed.Mu0GeV > 0 && a.Seed.Lambda12GeV > 9e13 && a.Seed.T12 > 20 && len(a.Seed.InitialVector) == 13, Detail: FormatSeed(a.Seed)},
			{Name: "define E72 scale function at Lambda12", Passed: a.Function.KSum > 0.049 && a.Function.Lambda < 0 && math.Abs(a.Function.AbsLambda-0.0497009420776833) < 5e-14 && math.Abs(a.Function.GaugeResidual-0.0509933868964996) < 5e-14 && math.Abs(a.Function.E72-8.525835107353608e-10) < 5e-16, Detail: FormatFunction(a.Function)},
			{Name: "audit first derivative and classify crossing", Passed: a.Derivative.DE72DtAnalytic > 9e-4 && math.Abs(a.Derivative.DE72DtAnalytic-a.Derivative.DE72DtFiniteDifference) < 1e-10 && !a.Derivative.Stationary && a.Derivative.ZeroCrossingNotStationary, Detail: FormatDerivative(a.Derivative)},
			{Name: "compute beta-balance equation", Passed: a.BetaBalance.BalanceLeft < -9e-4 && a.BetaBalance.RequiredDGaugeDt < -0.005 && a.BetaBalance.ActualDGaugeDt < -0.01 && a.BetaBalance.RequiredMinusActual > 0.009 && !a.BetaBalance.StationarityWouldRequire, Detail: FormatBetaBalance(a.BetaBalance)},
			{Name: "audit local curvature and zero-crossing width", Passed: a.Curvature.SecondDerivative > 7e-5 && strings.Contains(a.Curvature.LocalShape, "zero-crossing") && a.Curvature.ThresholdWidth1eMinus6 > 0.002 && a.Curvature.ThresholdWidth1eMinus4 > 0.2, Detail: FormatCurvature(a.Curvature)},
			{Name: "solve zero-scale offset", Passed: math.Abs(a.ZeroScale.DeltaLogFromLambda12) < 1e-5 && math.Abs(a.ZeroScale.MuZeroOverLambda12-0.9999991071689) < 1e-9 && math.Abs(a.ZeroScale.E72AtZero) < 1e-12 && a.ZeroScale.ClosureZeroAligned, Detail: FormatZero(a.ZeroScale)},
			{Name: "audit best weight versus scale", Passed: len(a.WeightScale.Rows) == 3 && math.Abs(a.WeightScale.Rows[1].WBestMinus7Over72) < 1e-6 && a.WeightScale.WeightIsSharpAtLambda12 && a.WeightScale.CrossesSevenOver72NearLambda12, Detail: FormatWeightScale(a.WeightScale)},
			{Name: "audit orientation approximation stationarity shift", Passed: math.Abs(a.Orientation.OrientationE72AtLambda12-2.77672572136e-06) < 1e-16 && math.Abs(a.Orientation.OrientationZeroDeltaLog) > 0.002 && math.Abs(a.Orientation.OrientationWBestAtLambda12-0.09937065106106) < 5e-14, Detail: FormatOrientation(a.Orientation)},
			{Name: "classify source type and preserve firewalls", Passed: len(a.Source.Classification) == 4 && !a.Discipline.ClaimsNativeScaleSelection && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsFullUncertaintyPropagation && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsCKMPMNSDerivation && a.Discipline.Verdict == StatusGate663Boundary, Detail: FormatSource(a.Source) + " | " + FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
