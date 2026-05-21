package generation2r3sectorledgerrequirementsunderboundaryalphasealaudit

import (
	"strings"
	"testing"
)

func TestGate882SuppliesButNotR3(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Supplies.BoundaryAlphaSeal || !a.Supplies.YDaggerYPositiveReadout || !a.Supplies.DiagnosticOnly {
		t.Fatalf("bad supplies: %s", FormatSupplies(a.Supplies))
	}
	if a.Eligibility.NativeR3 || a.Eligibility.NativeR4 || a.Eligibility.Classification != R2Status {
		t.Fatalf("bad eligibility: %s", FormatEligibility(a.Eligibility))
	}
}

func TestGate882RequirementsAndBlockers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Requirements.NativeR3Ready || a.Requirements.NonCircularAlphaSource || a.Requirements.TypedSectorProjectors {
		t.Fatalf("requirements wrongly satisfied: %s", FormatRequirements(a.Requirements))
	}
	if len(a.Blockers) != 5 || a.Blockers[0].Name != BlockerBoundaryIncidenceFunctor || a.Blockers[1].Name != BlockerSectorTraceLedgerMap {
		t.Fatalf("bad blockers: %s", FormatBlockers(a.Blockers))
	}
}

func TestGate882OfficialFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Frozen || !a.Ledger.DiagnosticOnly || a.Ledger.CanUpdate {
		t.Fatalf("official freeze leaked: %s", FormatLedger(a.Ledger))
	}
	if near(a.Ledger.OperatorNEff, a.Ledger.OfficialNEff) || near(a.Ledger.OperatorCYukawa, a.Ledger.OfficialCYukawa) {
		t.Fatalf("operator and official ledger collapsed: %s", FormatLedger(a.Ledger))
	}
}

func TestGate882Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate882Theorem(t *testing.T) {
	res := Generation2R3SectorLedgerRequirementsUnderBoundaryAlphaSealAuditTheorem().Verify()
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
