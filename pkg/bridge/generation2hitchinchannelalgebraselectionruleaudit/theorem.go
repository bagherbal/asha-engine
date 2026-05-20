package generation2hitchinchannelalgebraselectionruleaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HitchinChannelAlgebraSelectionRuleAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 649 — Hitchin AAA/AAB Channel Algebra Selection Rule Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate649 Hitchin channel algebra selection-rule audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate648 slot-source result and firewalls", Passed: a.Inherited.SlotMultiplicityInherited && a.Inherited.PositiveDim == 4 && a.Inherited.NegativeDim == 3 && a.Inherited.CubicDegree == 3 && a.Inherited.SlotSourceSupported && a.Inherited.ASHADimEqualsDegree && !a.Inherited.GeneralPQDimensionTheorem && !a.Inherited.CubicSlotTheoremCertified && !a.Inherited.FullSymbolicHitchin && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate648FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit two-component A/B tensor support", Passed: a.Support.OnlyAAndBSupported && len(a.Support.Rows) == 4 && a.Support.AName == "Ω++-" && a.Support.BName == "Ω---", Detail: FormatSupport(a.Support)},
			{Name: "expand ordered cubic AAA/AAB/ABB/BBB channels", Passed: a.Expansion.RouteCount == 3 && a.Expansion.ChannelsPerRoute == 8 && a.Expansion.NonzeroChannelsPerRoute == 4 && a.Expansion.AAAOnlyPositive && a.Expansion.AABOnlyNegative && a.Expansion.ABBBBBClean && a.Expansion.MixedBlocksClean && len(a.Expansion.Rows) == 24, Detail: FormatExpansion(a.Expansion)},
			{Name: "audit AAA positive channel", Passed: a.PositiveAAA.AAAContributesUnit && a.PositiveAAA.AAAContributesOnlyPlus && len(a.PositiveAAA.Rows) == 3, Detail: FormatPositiveAAA(a.PositiveAAA)},
			{Name: "audit AAB/ABA/BAA negative channels", Passed: a.NegativeAAB.EachAABContributesMinusUnit && a.NegativeAAB.NegativeOrderedChannelCount == 3 && len(a.NegativeAAB.Rows) == 9, Detail: FormatNegativeAAB(a.NegativeAAB)},
			{Name: "audit ABB/BAB/BBA/BBB vanishing or project-away routes", Passed: a.Vanishing.AllVanishOrProjectAway && !a.Vanishing.SymbolicMechanismCertified && len(a.Vanishing.Rows) == 12, Detail: FormatVanishing(a.Vanishing)},
			{Name: "audit off-block cancellation", Passed: a.OffBlock.ChannelwiseZero && a.OffBlock.MaxMixedFrobenius < tol, Detail: FormatOffBlock(a.OffBlock)},
			{Name: "derive slot-supported angle formula", Passed: a.SlotFormula.SlotMultiplicity == 3 && a.SlotFormula.NormSquared == 31 && a.SlotFormula.RecoversGate642Angle, Detail: FormatSlotFormula(a.SlotFormula)},
			{Name: "separate slot source from dimension coincidence", Passed: a.Coincidence.EqualInASHACarrier && a.Coincidence.SupportsSlotTheoremOnly && !a.Coincidence.SupportsDimensionTheorem, Detail: FormatCoincidence(a.Coincidence)},
			{Name: "preserve symbolic theorem gap", Passed: a.Readiness.FiniteChannelRuleSupported && !a.Readiness.FullSymbolicTheoremCertified, Detail: FormatReadiness(a.Readiness)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicChannelSelection && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate649Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate649 sharpens Gate648: the directly witnessed source is the AAA/AAB Hitchin channel algebra, not a general p,q dimension theorem.  The equality d=q=3 is recorded as an ASHA carrier coincidence.")
		if a.Readiness.FullSymbolicTheoremCertified || a.Firewalls.ClaimsFullSymbolicChannelSelection {
			notes = append(notes, "WARNING_SYMBOLIC_CHANNEL_SELECTION_THEOREM_PROMOTION_BLOCKED")
		}
		if !strings.Contains(a.Readiness.Verdict, StatusNoFullSymbolicChannelSelectionTheorem) {
			notes = append(notes, "WARNING_MISSING_SYMBOLIC_CHANNEL_SELECTION_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
