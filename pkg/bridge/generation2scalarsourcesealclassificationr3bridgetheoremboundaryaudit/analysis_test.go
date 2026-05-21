package generation2scalarsourcesealclassificationr3bridgetheoremboundaryaudit

import (
	"strings"
	"testing"
)

func TestGate947BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Boundary.BridgeTestPassed || !a.Boundary.ScalarSourceSealed || a.Boundary.NativeR3 || a.Boundary.OfficialLedgerUpdate {
		t.Fatalf("bad boundary classification: %#v", a.Boundary)
	}
	if len(a.Items) != 6 {
		t.Fatalf("expected six boundary items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Failures, ItemFailures(a.Items)), "\n")
	for _, want := range []string{
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing failure %s", want)
		}
	}
}

func TestGate947Theorem(t *testing.T) {
	res := Generation2ScalarSourceSealClassificationR3BridgeTheoremBoundaryAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{
		Verdict,
		Classification,
		ShortStatus,
		"TEST_PASSED_SCALAR_SOURCE_SEALED_BRIDGE_THEOREM_CANDIDATE",
		"CONDITIONAL_SUPPORT_R3_TRACEBRIDGE_CAN_BE_CLASSIFIED_AS_SCALAR_SOURCE_SEALED_BRIDGE_THEOREM_CANDIDATE",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		NextGateA,
		NextGateB,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
