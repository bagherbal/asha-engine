package generation2koidewalloffsetratioclosureaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("source=%q epsMZ=%.15g epsL12=%.15g Rmz=%.15g Rl12=%.15g verdict=%q", a.Source, a.MZEpsilonDeg, a.LambdaEpsilonDeg, a.MZPlaneAmplitudeR, a.LambdaPlaneAmplitudeR, a.Verdict)
}

func FormatModel(a ExactR1WallModel) string {
	return fmt.Sprintf("chamber=%q E=%q M=%q T=%q ratio=%q unknown=%q domain=[%.15g,%.15g] verdict=%q", a.Chamber, a.ElectronOverAFormula, a.MuonOverAFormula, a.TauOverAFormula, a.RatioEquation, a.Unknown, a.UniqueDomainDeg[0], a.UniqueDomainDeg[1], a.Verdict)
}

func FormatPrediction(a RatioPrediction) string {
	return fmt.Sprintf("input=%q inputRatio=%.15g solvedEpsDeg=%.15g observedEpsDeg=%.15g epsResidualDeg=%.15g predicts=%q predictedRoot=%.15g observedRoot=%.15g rootResidual=%.15g relRootResidual=%.15g predictedMass=%.15g observedMass=%.15g massResidual=%.15g within=%t verdict=%q", a.InputRatioName, a.InputRatio, a.SolvedEpsilonDeg, a.ObservedEpsilonDeg, a.EpsilonResidualDeg, a.PredictedRatioName, a.PredictedRootRatio, a.ObservedRootRatio, a.RootResidual, a.RelativeRootResidual, a.PredictedMassRatio, a.ObservedMassRatio, a.MassResidual, a.WithinClosureTolerance, a.Verdict)
}

func FormatClosure(a RatioClosure) string {
	return fmt.Sprintf("frame=%q eps=%.15g R=%.15g actualEMu=%.15g actualMuTau=%.15g exactAtEpsEMu=%.15g exactAtEpsMuTau=%.15g exactEMuResidual=%.15g exactMuTauResidual=%.15g fromEMu={%s} fromMuTau={%s} certified=%t verdict=%q", a.Frame, a.ObservedEpsilonDeg, a.ObservedPlaneAmplitudeR, a.ActualElectronMuonRootRatio, a.ActualMuonTauRootRatio, a.ExactR1AtObservedEpsilonElectronMuonRootRatio, a.ExactR1AtObservedEpsilonMuonTauRootRatio, a.ExactR1AtObservedEpsilonElectronMuonResidual, a.ExactR1AtObservedEpsilonMuonTauResidual, FormatPrediction(a.FromElectronMuon), FormatPrediction(a.FromMuonTau), a.ClosureCertified, a.Verdict)
}

func FormatTransport(a TransportAudit) string {
	return fmt.Sprintf("epsFromEMuMZ=%.15g epsFromEMuL12=%.15g drift=%.15g mzPredResidual=%.15g l12PredResidual=%.15g improves=%t stable=%t verdict=%q", a.MZEpsilonFromEMuDeg, a.LambdaEpsilonFromEMuDeg, a.EMuSolvedEpsilonDriftDeg, a.MZPredictionResidual, a.LambdaPredictionResidual, a.ResidualImprovesAtBoundary, a.ClosureStable, a.Verdict)
}

func FormatQuarks(a QuarkClosureAudit) string {
	return fmt.Sprintf("upR=%.15g downR=%.15g upOnKoide=%t downOnKoide=%t closure=%t interpretation=%q verdict=%q", a.UpR, a.DownR, a.UpOnKoideCircle, a.DownOnKoideCircle, a.OneParameterClosure, a.Interpretation, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t epsilon=%t masses=%t yukawas=%t ckm=%t pmns=%t generations=%t carrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesEpsilon, a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.AddsNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q input=%q epsMZ=%.15g predMuTau=%.15g obsMuTau=%.15g residual=%.15g epsL12=%.15g residualL12=%.15g closure=%t native=%t remaining=%q verdict=%q", a.SealName, a.MZInputRatio, a.MZSolvedEpsilonDeg, a.MZPredictedMuonTauRatio, a.MZObservedMuonTauRatio, a.MZPredictionResidual, a.LambdaSolvedEpsilonDeg, a.LambdaPredictionResidual, a.OneParameterClosure, a.NativeDerivationCertified, a.MinimalRemainingSeal, a.Verdict)
}
