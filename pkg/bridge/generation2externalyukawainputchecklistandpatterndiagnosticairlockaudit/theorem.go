package generation2externalyukawainputchecklistandpatterndiagnosticairlockaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-797-EXTERNAL-YUKAWA-INPUT-CHECKLIST-PATTERN-DIAGNOSTIC-AIRLOCK"
	theoremName = "Gate 797 — External Yukawa Input Checklist and Pattern-Diagnostic Airlock Audit"
)

func Generation2ExternalYukawaInputChecklistAndPatternDiagnosticAirlockAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 797 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate797}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 796 external intake airlock", Passed: a.Gate796.Inherited && a.Gate796.Verdict == StatusGate796Inherited, Detail: a.Gate796.Verdict},
			{Name: "classify pattern diagnostics as read-only tests", Passed: a.Koide.ReadOnly && a.FN.ReadOnly && a.BTau.ReadOnly && !a.Koide.CanPopulateAtomLedger && !a.FN.CanPopulateAtomLedger && !a.BTau.CanPopulateAtomLedger, Detail: StatusPatternsReadOnly},
			{Name: "define external Yukawa input checklist", Passed: a.Checklist.Defined && a.Checklist.RejectImplicit && containsAll(a.Checklist.Fields, []string{"source_label", "scale_mu", "scheme", "color_convention", "neutrino_convention", "y_u", "y_e", "uncertainties"}) && containsAll(a.Checklist.RequiredStatuses, []string{"scale_mu", "normalization", "color_convention", "neutrino_convention"}), Detail: FormatChecklist(a.Checklist)},
			{Name: "define atom construction protocol", Passed: a.AtomProtocol.Defined && a.AtomProtocol.AtomFormula == "x_f=y_f^2; x_f^2=y_f^4" && a.AtomProtocol.RequiresConventionLock && containsAll(a.AtomProtocol.CoefficientColorRules, []string{"a_u=3", "b_u=3"}) && containsAll(a.AtomProtocol.Computes, []string{"a_ext", "b_ext", "N_eff_ext", "C_Yukawa_ext"}), Detail: FormatAtomProtocol(a.AtomProtocol)},
			{Name: "define aggregate validation protocol", Passed: a.Validation.Defined && closeAbs(a.Validation.InheritedA, aInherited, 1e-16) && closeAbs(a.Validation.InheritedB, bInherited, 1e-16) && closeAbs(a.Validation.InheritedNEff, nEffInherited, 5e-16) && !a.Validation.SilentRescaleAllowed && containsAll(a.Validation.ClassifiedFailures, []string{"scale", "scheme", "normalization", "neutrino", "color"}), Detail: FormatValidation(a.Validation)},
			{Name: "define Koide diagnostic airlock", Passed: a.Koide.Defined && a.Koide.ReadOnly && !a.Koide.CanPopulateAtomLedger && containsAll(a.Koide.Requires, []string{"y_e", "y_mu", "y_tau"}) && containsAll(a.Koide.BlockedUse, []string{"populate", "derive N_eff", "PMNS"}), Detail: FormatPattern(a.Koide)},
			{Name: "define Froggatt-Nielsen diagnostic airlock", Passed: a.FN.Defined && a.FN.ReadOnly && !a.FN.CanPopulateAtomLedger && containsAll(a.FN.Requires, []string{"full sector Yukawa ledger", "epsilon"}) && containsAll(a.FN.BlockedUse, []string{"derive Yukawa", "epsilon", "native charge"}), Detail: FormatPattern(a.FN)},
			{Name: "define b-tau unification diagnostic airlock", Passed: a.BTau.Defined && a.BTau.ReadOnly && !a.BTau.CanPopulateAtomLedger && containsAll(a.BTau.Requires, []string{"y_b", "y_tau", "RG", "threshold"}) && containsAll(a.BTau.BlockedUse, []string{"single-scale", "derive full Yukawa", "derive N_eff"}), Detail: FormatPattern(a.BTau)},
			{Name: "record pattern priority classification", Passed: a.Priority.Recorded && containsAll(a.Priority.Ranking, []string{"full sector/atom ledger", "Froggatt-Nielsen", "Koide", "b-tau"}) && strings.HasPrefix(FormatPriority(a.Priority), "full sector/atom ledger"), Detail: FormatPriority(a.Priority)},
			{Name: "record C_Higgs impact", Passed: a.Impact.Recorded && a.Impact.ValidatedLedgerCanConfirm && !a.Impact.PatternsModifyFormula && a.Impact.CHiggsLevel == "Level B" && closeAbs(a.Impact.CurrentCHiggs, cHiggsLevelB, 1e-16), Detail: FormatImpact(a.Impact)},
			{Name: "record branch decision", Passed: a.Branch.Recorded && strings.Contains(a.Branch.Recommended, "Holding Pattern") && !a.Branch.ExternalLedger && !a.Branch.NativeYukawa && !a.Branch.D4Package && containsAll(a.Branch.Alternatives, []string{"External Yukawa Ledger Validation", "Native Yukawa", "D4 Triality"}), Detail: FormatBranch(a.Branch)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.KoideNativeYukawa && !a.Firewalls.FNNativeCharge && !a.Firewalls.BTauNativeUnification && !a.Firewalls.PatternFitIsTraceAtomData && !a.Firewalls.ExternalLedgerNative && !a.Firewalls.NEffGenerationTheorem && !a.Firewalls.NEffD4Triality && !a.Firewalls.CHiggsLevelC && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate797, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatChecklist(a.Checklist), FormatAtomProtocol(a.AtomProtocol), FormatValidation(a.Validation), FormatPattern(a.Koide), FormatPattern(a.FN), FormatPattern(a.BTau), FormatPriority(a.Priority), FormatImpact(a.Impact), FormatBranch(a.Branch), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
