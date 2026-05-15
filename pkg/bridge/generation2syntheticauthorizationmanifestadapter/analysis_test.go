package generation2syntheticauthorizationmanifestadapter

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "quarantine-only dry-run") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Authorization.DryRunAuthorizationArmed || a.Authorization.LiveComparatorAuthorization || a.Firewall.NativeRegistryWritten || a.Firewall.RealSchwingerSourceImported {
		t.Fatalf("Gate543 leaked: auth=%+v firewall=%+v", a.Authorization, a.Firewall)
	}
}

func TestManifestRowsAndChecksum(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Import.AcceptedRows != 14 || a.Import.RejectedRows != 0 || len(a.Import.MissingRows) != 0 || len(a.Import.DuplicateRows) != 0 {
		t.Fatalf("manifest row mismatch: %+v", a.Import)
	}
	if !a.Import.ChecksumVerified || a.Import.ChecksumExpected == "" || a.Import.ChecksumActual != a.Import.ChecksumExpected {
		t.Fatalf("checksum mismatch: %+v", a.Import)
	}
}

func TestMetadataSieve(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllComparatorOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || !a.Import.AllSourceTagged || !a.Import.AllConventionTagged {
		t.Fatalf("metadata not enforced: %+v", a.Import)
	}
	if a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyObservedClaim {
		t.Fatalf("metadata leaked native/physical claims: %+v", a.Import)
	}
}

func TestFirewallBlocksPhysicalPromotion(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Authorization.CanImportRealSource || a.Authorization.PhysicalSchwingerDerived || a.Authorization.OSPositivityProven || a.Authorization.WickRotationSelected || a.Authorization.HamiltonianDerived {
		t.Fatalf("authorization leaked physics: %+v", a.Authorization)
	}
	if a.Firewall.ComparatorExecutionPerformed || a.Firewall.LiveComparatorAuthorized || a.Firewall.NativeSchwingerFunctionWrite || a.Firewall.NativeHamiltonianWrite || a.Firewall.NativeTimeArrowWrite {
		t.Fatalf("firewall leaked native write: %+v", a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticComparatorAuthorizationManifestAdapterDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
