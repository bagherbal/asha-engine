package generation2kinetictoconnectionamplitudeairlockaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate666Inheritance) string {
	return fmt.Sprintf("inherited=%t class=%q ampPass=%t invPass=%t rAmp=%.15g rInv=%.15g inv/amp=%.15g dwAmp=%.15g dwInv=%.15g missingAirlock=%t no7=%t noDual=%t noTransport=%t noBoundary=%t verdict=%q", x.AmplitudeSealInherited, x.Classification, x.AmplitudeLayerPasses, x.InverseKineticLayerPasses, x.AmplitudeResidual, x.InverseFractionalWound, x.InverseOverAmplitude, x.AmplitudeWBestMinus7Over72, x.InverseWBestMinus7Over72, x.MissingAirlockTheorem, x.NoNativeSevenOver72, x.NoNativeDualRoot, x.NoNativeTransport, x.NoBoundaryStress, x.Verdict)
}

func FormatKinetic(x KineticCoordinateAudit) string {
	return fmt.Sprintf("native=%q kinetic=%q rg=%q reason=%q closure=%q verdict=%q", x.NativeCoordinate, x.SpectralKineticForm, x.RGVariable, x.WhyNativeForOneLoopRG, x.ClosureStatus, x.Verdict)
}

func FormatRescaling(x CanonicalFieldRescalingAudit) string {
	return fmt.Sprintf("u=%q map=%q amp=%q relation=%q derivative=%q typed=%t source=%q verdict=%q", x.KineticCoefficientU, x.CanonicalMap, x.ConnectionAmplitude, x.AlgebraicRelation, x.DerivativeRelation, x.AmplitudeCoordinateTyped, x.SourceStatement, x.Verdict)
}

func FormatCoordinateRow(x GaugeCoordinateLayerRow) string {
	return fmt.Sprintf("%s[%s]: coord=%q residual=%.15g w=%.15g dw=%.15g passes=%t verdict=%q", x.Layer, x.ResidualName, x.Coordinate, x.Residual, x.WBest, x.WBestMinus7Over72, x.PassesSevenOver72, x.Verdict)
}

func FormatCoordinates(x GaugeCoordinateComparisonAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatCoordinateRow(r))
	}
	return fmt.Sprintf("rows=%d ampOnly=%t inverseFails=%t closure=%q verdict=%q ledger=[%s]", len(x.Rows), x.AmplitudeOnlyPasses, x.InverseKineticFails, x.ClosureCoordinate, x.Verdict, strings.Join(rows, "; "))
}

func FormatHessian(x ElectroweakHessianSocketAudit) string {
	return fmt.Sprintf("D=%q neutral=%q charged=%q ampObjs=[%s] compatible=%t limitation=%q verdict=%q", x.CovariantDerivativeSocket, x.NeutralHessianShape, x.ChargedWSocket, strings.Join(x.AmplitudeObjects, ", "), x.CompatibleWithClosure, x.Limitation, x.Verdict)
}

func FormatScalar(x ScalarSideTypeAudit) string {
	return fmt.Sprintf("scalar=%q comparedTo=%q typedAs=%q nativeAmp=%t limitation=%q verdict=%q", x.ScalarObject, x.ComparedTo, x.TypedAs, x.NativeAmplitude, x.Limitation, x.Verdict)
}

func FormatPattern(x RootAmplitudeRecurrenceAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, fmt.Sprintf("%s: works=%q blocked=%q airlock=%q", r.Lane, r.WorkingCoordinate, r.BlockedOrUncertifiedCoordinate, r.AirlockReading))
	}
	return fmt.Sprintf("rows=%d pattern=%q verdict=%q ledger=[%s]", len(x.Rows), x.Pattern, x.Verdict, strings.Join(rows, "; "))
}

func FormatTarget(x AirlockTheoremTarget) string {
	return fmt.Sprintf("name=%q domain=%q airlock=%q codomain=%q content=%q status=%q verdict=%q", x.Name, x.Domain, x.Airlock, x.Codomain, x.CandidateContent, x.Status, x.Verdict)
}

func FormatSource(x SourceTypeVerdict) string {
	return fmt.Sprintf("classification=%q statements=[%s] verdict=%q", x.Classification, strings.Join(x.Statements, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsAirlock=%t claims7=%t claimsDual=%t claimsTransport=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsGauge=%t claimsCKM=%t verdict=%q", x.ClaimsNativeKineticAmplitudeTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativeDualRootTheorem, x.ClaimsNativeTransportTheorem, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsGaugeUnification, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
