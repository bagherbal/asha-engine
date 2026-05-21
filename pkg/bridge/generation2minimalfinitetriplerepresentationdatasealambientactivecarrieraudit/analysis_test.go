package generation2minimalfinitetriplerepresentationdatasealambientactivecarrieraudit

import (
	"strings"
	"testing"
)

func TestGate851CarrierFork(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Carrier.AmbientActiveSeparated || !a.Carrier.MinimalBranchSelected || !a.Carrier.ExtendedBranchKept {
		t.Fatalf("carrier fork not explicit: %s", FormatCarrier(a.Carrier))
	}
	if a.Carrier.HPartAmbientRank != 16 || a.Carrier.HFAmbientRank != 32 || a.Carrier.HPartMinRank != 15 || a.Carrier.HFMinRank != 30 {
		t.Fatalf("bad ambient/minimal ranks: %s", FormatCarrier(a.Carrier))
	}
	if a.Carrier.Puncture.Name != "e_+ tensor P_1" || !a.Carrier.Puncture.Absent || a.Carrier.Puncture.InMinimalCarrier || a.Carrier.Puncture.PhysicalName {
		t.Fatalf("puncture overpromoted or not outside minimal carrier: %s", FormatCell(a.Carrier.Puncture))
	}
}

func TestGate851RhoFSealClosureButNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RhoF.PreservesMinimalCarrier || !a.RhoF.AbsentCellClosureSafe || !a.RhoF.RightSocketsSealed || !a.RhoF.LeptoColorBlocksPreserved || !a.RhoF.WeakDoubletPreserved {
		t.Fatalf("rho_F seal does not preserve minimal support: %s", FormatRhoF(a.RhoF))
	}
	if a.RhoF.CompleteActionLedger || a.RhoF.NativeProof || a.RhoF.KernelStableUnderFullAction {
		t.Fatalf("rho_F overpromoted: %s", FormatRhoF(a.RhoF))
	}
	if !containsAll(a.RhoF.Failures, []string{FailureDataSealNotNativeFiniteTripleProof, FailureWeakSocketSplitNotNative, FailureKernelStabilityNotCertified}) {
		t.Fatalf("missing rho_F firewalls: %s", FormatRhoF(a.RhoF))
	}
}

func TestGate851DataSealPreparesButDoesNotProveFirstOrder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.ObjectsPrepared || !a.FirstOrder.HasRhoSeal || !a.FirstOrder.HasJSeal || !a.FirstOrder.HasGammaSeal || !a.FirstOrder.HasDFSymSeal {
		t.Fatalf("first-order objects not prepared: %s", FormatFirstOrder(a.FirstOrder))
	}
	if a.FirstOrder.CanCalculateFirstOrderNow || a.FirstOrder.NativeRhoF || a.FirstOrder.NativeJ || a.FirstOrder.NativeGamma || a.FirstOrder.NativeDF {
		t.Fatalf("first-order overpromoted: %s", FormatFirstOrder(a.FirstOrder))
	}
	if a.Impact.NativeFiniteTriple || a.Impact.FirstOrderCertified || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("impact overpromoted: %s", FormatImpact(a.Impact))
	}
}

func TestGate851DFSymAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DFSym.ExtendedToJCopy || !a.DFSym.SupportOnly || a.DFSym.OperatorValued || a.DFSym.YukawaMagnitudeSource {
		t.Fatalf("D_F sym overpromoted: %s", FormatDFSym(a.DFSym))
	}
	if !a.Ledger.OfficialFrozen || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger overpromoted: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate851Theorem(t *testing.T) {
	res := Generation2MinimalFiniteTripleRepresentationDataSealAmbientActiveCarrierAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
