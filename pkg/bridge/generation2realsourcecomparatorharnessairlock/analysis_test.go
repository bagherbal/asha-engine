package generation2realsourcecomparatorharnessairlock

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "no comparator executes") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Guard.HarnessDefined || a.Guard.ComparatorExecutionPerformed || a.Firewall.NativeRegistryWritten || a.Firewall.RealSchwingerSourceImported {
		t.Fatalf("Gate544 leaked: guard=%+v firewall=%+v", a.Guard, a.Firewall)
	}
}

func TestSchemaRows(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Schema.RequiredRows != 16 || a.Schema.SourceRows != 7 || a.Schema.InputContractRows != 5 || a.Schema.OutputContractRows != 2 || a.Schema.QuarantineRows != 4 || a.Schema.NativeWriteLockRows != 1 {
		t.Fatalf("schema mismatch: %+v", a.Schema)
	}
}

func TestExecutionGuardBlocksComparator(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Guard.RealSourceLoaded || a.Guard.AuthorizationManifestLoaded || a.Guard.ComparatorExecutionAuthorized || a.Guard.ComparatorExecutionPerformed || a.Guard.DryRunComparatorExecution || a.Guard.LiveComparatorExecution {
		t.Fatalf("comparator guard leaked: %+v", a.Guard)
	}
	if !a.Guard.NativeWriteLocked || a.Guard.NativeWriteAuthorization || !a.Guard.AbortConditionsDefined || !a.Guard.AbortTriggeredByNoSource {
		t.Fatalf("guard did not fail closed: %+v", a.Guard)
	}
}

func TestFirewallBlocksNativePromotion(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeTimeArrowWrite {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2RealSourceComparatorExecutionHarnessAirlockPreflightTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
