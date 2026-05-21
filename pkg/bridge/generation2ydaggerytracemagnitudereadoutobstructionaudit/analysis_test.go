package generation2ydaggerytracemagnitudereadoutobstructionaudit

import (
	"strings"
	"testing"
)

func TestGate864YDaggerYCarrierShape(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Readout.Positive || !a.Readout.CorrectActiveSupport || !a.Readout.PunctureAbsent || !a.Readout.LeftKernelExcluded || !a.Readout.CarrierWiseMatch {
		t.Fatalf("bad carrier shape: %s", FormatReadout(a.Readout))
	}
}

func TestGate864RequiredSocketMagnitudes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !weightsOK(a.Readout.Weights) {
		t.Fatalf("bad required socket magnitudes: %s", FormatWeights(a.Readout.Weights))
	}
	if !near(a.Readout.Weights[1].RequiredMagnitude, AlphaB*(1-AlphaB)) || !near(a.Readout.Weights[2].RequiredMagnitude, 3*AlphaB*AlphaB) {
		t.Fatalf("required alpha weights not matched: %s", FormatWeights(a.Readout.Weights))
	}
}

func TestGate864ConditionalTraceReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !near(a.Readout.TraceIfMatched, 3+3*AlphaB) || !near(a.Readout.SquareTraceIfMatched, 3+3*AlphaB*AlphaB-6*AlphaB*AlphaB*AlphaB+12*AlphaB*AlphaB*AlphaB*AlphaB) {
		t.Fatalf("bad conditional trace: %s", FormatReadout(a.Readout))
	}
	if a.Readout.MagnitudeWiseMatch || a.Readout.SocketMagnitudesDerived || !a.Readout.RequiresInsertedSocketValues {
		t.Fatalf("readout overpromoted: %s", FormatReadout(a.Readout))
	}
}

func TestGate864ObstructionAndR3Block(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Obstruction.YSocketMagnitudeSourceMissing || !a.Obstruction.AlphaSourceMissing || a.Obstruction.TraceReadoutNative || a.Obstruction.NonCircularMagnitudes || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		t.Fatalf("obstruction/R3 malformed: %s | %s", FormatObstruction(a.Obstruction), FormatR3(a.R3))
	}
}

func TestGate864LedgerFrozen(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate864Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate864Theorem(t *testing.T) {
	res := Generation2YDaggerYTraceMagnitudeReadoutObstructionAuditTheorem().Verify()
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
