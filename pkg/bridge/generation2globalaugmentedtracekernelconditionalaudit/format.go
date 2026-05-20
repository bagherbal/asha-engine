package generation2globalaugmentedtracekernelconditionalaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate679Inheritance) string {
	return fmt.Sprintf("inherited=%t h72=%d kernel=%d q=%d k7=%d sSplit=%.15g dBase=%.15g tauGlobal=%.15g tauKernel=%.15g tauFinite=%.15g tauHalf=%.15g residual=%.15g firewall=%t verdict=%q", x.RelativeTraceInherited, x.H72Dimension, x.KernelDimension, x.QuotientDimension, x.K7Rank, x.SSplit, x.DBase, x.TauGlobal, x.TauKernel, x.TauFinite, x.TauHalf, x.ResidualGlobal, x.FirewallPreserved, x.Verdict)
}
func FormatSequence(x ShortExactProjectionSequence) string {
	return fmt.Sprintf("sequence=%q kernel=%q dimKernel=%d dimAmbient=%d dimQ=%d exactByDim=%t verdict=%q", x.Sequence, x.KernelFormula, x.KernelDimension, x.AmbientDimension, x.QuotientDimension, x.ExactByDimension, x.Verdict)
}
func FormatDefect(x DefectInclusion) string {
	return fmt.Sprintf("inclusion=%q rank=%d kernel=%d ambient=%d tauKernel=%.15g tauGlobal=%.15g tauFinite=%.15g fullKernel=%t verdict=%q", x.Inclusion, x.Rank, x.KernelDimension, x.AmbientDimension, x.KernelConditionalDensity, x.GlobalDensity, x.FiniteDensity, x.FullKernel, x.Verdict)
}
func FormatNormalization(x TraceNormalization) string {
	return fmt.Sprintf("%s=%s den=%d value=%.15g pred=%.15g residual=%.15g abs=%.15g finite=%t anti=%t quotient=%t class=%q verdict=%q", x.Name, x.Formula, x.Denominator, x.Value, x.PredictedDBase, x.Residual, x.AbsResidual, x.IncludesFiniteChamber, x.IncludesBoundaryAntiLine, x.IncludesQuotientLine, x.Classification, x.Verdict)
}
func FormatNormalizations(xs []TraceNormalization) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatNormalization(x))
	}
	return strings.Join(parts, "; ")
}
func FormatCompatibility(x ResponseCompatibility) string {
	return fmt.Sprintf("domain=%q qIncluded=%t globalKeepsQ=%t kernelExcludesQ=%t finiteExcludesBoundary=%t best=%q verdict=%q", x.ResponseDomain, x.QuotientLineIncludedInSystem, x.GlobalKeepsQuotientInput, x.KernelExcludesQuotientInput, x.FiniteExcludesBoundarySystem, x.BestNormalization, x.Verdict)
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] new=%q support=[%s] verdict=%q", strings.Join(x.Missing, "; "), x.NewPreciseMissingPrinciple, strings.Join(x.AllowedSupport, "; "), x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsKernel=%t claimsFinite=%t claimsGlobalPrinciple=%t claims7=%t claimsTraceQuotient=%t claimsBoundary=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsKernelConditionalTrace, x.ClaimsFiniteOnlyTrace, x.ClaimsNativeGlobalPrinciple, x.ClaimsNativeSevenOver72, x.ClaimsNativeTraceQuotient, x.ClaimsBoundaryStress, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
