package generation2koidereactorrobustnessrdefectsensitivityaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("epsObsRad=%.15g epsObsDeg=%.15g RObs=%.15g dR=%.15g QObs=%.15g Qres=%.15g kObs=%.15g epsR1Rad=%.15g epsR1Deg=%.15g kR1=%.15g deltaK=%.15g deltaEps=%.15g source=%q verdict=%q", a.EpsilonObsRad, a.EpsilonObsDeg, a.RObs, a.RDefect, a.QObs, a.QResidual, a.KappaObs, a.EpsilonR1Rad, a.EpsilonR1Deg, a.KappaR1, a.KappaShiftObsMinusR1, a.EpsilonShiftR1MinusObs, a.Source, a.Verdict)
}

func FormatReactor(a ReactorInput) string {
	return fmt.Sprintf("source=%q version=%q variant=%q ordering=%q sin2=%.15g(+%.15g/-%.15g) candidate=%.15g range=[%.15g,%.15g] theta=%.15g range=[%.15g,%.15g] verdict=%q", a.SourceName, a.SourceVersion, a.Variant, a.MassOrdering, a.Sin2Theta13, a.Sin2Theta13Plus, a.Sin2Theta13Minus, a.Candidate, a.CandidateMin, a.CandidateMax, a.Theta13CentralDeg, a.Theta13LowDeg, a.Theta13HighDeg, a.Verdict)
}

func FormatKappaComparison(a KappaComparison) string {
	return fmt.Sprintf("name=%q eps=%.15g kappa=%.15g candidate=%.15g range=[%.15g,%.15g] signed=%.15g abs=%.15g rel=%.15g covered=%t certified=%t sin2Pred=%.15g thetaPred=%.15g thetaRes=%.15g withinSin2=%t withinTheta=%t epsPred=%.15g epsRes=%.15g verdict=%q", a.Name, a.EpsilonRad, a.Kappa, a.Candidate, a.CandidateMin, a.CandidateMax, a.SignedResidual, a.AbsResidual, a.RelativeResidual, a.CoveredByOneSigma, a.Certified, a.Sin2Theta13Pred, a.Theta13PredDeg, a.Theta13ResidualDeg, a.WithinSin2OneSigma, a.WithinThetaOneSigma, a.EpsilonPredRad, a.EpsilonResidualRad, a.Verdict)
}

func FormatRobustness(a ReactorRobustnessAudit) string {
	return fmt.Sprintf("observed={%s} exactR1={%s} observedBetter=%t weakerFactor=%.15g interpretation=%q verdict=%q", FormatKappaComparison(a.Observed), FormatKappaComparison(a.ExactR1), a.ObservedBetter, a.ExactR1WeakerFactor, a.Interpretation, a.Verdict)
}

func FormatRDefect(a RDefectCorrectionAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, fmt.Sprintf("%s:value=%.15g pred=%.15g signed=%.15g rel=%.15g closer=%t", c.Name, c.Value, c.PredictedShift, c.SignedResidual, c.RelativeResidual, c.CloserThanNoShift))
	}
	return fmt.Sprintf("dR=%.15g Rminus1=%.15g deltaK=%.15g requiredC=%.15g best={%s:value=%.15g rel=%.15g} certified=%t candidates=[%s] interpretation=%q verdict=%q", a.DROneMinusR, a.RMinusOne, a.KappaObsMinusR1, a.RequiredC, a.BestCandidate.Name, a.BestCandidate.Value, a.BestCandidate.RelativeResidual, a.BestCandidateCertified, strings.Join(parts, "; "), a.Interpretation, a.Verdict)
}

func FormatShift(a ShiftControlAudit) string {
	return fmt.Sprintf("deltaK=%.15g deltaEps=%.15g 8piDeltaEps=%.15g dR=%.15g ratioD=%.15g Qres=%.15g ratioQ=%.15g epsControl=%t rTyped=%t qTyped=%t verdict=%q", a.KappaShift, a.EpsilonShiftR1MinusObs, a.EightPiEpsilonShift, a.DROneMinusR, a.RatioToDROneMinusR, a.QResidual, a.RatioToAbsQResidual, a.ControlledByEpsilonShift, a.ControlledByRDefectTyped, a.ControlledByQResidualTyped, a.Verdict)
}

func FormatOperator(a OperatorAudit) string {
	return fmt.Sprintf("koideReactorOp=%t rDefectOp=%t rootTrace=%t derivesTheta13=%t derivesKappa=%t derivesEps=%t verdict=%q", a.NativeKoideReactorOperatorPresent, a.NativeRDefectCorrectionOperatorPresent, a.NativeRootTraceOperatorPresent, a.DerivesTheta13, a.DerivesKappa, a.DerivesEpsilon, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t theta13=%t pmns=%t neutrino=%t masses=%t flavor=%t observedNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesTheta13, a.DerivesPMNS, a.DerivesNeutrinoPhysics, a.DerivesLeptonMasses, a.DerivesFlavorLaw, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("observedBetter=%t obs1sigma=%t r1sigma=%t rDefectRequired=%t typedRCorrection=%t nativeOp=%t environmental=%t decision=%q verdict=%q", a.ReactorMatchesObservedBetter, a.ObservedInsideOneSigma, a.ExactR1InsideOneSigma, a.RDefectRequiredForBestMatch, a.TypedRDefectCorrectionPresent, a.NativeOperatorPresent, a.RelationRemainsEnvironmental, a.Decision, a.Verdict)
}
