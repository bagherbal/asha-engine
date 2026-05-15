package generation2oswickhilbertsectorclosureledger

import (
	"strings"
	"testing"
)

func TestBuildDefaultSectorClosureLedger(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.FrontierConsistent || a.Ledger.ClosedRows != 8 || a.Ledger.FailedRoutes != 9 {
		t.Fatalf("bad frontier ledger: %+v", a.Ledger)
	}
	if !a.Firewall.MatrixComplete || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeWickWrite || a.Firewall.NativeTimeArrowWrite || a.Firewall.NativeInternalGaugeWrite {
		t.Fatalf("closure ledger promoted native dynamics: %+v", a.Firewall)
	}
	if !strings.Contains(a.Truth, "Schwinger functions") || !strings.Contains(a.Firewall.Verdict, StatusFailedClosureNotHamiltonian) {
		t.Fatalf("missing final frontier language: truth=%s firewall=%s", a.Truth, a.Firewall.Verdict)
	}
}

func TestFrontierRowsCarryAllThreeLanes(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range a.Ledger.Rows {
		if row.Native == "" || row.BridgeSocket == "" || row.Environmental == "" || row.FailedRoute == "" || !row.Closed {
			t.Fatalf("incomplete frontier row: %+v", row)
		}
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2OSWickHilbertSectorClosureLedgerFrontierMapTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
