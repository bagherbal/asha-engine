package generation2relativerestmagnitudeoperatorandboundaryalphaactivationmapaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate825RelativeOperatorTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	alpha := a.Ledger.AlphaB
	if math.Abs(a.Operator.ARestOverT-3*alpha) > 1e-15 {
		t.Fatalf("bad rest trace: %s", FormatOperator(a.Operator))
	}
	wantBRest := 3*alpha*alpha - 6*math.Pow(alpha, 3) + 12*math.Pow(alpha, 4)
	if math.Abs(a.Operator.BRestOverT2-wantBRest) > 1e-21 {
		t.Fatalf("bad rest quartic: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.QRest-(1.0/3.0-(2.0/3.0)*alpha+(4.0/3.0)*alpha*alpha)) > 1e-15 {
		t.Fatalf("bad q_rest: %s", FormatOperator(a.Operator))
	}
}

func TestGate825FifthOrderClosureAndImpact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Operator.NEffOperator-a.Ledger.NEffBFN) > 1e-15 {
		t.Fatalf("operator should reproduce BFN closure within float precision: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Operator.SymbolicResidual+2.107593378826735e-16) > 1e-27 {
		t.Fatalf("bad symbolic residual: %s", FormatOperator(a.Operator))
	}
	if math.Abs(a.Impact.CYukawaCandidate-0.9992248096922658) > 1e-15 || math.Abs(a.Impact.CHiggsCandidate-1.0372205108665145) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
}

func TestGate825StatusAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Status.Level, "R2+ candidate") || a.Status.CanUpdateCYukawa || a.Status.HasCertifiedProjectors || a.Status.HasBoundaryActivationMap || a.Status.HasSectorLedger || a.Status.NativeYukawaTheorem {
		t.Fatalf("bad status: %+v", a.Status)
	}
	res := Generation2RelativeRestMagnitudeOperatorAndBoundaryAlphaActivationMapAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
