package generation2hitchinchannelsignequalunitcalibrationaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate650Inheritance) string {
	return fmt.Sprintf("degree=%t ledger=%t positiveAAA=%t negativeAAB=%t mixedZero=%t signGap=%t slot=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.DegreeSelectionInherited, x.SectorLedgerDefined, x.PositiveAAAOnly, x.NegativeAABOnly, x.MixedZeroByDegree, x.SignCalibrationGapInherited, x.SlotFormulaRecovered, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate650FirewallPreserved, x.Verdict)
}

func FormatOrientation(x OrientationVolumeAudit) string {
	return fmt.Sprintf("p=%d q=%d volume=%q interior=%q wedge=%q hitchin=%q SK=%q oct=%q compatible=%t conventionSign=%t verdict=%q", x.PositiveDim, x.NegativeDim, x.VolumeForm, x.InteriorConvention, x.WedgeConvention, x.HitchinNormalization, x.SKAction, x.OctonionicOrientation, x.OrientationCompatible, x.ConventionDependentSign, x.Verdict)
}

func formatChannelRows(rows []ChannelBilinearRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s/%s mean=%.6g target=%.6g trace=%.6g rank=%d anis=%.3g off=%.3g match=%t", r.Channel, r.Block, r.MeanCoefficient, r.TargetCoefficient, r.Trace, r.Rank, r.AnisotropyResidual, r.OffBlockResidual, r.MatchesCalibratedRay))
	}
	return strings.Join(parts, "; ")
}

func FormatMaps(x SurvivingChannelBilinearMaps) string {
	return fmt.Sprintf("computed=%t survivors=%v verdict=%q rows=%s", x.MapsComputed, x.OnlySurvivors, x.Verdict, formatChannelRows(x.Rows))
}

func FormatPositive(x PositiveUnitAudit) string {
	return fmt.Sprintf("c+=%.6g scalar=%t anis=%.3g onlyPlus=%t verdict=%q row=%s", x.CPlus, x.ScalarMultipleOfP, x.AnisotropyResidual, x.ContributesOnlyPlus, x.Verdict, formatChannelRows([]ChannelBilinearRow{x.AAA}))
}

func FormatNegative(x NegativeEqualUnitAudit) string {
	return fmt.Sprintf("c+=%.6g cAAB=%.6g cABA=%.6g cBAA=%.6g equalMinus=%t scalar=%t combined=%.6g anis=%.3g verdict=%q rows=%s", x.CPlus, x.CAAB, x.CABA, x.CBAA, x.EqualToMinusCPlus, x.EachScalarMultipleOfP, x.CombinedCoefficient, x.CombinedAnisotropy, x.Verdict, formatChannelRows(x.Rows))
}

func FormatSign(x SignSourceAudit) string {
	return fmt.Sprintf("negative=%t equal=%t primary=%q secondary=%v basisFree=%t needsProof=%t verdict=%q", x.FiniteNegativeSignObserved, x.FiniteEqualUnitObserved, x.PrimaryTypedSourceCandidate, x.SecondarySourceCandidates, x.BasisFreeSourceCertified, x.RequiresCalibrationIdentityProof, x.Verdict)
}

func FormatRoutes(x RouteUniversalityAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s AAA=%.3g AAB=%.3g ABA=%.3g BAA=%.3g equal=%t ray=%t c=%.3g", r.RouteName, r.AAAUnit, r.AABUnit, r.ABAUnit, r.BAAUnit, r.EqualPattern, r.ReconstructsProjectorRay, r.RouteNormalizedCoefficient))
	}
	return fmt.Sprintf("all=%t sameAfterNorm=%t routeDependent=%t verdict=%q rows=%s", x.AllRoutesPass, x.SamePatternAfterNorm, x.RouteDependentMagnitude, x.Verdict, strings.Join(parts, "; "))
}

func FormatReconstruction(x ReconstructionAudit) string {
	return fmt.Sprintf("c=%.6g pos=%.6g neg=%.6g norm=%.6g cos=%.12g rho2=%.12g reconstruct=%t recovers=%t formula=%q verdict=%q", x.C, x.PositiveCoefficient, x.NegativeCoefficient, x.NormSquared, x.Cosine, x.ResidualSquared, x.ReconstructsPPlusMinus3P, x.RecoversGate642Angle, x.Formula, x.Verdict)
}

func FormatTheorem(x TheoremTarget) string {
	return fmt.Sprintf("finite=%t symbolic=%t gap=%q theorem=%q verdict=%q", x.FiniteCalibrationIdentityPasses, x.FullSymbolicCalibrationTheorem, x.RemainingGap, x.CandidateTheorem, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicCalibration, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
