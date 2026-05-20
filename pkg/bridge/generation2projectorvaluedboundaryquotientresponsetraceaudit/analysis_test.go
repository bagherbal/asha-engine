package generation2projectorvaluedboundaryquotientresponsetraceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate683Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ResponseFiberInherited || !a.Inherited.HomNotNativeSubspace || a.Inherited.K7Dimension != 7 || a.Inherited.QBoundaryDimension != 1 || a.Inherited.H72Dimension != 72 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if a.Firewall.HomIsNativeSubspace || !a.Firewall.ProjectorRouteAllowed {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	if !a.Projector.ProjectorInEndH72 || !a.Projector.ResponseInEndH72 || a.Projector.ProjectorRank != 7 {
		t.Fatalf("bad projector response: %+v", a.Projector)
	}
	if math.Abs(a.Ordinary.Coefficient-7.0/72.0) > 1e-15 || math.Abs(a.Ordinary.Residual) > 1e-8 {
		t.Fatalf("bad ordinary trace: %+v", a.Ordinary)
	}
}

func TestOrdinaryTraceBeatsSignedAndAlternatives(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hodge.ActiveUsesOrdinary || !a.Hodge.SignedFailsActive {
		t.Fatalf("expected ordinary rank trace to be active and signed trace to fail: %+v", a.Hodge)
	}
	if a.Hodge.SignedTrace != 1 || math.Abs(a.Hodge.SignedCoefficient-1.0/72.0) > 1e-15 {
		t.Fatalf("bad signed trace: %+v", a.Hodge)
	}
	if a.Alternatives.BestName != "tau_global" {
		t.Fatalf("expected global H72 trace to be best typed response: %+v", a.Alternatives)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ProjectorValuedBoundaryQuotientResponseTraceAuditTheorem().Verify()
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
