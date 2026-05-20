package generation2realformtrialityairlockandnativestatusfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate801InheritanceAndStatusLevels(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inheritance.VolumeSquare != -1 || a.Inheritance.RealChiralityCertified || !strings.Contains(a.Inheritance.Outcome, "Outcome C") {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.StatusLevels.Defined || !strings.Contains(a.StatusLevels.CurrentStatus, "T1") || !a.StatusLevels.NotNative || !a.StatusLevels.NotYukawa {
		t.Fatalf("bad status levels: %s", FormatLevels(a.StatusLevels))
	}
}

func TestGate801AirlocksAndDescent(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ComplexAirlock.Defined || !containsAll(a.ComplexAirlock.Supports, []string{StatusComplexAuxSearch}) || !containsAll(a.ComplexAirlock.Failures, []string{StatusComplexAirlockNotNative}) {
		t.Fatalf("bad complex airlock: %s", FormatAirlock(a.ComplexAirlock))
	}
	if !a.CompactAirlock.Defined || !containsAll(a.CompactAirlock.Failures, []string{StatusCompactNotNative, StatusNoWickTransport}) {
		t.Fatalf("bad compact airlock: %s", FormatAirlock(a.CompactAirlock))
	}
	if !a.SplitAirlock.Defined || !containsAll(a.SplitAirlock.Failures, []string{StatusSplitNotNative}) {
		t.Fatalf("bad split airlock: %s", FormatAirlock(a.SplitAirlock))
	}
	if !a.Descent.Defined || a.Descent.NativeImport || !containsAll(a.Descent.Failures, []string{StatusNoNativeImport, StatusAuxCannotBeNative}) {
		t.Fatalf("bad descent: %s", FormatDescent(a.Descent))
	}
}

func TestGate801TrilinearNEffAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Trilinear.Refined || !a.Trilinear.NotYukawa || !a.Trilinear.NoReadout || !containsAll(a.Trilinear.Failures, []string{StatusTrilinearNotYukawa, StatusNoTrialityYukawaReadout}) {
		t.Fatalf("bad trilinear: %s", FormatTrilinear(a.Trilinear))
	}
	if !a.NEff.Preserved || a.NEff.AirlockChangesC || a.NEff.ExplainsDelta || !containsAll(a.NEff.Failures, []string{StatusAirlockNoNEff, StatusAirlockNoNEffMinus3}) {
		t.Fatalf("bad N_eff firewall: %s", FormatNEff(a.NEff))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoYukawa || !a.Firewalls.NoNEff || !a.Firewalls.NoPoleMass || a.Firewalls.Verdict != StatusFirewallPreservedGate801 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate801TheoremStatusesAndFinal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Final, "airlocked auxiliary") || !strings.Contains(a.Final, "not a native Cl(1,7) theorem") || !strings.Contains(a.Branch.NextNative, "Gate 802") {
		t.Fatalf("bad final/branch: %s / %+v", a.Final, a.Branch)
	}
	res := Generation2RealFormTrialityAirlockAndNativeStatusFirewallAuditTheorem().Verify()
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
