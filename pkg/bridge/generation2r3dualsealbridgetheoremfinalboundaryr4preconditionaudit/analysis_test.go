package generation2r3dualsealbridgetheoremfinalboundaryr4preconditionaudit

import (
	"strings"
	"testing"
)

func TestGate949BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Boundary.TracebridgeTestPassed || !a.Boundary.ClosureFactored || !a.Boundary.ScalarSourceSeal || !a.Boundary.PostOrientationSeal {
		t.Fatalf("bad boundary flags: %#v", a.Boundary)
	}
	if a.Boundary.NativeR3 || a.Boundary.OfficialLedgerUpdate || a.Boundary.PhysicalAssignment || a.Boundary.GenerationCarrier || a.Boundary.FlavorOrientation || a.Boundary.IndividualYukawa || a.Boundary.R4NativeSpectrum {
		t.Fatalf("Gate 949 overclaimed: %#v", a.Boundary)
	}
	if !a.Policy.MayProceedUnderSeal || !a.Policy.RequiresScalarSourceSeal || !a.Policy.RequiresPostOrientSeal {
		t.Fatalf("bad policy: %#v", a.Policy)
	}
	if a.Policy.AllowsPhysicalAssignment || a.Policy.AllowsOfficialLedger || a.Policy.AllowsNativeSpectrumClaim {
		t.Fatalf("policy overclaims: %#v", a.Policy)
	}
	if len(a.Items) != 6 {
		t.Fatalf("expected six items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Failures, ItemFailures(a.Items)), "\n")
	for _, want := range []string{
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_R3_TRACEBRIDGE_NOT_NATIVE_THEOREM",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED",
		"FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_SPECTRUM_THEOREM",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing failure %s", want)
		}
	}
}

func TestGate949Theorem(t *testing.T) {
	res := Generation2R3DualSealBridgeTheoremFinalBoundaryR4PreconditionAuditTheorem().Verify()
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
		"CONDITIONAL_SUPPORT_R4_WORK_MAY_PROCEED_ONLY_UNDER_EXPLICIT_DUAL_SEAL",
		"CONDITIONAL_SUPPORT_NATIVE_R3_AND_R4_FLAVOR_REMAIN_SEPARATED",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		"FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
