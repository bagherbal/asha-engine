package generation2conditionalyukawatraceproxybranchclosurenextfrontieraudit

import (
	"strings"
	"testing"
)

func TestGate881BranchClosedNotR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Closure.BranchClosed || !a.Closure.ConditionalProxyMature || a.Closure.EligibleForNativeR3 || a.Closure.EligibleForR4 || a.Closure.Status != R2Status {
		t.Fatalf("bad closure: %s", FormatClosure(a.Closure))
	}
}

func TestGate881LedgerFrozen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || !a.Ledger.DiagnosticOnly || a.Ledger.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatLedger(a.Ledger))
	}
	if near(a.Ledger.OperatorNEff, a.Ledger.OfficialNEff) || near(a.Ledger.OperatorCYukawa, a.Ledger.OfficialCYukawa) {
		t.Fatalf("operator and official ledgers collapsed: %s", FormatLedger(a.Ledger))
	}
}

func TestGate881NativeWallFiled(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Wall.Name != MissingNativeTheoremName || a.Wall.Native || !a.Wall.RequiredForR3 || !a.Wall.BlocksOfficialLedger {
		t.Fatalf("bad native wall: %s", FormatWall(a.Wall))
	}
	if !containsAll(a.Wall.Failures, []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed}) {
		t.Fatalf("missing wall firewalls: %s", FormatWall(a.Wall))
	}
}

func TestGate881NextFrontierDecision(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Decision.RecommendedFrontier != RecommendedNextFrontier || !a.Decision.AvoidAlphaLoop || !a.Decision.AvoidPhysicalYukawaJump {
		t.Fatalf("bad decision: %s", FormatDecision(a.Decision))
	}
	if len(a.Decision.Frontiers) != 3 || !a.Decision.Frontiers[1].Recommended || a.Decision.Frontiers[0].Recommended || a.Decision.Frontiers[2].Recommended {
		t.Fatalf("bad frontier set: %s", FormatDecision(a.Decision))
	}
}

func TestGate881Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate881Theorem(t *testing.T) {
	res := Generation2ConditionalYukawaTraceProxyBranchClosureNextFrontierAuditTheorem().Verify()
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
