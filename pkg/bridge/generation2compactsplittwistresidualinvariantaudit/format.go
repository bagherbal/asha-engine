package generation2compactsplittwistresidualinvariantaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return fmt.Sprintf("%g", x)
	}
	return fmt.Sprintf("%.12g", x)
}

func FormatInherited(x Gate638Inheritance) string {
	return fmt.Sprintf("gate638=%q gOmegaGK=%t scale=%s gOmegaRes=%s BKscaled=%t BKres=%s fused=%t twist=%t splitG2=%t boundary=%t sevenOver72=%t firewall=%t verdict=%q", x.Gate638Verdict, x.GOmegaAlignedWithGK, f64(x.GOmegaScaleToGK), f64(x.GOmegaRelativeResidualToGK), x.BKAsScaledGOmegaSK, f64(x.BKScaledResidual), x.CompactOmegaAndBKFused, x.NativeSplitCompatibleTwist, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.Gate638FirewallPreserved, x.Verdict)
}

func FormatRoute(x ResidualRoute) string {
	return fmt.Sprintf("%s[%s] formula=%q inertia=%s residual=%s stable=%t split=%t compatible=%t cluster=%t", x.Name, x.SourceGate, x.Formula, x.Inertia, f64(x.RelativeResidualToBK), x.Stable, x.SplitInertia, x.CompatibleWithBK, x.IncludedInRhoCluster)
}

func FormatRepetition(x RepeatedResidualAudit) string {
	return fmt.Sprintf("rho=%s cluster=%s count=%d min=%s max=%s spread=%s relSpread=%s repeated=%t verdict=%q", f64(x.RhoTwist), strings.Join(x.ClusterRouteNames, ","), x.ClusterCount, f64(x.MinClusterResidual), f64(x.MaxClusterResidual), f64(x.Spread), f64(x.RelativeSpread), x.RepeatedAcrossRoutes, x.Verdict)
}

func FormatProbe(x InvarianceProbe) string {
	return fmt.Sprintf("%s baseline=%s transformed=%s drift=%s invariant=%t reason=%q", x.Name, f64(x.BaselineRho), f64(x.TransformedRho), f64(x.AbsoluteDrift), x.Invariant, x.Reason)
}

func FormatInvariance(x ResidualInvarianceAudit) string {
	return fmt.Sprintf("basis=%t omegaScale=%t targetSign=%t SKflip=%t detVol=%t traceFree=%t maxDrift=%s all=%t verdict=%q", x.BasisChangeInvariant, x.OmegaRescaleInvariant, x.TargetSignInvariant, x.SKOrientationInvariant, x.DeterminantVolumeStable, x.TraceFreeStable, f64(x.MaxDrift), x.AllProjectiveTestsPass, x.Verdict)
}

func FormatSourceSweep(x SourceSweepAudit) string {
	return fmt.Sprintf("bestCompact=%s bestCompactRes=%s compactMetricRes=%s bestSplitTwist=%s compactRemovesRho=%t verdict=%q", x.BestCompactSourceName, f64(x.BestCompactSourceResidual), f64(x.CompactPullbackResidual), f64(x.BestSplitTwistResidual), x.CompactSourcesRemoveRho, x.Verdict)
}

func FormatClassification(x ResidualClassificationAudit) string {
	return fmt.Sprintf("rho=%s rho2=%s angleRad=%s angleDeg=%s artifact=%t orbitDistance=%t obstruction=%t interpretation=%q verdict=%q", f64(x.RhoTwist), f64(x.RhoSquared), f64(x.AngleRadiansFromCosModel), f64(x.AngleDegreesFromCosModel), x.ClassifiedAsArtifact, x.ClassifiedAsOrbitDistance, x.ClassifiedAsObstruction, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("physical=%t boundary=%t sevenOver72=%t scalarRG=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t splitG2=%t verdict=%q", x.ClaimsPhysicalSpacetime, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarRG, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.ClaimsSplitG2, x.Verdict)
}
