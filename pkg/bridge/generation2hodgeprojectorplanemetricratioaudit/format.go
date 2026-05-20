package generation2hodgeprojectorplanemetricratioaudit

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

func FormatInherited(x Gate643Inheritance) string {
	return fmt.Sprintf("cos=%s sin=%s residual=%t sameSector=%t offRejected=%t blocks(pp=%s mm=%s 2pm=%s) nativeTrace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t firewall=%t verdict=%q", f64(x.CosTheta), f64(x.SinTheta), x.ResidualTensorCertified, x.SameSectorHodgeDiagonal, x.OffSectorCarrierRejected, f64(x.RPlusPlusFrobSquared), f64(x.RMinusMinusFrobSquared), f64(x.TwiceRPlusMinusFrobSq), x.NativeTraceIdentityFound, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalAngle, x.PhysicalMetric, x.Gate643FirewallPreserved, x.Verdict)
}

func FormatDefinitions(x ProjectorPlaneDefinition) string {
	return fmt.Sprintf("B=%q G=%q Bweights=(%s,%s) Gweights=(%s,%s) norms(B=%s G=%s) residuals(B=%s Gnorm=%s) certified=%t verdict=%q", x.BHatFormula, x.GHatFormula, f64(x.BHatPlusWeight), f64(x.BHatMinusWeight), f64(x.GHatPlusWeight), f64(x.GHatMinusWeight), f64(x.BHatNorm), f64(x.GHatNorm), f64(x.BHatTargetResidual), f64(x.ProjectorPlaneTargetResidual), x.ProjectorPlaneMetricsCertified, x.Verdict)
}

func FormatRoute(x MetricRatioRoute) string {
	return fmt.Sprintf("%s inertia=%s cos=%s rho=%s planeResidual=%s reconstructResidual=%s plusMean=%s minusMean=%s ratio=%s plusSpread=%s minusSpread=%s offNorm=%s certified=%t plusEig=%s minusEig=%s formula=%q comment=%q", x.Name, x.Inertia, f64(x.Cosine), f64(x.Rho), f64(x.GHatToProjectorPlaneResidual), f64(x.GHatToReconstructedResidual), f64(x.PlusBlockMean), f64(x.MinusBlockMean), f64(x.ObservedMinusToPlusRatio), f64(x.PlusBlockSpread), f64(x.MinusBlockSpread), f64(x.PlusMinusFrobNorm), x.Ratio1ToMinus3Certified, f64s(x.PlusBlockEigenvalues), f64s(x.MinusBlockEigenvalues), x.Formula, x.Comment)
}

func FormatMetricRatio(x MetricRatioAudit) string {
	return fmt.Sprintf("routes=%d maxPlaneResidual=%s maxReconstructResidual=%s maxPlusSpread=%s maxMinusSpread=%s maxOffNorm=%s maxRatioDrift=%s certified=%t verdict=%q", len(x.Routes), f64(x.MaxProjectorPlaneResidual), f64(x.MaxReconstructedResidual), f64(x.MaxPlusSpread), f64(x.MaxMinusSpread), f64(x.MaxOffDiagonalNorm), f64(x.MaxRatioDrift), x.AllRoutesRatioCertified, x.Verdict)
}

func FormatMetricRatioDetails(x MetricRatioAudit) string {
	parts := []string{FormatMetricRatio(x)}
	for _, r := range x.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func FormatAngleFromPlane(x ProjectiveAngleFromPlane) string {
	return fmt.Sprintf("dims=(%d,%d) B=%q G=%q formula=%q cos=%s expectedCos=%s sin2=%s expectedSin2=%s derived=%t nativeTrace=%t verdict=%q", x.PlusDim, x.MinusDim, x.BHatWeights, x.GHatWeights, x.InnerProductFormula, f64(x.ComputedCosine), f64(x.ExpectedCosine), f64(x.ComputedSinSquared), f64(x.ExpectedSinSquared), x.AngleDerivedFromPlane, x.NativeTraceIdentityFound, x.Verdict)
}

func FormatMinusThree(x MinusThreeSourceAudit) string {
	return fmt.Sprintf("negativeDim=%d dimCandidate=%q traceCandidate=%q twistCandidate=%q certified=%t verdict=%q", x.NegativeHodgeSectorDim, x.CandidateFromDimK7Minus, x.CandidateFromTraceBalance, x.CandidateFromTwistOperation, x.CertifiedNativeSource, x.Verdict)
}

func FormatInterpretation(x Interpretation) string {
	return fmt.Sprintf("inherited=%t reconstructed=%t ratio=%t angle=%t minus3=%t nativeTrace=%t interpretation=%q verdict=%q", x.Gate643Inherited, x.GHatReconstructed, x.RatioCertified, x.AngleFromPlane, x.MinusThreeSourceFound, x.NativeTraceIdentityFound, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("trace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsNativeTraceIdentity, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalAngle, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
