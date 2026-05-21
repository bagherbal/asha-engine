package generation2alternativegenerationcarriersearchaudit

import (
	"strings"
	"testing"
)

func TestGate958BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.R3DualSealRequired || a.NativeR3 || a.GenerationCarrierCertified || a.Decision.NativeGenerationCarrier || a.Decision.AlternativeCandidateFound {
		t.Fatalf("overclaimed native/generation status: %#v", a.Decision)
	}
	if a.Decision.FlavorOrientationCertified || a.Decision.IndividualYukawaCertified || a.Decision.PhysicalAssignmentCertified || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed downstream status: %#v", a.Decision)
	}
	if !a.Decision.ExternalSealAvailable || !a.Decision.ParentAirlockRequired {
		t.Fatalf("expected remaining routes to be external seal or parent airlock")
	}
	if a.K7PlusDimension != 4 || a.K7MinusDimension != 3 || a.K7PlusDimension+a.K7MinusDimension != K7Dim {
		t.Fatalf("bad K7 dimensions")
	}
	if a.BoundaryPairDimension != 2 || a.R3TraceRows != R3TraceRows {
		t.Fatalf("bad inherited B2/R3 context")
	}
	if len(a.Candidates) != 5 {
		t.Fatalf("expected 5 candidates, got %d", len(a.Candidates))
	}
	for _, c := range a.Candidates {
		if c.Status == CandidateStronglyValid || c.TypedTracebodyMap || c.UsesFlavorBacksolve || c.UsesR3RowsAsLabels {
			t.Fatalf("candidate overclaimed or used forbidden route: %#v", c)
		}
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, CandidateSupports(a.Candidates), CandidateFailures(a.Candidates)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate958Theorem(t *testing.T) {
	res := Generation2AlternativeGenerationCarrierSearchAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP", "FAILED_ROUTE_NO_CANONICAL_ALTERNATIVE_THREE_CARRIER_FOUND", "FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER", "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
