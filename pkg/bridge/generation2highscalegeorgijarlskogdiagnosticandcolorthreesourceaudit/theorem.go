package generation2highscalegeorgijarlskogdiagnosticandcolorthreesourceaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-798-HIGH-SCALE-GEORGI-JARLSKOG-DIAGNOSTIC-COLOR-THREE-SOURCE"
	theoremName = "Gate 798 — High-Scale Georgi-Jarlskog Diagnostic and Color-Three Source Audit"
)

func Generation2HighScaleGeorgiJarlskogDiagnosticAndColorThreeSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 798 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate798}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 797 pattern airlock", Passed: a.Gate797.Inherited && a.Gate797.NEffThreeStatusKnown && a.Gate797.CurrentThreeSource == "color-tripled top dominance", Detail: a.Gate797.Verdict},
			{Name: "inherit N_eff three source status", Passed: a.Gate797.NotGenerationTriality && a.Gate797.NotD4Triality && a.Gate797.NotNativeYukawaTheorem && closeAbs(a.Gate797.NEff, nEffSnapshot, 1e-15) && closeAbs(a.Gate797.CYukawa, cYukawaSnapshot, 1e-15), Detail: StatusNEffThreeInherited},
			{Name: "define Georgi-Jarlskog diagnostic hypothesis", Passed: a.Hypothesis.Defined && strings.Contains(a.Hypothesis.LowScale, "color") && strings.Contains(a.Hypothesis.HighScale, "Georgi") && containsAll(a.Hypothesis.Blocked, []string{"native Yukawa", "generation", "D4"}), Detail: FormatHypothesis(a.Hypothesis)},
			{Name: "define multi-scale Yukawa ledger requirement", Passed: a.Requirement.Defined && !a.Requirement.SingleScaleOK && containsAll(a.Requirement.Fields, []string{"low_scale_mu", "high_scale_mu", "RG_scheme", "threshold", "normalization"}) && containsAll(a.Requirement.MinimumValues, []string{"y_d", "y_s", "y_b", "y_e", "y_mu", "y_tau"}), Detail: FormatRequirement(a.Requirement)},
			{Name: "define Georgi-Jarlskog ratio diagnostics", Passed: a.GJ.Defined && a.GJ.HighScale && containsAll(a.GJ.Ratios, []string{"y_b/y_tau", "y_mu/(3y_s)", "(3y_e)/y_d"}) && strings.Contains(a.GJ.ClosureNorm, "Delta_GJ"), Detail: FormatGJ(a.GJ)},
			{Name: "type N_eff and GJ threes as distinct readouts", Passed: a.Comparison.Recorded && a.Comparison.LawfulComparison && !a.Comparison.SameTypedObject && containsAll(a.Comparison.BlockedShortcuts, []string{"N_eff", "GJ", "triality"}), Detail: FormatComparison(a.Comparison)},
			{Name: "define FN compatibility check", Passed: a.FN.Defined && containsAll(a.FN.Requires, []string{"full Yukawa ledger", "epsilon"}) && containsAll(a.FN.Blocked, []string{"native charge", "silent epsilon", "invent trace atoms"}), Detail: FormatDiagnostic(a.FN)},
			{Name: "define Koide scale compatibility check", Passed: a.Koide.Defined && containsAll(a.Koide.Requires, []string{"y_e", "y_mu", "y_tau"}) && containsAll(a.Koide.Blocked, []string{"derive charged-lepton", "generation", "N_eff"}), Detail: FormatDiagnostic(a.Koide)},
			{Name: "audit hexagram motif firewall", Passed: a.Hexagram.Audited && containsAll(a.Hexagram.LawfulReadings, []string{"A2", "SU(3)", "hexagon"}) && containsAll(a.Hexagram.ForbiddenUse, []string{"visual proof", "D4", "Yukawa"}) && containsAll(a.Hexagram.RequiredTheorem, []string{"typed carrier", "trace-readout"}), Detail: FormatHexagram(a.Hexagram)},
			{Name: "define diagnostic outcome table", Passed: a.Outcomes.Defined && len(a.Outcomes.Outcomes) == 4 && containsAll(a.Outcomes.Outcomes, []string{"GJ ratios close", "N_eff near 3 but GJ fails", "both fail"}), Detail: FormatOutcomes(a.Outcomes)},
			{Name: "preserve C_Higgs formula firewall", Passed: a.Impact.Recorded && a.Impact.Formula == "C_Higgs=(3/N_eff)C_History" && a.Impact.ValidatedLedgerCanSetC && !a.Impact.PatternsModifyFormula && a.Impact.CHiggsLevel == "Level B" && closeAbs(a.Impact.CurrentCHiggs, cHiggsSnapshot, 1e-15), Detail: FormatImpact(a.Impact)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Recommended, "Native Three-Source") && !a.Branch.MultiScaleLedger && !a.Branch.SingleScaleLedger && !a.Branch.AnyLedger && containsAll(a.Branch.Alternatives, []string{"Georgi-Jarlskog", "Sector Contribution"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.GJNativeYukawa && !a.Firewalls.GUTUnificationTheorem && !a.Firewalls.FNChargeTheorem && !a.Firewalls.KoideNativeYukawa && !a.Firewalls.VisualMotifProof && !a.Firewalls.NEffGenerationTheorem && !a.Firewalls.NEffD4Triality && !a.Firewalls.CHiggsLevelC && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate798, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth, FormatHypothesis(a.Hypothesis), FormatRequirement(a.Requirement), FormatGJ(a.GJ), FormatComparison(a.Comparison), FormatDiagnostic(a.FN), FormatDiagnostic(a.Koide), FormatHexagram(a.Hexagram), FormatOutcomes(a.Outcomes), FormatImpact(a.Impact), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
