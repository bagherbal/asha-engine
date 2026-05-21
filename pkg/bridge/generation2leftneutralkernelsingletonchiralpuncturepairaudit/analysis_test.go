package generation2leftneutralkernelsingletonchiralpuncturepairaudit

import (
	"strings"
	"testing"
)

func TestGate849LeftNeutralComplement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Image.RightDomainRank != HRMinRank || a.Image.LeftTargetRank != HLRank || !a.Image.FullSupportRank || a.Image.ImageRank != 7 || a.Image.ComplementRank != 1 {
		t.Fatalf("bad image rank anatomy: %s", FormatImage(a.Image))
	}
	if a.Image.LeftComplement.Name != "h_+ tensor P_1" || !a.Image.LeftComplement.Kernel || a.Image.LeftComplement.InYImage || a.Image.LeftComplement.PhysicalAssignment {
		t.Fatalf("bad left complement: %s", FormatCell(a.Image.LeftComplement))
	}
	if activeCellsRank(a.Image.ActiveTargetCells) != 7 || !activeCellsExclude(a.Image.ActiveTargetCells, "h_+ tensor P_1") {
		t.Fatalf("bad active image cells: %s", FormatImage(a.Image))
	}
}

func TestGate849SymbolicDiracKernel(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Kernel.SupportOnly || a.Kernel.NativeDFMatrix || a.Kernel.NumericalDFMatrix {
		t.Fatalf("kernel promoted incorrectly: %s", FormatKernel(a.Kernel))
	}
	if a.Kernel.YRank != 7 || a.Kernel.DFRank != 14 || a.Kernel.TotalRank != 15 || a.Kernel.KernelDim != 1 || a.Kernel.RightKernelDim != 0 || a.Kernel.LeftKernelDim != 1 {
		t.Fatalf("bad kernel dimensions: %s", FormatKernel(a.Kernel))
	}
	if a.Kernel.KernelSupport.Name != "h_+ tensor P_1" || !a.Kernel.KernelSupport.Kernel || a.Kernel.KernelSupport.PhysicalAssignment {
		t.Fatalf("bad kernel support: %s", FormatCell(a.Kernel.KernelSupport))
	}
}

func TestGate849ChiralNeutralPairFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Pair.PairCandidate || a.Pair.PhysicalParticleTheorem || !a.Pair.SameLeptonSupport || !a.Pair.SamePlusSocket || !a.Pair.DifferentChirality {
		t.Fatalf("bad chiral pair: %s", FormatPair(a.Pair))
	}
	if a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing {
		t.Fatalf("impact overpromoted: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NoPhysicalNeutrino || !a.Firewalls.NoRightNeutrino || !a.Firewalls.NoMasslessnessTheorem || !a.Firewalls.KernelNotYukawaMagnitude || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || a.Firewalls.Verdict != StatusFirewallGate849 {
		t.Fatalf("firewalls invalid: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate849Theorem(t *testing.T) {
	res := Generation2LeftNeutralKernelSingletonChiralPuncturePairAuditTheorem().Verify()
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
