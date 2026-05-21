package generation2socketmagnitudesourcebernoullibminusltransferaudit

import (
	"strings"
	"testing"
)

func TestGate865RequiredSocketMagnitudesSourceTyped(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !magnitudesOK(a.Transfer.Magnitudes) || !a.Transfer.DominantNormalization || !a.Transfer.RestBMinusLTransfer || !a.Transfer.BernoulliComplement || !a.Transfer.TripletQuadraticTransfer {
		t.Fatalf("bad magnitude source typing: %s", FormatTransfer(a.Transfer))
	}
}

func TestGate865TracePreservation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Trace.RestTrace, 3*AlphaB) || !near(a.Transfer.Magnitudes[1].Magnitude*3+a.Transfer.Magnitudes[2].Magnitude, 3*AlphaB) || !a.Transfer.TracePreserving {
		t.Fatalf("rest trace not preserved: %s | %s", FormatTrace(a.Trace), FormatTransfer(a.Transfer))
	}
}

func TestGate865SquareTraceReproduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	rc := AlphaB * (1 - AlphaB)
	rl := 3 * AlphaB * AlphaB
	expected := 3*rc*rc + rl*rl
	if !near(a.Trace.RestSquareTrace, expected) || !near(a.Trace.RestSquareTrace, 3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("bad square trace: %s", FormatTrace(a.Trace))
	}
}

func TestGate865NonCircularityFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Transfer.Native || a.Transfer.NonCircular || a.Obstruction.LayerB_AlphaDerived || a.Obstruction.SocketMagnitudeNative || a.R3.EligibleForR3 {
		t.Fatalf("overpromoted: %s | %s | %s", FormatTransfer(a.Transfer), FormatObstruction(a.Obstruction), FormatR3(a.R3))
	}
}

func TestGate865LedgerFrozen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate865Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate865Theorem(t *testing.T) {
	res := Generation2SocketMagnitudeSourceBernoulliBMinusLTransferAuditTheorem().Verify()
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
