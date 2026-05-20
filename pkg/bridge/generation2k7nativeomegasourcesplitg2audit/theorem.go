package generation2k7nativeomegasourcesplitg2audit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2K7NativeOmegaSourceSplitG2CompatibilityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 637 — K7 Native Omega Source and Split-G2 Compatibility Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate637 native Omega source audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate636 split bilinear carrier", Passed: a.Inherited.Verdict == StatusGate636Inherited && a.Inherited.K7Dimension == 7 && a.Inherited.BKInertiaPlus == 4 && a.Inherited.BKInertiaMinus == 3 && a.Inherited.BKInertiaZero == 0 && math.Abs(a.Inherited.BKTrace-1) < 1e-10 && math.Abs(a.Inherited.BKDeterminant+1) < 1e-8 && a.Inherited.NativeSplitSignature && a.Inherited.BilinearNotSelector && a.Inherited.NoFockSelectorMap && a.Inherited.NoSplitG2Yet && a.Inherited.NoBoundaryAssignment && a.Inherited.NoSevenOver72Theorem && a.Inherited.Gate636FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit octonionic calibration as strongest native Omega source", Passed: a.Source.Verdict == StatusOctonionicCalibrationSource && a.Source.PGSectorDimension == 14 && a.Source.RawCalibrationRows == 70 && a.Source.RawCalibrationColumns == 14 && a.Source.AssociativeFanoTerms == 7 && a.Source.CoassociativeTerms == 7 && a.Source.K7ToPGCoordinatesComputed && !a.Source.UsesArbitrarySplitG2Normal && !a.Source.UsesExternalThreeForm && !a.Source.HodgePolarityAloneSufficient, Detail: FormatSource(a.Source)},
			{Name: "compute P_G pullback Omega candidates", Passed: a.Candidates.PullbackCandidatesComputed && len(a.Candidates.Candidates) == 4 && a.Candidates.NonZeroStableCandidates >= 3 && a.Candidates.CandidateStabilityCertified, Detail: FormatCandidateSummary(a.Candidates)},
			{Name: "certify tensor antisymmetry and Hitchin metrics", Passed: allCandidatesAntisymmetricAndMetricComputed(a.Candidates), Detail: candidateDetails(a.Candidates)},
			{Name: "detect compact positive Hitchin metric rather than B_K split metric", Passed: a.Compatibility.Verdict == join(StatusOmegaCompactNotSplitBK, StatusNoCompatibleOmegaK) && !a.Compatibility.GomegaProportionalToBK && !a.Compatibility.GomegaSignatureMatchesBK && a.Compatibility.BestRelativeResidualToBK > 0.9 && a.Compatibility.BestOmegaInertia == "(7,0,0)", Detail: FormatCompatibility(a.Compatibility)},
			{Name: "block BK-compatible cross product", Passed: a.CrossProduct.Verdict == StatusNoCrossProductIdentity && !a.CrossProduct.OmegaCompatibleWithBK && !a.CrossProduct.CrossProductDefined && !a.CrossProduct.BKPairingIdentityCertified && !a.CrossProduct.SplitCrossProductIdentity, Detail: FormatCrossProduct(a.CrossProduct)},
			{Name: "block split-G2 stabilizer without compatible Omega_K", Passed: strings.Contains(a.Stabilizer.Verdict, StatusSplitSignatureAloneNoSplitG2) && strings.Contains(a.Stabilizer.Verdict, StatusNoCertifiedSplitG2) && a.Stabilizer.BilinearStabilizerCandidate == "O(4,3)" && a.Stabilizer.ExpectedSplitG2Dimension == 14 && !a.Stabilizer.StabilizerDimensionComputed && !a.Stabilizer.SplitG2Certified, Detail: FormatStabilizer(a.Stabilizer)},
			{Name: "classify native Omega status without promotion", Passed: strings.Contains(a.NativeStatus.Verdict, StatusNoCompatibleOmegaK) && strings.Contains(a.NativeStatus.Verdict, StatusNoCertifiedSplitG2) && a.NativeStatus.NativePullbackTensorExists && !a.NativeStatus.CompatibleOmegaKCertified && !a.NativeStatus.SplitG2CandidateCertified && !a.NativeStatus.BoundaryStressAssignment && !a.NativeStatus.SevenOver72TraceTheorem && !a.NativeStatus.PhysicalSpacetimeMetricClaimed && !a.NativeStatus.FockSelectorClaimed && !a.NativeStatus.ScalarRGMatchingClaimed && !a.NativeStatus.FlavorClaimed && !a.NativeStatus.GaugeUnificationClaimed, Detail: FormatNativeStatus(a.NativeStatus)},
			{Name: "preserve physical, selector, boundary, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsPhysicalSpacetimeMetric && !a.Firewalls.ClaimsFockSelector && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72Theorem && !a.Firewalls.ClaimsScalarRGMatching && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsFlavor && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && !a.Firewalls.ClaimsSplitG2WithoutOmega && a.Firewalls.Verdict == StatusGate637Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Computed result: the P_G-sourced octonionic pullback 3-form is stable, but g_Ω has inertia (7,0,0), not B_K inertia (4,3,0); no compatible native Ω_K or split-G2 carrier is certified.")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func allCandidatesAntisymmetricAndMetricComputed(s OmegaCandidateSummary) bool {
	for _, c := range s.Candidates {
		if !c.FullyAntisymmetric || c.AntisymmetryResidual > 1e-9 || !c.HitchinMetricComputed {
			return false
		}
		if c.NonZero && !c.Stable {
			return false
		}
	}
	return true
}

func candidateDetails(s OmegaCandidateSummary) string {
	parts := make([]string, 0, len(s.Candidates)+1)
	parts = append(parts, FormatCandidateSummary(s))
	for _, c := range s.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}
