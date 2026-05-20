package generation2externalyukawainputchecklistandpatterndiagnosticairlockaudit

import (
	"strings"
	"testing"
)

func TestGate797ChecklistAndAtomProtocol(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate796.Inherited {
		t.Fatalf("bad inheritance: %+v", a.Gate796)
	}
	if !a.Checklist.Defined || !a.Checklist.RejectImplicit || !containsAll(a.Checklist.Fields, []string{"source_label", "scale_mu", "scheme", "color_convention", "neutrino_convention", "y_u", "y_e", "uncertainties"}) || !containsAll(a.Checklist.RequiredStatuses, []string{"scale_mu", "normalization", "color_convention", "neutrino_convention"}) {
		t.Fatalf("bad checklist: %s", FormatChecklist(a.Checklist))
	}
	if !a.AtomProtocol.Defined || a.AtomProtocol.AtomFormula != "x_f=y_f^2; x_f^2=y_f^4" || !a.AtomProtocol.RequiresConventionLock || !containsAll(a.AtomProtocol.CoefficientColorRules, []string{"a_u=3", "b_u=3"}) || !containsAll(a.AtomProtocol.Computes, []string{"a_ext", "b_ext", "N_eff_ext", "C_Yukawa_ext"}) {
		t.Fatalf("bad atom protocol: %s", FormatAtomProtocol(a.AtomProtocol))
	}
	if !a.Validation.Defined || a.Validation.SilentRescaleAllowed || !closeAbs(a.Validation.InheritedNEff, nEffInherited, 5e-16) || !containsAll(a.Validation.ClassifiedFailures, []string{"scale", "scheme", "normalization", "neutrino", "color"}) {
		t.Fatalf("bad validation: %s", FormatValidation(a.Validation))
	}
}

func TestGate797PatternDiagnosticsAreReadOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	diagnostics := []PatternDiagnostic{a.Koide, a.FN, a.BTau}
	for _, d := range diagnostics {
		if !d.Defined || !d.ReadOnly || d.CanPopulateAtomLedger {
			t.Fatalf("diagnostic not read-only: %s", FormatPattern(d))
		}
	}
	if !containsAll(a.Koide.Requires, []string{"y_e", "y_mu", "y_tau"}) || !containsAll(a.Koide.BlockedUse, []string{"populate", "derive N_eff", "PMNS"}) {
		t.Fatalf("bad Koide diagnostic: %s", FormatPattern(a.Koide))
	}
	if !containsAll(a.FN.Requires, []string{"full sector Yukawa ledger", "epsilon"}) || !containsAll(a.FN.BlockedUse, []string{"derive Yukawa", "epsilon", "native charge"}) {
		t.Fatalf("bad FN diagnostic: %s", FormatPattern(a.FN))
	}
	if !containsAll(a.BTau.Requires, []string{"y_b", "y_tau", "RG", "threshold"}) || !containsAll(a.BTau.BlockedUse, []string{"single-scale", "derive full Yukawa", "derive N_eff"}) {
		t.Fatalf("bad b-tau diagnostic: %s", FormatPattern(a.BTau))
	}
}

func TestGate797PriorityImpactBranchAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Priority.Recorded || !strings.HasPrefix(FormatPriority(a.Priority), "full sector/atom ledger") || !containsAll(a.Priority.Ranking, []string{"Froggatt-Nielsen", "Koide", "b-tau"}) {
		t.Fatalf("bad priority: %s", FormatPriority(a.Priority))
	}
	if !a.Impact.Recorded || !a.Impact.ValidatedLedgerCanConfirm || a.Impact.PatternsModifyFormula || a.Impact.CHiggsLevel != "Level B" || !closeAbs(a.Impact.CurrentCHiggs, cHiggsLevelB, 1e-16) {
		t.Fatalf("bad C_Higgs impact: %s", FormatImpact(a.Impact))
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "Holding Pattern") || a.Branch.ExternalLedger || a.Branch.NativeYukawa || a.Branch.D4Package || !containsAll(a.Branch.Alternatives, []string{"External Yukawa Ledger Validation", "Native Yukawa", "D4 Triality"}) {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.KoideNativeYukawa || a.Firewalls.FNNativeCharge || a.Firewalls.BTauNativeUnification || a.Firewalls.PatternFitIsTraceAtomData || a.Firewalls.ExternalLedgerNative || a.Firewalls.NEffGenerationTheorem || a.Firewalls.NEffD4Triality || a.Firewalls.CHiggsLevelC || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate797TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not import") || !strings.Contains(a.FinalStatement, "secondary read-only diagnostics") || !strings.Contains(a.FinalStatement, "full validated atom ledger remains the primary need") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2ExternalYukawaInputChecklistAndPatternDiagnosticAirlockAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
