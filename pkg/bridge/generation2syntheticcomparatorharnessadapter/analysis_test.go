package generation2syntheticcomparatorharnessadapter

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "quarantine output") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.DryRun.DryRunComparatorExecuted || !a.DryRun.QuarantineOutputWritten || a.Firewall.NativeRegistryWritten || a.Firewall.RealSchwingerSourceImported {
		t.Fatalf("Gate545 leaked: dryrun=%+v firewall=%+v", a.DryRun, a.Firewall)
	}
}

func TestSyntheticBundleRowsAndChecksum(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Import.AcceptedRows != 16 || a.Import.RejectedRows != 0 || len(a.Import.MissingRows) != 0 || len(a.Import.DuplicateRows) != 0 {
		t.Fatalf("row sieve mismatch: %+v", a.Import)
	}
	if !a.Import.ChecksumVerified || a.Import.ChecksumActual != a.Import.ChecksumExpected {
		t.Fatalf("checksum mismatch: %+v", a.Import)
	}
}

func TestDryRunQuarantineOnly(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if !a.DryRun.DryRunComparatorExecuted || a.DryRun.LiveComparatorExecuted || !a.DryRun.BridgeQuarantineOnly || !a.DryRun.QuarantineOutputWritten {
		t.Fatalf("dry-run state mismatch: %+v", a.DryRun)
	}
	if !a.DryRun.NativeWriteLocked || a.DryRun.NativeWriteAuthorization || !a.DryRun.AbortTriggered || !a.DryRun.RollbackTracePresent || !a.DryRun.HumanReviewRequired {
		t.Fatalf("quarantine lock mismatch: %+v", a.DryRun)
	}
}

func TestPhysicalClaimsRemainFalse(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Import.RealSource || a.Import.AuthenticatedReal || a.Import.ObservedLoaded || a.Import.MeasureLoaded || a.Import.OSCertLoaded || a.Import.WickMapLoaded || a.Import.HamiltonianLoaded || a.Import.NativeWrite {
		t.Fatalf("import leaked physical source: %+v", a.Import)
	}
	if a.DryRun.PhysicalOSProof || a.DryRun.PhysicalWickMap || a.DryRun.PhysicalHilbertSpace || a.DryRun.PhysicalHamiltonian || a.DryRun.PhysicalUnitaryDynamics || a.DryRun.PhysicalGlobalCausality || a.DryRun.PhysicalArrowOfTime {
		t.Fatalf("dry-run leaked physical claims: %+v", a.DryRun)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticComparatorHarnessResultAdapterDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
