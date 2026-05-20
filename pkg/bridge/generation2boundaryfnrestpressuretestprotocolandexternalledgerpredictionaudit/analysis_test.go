package generation2boundaryfnrestpressuretestprotocolandexternalledgerpredictionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate815FrozenHypothesisAndAggregateDiagnostics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate814Inherited || !a.Hypothesis.Frozen {
		t.Fatalf("bad inheritance/hypothesis: %+v %+v", a.Inheritance, a.Hypothesis)
	}
	if math.Abs(a.Inheritance.DeltaNBFN-0.002327375081808316) > 1e-18 {
		t.Fatalf("bad DeltaN_BFN: %s", FormatInheritance(a.Inheritance))
	}
	if math.Abs(a.Aggregate.C2Observed-5.82999157225) > 1e-9 {
		t.Fatalf("bad c2 observed: %s", FormatAggregate(a.Aggregate))
	}
	if !containsAll(a.Hypothesis.Failures, []string{StatusNoRetuning, StatusTypedNotNative}) {
		t.Fatalf("missing frozen failures: %+v", a.Hypothesis.Failures)
	}
}

func TestGate815TopRestBandSpurionAndImpact(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.TopRest.AlphaMinOverS-0.300067468178) > 1e-10 {
		t.Fatalf("bad alpha min over s: %s", FormatTopRest(a.TopRest))
	}
	if math.Abs(a.TopRest.AlphaMaxOverS-0.30024229794) > 1e-9 {
		t.Fatalf("bad alpha max over s: %s", FormatTopRest(a.TopRest))
	}
	if math.Abs(a.Spurion.EpsilonBFN-0.2196426096400638) > 1e-15 {
		t.Fatalf("bad epsilon BFN: %+v", a.Spurion)
	}
	if math.Abs(a.Impact.CYukawaBFN-0.9992248096922658) > 1e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	if math.Abs(a.Impact.OfficialCYukawa-CYukawa) > 1e-15 {
		t.Fatalf("official ledger changed: %s", FormatImpact(a.Impact))
	}
}

func TestGate815TheoremStatusesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate815 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2BoundaryFNRestPressureTestProtocolAndExternalLedgerPredictionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
