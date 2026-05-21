package generation2conditionalyukawatraceproxyledgerofficialfreezeaudit

import (
	"strings"
	"testing"
)

func TestGate874ConditionalProxyChain(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Chain.CoherentGivenSeal || !a.Chain.ConditionalYukawaLike || a.Chain.AlphaNative {
		t.Fatalf("bad chain: %s", FormatChain(a.Chain))
	}
	if !containsAll(a.Chain.Failures, []string{FailureAlphaStillSealed, FailureNoNativeTargetSelection, FailureNoNativeSocketMagnitudeSource}) {
		t.Fatalf("missing chain failures: %s", FormatChain(a.Chain))
	}
}

func TestGate874DiagnosticLedgerSeparated(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.CYukawaMatchesThreeOverNEff || !a.Ledger.OfficialFrozen {
		t.Fatalf("bad diagnostic ledger: %s", FormatLedger(a.Ledger))
	}
	if a.Ledger.OperatorEqualsOfficialNEff || a.Ledger.OperatorEqualsOfficialCYukawa || a.Ledger.OperatorEqualsOfficialCHiggs {
		t.Fatalf("operator and official ledgers collapsed: %s", FormatLedger(a.Ledger))
	}
}

func TestGate874PromotionRequirementsBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.PromotionRequirements.AllRequirementsMet || a.PromotionRequirements.EligibleForR3 || a.PromotionRequirements.EligibleForR4 {
		t.Fatalf("promotion leaked: %s", FormatRequirements(a.PromotionRequirements))
	}
	if !containsAll(a.PromotionRequirements.Failures, []string{FailureNoNativeTargetSelection, FailureNoCrossLaneExclusion, FailureNoNativeSectorTraceMagnitude}) {
		t.Fatalf("missing promotion failures: %s", FormatRequirements(a.PromotionRequirements))
	}
}

func TestGate874OfficialFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("freeze violated: %s", FormatImpact(a.Impact))
	}
}

func TestGate874Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate874Theorem(t *testing.T) {
	res := Generation2ConditionalYukawaTraceProxyLedgerOfficialFreezeAuditTheorem().Verify()
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
