package generation2twistresidualrationalcompressionaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2TwistResidualRationalCompressionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 640 — TwistResidual RationalCompression Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate640 twist residual rational compression audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate639 rho_twist obstruction invariant", Passed: a.Inherited.Verdict == StatusGate639RhoInherited && a.Inherited.RepeatedAcrossRoutes && a.Inherited.ResidualInvariant && a.Inherited.CompactSplitObstruction && !a.Inherited.Gate639ClassifiedAsArtifact && !a.Inherited.Gate639SplitG2Certified && !a.Inherited.Gate639BoundaryStressAssignment && !a.Inherited.Gate639SevenOver72Theorem && !a.Inherited.Gate639ScalarFlavorTransport && a.Inherited.Gate639FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "compute rho_twist^2 and compare with 48/217", Passed: a.Compression.Compressed && a.Compression.Verdict == StatusRhoSquaredEquals48Over217 && math.Abs(a.Compression.ResidualSquared) < rationalTolerance && math.Abs(a.Compression.RhoResidual) < rationalTolerance, Detail: FormatCompression(a.Compression)},
			{Name: "verify rational compression repeats across Gate639 cluster routes", Passed: a.Routes.AllClusterRoutesCompress && len(a.Routes.Routes) >= 3 && a.Routes.MaxSquaredDelta < routeTolerance && a.Routes.Verdict == StatusRouteCompressionRepeated, Detail: routeDetails(a.Routes)},
			{Name: "type 48/217 by K7 Hodge polarity and ambient self-dual complement", Passed: a.Skeleton.NumeratorMatches && a.Skeleton.DenominatorMatches && strings.Contains(a.Skeleton.Verdict, StatusNumerator48TypedCandidate) && strings.Contains(a.Skeleton.Verdict, StatusDenominator217TypedCandidate), Detail: FormatSkeleton(a.Skeleton)},
			{Name: "audit trace/projector contraction candidates without certifying derivation", Passed: strings.Contains(a.Projectors.Verdict, StatusProjectorContractionsAudited) && strings.Contains(a.Projectors.Verdict, StatusNoTraceDerivation) && !a.Projectors.TraceDerivationCertified && a.Projectors.BestCandidateResidual < rationalTolerance, Detail: projectorDetails(a.Projectors)},
			{Name: "classify rational compression as obstruction-only candidate", Passed: a.Classification.CompressionCandidate && !a.Classification.ExactFromFiniteMatrixClaim && !a.Classification.ConsequenceOfHodgeSplitClaim && !a.Classification.ArtifactClaim && a.Classification.ObstructionOnly && a.Classification.Verdict == StatusRhoSquaredEquals48Over217, Detail: FormatClassification(a.Classification)},
			{Name: "preserve trace, split-G2, boundary, 7/72, scalar/flavor, and physical firewalls", Passed: !a.Firewalls.ClaimsExactTraceTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate640Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: rho_twist^2≈"+f64(a.Compression.RhoSquared)+" matches 48/217≈"+f64(a.Compression.CandidateRatio)+" to matrix float tolerance; 48=4^2*3 and 217=7*(35-4) are conditionally typed, but no native trace derivation is certified.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func routeDetails(s RouteCompressionAudit) string {
	parts := []string{FormatRoutes(s)}
	for _, r := range s.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func projectorDetails(s ProjectorContractionAudit) string {
	parts := []string{FormatProjectors(s)}
	for _, c := range s.Candidates {
		parts = append(parts, FormatProjectorCandidate(c))
	}
	return strings.Join(parts, "\n")
}
