package generation2boundaryquotientprojectionkernelaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate678Inheritance) string {
	return fmt.Sprintf("diagram=%t weaker=%t strict=%t tau=%.15g sSplit=%.15g dBase=%.15g residual=%.15g h72=%d k7=%d qdim=%d missingExact=%t missingTrace=%t missing7=%t firewall=%t verdict=%q", x.AugmentedDiagramInherited, x.WeakerDiagramLawful, x.StrictExactnessCertified, x.TauDefect, x.SSplit, x.DBase, x.Residual, x.H72Dimension, x.K7Rank, x.BoundaryQuotientDimension, x.MissingExactSequenceTheorem, x.MissingTraceQuotientTheorem, x.MissingSevenOver72, x.FirewallPreserved, x.Verdict)
}
func FormatProjection(x BoundaryProjection) string {
	return fmt.Sprintf("name=%q domain=%q codomain=%q formula=%q pr=%q sigma=%q dim=%d->%d surjective=%t verdict=%q", x.Name, x.Domain, x.Codomain, x.Formula, x.BoundaryProjection, x.SplitFunctional, x.DomainDimension, x.CodomainDimension, x.Surjective, x.Verdict)
}
func FormatKernel(x ProjectionKernel) string {
	return fmt.Sprintf("formula=%q finite=%d anti=%d total=%d k7=%d full=%t inside=%t qdim=%d verdict=%q", x.Formula, x.FiniteKernelDimension, x.AntiLineDimension, x.TotalKernelDimension, x.K7Rank, x.K7IsFullKernel, x.K7InsideKernel, x.QuotientDimension, x.Verdict)
}
func FormatDefect(x DefectInsideKernel) string {
	return fmt.Sprintf("defect=%q projector=%q rank=%d kernel=%d ambient=%d relKernel=%.15g relAmbient=%.15g distinguished=%t full=%t verdict=%q", x.Defect, x.Projector, x.Rank, x.KernelDimension, x.AmbientDimension, x.RelativeToKernel, x.RelativeToAmbient, x.Distinguished, x.FullKernel, x.Verdict)
}
func FormatTrace(x RelativeTraceResponse) string {
	return fmt.Sprintf("tauGlobal=%.15g tauKernel=%.15g tauFinite=%.15g tauHalf=%.15g sSplit=%.15g dBase=%.15g pred=%.15g residual=%.15g abs=%.15g best=%q global=%t verdict=%q", x.TauGlobal, x.TauKernel, x.TauFinite, x.TauHalf, x.SSplit, x.DBase, x.PredictedDBase, x.Residual, x.AbsResidual, x.BestAlternative, x.UsesGlobalH72, x.Verdict)
}
func FormatAlternative(x TraceAlternative) string {
	return fmt.Sprintf("%s=%s value=%.15g pred=%.15g residual=%.15g abs=%.15g class=%q verdict=%q", x.Name, x.Formula, x.Value, x.PredictedDBase, x.Residual, x.AbsResidual, x.Classification, x.Verdict)
}
func FormatAlternatives(xs []TraceAlternative) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatAlternative(x))
	}
	return strings.Join(parts, "; ")
}
func FormatCondition(x NonTautologyCondition) string {
	return fmt.Sprintf("%s status=%q comment=%q", x.Principle, x.Status, x.Comment)
}
func FormatConditions(xs []NonTautologyCondition) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCondition(x))
	}
	return strings.Join(parts, "; ")
}
func FormatCandidate(x SourceCandidate) string {
	return fmt.Sprintf("%s status=%q class=%q comment=%q", x.Candidate, x.Status, x.Classification, x.Comment)
}
func FormatCandidates(xs []SourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCandidate(x))
	}
	return strings.Join(parts, "; ")
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] new=%q support=[%s] verdict=%q", strings.Join(x.Missing, "; "), x.NewPreciseMissingPrinciple, strings.Join(x.AllowedSupport, "; "), x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsK7Kernel=%t claimsExact=%t claimsGlobalTrace=%t claimsTraceQuotient=%t claims7=%t claimsWall=%t claimsBoundary=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsK7KernelOfPiSplit, x.ClaimsLiteralExactSequence, x.ClaimsNativeGlobalTraceTheorem, x.ClaimsNativeTraceQuotientTheorem, x.ClaimsNativeSevenOver72, x.ClaimsWallDistanceAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
