package generation2projectorvaluedboundaryquotientresponsetraceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate682Inheritance) string {
	return fmt.Sprintf("inherited=%t homFirewall=%t dBase=%.15g sSplit=%.15g k7=%d q=%d h72=%d fiber=%d priorResidual=%.15g priorFirewall=%t verdict=%q", x.ResponseFiberInherited, x.HomNotNativeSubspace, x.DBase, x.SSplit, x.K7Dimension, x.QBoundaryDimension, x.H72Dimension, x.FiberDimension, x.PriorResidual, x.PriorFirewallPreserved, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("blocked=%q reason=%q h72=%q hom=%q nativeSubspace=%t projectorRoute=%t verdict=%q", x.BlockedClaim, x.Reason, x.H72Type, x.HomType, x.HomIsNativeSubspace, x.ProjectorRouteAllowed, x.Verdict)
}
func FormatProjector(x ProjectorValuedResponse) string {
	return fmt.Sprintf("projector=%q inEnd=%t rank=%d coord=%q sSplit=%.15g response=%q responseInEnd=%t interp=%q verdict=%q", x.Projector, x.ProjectorInEndH72, x.ProjectorRank, x.BoundaryCoordinate, x.SSplit, x.ResponseEndomorphism, x.ResponseInEndH72, x.Interpretation, x.Verdict)
}
func FormatOrdinary(x OrdinaryTraceResponse) string {
	return fmt.Sprintf("TrP7=%d TrI=%d coeff=%.15g sSplit=%.15g pred=%.15g dBase=%.15g residual=%.15g verdict=%q", x.TraceP7, x.TraceIdentity, x.Coefficient, x.SSplit, x.PredictedDBase, x.DBase, x.Residual, x.Verdict)
}
func FormatHodge(x HodgePolarizedTraceAudit) string {
	return fmt.Sprintf("plus=%d minus=%d ordinaryTrace=%d signedTrace=%d ordinaryCoeff=%.15g signedCoeff=%.15g ordinaryPred=%.15g signedPred=%.15g dBase=%.15g ordinaryResidual=%.15g signedResidual=%.15g activeOrdinary=%t signedFails=%t verdict=%q", x.K7PlusDimension, x.K7MinusDimension, x.OrdinaryTrace, x.SignedTrace, x.OrdinaryCoefficient, x.SignedCoefficient, x.OrdinaryPrediction, x.SignedPrediction, x.DBase, x.OrdinaryResidual, x.SignedResidual, x.ActiveUsesOrdinary, x.SignedFailsActive, x.Verdict)
}
func FormatAlternatives(x DenominatorAlternativeAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s coeff=%.15g pred=%.15g residual=%.15g meaning=%s", a.Name, a.Coefficient, a.Prediction, a.Residual, a.TypedMeaning))
	}
	return fmt.Sprintf("best=%s bestResidual=%.15g alternatives=[%s] verdict=%q", x.BestName, x.BestResidual, strings.Join(parts, " | "), x.Verdict)
}
func FormatClassification(x SourceTypeClassification) string {
	return fmt.Sprintf("sSplit=%q p7=%q rSplit=%q trace=%q dBase=%q verdict=%q", x.SSplit, x.P7, x.RSplit, x.Trace, x.DBase, x.Verdict)
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsHom=%t claimsProjector=%t claimsTrace=%t claims7=%t claimsBoundary=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsHomSubspace, x.ClaimsNativeProjectorTheorem, x.ClaimsNativeTraceTheorem, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStress, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
