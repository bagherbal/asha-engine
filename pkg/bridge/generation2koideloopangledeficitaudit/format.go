package generation2koideloopangledeficitaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("epsDeg=%.15g epsRad=%.15g loopUnit=%.15g gate585Best=%q gate585Rel=%.15g kappaFromGate585=%.15g J=%.15g sqrtJ=%.15g alpha2=%.15g alpha2Over2pi=%.15g projDrift=%.15g deltaPhi=%.15g deltaTheta=%.15g deltaEps=%.15g Rminus1=%.15g mu0=%.15g lambda12=%.15g verdict=%q", a.EpsilonDeg, a.EpsilonRad, a.LoopUnit, a.Gate585BestCandidate, a.Gate585BestRelative, a.KappaFromGate585, a.JCKM, a.SqrtJCKM, a.Alpha2MZ, a.Alpha2Over2Pi, a.ProjectiveDriftRad, a.DeltaPhiRad, a.DeltaThetaRad, a.DeltaEpsilonRad, a.KoideAmplitudeResidual, a.Mu0GeV, a.Lambda12GeV, a.Verdict)
}

func FormatDefinition(a DeficitDefinition) string {
	return fmt.Sprintf("formula=%q L=%.15g epsRad=%.15g epsDeg=%.15g kappa=%.15g kappaPct=%.15g reconstructed=%.15g error=%.15g interp=%q near=%.15g certified=%.15g verdict=%q", a.Formula, a.LoopUnit, a.EpsilonRad, a.EpsilonDeg, a.Kappa, a.KappaPercent, a.ReconstructedEpsRad, a.ReconstructionError, a.Interpretation, a.NearTolerance, a.CertifiedTolerance, a.Verdict)
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
	return fmt.Sprintf("target=%.15g count=%d best={%s} near=[%s] certified=[%s] verdict=%q", a.TargetKappa, a.CandidateCount, FormatCandidate(a.Best), strings.Join(near, ","), strings.Join(cert, ","), a.Verdict)
}

func FormatOrientation(a OrientationAudit) string {
	return fmt.Sprintf("J={%s} sqrtJ={%s} best={%s} nearOnly=%t pmnsInput=%t verdict=%q", FormatCandidate(a.JCKM), FormatCandidate(a.SqrtJCKM), FormatCandidate(a.BestOrientation), a.NearButNotSource, a.PMNSRuntimeInput, a.Verdict)
}

func FormatTransport(a TransportAudit) string {
	return fmt.Sprintf("projective={%s} dphi={%s} dtheta={%s} deps={%s} Rminus1={%s} best={%s} certified=%t verdict=%q", FormatCandidate(a.ProjectiveDrift), FormatCandidate(a.DeltaPhi), FormatCandidate(a.DeltaTheta), FormatCandidate(a.DeltaEpsilon), FormatCandidate(a.KoideRMinusOne), FormatCandidate(a.BestTransport), a.Certified, a.Verdict)
}

func FormatCorrections(a CorrectionScaleAudit) string {
	return fmt.Sprintf("alpha2Over2Pi={%s} alphaEM={%s} alphaEMOverPi={%s} strongOver2Pi={%s} deltaSin2Over8Pi={%s} lambdaOver2Pi={%s} best={%s} certified=%t verdict=%q", FormatCandidate(a.Alpha2Over2Pi), FormatCandidate(a.AlphaEM), FormatCandidate(a.AlphaEMOverPi), FormatCandidate(a.StrongMismatchOver2Pi), FormatCandidate(a.DeltaSin2Over8Pi), FormatCandidate(a.LambdaOver2Pi), FormatCandidate(a.BestCorrection), a.Certified, a.Verdict)
}

func FormatDecision(a SourceDecision) string {
	return fmt.Sprintf("best=%q bestValue=%.15g bestRel=%.15g bestAbs=%.15g near=%t certified=%t meaning=%q requirement=%q decision=%q verdict=%q", a.BestCandidateName, a.BestCandidateValue, a.BestRelativeResidual, a.BestAbsResidual, a.NearClue, a.CertifiedSource, a.CandidateMeaning, a.MinimalNextRequirement, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("kappa=%t epsilon=%t koide=%t masses=%t yukawas=%t ckm=%t pmns=%t carrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesKappa, a.DerivesEpsilon, a.DerivesKoide, a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.AddsNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q epsRad=%.15g L=%.15g kappa=%.15g best=%q bestValue=%.15g bestRel=%.15g certified=%t orientClue=%t couplingClue=%t native=%t remaining=%q verdict=%q", a.SealName, a.EpsilonRad, a.LoopUnit, a.Kappa, a.BestCandidate, a.BestCandidateValue, a.BestCandidateRelativeDiff, a.CandidateCertified, a.NearOrientationClue, a.NearCouplingClue, a.NativeDerivationCertified, a.RemainingSeal, a.Verdict)
}
