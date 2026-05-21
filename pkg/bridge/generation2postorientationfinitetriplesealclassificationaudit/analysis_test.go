package generation2postorientationfinitetriplesealclassificationaudit

import (
	"strings"
	"testing"
)

func TestGate863LayerClassification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !layersOK(a.Layers) {
		t.Fatalf("bad layer classification: %s", FormatLayers(a.Layers))
	}
}

func TestGate863CarrierFork(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Carrier.AmbientPartRank != 16 || a.Carrier.AmbientFullRank != 32 || a.Carrier.MinimalPartRank != 15 || a.Carrier.MinimalFullRank != 30 || !a.Carrier.MinimalBranchSealed || a.Carrier.AmbientBranchNative {
		t.Fatalf("bad carrier fork: %s", FormatCarrier(a.Carrier))
	}
}

func TestGate863EdgeOperatorSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Edge.ColorCentral || !a.Edge.ScalarSocket || !a.Edge.CharacterMatchedBySeal || !a.Edge.FirstOrderConditional || a.Edge.NumericalYukawa || a.Edge.NativeFiniteTriple || a.Edge.YRank != 7 || a.Edge.DSymRank != 14 || a.Edge.KernelRank != 1 {
		t.Fatalf("bad edge classification: %s", FormatEdge(a.Edge))
	}
}

func TestGate863R3BlockedNextWoundYdaggerY(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.R3.FiniteBodySealPresent || a.R3.SectorTraceLedgerPresent || a.R3.TraceMagnitudeReadoutPresent || !a.R3.YDaggerYShapeCandidate || a.R3.YSocketMagnitudesDerived || a.R3.AlphaNative || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		t.Fatalf("R3 overpromoted or next wound missing: %s", FormatR3(a.R3))
	}
}

func TestGate863LedgerFrozen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate863Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate863Theorem(t *testing.T) {
	res := Generation2PostOrientationFiniteTripleSealClassificationAuditTheorem().Verify()
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
