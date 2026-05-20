package generation2twistresidualcomplementanglesourceaudit

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

func FormatInherited(x Gate640Inheritance) string {
	return fmt.Sprintf("rho=%s rho2=%s ratio48_217=%s compressed=%t routes=%t skeleton=%t trace640=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", f64(x.RhoTwist), f64(x.RhoSquared), f64(x.RhoSquaredRatio48Over217), x.RhoSquaredCompresses, x.RouteCompressionRepeated, x.DimensionalSkeletonTyped, x.TraceDerivationCertifiedByGate640, x.SplitG2CertifiedByGate640, x.BoundaryStressByGate640, x.SevenOver72TheoremByGate640, x.ScalarFlavorByGate640, x.PhysicalMetricByGate640, x.Gate640FirewallPreserved, x.Verdict)
}

func FormatComplement(x ComplementAngleAudit) string {
	return fmt.Sprintf("rho2=%s complement=%s candidate=%d/%d=%s delta=%s rel=%s sin=%s cos=%s tan=%s thetaRad=%s thetaDeg=%s pyth=%s identified=%t angleCandidate=%t verdict=%q", f64(x.RhoSquared), f64(x.Complement), x.AlignmentNumerator, x.Denominator, f64(x.CandidateComplement), f64(x.ComplementResidual), f64(x.RelativeResidual), f64(x.SinTheta), f64(x.CosTheta), f64(x.TanTheta), f64(x.ThetaRadians), f64(x.ThetaDegrees), f64(x.PythagoreanResidual), x.ComplementIdentified, x.FiniteAngleCandidate, x.Verdict)
}

func FormatContraction(x NormalizedFrobeniusContraction) string {
	return fmt.Sprintf("%s sin2=%s cos2=%s sine=%s cosine=%s tan=%s <g,b>norm=%s ||g||=%s ||b||=%s pyth=%s matches13=%t comment=%q", x.RouteName, f64(x.SinSquared), f64(x.CosSquared), f64(x.Sine), f64(x.Cosine), f64(x.Tangent), f64(x.InnerProductNormalized), f64(x.NormGTwist), f64(x.NormBK), f64(x.PythagoreanResidual), x.Matches13Squared, x.Comment)
}

func FormatProjective(x ProjectiveAlignmentAudit) string {
	return fmt.Sprintf("routes=%d bestCos2=%s candidateCos2=%s maxDelta=%s maxPyth=%s all=%t verdict=%q", len(x.Contractions), f64(x.BestCosSquared), f64(x.CandidateCosSquared), f64(x.MaxCosSquaredDelta), f64(x.MaxPythagoreanResidual), x.AllRoutesAlign, x.Verdict)
}

func FormatThirteenCandidate(x ThirteenSourceCandidate) string {
	return fmt.Sprintf("%s formula=%q value=%d matches13=%t strength=%q certified=%t reason=%q", x.Name, x.Formula, x.Value, x.Matches13, x.Strength, x.Certified, x.Reason)
}

func FormatThirteen(x ThirteenSourceAudit) string {
	return fmt.Sprintf("strongest=%q value=%d typed=%t traceIdentity=%t verdict=%q", x.StrongestCandidateName, x.StrongestCandidateValue, x.StrongestCandidateTyped, x.TraceIdentityCertified, x.Verdict)
}

func FormatTraceCandidate(x TraceIdentityCandidate) string {
	return fmt.Sprintf("%s formula=%q value=%s matches=%t nativeIdentity=%t reason=%q", x.Name, x.Formula, f64(x.Value), x.Matches169Over217, x.NativeIdentity, x.Reason)
}

func FormatTraceIdentity(x ProjectorTraceIdentityAudit) string {
	return fmt.Sprintf("best=%q bestResidual=%s nativeTraceIdentity=%t verdict=%q", x.BestCandidateName, f64(x.BestCandidateResidual), x.NativeTraceIdentityFound, x.Verdict)
}

func FormatClassification(x ClassificationAudit) string {
	return fmt.Sprintf("sin48=%t cos169=%t finiteAngle=%t traceAngle=%t artifact=%t obstructionOnly=%t interpretation=%q verdict=%q", x.SinSquared48Over217, x.CosSquared169Over217, x.FiniteAngleCandidate, x.TraceAngleDecomposition, x.NormalizationArtifact, x.ObstructionOnly, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("trace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsNativeTraceIdentity, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalAngle, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}

func FormatProjectiveDetails(x ProjectiveAlignmentAudit) string {
	parts := []string{FormatProjective(x)}
	for _, c := range x.Contractions {
		parts = append(parts, FormatContraction(c))
	}
	return strings.Join(parts, "\n")
}

func FormatThirteenDetails(x ThirteenSourceAudit) string {
	parts := []string{FormatThirteen(x)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatThirteenCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatTraceDetails(x ProjectorTraceIdentityAudit) string {
	parts := []string{FormatTraceIdentity(x)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatTraceCandidate(c))
	}
	return strings.Join(parts, "\n")
}
