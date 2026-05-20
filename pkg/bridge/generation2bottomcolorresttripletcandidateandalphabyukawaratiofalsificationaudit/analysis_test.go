package generation2bottomcolorresttripletcandidateandalphabyukawaratiofalsificationaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate820TripletRatios(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Ledger.BOverT-0.0003877453837799576) > 1e-18 {
		t.Fatalf("bad B/T: %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.SqrtBOverT-0.019691251452864992) > 1e-16 {
		t.Fatalf("bad sqrt(B/T): %s", FormatLedger(a.Ledger))
	}
	if math.Abs(a.Ledger.DustOverT-4.513895642851888e-7) > 1e-19 || math.Abs(a.Ledger.SqrtDustOverT-0.0006718553149936292) > 1e-16 {
		t.Fatalf("bad dust scale: %s", FormatLedger(a.Ledger))
	}
}

func TestGate820CandidatesProtocolAndStatus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Candidates) != 4 || a.Candidates[0].AllowedToIdentifyNow || a.Candidates[1].AllowedToIdentifyNow || !a.Candidates[3].AllowedToIdentifyNow {
		t.Fatalf("bad candidate table: %s", FormatCandidates(a.Candidates))
	}
	if !a.Protocol.CanFalsify || !a.Protocol.CanUpgradeExternalR3 || !strings.Contains(strings.Join(a.Protocol.Tests, " "), "D_ext/T") {
		t.Fatalf("bad protocol: %s", FormatProtocol(a.Protocol))
	}
	if a.Status.NativeSourceFound || a.Status.ExternalLedgerSupplied || a.Status.CanUpdateCYukawa || !strings.Contains(a.Status.Level, "strengthened partial R2") {
		t.Fatalf("bad status: %+v", a.Status)
	}
}

func TestGate820TheoremAndVerdicts(t *testing.T) {
	res := Generation2BottomColorRestTripletCandidateAndAlphaBYukawaRatioFalsificationAuditTheorem().Verify()
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
