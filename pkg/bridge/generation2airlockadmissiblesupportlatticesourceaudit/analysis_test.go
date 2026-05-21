package generation2airlockadmissiblesupportlatticesourceaudit

import (
	"strings"
	"testing"
)

func TestGate931SupportLatticeSources(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != Gate930ShortStatus {
		t.Fatalf("bad inherited: %s", a.Inherited)
	}
	if !a.PunctureRoot.RootedAtPuncture || a.PunctureRoot.F0Rank != RankF0 || !a.PunctureRoot.RelativeTargets || a.PunctureRoot.NativeTheorem {
		t.Fatalf("bad puncture root: %s", FormatPunctureRoot(a.PunctureRoot))
	}
	if !a.SameSocket.CompletesSameSocket || !a.SameSocket.F1EqualsPhaseW || a.SameSocket.F1Rank != RankF1 || a.SameSocket.F1OverF0Rank != RankF1OverF0 || a.SameSocket.NativeTheorem {
		t.Fatalf("bad same socket: %s", FormatSameSocket(a.SameSocket))
	}
	if !a.TensorIntegrity.StructuredCompletions || a.TensorIntegrity.ArbitrarySubspaces || !a.TensorIntegrity.PreservesSocketW || a.TensorIntegrity.NativeTheorem {
		t.Fatalf("bad tensor integrity: %s", FormatTensorIntegrity(a.TensorIntegrity))
	}
	if !a.NoOrphan.ExcludesOppositeLeptonSingleton || !a.NoOrphan.ExcludesOppositeColorFragment || !a.NoOrphan.OppositeOnlyAtFullSaturation || a.NoOrphan.NativeTheorem {
		t.Fatalf("bad no orphan: %s", FormatNoOrphan(a.NoOrphan))
	}
}

func TestGate931SaturationZ2MeasureAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Saturation.FullPairForcesF2 || !a.Saturation.FullRectangle || !a.Saturation.RejectsStopAtF1 || a.Saturation.F2Rank != RankF2 || a.Saturation.F2OverF0Rank != RankF2OverF0 || a.Saturation.NativeTheorem {
		t.Fatalf("bad saturation: %s", FormatSaturation(a.Saturation))
	}
	if !a.MinimalChain.MinimalSufficient || !a.MinimalChain.NoExtraIntermediate || !a.MinimalChain.ThreeLevelCollapse || a.MinimalChain.NativeUniquenessProof {
		t.Fatalf("bad minimal chain: %s", FormatMinimalChain(a.MinimalChain))
	}
	if !a.Z2Lattice.DescendsToZ2Class || !a.Z2Lattice.PhaseFlipExchangesF0F1 || !a.Z2Lattice.PhaseFlipFixesF2 || !a.Z2Lattice.RanksRepresentativeFree || a.Z2Lattice.NativePhaseTheorem {
		t.Fatalf("bad z2 lattice: %s", FormatZ2Lattice(a.Z2Lattice))
	}
	if !a.Measure.ThetaRecovered || !a.Measure.ClosureDomainSourced || !a.Measure.MuBReconstructed || a.Measure.NativeAlpha || a.Measure.ThetaOneRank != RankF1OverF0 || a.Measure.ThetaTwoRank != RankF2OverF0 {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate931Theorem(t *testing.T) {
	res := Generation2AirlockAdmissibleSupportLatticeSourceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, AmbientRightSupportRectangle, CR2Decomposition, WDecomposition, AtomicCells, Z2PunctureClass, RepresentativePuncture, AdmissibleSupportChain, Z2AdmissibleSupportLattice, ClosureOperatorConsequence, ThetaFromSupportLattice, MuBFromSupportLattice, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
