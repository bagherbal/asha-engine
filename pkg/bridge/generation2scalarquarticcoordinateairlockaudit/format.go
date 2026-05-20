package generation2scalarquarticcoordinateairlockaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate667Inheritance) string {
	return fmt.Sprintf("inherited=%t class=%q closure=%q ampPass=%t invFails=%t scalarRuntimeShadow=%t missingKineticAirlock=%t no7=%t noTransport=%t noBoundary=%t verdict=%q", x.ConnectionAmplitudeInherited, x.Classification, x.ClosureCoordinate, x.AmplitudeOnlyPasses, x.InverseKineticFails, x.ScalarSideWasRuntimeShadow, x.MissingKineticAirlock, x.NoNativeSevenOver72, x.NoNativeTransport, x.NoBoundaryStress, x.Verdict)
}

func FormatScalarRow(r ScalarCoordinateRow) string {
	return fmt.Sprintf("%s expr=%q value=%.15g layer=%q typed=%q role=%q verdict=%q", r.Name, r.Expression, r.Value, r.Layer, r.TypedStatus, r.ClosureRole, r.Verdict)
}

func FormatScalars(x ScalarCoordinateFamilyAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatScalarRow(r))
	}
	return fmt.Sprintf("rows=%d active=%q hessian=%q massAmp=%q verdict=%q ledger=[%s]", len(x.Rows), x.ActiveScalarCoordinate, x.HessianCoordinate, x.MassAmplitudeCoordinate, x.Verdict, strings.Join(rows, "; "))
}

func FormatHessian(x HessianDoublingAudit) string {
	return fmt.Sprintf("potential=%q relation=%q hessian=%.15g rAmp=%.15g rInv=%.15g 2rAmp=%.15g invMinus2r=%.15g hessMinusInv=%.15g hessMinus2r=%.15g typed=%t verdict=%q", x.PotentialConvention, x.LowScaleRelation, x.HessianCoordinate, x.AmplitudeResidual, x.InverseKineticWound, x.TwoTimesAmplitudeResidual, x.InverseMinusTwoAmplitude, x.HessianMinusInverse, x.HessianMinusTwoAmplitude, x.TypedAsHessianLayer, x.Verdict)
}

func FormatPairingRow(r GaugeScalarPairingRow) string {
	return fmt.Sprintf("%s gauge=%q %.15g scalar=%q %.15g diff=%.15g rel=%.15g w=%.15g dw7=%.15g passes=%t interp=%q verdict=%q", r.Name, r.GaugeCoordinate, r.GaugeValue, r.ScalarCoordinate, r.ScalarValue, r.SignedDifference, r.RelativeDifference, r.WBest, r.WBestMinus7Over72, r.PassesSevenOver72, r.Interpretation, r.Verdict)
}

func FormatPairings(x GaugeScalarPairingAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, FormatPairingRow(r))
	}
	return fmt.Sprintf("rows=%d ampPass=%t invShadow=%t invClosure=%t massPass=%t verdict=%q ledger=[%s]", len(x.Rows), x.AmplitudePairPasses, x.InverseHessianShadowMagnitude, x.InverseHessianClosurePasses, x.MassAmplitudePairPasses, x.Verdict, strings.Join(rows, "; "))
}

func FormatRetest(x ClosureCoordinateRetest) string {
	return fmt.Sprintf("best=%q diff=%.15g w=%.15g dw7=%.15g selectedBy=%q inverseStatus=%q verdict=%q", x.BestTypedPair, x.BestTypedResidual, x.BestTypedWBest, x.BestTypedWBestMinus7, x.SevenOver72SelectedBy, x.InverseHessianStatus, x.Verdict)
}

func FormatSource(x SourceTypeResult) string {
	return fmt.Sprintf("classification=%q statements=[%s] verdict=%q", x.Classification, strings.Join(x.Statements, "; "), x.Verdict)
}

func FormatPattern(x RootAmplitudeRecurrenceAudit) string {
	rows := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		rows = append(rows, fmt.Sprintf("%s:%s passes=%t", r.Layer, r.Coordinate, r.PassesSevenOver72))
	}
	return fmt.Sprintf("rows=%d pattern=%q verdict=%q ledger=[%s]", len(x.Rows), x.Pattern, x.Verdict, strings.Join(rows, "; "))
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsScalarAirlock=%t claimsBoundary=%t claims7=%t claimsTransport=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeScalarAirlockTheorem, x.ClaimsNativeBoundaryStressTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativeTransportTheorem, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
