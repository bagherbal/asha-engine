package generation2hitchinsectordegreetopformselectionaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HitchinSectorDegreeTopFormSelectionRuleAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 650 — Hitchin Sector-Degree Top-Form Selection Rule Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate650 Hitchin sector-degree top-form selection audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate649 channel algebra and firewalls", Passed: a.Inherited.ChannelAlgebraInherited && a.Inherited.TwoComponentSupport && a.Inherited.AAAChannelAudited && a.Inherited.AABChannelAudited && a.Inherited.VanishingAudited && a.Inherited.OffBlockAudited && a.Inherited.SlotFormulaDerived && a.Inherited.DEqualsQCoincidence && !a.Inherited.FullSymbolicChannelTheorem && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate649FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "define sector-degree ledger", Passed: a.Ledger.PositiveDim == 4 && a.Ledger.NegativeDim == 3 && a.Ledger.TopDegree == (Degree{4, 3}) && a.Ledger.AHasDegree21 && a.Ledger.BHasDegree03 && len(a.Ledger.Rows) == 2, Detail: FormatLedger(a.Ledger)},
			{Name: "audit positive block AAA-only by top-form degree", Passed: a.Positive.AAAOnlySurvives && len(a.Positive.SurvivingChannels) == 1 && a.Positive.SurvivingChannels[0] == "AAA", Detail: FormatPositive(a.Positive)},
			{Name: "audit negative block AAB/ABA/BAA by top-form degree", Passed: a.Negative.AABPlacementsOnly && a.Negative.AllowedPlacementCount == 3 && containsAll(a.Negative.SurvivingChannels, []string{"AAB", "ABA", "BAA"}), Detail: FormatNegative(a.Negative)},
			{Name: "audit mixed block zero by top-form degree", Passed: !a.Mixed.AnySurvivesByDegree && a.Mixed.MixedBlockZeroByDegree, Detail: FormatMixed(a.Mixed)},
			{Name: "preserve sign/equal-unit calibration gap", Passed: a.Sign.DegreeRuleCertifiesSupport && a.Sign.Gate649CertifiesFiniteSigns && !a.Sign.EqualUnitWeightCertifiedByDegree && a.Sign.RequiresCalibrationIdentity, Detail: FormatSign(a.Sign)},
			{Name: "state degree-selection theorem target without promotion", Passed: a.Theorem.DegreeSelectionSupported && !a.Theorem.FullSymbolicTheoremCertified, Detail: FormatTheorem(a.Theorem)},
			{Name: "derive slot formula only conditional on calibration identity", Passed: a.Slot.PositiveDim == 4 && a.Slot.NegativeDim == 3 && a.Slot.SlotMultiplicity == 3 && a.Slot.NormSquared == 31 && a.Slot.RecoversGate642Angle, Detail: FormatSlot(a.Slot)},
			{Name: "record d=q ASHA carrier resonance", Passed: a.Resonance.CubicDegree == 3 && a.Resonance.NegativeDim == 3 && a.Resonance.EqualInASHACarrier, Detail: FormatResonance(a.Resonance)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsFullSymbolicDegreeTheorem && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate650Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate650 upgrades Gate649: AAA/AAB support is explained by 4|3 top-form degree saturation, while sign and equal unit weights remain calibration-dependent.")
		if a.Theorem.FullSymbolicTheoremCertified || a.Firewalls.ClaimsFullSymbolicDegreeTheorem {
			notes = append(notes, "WARNING_DEGREE_SELECTION_THEOREM_PROMOTION_BLOCKED")
		}
		if !strings.Contains(a.Sign.Verdict, StatusSignUnitRequiresCalibration) {
			notes = append(notes, "WARNING_MISSING_CALIBRATION_GAP_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
