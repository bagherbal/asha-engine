package generation2realschwingerimportswitchairlock

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "switch") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Switch.DefaultOff || a.Switch.RealSourceImportEnabled || a.Switch.ComparatorExecutionAllowed {
		t.Fatalf("switch is not fail-closed: %+v", a.Switch)
	}
}

func TestSwitchSchema(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Schema.Rows) != 12 || a.Schema.RequiredRows != 12 || a.Schema.DefaultOffRows != 12 || a.Schema.NativeWriteRows != 0 {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if !a.Schema.OperatorIntentRow || !a.Schema.SourceURIRow || !a.Schema.ChecksumProofRow || !a.Schema.NativeWriteLockRow {
		t.Fatalf("missing essential switch row: %+v", a.Schema)
	}
}

func TestNoComparatorOrNativeWrites(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Guard.ComparatorExecutionPerformed || a.Firewall.NativeRegistryWritten || a.Firewall.RealSchwingerSourceImported || a.Firewall.NativeHamiltonianWrite {
		t.Fatalf("firewall failed: guard=%+v firewall=%+v", a.Guard, a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2RealSchwingerSourceImportSwitchAirlockPreflightTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
