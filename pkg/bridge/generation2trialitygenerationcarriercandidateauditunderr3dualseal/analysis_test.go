package generation2trialitygenerationcarriercandidateauditunderr3dualseal

import (
	"strings"
	"testing"
)

func TestGate952BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.DualSealRequired {
		t.Fatalf("dual seal must be required")
	}
	if a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		t.Fatalf("overclaimed downstream status: %#v", a)
	}
	if len(a.Items) != 5 {
		t.Fatalf("expected 5 items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ItemSupports(a.Items), ItemFailures(a.Items)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate952Theorem(t *testing.T) {
	res := Generation2TrialityGenerationCarrierCandidateAuditUnderR3DualSealTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
