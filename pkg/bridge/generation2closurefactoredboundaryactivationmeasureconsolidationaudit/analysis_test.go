package generation2closurefactoredboundaryactivationmeasureconsolidationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate934BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Classification != Classification || a.ShortStatus != ShortStatus || a.Truth != FinalTruth {
		t.Fatalf("bad identity: %#v", a)
	}
	if !componentsOK(a.Components) {
		t.Fatalf("bad components: %#v", a.Components)
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
	if math.Abs(a.Numeric.AlphaLinear+a.Numeric.AlphaQuadratic-a.Numeric.AlphaB) > 1e-18 {
		t.Fatalf("alpha sum mismatch")
	}
}

func TestGate934Theorem(t *testing.T) {
	res := Generation2ClosureFactoredBoundaryActivationMeasureConsolidationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, InheritedStatus, F0, F1, F2, AdmissibleSupportChain, ClosureOperator, ThetaFunctor, BoundaryActivationMeasure, AlphaFormula, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
