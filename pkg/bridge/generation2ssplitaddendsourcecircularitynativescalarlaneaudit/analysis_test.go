package generation2ssplitaddendsourcecircularitynativescalarlaneaudit

import (
	"strings"
	"testing"
)

func TestGate945BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.Expression != "S_split=(R_3-1)+lambda(Lambda_12)" || a.Value != Ssplit {
		t.Fatalf("bad expression/value: %s %.19g", a.Expression, a.Value)
	}
	if len(a.Addends) != 2 {
		t.Fatalf("expected two addends, got %d", len(a.Addends))
	}
	if !a.Addends[0].CircularRisk || a.Addends[0].NativeH72Scalar {
		t.Fatalf("R3-1 addend should be circular-risk and non-native: %#v", a.Addends[0])
	}
	if !a.Addends[1].BridgeHistoryScalar || a.Addends[1].NativeH72Scalar {
		t.Fatalf("lambda addend should be bridge/history and non-native: %#v", a.Addends[1])
	}
	if !a.ScalarLane.CanonicalAddition || a.ScalarLane.SsplitNative || a.ScalarLane.BothAddendsNative {
		t.Fatalf("bad scalar lane: %#v", a.ScalarLane)
	}
	if !containsCurrentCircularOrBridgeRow(a.TruthTable) {
		t.Fatalf("truth table did not mark current bridge/circular state: %#v", a.TruthTable)
	}
	if !a.Certificate.TransportLayerStrong || !a.Certificate.CentralH72Compatible || a.Certificate.AddendSourceNative || a.Certificate.CertificatePassed {
		t.Fatalf("bad certificate status: %#v", a.Certificate)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing support/failure markers")
	}
}

func TestGate945Theorem(t *testing.T) {
	res := Generation2SSplitAddendSourceCircularityNativeScalarLaneAuditTheorem().Verify()
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
		"CONDITIONAL_SUPPORT_NATIVE_S_SPLIT_REQUIRES_COMMON_H72_SCALAR_LANE_CERTIFICATE",
		"FAILED_ROUTE_R3_MINUS_ONE_AS_INPUT_TO_R3_PROMOTION_IS_POTENTIALLY_CIRCULAR",
		"FAILED_ROUTE_LAMBDA_LAMBDA12_NOT_DERIVED_AS_NATIVE_H72_SCALAR",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED_BECAUSE_ADDEND_SOURCES_NOT_NATIVE",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
