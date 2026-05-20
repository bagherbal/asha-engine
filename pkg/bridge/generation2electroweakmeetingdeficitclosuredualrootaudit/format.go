package generation2electroweakmeetingdeficitclosuredualrootaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate663Inheritance) string {
	return fmt.Sprintf("zeroInherited=%t Lambda12=%.12g t12=%.15g Ksum=%.15g E12=%.15g dE=%.15g muZeroRatio=%.15g dlogZero=%.15g noStationary=%t noScale=%t no7=%t noUnc=%t noBoundary=%t verdict=%q", x.ZeroCrossingInherited, x.Lambda12GeV, x.T12, x.KSum, x.E72AtLambda12, x.DE72Dt, x.MuZeroOverLambda12, x.DeltaLogZero, x.NoStationaryClaim, x.NoNativeScale, x.NoNativeSevenOver72, x.NoUncertainty, x.NoBoundaryStress, x.Verdict)
}

func FormatSeed(x TransportSeed) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g t12=%.15g g1=%.12g gy=%.12g g2=%.12g g3=%.12g init=%d verdict=%q", x.Mu0GeV, x.Lambda12GeV, x.T12, x.G1MZ, x.GYMZ, x.G2MZ, x.G3MZ, len(x.InitialVector), x.Verdict)
}

func FormatMeeting(x ElectroweakMeetingAudit) string {
	return fmt.Sprintf("F=%q U=%q t12=%.15g mu12=%.12g Froot=%.15g Uroot=%.15g convention=%q verdict=%q", x.DefinitionF12, x.DefinitionU12, x.T12Analytic, x.Mu12GeV, x.F12AtRoot, x.U12AtRoot, x.GaugeConvention, x.Verdict)
}

func FormatClosureRoot(x ClosureRootAudit) string {
	return fmt.Sprintf("tE=%.15g muE=%.12g Ezero=%.15g bracket=%.3g transverse=%t verdict=%q", x.TClosureZero, x.MuClosureZeroGeV, x.E72AtClosureZero, x.BracketHalfWidth, x.ClosureIsTransverse, x.Verdict)
}

func FormatDualRoot(x DualRootOffsetAudit) string {
	return fmt.Sprintf("dlog=%.15g ratio=%.15g absDelta=%.12g aligned=%t verdict=%q", x.DeltaLogMuEOverMu12, x.MuEOverMu12, x.AbsoluteScaleDelta, x.AlignedInV1, x.Verdict)
}

func FormatTransversality(x TransversalityAudit) string {
	return fmt.Sprintf("dF=%.15g dU=%.15g dE=%.15g Ftrans=%t Utrans=%t Etrans=%t slopeTied=%t verdict=%q", x.DF12DtAtLambda12, x.DU12DtAtLambda12, x.DE72DtAtLambda12, x.F12Transverse, x.U12Transverse, x.E72Transverse, x.SlopeTied, x.Verdict)
}

func FormatProportionality(x ProportionalityAudit) string {
	return fmt.Sprintf("window=%.3g samples=%d cF=%.15g relF=%.15g cU=%.15g relU=%.15g propF=%t propU=%t verdict=%q", x.Window, x.Samples, x.CForF12, x.RelativeResidualF12, x.CForU12, x.RelativeResidualU12, x.ProportionalToF12, x.ProportionalToU12, x.Verdict)
}

func FormatConventionRow(x ResidualConventionRow) string {
	return fmt.Sprintf("%s: G=%.15g E12=%.15g dlog=%.15g ratio=%.15g found=%t def=%q", x.Name, x.GaugeResidualAtT12, x.E72AtT12, x.DeltaLogFromT12, x.MuZeroOverLambda12, x.RootFoundNearLambda12, x.Definition)
}

func FormatConventions(x ResidualConventionAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatConventionRow(r))
	}
	return fmt.Sprintf("rows=%d directPass=%d inversePass=%t stable=%t verdict=%q ledger=[%s]", len(x.Rows), x.DirectCouplingConventionsPass, x.InverseConventionPasses, x.ConventionStable, x.Verdict, strings.Join(rows, "; "))
}

func FormatWeightRow(x WeightRootRow) string {
	return fmt.Sprintf("dlog=%.3g wBest=%.15g dw=%.15g E7=%.15g", x.DeltaLog, x.WBest, x.WBestMinus7Over72, x.E72At7Over72)
}

func FormatWeightRoot(x WeightRootAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatWeightRow(r))
	}
	return fmt.Sprintf("w12=%.15g dw12=%.15g crosses=%t independent=%t verdict=%q ledger=[%s]", x.WBestAtLambda12, x.WBestMinus7Over72AtLambda12, x.CrossesSevenOver72NearLambda, x.WeightIndependentlySelected, x.Verdict, strings.Join(rows, "; "))
}

func FormatSource(x SourceTypeClassification) string {
	return fmt.Sprintf("outcomes=[%s] verdict=%q", strings.Join(x.Outcomes, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsDualRoot=%t claims7=%t claimsUnc=%t claimsBoundary=%t claimsTransport=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeDualRootTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsFullUncertaintyPropagation, x.ClaimsBoundaryStressDerivation, x.ClaimsNativeTransportTheorem, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
