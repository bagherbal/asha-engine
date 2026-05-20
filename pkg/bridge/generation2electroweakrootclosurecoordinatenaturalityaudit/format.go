package generation2electroweakrootclosurecoordinatenaturalityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate664Inheritance) string {
	return fmt.Sprintf("dualRoot=%t Lambda12=%.12g t12=%.15g Ksum=%.15g E12=%.15g rootRatio=%.15g dlog=%.15g dE=%.15g transverse=%t noDual=%t no7=%t noUnc=%t noBoundary=%t verdict=%q", x.DualRootInherited, x.Lambda12GeV, x.T12, x.KSum, x.E72AtLambda12, x.ClosureRootRatio, x.ClosureRootDeltaLog, x.DE72Dt, x.TransverseCrossing, x.NoNativeDualRoot, x.NoNativeSevenOver72, x.NoFullUncertainty, x.NoBoundaryStress, x.Verdict)
}

func FormatSeed(x TransportSeed) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g t12=%.15g g1=%.12g gy=%.12g g2=%.12g g3=%.12g init=%d verdict=%q", x.Mu0GeV, x.Lambda12GeV, x.T12, x.G1MZ, x.GYMZ, x.G2MZ, x.G3MZ, len(x.InitialVector), x.Verdict)
}

func FormatCommonRoot(x CommonRootAudit) string {
	return fmt.Sprintf("F=%q E=%q t12=%.15g mu12=%.12g Froot=%.15g Eamp=%.15g wBest=%.15g dw=%.15g pass=%t verdict=%q", x.DefinitionF12, x.DefinitionE72, x.T12Analytic, x.Mu12GeV, x.F12AtRoot, x.E72AmplitudeAtRoot, x.WBestAtRoot, x.WBestMinus7Over72, x.ConditionalRootPass, x.Verdict)
}

func FormatFactorization(x LocalFactorizationAudit) string {
	return fmt.Sprintf("window=%.3g samples=%d cF=%.15g bF=%.15g relF=%.15g cU=%.15g bU=%.15g relU=%.15g ampF=%t invU=%t verdict=%q", x.Window, x.Samples, x.CAmplitudeForF12, x.InterceptAmplitudeF12, x.RelativeResidualF12, x.CAmplitudeForU12, x.InterceptAmplitudeU12, x.RelativeResidualU12, x.AmplitudeFactorLikeF12, x.InverseFactorLikeU12, x.Verdict)
}

func FormatCoordinateRow(x CoordinateRow) string {
	return fmt.Sprintf("%s[%s]: G=%.15g w=%.15g dw=%.15g E7=%.15g dlog=%.15g ratio=%.15g root=%t near7=%t def=%q", x.Name, x.CoordinateClass, x.GaugeResidualAtT12, x.WBestAtT12, x.WBestMinus7Over72, x.E72AtSevenOver72, x.DeltaLogRootFromT12, x.MuZeroOverLambda12, x.RootFoundNearLambda12, x.NearSevenOver72, x.Definition)
}

func FormatCoordinates(x CoordinateFamilyAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatCoordinateRow(r))
	}
	return fmt.Sprintf("rows=%d ampNear=%d invNear=%d robust=%t ampNatural=%t inverseNatural=%t verdict=%q ledger=[%s]", len(x.Rows), x.AmplitudeRowsNearWeight, x.InverseRowsNearWeight, x.CoordinateRobust, x.AmplitudeNatural, x.RGNativeInverseNatural, x.Verdict, strings.Join(rows, "; "))
}

func FormatCoordinateSeal(x CoordinateNaturalityVerdict) string {
	return fmt.Sprintf("classification=%q outcomes=[%s] verdict=%q", x.Classification, strings.Join(x.Outcomes, "; "), x.Verdict)
}

func FormatSource(x SourceTypeInterpretation) string {
	return fmt.Sprintf("interpretations=[%s] verdict=%q", strings.Join(x.Interpretations, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsDual=%t claims7=%t claimsUnc=%t claimsBoundary=%t claimsTransport=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeDualRootTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsFullUncertaintyPropagation, x.ClaimsBoundaryStressDerivation, x.ClaimsNativeTransportTheorem, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
