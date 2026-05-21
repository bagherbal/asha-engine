package generation2k7minustrialitytracebodycouplingpreconditionaudit

import (
	"strings"
	"testing"
)

func TestGate954BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.DualSealRequired || !a.Coupling.DualSealRequired {
		t.Fatalf("dual seal must be required")
	}
	if a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		t.Fatalf("overclaimed downstream status: %#v", a)
	}
	if a.Coupling.GenerationMapCertified || a.Coupling.FlavorMapCertified || a.Coupling.IntertwinerCertified || a.Coupling.IndividualYukawaAllowed || a.Coupling.FlavorBacksolveAllowed {
		t.Fatalf("overclaimed coupling map or flavor permission: %#v", a.Coupling)
	}
	if a.Coupling.CarrierDimension != 3 || len(a.Coupling.ActionLanes) != 3 {
		t.Fatalf("bad K7/triality shape: %#v", a.Coupling)
	}
	if len(a.Items) != 4 {
		t.Fatalf("expected 4 items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ItemSupports(a.Items), ItemFailures(a.Items)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate954Theorem(t *testing.T) {
	res := Generation2K7MinusTrialityTracebodyCouplingPreconditionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "K7MinusTrialityTracebodyIntertwiner"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
