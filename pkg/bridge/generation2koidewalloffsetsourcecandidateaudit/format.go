package generation2koidewalloffsetsourcecandidateaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("source=%q eps583Deg=%.15g eps583Rad=%.15g eps584Deg=%.15g eps584Rad=%.15g predResid=%.15g mu0=%.15g lambda12=%.15g verdict=%q", a.Source, a.Gate583EpsilonDeg, a.Gate583EpsilonRad, a.Gate584SolvedEpsilonDeg, a.Gate584SolvedEpsilonRad, a.Gate584PredictionResid, a.Mu0GeV, a.Lambda12GeV, a.Verdict)
}

func FormatTarget(a EpsilonTarget) string {
	return fmt.Sprintf("definition=%q epsDeg=%.15g epsRad=%.15g closureDeg=%.15g closureRad=%.15g diffDeg=%.15g diffRad=%.15g sieve=%q nearRel=%.15g certifiedRel=%.15g verdict=%q", a.PrimaryDefinition, a.PrimaryEpsilonDeg, a.PrimaryEpsilonRad, a.ExactR1RatioClosureDeg, a.ExactR1RatioClosureRad, a.DifferenceDeg, a.DifferenceRad, a.UseForCandidateSieve, a.NearToleranceRelative, a.CertifiedToleranceRelative, a.Verdict)
}

func FormatCandidate(a Candidate) string {
	return fmt.Sprintf("name=%q class=%q equation=%q value=%.15g signedResidual=%.15g absResidual=%.15g relativeResidual=%.15g near=%t certified=%t interpretation=%q", a.Name, a.Class, a.Equation, a.Value, a.SignedResidual, a.AbsResidual, a.RelativeResidual, a.Near, a.Certified, a.Interpretation)
}

func FormatCandidateSet(a CandidateSet) string {
	near := make([]string, 0, len(a.NearCandidates))
	for _, c := range a.NearCandidates {
		near = append(near, c.Name)
	}
	cert := make([]string, 0, len(a.CertifiedCandidates))
	for _, c := range a.CertifiedCandidates {
		cert = append(cert, c.Name)
	}
	return fmt.Sprintf("target=%.15g count=%d best={%s} near=[%s] certified=[%s] verdict=%q", a.TargetEpsilonRad, a.CandidateCount, FormatCandidate(a.Best), strings.Join(near, ","), strings.Join(cert, ","), a.Verdict)
}

func FormatLoop(a LoopFactorAudit) string {
	return fmt.Sprintf("oneOver8Pi={%s} oneOver4Pi={%s} oneOver16Pi={%s} best={%s} requiredCorrection=%.15g requiredCorrectionPct=%.15g nearOnly=%t verdict=%q", FormatCandidate(a.OneOver8Pi), FormatCandidate(a.OneOver4Pi), FormatCandidate(a.OneOver16Pi), FormatCandidate(a.BestLoop), a.RequiredCorrection, a.RequiredCorrectionPct, a.NearButNotCertified, a.Verdict)
}

func FormatCouplings(a CouplingAudit) string {
	return fmt.Sprintf("alphaEM={%s} sqrtAlphaEM={%s} alphaEMOverPi={%s} gStar2Over8Pi2={%s} alpha2={%s} best={%s} certified=%t verdict=%q", FormatCandidate(a.AlphaEMMZ), FormatCandidate(a.SqrtAlphaEMMZ), FormatCandidate(a.AlphaEMOverPiMZ), FormatCandidate(a.GStarSquaredOver8Pi2), FormatCandidate(a.Alpha2MZ), FormatCandidate(a.BestCoupling), a.Certified, a.Verdict)
}

func FormatResiduals(a ResidualAudit) string {
	return fmt.Sprintf("strong={%s} lambda={%s} deltaSin2={%s} J={%s} sqrtJ={%s} best={%s} certified=%t verdict=%q", FormatCandidate(a.StrongMismatch), FormatCandidate(a.AbsLambdaL12), FormatCandidate(a.AbsDeltaSin2), FormatCandidate(a.JCKM), FormatCandidate(a.SqrtJCKM), FormatCandidate(a.BestResidual), a.Certified, a.Verdict)
}

func FormatDecision(a SourceDecision) string {
	return fmt.Sprintf("best=%q bestValue=%.15g bestRel=%.15g bestAbs=%.15g near=%t certified=%t requirement=%q decision=%q verdict=%q", a.BestCandidateName, a.BestCandidateValue, a.BestRelativeResidual, a.BestAbsResidual, a.NearClue, a.CertifiedSource, a.MinimalNextRequirement, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("epsilon=%t koide=%t masses=%t yukawas=%t ckm=%t pmns=%t generations=%t carrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesEpsilon, a.DerivesKoide, a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.AddsNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q epsRad=%.15g epsDeg=%.15g best=%q bestValue=%.15g bestRel=%.15g certified=%t nearLoop=%t native=%t remaining=%q verdict=%q", a.SealName, a.EpsilonRad, a.EpsilonDeg, a.BestCandidate, a.BestCandidateValue, a.BestCandidateRelativeDiff, a.CandidateCertified, a.NearLoopScaleClue, a.NativeDerivationCertified, a.RemainingSeal, a.Verdict)
}
