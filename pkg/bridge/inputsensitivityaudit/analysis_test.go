package inputsensitivityaudit

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if a.Summary.Status != StatusConditionalPhenomenology {
		t.Fatalf("status=%s summary=%s", a.Summary.Status, FormatSummary(a.Summary))
	}
	if !a.Completeness.BottomYukawaIncluded || !a.Completeness.TauYukawaIncluded {
		t.Fatalf("bottom/tau not included: %s", FormatCompleteness(a.Completeness))
	}
	if !a.CentralFit.MatchingPlausible || a.CentralFit.ResidualOverEpsilon >= 1 {
		t.Fatalf("central fit not plausible: %s", FormatFit(a.CentralFit))
	}
	if a.Sensitivity.BrokenEnvelopeCases != 0 || a.Sensitivity.PlausibleCases != a.Sensitivity.CasesAudited {
		t.Fatalf("sensitivity broken: %s", FormatSensitivity(a.Sensitivity))
	}
	if a.Firewall.InputsTunedToForceZeroResidual || a.Firewall.MatchingCorrectionsDerived || a.Firewall.PhysicalPredictionClaimed {
		t.Fatalf("firewall broken: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	r := InputSensitivityBottomTauYukawaCompletenessAuditTheorem().Run()
	if !r.Passed() {
		t.Fatalf("theorem failed:\n%s", r.Details())
	}
}
