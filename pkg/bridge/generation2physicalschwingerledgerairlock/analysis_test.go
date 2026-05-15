package generation2physicalschwingerledgerairlock

import (
	"strings"
	"testing"
)

func TestBuildDefaultPhysicalSchwingerAirlock(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Schema.Executed || a.Schema.RequiredRows != 19 || a.Schema.NativeWriteRows != 0 || !a.Schema.ReflectionPositivityCertificateRequired {
		t.Fatalf("bad Schwinger schema: %+v", a.Schema)
	}
	if a.Guard.ComparatorExecutionPerformed || a.Guard.PhysicalSchwingerDerived || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected {
		t.Fatalf("preflight guard promoted physics: %+v", a.Guard)
	}
	if !strings.Contains(a.Firewall.Verdict, StatusFailedSchemaNotHamiltonian) || !strings.Contains(a.Truth, "Schwinger") {
		t.Fatalf("missing firewall/truth language: firewall=%s truth=%s", a.Firewall.Verdict, a.Truth)
	}
}

func TestSchemaRowsAreBridgeOnlyAndRequired(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	for _, row := range a.Schema.Rows {
		if row.Name == "" || !row.Required || !row.BridgeOnly || row.NativeWrite || row.Reason == "" {
			t.Fatalf("bad schema row: %+v", row)
		}
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2PhysicalSchwingerFunctionSourceLedgerAirlockTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
