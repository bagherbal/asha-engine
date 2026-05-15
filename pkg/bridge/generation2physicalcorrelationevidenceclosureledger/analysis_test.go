package generation2physicalcorrelationevidenceclosureledger

import "testing"

func TestGate551PhysicalCorrelationEvidenceBoardClosureLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Inheritance.Gate550SyntheticBlocked || !a.Inheritance.Gate550NoBridgeEvidence || !a.Inheritance.Gate550NoRealSource || !a.Inheritance.Gate550NativeWriteLocked || !a.Inheritance.Gate550NativeDeltaZero {
		t.Fatalf("Gate550 inheritance leaked: %+v", a.Inheritance)
	}
	if a.Closure.RowCount != 15 || !a.Closure.NativeFrontierFrozen || !a.Closure.BridgeFrontierMapped || !a.Closure.EnvironmentalMapped || !a.Closure.EvidenceBoardClosed {
		t.Fatalf("closure ledger incomplete: %+v", a.Closure)
	}
	if a.Guard.AuthenticatedNonSyntheticSource || a.Guard.ComparatorExecutedOnRealSource || a.Guard.BridgeEvidenceReleased || a.Guard.EvidenceBoardEntryAccepted || a.Guard.EvidenceEntriesAccepted != 0 || a.Guard.NativeRegistryWrite || !a.Guard.NativeWriteLocked || !a.Guard.ClosureOnly {
		t.Fatalf("sector guard leaked: %+v", a.Guard)
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidence || a.Firewall.BoardedBridgeEvidence || a.Firewall.NativeEvidenceBoardWrite || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestGate551TheoremPasses(t *testing.T) {
	result := Generation2PhysicalCorrelationEvidenceBoardSectorClosureLedgerTheorem().Run()
	if !result.Passed() {
		t.Fatalf("theorem failed:\n%s", result.Details())
	}
}
