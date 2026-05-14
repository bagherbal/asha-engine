package ewcartanledger

import "testing"

func TestGate254RetrievesNearbyLedgersButRejectsCarrierConflation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.Summary.Gate253DictionaryInherited || !a.Summary.RegistrySearchCompleted || !a.Summary.FockLedgersRetrieved {
		t.Fatalf("expected inherited dictionary and retrieved Fock ledgers: %+v", a.Summary)
	}
	if a.Summary.T3LNumberLedgerRetrieved || a.Summary.YPhiNumberLedgerRetrieved || a.Summary.T3LYPhiSO8Coordinates {
		t.Fatalf("physical EW coordinates must remain absent: %+v", a.Summary)
	}
	if !a.CarrierTyping.ConflationRejected || !a.CarrierTyping.MatterT3RNumberLedger {
		t.Fatalf("expected carrier separation with matter T3R diagnostic: %s", FormatCarrierTyping(a.CarrierTyping))
	}
	if a.Firewall.ConflatedT3RWithT3L || a.Firewall.ConflatedScalarYPhiWithFockY || a.Firewall.ImportedSMHyperchargeAsLedger {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}

func TestGate254WeakCartansAuditedButNotSelected(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if len(a.WeakCartans) != 6 {
		t.Fatalf("expected six two-mode Cartan candidates, got %d", len(a.WeakCartans))
	}
	if countSpatialU1(a.WeakCartans) != 3 {
		t.Fatalf("expected three pure-spatial U1-preserving candidates: %s", FormatWeakCartans(a.WeakCartans))
	}
	if anySelectedWeakCartan(a.WeakCartans) || a.Firewall.ForcedWeakPlane {
		t.Fatalf("weak plane was incorrectly selected: weak=%s firewall=%s", FormatWeakCartans(a.WeakCartans), FormatFirewall(a.Firewall))
	}
	if a.Kernel.Q8vCConstructed || a.Kernel.ThreePlaneDerived || a.Firewall.ForcedKernelDim3 {
		t.Fatalf("kernel firewall violated: kernel=%s firewall=%s", FormatKernel(a.Kernel), FormatFirewall(a.Firewall))
	}
}
