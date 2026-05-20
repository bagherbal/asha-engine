package generation2flavorbranchcompatibilityselectoraudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate601BranchEnumeration(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.LeptonBranches) != 6 {
		t.Fatalf("expected 6 lepton branches, got %d", len(a.LeptonBranches))
	}
	for _, b := range a.LeptonBranches {
		if !b.PositiveChamber || math.Abs(b.EpsilonDeg-2.26718003289167) > 1e-9 {
			t.Fatalf("unexpected branch epsilon/positivity: %+v", b)
		}
	}
	if len(a.PMNSOverlaps) != 3 || math.Abs(a.PMNSOverlaps[2].Li-0.0055375) > 1e-15 {
		t.Fatalf("bad PMNS overlaps: %+v", a.PMNSOverlaps)
	}
	if len(a.CKMSigns) != 2 || a.CKMSigns[0].Sign != 1 || a.CKMSigns[1].Sign != -1 {
		t.Fatalf("bad CKM signs: %+v", a.CKMSigns)
	}
}

func TestGate601BranchBalanceSelector(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.BalanceTable) != 36 {
		t.Fatalf("expected 36 branch rows, got %d", len(a.BalanceTable))
	}
	if a.ObservedRank.Rank != 1 || a.ObservedRank.MinimalClassSize != 6 || a.ObservedRank.Unique {
		t.Fatalf("observed branch should be in sixfold minimal class: %+v", a.ObservedRank)
	}
	if !a.SelectorVerdict.SelectsNeutrinoThirdProjector || !a.SelectorVerdict.SelectsPositiveCKMSign || a.SelectorVerdict.SelectsChargedLeptonOrdering || a.SelectorVerdict.UniqueBranchSelector {
		t.Fatalf("bad selector verdict: %+v", a.SelectorVerdict)
	}
	if a.Gap.GapToNextDistinct < 1e-5 {
		t.Fatalf("gap too small: %+v", a.Gap)
	}
}

func TestGate601TheoremAndFirewalls(t *testing.T) {
	res := Generation2FlavorBranchCompatibilitySelectorAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusBranchBalanceDefined, StatusBranchSpaceEnumerated, StatusObservedBranchHasMinimumResidual, StatusBalanceSelectsP3AndPositiveJ, StatusChargedLeptonPermutationDegeneracy, StatusBranchSelectorNotUnique, StatusNoNativeBranchSelectionTheorem, StatusGate600Boundary, StatusGate601Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
