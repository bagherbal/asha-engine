package brokenaction

import "testing"

func TestBrokenSectorActionSecondVariationSearch(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CandidatePositive {
		t.Fatalf("candidate diag(1,1,4) should be positive")
	}
	if !a.WhiteningExact {
		t.Fatalf("expected whitening diagnostic to be exact")
	}
	if a.Diag114SelectedByAction {
		t.Fatalf("Gate 95 must not mark diag(1,1,4) action-selected")
	}
	if a.SecondVariationComputed {
		t.Fatalf("second variation should remain open")
	}
	if len(a.ActionSlots) < 5 {
		t.Fatalf("expected finite action slots to be audited")
	}
}
