package generation2compactomegahodgesplitpolarizationtwistaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CompactOmegaHodgeSplitPolarizationTwistAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 638 — Compact Omega / Hodge Split Polarization and Twist-Admissibility Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate638 compact Omega / Hodge split twist audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate637 compact Omega versus B_K conflict", Passed: a.Inherited.Verdict == StatusGate637Inherited && a.Inherited.K7Dimension == 7 && a.Inherited.BKInertia == "(4,3,0)" && a.Inherited.BestOmegaInertia == "(7,0,0)" && a.Inherited.NativePullbackTensorExists && !a.Inherited.CompatibleOmegaKCertified && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && a.Inherited.CompactOmegaAndBKConflict && a.Inherited.Gate637FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit g_Omega alignment with inherited g_K", Passed: strings.Contains(a.MetricAlign.Verdict, StatusGOmegaToGKAlignment) && strings.Contains(a.MetricAlign.Verdict, StatusGOmegaAlignedCompactGK) && a.MetricAlign.GOmegaInertiaPlus == 7 && a.MetricAlign.GOmegaInertiaMinus == 0 && a.MetricAlign.GOmegaInertiaZero == 0 && a.MetricAlign.RelativeResidualToGK < 1e-8 && a.MetricAlign.AlignedWithGK && a.MetricAlign.CompactPositive, Detail: FormatMetricAlignment(a.MetricAlign)},
			{Name: "certify B_K as Hodge-polarized compact metric", Passed: a.Reconstruction.Verdict == StatusBKEqualsGKSK && a.Reconstruction.BKEqualsGKSK && a.Reconstruction.BKResidual < 1e-12 && a.Reconstruction.BKEqualsScaledGOmegaSK && a.Reconstruction.ScaledGOmegaSKResidual < 1e-8, Detail: FormatReconstruction(a.Reconstruction)},
			{Name: "audit S_K action on Omega_0", Passed: a.SKAction.Verdict == StatusSKActionOnOmegaAudited && a.SKAction.SKOrthogonalForGOmega && a.SKAction.OrthogonalityResidual < 1e-8 && (a.SKAction.Omega3Inertia == "(7,0,0)" || a.SKAction.Omega3Inertia == "(0,7,0)"), Detail: FormatSKAction(a.SKAction)},
			{Name: "audit admissible S_K twists with antisymmetrization", Passed: strings.Contains(a.Twists.Verdict, StatusTwistAdmissibilityAudited) && strings.Contains(a.Twists.Verdict, StatusNoSKTwistMatchesBK) && len(a.Twists.Candidates) == 4 && a.Twists.AdmissibleAlternatingCandidates == 4 && a.Twists.StableCandidates >= 1 && a.Twists.SplitCompatibleCandidates == 0 && !a.Twists.NativeSKTwistMatchesBK, Detail: twistDetails(a.Twists)},
			{Name: "audit compact cross-product paired with B_K", Passed: strings.Contains(a.CrossProduct.Verdict, StatusCrossProductAudited) && strings.Contains(a.CrossProduct.Verdict, StatusNoSKTwistMatchesBK) && a.CrossProduct.CompactCrossProductDefined && !a.CrossProduct.OmegaBMatchesBK, Detail: FormatCrossProduct(a.CrossProduct)},
			{Name: "classify compact Omega and Hodge split B_K as unfused", Passed: strings.Contains(a.Interpretation.Verdict, StatusCompactOmegaBKDoNotFuse) && strings.Contains(a.Interpretation.Verdict, StatusNoCertifiedSplitG2) && a.Interpretation.GOmegaAlignedWithGK && a.Interpretation.BKIsHodgePolarizedCompactMetric && !a.Interpretation.NativeSplitCompatibleTwistFound && !a.Interpretation.CompactOmegaAndBKFused, Detail: FormatInterpretation(a.Interpretation)},
			{Name: "preserve physical, boundary, 7/72, scalar, flavor, and gauge firewalls", Passed: !a.Firewalls.ClaimsPhysicalSpacetime && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsScalarRG && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsSplitG2 && a.Firewalls.Verdict == StatusGate638Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: the native compact Omega_0 aligns with g_K, and B_K is the S_K-polarization of that compact metric, but no admissible S_K twist or B_K-paired compact cross-product tensor induces B_K; the compact calibration and split bilinear remain unfused.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func twistDetails(s TwistAdmissibilityAudit) string {
	parts := []string{FormatTwists(s)}
	for _, c := range s.Candidates {
		parts = append(parts, FormatTwistCandidate(c))
	}
	return strings.Join(parts, "\n")
}
