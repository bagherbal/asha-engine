package hyperaudit

import "testing"

func TestHyperchargeTableAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ChiralOrientationSelected {
		t.Fatalf("expected hypercharge audit to select a chiral orientation")
	}
	if !a.Odd.MatchesRightSingletConjugateTable {
		t.Fatalf("expected odd branch to match right-singlet/conjugate hypercharge table, got %s", FormatCounts(a.Odd.HyperchargeCounts))
	}
	if a.Even.MatchesRightSingletConjugateTable {
		t.Fatalf("even branch should not match the right-singlet/conjugate table")
	}
	if a.FullStandardModelTableDerived {
		t.Fatalf("full left-handed Standard Model table should remain open")
	}
}
