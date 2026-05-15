package generation2realsourcecomparatorauthorizationairlock

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "bridge-quarantine-only") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if a.Authorization.ComparatorExecutionPerformed || a.Firewall.NativeRegistryWritten || a.Firewall.RealSchwingerSourceImported {
		t.Fatalf("authorization airlock leaked: auth=%+v firewall=%+v", a.Authorization, a.Firewall)
	}
}

func TestManifestRows(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Schema.Rows) != 14 || a.Schema.RequiredRows != 14 || a.Schema.NativeWriteRows != 0 || !a.Schema.NativeWriteLockRow || !a.Schema.HumanReviewRow {
		t.Fatalf("manifest schema incomplete: %+v", a.Schema)
	}
	if !a.Authorization.BridgeQuarantineOnly || !a.Authorization.NativeWriteLocked || a.Authorization.ManifestImported {
		t.Fatalf("authorization state not fail-closed: %+v", a.Authorization)
	}
}

func TestFirewallBlocksPhysicalPromotion(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Guard.ComparatorCanRunInPreflight || a.Guard.OSPositivityProven || a.Guard.WickRotationSelected || a.Guard.HamiltonianDerived {
		t.Fatalf("guard leaked physics: %+v", a.Guard)
	}
	if a.Firewall.ComparatorExecutionPerformed || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeTimeArrowWrite {
		t.Fatalf("firewall leaked native write: %+v", a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2RealSourceComparatorAuthorizationManifestAirlockTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
