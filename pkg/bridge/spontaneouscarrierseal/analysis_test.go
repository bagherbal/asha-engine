package spontaneouscarrierseal

import "testing"

func TestGate256SpontaneousCarrierSealAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.Gate255NoGoInherited || !a.Summary.SpontaneousSealRecorded || !a.Summary.ConditionalIntertwinerSchema {
		t.Fatalf("expected inherited no-go and recorded seal schema: %s", FormatSummary(a.Summary))
	}
	if a.Seal.DerivedFromFiniteGeometry || a.Seal.OverridesNativeNoGo || a.Seal.PollutesFiniteCore {
		t.Fatalf("seal must remain quarantined and non-derived: %s", FormatSeal(a.Seal))
	}
	if a.Intertwiner.OperationalIntertwinerBuilt || a.Intertwiner.ProvidedDataCount != 0 || a.Intertwiner.AllRequiredDataProvided {
		t.Fatalf("Gate 256 must not invent sealed embedding values: %s", FormatIntertwiner(a.Intertwiner))
	}
	if !a.UnifiedLedger.SymbolicSchemaConstructed || a.UnifiedLedger.OperationalUnifiedLedgerBuilt || a.UnifiedLedger.ConcreteT3LNumberLedgerAvailable || a.UnifiedLedger.ConcreteYPhiNumberLedgerAvailable || a.UnifiedLedger.ConcreteQNumberLedgerAvailable {
		t.Fatalf("expected symbolic-only ledger schema: %s", FormatUnifiedLedger(a.UnifiedLedger))
	}
	if !a.SO8.SymbolicSchemaAvailable || a.SO8.ConcreteT3LSO8 || a.SO8.ConcreteYPhiSO8 || a.SO8.ConcreteQSO8 {
		t.Fatalf("expected symbolic-only so(8) schema: %s", FormatSO8(a.SO8))
	}
	if a.TrialityKernel.PhysicalBranchSelected || a.TrialityKernel.Q8vCConstructed || a.TrialityKernel.KernelDimensionKnown || a.TrialityKernel.NeutralThreePlaneDerived {
		t.Fatalf("triality/kernel must remain blocked: %s", FormatTrialityKernel(a.TrialityKernel))
	}
	if a.Firewall.InventedEmbeddingValues || a.Firewall.ImportedSMHyperchargeConvention || a.Firewall.ForcedWeakPlane || a.Firewall.SelectedTrialityByKernel || a.Firewall.PollutedFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
