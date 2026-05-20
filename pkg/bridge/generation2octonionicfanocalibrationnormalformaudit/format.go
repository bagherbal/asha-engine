package generation2octonionicfanocalibrationnormalformaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate651Inheritance) string {
	return fmt.Sprintf("calibration=%t AAA=%t AAB=%t reconstruct=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.CalibrationInherited, x.AAAUnit, x.AABEqualNegativeUnits, x.ReconstructsPPlusMinus3, x.FullSymbolicCalibration, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate651FirewallPreserved, x.Verdict)
}

func FormatBVolume(x NegativeVolumeFormAudit) string {
	return fmt.Sprintf("basis=%v beta=%.6g sign=%d residual=%.3g volume=%t verdict=%q", x.Basis, x.Beta, x.OrientationSign, x.ResidualAgainstVolume, x.BIsVolumeForm, x.Verdict)
}

func formatTwoFormRows(rows []TwoFormRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s∧%s norm2=%.6g innerOther=%.3g wedgeSelf=%.6g cross=%.3g selfDual=%+d extracted=%t", r.Name, r.Eta, r.NormSquared, r.InnerWithOthers, r.WedgeSelfCoefficient, r.WedgeCrossResidual, r.SelfDualSign, r.Extracted))
	}
	return strings.Join(parts, "; ")
}

func FormatAExtract(x ATwoFormExtractionAudit) string {
	return fmt.Sprintf("formula=%q all=%t orthogonal=%t equalNorms=%t wedge=%t residual=%.3g verdict=%q rows=%s", x.Formula, x.AllExtracted, x.OrthogonalTriple, x.EqualNorms, x.WedgeOrthonormal, x.Residual, x.Verdict, formatTwoFormRows(x.Rows))
}

func FormatQuaternionic(x QuaternionicTripleAudit) string {
	return fmt.Sprintf("endomorphisms=%t wedge=%t quaternionic=%t residual=%.3g convention=%q verdict=%q", x.FormsDefineEndomorphisms, x.WedgeIdentityPasses, x.QuaternionicIdentities, x.IdentityResidual, x.OrientationConvention, x.Verdict)
}

func FormatAAA(x AAAChannelDerivation) string {
	return fmt.Sprintf("alpha=%.6g beta=%.6g c+=%.6g scalar=%t anis=%.3g source=%q input=%q hitchin=%q verdict=%q", x.Alpha, x.Beta, x.CPositive, x.ScalarMultipleOfP, x.AnisotropyResidual, x.DerivationSource, x.InputNormalForm, x.HitchinFactor, x.Verdict)
}

func formatNegRows(rows []NegativeChannelDerivation) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s coeff=%.6g target=%.6g scalar=%t anis=%.3g sign=%q", r.Channel, r.Coefficient, r.Target, r.ScalarMultipleOfP, r.AnisotropyResidual, r.SignSource))
	}
	return strings.Join(parts, "; ")
}

func FormatAAB(x AABChannelDerivationAudit) string {
	return fmt.Sprintf("c+=%.6g equalMinus=%t combined=%.6g residual=%.3g verdict=%q rows=%s", x.CPositive, x.EqualToMinusPositive, x.CombinedCoefficient, x.CombinedResidual, x.Verdict, formatNegRows(x.Rows))
}

func FormatEqualUnit(x EqualUnitSourceAudit) string {
	return fmt.Sprintf("sameAlphaBeta=%t fano=%t quaternionic=%t routeOnly=%t basisFree=%t verdict=%q", x.SameAlphaBetaNormalization, x.FanoIncidenceSymmetry, x.QuaternionicNormalization, x.RouteSpecificOnly, x.BasisFreeProofCertified, x.Verdict)
}

func FormatRoutes(x RouteUniversalityAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s b=%.3g a=%.3g wedge=%.3g norm=%t pattern=%q", r.RouteName, r.BVolumeResidual, r.AExtractionResidual, r.WedgeIdentityResidual, r.ReducesAfterNorm, r.ChannelPattern))
	}
	return fmt.Sprintf("all=%t same=%t routeScale=%t verdict=%q rows=%s", x.AllRoutesReduce, x.SameNormalFormAfterNorm, x.RouteDependentScale, x.Verdict, strings.Join(parts, "; "))
}

func FormatTheorem(x TheoremTarget) string {
	return fmt.Sprintf("finite=%t symbolic=%t gap=%q theorem=%q verdict=%q", x.FiniteNormalFormIdentitiesPass, x.FullSymbolicOctonionicTheorem, x.RemainingGap, x.CandidateTheorem, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicOctonionicTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
