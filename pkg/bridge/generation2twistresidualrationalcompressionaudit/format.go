package generation2twistresidualrationalcompressionaudit

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

func FormatInherited(x Gate639Inheritance) string {
	return fmt.Sprintf("rho=%s rho2=%s repeated=%t invariant=%t obstruction=%t artifact=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t firewall=%t verdict=%q", f64(x.RhoTwist), f64(x.RhoSquared), x.RepeatedAcrossRoutes, x.ResidualInvariant, x.CompactSplitObstruction, x.Gate639ClassifiedAsArtifact, x.Gate639SplitG2Certified, x.Gate639BoundaryStressAssignment, x.Gate639SevenOver72Theorem, x.Gate639ScalarFlavorTransport, x.Gate639FirewallPreserved, x.Verdict)
}

func FormatCompression(x RationalCompressionAudit) string {
	return fmt.Sprintf("rho=%s rho2=%s candidate=%d/%d=%s sqrt=%s delta2=%s rel=%s deltaRho=%s compressed=%t verdict=%q", f64(x.RhoTwist), f64(x.RhoSquared), x.CandidateNumerator, x.CandidateDenominator, f64(x.CandidateRatio), f64(x.CandidateSqrt), f64(x.ResidualSquared), f64(x.RelativeResidual), f64(x.RhoResidual), x.Compressed, x.Verdict)
}

func FormatRoute(x RouteCompression) string {
	return fmt.Sprintf("%s rho=%s rho2=%s delta2=%s deltaRho=%s compressed=%t", x.Name, f64(x.Residual), f64(x.ResidualSquared), f64(x.DeltaTo48Over217), f64(x.RhoDelta), x.Compressed)
}

func FormatRoutes(x RouteCompressionAudit) string {
	return fmt.Sprintf("routes=%s maxDelta2=%s maxDeltaRho=%s all=%t verdict=%q", strings.Join(x.CompressedRouteNames, ","), f64(x.MaxSquaredDelta), f64(x.MaxRhoDelta), x.AllClusterRoutesCompress, x.Verdict)
}

func FormatSkeleton(x DimensionalSkeletonAudit) string {
	return fmt.Sprintf("K7=%d K7+=(%d) K7-=(%d) Lambda4+=(%d) Lambda4-=(%d) complement31=%d numerator=%d denominator=%d matches=(%t,%t) formula=%q verdict=%q", x.K7Dim, x.K7PlusDim, x.K7MinusDim, x.Lambda4SelfDualDim, x.Lambda4AntiSelfDualDim, x.SelfDualComplementToK7PlusDim, x.NumeratorFromHodgePolarity, x.DenominatorFromSelfDualGap, x.NumeratorMatches, x.DenominatorMatches, x.Formula, x.Verdict)
}

func FormatProjectorCandidate(x ProjectorContractionCandidate) string {
	return fmt.Sprintf("%s formula=%q value=%s matches=%t nativeDerivation=%t reason=%q", x.Name, x.Formula, f64(x.Value), x.Matches48Over217, x.NativeDerivation, x.Reason)
}

func FormatProjectors(x ProjectorContractionAudit) string {
	return fmt.Sprintf("P+=(%d) P-=(%d) PK7+=(%d) PK7-=(%d) best=%q bestResidual=%s traceDerivation=%t verdict=%q", x.PPlusDimension, x.PMinusDimension, x.PK7PlusDimension, x.PK7MinusDimension, x.BestCandidateName, f64(x.BestCandidateResidual), x.TraceDerivationCertified, x.Verdict)
}

func FormatClassification(x ClassificationAudit) string {
	return fmt.Sprintf("rho2=%s ratio=%s candidate=%t exactClaim=%t hodgeSplitClaim=%t artifact=%t obstructionOnly=%t interpretation=%q verdict=%q", f64(x.RhoSquared), f64(x.Ratio48Over217), x.CompressionCandidate, x.ExactFromFiniteMatrixClaim, x.ConsequenceOfHodgeSplitClaim, x.ArtifactClaim, x.ObstructionOnly, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("trace=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsExactTraceTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
