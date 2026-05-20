package generation2tracefunctionalnontautologyaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate674Inheritance) string {
	return fmt.Sprintf("traceInherited=%t chamber=%t rank7=%t scalarTrace=%t alternatives=%t residual=%.15g dBase=%.15g sSplit=%.15g qTrace=%.15g rank=%d dimH72=%d k7MapFailed=%t noTrace=%t no7=%t noBoundary=%t firewall=%t verdict=%q", x.TraceCandidateInherited, x.AugmentedChamberDefined, x.RankSevenSourceAudited, x.ScalarTraceCandidateDefined, x.DenominatorAlternativesDone, x.TraceResponseResidual, x.DBase, x.SSplit, x.QTrace, x.RankDefect, x.DimH72, x.FullK7BoundaryMapFailed, x.NoNativeTraceResponseTheorem, x.NoNativeSevenOver72Theorem, x.NoBoundaryStressDerivation, x.FirewallPreserved, x.Verdict)
}

func FormatProjector(x AugmentedChamberProjectorAudit) string {
	return fmt.Sprintf("chamber=%q lambda4=%d boundary=%d total=%d projector=%q boundaryRank=%d rank=%d trace=%q trI=%d trP=%d vectorMapNeeded=%t verdict=%q", x.Chamber, x.Lambda4Dimension, x.BoundaryDimension, x.TotalDimension, x.Projector, x.BoundaryActionRank, x.RankPDefect, x.TraceIdentity, x.TraceOfIdentity, x.TraceOfPDefect, x.BoundaryVectorMapNeeded, x.Verdict)
}

func FormatTrace(x NormalizedDefectTraceAudit) string {
	return fmt.Sprintf("trP=%.15g trI=%.15g tau=%.15g candidate=%q verdict=%q", x.TracePDefect, x.TraceIdentity, x.TauDefect, x.Candidate, x.Verdict)
}

func FormatBoundaryLineCandidate(x BoundaryLineCandidate) string {
	return fmt.Sprintf("%s vector=%s value=%.15g typing=%q class=%q", x.Name, x.Vector, x.Value, x.Typing, x.Classification)
}

func FormatBoundaryLine(x BoundarySplitLineAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatBoundaryLineCandidate(c))
	}
	return fmt.Sprintf("pair=%q chosen={%s} lambda=%.15g r3=%.15g sSplit=%.15g candidates=[%s] verdict=%q", x.BoundaryPair, FormatBoundaryLineCandidate(x.ChosenLine), x.Lambda, x.R3Minus1, x.SSplit, strings.Join(parts, "; "), x.Verdict)
}

func FormatAnsatz(x TraceResponseAnsatzAudit) string {
	return fmt.Sprintf("dBase=%.15g sSplit=%.15g tau=%.15g pred=%.15g residual=%.15g abs=%.15g qPull=%.15g scalarFunctional=%t vectorMap=%t verdict=%q", x.DBase, x.SSplit, x.TauDefect, x.PredictedDBase, x.Residual, x.AbsResidual, x.QPull, x.RequiresScalarFunctional, x.RequiresVectorBoundaryMap, x.Verdict)
}

func FormatCriterion(x NonTautologyCriterion) string {
	return fmt.Sprintf("%s status=%q certified=%t comment=%q", x.Criterion, x.Status, x.Certified, x.Comment)
}

func FormatNonTautology(x NonTautologyAudit) string {
	parts := make([]string, 0, len(x.Criteria))
	for _, c := range x.Criteria {
		parts = append(parts, FormatCriterion(c))
	}
	return fmt.Sprintf("certified=%d required=%d promotable=%t conclusion=%q criteria=[%s] verdict=%q", x.CertifiedCriteriaCount, x.RequiredCriteriaCount, x.PromotableToNativeTheorem, x.Conclusion, strings.Join(parts, "; "), x.Verdict)
}

func FormatSource(x SourceRouteAudit) string {
	return fmt.Sprintf("%s support=%q status=%q class=%q", x.Route, x.Support, x.Status, x.Classification)
}

func FormatSources(xs []SourceRouteAudit) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, FormatSource(x))
	}
	return strings.Join(parts, "; ")
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("targets=[%s] missing=[%s] support=[%s] verdict=%q", strings.Join(x.NativeTheoremTargets, "; "), strings.Join(x.MissingTheorems, "; "), strings.Join(x.AllowedSupport, "; "), x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsTrace=%t claimsActsOnLine=%t claimsAirlock=%t claims7=%t claimsStress=%t claimsK7Map=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsGauge=%t claimsFlavor=%t claimsCKM=%t verdict=%q", x.ClaimsNativeTraceResponse, x.ClaimsTraceActsOnSplitLine, x.ClaimsNativeWallDistanceAirlock, x.ClaimsNativeSevenOver72, x.ClaimsNativeStressSplitPullback, x.ClaimsFullK7BoundaryMap, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsMassPrediction, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.Verdict)
}
