package generation2booleanoctonionicintersectionsupportprojectorselectionaudit

import (
	"strings"
	"testing"
)

func TestGate685Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.RankDegeneracyInherited || !a.Inherited.OrdinaryTraceRankOnly || !a.Inherited.RankSevenSelected {
		t.Fatalf("bad Gate684 inheritance: %+v", a.Inherited)
	}
	if !a.Chamber.DimensionalLedgerOK || a.Chamber.UPlusVDim != 63 || a.Chamber.OrthogonalW7Dim != 7 {
		t.Fatalf("bad chamber dimension ledger: %+v", a.Chamber)
	}
	if !a.Support.ImpliesImageInIntersection || a.Support.IntersectionDimension != 7 {
		t.Fatalf("support constraints should imply K7 image containment: %+v", a.Support)
	}
}

func TestSupportSelectionProof(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Selection.ImageSubsetK7 || !a.Selection.ImageEqualsK7 || !a.Selection.OrthogonalProjectorUnique {
		t.Fatalf("rank plus support should select unique P_K7: %+v", a.Selection)
	}
	if a.Selection.SelectedProjector != "P_K7" {
		t.Fatalf("wrong selected projector: %+v", a.Selection)
	}
}

func TestCandidateSupportRejection(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidates.PK7Passes || !a.Candidates.AllPassingArePK7 {
		t.Fatalf("P_K7 should be the only passing candidate: %+v", a.Candidates)
	}
	for _, want := range []string{"P_W7", "P_Uonly7", "P_Vonly7", "P_mixed_K7_W7", "P_boundary_mixed"} {
		if !containsCandidate(a.Candidates.RejectedRankSeven, want) {
			t.Fatalf("expected %s to be rejected: %+v", want, a.Candidates.RejectedRankSeven)
		}
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BooleanOctonionicIntersectionSupportProjectorSelectionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
