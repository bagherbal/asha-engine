package generation2koideyukawasquarerootconesealaudit

import (
	"fmt"
	"strings"
)

func FormatRuntime(a RuntimeInheritance) string {
	return fmt.Sprintf("mu0=%.12g Lambda12=%.12g J=%.12g KoideQe=%.15g source=%q verdict=%q", a.Mu0GeV, a.Lambda12GeV, a.JCKM, a.KoideQe, a.Source, a.Verdict)
}

func FormatGeometry(a GeometryDefinition) string {
	return fmt.Sprintf("root=%q Q=%q n=%q cos=%q equivalence=%q targetQ=%.12g targetAngle=%.12g positive=%t endpoint=%t verdict=%q", a.RootVectorFormula, a.KoideFormula, a.DemocraticAxis, a.CosineFormula, a.ConeEquivalence, a.TargetQ, a.TargetAngleDeg, a.PositiveConeOnly, a.UsesObservedEndpoint, a.Verdict)
}

func FormatConePoint(p ConePoint) string {
	return fmt.Sprintf("scale=%s sector=%s labels=%v y=[%s] root=[%s] rho=%.12g parallel=%.12g perp=%.12g perp/parallel=%.12g phiDeg=%.12g Q=%.15g delta=%.15g thetaDeg=%.12g thetaDelta=%.12g on1e4=%t on1e5=%t verdict=%q", p.Scale, p.Sector, p.Labels, floats(p.Yukawas), floats(p.RootVector), p.Rho, p.DemocraticParallel, p.PerpendicularNorm, p.PerpOverParallel, p.AzimuthDeg, p.Q, p.DeltaFromTwoThirds, p.AngleDeg, p.AngleDeltaDeg, p.OnKoideCone1e4, p.OnKoideCone1e5, p.Verdict)
}

func FormatComparison(a SectorComparison) string {
	items := []string{}
	for _, p := range a.Points {
		items = append(items, fmt.Sprintf("%s/%s:Q=%.12g,theta=%.8g,delta=%.4g", p.Scale, p.Sector, p.Q, p.AngleDeg, p.DeltaFromTwoThirds))
	}
	return fmt.Sprintf("points=[%s] leptonMZ=%t leptonLambda=%t upKoide=%t downKoide=%t stable=%t universal=%t best=%q verdict=%q", strings.Join(items, "; "), a.ChargedLeptonMZSharp, a.ChargedLeptonLambda12Sharp, a.UpQuarksOnKoideCone, a.DownQuarksOnKoideCone, a.ChargedLeptonTransportStable, a.KoideUniversalAcrossSectors, a.BestSector, a.Verdict)
}

func FormatSeal(a MinimalEnvironmentalSeal) string {
	return fmt.Sprintf("name=%q carrier=%q constraint=%q coords=%v original=%d constraints=%d remaining=%d native=%t bridge=%t solves=%q verdict=%q", a.Name, a.Carrier, a.SealConstraint, a.ReducedCoordinates, a.OriginalPositiveMagnitudes, a.ConeConstraintCount, a.RemainingContinuousCoordinates, a.NativeDerivation, a.BridgeOnly, a.SolvesFirstLogicalSealAs, a.Verdict)
}

func FormatGate352(a Gate352Inheritance) string {
	return fmt.Sprintf("gate=%d empirical=%t nativePromotion=%t rootTraceNative=%t pfaffian=%t required=%q verdict=%q", a.Gate, a.EmpiricalAlignment, a.NativePromotion, a.RootTraceNative, a.PfaffianCanGenerate, a.RequiredNewObject, a.Verdict)
}

func FormatFirewalls(a FirewallAudit) string {
	return fmt.Sprintf("masses=%t yukawaEigen=%t ckm=%t pmns=%t generation=%t importNative=%t newCarrier=%t gate352=%t verdict=%q", a.DerivesChargedLeptonMasses, a.DerivesYukawaEigenvalues, a.DerivesCKM, a.DerivesPMNS, a.DerivesGenerationHierarchy, a.ImportsObservedAsNative, a.AddsNewCarrier, a.PreservesGate352, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("seal=%q geometry=%q Q_MZ=%.15g delta_MZ=%.15g theta_MZ=%.12g Q_Lambda=%.15g delta_Lambda=%.15g native=%t next=%q verdict=%q", a.FirstLogicalSeal, a.StrongestRuntimeGeometry, a.KoideQeMZ, a.KoideDeltaMZ, a.KoideAngleMZDeg, a.KoideQeLambda12, a.KoideDeltaLambda12, a.NativeASHAFlavorDerivation, a.NextRequiredTheorem, a.Verdict)
}

func floats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = fmt.Sprintf("%.12g", x)
	}
	return strings.Join(parts, ",")
}
