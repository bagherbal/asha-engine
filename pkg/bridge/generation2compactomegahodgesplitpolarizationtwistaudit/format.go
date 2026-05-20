package generation2compactomegahodgesplitpolarizationtwistaudit

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

func FormatInherited(x Gate637Inheritance) string {
	return fmt.Sprintf("K7=%d BK=%s bestOmega=%s omegaInertia=%s pullback=%t compatibleOmega=%t splitG2=%t boundary=%t sevenOver72=%t conflict=%t firewall=%t verdict=%q", x.K7Dimension, x.BKInertia, x.BestOmegaName, x.BestOmegaInertia, x.NativePullbackTensorExists, x.CompatibleOmegaKCertified, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.CompactOmegaAndBKConflict, x.Gate637FirewallPreserved, x.Verdict)
}
func FormatMetricAlignment(x MetricAlignmentAudit) string {
	return fmt.Sprintf("omega=%s inertia=(%d,%d,%d) det=%s scaleToGK=%s relResGK=%s aligned=%t compact=%t verdict=%q", x.OmegaName, x.GOmegaInertiaPlus, x.GOmegaInertiaMinus, x.GOmegaInertiaZero, f64(x.GOmegaDeterminant), f64(x.BestScaleToGK), f64(x.RelativeResidualToGK), x.AlignedWithGK, x.CompactPositive, x.Verdict)
}
func FormatReconstruction(x HodgeBilinearReconstructionAudit) string {
	return fmt.Sprintf("BK=gKSK=%t BKres=%s scale=%s scaledGOmegaSK=%t scaledRes=%s interpretation=%q verdict=%q", x.BKEqualsGKSK, f64(x.BKResidual), f64(x.GOmegaScaleToGK), x.BKEqualsScaledGOmegaSK, f64(x.ScaledGOmegaSKResidual), x.Interpretation, x.Verdict)
}
func FormatSKAction(x SKActionOnOmegaAudit) string {
	return fmt.Sprintf("SK_orthogonal=%t orthRes=%s omega3PlusRes=%s omega3MinusRes=%s sign=%s omega3Inertia=%s compactOrbit=%t verdict=%q", x.SKOrthogonalForGOmega, f64(x.OrthogonalityResidual), f64(x.Omega3RelativeResidualPlus), f64(x.Omega3RelativeResidualMinus), x.Omega3SignClassification, x.Omega3Inertia, x.Omega3RemainsCompactOrbit, x.Verdict)
}
func FormatTwistCandidate(x TwistCandidateAudit) string {
	return fmt.Sprintf("name=%s formula=%q antisymmetrized=%t antiRes=%s norm=%s hitchin=%t inertia=(%d,%d,%d) det=%s stable=%t scaleToBK=%s relResBK=%s splitCompatible=%t verdict=%q", x.Name, x.Formula, x.Antisymmetrized, f64(x.AntisymmetryResidual), f64(x.TensorNorm), x.HitchinMetricComputed, x.InertiaPlus, x.InertiaMinus, x.InertiaZero, f64(x.Determinant), x.Stable, f64(x.ScaleToBK), f64(x.RelativeResidualToBK), x.SplitCompatibleWithBK, x.Verdict)
}
func FormatTwists(x TwistAdmissibilityAudit) string {
	return fmt.Sprintf("count=%d admissible=%d stable=%d compatible=%d best=%s bestRes=%s bestInertia=%s nativeMatch=%t verdict=%q", len(x.Candidates), x.AdmissibleAlternatingCandidates, x.StableCandidates, x.SplitCompatibleCandidates, x.BestCandidateName, f64(x.BestRelativeResidualToBK), x.BestInertia, x.NativeSKTwistMatchesBK, x.Verdict)
}
func FormatCrossProduct(x CrossProductCompatibilityAudit) string {
	return fmt.Sprintf("compactCross=%t omegaBAlt=%t antiRes=%s norm=%s inertia=%s stable=%t scaleToBK=%s relResBK=%s matchesBK=%t verdict=%q", x.CompactCrossProductDefined, x.OmegaBAlternating, f64(x.OmegaBAntisymmetryResidual), f64(x.OmegaBNorm), x.OmegaBInertia, x.OmegaBStable, f64(x.OmegaBScaleToBK), f64(x.OmegaBRelativeResidualToBK), x.OmegaBMatchesBK, x.Verdict)
}
func FormatInterpretation(x InterpretationAudit) string {
	return fmt.Sprintf("gOmegaGK=%t BKpolarized=%t twistFound=%t fused=%t classification=%q verdict=%q", x.GOmegaAlignedWithGK, x.BKIsHodgePolarizedCompactMetric, x.NativeSplitCompatibleTwistFound, x.CompactOmegaAndBKFused, x.Classification, x.Verdict)
}
func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("physical=%t boundary=%t sevenOver72=%t flavor=%t scalarRG=%t higgs=%t ckm_pmns=%t gauge=%t splitG2=%t verdict=%q", x.ClaimsPhysicalSpacetime, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsFlavor, x.ClaimsScalarRG, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.ClaimsSplitG2, x.Verdict)
}
