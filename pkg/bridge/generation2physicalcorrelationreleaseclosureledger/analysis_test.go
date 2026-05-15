package generation2physicalcorrelationreleaseclosureledger

import "testing"

func TestGate548PhysicalCorrelationClosureLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Inheritance.Gate547SyntheticBlocked || !a.Inheritance.Gate547NoBridgeEvidence || !a.Inheritance.Gate547NoRealSource || !a.Inheritance.Gate547NativeWriteLocked {
		t.Fatalf("Gate547 inheritance leaked: %+v", a.Inheritance)
	}
	if a.Closure.RowCount != 12 || !a.Closure.NativeFrontierFrozen || !a.Closure.BridgeFrontierMapped || !a.Closure.EnvironmentalMapped {
		t.Fatalf("closure ledger incomplete: %+v", a.Closure)
	}
	if a.Guard.AuthenticatedNonSyntheticSource || a.Guard.ComparatorExecutedOnRealSource || a.Guard.BridgeEvidenceReleased || a.Guard.NativeRegistryWrite || !a.Guard.NativeWriteLocked || !a.Guard.ClosureOnly {
		t.Fatalf("sector guard leaked: %+v", a.Guard)
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidence || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestGate548TheoremPasses(t *testing.T) {
	result := Generation2PhysicalCorrelationImportReleaseSectorClosureLedgerTheorem().Run()
	if !result.Passed() {
		t.Fatalf("theorem failed:\n%s", result.Details())
	}
}
