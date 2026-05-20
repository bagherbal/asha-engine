package generation2defecttodefecttraceoperatoraudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate677Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.BoundaryQuotientInherited || !a.Inherited.TraceActsOnBoundaryQuotient || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Domain.Dimension != 1 || !a.Domain.CanonicalFromGate676 || math.Abs(a.Domain.SSplit-0.0012924448188163) > 1e-15 {
		t.Fatalf("bad domain: %+v", a.Domain)
	}
	if a.Codomain.Dimension != 1 || math.Abs(a.Codomain.DBase-0.0001256552099684) > 1e-15 {
		t.Fatalf("bad codomain: %+v", a.Codomain)
	}
	if !a.Operator.Linear || !a.Operator.ScalarFunctionalOnly || a.Operator.RequiresVectorMap || math.Abs(a.Operator.TauDefect-7.0/72.0) > 1e-15 {
		t.Fatalf("bad operator: %+v", a.Operator)
	}
	if math.Abs(a.Test.Residual-8.52583439801e-10) > 1e-14 {
		t.Fatalf("bad residual: %+v", a.Test)
	}
	if a.Discipline.ClaimsNativeTraceCouplesDefects || a.Discipline.ClaimsNativeTraceOperator || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.ClaimsNativeWallAirlock || a.Discipline.ClaimsBoundaryStressDerivation || a.Discipline.ClaimsFullK7BoundaryMap || a.Discipline.Verdict != StatusGate677Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestOperatorIsScalarFunctionalNotVectorMap(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Operator.ScalarFunctionalOnly || a.Operator.RequiresVectorMap {
		t.Fatalf("operator should be scalar functional only: %+v", a.Operator)
	}
	pred := a.Operator.TauDefect * a.Domain.SSplit
	if math.Abs(pred-a.Test.PredictedDBase) > 1e-15 {
		t.Fatalf("prediction mismatch")
	}
	if math.Abs(a.Codomain.DBase-pred-a.Test.Residual) > 1e-15 {
		t.Fatalf("residual mismatch")
	}
}

func TestNonTautologyRequirements(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Requirements) != 5 {
		t.Fatalf("bad requirements length")
	}
	supplied := 0
	missing := 0
	for _, r := range a.Requirements {
		switch r.Status {
		case "supplied", "partially supplied":
			supplied++
		case "missing theorem":
			missing++
		}
	}
	if supplied != 4 || missing != 1 {
		t.Fatalf("unexpected requirement status counts: supplied=%d missing=%d reqs=%+v", supplied, missing, a.Requirements)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2DefectToDefectTraceCouplingOperatorAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
