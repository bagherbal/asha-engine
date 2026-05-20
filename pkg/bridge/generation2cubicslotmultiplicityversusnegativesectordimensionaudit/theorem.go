package generation2cubicslotmultiplicityversusnegativesectordimensionaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CubicSlotMultiplicityVersusNegativeSectorDimensionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 648 — Cubic Slot Multiplicity versus Negative-Sector Dimension Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate648 cubic slot multiplicity versus negative-sector dimension audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate647 finite contraction ledger and firewalls", Passed: a.Inherited.ContractionLedgerInherited && a.Inherited.RouteCount == 3 && a.Inherited.PositiveDim == 4 && a.Inherited.NegativeDim == 3 && a.Inherited.CubicDegree == 3 && a.Inherited.HasThreeNegativeChannels && !a.Inherited.GeneralPQDimensionClaim && !a.Inherited.FullSymbolicTheorem && !a.Inherited.SplitG2Certified && !a.Inherited.BoundaryStressAssignment && !a.Inherited.SevenOver72Theorem && !a.Inherited.ScalarFlavorTransport && !a.Inherited.PhysicalMetric && a.Inherited.Gate647FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "compute per-direction and total trace separation", Passed: a.TraceAudit.AllRoutesPass && len(a.TraceAudit.Rows) == 3, Detail: FormatTraceAudit(a.TraceAudit)},
			{Name: "compute ordered cubic slot contributions", Passed: a.SlotAudit.SlotSourceSupported && a.SlotAudit.EachNegativeChannelUnit && a.SlotAudit.NegativeChannelCount == 3 && len(a.SlotAudit.NegativeChannels) == 9 && len(a.SlotAudit.PositiveChannels) == 3, Detail: FormatSlotAudit(a.SlotAudit)},
			{Name: "compute negative-index contribution diagnostic", Passed: a.NegativeIndexAudit.AllRoutesUniformByIndex && len(a.NegativeIndexAudit.Rows) == 3, Detail: FormatNegativeIndex(a.NegativeIndexAudit)},
			{Name: "disambiguate dimension formula from cubic-slot formula", Passed: a.FormulaAudit.FinalRayCannotDistinguish && a.FormulaAudit.LedgerSelectsSlotSource && !a.FormulaAudit.GeneralPQDimensionTheorem && a.FormulaAudit.Row.CoincideInASHA, Detail: FormatFormula(a.FormulaAudit)},
			{Name: "record diagnostic ablations as non-native source probes", Passed: a.Ablations.AllDiagnosticOnly && a.Ablations.SlotSourceDominates && len(a.Ablations.Diagnostics) >= 4, Detail: FormatAblations(a.Ablations)},
			{Name: "refine theorem target without promoting a general p,q theorem", Passed: a.TheoremTarget.ASHACoincidenceCertified && !a.TheoremTarget.GeneralPQDimensionTheorem && !a.TheoremTarget.CubicSlotTheoremCertified, Detail: FormatTheoremTarget(a.TheoremTarget)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsGeneralPQDimensionTheorem && !a.Firewalls.ClaimsCubicSlotTheorem && !a.Firewalls.ClaimsFullSymbolicHitchin && !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate648Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate648 corrects the Gate647 interpretation: the finite ledger directly certifies three ordered cubic negative slots as the immediate source of the -3 coefficient.  The equality dim(K_7^-)=3 is recorded as an ASHA carrier coincidence with the cubic Hitchin degree, not as a general p,q dimension theorem.")
		if a.TheoremTarget.GeneralPQDimensionTheorem || a.Firewalls.ClaimsGeneralPQDimensionTheorem {
			notes = append(notes, "WARNING_GENERAL_P_Q_DIMENSION_THEOREM_PROMOTION_BLOCKED")
		}
		if !strings.Contains(a.TheoremTarget.Verdict, StatusNoGeneralPQDimensionTheorem) {
			notes = append(notes, "WARNING_MISSING_NO_GENERAL_P_Q_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
