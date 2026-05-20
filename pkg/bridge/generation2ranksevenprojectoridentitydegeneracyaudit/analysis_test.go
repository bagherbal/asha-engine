package generation2ranksevenprojectoridentitydegeneracyaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate684Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ProjectorResponseInherited || a.Inherited.K7Rank != 7 || a.Inherited.H72Dimension != 72 {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.RankLaw.DependsOnlyOnRank || a.RankLaw.CanSelectIdentity {
		t.Fatalf("rank law should be rank-only and identity-degenerate: %+v", a.RankLaw)
	}
	if a.Candidates.BestRank != 7 || math.Abs(a.Candidates.BestResidual) > 1e-8 {
		t.Fatalf("expected rank-seven best response: %+v", a.Candidates)
	}
}

func TestTraceDegeneracy(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !contains(a.Candidates.RankSevenCandidates, "P_K7") || !contains(a.Candidates.RankSevenCandidates, "P_W7") {
		t.Fatalf("expected P_K7 and P_W7 to share rank-seven trace response: %+v", a.Candidates.RankSevenCandidates)
	}
	if a.Degeneracy.PK7UniquelySelected || !a.Degeneracy.OrdinaryTraceRankOnly {
		t.Fatalf("ordinary trace must not select K7 identity: %+v", a.Degeneracy)
	}
	var pplus, pminus ProjectorCandidate
	for _, c := range a.Candidates.Candidates {
		switch c.Name {
		case "P_+":
			pplus = c
		case "P_-":
			pminus = c
		}
	}
	if math.Abs(pplus.Residual) < 1e-5 || math.Abs(pminus.Residual) < 1e-5 {
		t.Fatalf("Hodge blocks should not match active closure: P+ %+v P- %+v", pplus, pminus)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2RankSevenProjectorIdentityDegeneracyAuditTheorem().Verify()
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
