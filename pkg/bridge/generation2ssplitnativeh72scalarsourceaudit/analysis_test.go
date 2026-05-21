package generation2ssplitnativeh72scalarsourceaudit

import (
	"strings"
	"testing"
)

func TestGate944BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Expression.Dimensionless || !a.Expression.H72Compatible || a.Expression.NativeCertified {
		t.Fatalf("bad expression status: %#v", a.Expression)
	}
	if len(a.Components) != 2 {
		t.Fatalf("expected two components, got %d", len(a.Components))
	}
	for _, c := range a.Components {
		if c.NativeCertified || !c.BridgeHistoryUse {
			t.Fatalf("component should remain bridge/history and non-native: %#v", c)
		}
	}
	if a.Criteria.Status != SourceNativeMissing || len(a.Criteria.Blocked) == 0 {
		t.Fatalf("bad criteria: %#v", a.Criteria)
	}
	if !a.Certificate.DescentSupported || a.Certificate.SourceNative || a.Certificate.CertificatePassed {
		t.Fatalf("bad certificate status: %#v", a.Certificate)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing supports or failures")
	}
}

func TestGate944Theorem(t *testing.T) {
	res := Generation2SSplitNativeH72ScalarSourceAuditTheorem().Verify()
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
		"S_split=(R_3-1)+lambda(Lambda_12)",
		"R_3-1",
		"lambda(Lambda_12)",
		"CONDITIONAL_SUPPORT_TRANSPORT_NO_LONGER_MAIN_WOUND_AFTER_GATE943",
		"FAILED_ROUTE_R3_MINUS_ONE_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
		"FAILED_ROUTE_LAMBDA_LAMBDA12_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
		"FAILED_ROUTE_S_SPLIT_REMAINS_BRIDGE_HISTORY_SCALAR_INPUT",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
