package generation2koideazimuthenvironmentalorientationaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g Gate577Qe=%.15g Gate577PhiMZ=%.12g Gate577PhiLambda=%.12g source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.Gate577QeMZ, a.Gate577PhiMZDeg, a.Gate577PhiLDeg, a.Source, a.Verdict)
}

func FormatFrame(a AzimuthFrame) string {
	return fmt.Sprintf("n=[%s] e1=[%s] e2=[%s] dot(n,e1)=%.3g dot(n,e2)=%.3g dot(e1,e2)=%.3g norms=(%.12g,%.12g,%.12g) rightHanded=%t convention=%q verdict=%q", floats(a.DemocraticAxis), floats(a.E1), floats(a.E2), a.DotNE1, a.DotNE2, a.DotE1E2, a.NormN, a.NormE1, a.NormE2, a.RightHanded, a.Convention, a.Verdict)
}

func FormatAzimuthPoint(a AzimuthPoint) string {
	return fmt.Sprintf("scale=%s labels=%v y=[%s] root=[%s] rho=%.15g Q=%.15g deltaQ=%.15g theta=%.12g thetaDelta=%.12g parallel=%.12g perp=%.12g coordE1=%.12g coordE2=%.12g phiSigned=%.12g phi=%.12g verdict=%q", a.Scale, a.Labels, floats(a.Yukawas), floats(a.RootVector), a.Rho, a.Q, a.DeltaFromTwoThirds, a.ThetaDeg, a.ThetaDeltaDeg, a.Parallel, a.PerpendicularNorm, a.CoordE1, a.CoordE2, a.PhiSignedDeg, a.PhiDeg, a.Verdict)
}

func FormatTransport(a AzimuthTransport) string {
	return fmt.Sprintf("MZ={%s} Lambda12={%s} deltaPhi=%.12g absDeltaPhi=%.12g deltaQ=%.15g stable1e-3=%t stable1e-2=%t verdict=%q", FormatAzimuthPoint(a.MZ), FormatAzimuthPoint(a.Lambda12), a.DeltaPhiDeg, a.AbsDeltaPhiDeg, a.DeltaQ, a.StableAt1eMinus3Deg, a.StableAt1eMinus2Deg, a.Verdict)
}

func FormatCandidate(a PhaseCandidate) string {
	return fmt.Sprintf("name=%q formula=%q candidate=%.12g distance=%.12g threshold=%.12g certified=%t verdict=%q", a.Name, a.Formula, a.CandidateDeg, a.DistanceDeg, a.ThresholdDeg, a.Certified, a.Verdict)
}

func FormatCandidates(a CandidateAudit) string {
	items := []string{}
	for _, c := range a.Candidates {
		items = append(items, FormatCandidate(c))
	}
	return fmt.Sprintf("phi=%.12g driftScale=%.12g threshold=%.12g nearest=%s distance=%.12g anyCertified=%t candidates=[%s] verdict=%q", a.PointPhiDeg, a.DriftScaleDeg, a.CertificationThresholdDeg, a.NearestRationalTurn, a.NearestRationalDistanceDeg, a.AnyCertified, strings.Join(items, "; "), a.Verdict)
}

func FormatSeal(a AzimuthSeal) string {
	return fmt.Sprintf("name=%q carrier=%q coords=%v cone=%q azimuth=%q rhoMZ=%.15g phiMZ=%.12g phiLambda=%.12g drift=%.12g original=%d constraints=%d remaining=%d native=%t bridge=%t verdict=%q", a.Name, a.Carrier, a.Coordinates, a.ConeConstraint, a.AzimuthDefinition, a.RadiusMZ, a.PhiMZDeg, a.PhiLambda12Deg, a.DriftDeg, a.OriginalPositiveMagnitudes, a.Constraints, a.RemainingContinuousCoordinates, a.NativeDerivation, a.BridgeOnly, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("masses=%t yukawas=%t ckm=%t pmns=%t generation=%t ashaPhase=%t importNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesChargedLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.IdentifiesWithASHAProjectivePhase, a.ImportsObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q phiMZ=%.12g phiSignedMZ=%.12g phiLambda=%.12g deltaPhi=%.12g nearest=%q nearestDistance=%.12g certified=%t native=%t next=%q verdict=%q", a.SealName, a.PhiMZDeg, a.PhiSignedMZDeg, a.PhiLambda12Deg, a.DeltaPhiDeg, a.NearestSimpleCandidate, a.NearestSimpleDistanceDeg, a.CertifiedSimplePhase, a.NativeASHAFlavorDerivation, a.NextRequiredTheorem, a.Verdict)
}

func floats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return strings.Join(parts, ",")
}
