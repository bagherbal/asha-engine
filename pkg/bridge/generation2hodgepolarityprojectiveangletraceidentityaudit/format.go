package generation2hodgepolarityprojectiveangletraceidentityaudit

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

func FormatInherited(x Gate641Inheritance) string {
	return fmt.Sprintf("sin=%s cos=%s tan=%s sin2=%s cos2=%s failure=%d alignRoot=%d alignNum=%d den=%d complement=%t projective=%t source13=%t trace641=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t firewall=%t verdict=%q", f64(x.SinTheta), f64(x.CosTheta), f64(x.TanTheta), f64(x.SinSquared), f64(x.CosSquared), x.FailureNumerator, x.AlignmentRoot, x.AlignmentNumerator, x.Denominator, x.ComplementIdentified, x.ProjectiveAngleAudited, x.ThirteenSourcesAudited, x.TraceIdentityCertifiedByGate641, x.SplitG2CertifiedByGate641, x.BoundaryStressByGate641, x.SevenOver72TheoremByGate641, x.ScalarFlavorByGate641, x.PhysicalAngleByGate641, x.Gate641FirewallPreserved, x.Verdict)
}

func FormatRawContraction(x RawFrobeniusContraction) string {
	return fmt.Sprintf("%s inner=%s inner2=%s fail2=%s intInner2=%d intFail2=%d intNorm2=%d matches=%t ratio=%q comment=%q", x.RouteName, f64(x.NormalizedInnerProduct), f64(x.NormalizedInnerProductSquare), f64(x.NormalizedFailureSquare), x.IntegerInnerProductSquare, x.IntegerFailureSquare, x.IntegerProductNormSquare, x.ProjectivePairMatches, x.RatioStatement, x.Comment)
}

func FormatRawContractions(x RawFrobeniusContractionAudit) string {
	return fmt.Sprintf("routes=%d cos2Target=%s sin2Target=%s maxCosDelta=%s maxSinDelta=%s integerRatio=%t nativeTrace=%t verdict=%q", len(x.Contractions), f64(x.CandidateCosSquared), f64(x.CandidateSinSquared), f64(x.MaxCosSquaredDelta), f64(x.MaxSinSquaredDelta), x.IntegerRatioVerified, x.NativeTraceIdentityFound, x.Verdict)
}

func FormatRawContractionsDetails(x RawFrobeniusContractionAudit) string {
	parts := []string{FormatRawContractions(x)}
	for _, c := range x.Contractions {
		parts = append(parts, FormatRawContraction(c))
	}
	return strings.Join(parts, "\n")
}

func FormatSectorBlock(x HodgeSectorBlock) string {
	return fmt.Sprintf("%s carrier=%q dim=%d signedTrace=%d contribution=%d formula=%q typed=%t native=%t reason=%q", x.Name, x.Carrier, x.Dimension, x.SignedTrace, x.Contribution, x.Formula, x.Typed, x.NativeContraction, x.Reason)
}

func FormatSectorBlocks(x HodgeSectorBlockAudit) string {
	return fmt.Sprintf("p=%d q=%d align=%d failure=%d failureText=%q den=%d blocks=%d matches=%t nativeTrace=%t verdict=%q", x.PDim, x.QDim, x.AlignmentAmplitude, x.FailureAmplitudeSquared, x.FailureAmplitudeText, x.Denominator, len(x.Blocks), x.BlockSkeletonMatches, x.NativeTraceIdentity, x.Verdict)
}

func FormatSectorBlocksDetails(x HodgeSectorBlockAudit) string {
	parts := []string{FormatSectorBlocks(x)}
	for _, b := range x.Blocks {
		parts = append(parts, FormatSectorBlock(b))
	}
	return strings.Join(parts, "\n")
}

func FormatProjectivePair(x ProjectivePairAudit) string {
	return fmt.Sprintf("alignment=%d failureSq=%d failureText=%q den=%d residual=%d tan2=%d/%d formula=%q matches=%t nativeTrace=%t verdict=%q", x.AlignmentAmplitude, x.FailureAmplitudeSquared, x.FailureAmplitudeText, x.Denominator, x.PythagoreanIntegerResidual, x.TanSquaredNumerator, x.TanSquaredDenominator, x.Formula, x.PairMatches, x.DerivedFromNativeTraceIdentity, x.Verdict)
}

func FormatTraceCandidate(x TraceIdentityCandidate) string {
	return fmt.Sprintf("%s formula=%q value=%s target=%s residual=%s matches=%t native=%t reason=%q", x.Name, x.Formula, f64(x.Value), f64(x.Target), f64(x.Residual), x.MatchesTarget, x.NativeTraceIdentity, x.Reason)
}

func FormatTraceIdentity(x TraceIdentityAudit) string {
	return fmt.Sprintf("best=%q bestResidual=%s nativeTrace=%t verdict=%q", x.BestCandidateName, f64(x.BestCandidateResidual), x.NativeTraceIdentityFound, x.Verdict)
}

func FormatTraceDetails(x TraceIdentityAudit) string {
	parts := []string{FormatTraceIdentity(x)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatTraceCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatClassification(x ClassificationAudit) string {
	return fmt.Sprintf("angleInherited=%t raw=%t sector=%t skeleton=%t nativeTrace=%t obstructionOnly=%t interpretation=%q verdict=%q", x.ProjectiveAngleInherited, x.RawContractionsComputed, x.HodgeSectorBlocksComputed, x.BlockSkeletonSupported, x.NativeTraceIdentityCertified, x.ObstructionOnly, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("trace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalAngle=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsNativeTraceIdentity, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalAngle, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
