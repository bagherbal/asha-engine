package generation2koidechamberwalloffsetaudit

import "fmt"

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g lambda12=%.12g gate582_delta_mz=%.15g gate582_delta_l12=%.15g Rmz=%.15g Rl12=%.15g source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.Gate582DeltaMZDeg, a.Gate582DeltaL12Deg, a.Gate582RMZ, a.Gate582RL12, a.RuntimeSource, a.Verdict)
}

func FormatChamber(a ChamberGeometryAudit) string {
	return fmt.Sprintf("formula=%q koide=%q order=%v chamber=[%.15g,%.15g] lower=%q upper=%q upperZero=%.15g lowerZero=%.15g epsilon=%q nearWall=%q verdict=%q", a.Formula, a.KoideCircleCondition, a.CanonicalOrder, a.PositiveChamberDeg[0], a.PositiveChamberDeg[1], a.LowerWallLabel, a.UpperWallLabel, a.UpperWallZeroCheck, a.LowerWallZeroCheck, a.WallOffsetDefinition, a.NearWallElectronFormula, a.Verdict)
}

func FormatWallPoint(a WallPoint) string {
	return fmt.Sprintf("name=%q delta=%.15g R=%.15g epsilonDeg=%.15g epsilonRad=%.15g chamberCoord=%.15g distUpperFrac=%.15g inside=%t xeA=%.15g xmuA=%.15g xtauA=%.15g exactWallXeA=%.15g linXeA=%.15g quadXeA=%.15g exactResidual=%.15g linResidual=%.15g quadResidual=%.15g xe/xmu=%.15g me/mmu=%.15g xe/xtau=%.15g me/mtau=%.15g verdict=%q", a.Name, a.DeltaDeg, a.PlaneAmplitudeR, a.EpsilonDeg, a.EpsilonRad, a.NormalizedChamberCoordinate, a.NormalizedDistanceUpperWall, a.InsideCanonicalChamber, a.ElectronRootOverA, a.MuonRootOverA, a.TauRootOverA, a.ExactKoideWallElectronOverA, a.LinearElectronOverA, a.QuadraticElectronOverA, a.ExactWallResidual, a.LinearResidual, a.QuadraticResidual, a.ElectronMuonRootRatio, a.ElectronMuonMassRatio, a.ElectronTauRootRatio, a.ElectronTauMassRatio, a.Verdict)
}

func FormatTransport(a WallTransportAudit) string {
	return fmt.Sprintf("epsMZ=%.15g epsL12=%.15g drift=%.15g absDrift=%.15g RmzMinus1=%.15g Rl12Minus1=%.15g amplitudeTowardOne=%t stable=%t chamberPreserved=%t verdict=%q", a.MZEpsilonDeg, a.LambdaEpsilonDeg, a.SignedDriftDeg, a.AbsDriftDeg, a.MZAmplitudeResidual, a.L12AmplitudeResidual, a.AmplitudeMovesToward1, a.EpsilonStable, a.ChamberPreserved, a.Verdict)
}

func FormatSector(a SectorWallAudit) string {
	return fmt.Sprintf("sector=%q labels=%v delta=%.15g R=%.15g Q=%.15g epsilonTo135=%.15g koideLike=%t wallSeal=%t interpretation=%q", a.Sector, a.Labels, a.DeltaDeg, a.R, a.Q, a.EpsilonTo135, a.KoideLike, a.WallSealValid, a.Interpretation)
}

func FormatQuarks(a QuarkAnalogyAudit) string {
	return fmt.Sprintf("up={%s} down={%s} conclusion=%q verdict=%q", FormatSector(a.Up), FormatSector(a.Down), a.Conclusion, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("masses=%t yukawas=%t koide=%t epsilon=%t ckm=%t pmns=%t generation=%t carrier=%t observedNative=%t gate352=%t verdict=%q", a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesKoide, a.DerivesEpsilon, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.AddsNewCarrier, a.PromotesObservedAsNative, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q epsMZ=%.15g epsL12=%.15g Rmz=%.15g Rl12=%.15g chamber=%q stable=%t nearWall=%t quarkSeal=%t native=%t next=%q verdict=%q", a.SealName, a.MZEpsilonDeg, a.LambdaEpsilonDeg, a.MZPlaneAmplitudeR, a.LambdaPlaneAmplitudeR, a.Chamber, a.WallOffsetStableInV1, a.HierarchyNearWall, a.QuarkWallSealCertified, a.NativeSelectorCertified, a.MinimalNextRequirement, a.Verdict)
}
