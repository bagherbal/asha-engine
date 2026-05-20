package generation2k7nativeomegasourcesplitg2audit

import (
	"fmt"
	"math"
)

func f64(x float64) string {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return fmt.Sprintf("%g", x)
	}
	return fmt.Sprintf("%.12g", x)
}

func FormatInherited(x Gate636Inheritance) string {
	return fmt.Sprintf("K7=%d inertia=(%d,%d,%d) tr=%s det=%s native_split=%t bilinear_not_selector=%t no_fock=%t no_splitG2=%t no_boundary=%t no_7over72=%t firewall=%t verdict=%q", x.K7Dimension, x.BKInertiaPlus, x.BKInertiaMinus, x.BKInertiaZero, f64(x.BKTrace), f64(x.BKDeterminant), x.NativeSplitSignature, x.BilinearNotSelector, x.NoFockSelectorMap, x.NoSplitG2Yet, x.NoBoundaryAssignment, x.NoSevenOver72Theorem, x.Gate636FirewallPreserved, x.Verdict)
}

func FormatSource(x OmegaSourceAudit) string {
	return fmt.Sprintf("source=%q PGdim=%d raw=%dx%d fano_terms=%d coassoc_terms=%d coords=%t arbitrary_splitG2=%t external=%t hodge_alone=%t verdict=%q", x.StrongestCandidateSource, x.PGSectorDimension, x.RawCalibrationRows, x.RawCalibrationColumns, x.AssociativeFanoTerms, x.CoassociativeTerms, x.K7ToPGCoordinatesComputed, x.UsesArbitrarySplitG2Normal, x.UsesExternalThreeForm, x.HodgePolarityAloneSufficient, x.Verdict)
}

func FormatCandidate(x CandidateTensorAudit) string {
	return fmt.Sprintf("name=%s formula=%q source=%q antisym=%t antiRes=%s norm=%s max=%s hitchin=%t bNorm=%s det=%s inertia=(%d,%d,%d) stable=%t scaleToBK=%s relResBK=%s compatible=%t verdict=%q", x.Name, x.Formula, x.Source, x.FullyAntisymmetric, f64(x.AntisymmetryResidual), f64(x.TensorNorm), f64(x.MaxAbsComponent), x.HitchinMetricComputed, f64(x.HitchinMetricFrobenius), f64(x.HitchinMetricDeterminant), x.HitchinMetricInertiaPlus, x.HitchinMetricInertiaMinus, x.HitchinMetricInertiaZero, x.Stable, f64(x.ScaleToBK), f64(x.RelativeResidualToBK), x.CompatibleWithBK, x.Verdict)
}

func FormatCandidateSummary(x OmegaCandidateSummary) string {
	return fmt.Sprintf("count=%d stable=%d compatible=%d best=%s bestRes=%s bestInertia=%s computed=%t stable_cert=%t compatible_cert=%t verdict=%q", len(x.Candidates), x.NonZeroStableCandidates, x.CompatibleWithBKCount, x.BestCandidateName, f64(x.BestRelativeResidualToBK), x.BestHitchinInertia, x.PullbackCandidatesComputed, x.CandidateStabilityCertified, x.CompatibleNativeOmegaCertified, x.Verdict)
}

func FormatCompatibility(x MetricCompatibilityAudit) string {
	return fmt.Sprintf("BK=%s best=%s omegaInertia=%s scale=%s relRes=%s proportional=%t signature_match=%t certified_scale=%t reason=%q verdict=%q", x.BKInertia, x.BestOmegaName, x.BestOmegaInertia, f64(x.BestScaleToBK), f64(x.BestRelativeResidualToBK), x.GomegaProportionalToBK, x.GomegaSignatureMatchesBK, x.CertifiedScaleNotFitted, x.Reason, x.Verdict)
}

func FormatCrossProduct(x CrossProductAudit) string {
	return fmt.Sprintf("omega_BK=%t cross=%t pairing=%t split_identity=%t reason=%q verdict=%q", x.OmegaCompatibleWithBK, x.CrossProductDefined, x.BKPairingIdentityCertified, x.SplitCrossProductIdentity, x.Reason, x.Verdict)
}

func FormatStabilizer(x StabilizerAudit) string {
	return fmt.Sprintf("bilinear=%s omega=%q dim_computed=%t expected_dim=%d splitG2=%t reason=%q verdict=%q", x.BilinearStabilizerCandidate, x.OmegaStabilizerCandidate, x.StabilizerDimensionComputed, x.ExpectedSplitG2Dimension, x.SplitG2Certified, x.Reason, x.Verdict)
}

func FormatNativeStatus(x NativeOmegaStatus) string {
	return fmt.Sprintf("pullback=%t compatibleOmega=%t splitG2=%t boundary=%t sevenOver72=%t physical=%t fock=%t scalarRG=%t flavor=%t gauge=%t statement=%q verdict=%q", x.NativePullbackTensorExists, x.CompatibleOmegaKCertified, x.SplitG2CandidateCertified, x.BoundaryStressAssignment, x.SevenOver72TraceTheorem, x.PhysicalSpacetimeMetricClaimed, x.FockSelectorClaimed, x.ScalarRGMatchingClaimed, x.FlavorClaimed, x.GaugeUnificationClaimed, x.Statement, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("physical=%t fock=%t boundary=%t sevenOver72=%t scalarRG=%t higgs=%t flavor=%t ckm_pmns=%t gauge=%t splitG2_without_omega=%t verdict=%q", x.ClaimsPhysicalSpacetimeMetric, x.ClaimsFockSelector, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsFlavor, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.ClaimsSplitG2WithoutOmega, x.Verdict)
}
