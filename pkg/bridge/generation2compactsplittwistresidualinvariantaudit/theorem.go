package generation2compactsplittwistresidualinvariantaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CompactSplitTwistResidualInvariantAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 639 — CompactSplitTwistResidual Invariant Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate639 compact/split residual invariant audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate638 unfused compact Omega and Hodge split polarization", Passed: a.Inherited.Verdict == StatusGate638UnfusedInherited && a.Inherited.GOmegaAlignedWithGK && a.Inherited.BKAsScaledGOmegaSK && !a.Inherited.CompactOmegaAndBKFused && !a.Inherited.NativeSplitCompatibleTwist && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && a.Inherited.Gate638FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "certify rho_twist repetition across independent split routes", Passed: a.Repetition.Verdict == StatusTwistResidualRepeated && a.Repetition.RepeatedAcrossRoutes && a.Repetition.ClusterCount >= 3 && a.Repetition.Spread < repeatedTolerance && a.Repetition.RhoTwist > 0.47 && a.Repetition.RhoTwist < 0.471, Detail: repetitionDetails(a.Repetition)},
			{Name: "compute projective residual invariance tests", Passed: strings.Contains(a.Invariance.Verdict, StatusResidualInvarianceTests) && strings.Contains(a.Invariance.Verdict, StatusResidualNotNormalization) && strings.Contains(a.Invariance.Verdict, StatusProjectiveResidualAudited) && a.Invariance.AllProjectiveTestsPass && a.Invariance.MaxDrift < invarianceTolerance, Detail: invarianceDetails(a.Invariance)},
			{Name: "audit compact source sweep without removing rho_twist", Passed: strings.Contains(a.SourceSweep.Verdict, StatusResidualInvarianceTests) && strings.Contains(a.SourceSweep.Verdict, StatusResidualNotNormalization) && !a.SourceSweep.CompactSourcesRemoveRho && a.SourceSweep.BestCompactSourceResidual > a.SourceSweep.BestSplitTwistResidual, Detail: sourceSweepDetails(a.SourceSweep)},
			{Name: "classify rho_twist as compact/split obstruction witness", Passed: a.Classification.Verdict == StatusCompactSplitObstruction && !a.Classification.ClassifiedAsArtifact && a.Classification.ClassifiedAsOrbitDistance && a.Classification.ClassifiedAsObstruction && a.Classification.RhoTwist > 0.47 && a.Classification.RhoTwist < 0.471, Detail: FormatClassification(a.Classification)},
			{Name: "preserve split-G2, boundary, 7/72, scalar, flavor, and physical firewalls", Passed: !a.Firewalls.ClaimsPhysicalSpacetime && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarRG && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsSplitG2 && a.Firewalls.Verdict == StatusGate639Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: rho_twist≈"+f64(a.Repetition.RhoTwist)+" repeats across omega_1_alt, omega_2_alt, and the antisymmetrized B_K-paired cross-product route; it survives projective normalization tests and remains an internal compact/split obstruction witness, not a physical theorem.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func repetitionDetails(s RepeatedResidualAudit) string {
	parts := []string{FormatRepetition(s)}
	for _, r := range s.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func invarianceDetails(s ResidualInvarianceAudit) string {
	parts := []string{FormatInvariance(s)}
	for _, p := range s.Probes {
		parts = append(parts, FormatProbe(p))
	}
	return strings.Join(parts, "\n")
}

func sourceSweepDetails(s SourceSweepAudit) string {
	parts := []string{FormatSourceSweep(s)}
	for _, r := range s.CandidateResiduals {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}
