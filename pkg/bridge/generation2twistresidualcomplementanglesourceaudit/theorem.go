package generation2twistresidualcomplementanglesourceaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2TwistResidualComplementAngleSourceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 641 — TwistResidual ComplementAngle Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate641 twist residual complement angle source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate640 rho_twist^2 = 48/217 obstruction compression", Passed: a.Inherited.Verdict == StatusGate640RhoSquaredInherited && a.Inherited.RhoSquaredCompresses && a.Inherited.RouteCompressionRepeated && a.Inherited.DimensionalSkeletonTyped && !a.Inherited.TraceDerivationCertifiedByGate640 && !a.Inherited.SplitG2CertifiedByGate640 && !a.Inherited.BoundaryStressByGate640 && !a.Inherited.SevenOver72TheoremByGate640 && !a.Inherited.ScalarFlavorByGate640 && !a.Inherited.PhysicalMetricByGate640 && a.Inherited.Gate640FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "identify complement 1-rho_twist^2 as 169/217", Passed: a.Complement.ComplementIdentified && a.Complement.FiniteAngleCandidate && math.Abs(a.Complement.ComplementResidual) < complementTolerance && math.Abs(a.Complement.PythagoreanResidual) < complementTolerance && strings.Contains(a.Complement.Verdict, StatusComplement169Identified) && strings.Contains(a.Complement.Verdict, StatusAlignment13SquaredCandidate), Detail: FormatComplement(a.Complement)},
			{Name: "audit projective alignment angle and normalized Frobenius contractions", Passed: a.Projective.AllRoutesAlign && len(a.Projective.Contractions) >= 3 && a.Projective.MaxCosSquaredDelta < routeComplementTolerance && a.Projective.MaxPythagoreanResidual < routeComplementTolerance && strings.Contains(a.Projective.Verdict, StatusProjectiveAlignmentAngleAudited) && strings.Contains(a.Projective.Verdict, StatusRawFrobeniusContractionsAudited), Detail: FormatProjectiveDetails(a.Projective)},
			{Name: "audit typed source candidates for integer 13", Passed: a.Thirteen.StrongestCandidateTyped && a.Thirteen.StrongestCandidateValue == candidateAlignmentRoot && !a.Thirteen.TraceIdentityCertified && strings.Contains(a.Thirteen.Verdict, StatusThirteenSourceCandidatesAudited) && strings.Contains(a.Thirteen.Verdict, StatusNoNativeTraceIdentityFor13), Detail: FormatThirteenDetails(a.Thirteen)},
			{Name: "search projector-trace identities without certifying derivation", Passed: !a.TraceIdentity.NativeTraceIdentityFound && a.TraceIdentity.BestCandidateResidual < complementTolerance && strings.Contains(a.TraceIdentity.Verdict, StatusTraceIdentitySearched) && strings.Contains(a.TraceIdentity.Verdict, StatusNoNativeTraceIdentityFor13), Detail: FormatTraceDetails(a.TraceIdentity)},
			{Name: "classify complement angle as internal obstruction-only candidate", Passed: a.Classification.SinSquared48Over217 && a.Classification.CosSquared169Over217 && a.Classification.FiniteAngleCandidate && !a.Classification.TraceAngleDecomposition && !a.Classification.NormalizationArtifact && a.Classification.ObstructionOnly && a.Classification.Verdict == StatusAlignment13SquaredCandidate, Detail: FormatClassification(a.Classification)},
			{Name: "preserve trace, split-G2, boundary, 7/72, scalar/flavor, and physical firewalls", Passed: !a.Firewalls.ClaimsNativeTraceIdentity && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalAngle && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate641Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: rho_twist^2≈"+f64(a.Complement.RhoSquared)+", 1-rho_twist^2≈"+f64(a.Complement.Complement)+" matches 169/217≈"+f64(a.Complement.CandidateComplement)+"; theta has sin≈"+f64(a.Complement.SinTheta)+", cos≈"+f64(a.Complement.CosTheta)+", tan≈"+f64(a.Complement.TanTheta)+".  The strongest typed 13-source candidate is dim(Im(P_G))-tr(S_K)=14-1, but no native trace identity is certified.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
