package generation2unsealedleptonwallpmnsrowbranchselectoraudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate602Enumeration(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.WallCandidates) != 18 {
		t.Fatalf("expected 18 wall candidates, got %d", len(a.WallCandidates))
	}
	if len(a.PMNSOverlaps) != 9 {
		t.Fatalf("expected 9 PMNS overlaps, got %d", len(a.PMNSOverlaps))
	}
	if len(a.CKMSigns) != 2 {
		t.Fatalf("expected 2 CKM signs, got %d", len(a.CKMSigns))
	}
	for _, row := range a.PMNSOverlaps {
		if row.Alpha == "e" && row.Index == 3 && math.Abs(row.Li-0.0055375) > 1e-15 {
			t.Fatalf("bad electron reactor overlap: %+v", row)
		}
	}
}

func TestGate602Selector(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.BalanceTable) != 108 {
		t.Fatalf("expected 108 branch-row rows, got %d", len(a.BalanceTable))
	}
	best := a.BalanceTable[0]
	if best.Alpha != "e" || best.NeutrinoI != 3 || best.CKMSign != +1 {
		t.Fatalf("best should select electron row, P3, +J: %+v", best)
	}
	if math.Abs(best.AbsBFlav-2.77587313788925e-06) > 1e-15 {
		t.Fatalf("unexpected best residual %.18g", best.AbsBFlav)
	}
	if a.ObservedRank.Rank != 1 || a.ObservedRank.MinimalClassSize != 6 || a.ObservedRank.Unique {
		t.Fatalf("observed tuple should be in sixfold minimal class: %+v", a.ObservedRank)
	}
	if !a.Degeneracy.ElectronRowSelected || !a.Degeneracy.P3Selected || !a.Degeneracy.PositiveJSelected || !a.Degeneracy.SigmaStillDegenerate {
		t.Fatalf("bad degeneracy ledger: %+v", a.Degeneracy)
	}
	if a.SelectorVerdict.SelectsFullChargedLeptonSigma || a.SelectorVerdict.UniqueSelector || a.SelectorVerdict.NativeSelector {
		t.Fatalf("selector overclaims: %+v", a.SelectorVerdict)
	}
}

func TestGate602TheoremAndFirewalls(t *testing.T) {
	res := Generation2UnsealedLeptonWallPMNSRowBranchSelectorAuditTheorem().Verify()
	if !res.Passed() {
		t.Fatalf("theorem checks failed: %+v", res)
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusBranchRowBalanceDefined, StatusWallRowsAndSignsEnumerated, StatusSelectsElectronRow, StatusSelectsP3AndPositiveJ, StatusNotFullOrderingSelector, StatusNoNativeBranchSelectionTheorem, StatusGate600Boundary, StatusGate602Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
