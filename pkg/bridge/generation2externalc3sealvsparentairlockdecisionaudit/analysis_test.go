package generation2externalc3sealvsparentairlockdecisionaudit

import (
	"strings"
	"testing"
)

func TestGate959BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.R3DualSealRequired || a.NativeR3 || a.GenerationCarrierCertified || a.Decision.NativeGenerationCarrierCertified {
		t.Fatalf("overclaimed native status: %#v", a.Decision)
	}
	if !a.Decision.ActiveBoardExhausted || !a.Decision.ExternalC3SealAllowed || !a.Decision.ParentAirlockOnlyNativeRoute {
		t.Fatalf("bad fork decision: %#v", a.Decision)
	}
	if a.Decision.FlavorDerivationAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed downstream claims: %#v", a.Decision)
	}
	if len(a.Routes) != 4 {
		t.Fatalf("expected 4 routes, got %d", len(a.Routes))
	}
	for _, r := range a.Routes {
		if r.Status == RouteCertifiedNative || r.NativeCarrier || r.UsesFlavorBacksolve || r.UsesR3RowsAsLabels {
			t.Fatalf("route overclaimed or used forbidden input: %#v", r)
		}
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, RouteSupports(a.Routes), RouteFailures(a.Routes)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate959Theorem(t *testing.T) {
	res := Generation2ExternalC3SealVsParentAirlockDecisionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "FAILED_ROUTE_ACTIVE_BOARD_NATIVE_GENERATION_CARRIER_EXHAUSTED_IN_CURRENT_CERTIFICATE", "FAILED_ROUTE_EXTERNAL_C3_NOT_NATIVE_GENERATION_CARRIER", "FAILED_ROUTE_NO_PARENT_AIRLOCK_CERTIFIED_YET", "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
