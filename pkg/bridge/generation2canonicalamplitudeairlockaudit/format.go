package generation2canonicalamplitudeairlockaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate665Inheritance) string {
	return fmt.Sprintf("inherited=%t class=%q ampNat=%t robust=%t invNat=%t wAmp=%.15g dwAmp=%.15g wInv=%.15g dwInv=%.15g eAmp=%.15g eInv=%.15g noDual=%t no7=%t noTransport=%t noBoundary=%t verdict=%q", x.CoordinateSealInherited, x.Classification, x.AmplitudeNatural, x.CoordinateRobust, x.RGNativeInverseNatural, x.AmplitudeWBest, x.AmplitudeWBestMinus7Over72, x.InverseWBest, x.InverseWBestMinus7Over72, x.AmplitudeE72, x.InverseE72, x.NoNativeDualRoot, x.NoNativeSevenOver72, x.NoNativeTransport, x.NoBoundaryStress, x.Verdict)
}

func FormatCoordinateRow(x CoordinateLayerRow) string {
	return fmt.Sprintf("%s[%s]: coord=%q G=%.15g w=%.15g dw=%.15g E7=%.15g root=%t near7=%t verdict=%q", x.Layer, x.ResidualName, x.Coordinate, x.ResidualAtLambda12, x.WBest, x.WBestMinus7Over72, x.E72AtSevenOver72, x.RootAligned, x.NearSevenOver72, x.Verdict)
}

func FormatCoordinateStack(x CoordinateStackAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatCoordinateRow(r))
	}
	return fmt.Sprintf("rows=%d ampPass=%t inversePass=%t strengthPass=%t logPass=%t verdict=%q ledger=[%s]", len(x.Rows), x.AmplitudeLayerPasses, x.InverseKineticLayerPasses, x.StrengthLayerPasses, x.LogLayerPasses, x.Verdict, strings.Join(rows, "; "))
}

func FormatKineticToAmplitude(x KineticToAmplitudeAudit) string {
	return fmt.Sprintf("r_g=%.15g invWound=%.15g inv/amp=%.15g g2=%.15g g2/amp=%.15g log=%.15g absLambda=%.15g ampGap=%.15g invGap=%.15g statement=%q verdict=%q", x.AmplitudeResidual, x.InverseFractionalWound, x.InverseOverAmplitude, x.SquaredResidual, x.SquaredOverAmplitude, x.LogResidual, x.ScalarWoundAbsLambda, x.AmplitudeScalarScaleGap, x.InverseScalarScaleGap, x.FirstOrderStatement, x.Verdict)
}

func FormatPatternRow(x RecurringPatternRow) string {
	return fmt.Sprintf("%s: works=%q blocks=%q airlock=%q", x.Lane, x.WorkingCoordinate, x.BlockedOrUncertifiedCoordinate, x.AirlockReading)
}

func FormatPattern(x RecurringAmplitudePatternAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatPatternRow(r))
	}
	return fmt.Sprintf("rows=%d pattern=%q verdict=%q ledger=[%s]", len(x.Rows), x.Pattern, x.Verdict, strings.Join(rows, "; "))
}

func FormatTarget(x AirlockTheoremTarget) string {
	return fmt.Sprintf("native=%q airlock=%q closure=%q missing=%q theorem=%q verdict=%q", x.NativeCoordinate, x.AirlockCoordinate, x.ClosureCoordinate, x.MissingMap, x.CandidateTheorem, x.Verdict)
}

func FormatSource(x SourceTypeVerdict) string {
	return fmt.Sprintf("classification=%q statements=[%s] verdict=%q", x.Classification, strings.Join(x.Statements, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsAirlock=%t claimsDual=%t claims7=%t claimsTransport=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeAmplitudeAirlockTheorem, x.ClaimsNativeDualRootTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativeTransportTheorem, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
