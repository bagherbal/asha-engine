package generation2koidereactorckmorientationcombinationaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("epsObsRad=%.15g epsObsDeg=%.15g RObs=%.15g Rminus1=%.15g dR=%.15g QObs=%.15g Qres=%.15g kappa=%.15g loop=%.15g source=%q verdict=%q", a.EpsilonObsRad, a.EpsilonObsDeg, a.RObs, a.RMinusOne, a.RDefect, a.QObs, a.QResidual, a.KappaObs, a.LoopUnit, a.Source, a.Verdict)
}

func FormatReactor(a ReactorInput) string {
	return fmt.Sprintf("source=%q version=%q dataThrough=%q variant=%q ordering=%q sin2=%.15g(+%.15g/-%.15g) theta=%.15g range=[%.15g,%.15g] verdict=%q", a.SourceName, a.SourceVersion, a.DataThrough, a.Variant, a.MassOrdering, a.Sin2Theta13, a.Sin2Theta13Plus, a.Sin2Theta13Minus, a.Theta13CentralDeg, a.Theta13LowDeg, a.Theta13HighDeg, a.Verdict)
}

func FormatCKM(a CKMInput) string {
	return fmt.Sprintf("source=%q version=%q J=%.15g uncertainty=%.15g hasUncertainty=%t verdict=%q", a.SourceName, a.SourceVersion, a.JCKM, a.JCKMUncertainty, a.HasJCKMUncertainty, a.Verdict)
}

func FormatCandidate(a CandidateComparison) string {
	return fmt.Sprintf("name=%q equation=%q value=%.15g range=[%.15g,%.15g] signed=%.15g abs=%.15g rel=%.15g covers=%t certified=%t verdict=%q", a.Name, a.Equation, a.Value, a.Min1Sigma, a.Max1Sigma, a.SignedResidual, a.AbsResidual, a.RelativeResidual, a.CoversKappa, a.Certified, a.Verdict)
}

func FormatCombination(a CombinationAudit) string {
	return fmt.Sprintf("A={%s} B={%s} improvementFactor=%.15g outperforms=%t interpretation=%q verdict=%q", FormatCandidate(a.AReactorQuarter), FormatCandidate(a.BReactorMinusCKM), a.BImprovementFactor, a.BOutperformsA, a.Interpretation, a.Verdict)
}

func FormatEpsilon(a EpsilonPrediction) string {
	return fmt.Sprintf("epsObs=%.15g rad %.15g deg epsA=%.15g rad %.15g deg resA=%.15g rad %.15g deg epsB=%.15g rad %.15g deg resB=%.15g rad %.15g deg improvement=%.15g verdict=%q", a.EpsilonObservedRad, a.EpsilonObservedDeg, a.EpsilonPredA_rad, a.EpsilonPredA_deg, a.ResidualA_rad, a.ResidualA_deg, a.EpsilonPredB_rad, a.EpsilonPredB_deg, a.ResidualB_rad, a.ResidualB_deg, a.ImprovementFactor, a.Verdict)
}

func FormatInverse(a InversePrediction) string {
	return fmt.Sprintf("sin2Pred=%.15g thetaPred=%.15g sin2Central=%.15g range=[%.15g,%.15g] thetaCentral=%.15g range=[%.15g,%.15g] sin2Res=%.15g thetaRes=%.15g withinSin2=%t withinTheta=%t verdict=%q", a.Sin2Theta13Pred, a.Theta13PredDeg, a.Sin2Central, a.Sin2Low, a.Sin2High, a.Theta13CentralDeg, a.Theta13LowDeg, a.Theta13HighDeg, a.Sin2Residual, a.ThetaResidualDeg, a.WithinSin2OneSigma, a.WithinThetaOneSigma, a.Verdict)
}

func FormatUncertainty(a UncertaintyAudit) string {
	return fmt.Sprintf("Btheta13Range=[%.15g,%.15g] coversKappa=%t ckmUncertaintyPresent=%t ckmUncertainty=%.15g fullCertified=%t limitedBy=%q verdict=%q", a.Theta13CandidateBMin, a.Theta13CandidateBMax, a.CoversKappaWithTheta13, a.CKMUncertaintyPresent, a.CKMUncertaintyValue, a.FullUncertaintyCertified, a.PrecisionLimitedBy, a.Verdict)
}

func FormatLawfulness(a SectorLawfulnessAudit) string {
	return fmt.Sprintf("crossSectorIntertwiner=%t leptonKoideOp=%t ckmWallOp=%t rootTrace=%t derivesKappa=%t derivesEps=%t derivesTheta13=%t derivesJ=%t verdict=%q", a.CrossSectorOrientationIntertwinerPresent, a.LeptonOrientationToKoideOperatorPresent, a.CKMToChargedLeptonWallOperatorPresent, a.NativeRootTraceOperatorPresent, a.DerivesKappa, a.DerivesEpsilon, a.DerivesTheta13, a.DerivesJCKM, a.Verdict)
}

func FormatResidual(a ResidualControlAudit) string {
	return fmt.Sprintf("combinedResidual=%.15g Rminus1=%.15g dR=%.15g Qres=%.15g epsShift=%.15g ratioR=%.15g ratioQ=%.15g ratioEps=%.15g requiredRcoef=%.15g requiredQcoef=%.15g typed=%t interpretation=%q verdict=%q", a.CombinedResidual, a.RMinusOne, a.RDefect, a.QResidual, a.EpsilonR1MinusObs, a.RatioToAbsRMinusOne, a.RatioToAbsQResidual, a.RatioToEpsilonShift, a.RequiredRDefectCoefficient, a.RequiredQResidualCoefficient, a.TypedCoefficientPresent, a.Interpretation, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t theta13=%t neutrino=%t masses=%t hierarchy=%t observedNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesTheta13, a.DerivesNeutrinoPhysics, a.DerivesChargedLeptonMasses, a.DerivesFlavorHierarchy, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("BOutperformsA=%t inverseWithin=%t crossSectorBridge=%t environmental=%t residual=%.15g rel=%.15g decision=%q verdict=%q", a.BOutperformsA, a.InverseTheta13WithinOneSigma, a.CrossSectorBridgePresent, a.KappaRemainsEnvironmental, a.RemainingResidual, a.RemainingRelativeResidual, a.Decision, a.Verdict)
}
