package generation2flavororientationmappreconditionauditunderexternalc3andr3dualseal

import (
	"strings"
	"testing"
)

func TestGate961BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.ExternalC3DomainAvailable || a.Decision.ExternalC3NativeGeneration {
		t.Fatalf("bad external C3 domain typing: %#v", a.Decision)
	}
	if !a.Decision.LedgerCodomainTyped || !a.Decision.FlavorOrientationMapRequired || a.Decision.FlavorOrientationMapCertified {
		t.Fatalf("bad map requirement state: %#v", a.Decision)
	}
	if a.Decision.ObservedFlavorInputAllowed || a.Decision.FlavorFormulaBacksolveAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.CKMPMNSAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed flavor inputs or downstream outputs: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalC3SealPreserved || !a.Decision.CanProceedToConstructionAudit {
		t.Fatalf("seals or next construction audit not preserved: %#v", a.Decision)
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, AuditSupports(a.Audits), AuditFailures(a.Audits)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate961Theorem(t *testing.T) {
	res := Generation2FlavorOrientationMapPreconditionAuditUnderExternalC3AndR3DualSealTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "CONDITIONAL_SUPPORT_FLAVOR_ORIENTATION_MAP_IS_NEXT_REQUIRED_OBJECT", "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP_CERTIFIED_YET", "FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR", "FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
