package generation2defectquotientresponsefibertypingaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate681Inheritance) string {
	return fmt.Sprintf("inherited=%t k7=%d q=%d h72=%d density=%.15g dBase=%.15g sSplit=%.15g residual=%.15g firewall=%t verdict=%q", x.PrimitiveDensityInherited, x.K7Dimension, x.QBoundaryDimension, x.H72Dimension, x.Density, x.DBase, x.SSplit, x.Residual, x.FirewallPreserved, x.Verdict)
}
func FormatFiber(x ResponseFiberCandidate) string {
	return fmt.Sprintf("fiber=%q dual=%q qDimOne=%t k7=%d q=%d dim=%d interp=%q verdict=%q", x.Fiber, x.DualForm, x.IsomorphicSinceQDimOne, x.K7Dimension, x.QBoundaryDimension, x.FiberDimension, x.Interpretation, x.Verdict)
}
func FormatProductDensity(x ProductDensityAudit) string {
	return fmt.Sprintf("k7=%d q=%d h72=%d product=%d density=%.15g matches=%t verdict=%q", x.K7Dimension, x.QBoundaryDimension, x.H72Dimension, x.ProductDimension, x.Density, x.MatchesGate681Density, x.Verdict)
}
func FormatDirectTensor(x DirectSumTensorProductAudit) string {
	return fmt.Sprintf("h72=%q k7Certified=%t qCertified=%t fiberNative=%t requiresMap=%t class=%q verdict=%q", x.H72Structure, x.K7SubspaceCertified, x.QBoundaryQuotientCertified, x.FiberIsNativeSubspace, x.RequiresCouplingMap, x.Classification, x.Verdict)
}
func FormatTrace(x TraceDensityAudit) string {
	return fmt.Sprintf("bare=%q bareRank=%d respRank=%d h72=%d bare=%.15g resp=%.15g same=%t gain=%q verdict=%q", x.BareProjector, x.BareProjectorRank, x.ResponseFiberRank, x.H72Dimension, x.BareTraceDensity, x.ResponseFiberTraceDensity, x.SameNumericalDensity, x.TypeGain, x.Verdict)
}
func FormatAction(x SplitActionAudit) string {
	return fmt.Sprintf("operator=%q domain=%q codomain=%q coeff=%.15g sSplit=%.15g pred=%.15g dBase=%.15g residual=%.15g interp=%q verdict=%q", x.Operator, x.DomainCoordinate, x.Codomain, x.Coefficient, x.SSplit, x.PredictedDBase, x.DBase, x.Residual, x.Interpretation, x.Verdict)
}
func FormatCriteria(x NonTautologyCriteria) string {
	return fmt.Sprintf("qCanonical=%t k7Canonical=%t fiberCanonical=%t h72Canonical=%t controlsD=%t missing=[%s] verdict=%q", x.CanonicalQBoundary, x.CanonicalK7Carrier, x.CanonicalResponseFiber, x.CanonicalH72Normalization, x.TypedReasonControlsDHistory, strings.Join(x.StillMissing, "; "), x.Verdict)
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] new=%q support=[%s] verdict=%q", strings.Join(x.Missing, "; "), x.NewPreciseMissingPrinciple, strings.Join(x.AllowedSupport, "; "), x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsFiber=%t claimsNativeSubspace=%t claimsTraceQuotient=%t claims7=%t claimsBoundary=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsResponseFiberTheorem, x.ClaimsNativeSubspace, x.ClaimsTraceQuotientTheorem, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStress, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
