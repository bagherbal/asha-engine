package generation2syntheticevidenceboardadapter

import "testing"

func TestGate550SyntheticEvidenceBoardAdapter(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error = %v", err)
	}
	if !a.Inheritance.Gate549AirlockDefined || a.Inheritance.Gate549SchemaRows != 17 || !a.Inheritance.Gate549NoBoardEvidence || !a.Inheritance.Gate549NativeWriteLocked || !a.Inheritance.Gate549RedirectsToGate550 {
		t.Fatalf("Gate549 inheritance leaked: %+v", a.Inheritance)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 17 || a.Import.RejectedRows != 0 || len(a.Import.MissingRows) != 0 || len(a.Import.DuplicateRows) != 0 || !a.Import.ChecksumVerified {
		t.Fatalf("manifest parser/checksum failed: %+v", a.Import)
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllEvidenceBoardOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem || a.Import.AnyNativePromotion || a.Import.AnyNativeWrite || a.Import.AnyPhysicalClaim || a.Import.AnyBridgeEvidence || a.Import.AnyObservedClaim {
		t.Fatalf("metadata sieve failed: %+v", a.Import)
	}
	if !a.Board.CitationScopeParsed || !a.Board.UncertaintyBudgetParsed || !a.Board.ReproducibilityParsed || !a.Board.RevocationHooksParsed || !a.Board.VersionedIndexParsed || !a.Board.NativeDeltaZero || !a.Board.SyntheticUnderlyingEvidence || a.Board.AuthenticatedSourceChain || a.Board.AuthenticatedBridgeEvidence || a.Board.AcceptanceAllowed || a.Board.EvidenceEntriesAccepted != 0 || a.Board.BoardedAsBridgeEvidence || !a.Board.NativeWriteLocked || a.Board.NativeWriteAuthorization || a.Board.NativeRegistryWrite || !a.Board.BlockedBecauseSynthetic {
		t.Fatalf("board guard leaked: %+v", a.Board)
	}
	if a.Firewall.BridgeEvidenceBoarded || a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestGate550TheoremPasses(t *testing.T) {
	result := Generation2SyntheticEvidenceBoardAdapterDryRunTheorem().Run()
	if !result.Passed() {
		t.Fatalf("theorem failed:\n%s", result.Details())
	}
}
