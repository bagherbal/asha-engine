package generation2hitchinchannelalgebraselectionruleaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate648Inheritance) string {
	return fmt.Sprintf("slotInherited=%t p=%d q=%d d=%d slotSupported=%t d_eq_q=%t pqTheorem=%t cubicTheorem=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.SlotMultiplicityInherited, x.PositiveDim, x.NegativeDim, x.CubicDegree, x.SlotSourceSupported, x.ASHADimEqualsDegree, x.GeneralPQDimensionTheorem, x.CubicSlotTheoremCertified, x.FullSymbolicHitchin, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate648FirewallPreserved, x.Verdict)
}

func FormatSupport(x TwoComponentTensorSupportAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s(alias=%s support=%t norms=[%.6g %.6g %.6g] leak=%t role=%s)", r.Family, r.Alias, r.Supported, r.Omega1NormSq, r.Omega2NormSq, r.OmegaBNormSq, r.ResidualLeak, r.Role))
	}
	return fmt.Sprintf("A=%s B=%s onlyAB=%t verdict=%q rows=%s", x.AName, x.BName, x.OnlyAAndBSupported, x.Verdict, strings.Join(parts, "; "))
}

func FormatExpansion(x OrderedCubicExpansionAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s/%s class=%s nonzero=%t plus=%.6g minus=%.6g mixed=%.3g status=%s", r.RouteName, r.Channel, r.Class, r.Nonzero, r.PlusMeanUnit, r.MinusMeanUnit, r.MixedFrobenius, r.SelectionRuleStatus))
	}
	return fmt.Sprintf("routes=%d channelsPerRoute=%d nonzeroPerRoute=%d AAA=%t AAB=%t ABB_BBB=%t mixed=%t verdict=%q rows=%s", x.RouteCount, x.ChannelsPerRoute, x.NonzeroChannelsPerRoute, x.AAAOnlyPositive, x.AABOnlyNegative, x.ABBBBBClean, x.MixedBlocksClean, x.Verdict, strings.Join(parts, "; "))
}

func FormatPositiveAAA(x PositiveChannelAudit) string {
	return fmt.Sprintf("rows=%d unit=%t onlyPlus=%t coeff=%.6g source=%q verdict=%q", len(x.Rows), x.AAAContributesUnit, x.AAAContributesOnlyPlus, x.UnitCoefficient, x.SourceClassification, x.Verdict)
}

func FormatNegativeAAB(x NegativeChannelAudit) string {
	parts := []string{}
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s/%s minus=%.6g plus=%.6g", r.RouteName, r.Class, r.MinusMeanUnit, r.PlusMeanUnit))
	}
	return fmt.Sprintf("count=%d eachUnit=%t combined=%.6g source=%q verdict=%q rows=%s", x.NegativeOrderedChannelCount, x.EachAABContributesMinusUnit, x.CombinedNegativeCoefficient, x.SourceClassification, x.Verdict, strings.Join(parts, "; "))
}

func FormatVanishing(x VanishingCancellationAudit) string {
	parts := []string{}
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s/%s vanishing=%t source=%s", r.RouteName, r.Class, r.Vanishing, r.SourceCandidate))
	}
	return fmt.Sprintf("all=%t symbolic=%t mechanism=%q verdict=%q rows=%s", x.AllVanishOrProjectAway, x.SymbolicMechanismCertified, x.MechanismCandidate, x.Verdict, strings.Join(parts, "; "))
}

func FormatOffBlock(x OffBlockCancellationAudit) string {
	return fmt.Sprintf("maxMixed=%.6g channelwise=%t pairwise=%t fullSum=%t mechanism=%q verdict=%q", x.MaxMixedFrobenius, x.ChannelwiseZero, x.PairwiseCancellation, x.FullSumCancellation, x.MechanismClassification, x.Verdict)
}

func FormatSlotFormula(x SlotFormulaDerivation) string {
	return fmt.Sprintf("p=%d q=%d d=%d formula=%q normSq=%.12g cos=%q %.15g rho=%q %.15g recovers=%t verdict=%q", x.PositiveDim, x.NegativeDim, x.SlotMultiplicity, x.GSlotFormula, x.NormSquared, x.CosineFormula, x.Cosine, x.ResidualSquaredFormula, x.ResidualSquared, x.RecoversGate642Angle, x.Verdict)
}

func FormatCoincidence(x DimensionCoincidenceAudit) string {
	return fmt.Sprintf("d=%d q=%d equal=%t slotOnly=%t dimTheorem=%t interpretation=%q verdict=%q", x.SlotMultiplicity, x.NegativeDim, x.EqualInASHACarrier, x.SupportsSlotTheoremOnly, x.SupportsDimensionTheorem, x.Interpretation, x.Verdict)
}

func FormatReadiness(x SymbolicTheoremReadiness) string {
	return fmt.Sprintf("candidate=%q finite=%t symbolic=%t missing=%q verdict=%q", x.CandidateTheorem, x.FiniteChannelRuleSupported, x.FullSymbolicTheoremCertified, x.MissingProofObject, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicChannelSelection, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
