package pevobservabilityaudit

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Summary.EWPOSafe || !a.Summary.HiggsLoopSafe || !a.Summary.DirectReachSafe {
		t.Fatalf("unexpected precision/direct failure: %s", FormatSummary(a.Summary))
	}
	if !a.Summary.CosmologyWarning || !a.Cosmology.StableColoredRelicWarning {
		t.Fatalf("missing cosmology warning: %s", FormatCosmology(a.Cosmology))
	}
	if a.Firewall.DecayOperatorInvented || a.Firewall.HeavyHiggsYukawaInvented || a.Firewall.DarkMatterClaimed || a.Firewall.PhysicalObservationClaimed {
		t.Fatalf("firewall broken: %s", FormatFirewall(a.Firewall))
	}
}

func TestTheorem(t *testing.T) {
	r := PeVThresholdIndirectSignatureObservabilityAuditTheorem().Run()
	if !r.Passed() {
		t.Fatalf("theorem failed:\n%s", r.Details())
	}
}
