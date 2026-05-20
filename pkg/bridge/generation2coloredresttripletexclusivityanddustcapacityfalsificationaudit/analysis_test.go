package generation2coloredresttripletexclusivityanddustcapacityfalsificationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate821DustCapacityNumbers(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.TotalRestOverT-0.001163687540904158) > 1e-18 {
		t.Fatalf("bad total rest: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.TripletTraceOverT-0.001163236151339873) > 1e-18 {
		t.Fatalf("bad triplet trace: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.DustOverT-4.513895642851889e-7) > 1e-19 {
		t.Fatalf("bad dust capacity: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.SecondColoredPerColorBound-1.5046318809506294e-7) > 1e-20 || math.Abs(a.Ledger.SecondColoredSqrtBound-0.0003878958469680527) > 1e-18 {
		t.Fatalf("bad second colored bound: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.UncoloredAtomBound-4.513895642851889e-7) > 1e-19 || math.Abs(a.Ledger.UncoloredSqrtBound-0.0006718553149936293) > 1e-16 {
		t.Fatalf("bad uncolored bound: %s", FormatLedger(a.Ledger))
	}
}

func TestGate821CapacityHelpers(t *testing.T) {
	alpha := AlphaB(SBoundary)
	largest, ok := CheckColoredDustCapacity([]float64{BOverT(alpha), alpha * alpha * 0.5, alpha * alpha * 0.1}, alpha)
	if !ok || math.Abs(largest-BOverT(alpha)) > 1e-18 {
		t.Fatalf("expected colored capacity pass largest=%g ok=%t", largest, ok)
	}
	_, ok = CheckColoredDustCapacity([]float64{BOverT(alpha), alpha * alpha * 1.01}, alpha)
	if ok {
		t.Fatalf("expected colored capacity failure")
	}
	if !CheckUncoloredDustCapacity([]float64{DustOverT(alpha) * 0.9}, alpha) {
		t.Fatalf("expected uncolored capacity pass")
	}
	if CheckUncoloredDustCapacity([]float64{DustOverT(alpha) * 1.01}, alpha) {
		t.Fatalf("expected uncolored capacity failure")
	}
}

func TestGate821ProtocolStatusAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Branches) != 4 || !strings.Contains(FormatBranches(a.Branches), "bottom-color") || !strings.Contains(FormatBranches(a.Branches), "failure branch") {
		t.Fatalf("bad branches: %s", FormatBranches(a.Branches))
	}
	if !a.Protocol.CanFalsify || !strings.Contains(FormatProtocol(a.Protocol), "dust") || !strings.Contains(FormatProtocol(a.Protocol), "R_k <= alpha_B^2") {
		t.Fatalf("bad protocol: %s", FormatProtocol(a.Protocol))
	}
	if a.Status.NativeSourceFound || a.Status.ExternalLedgerSupplied || a.Status.CanUpdateCYukawa || !strings.Contains(a.Status.Level, "strengthened partial R2") {
		t.Fatalf("bad status: %+v", a.Status)
	}
	res := Generation2ColoredRestTripletExclusivityAndDustCapacityFalsificationAuditTheorem().Verify()
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
