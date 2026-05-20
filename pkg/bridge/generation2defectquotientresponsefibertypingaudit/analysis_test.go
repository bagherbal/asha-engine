package generation2defectquotientresponsefibertypingaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate682Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.PrimitiveDensityInherited || a.Inherited.K7Dimension != 7 || a.Inherited.QBoundaryDimension != 1 || a.Inherited.H72Dimension != 72 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Fiber.FiberDimension != 7 || !a.Fiber.IsomorphicSinceQDimOne || !strings.Contains(a.Fiber.DualForm, "Hom") {
		t.Fatalf("bad response fiber: %+v", a.Fiber)
	}
	if !a.ProductDensity.MatchesGate681Density || math.Abs(a.ProductDensity.Density-7.0/72.0) > 1e-15 {
		t.Fatalf("bad product density: %+v", a.ProductDensity)
	}
	if a.DirectTensor.FiberIsNativeSubspace || !a.DirectTensor.RequiresCouplingMap {
		t.Fatalf("direct/tensor firewall failed: %+v", a.DirectTensor)
	}
	if math.Abs(a.Action.Residual-a.Inherited.Residual) > 1e-18 || math.Abs(a.Action.Residual) > 1e-8 {
		t.Fatalf("bad action residual: %+v inherited=%g", a.Action, a.Inherited.Residual)
	}
	if a.Discipline.ClaimsResponseFiberTheorem || a.Discipline.ClaimsNativeSubspace || a.Discipline.ClaimsNativeSevenOver72 || a.Discipline.Verdict != StatusGate682Boundary {
		t.Fatalf("firewall breach: %+v", a.Discipline)
	}
}

func TestResponseFiberIsTypeGainNotNewSubspace(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Trace.SameNumericalDensity || a.Trace.BareProjectorRank != a.Trace.ResponseFiberRank {
		t.Fatalf("expected same rank density: %+v", a.Trace)
	}
	if a.DirectTensor.FiberIsNativeSubspace {
		t.Fatalf("response fiber should not be certified as native H72 subspace: %+v", a.DirectTensor)
	}
	if !strings.Contains(a.Missing.Verdict, StatusNoNativeResponseFiberCouplingMap) {
		t.Fatal("missing response-fiber coupling firewall")
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2DefectQuotientResponseFiberTypingAuditTheorem().Verify()
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
