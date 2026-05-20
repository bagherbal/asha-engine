package generation2koidenaturalframeaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g gate578PhiMZ=%.12g gate578PhiLambda=%.12g leptonScheme=%q leptonTransport=%q source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.Gate578PhiMZDeg, a.Gate578PhiLDeg, a.LeptonMassScheme, a.LeptonTransport, a.RuntimeSource, a.Verdict)
}

func FormatFrame(a KoideFrame) string {
	return fmt.Sprintf("n=[%s] e1=[%s] e2=[%s] dot(n,e1)=%.3g dot(n,e2)=%.3g dot(e1,e2)=%.3g norms=(%.12g,%.12g,%.12g) rightHanded=%t verdict=%q", floats(a.DemocraticAxis), floats(a.E1), floats(a.E2), a.DotNE1, a.DotNE2, a.DotE1E2, a.NormN, a.NormE1, a.NormE2, a.RightHanded, a.Verdict)
}

func FormatPoint(a FramePoint) string {
	return fmt.Sprintf("name=%q carrier=%q labels=%v values=[%s] root=[%s] rho=%.15g Q=%.15g deltaQ=%.15g theta=%.12g thetaDelta=%.12g phiSigned=%.12g phi=%.12g e1=%.12g e2=%.12g scaleFromPole=%.15g verdict=%q", a.Name, a.Carrier, a.Labels, floats(a.Values), floats(a.RootVector), a.Rho, a.Q, a.DeltaQ, a.ThetaDeg, a.ThetaDeltaDeg, a.PhiSignedDeg, a.PhiDeg, a.CoordinateE1, a.CoordinateE2, a.UniformScaleFromPole, a.Verdict)
}

func FormatComparison(a FrameComparison) string {
	return fmt.Sprintf("pole={%s} MZ={%s} Lambda12={%s} deltaPhiPoleMZ=%.12g deltaPhiMZLambda=%.12g deltaQPoleMZ=%.15g deltaQMZLambda=%.15g lambdaCloser=%t mzEqualPole=%t azimuthStable=%t best=%q verdict=%q", FormatPoint(a.Pole), FormatPoint(a.MZ), FormatPoint(a.Lambda12), a.DeltaPhiPoleToMZDeg, a.DeltaPhiMZToLambdaDeg, a.DeltaQPoleToMZ, a.DeltaQMZToLambda, a.LambdaCloserThanMZ, a.MZEqualPole, a.AzimuthStable, a.BestKoideFrame, a.Verdict)
}

func FormatNatural(a NaturalFrameAudit) string {
	return fmt.Sprintf("question=%q poleNatural=%t mzIndependent=%t boundaryCertified=%t boundaryCleanerV1=%t reason=%q missing=%q verdict=%q", a.Question, a.PoleMassFrameNatural, a.MZYukawaFrameIndependent, a.BoundaryFrameCertified, a.BoundaryFrameCleanerInV1, a.Reason, a.MissingTheorem, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("masses=%t yukawas=%t pmns=%t ckm=%t generation=%t observedNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesPMNS, a.DerivesCKM, a.DerivesGenerationHierarchy, a.PromotesObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q best=%q phiPole=%.12g phiMZ=%.12g phiLambda=%.12g deltaQPole=%.15g deltaQMZ=%.15g deltaQLambda=%.15g certified=%t next=%q verdict=%q", a.SealName, a.BestKoideResidualFrame, a.PolePhiDeg, a.MZPhiDeg, a.Lambda12PhiDeg, a.PoleDeltaQ, a.MZDeltaQ, a.Lambda12DeltaQ, a.NaturalFrameCertified, a.MinimalNextRequirement, a.Verdict)
}

func floats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return strings.Join(parts, ",")
}
