package generation2koideloopdeficitpmnsorientationaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("eps=%.15g L=%.15g kappa=%.15g sqrtJCKM=%.15g JCKM=%.15g alpha2Over2pi=%.15g midpoint=%.15g midpointRel=%.15g mu0=%.15g lambda12=%.15g verdict=%q", a.EpsilonRad, a.LoopUnit, a.Kappa, a.SqrtJCKM, a.JCKM, a.Alpha2Over2Pi, a.CKMAlpha2Midpoint, a.CKMAlpha2MidpointRel, a.Mu0GeV, a.Lambda12GeV, a.Verdict)
}

func FormatPMNSInput(a PMNSInput) string {
	return fmt.Sprintf("source=%q version=%q through=%q variant=%q ordering=%q convention=%q sin2t12=%.15g(+%.15g/-%.15g) sin2t23=%.15g(+%.15g/-%.15g) sin2t13=%.15g(+%.15g/-%.15g) delta=%.15g(+%.15g/-%.15g) thetaDeg=(%.15g,%.15g,%.15g) note=%q verdict=%q", a.SourceName, a.SourceVersion, a.DataThrough, a.Variant, a.MassOrdering, a.Convention, a.Sin2Theta12, a.Sin2Theta12Plus, a.Sin2Theta12Minus, a.Sin2Theta23, a.Sin2Theta23Plus, a.Sin2Theta23Minus, a.Sin2Theta13, a.Sin2Theta13Plus, a.Sin2Theta13Minus, a.DeltaCPDeg, a.DeltaCPPlusDeg, a.DeltaCPMinusDeg, a.Theta12Deg, a.Theta23Deg, a.Theta13Deg, a.SourceNote, a.Verdict)
}

func FormatInvariants(a PMNSInvariants) string {
	return fmt.Sprintf("s12=%.15g c12=%.15g s23=%.15g c23=%.15g s13=%.15g c13=%.15g deltaRad=%.15g J=%.15g absJ=%.15g sqrtAbsJ=%.15g s13sq=%.15g alpha2Over2pi=%.15g verdict=%q", a.S12, a.C12, a.S23, a.C23, a.S13, a.C13, a.DeltaRad, a.JPMNS, a.AbsJPMNS, a.SqrtAbsJ, a.S13Squared, a.Alpha2Over2Pi, a.Verdict)
}

func FormatCandidate(a Candidate) string {
	return fmt.Sprintf("name=%q class=%q eq=%q value=%.15g range=[%.15g,%.15g] signed=%.15g abs=%.15g rel=%.15g covers=%t near=%t certified=%t interp=%q", a.Name, a.Class, a.Equation, a.Value, a.Min1Sigma, a.Max1Sigma, a.SignedResidual, a.AbsResidual, a.RelativeResidual, a.CoversKappa, a.Near, a.Certified, a.Interpretation)
}

func FormatCandidateSet(a CandidateSet) string {
	return fmt.Sprintf("target=%.15g count=%d best={%s} bestDirect={%s} bestPMNS={%s} sqrtJ={%s} absJ={%s} s13sq={%s} alpha2={%s} certifiedCount=%d verdict=%q", a.TargetKappa, a.CandidateCount, FormatCandidate(a.Best), FormatCandidate(a.BestDirectPMNS), FormatCandidate(a.BestPMNSAssisted), FormatCandidate(a.SqrtJPMNS), FormatCandidate(a.AbsJPMNS), FormatCandidate(a.S13Squared), FormatCandidate(a.Alpha2Over2Pi), len(a.CertifiedCandidates), a.Verdict)
}

func FormatUncertainty(a UncertaintyAudit) string {
	return fmt.Sprintf("sqrtJRange={%s} alpha2DivC13={%s} anyCovers=%t certifiedUnderUncertainty=%t verdict=%q", FormatCandidate(a.SqrtJPMNSRange), FormatCandidate(a.Alpha2Over2PiDivC13), a.AnyCandidateCovers, a.CertifiedUnderUncertainty, a.Verdict)
}

func FormatCKM(a CKMComparison) string {
	return fmt.Sprintf("sqrtJCKM={%s} midpoint={%s} bestPMNS={%s} directBetter=%t assistedBetter=%t midpointClosest=%t interp=%q verdict=%q", FormatCandidate(a.SqrtJCKM), FormatCandidate(a.CKMAlpha2Midpoint), FormatCandidate(a.BestPMNSAssisted), a.DirectPMNSBetterThanSqrtJCKM, a.PMNSAssistedBetterThanSqrtJCKM, a.MidpointStillClosestNumeric, a.Interpretation, a.Verdict)
}

func FormatDecision(a SourceDecision) string {
	return fmt.Sprintf("pmnsBetter=%t certified=%t midpointSurvives=%t remainsSeal=%t best=%q bestValue=%.15g bestRel=%.15g decision=%q verdict=%q", a.PMNSProducesBetterTypedCandidate, a.AnyCandidateCertified, a.CKMMidpointSurvives, a.KappaRemainsSeal, a.BestCandidateName, a.BestCandidateValue, a.BestCandidateRelativeResidual, a.Decision, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("kappa=%t eps=%t koide=%t pmns=%t nuMass=%t texture=%t chargedMass=%t observedNative=%t carrier=%t gate352=%t verdict=%q", a.DerivesKappa, a.DerivesEpsilon, a.DerivesKoide, a.DerivesPMNS, a.DerivesNeutrinoMasses, a.DerivesFlavorTexture, a.DerivesChargedLeptonMasses, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("kappa=%.15g directCert=%t bestDirect=%q bestDirectRel=%.15g bestPMNS=%q bestPMNSRel=%.15g midpointRel=%.15g anyCert=%t remaining=%q verdict=%q", a.Kappa, a.DirectPMNSCertified, a.BestDirectPMNSName, a.BestDirectPMNSRelativeResidual, a.BestPMNSAssistedName, a.BestPMNSAssistedRelativeResidual, a.CKMMidpointRelativeResidual, a.AnyCertified, a.RemainingSeal, a.Verdict)
}
