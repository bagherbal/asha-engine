package generation2cubicslotmultiplicityversusnegativesectordimensionaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return fmt.Sprintf("%g", x)
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(x Gate647Inheritance) string {
	return fmt.Sprintf("ledger=%t routes=%d p=%d q=%d degree=%d ray=%q threeNegative=%t pqClaim=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.ContractionLedgerInherited, x.RouteCount, x.PositiveDim, x.NegativeDim, x.CubicDegree, x.BlockRay, x.HasThreeNegativeChannels, x.GeneralPQDimensionClaim, x.FullSymbolicTheorem, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate647FirewallPreserved, x.Verdict)
}

func FormatTraceAudit(x PerDirectionAndTotalTraceAudit) string {
	parts := []string{fmt.Sprintf("all=%t verdict=%q interpretation=%q", x.AllRoutesPass, x.Verdict, x.Interpretation)}
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("[%s c+=%s c-=%s c-/c+=%s T+=%s T-=%s T-/T+=%s expected=%s pass=%t]", r.RouteName, f64(r.PositiveMean), f64(r.NegativeMean), f64(r.PerDirectionRatio), f64(r.PositiveTrace), f64(r.NegativeTrace), f64(r.TotalTraceRatio), f64(r.ExpectedTraceRatio), r.Passed))
	}
	return strings.Join(parts, "\n")
}

func FormatSlotAudit(x OrderedSlotContributionAudit) string {
	parts := []string{fmt.Sprintf("negativeChannels=%d expected=%d eachUnit=%t removeOneDelta=%s slotSource=%t verdict=%q", x.NegativeChannelCount, x.ExpectedCubicSlotCount, x.EachNegativeChannelUnit, f64(x.RemovingOneChannelDelta), x.SlotSourceSupported, x.Verdict)}
	for _, r := range x.PositiveChannels {
		parts = append(parts, fmt.Sprintf("positive[%s %s plusUnit=%s minusUnit=%s]", r.RouteName, r.Channel, f64(r.PositiveMeanUnit), f64(r.NegativeMeanUnit)))
	}
	for _, r := range x.NegativeChannels {
		parts = append(parts, fmt.Sprintf("negative[%s %s plusUnit=%s minusUnit=%s unitNegative=%t]", r.RouteName, r.Channel, f64(r.PositiveMeanUnit), f64(r.NegativeMeanUnit), r.ContributesUnitNegative))
	}
	return strings.Join(parts, "\n")
}

func FormatNegativeIndex(x NegativeIndexContributionAudit) string {
	parts := []string{fmt.Sprintf("uniform=%t verdict=%q conclusion=%q", x.AllRoutesUniformByIndex, x.Verdict, x.SlotVsDirectionConclusion)}
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("[%s dirs=%d perDir=%s perChannelPerDir=%s totalTraceWeight=%s dimensionTraceOnly=%t]", r.RouteName, r.NegativeDirections, f64(r.PerNegativeDirectionWeight), f64(r.PerChannelPerDirectionWeight), f64(r.TotalNegativeTraceWeight), r.DimensionChangesTraceOnly))
	}
	return strings.Join(parts, "\n")
}

func FormatFormula(x FormulaDisambiguationAudit) string {
	r := x.Row
	return fmt.Sprintf("p=%d q=%d degree=%d dimFormula=%q slotFormula=%q dimNormSq=%s slotNormSq=%s coincide=%t supported=%q finalRayCannotDistinguish=%t ledgerSelectsSlot=%t pqTheorem=%t verdict=%q", r.PositiveDim, r.NegativeDim, r.CubicDegree, r.DimensionFormula, r.SlotFormula, f64(r.DimensionNormSq), f64(r.SlotNormSq), r.CoincideInASHA, r.SupportedFormula, x.FinalRayCannotDistinguish, x.LedgerSelectsSlotSource, x.GeneralPQDimensionTheorem, x.Verdict)
}

func FormatAblations(x SyntheticAblativeDiagnostics) string {
	parts := []string{fmt.Sprintf("diagnosticOnly=%t slotDominates=%t verdict=%q", x.AllDiagnosticOnly, x.SlotSourceDominates, x.Verdict)}
	for _, d := range x.Diagnostics {
		parts = append(parts, fmt.Sprintf("[%s diagnostic=%t expected=%q coeff=%s slot=%t dimensionOnly=%t]", d.Name, d.DiagnosticOnly, d.ExpectedEffect, f64(d.ObservedCoefficient), d.SupportsSlotSource, d.SupportsDimensionOnly))
	}
	return strings.Join(parts, "\n")
}

func FormatTheoremTarget(x SymbolicTheoremTargetUpdate) string {
	return fmt.Sprintf("old=%q refined=%q support=%q pqTheorem=%t cubicSlotTheorem=%t coincidence=%t missing=%q verdict=%q", x.OldTarget, x.RefinedTarget, x.FiniteDataSupport, x.GeneralPQDimensionTheorem, x.CubicSlotTheoremCertified, x.ASHACoincidenceCertified, x.MissingProofObject, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("pq=%t cubicSlot=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsGeneralPQDimensionTheorem, x.ClaimsCubicSlotTheorem, x.ClaimsFullSymbolicHitchin, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
