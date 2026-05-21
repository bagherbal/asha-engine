package generation2boundaryalphaexteriorsealr3eligibilityaudit

import (
	"strings"
	"testing"
)

func TestGate873AlphaSealClassification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AlphaSeal.ShapeTyped || !a.AlphaSeal.RankSourcesTyped || a.AlphaSeal.Native || !near(a.AlphaSeal.Alpha, AlphaB) {
		t.Fatalf("bad alpha seal: %s", FormatAlphaSeal(a.AlphaSeal))
	}
	if !containsAll(a.AlphaSeal.Failures, []string{FailureAlphaStillSealed, FailureAlphaNotNativeWithoutTargetTheorem, FailureNoNativeTargetSelection}) {
		t.Fatalf("missing alpha failures: %s", FormatAlphaSeal(a.AlphaSeal))
	}
}

func TestGate873ConditionalChain(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Chain.CoherentGivenAlphaSeal || !a.Chain.YDagYReadoutReady || !a.Chain.HaggReconstructedGivenAlpha || a.Chain.AlphaNative {
		t.Fatalf("bad chain: %s", FormatChain(a.Chain))
	}
}

func TestGate873OperatorNEffFrozenLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Readout.NEffMatchesGate829 || a.Readout.OfficialEqualsOperator || !a.Readout.Conditional {
		t.Fatalf("bad readout: %s", FormatReadout(a.Readout))
	}
	if a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger update leaked: %s", FormatImpact(a.Impact))
	}
}

func TestGate873R3EligibilityBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.R3.EligibleForConditionalR3Candidate || a.R3.EligibleForOfficialR3 || a.R3.EligibleForR4 || a.Impact.CanPromoteToR3 {
		t.Fatalf("R3 overpromoted: %s | %s", FormatR3(a.R3), FormatImpact(a.Impact))
	}
}

func TestGate873Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate873Theorem(t *testing.T) {
	res := Generation2BoundaryAlphaExteriorSealR3EligibilityAuditTheorem().Verify()
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
