package heavycarrierdecayaudit

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Operators.NativePortalSearchFailed || a.Operators.DecayOperatorsDerived != 0 {
		t.Fatalf("unexpected decay portal result: %s", FormatOperatorAudit(a.Operators))
	}
	if a.Lifetime.PassesBBN || !a.Lifetime.FailsBBNByOperatorAbsence {
		t.Fatalf("unexpected BBN lifetime result: %s", FormatLifetime(a.Lifetime))
	}
	if !a.Summary.FatalPathology || a.Summary.CosmologyCleared {
		t.Fatalf("cosmology should fail: %s", FormatSummary(a.Summary))
	}
	if a.Firewall.DecayOperatorInvented || a.Firewall.MassSplittingInvented || a.Firewall.DarkMatterClaimed || a.Firewall.RelicAbundanceComputed {
		t.Fatalf("firewall broken: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	r := HeavyCarrierDecayRelicSafetyAuditTheorem().Run()
	if !r.Passed() {
		t.Fatalf("theorem failed:\n%s", r.Details())
	}
}
