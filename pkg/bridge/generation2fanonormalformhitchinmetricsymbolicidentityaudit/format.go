package generation2fanonormalformhitchinmetricsymbolicidentityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate652Inheritance) string {
	return fmt.Sprintf("fano=%t Bvol=%t Atrip=%t wedge=%t quat=%t AAA=%t AAB=%t finite=%t pgFano=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.FanoNormalFormInherited, x.BVolumeForm, x.ATwoFormTriple, x.WedgeOrthonormality, x.QuaternionicTriple, x.AAAChannelFinite, x.AABChannelsFinite, x.FiniteNormalFormIdentities, x.FullBasisFreeFanoTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.Gate652FirewallPreserved, x.Verdict)
}

func FormatPositive(x SymbolicPositiveBlockDerivation) string {
	return fmt.Sprintf("domain=%q c+=%.6g target=%q scalar=%t anis=%.3g normalForm=%t wedge=%t symbolic=%t expr=%q hitchin=%q verdict=%q", x.Domain, x.CPositive, x.Target, x.ScalarMultipleOfP, x.AnisotropyResidual, x.UsesNormalForm, x.UsesWedgeIdentity, x.SymbolicDerivation, x.Expression, x.HitchinFactor, x.Verdict)
}

func formatNegativeRows(rows []NegativeSymbolicChannel) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, fmt.Sprintf("%s coeff=%.6g target=%.6g scalar=%t vol=%t wedge=%t anis=%.3g sign=%q expr=%q", r.Channel, r.Coefficient, r.Target, r.ScalarMultipleOfP, r.UsesVolumeForm, r.UsesWedgeIdentity, r.AnisotropyResidual, r.SignSource, r.Expression))
	}
	return strings.Join(parts, "; ")
}

func FormatNegative(x SymbolicNegativeBlockDerivation) string {
	return fmt.Sprintf("c+=%.6g eachMinus=%t combined=%.6g target=%.6g residual=%.3g signLocated=%t sign=%q symbolic=%t verdict=%q rows=%s", x.CPositive, x.EachEqualsMinusC, x.CombinedCoefficient, x.CombinedTarget, x.CombinedResidual, x.NegativeSignLocated, x.SignLocation, x.SymbolicDerivation, x.Verdict, formatNegativeRows(x.Rows))
}

func FormatMixed(x SymbolicMixedBlockVanishing) string {
	return fmt.Sprintf("cases=%v norm=%.3g zero=%t degree=%q support=%q verdict=%q", x.Cases, x.MixedBlockNorm, x.SymbolicallyZero, x.DegreeSaturationReason, x.ContractionSupportReason, x.Verdict)
}

func FormatNormalization(x NormalizationAudit) string {
	return fmt.Sprintf("c+=%.6g cAAB=%.6g cABA=%.6g cBAA=%.6g absEqual=%t rescale=%t residual=%.3g scale=%q verdict=%q", x.CPositive, x.CAAB, x.CABA, x.CBAA, x.AllEqualAbs, x.RequiresRescale, x.Residual, x.CommonScale, x.Verdict)
}

func FormatRoutes(x RouteIndependenceAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s scale=%.6g same=%t residual=%.3g form=%q identity=%q", r.RouteName, r.RouteScale, r.ReducesToSameForm, r.Residual, r.NormalForm, r.SymbolicIdentity))
	}
	return fmt.Sprintf("all=%t sameIdentity=%t routeOnly=%t verdict=%q rows=%s", x.AllRoutesReduce, x.SameSymbolicIdentity, x.RouteDependentOnly, x.Verdict, strings.Join(parts, "; "))
}

func FormatFinalIdentity(x FinalSymbolicIdentity) string {
	return fmt.Sprintf("positive=%t negative=%t mixed=%t equalC=%t closed=%t pgFano=%t identity=%q ghat=%q cos=%q rho=%q verdict=%q", x.PositiveBlockPasses, x.NegativeBlockPasses, x.MixedBlockPasses, x.EqualNormalizationPasses, x.InternalMechanismClosed, x.FullPGToFanoSourceTheorem, x.Identity, x.GHat, x.CosTheta, x.RhoSquared, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
