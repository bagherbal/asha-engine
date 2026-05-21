package generation2weakdoublethiggsorientationsealaudit

import (
	"strings"
	"testing"
)

func TestGate853QuaternionicFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	q := a.QuaternionicFirewall
	if !q.FullModuleHStable || q.NativeRankOneEigensplit || q.GenericHActionPreservesLines {
		t.Fatalf("quaternionic firewall violated: %s", FormatQuaternionicFirewall(q))
	}
	if !containsAll(q.Failures, []string{FailureWeakSplitNotNativeHEigensplit, FailureRankOneLinesNotGloballyHStable}) {
		t.Fatalf("missing quaternionic firewall failures: %s", FormatQuaternionicFirewall(q))
	}
}

func TestGate853OrientationSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	o := a.OrientationSeal
	if !o.DefinedAtSealLevel || !o.ProjectorsComplete || !o.Orthogonal || o.NativeDerivation || !o.RequiresGaugeOrientation {
		t.Fatalf("orientation seal inconsistent: %s", FormatOrientationSeal(o))
	}
	if o.HPlusRank != 1 || o.HMinusRank != 1 {
		t.Fatalf("wrong weak socket ranks: %s", FormatOrientationSeal(o))
	}
}

func TestGate853EdgeRewriteAndKernel(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.EdgeRewrite.CompatibleWithGate847Skeleton || !a.EdgeRewrite.PunctureEdgeAbsent || !a.EdgeRewrite.LeptoColorPreserved || len(a.EdgeRewrite.Edges) != 3 {
		t.Fatalf("edge rewrite invalid: %s", FormatEdgeRewrite(a.EdgeRewrite))
	}
	if !a.Kernel.OrientationRelative || !a.Kernel.StableUnderOrientationBlocks || a.Kernel.StableUnderFullRhoJ || a.Kernel.PhysicalNeutrino || a.Kernel.Masslessness {
		t.Fatalf("kernel overpromoted: %s", FormatKernel(a.Kernel))
	}
}

func TestGate853FirstOrderAndLedgerFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FirstOrderPreparation.OrientationSealAvailable || a.FirstOrderPreparation.OperatorRealizationReady || a.FirstOrderPreparation.FirstOrderExecutable || a.FirstOrderPreparation.FirstOrderCertified {
		t.Fatalf("first-order overpromoted: %s", FormatFirstOrderPreparation(a.FirstOrderPreparation))
	}
	if !a.Ledger.OfficialFrozen || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("ledger firewall violated: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
}

func TestGate853Theorem(t *testing.T) {
	res := Generation2WeakDoubletHiggsOrientationSealAuditTheorem().Verify()
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
