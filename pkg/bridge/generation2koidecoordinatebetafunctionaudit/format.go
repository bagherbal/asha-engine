package generation2koidecoordinatebetafunctionaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g logLambdaOverMZ=%.12g predecessor=%q source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.LogLambdaOverMZ, a.PredecessorGate, a.RuntimeSource, a.Verdict)
}

func FormatFormula(a CoordinateFormulaAudit) string {
	return fmt.Sprintf("state=%q rate=%q dlnrho=%q ds=%q dtheta=%q dphi=%q commonCancels=%t verdict=%q", a.StateEquation, a.RateDefinition, a.DLnRhoFormula, a.DSFormula, a.DThetaFormula, a.DPhiFormula, a.CommonRateCancels, a.Verdict)
}

func FormatEndpointBeta(a EndpointBeta) string {
	return fmt.Sprintf("name=%q rho=%.15g theta=%.12g thetaError=%.12g phi=%.12g y=[%.15g %.15g %.15g] rates=[%.15g %.15g %.15g] common=%.15g split=[%.15g %.15g %.15g] spread=%.15g relSpread=%.12g dlnrho_dt=%.15g dtheta_dt_deg=%.15g dphi_dt_deg=%.15g projectiveSpeedDeg=%.15g commonOnlyProjectiveRad=%.15g exactConeDthetaDeg=%.15g exactConeDphiDeg=%.15g exactConeInvariant=%t towardCone=%t phiSlow=%t verdict=%q", a.Name, a.Rho, a.ThetaDeg, a.ThetaErrorDeg, a.PhiDeg, a.Yukawas[0], a.Yukawas[1], a.Yukawas[2], a.Rates[0], a.Rates[1], a.Rates[2], a.CommonRate, a.FamilySplittings[0], a.FamilySplittings[1], a.FamilySplittings[2], a.RateSpread, a.RelativeSpreadToCommon, a.DLnRhoDT, a.DThetaDTDeg, a.DPhiDTDeg, a.ProjectiveSpeedDeg, a.CommonOnlyProjectiveSpeedRad, a.ExactConeDThetaDTDeg, a.ExactConeDPhiDTDeg, a.ExactConeInvariant, a.PointsTowardCone, a.PhiSlow, a.Verdict)
}

func FormatProjectiveSource(a ProjectiveSourceAudit) string {
	return fmt.Sprintf("commonRateDominates=%t projectiveRequiresSplitting=%t mzSpread=%.15g lambdaSpread=%.15g mzCommonOnlyProjective=%.15g lambdaCommonOnlyProjective=%.15g explanation=%q verdict=%q", a.CommonRateDominatesRadial, a.ProjectiveMotionRequiresRateSplitting, a.MZRateSpread, a.LambdaRateSpread, a.MZCommonOnlyProjectiveSpeed, a.LambdaCommonOnlyProjectiveSpeed, a.Explanation, a.Verdict)
}

func FormatCone(a ConeInvariantAudit) string {
	return fmt.Sprintf("testedTheta=%.12g mzExactConeDthetaDeg=%.15g lambdaExactConeDthetaDeg=%.15g invariant=%t attractor=%t explanation=%q verdict=%q", a.TestedAtExactThetaDeg, a.MZExactConeDThetaDTDeg, a.LambdaExactConeDThetaDTDeg, a.ConeInvariantInV1, a.AttractorCertified, a.Explanation, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t yukawas=%t ckm=%t pmns=%t generation=%t newCarrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.IntroducesNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q mzDthetaDeg=%.15g mzDphiDeg=%.15g lambdaDthetaDeg=%.15g lambdaDphiDeg=%.15g coneInvariant=%t attractor=%t next=%q verdict=%q", a.SealName, a.LocalMZDThetaDTDeg, a.LocalMZDPhiDTDeg, a.LocalLambdaDThetaDTDeg, a.LocalLambdaDPhiDTDeg, a.ConeInvariantInV1, a.AttractorCertified, a.MinimalNextRequirement, a.Verdict)
}
