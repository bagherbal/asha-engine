package generation2oneplusthreerestsimplexsourceminimalityandexternalledgerfalsificationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate819InheritedSimplexLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.DeltaNBFN-6*a.Ledger.AlphaB) > 1e-18 {
		t.Fatalf("Delta_BFN != 6 alpha_B: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.QSimplex-0.3330749367196054) > 1e-15 || math.Abs(a.Ledger.NEffSimplex-a.Ledger.NEffBFN) > 3e-16 {
		t.Fatalf("bad simplex inheritance: %s", FormatLedger(a.Ledger))
	}
}

func TestGate819SourceLanesAndProtocol(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.SourceSeal.CurrentSupplied || len(a.SourceSeal.Components) != 10 {
		t.Fatalf("source seal should be defined but not supplied: %s", FormatSourceSeal(a.SourceSeal))
	}
	if len(a.Candidates) != 4 || a.Candidates[0].SuppliesTraceReadout || a.Candidates[1].SuppliesTraceReadout || !a.Candidates[2].SuppliesAlpha || !a.Candidates[3].SuppliesSectorAtoms {
		t.Fatalf("bad candidate audit: %s", FormatCandidates(a.Candidates))
	}
	if !a.Protocol.CanUpgradeExternalR3 || len(a.Protocol.PrimaryTests) != 5 || !strings.Contains(strings.Join(a.Protocol.PrimaryTests, " "), "alpha_ext") {
		t.Fatalf("bad falsification protocol: %s", FormatProtocol(a.Protocol))
	}
}

func TestGate819StatusImpactAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Status.Level, "strengthened partial R2") || !a.Status.ExternalR3Ready || a.Status.NativeSourceFound || a.Status.CanUpdateCYukawa {
		t.Fatalf("bad status: %+v", a.Status)
	}
	if math.Abs(a.Impact.CYukawaCandidate-0.9992248096922658) > 1e-15 || math.Abs(a.Impact.CHiggsCandidate-1.0372205108665146) > 2e-15 {
		t.Fatalf("bad impact: %s", FormatImpact(a.Impact))
	}
	res := Generation2OnePlusThreeRestSimplexSourceMinimalityAndExternalLedgerFalsificationAuditTheorem().Verify()
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
