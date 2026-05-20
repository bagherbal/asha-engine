package generation2hitchincubicsectorcontractionmultiplicityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HitchinCubicSectorContractionMultiplicityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 647 — Hitchin Cubic Sector-Contraction Multiplicity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate647 Hitchin cubic sector-contraction multiplicity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate646 p,q projector-plane identity and firewalls", Passed: a.Inherited.ProjectorPlaneIdentityInherited && a.Inherited.RouteUniversal && a.Inherited.PositiveDim == 4 && a.Inherited.NegativeDim == 3 && !a.Inherited.FullSymbolicTheoremCertified && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate646FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "compute component-family tensor ledger for admissible antisymmetrized twists", Passed: a.Families.AllFamiliesAudited && len(a.Families.Families) == 4 && a.Families.AntisymmetrizedTwistsAudited && !a.Families.AnySymbolicFamilyTheorem, Detail: FormatFamilies(a.Families)},
			{Name: "expand cubic Hitchin contraction into ordered family-triple contribution ledgers", Passed: a.Contributions.AllRoutesReconstruct && a.Contributions.AllRoutesBlockRayCertified && a.Contributions.SameProjectorPlaneShadow && len(a.Contributions.Routes) == 3 && !a.Contributions.AnyRouteSymbolicCertified, Detail: FormatContributions(a.Contributions)},
			{Name: "audit positive unit coefficient and negative -q multiplicity source", Passed: a.PositiveUnit.AllRoutesUnitPositive && !a.PositiveUnit.SymbolicUnitTheoremCertified && a.NegativeMultiplicity.AllRoutesMinusQ && a.NegativeMultiplicity.CubicSectorMultiplicitySupported && !a.NegativeMultiplicity.SymbolicMultiplicityTheoremCertified, Detail: FormatPositive(a.PositiveUnit) + "\n" + FormatNegative(a.NegativeMultiplicity)},
			{Name: "audit mixed-block cancellation source without symbolic promotion", Passed: a.OffBlockCancellation.MaxTotalOffBlockNorm < blockTolerance && !a.OffBlockCancellation.StructuralCancellationCertified, Detail: FormatOffBlock(a.OffBlockCancellation)},
			{Name: "compare route universality at final ray and component-ledger levels", Passed: a.RouteUniversality.RouteCount == 3 && a.RouteUniversality.AllRoutesSameFinalRay && !a.RouteUniversality.ComponentContributionLedgersEqual, Detail: FormatRouteUniversality(a.RouteUniversality)},
			{Name: "sharpen candidate theorem while preserving no-symbolic-theorem firewall", Passed: a.TheoremReadiness.FiniteLedgerSupportsTheorem && a.TheoremReadiness.ComponentLedgerComputed && a.TheoremReadiness.BlockContributionComputed && !a.TheoremReadiness.FullSymbolicTheoremCertified, Detail: FormatTheoremReadiness(a.TheoremReadiness)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicHitchinTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate647Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate647 computes the ordered family-triple ledger of the cubic Hitchin contraction.  The finite ledger reconstructs the route-universal ray g_twist ∝ P_+ - 3P_- and sharpens the candidate theorem HitchinMetric(Ω_twist) ∝ P_+ - dim(K_7^-)P_-, but no basis-free symbolic contraction proof is certified.")
		if !strings.Contains(a.TheoremReadiness.Verdict, StatusNoFullSymbolicHitchinTheorem) {
			notes = append(notes, "WARNING_UNEXPECTED_SYMBOLIC_THEOREM_PROMOTION_BLOCKED")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
