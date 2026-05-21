package generation2trialityairlockgenerationcarrierroutebifurcationaudit

import (
	"strings"
	"testing"
)

func TestGate957BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.R3DualSealRequired || a.NativeR3 || a.GenerationCarrierCertified || a.Decision.GenerationCarrierCertified || a.Decision.FlavorOrientationCertified || a.Decision.IndividualYukawaCertified || a.Decision.PhysicalAssignmentCertified || a.Decision.OfficialLedgerUpdate {
		t.Fatalf("overclaimed status: %#v", a)
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 || a.K7MinusDimension+a.K7PlusDimension != K7Dim {
		t.Fatalf("bad K7 polarity dimensions")
	}
	if a.R3TraceRows != R3TraceRows {
		t.Fatalf("R3 rows must remain aggregate")
	}
	if a.Parent.Identified || a.Parent.NativeD4Spin8SourceCertified || a.Parent.ReplacesGate955AbstractC3 || a.Parent.HasOrderThreePermutation {
		t.Fatalf("parent triality board overcertified: %#v", a.Parent)
	}
	if a.Airlock.ToLambda4Certified || a.Airlock.ToK7Certified || a.Airlock.PreservesK7 || a.Airlock.PreservesK7Minus || a.Airlock.SelectsAlternativeThreeCarrier {
		t.Fatalf("triality airlock overcertified: %#v", a.Airlock)
	}
	if !a.Airlock.SelectsNoCanonicalThreeCarrier || a.Decision.K7MinusRouteReopened || a.Decision.AlternativeCarrierSelected || !a.Decision.AlternativeSearchRequired {
		t.Fatalf("bad route bifurcation decision: %#v %#v", a.Airlock, a.Decision)
	}
	if len(a.Items) != 7 {
		t.Fatalf("expected 7 items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ItemSupports(a.Items), ItemFailures(a.Items)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate957Theorem(t *testing.T) {
	res := Generation2TrialityAirlockGenerationCarrierRouteBifurcationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_AIRLOCK", "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP", "FAILED_ROUTE_R3_TRACE_ROWS_ARE_NOT_GENERATION_LABELS"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
