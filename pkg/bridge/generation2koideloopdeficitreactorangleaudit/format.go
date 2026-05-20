package generation2koideloopdeficitreactorangleaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("epsRad=%.15g epsDeg=%.15g L=%.15g kappa=%.15g alpha2Over2pi=%.15g sqrtJCKM=%.15g priorPMNS=%q %.15g rel=%.15g midpoint=%q %.15g rel=%.15g verdict=%q", a.EpsilonRad, a.EpsilonDeg, a.LoopUnit, a.Kappa, a.Alpha2Over2Pi, a.SqrtJCKM, a.PriorPMNSAssistedName, a.PriorPMNSAssistedValue, a.PriorPMNSAssistedRel, a.CKMAlpha2MidpointName, a.CKMAlpha2MidpointValue, a.CKMAlpha2MidpointRel, a.Verdict)
}

func FormatInput(a ReactorInput) string {
	return fmt.Sprintf("source=%q version=%q through=%q variant=%q ordering=%q convention=%q sin2t13=%.15g(+%.15g/-%.15g) theta13=%.15g range=[%.15g,%.15g] note=%q verdict=%q", a.SourceName, a.SourceVersion, a.DataThrough, a.Variant, a.MassOrdering, a.Convention, a.Sin2Theta13, a.Sin2Theta13Plus, a.Sin2Theta13Minus, a.Theta13Deg, a.Theta13LowDeg, a.Theta13HighDeg, a.SourceNote, a.Verdict)
}

func FormatCandidate(a ReactorCandidate) string {
	return fmt.Sprintf("target=%.15g eq=%q value=%.15g range=[%.15g,%.15g] signed=%.15g abs=%.15g rel=%.15g oneSigma=(-%.15g,+%.15g) covers=%t near=%t certified=%t verdict=%q", a.TargetKappa, a.Equation, a.Value, a.Min1Sigma, a.Max1Sigma, a.SignedResidual, a.AbsResidual, a.RelativeResidual, a.OneSigmaMinus, a.OneSigmaPlus, a.CoversKappa, a.Near, a.Certified, a.Verdict)
}

func FormatInverse(a InversePrediction) string {
	return fmt.Sprintf("sin2Pred=%.15g thetaPred=%.15g thetaCentral=%.15g thetaRange=[%.15g,%.15g] sin2Central=%.15g sin2Range=[%.15g,%.15g] thetaRes=%.15g sin2Res=%.15g withinSin2=%t withinTheta=%t verdict=%q", a.Sin2Theta13Pred, a.Theta13PredDeg, a.Theta13CentralDeg, a.Theta13LowDeg, a.Theta13HighDeg, a.Sin2Central, a.Sin2Low, a.Sin2High, a.ThetaResidualDeg, a.Sin2Residual, a.WithinSin2OneSigma, a.WithinThetaOneSigma, a.Verdict)
}

func FormatEpsilon(a EpsilonPrediction) string {
	return fmt.Sprintf("L=%.15g kappaTarget=%.15g kappaCandidate=%.15g epsTargetRad=%.15g epsTargetDeg=%.15g epsPredRad=%.15g epsPredDeg=%.15g signedRad=%.15g signedDeg=%.15g rel=%.15g epsRangeRad=[%.15g,%.15g] covers=%t verdict=%q", a.LoopUnit, a.KappaTarget, a.KappaCandidate, a.EpsilonTargetRad, a.EpsilonTargetDeg, a.EpsilonPredRad, a.EpsilonPredDeg, a.SignedResidualRad, a.SignedResidualDeg, a.RelativeResidual, a.OneSigmaMinRad, a.OneSigmaMaxRad, a.CoversTargetEpsilon, a.Verdict)
}

func FormatComparison(a ComparisonAudit) string {
	return fmt.Sprintf("reactorRel=%.15g priorPMNS=%q rel=%.15g sqrtJCKMRel=%.15g midpoint=%q rel=%.15g beatsPrior=%t beatsSqrtJ=%t beatsMidpoint=%t midpointStillClosest=%t interp=%q verdict=%q", a.ReactorQuarterRel, a.PriorPMNSAssistedName, a.PriorPMNSAssistedRel, a.SqrtJCKMRel, a.CKMAlpha2MidpointName, a.CKMAlpha2MidpointRel, a.BeatsPriorPMNSAssisted, a.BeatsSqrtJCKM, a.BeatsCKMAlpha2Midpoint, a.CKMMidpointStillClosest, a.Interpretation, a.Verdict)
}

func FormatOperator(a OperatorAudit) string {
	return fmt.Sprintf("quarterClue=%t leptonOp=%t weakOp=%t rootTrace=%t derivesTheta13=%t derivesKappa=%t derivesEps=%t verdict=%q", a.FactorOneQuarterInterpretedAsWeakNormalizationClue, a.NativeLeptonOrientationOperatorPresent, a.NativeWeakDoubletOperatorPresent, a.NativeRootTraceOperatorPresent, a.DerivesTheta13, a.DerivesKappa, a.DerivesEpsilon, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t chargedMasses=%t pmns=%t nuParams=%t theta13=%t texture=%t observedNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesChargedLeptonMasses, a.DerivesPMNS, a.DerivesNeutrinoParameters, a.DerivesTheta13, a.DerivesFlavorTexture, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("candidate=%.15g kappa=%.15g rel=%.15g kappaIn1sigma=%t thetaPredIn1sigma=%t betterPriorPMNS=%t nativeOp=%t remainsSeal=%t decision=%q verdict=%q", a.CandidateValue, a.Kappa, a.RelativeResidual, a.KappaWithinTheta13OneSigma, a.Theta13PredWithinOneSigma, a.BetterThanPriorPMNS, a.AnyNativeOperator, a.KappaRemainsSeal, a.Decision, a.Verdict)
}
