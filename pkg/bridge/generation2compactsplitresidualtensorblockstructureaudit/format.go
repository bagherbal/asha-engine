package generation2compactsplitresidualtensorblockstructureaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return fmt.Sprintf("%g", x)
	}
	return fmt.Sprintf("%.15g", x)
}

func f64s(xs []float64) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, f64(x))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatInherited(x Gate642Inheritance) string {
	return fmt.Sprintf("cos=%s sin=%s cos2=%s sin2=%s align=%d failure=%d den=%d skeleton=%t nativeTrace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t firewall=%t verdict=%q", f64(x.CosTheta), f64(x.SinTheta), f64(x.CosSquared), f64(x.SinSquared), x.AlignmentNumerator, x.FailureNumerator, x.Denominator, x.HodgePolaritySkeleton, x.NativeTraceIdentityCertified, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalAngle, x.PhysicalMetric, x.Gate642FirewallPreserved, x.Verdict)
}

func FormatRoute(x ResidualTensorRoute) string {
	return fmt.Sprintf("%s inertia=%s cos=%s rho=%s orthB=%s unit=%s sym=%s blocks(pp=%s mm=%s 2pm=%s sum=%s) ranks(pp=%d mm=%d pm=%d) traces(pp=%s mm=%s) dom=%q offDom=%t typed=%t ppEig=%s mmEig=%s pmSing=%s formula=%q comment=%q", x.Name, x.Inertia, f64(x.Cosine), f64(x.Rho), f64(x.OrthogonalityToBHat), f64(x.ResidualUnitNorm), f64(x.SymmetryResidual), f64(x.RPlusPlusFrobSquared), f64(x.RMinusMinusFrobSquared), f64(x.TwiceRPlusMinusFrobSq), f64(x.BlockNormSum), x.PlusPlusRank, x.MinusMinusRank, x.PlusMinusRank, f64(x.PlusPlusTrace), f64(x.MinusMinusTrace), x.DominantBlock, x.OffSectorDominant, x.TypedBlockProfile, f64s(x.PlusPlusEigenvalues), f64s(x.MinusMinusEigenvalues), f64s(x.PlusMinusSingularValues), x.Formula, x.Comment)
}

func FormatResidualTensor(x ResidualTensorAudit) string {
	return fmt.Sprintf("routes=%d cosTarget=%s rhoTarget=%s maxOrth=%s maxUnitDrift=%s maxCosDrift=%s maxRhoDrift=%s certified=%t verdict=%q", len(x.Routes), f64(x.CosineTarget), f64(x.RhoTarget), f64(x.MaxOrthogonalityToBHat), f64(x.MaxResidualUnitNormDrift), f64(x.MaxCosineDrift), f64(x.MaxRhoDrift), x.ResidualTensorsCertified, x.Verdict)
}

func FormatResidualTensorDetails(x ResidualTensorAudit) string {
	parts := []string{FormatResidualTensor(x)}
	for _, r := range x.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func FormatBlockSummary(x HodgeBlockSummary) string {
	return fmt.Sprintf("routes=%d meanPP=%s meanMM=%s mean2PM=%s offDominantRoutes=%d sameDominant=%t sameRank=%t typed=%t nativeTrace=%t skeleton=%q exactProfile=%q verdict=%q", x.RouteCount, f64(x.PlusPlusMeanFrobSquared), f64(x.MinusMinusMeanFrobSquared), f64(x.TwicePlusMinusMeanFrobSquared), x.OffSectorDominantRoutes, x.SameDominantBlockAllRoutes, x.SameRankProfileAllRoutes, x.HasTypedBlockStructure, x.NativeTraceIdentityFound, x.FailureSkeleton, x.ExactSameSectorProfile, x.Verdict)
}

func FormatInterpretation(x ResidualInterpretation) string {
	return fmt.Sprintf("angle=%t tensor=%t blocks=%t typed=%t offCandidate=%t nativeTrace=%t interpretation=%q verdict=%q", x.AnglePairInherited, x.ResidualTensorDefined, x.BlocksComputed, x.TypedBlockStructure, x.OffSectorCarrierCandidate, x.NativeTraceIdentityFound, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("trace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsNativeTraceIdentity, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalAngle, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
