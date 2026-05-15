package generation2reallookingschwingersourcenegativecontroladapter

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "default-deny") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Rejection.RejectedAsPhysicalSource || a.Rejection.ComparatorExecutionPerformed || a.Firewall.NativeRegistryWritten {
		t.Fatalf("negative control was not rejected cleanly: rejection=%+v firewall=%+v", a.Rejection, a.Firewall)
	}
}

func TestRealLookingFixtureParsesButDoesNotAuthenticate(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.RealLookingFixture || !a.Import.PhysicalSourceClaim || !a.Import.NonSyntheticClaim || a.Import.DeclaredSyntheticFixture {
		t.Fatalf("fixture did not present as real-looking negative control: %+v", a.Import)
	}
	if !a.Import.ChecksumVerified || !a.Import.RequiredSchemaRowsMatched || !a.Import.AllRowsNegativeControl {
		t.Fatalf("parser plumbing failed: %+v", a.Import)
	}
	if a.Rejection.PhysicalSourceAuthenticated || a.Rejection.PhysicalSourceImported {
		t.Fatalf("negative control authenticated/imported unexpectedly: %+v", a.Rejection)
	}
}

func TestDefaultSwitchBlocksComparator(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Rejection.SwitchOff || !a.Rejection.NoExplicitOperatorIntent || !a.Rejection.MissingLicenseOrAccessGrant || !a.Rejection.SourceURINotAuthenticated || !a.Rejection.InsufficientProvenance {
		t.Fatalf("missing required rejection reasons: %+v", a.Rejection)
	}
	if a.Rejection.ComparatorExecutionAllowed || a.Rejection.ComparatorExecutionPerformed || a.Firewall.ComparatorExecutionPerformed {
		t.Fatalf("comparator was not blocked: rejection=%+v firewall=%+v", a.Rejection, a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2RealLookingSchwingerSourceNegativeControlAdapterTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
