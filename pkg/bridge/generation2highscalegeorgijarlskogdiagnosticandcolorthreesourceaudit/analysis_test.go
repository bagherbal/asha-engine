package generation2highscalegeorgijarlskogdiagnosticandcolorthreesourceaudit

import (
	"strings"
	"testing"
)

func TestGate798HypothesisAndRequirements(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate797.Inherited || a.Gate797.CurrentThreeSource != "color-tripled top dominance" || !a.Gate797.NotGenerationTriality || !a.Gate797.NotD4Triality {
		t.Fatalf("bad inheritance: %+v", a.Gate797)
	}
	if !a.Hypothesis.Defined || !strings.Contains(a.Hypothesis.LowScale, "color") || !strings.Contains(a.Hypothesis.HighScale, "Georgi") || !containsAll(a.Hypothesis.Blocked, []string{"native Yukawa", "generation", "D4"}) {
		t.Fatalf("bad hypothesis: %s", FormatHypothesis(a.Hypothesis))
	}
	if !a.Requirement.Defined || a.Requirement.SingleScaleOK || !containsAll(a.Requirement.Fields, []string{"low_scale_mu", "high_scale_mu", "RG_scheme", "threshold", "normalization"}) || !containsAll(a.Requirement.MinimumValues, []string{"y_d", "y_s", "y_b", "y_e", "y_mu", "y_tau"}) {
		t.Fatalf("bad requirement: %s", FormatRequirement(a.Requirement))
	}
}

func TestGate798DiagnosticsAndDistinctReadouts(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.GJ.Defined || !a.GJ.HighScale || !containsAll(a.GJ.Ratios, []string{"y_b/y_tau", "y_mu/(3y_s)", "(3y_e)/y_d"}) || !strings.Contains(a.GJ.ClosureNorm, "Delta_GJ") {
		t.Fatalf("bad GJ: %s", FormatGJ(a.GJ))
	}
	if !a.Comparison.Recorded || !a.Comparison.LawfulComparison || a.Comparison.SameTypedObject || !containsAll(a.Comparison.BlockedShortcuts, []string{"N_eff", "GJ", "triality"}) {
		t.Fatalf("bad comparison: %s", FormatComparison(a.Comparison))
	}
	if !a.FN.Defined || !containsAll(a.FN.Requires, []string{"full Yukawa ledger", "epsilon"}) || !containsAll(a.FN.Blocked, []string{"native charge", "silent epsilon", "invent trace atoms"}) {
		t.Fatalf("bad FN: %s", FormatDiagnostic(a.FN))
	}
	if !a.Koide.Defined || !containsAll(a.Koide.Requires, []string{"y_e", "y_mu", "y_tau"}) || !containsAll(a.Koide.Blocked, []string{"derive charged-lepton", "generation", "N_eff"}) {
		t.Fatalf("bad Koide: %s", FormatDiagnostic(a.Koide))
	}
}

func TestGate798FirewallsImpactAndBranch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hexagram.Audited || !containsAll(a.Hexagram.LawfulReadings, []string{"A2", "SU(3)", "hexagon"}) || !containsAll(a.Hexagram.ForbiddenUse, []string{"visual proof", "D4", "Yukawa"}) {
		t.Fatalf("bad hexagram: %s", FormatHexagram(a.Hexagram))
	}
	if !a.Outcomes.Defined || len(a.Outcomes.Outcomes) != 4 || !containsAll(a.Outcomes.Outcomes, []string{"GJ ratios close", "N_eff near 3 but GJ fails", "both fail"}) {
		t.Fatalf("bad outcomes: %s", FormatOutcomes(a.Outcomes))
	}
	if !a.Impact.Recorded || a.Impact.PatternsModifyFormula || a.Impact.CHiggsLevel != "Level B" || !closeAbs(a.Impact.CurrentCHiggs, cHiggsSnapshot, 1e-15) {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "Native Three-Source") || a.Branch.MultiScaleLedger || a.Branch.SingleScaleLedger || a.Branch.AnyLedger {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.GJNativeYukawa || a.Firewalls.GUTUnificationTheorem || a.Firewalls.FNChargeTheorem || a.Firewalls.KoideNativeYukawa || a.Firewalls.VisualMotifProof || a.Firewalls.NEffGenerationTheorem || a.Firewalls.NEffD4Triality || a.Firewalls.CHiggsLevelC || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate798TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "does not claim") || !strings.Contains(a.FinalStatement, "low-scale color-three") || !strings.Contains(a.FinalStatement, "high-scale Georgi-Jarlskog") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2HighScaleGeorgiJarlskogDiagnosticAndColorThreeSourceAuditTheorem().Verify()
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
