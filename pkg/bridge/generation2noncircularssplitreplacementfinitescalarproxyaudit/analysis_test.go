package generation2noncircularssplitreplacementfinitescalarproxyaudit

import (
	"strings"
	"testing"
)

func TestGate946BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.TargetValue != Ssplit {
		t.Fatalf("bad target Ssplit: %.19g", a.TargetValue)
	}
	if len(a.Candidates) != 7 {
		t.Fatalf("expected seven candidates, got %d", len(a.Candidates))
	}
	if HasSuccessfulReplacement(a.Candidates) {
		t.Fatalf("no candidate should be a valid native replacement")
	}
	joined := strings.Join(appendAll(a.Failures, CandidateFailures(a.Candidates)), "\n")
	for _, want := range []string{
		"FAILED_ROUTE_D_BASE_REPLACEMENT_IS_REPARAMETERIZATION_OF_S_SPLIT",
		"FAILED_ROUTE_7_OVER_72_IS_NORMALIZATION_COEFFICIENT_NOT_S_SPLIT_SCALAR",
		"FAILED_ROUTE_FINITE_RANK_DATA_DO_NOT_CANONICALLY_GENERATE_S_SPLIT_MAGNITUDE",
		"FAILED_ROUTE_FIXED_POINT_SCALAR_RECOVERY_IS_CIRCULAR_WITH_TRACEBRIDGE_OUTPUT",
		"FAILED_ROUTE_ZERO_SCALAR_DOES_NOT_REPRODUCE_TRACEBRIDGE",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing failure %s", want)
		}
	}
}

func TestGate946Theorem(t *testing.T) {
	res := Generation2NonCircularSSplitReplacementFiniteScalarProxyAuditTheorem().Verify()
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
		"D_base rescaling",
		"pure rank ratio 7/72",
		"closure-measure fixed point",
		"FAILED_ROUTE_NO_NONCIRCULAR_NATIVE_S_SPLIT_PROXY_FOUND",
		"CONDITIONAL_SUPPORT_CERTIFICATE_II_REDUCES_TO_SCALAR_SOURCE_SEAL",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
