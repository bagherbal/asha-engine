package generation2boundaryweighteddeficitclosurestationarityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate662Inheritance) string {
	return fmt.Sprintf("scale=%t grid=%t local=%t weightNear=%t noScale=%t no7=%t noUnc=%t noTransport=%t noBoundary=%t Ksum=%.15g E12=%.15g wBest=%.15g dw=%.15g verdict=%q", x.ScaleSweepInherited, x.Lambda12SelectedInGrid, x.Lambda12SelectedLocally, x.ExactWeightNearSevenOver72, x.NoNativeScaleSelection, x.NoNativeSevenOver72, x.NoFullUncertainty, x.NoNativeTransport, x.NoBoundaryStress, x.KSum, x.E72AtLambda12, x.WBestExact, x.WBestMinusSevenOver72, x.Verdict)
}

func FormatSeed(x TransportSeed) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g t12=%.15g g1=%.12g gy=%.12g g2=%.12g g3=%.12g init=%d verdict=%q", x.Mu0GeV, x.Lambda12GeV, x.T12, x.G1MZ, x.GYMZ, x.G2MZ, x.G3MZ, len(x.InitialVector), x.Verdict)
}

func FormatFunction(x ScaleFunctionAudit) string {
	return fmt.Sprintf("Ksum=%.15g t=%.15g mu=%.12g lambda=%.15g |lambda|=%.15g G=%.15g W72=%.15g E72=%.15g gauge=%q verdict=%q", x.KSum, x.T, x.MuGeV, x.Lambda, x.AbsLambda, x.GaugeResidual, x.W72, x.E72, x.GaugeDefinition, x.Verdict)
}

func FormatDerivative(x FirstDerivativeAudit) string {
	return fmt.Sprintf("lambda=%.15g betaLambda=%.15g dAbs=%.15g dG=%.15g wAbs*dAbs=%.15g wG*dG=%.15g dE=%.15g finite=%.15g stationary=%t crossing=%t verdict=%q", x.Lambda, x.BetaLambda, x.DAbsLambdaDt, x.DGaugeResidualDt, x.DWeightedAbsLambdaDt, x.DWeightedGaugeDt, x.DE72DtAnalytic, x.DE72DtFiniteDifference, x.Stationary, x.ZeroCrossingNotStationary, x.Verdict)
}

func FormatBetaBalance(x BetaBalanceAudit) string {
	return fmt.Sprintf("balanceLeft=%.15g requiredDG=%.15g actualDG=%.15g requiredMinusActual=%.15g sign=%t stationaryReq=%t verdict=%q", x.BalanceLeft, x.RequiredDGaugeDt, x.ActualDGaugeDt, x.RequiredMinusActual, x.SignConsistent, x.StationarityWouldRequire, x.Verdict)
}

func FormatCurvature(x CurvatureAudit) string {
	return fmt.Sprintf("h=%.3g second=%.15g shape=%q width1e-6=%.15g width1e-5=%.15g width1e-4=%.15g slope=%.15g verdict=%q", x.Step, x.SecondDerivative, x.LocalShape, x.ThresholdWidth1eMinus6, x.ThresholdWidth1eMinus5, x.ThresholdWidth1eMinus4, x.FiniteSlopeMagnitude, x.Verdict)
}

func FormatZero(x ZeroScaleAudit) string {
	return fmt.Sprintf("tZero=%.15g dlog=%.15g muZero=%.12g ratio=%.15g Ezero=%.15g aligned=%t verdict=%q", x.TZero, x.DeltaLogFromLambda12, x.MuZeroGeV, x.MuZeroOverLambda12, x.E72AtZero, x.ClosureZeroAligned, x.Verdict)
}

func FormatWeightRow(x WeightVersusScaleRow) string {
	return fmt.Sprintf("dlog=%.3g mu=%.12g |lambda|=%.15g G=%.15g wBest=%.15g dw=%.15g E7over72=%.15g", x.DeltaLog, x.MuGeV, x.AbsLambda, x.GaugeResidual, x.WBestExact, x.WBestMinus7Over72, x.E72AtSevenOver72)
}

func FormatWeightScale(x WeightVersusScaleAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatWeightRow(r))
	}
	return fmt.Sprintf("rows=%d crossesNear=%t sharp=%t verdict=%q ledger=[%s]", len(x.Rows), x.CrossesSevenOver72NearLambda12, x.WeightIsSharpAtLambda12, x.Verdict, strings.Join(rows, "; "))
}

func FormatOrientation(x OrientationStationarityAudit) string {
	return fmt.Sprintf("kappaExact=%.15g kappaOrient=%.15g exactE=%.15g orientE=%.15g orientZeroDlog=%.15g wBestOrient=%.15g verdict=%q", x.KappaEExact, x.KappaEOrientation, x.ExactE72AtLambda12, x.OrientationE72AtLambda12, x.OrientationZeroDeltaLog, x.OrientationWBestAtLambda12, x.Verdict)
}

func FormatSource(x SourceTypeAudit) string {
	return fmt.Sprintf("classification=[%s] verdict=%q", strings.Join(x.Classification, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsScale=%t claims7=%t claimsUnc=%t claimsBoundary=%t claimsTransport=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeScaleSelection, x.ClaimsNativeSevenOver72Theorem, x.ClaimsFullUncertaintyPropagation, x.ClaimsBoundaryStressDerivation, x.ClaimsNativeTransportTheorem, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
