package generation2r2plusplusr3firewallanddualtripletsectorledgerobstructionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate831InheritsAggregateOperatorButNotR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.R2PlusPlusConsolidated || a.Ledger.R3SectorLedgerCertified || !a.Ledger.AlphaSealed {
		t.Fatalf("bad ledger status: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.OperatorNEff-3.002327375081808) > 5e-16 {
		t.Fatalf("bad operator N_eff: %s", FormatLedger(a.Ledger))
	}
	if a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("over-promoted impact: %s", FormatImpact(a.Impact))
	}
}

func TestGate831DualTripletAndSevenFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DualTriplet.SameDimension || a.DualTriplet.TypedMapCertified || a.DualTriplet.Identified {
		t.Fatalf("dual triplet firewall failed: %s", FormatDualTriplet(a.DualTriplet))
	}
	if !containsAll(a.DualTriplet.Failures, []string{FailureColorTripletNotFockTriplet, FailureNoSectorTraceLedgerMap}) {
		t.Fatalf("missing dual-triplet failures: %s", strings.Join(a.DualTriplet.Failures, ","))
	}
	if !a.Seven.CountMatchesK7 || !a.Seven.ClassifiedAsResonanceOnly || a.Seven.ProjectorTheoremCertified || a.Seven.AggregateToK7MapCertified {
		t.Fatalf("seven firewall failed: %s", FormatSeven(a.Seven))
	}
	if !containsAll(a.Seven.Failures, []string{FailureSevenAtomsNotK7, FailureNoAggregateToK7Map}) {
		t.Fatalf("missing seven failures: %s", strings.Join(a.Seven.Failures, ","))
	}
}

func TestGate831SectorLedgerRequirementsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Requirements.RequirementsSatisfied || a.Requirements.SectorLedgerCertified {
		t.Fatalf("sector ledger over-promoted: %s", FormatRequirements(a.Requirements))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.DualTripletSeparated || !a.Firewalls.SevenNotK7 || !a.Firewalls.NoSectorLedgerMap || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2R2PlusPlusR3FirewallAndDualTripletSectorLedgerObstructionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem construction failure: %+v", res)
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
