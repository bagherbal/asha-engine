package generation2firstorderjoppositecompatibilitycalculationaudit

import (
	"strings"
	"testing"
)

func TestGate852FirstOrderTargetTypedButNotExecutable(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrder.TypedAfterDataSeal || !a.FirstOrder.HasRhoSeal || !a.FirstOrder.HasJSeal || !a.FirstOrder.HasGammaSeal || !a.FirstOrder.HasDFSupport {
		t.Fatalf("first-order target not well typed: %s", FormatFirstOrder(a.FirstOrder))
	}
	if a.FirstOrder.ExecutableNow || a.FirstOrder.Certified || a.FirstOrder.HasOperatorRho || a.FirstOrder.HasOperatorJ || a.FirstOrder.HasOperatorDF {
		t.Fatalf("first-order overpromoted: %s", FormatFirstOrder(a.FirstOrder))
	}
}

func TestGate852WeakOrientationBlocker(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	w := a.WeakOrientation
	if !w.SplitDefinedAtSealLevel || !w.RequiresHiggsOrientation || !w.PrimaryFragilePoint {
		t.Fatalf("weak orientation blocker not isolated: %s", FormatWeakOrientation(w))
	}
	if w.StableUnderFullHAction || w.NativeHEigensplit {
		t.Fatalf("weak split overpromoted: %s", FormatWeakOrientation(w))
	}
	if !containsAll(w.Failures, []string{FailureWeakSocketSplitNotNative, FailureWeakOrientationNeedsHiggsSeal}) {
		t.Fatalf("missing weak-orientation failures: %s", FormatWeakOrientation(w))
	}
}

func TestGate852JOppositeAndKernelFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.JOpposite.JSealAvailable || a.JOpposite.OperatorLevelJ || a.JOpposite.OppositeCertified || a.JOpposite.CanBuildOppositeCommutator {
		t.Fatalf("J-opposite overpromoted: %s", FormatJOpposite(a.JOpposite))
	}
	if a.Kernel.StableUnderFullRhoJ || a.Kernel.PhysicalNeutrinoTheorem || a.Kernel.MasslessnessTheorem {
		t.Fatalf("kernel overpromoted: %s", FormatKernel(a.Kernel))
	}
	if !a.Kernel.StableUnderSchematicBlocks || !a.Kernel.KernelInsideMinimalCarrier || !a.Kernel.RightPunctureOutside {
		t.Fatalf("kernel support facts not preserved: %s", FormatKernel(a.Kernel))
	}
}

func TestGate852CarrierAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CarrierClosure.PreservedByBlockActionSeal || a.CarrierClosure.ClosedNatively || a.CarrierClosure.AbsentCellForcedBackBySchematicRhoF {
		t.Fatalf("carrier closure overpromoted: %s", FormatCarrierClosure(a.CarrierClosure))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger/R3 firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate852Theorem(t *testing.T) {
	res := Generation2FirstOrderJOppositeCompatibilityCalculationAuditTheorem().Verify()
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
