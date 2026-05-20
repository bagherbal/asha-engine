package generation2negativesectormultiplicityhitchinmetricsourceaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2NegativeSectorMultiplicityHitchinMetricSourceAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 645 — NegativeSectorMultiplicity HitchinMetric Source Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate645 negative-sector multiplicity Hitchin metric source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate644 projector-plane ratio and firewall", Passed: a.Inherited.Verdict == StatusGate644ProjectorPlaneInherited && a.Inherited.ProjectorPlaneRatioCertified && a.Inherited.MinusThreeCandidate && !a.Inherited.MinusThreeSourceFound && !a.Inherited.NativeTraceIdentityFound && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate644FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "decompose native Omega families over K7+ | K7- and construct admissible twists", Passed: a.Components.AllFamiliesAudited && a.Components.AntisymmetrizedTwistUsed && len(a.Components.Families) == 4 && a.Components.Omega0TotalNormSq > 0 && a.Components.Omega1AltTotalNormSq > 0, Detail: FormatComponentDetails(a.Components)},
			{Name: "compute Hitchin metric block form across repeated routes", Passed: a.HitchinBlocks.AllRoutesBlockCertified && len(a.HitchinBlocks.Routes) == 3 && a.HitchinBlocks.MaxPlusSpread < blockTolerance && a.HitchinBlocks.MaxMinusSpread < blockTolerance && a.HitchinBlocks.MaxOffDiagonalNorm < blockTolerance && a.HitchinBlocks.MaxRatioDrift < blockTolerance, Detail: FormatBlockDetails(a.HitchinBlocks)},
			{Name: "certify minus-three as finite block weight and conditionally type it by negative-sector multiplicity", Passed: a.Multiplicity.PerDirectionWeightCertified && a.Multiplicity.NegativeSectorDim == 3 && a.Multiplicity.ObservedNegativeWeight == -3 && !a.Multiplicity.DerivedBySymbolicTheorem && strings.Contains(a.Multiplicity.Verdict, StatusMinusThreeMultiplicityCandidate), Detail: FormatMultiplicity(a.Multiplicity)},
			{Name: "derive projective angle from Hitchin block trace", Passed: a.Angle.AngleFromBlockTrace && strings.Contains(a.Angle.Verdict, StatusProjectiveAngleFromHitchinBlockTrace), Detail: FormatAngle(a.Angle)},
			{Name: "preserve symbolic, split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsSymbolicMultiplicityTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate645Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate645 searches for the -3 weight inside the cubic Hitchin metric contraction b_Omega=(1/6)(i_x Omega)∧(i_y Omega)∧Omega.  The finite block audit certifies g_twist ∝ P_{K7+}-3P_{K7-} across the repeated routes and conditionally supports -3=-dim(K_7^-), while preserving the absence of a symbolic multiplicity theorem and all physics/boundary firewalls.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
