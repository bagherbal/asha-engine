package generation2koidetransportvectordecompositionaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g logLambdaOverMZ=%.12g predecessor=%q source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.LogLambdaOverMZ, a.PredecessorGate, a.RuntimeSource, a.Verdict)
}

func FormatEndpoint(a KoideCoordinateEndpoint) string {
	return fmt.Sprintf("name=%q rho=%.15g Q=%.15g deltaQ=%.15g theta=%.12g thetaError=%.12g phiSigned=%.12g phi=%.12g verdict=%q", a.Name, a.Rho, a.Q, a.DeltaQ, a.ThetaDeg, a.ThetaErrorDeg, a.PhiSignedDeg, a.PhiDeg, a.Verdict)
}

func FormatTransport(a TransportVector) string {
	return fmt.Sprintf("from={%s} to={%s} deltaT=%.15g deltaRho=%.15g deltaLnRho=%.15g dlnrho_dt=%.15g deltaThetaDeg=%.15g dtheta_dt_deg=%.15g deltaPhiDeg=%.15g dphi_dt_deg=%.15g avgThetaDeg=%.12g projectiveDeltaRad=%.15g projectiveDeltaDeg=%.15g projectiveSpeedRad=%.15g radialToProjective=%.12g thetaImprovement=%.12g qImprovement=%.12g towardCone=%t phiInvariant=%t radialDominant=%t verdict=%q", FormatEndpoint(a.From), FormatEndpoint(a.To), a.DeltaT, a.DeltaRho, a.DeltaLnRho, a.DLnRhoDT, a.DeltaThetaDeg, a.DThetaDTDeg, a.DeltaPhiDeg, a.DPhiDTDeg, a.AverageThetaDeg, a.ProjectiveAngularDelta, radToDeg(a.ProjectiveAngularDelta), a.ProjectiveAngularDT, a.RadialToProjectiveRatio, a.ThetaImprovementFactor, a.QImprovementFactor, a.MovesTowardCone, a.PhiNearlyInvariant, a.RadialDominant, a.Verdict)
}

func FormatDynamics(a DynamicalInterpretation) string {
	return fmt.Sprintf("mostlyRadial=%t coneVisible=%t coneAttractorCertified=%t azimuthPreserved=%t continuousBeta=%t explanation=%q missing=%q verdict=%q", a.MostlyRadialRescaling, a.ConeAttractionVisible, a.ConeAttractorCertified, a.AzimuthPreserved, a.ContinuousBetaCertified, a.Explanation, a.MissingTheorem, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("koide=%t masses=%t yukawas=%t ckm=%t pmns=%t generation=%t newCarrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesKoide, a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.IntroducesNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q dlnrho_dt=%.15g dtheta_dt_deg=%.15g dphi_dt_deg=%.15g thetaImprovement=%.12g qImprovement=%.12g projectiveDeltaDeg=%.15g radialToProjective=%.12g coneAttractorCertified=%t next=%q verdict=%q", a.SealName, a.DlnRhoDt, a.DThetaDtDeg, a.DPhiDtDeg, a.ThetaImprovementFactor, a.QImprovementFactor, a.ProjectiveAngularDeltaDeg, a.RadialToProjectiveRatio, a.ConeAttractorCertified, a.MinimalNextRequirement, a.Verdict)
}
