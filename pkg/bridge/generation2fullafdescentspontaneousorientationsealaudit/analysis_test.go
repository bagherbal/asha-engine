package generation2fullafdescentspontaneousorientationsealaudit

import (
	"strings"
	"testing"
)

func TestGate948BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Boundary.ScalarSourceSealed || !a.Boundary.FullAFDescentBlocked || !a.Boundary.StableInAFOrient || !a.Boundary.SpontaneousOrientSeal {
		t.Fatalf("bad boundary flags: %#v", a.Boundary)
	}
	if a.Boundary.NativeOrientation || a.Boundary.NativeR3 || a.Boundary.OfficialLedgerUpdate || a.Boundary.PhysicalAssignment || a.Boundary.GenerationFlavorClaims {
		t.Fatalf("Gate 948 overclaimed: %#v", a.Boundary)
	}
	if len(a.Items) != 6 {
		t.Fatalf("expected six items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Failures, ItemFailures(a.Items)), "\n")
	for _, want := range []string{
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NO_NATIVE_SPONTANEOUS_ORIENTATION_THEOREM",
		"FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing failure %s", want)
		}
	}
}

func TestGate948Theorem(t *testing.T) {
	res := Generation2FullAFDescentSpontaneousOrientationSealAuditTheorem().Verify()
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
		"CONDITIONAL_SUPPORT_TRACEBRIDGE_STABLE_IN_A_F_ORIENT_LAYER",
		"CONDITIONAL_SUPPORT_SPONTANEOUS_ORIENTATION_SEAL_IS_LAWFUL_BRIDGE_LAYER",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
