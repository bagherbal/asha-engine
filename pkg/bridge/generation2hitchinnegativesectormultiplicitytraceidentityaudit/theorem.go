package generation2hitchinnegativesectormultiplicitytraceidentityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HitchinNegativeSectorMultiplicityTraceIdentityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 646 — Hitchin Negative-Sector Multiplicity Trace Identity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate646 Hitchin negative-sector multiplicity trace identity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate645 negative-sector Hitchin metric result and firewalls", Passed: a.Inherited.NegativeWeightCertified && a.Inherited.ProjectiveAngleDerived && a.Inherited.ComponentAuditComputed && a.Inherited.RouteCount == 3 && a.Inherited.MinusThreeSourceCandidate && !a.Inherited.FullSymbolicTheoremCertified && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate645FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit Hodge-sector component families without promoting symbolic family theorem", Passed: a.Components.AllFamiliesAudited && len(a.Components.Families) == 4 && !a.Components.AnyFamilyContributionCertified && !a.Components.SymbolicComponentTheoremCertified, Detail: FormatComponents(a.Components)},
			{Name: "audit off-block cancellation as route-wise finite zero, not symbolic cancellation theorem", Passed: a.OffBlockCancellation.NumericalCancellation && a.OffBlockCancellation.MaxOffBlockFrobeniusNorm < blockTolerance && !a.OffBlockCancellation.SymbolicCancellationCertified, Detail: FormatOffBlock(a.OffBlockCancellation)},
			{Name: "audit positive unit block and negative multiplicity block", Passed: a.PositiveUnit.UnitWeightCertified && !a.PositiveUnit.SymbolicUnitWeightCertified && a.NegativeMultiplicity.MultiplicityWeightCertified && !a.NegativeMultiplicity.SymbolicMultiplicityCertified && a.NegativeMultiplicity.NegativeDim == 3 && a.NegativeMultiplicity.ObservedNegativeWeight == -3, Detail: FormatPositive(a.PositiveUnit) + "\n" + FormatNegative(a.NegativeMultiplicity)},
			{Name: "derive p,q projector-plane angle consequences from certified block identity", Passed: a.ProjectorIdentity.IdentityMatchesRouteData && !a.ProjectorIdentity.FullSymbolicTheoremCertified, Detail: FormatProjectorIdentity(a.ProjectorIdentity)},
			{Name: "audit route universality across omega_1_alt, omega_2_alt, and omega_B_alt", Passed: a.RouteUniversality.AllRoutesPass && a.RouteUniversality.RouteUniversalCandidate && len(a.RouteUniversality.Routes) == 3 && !a.RouteUniversality.RouteDependentFailure, Detail: FormatRouteUniversality(a.RouteUniversality)},
			{Name: "preserve symbolic, split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicHitchinTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate646Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate646 derives the route-supported p,q projector-plane identity G_hat=(P_+-qP_-)/sqrt(p+q^3), B_hat=(P_+-P_-)/sqrt(p+q).  For p=4 and q=3 this gives cos(theta)=13/sqrt(217) and rho^2=48/217.  The finite Hitchin block identity is route-universal, but no full symbolic cubic-contraction multiplicity theorem is certified.")
		if !strings.Contains(a.Interpretation.Verdict, StatusNoFullSymbolicHitchinTheorem) {
			notes = append(notes, "WARNING_UNEXPECTED_SYMBOLIC_THEOREM_PROMOTION_BLOCKED")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
