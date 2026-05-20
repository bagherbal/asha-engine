package generation2koidereactorckmresidualclosureuncertaintyaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("epsObs=%.15g rad %.15g deg kappa=%.15g R=%.15g dR=%.15g Rminus1=%.15g Q=%.15g Qres=%.15g loop=%.15g verdict=%q", a.EpsilonObsRad, a.EpsilonObsDeg, a.KappaObs, a.RObs, a.RDefect, a.RMinusOne, a.QObs, a.QResidual, a.LoopUnit, a.Verdict)
}

func FormatInputs(a OrientationInputs) string {
	return fmt.Sprintf("reactor=%q sin2=%.15g(+%.15g/-%.15g) Jruntime=%.15g Jsource=%q JsrcCentral=%.15g(+%.15g/-%.15g) verdict=%q", a.ReactorSource, a.Sin2Theta13, a.Sin2Theta13Plus, a.Sin2Theta13Minus, a.JCKMRuntime, a.JCKMUncertaintySrc, a.JCKMUncertaintyCtr, a.JCKMPlus, a.JCKMMinus, a.Verdict)
}

func FormatResidual(a CombinedResidual) string {
	return fmt.Sprintf("A=%.15g B=%.15g kappa=%.15g delta590=%.15g abs=%.15g rel=%.15g epsResidual=%.15g rad %.15g deg improvement=%.15g removed=%.15g%% verdict=%q", a.AReactorQuarter, a.BReactorMinusCKM, a.KappaObs, a.Delta590, a.AbsDelta590, a.RelativeDelta590, a.EpsilonResidualRad, a.EpsilonResidualDeg, a.ImprovementOverA, a.PercentMismatchRemoved, a.Verdict)
}

func FormatUncertainty(a UncertaintyAudit) string {
	return fmt.Sprintf("Bcentral=%.15g B1sigma=[%.15g,%.15g] residuals=[%.15g,%.15g,%.15g] covers=%t thetaWidths=[%.15g,%.15g] ckmWidths=[%.15g,%.15g] totalWidths=[%.15g,%.15g] sigmaFractions=[%.15g,%.15g] dominant=%q verdict=%q", a.BCentral, a.BMin1Sigma, a.BMax1Sigma, a.ResidualLow, a.ResidualCentral, a.ResidualHigh, a.CoversKappa, a.Theta13WidthMinus, a.Theta13WidthPlus, a.CKMWidthMinus, a.CKMWidthPlus, a.TotalWidthMinus, a.TotalWidthPlus, a.SigmaFractionMinus, a.SigmaFractionPlus, a.DominantUncertainty, a.Verdict)
}

func FormatInverse(a InversePrediction) string {
	return fmt.Sprintf("sin2Pred=%.15g rangeFromJ=[%.15g,%.15g] thetaPred=%.15g rangeFromJ=[%.15g,%.15g] NuFITsin2=%.15g range=[%.15g,%.15g] NuFITtheta=%.15g range=[%.15g,%.15g] thetaResidual=%.15g within=%t verdict=%q", a.Sin2PredCentral, a.Sin2PredLowFromJ, a.Sin2PredHighFromJ, a.ThetaPredCentralDeg, a.ThetaPredLowDeg, a.ThetaPredHighDeg, a.NuFITCentralSin2, a.NuFITLowSin2, a.NuFITHighSin2, a.NuFITCentralDeg, a.NuFITLowDeg, a.NuFITHighDeg, a.CentralResidualDeg, a.WithinOneSigma, a.Verdict)
}

func FormatDefects(a DefectScaleAudit) string {
	return fmt.Sprintf("delta590=%.15g Rdefect=%.15g Qres=%.15g absQ=%.15g delta/R=%.15g delta/absQ=%.15g smallerR=%t smallerQ=%t interpretation=%q verdict=%q", a.Delta590, a.RDefect, a.QResidual, a.AbsQResidual, a.DeltaOverRDefect, a.DeltaOverAbsQ, a.DeltaSmallerThanR, a.DeltaSmallerThanQ, a.Interpretation, a.Verdict)
}

func FormatCorrectionCandidate(a CorrectionCandidate) string {
	return fmt.Sprintf("name=%q source=%q eq=%q value=%.15g coeff=%.15g correctedDelta=%.15g abs=%.15g relToDelta=%.15g certified=%t", a.Name, a.Source, a.Equation, a.Value, a.Coefficient, a.CorrectedDelta, a.AbsResidual, a.RelativeToDelta, a.Certified)
}

func FormatCorrections(a CorrectionAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, FormatCorrectionCandidate(c))
	}
	return fmt.Sprintf("requiredRcoef=%.15g requiredQcoef=%.15g best={%s} anyCertified=%t candidates=[%s] interpretation=%q verdict=%q", a.RequiredRDefectCoefficient, a.RequiredQResidualCoefficient, FormatCorrectionCandidate(a.BestCandidate), a.AnyCertified, strings.Join(parts, "; "), a.Interpretation, a.Verdict)
}

func FormatLawfulness(a LawfulnessAudit) string {
	return fmt.Sprintf("crossSector=%t Rop=%t Qop=%t rootTrace=%t derivesDelta=%t derivesKappa=%t verdict=%q", a.CrossSectorOrientationIntertwinerPresent, a.RDefectToOrientationOperatorPresent, a.QDefectToOrientationOperatorPresent, a.NativeRootTraceOperatorPresent, a.DerivesDelta590, a.DerivesKappa, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t pmns=%t ckm=%t theta13=%t neutrino=%t masses=%t hierarchy=%t observedNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesPMNS, a.DerivesCKM, a.DerivesTheta13, a.DerivesNeutrinoPhysics, a.DerivesChargedLeptonMasses, a.DerivesFlavorHierarchy, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("deltaMeaningful=%t correctedCertified=%t crossSectorBridge=%t environmental=%t decision=%q verdict=%q", a.DeltaStatisticallyMeaningful, a.CorrectedFormulaCertified, a.CrossSectorBridgePresent, a.KappaRemainsEnvironmental, a.Decision, a.Verdict)
}
