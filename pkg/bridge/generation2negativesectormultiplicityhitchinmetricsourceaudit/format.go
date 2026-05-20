package generation2negativesectormultiplicityhitchinmetricsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsInf(x, 0) || math.IsNaN(x) {
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

func FormatInherited(x Gate644Inheritance) string {
	return fmt.Sprintf("ratio=%t G=%q B=%q minusCandidate=%t minusSource=%t nativeTrace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalMetric=%t firewall=%t verdict=%q", x.ProjectorPlaneRatioCertified, x.GHATFormula, x.BHATFormula, x.MinusThreeCandidate, x.MinusThreeSourceFound, x.NativeTraceIdentityFound, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate644FirewallPreserved, x.Verdict)
}

func FormatComponentFamily(x ComponentFamily) string {
	return fmt.Sprintf("%s minus=%d hint=%q omega0Norm2=%s omega1Norm2=%s omega2Norm2=%s survives=%t comment=%q", x.Family, x.MinusCount, x.DimensionHint, f64(x.Omega0NormSq), f64(x.Omega1NormSq), f64(x.Omega2NormSq), x.Survives, x.Comment)
}

func FormatComponents(x OmegaSectorDecomposition) string {
	return fmt.Sprintf("families=%d omega0Total=%s omega1AltTotal=%s omega2AltTotal=%s audited=%t antisymTwist=%t verdict=%q", len(x.Families), f64(x.Omega0TotalNormSq), f64(x.Omega1AltTotalNormSq), f64(x.Omega2AltTotalNormSq), x.AllFamiliesAudited, x.AntisymmetrizedTwistUsed, x.Verdict)
}

func FormatComponentDetails(x OmegaSectorDecomposition) string {
	parts := []string{FormatComponents(x)}
	for _, f := range x.Families {
		parts = append(parts, FormatComponentFamily(f))
	}
	return strings.Join(parts, "\n")
}

func FormatRoute(x HitchinRouteBlock) string {
	return fmt.Sprintf("%s inertia=%s plusMean=%s minusMean=%s ratio=%s plusTrace=%s minusTrace=%s plusSpread=%s minusSpread=%s offNorm=%s block=%t minus3=%t plusEig=%s minusEig=%s formula=%q comment=%q", x.Name, x.Inertia, f64(x.GHatPlusMean), f64(x.GHatMinusMean), f64(x.GHatMinusToPlusRatio), f64(x.GHatPlusTrace), f64(x.GHatMinusTrace), f64(x.PlusBlockSpread), f64(x.MinusBlockSpread), f64(x.PlusMinusFrobNorm), x.BlockFormCertified, x.MinusThreeCertified, f64s(x.PlusBlockEigenvalues), f64s(x.MinusBlockEigenvalues), x.Formula, x.Comment)
}

func FormatBlocks(x HitchinMetricBlockTraceAudit) string {
	return fmt.Sprintf("routes=%d maxPlusSpread=%s maxMinusSpread=%s maxOffNorm=%s maxRatioDrift=%s all=%t weights=(%s,%s) negativeMultiplicity=%d verdict=%q", len(x.Routes), f64(x.MaxPlusSpread), f64(x.MaxMinusSpread), f64(x.MaxOffDiagonalNorm), f64(x.MaxRatioDrift), x.AllRoutesBlockCertified, f64(x.PositiveSectorWeight), f64(x.NegativeSectorWeight), x.NegativeSectorMultiplicity, x.Verdict)
}

func FormatBlockDetails(x HitchinMetricBlockTraceAudit) string {
	parts := []string{FormatBlocks(x)}
	for _, r := range x.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func FormatMultiplicity(x MultiplicitySourceAudit) string {
	return fmt.Sprintf("dims=(plus=%d minus=%d) observedNegativeWeight=%s candidate=%q perDirectionCertified=%t traceMinusOverUnitPlus=%s symbolic=%t verdict=%q explanation=%q", x.PositiveSectorDim, x.NegativeSectorDim, f64(x.ObservedNegativeWeight), x.CandidateFormula, x.PerDirectionWeightCertified, f64(x.TraceTotalMinusOverUnitPlus), x.DerivedBySymbolicTheorem, x.Verdict, x.Explanation)
}

func FormatAngle(x ProjectiveAngleConsequence) string {
	return fmt.Sprintf("B=%q G=%q formula=%q cos=%s residual2=%s fromBlock=%t verdict=%q", x.BHatWeights, x.GHatWeights, x.InnerProductFormula, f64(x.ComputedCosine), f64(x.ComputedResidualSq), x.AngleFromBlockTrace, x.Verdict)
}

func FormatInterpretation(x Interpretation) string {
	return fmt.Sprintf("gate644=%t components=%t block=%t minus3=%t theorem=%t interpretation=%q verdict=%q", x.Gate644Inherited, x.ComponentsComputed, x.HitchinBlockCertified, x.MinusThreeCertified, x.MultiplicityTheoremFound, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsSymbolicMultiplicityTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
