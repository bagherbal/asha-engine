package generation2augmenteddefectexactsequenceaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate677Inheritance) string {
	return fmt.Sprintf("trace=%t domain=%t codomain=%t operator=%t sharper=%t tau=%.15g sSplit=%.15g dBase=%.15g residual=%.15g noCouple=%t noTrace=%t no7=%t noWall=%t firewall=%t verdict=%q", x.TraceOperatorInherited, x.DomainDefectDefined, x.CodomainDefectDefined, x.OperatorDefined, x.OperatorSharperThanFit, x.TauDefect, x.SSplit, x.DBase, x.Residual, x.NoNativeTraceCouplesDefects, x.NoNativeTraceResponseTheorem, x.NoNativeSevenOver72, x.NoNativeWallAirlock, x.FirewallPreserved, x.Verdict)
}
func FormatChamber(x AugmentedChamberObject) string {
	return fmt.Sprintf("name=%q formula=%q finite=%d boundary=%d total=%d containsK7=%t containsBoundary=%t verdict=%q", x.Name, x.Formula, x.FiniteDimension, x.BoundaryDimension, x.TotalDimension, x.ContainsK7, x.ContainsBoundary, x.Verdict)
}
func FormatDefect(x InternalDefectObject) string {
	return fmt.Sprintf("name=%q carrier=%q projector=%q rank=%d denom=%d tau=%.15g vectorMap=%t verdict=%q", x.Name, x.Carrier, x.Projector, x.Rank, x.TraceDenominator, x.TauDefect, x.VectorBoundaryMapCertified, x.Verdict)
}
func FormatBoundary(x BoundaryQuotientObject) string {
	return fmt.Sprintf("name=%q plane=%q anti=%q quotient=%q coord=%q dim=%d sSplit=%.15g verdict=%q", x.Name, x.BoundaryPlane, x.AntiLine, x.Quotient, x.Coordinate, x.Dimension, x.SSplit, x.Verdict)
}
func FormatHistory(x HistoryDefectObject) string {
	return fmt.Sprintf("name=%q coord=%q dim=%d dBase=%.15g verdict=%q", x.Name, x.Coordinate, x.Dimension, x.DBase, x.Verdict)
}
func FormatSequence(x ExactSequenceCandidate) string {
	return fmt.Sprintf("seq=%q strict=%t weaker=%t incl=%t proj=%t qToD=%t exact=%t compatible=%t verdict=%q", x.CandidateSequence, x.StrictExactSequenceCertified, x.WeakerDiagramLawful, x.InclusionK7ToH72Typed, x.ProjectionH72ToQBoundaryTyped, x.MapQBoundaryToDHistoryTyped, x.KernelCokernelExactnessCertified, x.DiagramObjectsCompatible, x.Verdict)
}
func FormatTrace(x TraceCompatibility) string {
	return fmt.Sprintf("tau=%.15g sSplit=%.15g dBase=%.15g pred=%.15g residual=%.15g abs=%.15g rel=%.15g quotientNorm=%t verdict=%q", x.TauDefect, x.SSplit, x.DBase, x.PredictedDBase, x.Residual, x.AbsResidual, x.RelativeToDBase, x.QuotientNormalized, x.Verdict)
}
func FormatRequirement(x NonTautologyRequirement) string {
	return fmt.Sprintf("%s status=%q source=%q comment=%q", x.Requirement, x.Status, x.Source, x.Comment)
}
func FormatRequirements(xs []NonTautologyRequirement) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatRequirement(x))
	}
	return strings.Join(parts, "; ")
}
func FormatCandidate(x DiagramSourceCandidate) string {
	return fmt.Sprintf("%s status=%q class=%q comment=%q", x.Candidate, x.Status, x.Classification, x.Comment)
}
func FormatCandidates(xs []DiagramSourceCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatCandidate(x))
	}
	return strings.Join(parts, "; ")
}
func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("targets=[%s] missing=[%s] support=[%s] verdict=%q", strings.Join(x.NativeTheoremTargets, "; "), strings.Join(x.MissingTheorems, "; "), strings.Join(x.AllowedSupport, "; "), x.Verdict)
}
func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsExact=%t claimsTraceQuotient=%t claims7=%t claimsWall=%t claimsBoundary=%t claimsK7Map=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeExactSequenceTheorem, x.ClaimsNativeTraceToQuotientTheorem, x.ClaimsNativeSevenOver72, x.ClaimsNativeWallAirlock, x.ClaimsBoundaryStressDerivation, x.ClaimsFullK7BoundaryMap, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
