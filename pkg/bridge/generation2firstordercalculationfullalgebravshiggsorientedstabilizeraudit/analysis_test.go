package generation2firstordercalculationfullalgebravshiggsorientedstabilizeraudit

import (
	"strings"
	"testing"
)

func TestGate855CarrierAndTarget(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Carrier.HLRank != 8 || a.Carrier.HRMinRank != 7 || a.Carrier.HPartMinRank != 15 || a.Carrier.HFMinRank != 30 {
		t.Fatalf("bad carrier: %s", FormatCarrier(a.Carrier))
	}
	if !a.Target.WellTyped || !a.Target.SupportExecutable || a.Target.OperatorTheoremExecutable || a.Target.Certified {
		t.Fatalf("bad first-order target flags: %s", FormatTarget(a.Target))
	}
}

func TestGate855FullAlgebraBranchBlocksOnWeakOrientation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Full.GenericHMixesHPlusHMinus || !a.Full.DRequiresOrientedWeakSockets || a.Full.OrientedSupportPreservedByFullAF || a.Full.FirstOrderSupportZero || a.Full.FirstOrderCertified {
		t.Fatalf("full branch overpromoted: %s", FormatFullBranch(a.Full))
	}
	if !containsAll(a.Full.Failures, []string{FailureFullHActionMixesWeakSockets, FailureFullAFBlockedByHiggsOrientation}) {
		t.Fatalf("missing weak-orientation failures: %+v", a.Full.Failures)
	}
}

func TestGate855StabilizerBranchSupportOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Stabilizer.PreservesHPlusHMinus || !a.Stabilizer.FirstOrderSupportCompatible || a.Stabilizer.FirstOrderOperatorTheorem || a.Stabilizer.FullUnbrokenAFTheorem {
		t.Fatalf("stabilizer branch invalid: %s", FormatStabilizer(a.Stabilizer))
	}
	if !a.Impact.PostOrientationSupportObject || !a.Impact.StabilizerSupportPass || a.Impact.FullAFPass || a.Impact.FirstOrderCertified || a.Impact.NativeFiniteTripleProof {
		t.Fatalf("impact overpromoted: %s", FormatImpact(a.Impact))
	}
}

func TestGate855KernelAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Kernel.LeftKernelStableFullAF || !a.Kernel.LeftKernelStableStabilizer || a.Kernel.PhysicalNeutrinoTheorem || a.Kernel.MasslessnessTheorem {
		t.Fatalf("kernel overpromoted: %s", FormatKernel(a.Kernel))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate855Theorem(t *testing.T) {
	res := Generation2FirstOrderCalculationFullAlgebraVsHiggsOrientedStabilizerAuditTheorem().Verify()
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
