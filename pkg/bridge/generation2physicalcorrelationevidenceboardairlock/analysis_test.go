package generation2physicalcorrelationevidenceboardairlock

import "testing"

func TestGate549PhysicalCorrelationEvidenceBoardAirlock(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Inheritance.Gate548ClosureEmitted || a.Inheritance.Gate548RowsClosed != 12 || !a.Inheritance.Gate548NoRealSource || !a.Inheritance.Gate548NoBridgeEvidence || !a.Inheritance.Gate548NativeWriteLocked || !a.Inheritance.Gate548RedirectsToGate549 {
		t.Fatalf("Gate548 inheritance leaked: %+v", a.Inheritance)
	}
	if a.Schema.RowCount != 17 || !a.Schema.AllRowsRequired || !a.Schema.CitationScopeDefined || !a.Schema.UncertaintyBudgetRequired || !a.Schema.ReproducibilityRecordRequired || !a.Schema.EnvironmentalClassRequired || !a.Schema.RevocationHooksRequired || !a.Schema.NativeDeltaZeroRequired || !a.Schema.BridgeOnly || !a.Schema.NativePromotionRejected {
		t.Fatalf("evidence-board schema incomplete: %+v", a.Schema)
	}
	if a.State.ReleasedBridgeEvidenceAvailable || a.State.EvidenceBoardManifestImported || a.State.EvidenceEntriesAccepted != 0 || a.State.BoardedAsBridgeEvidence || a.State.NativeRegistryWrite || !a.State.NativeWriteLocked || !a.State.PrefightOnly {
		t.Fatalf("evidence-board state leaked: %+v", a.State)
	}
	if a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.PhysicalOSCertificateLoaded || a.Firewall.PhysicalWickMapLoaded || a.Firewall.PhysicalHilbertSpaceLoaded || a.Firewall.PhysicalHamiltonianLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.ReleasedBridgeEvidenceLoaded || a.Firewall.EvidenceBoardEntryWritten || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestGate549TheoremPasses(t *testing.T) {
	result := Generation2PhysicalCorrelationEvidenceBoardAirlockTheorem().Run()
	if !result.Passed() {
		t.Fatalf("theorem failed:\n%s", result.Details())
	}
}
