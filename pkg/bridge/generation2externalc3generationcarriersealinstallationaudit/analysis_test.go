package generation2externalc3generationcarriersealinstallationaudit

import (
	"strings"
	"testing"
)

func TestGate960BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.GenerationCarrierAvailable || !a.Decision.ExternalC3SealInstalled || !a.Decision.R4MayProceedUnderSeals {
		t.Fatalf("external seal not installed: %#v", a.Decision)
	}
	if a.NativeMultiplicityTheorem || a.Decision.NativeGenerationCarrier {
		t.Fatalf("overclaimed native generation: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ParentAirlockStillOpen {
		t.Fatalf("missing seal preservation or parent route: %#v", a.Decision)
	}
	if a.Decision.FlavorOrientationMapAvailable || a.Decision.IndividualYukawaAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed downstream claims: %#v", a.Decision)
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, AuditSupports(a.Audits), AuditFailures(a.Audits)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate960Theorem(t *testing.T) {
	res := Generation2ExternalC3GenerationCarrierSealInstallationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER", "FAILED_ROUTE_R3_TRACEBRIDGE_REMAINS_DUALSEALED_NOT_NATIVE", "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP", "FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
