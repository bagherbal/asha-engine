package generation2r3alphayukawatracebridgepretestexecutionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate937BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Classification != Classification || a.ShortStatus != ShortStatus || a.Truth != FinalTruth {
		t.Fatalf("bad identity: %#v", a)
	}
	if !positiveOK(a.Positive) {
		t.Fatalf("positive checks failed: %#v", a.Positive)
	}
	if !negativeOK(a.Negative) {
		t.Fatalf("negative checks failed: %#v", a.Negative)
	}
	if !containsAll(a.Supports, Supports()) {
		t.Fatalf("supports missing")
	}
	if !containsAll(a.Failures, Failures()) {
		t.Fatalf("failures missing")
	}
	if !numericOK(a.Numeric) {
		t.Fatalf("bad numeric: %s", FormatNumeric(a.Numeric))
	}
}

func TestGate937Numerics(t *testing.T) {
	n := ExecuteNumeric()
	if math.Abs(n.AlphaLinear+n.AlphaQuadratic-n.AlphaB) > floatTolTight {
		t.Fatalf("alpha components mismatch")
	}
	if math.Abs(n.NEff-NEffOperator) > floatTol || math.Abs(n.CYukawa-CYukawaOperator) > floatTol {
		t.Fatalf("operator diagnostics mismatch: %s", FormatTrace(n))
	}
	if n.BareLinear8 == n.AlphaLinear || n.BareQuadratic70 == n.AlphaQuadratic || n.CrossLanePolluted == n.AlphaB || n.CommonDenom10 == n.AlphaB {
		t.Fatalf("negative numeric route did not differ: %s", FormatNumeric(n))
	}
	if !traceWeightsPositive(n.AlphaB) {
		t.Fatalf("trace weights not positive")
	}
}

func TestGate937Theorem(t *testing.T) {
	res := Generation2R3AlphaYukawaTraceBridgePreTestExecutionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(append(Statuses(), Supports()...), Failures()...), []string{
		FinalTruth,
		Classification,
		ShortStatus,
		InheritedStatus,
		"PASS_REDUCED_B2_RESPONSE_EXPANDS_TO_S_PLUS_S2",
		"PASS_BOUNDARY_ACTIVATION_MEASURE_RECONSTRUCTS_ALPHA_B",
		"PASS_OPERATOR_N_EFF_RECONSTRUCTED",
		"REJECT_CROSS_LANE_POLLUTED_ALPHA",
		"REJECT_BARE_DENOMINATOR_8",
		NextGate,
	}...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
