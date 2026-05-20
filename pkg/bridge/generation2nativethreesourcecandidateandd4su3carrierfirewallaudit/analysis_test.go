package generation2nativethreesourcecandidateandd4su3carrierfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate799InheritanceAndRequirements(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate798.Inherited || a.Gate798.CurrentCertifiedSource != "color-tripled top dominance" || !a.Gate798.NotGenerationTheorem || !a.Gate798.NotD4Theorem || !a.Gate798.NotNativeYukawaTheorem {
		t.Fatalf("bad inheritance: %+v", a.Gate798)
	}
	if !a.Requirement.Defined || !a.Requirement.RejectNoMap || !a.Requirement.RejectNoOps || !containsAll(a.Requirement.Fields, []string{"typed carrier", "trace/readout", "breaking", "noncircularity"}) {
		t.Fatalf("bad requirement: %s", FormatRequirement(a.Requirement))
	}
}

func TestGate799Candidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	color, ok := CandidateByName(a.Candidates, "Color SU(3)")
	if !ok || color.Rank != 1 || !strings.Contains(color.TypedSource, "trace") || !containsAll(color.Strengths, []string{"a_u=3", "N_eff_top=3"}) || !containsAll(color.Failures, []string{StatusColorNoEigenvalues, StatusColorNoGeneration}) {
		t.Fatalf("bad color candidate: %s", FormatCandidate(color))
	}
	d4, ok := CandidateByName(a.Candidates, "D4")
	if !ok || d4.Rank != 3 || !containsAll(d4.RequiredMap, []string{"D4TrialityCarrierPackage", "real-form", "trace-readout", "breaking"}) || !containsAll(d4.Failures, []string{StatusNoD4Package, StatusNoD4TraceMap, StatusCompactSpin8Firewall}) {
		t.Fatalf("bad D4 candidate: %s", FormatCandidate(d4))
	}
	a2, ok := CandidateByName(a.Candidates, "SU(3) / A2")
	if !ok || a2.Rank != 6 || !containsAll(a2.Failures, []string{StatusHexMotifNotEvidence, StatusColorNotFlavorSU3, StatusNoA2TraceMap}) {
		t.Fatalf("bad A2 candidate: %s", FormatCandidate(a2))
	}
	k7, ok := CandidateByName(a.Candidates, "K7")
	if !ok || k7.Rank != 7 || !containsAll(k7.Failures, []string{StatusK7MinusNotGeneration, StatusNoK7YukawaMap, StatusProjectiveNotYukawa}) {
		t.Fatalf("bad K7/projective candidate: %s", FormatCandidate(k7))
	}
}

func TestGate799RankingBranchFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ranking.Recorded || len(a.Ranking.Ranks) != 7 || !containsAll(a.Ranking.Ranks, []string{"1 Color", "2 External", "3 D4", "4 Generation", "5 Georgi", "6 SU(3)", "7 K7"}) {
		t.Fatalf("bad ranking: %s", FormatRanking(a.Ranking))
	}
	if !a.Branch.Recorded || !strings.Contains(a.Branch.Recommended, "D4 Triality") || !strings.Contains(a.Branch.RecommendationWhy, "Cl(1,7)") || !strings.Contains(a.Branch.ForbiddenPath, "symbolic") {
		t.Fatalf("bad branch: %s", FormatBranch(a.Branch))
	}
	if !a.Firewalls.Enforced || a.Firewalls.ThreeIsProof || a.Firewalls.NEffD4Theorem || a.Firewalls.NEffGenerationTheorem || a.Firewalls.HexagramEvidence || a.Firewalls.CHiggsLevelC || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate799TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Final, "does not prove") || !strings.Contains(a.Final, "color-tripled top dominance") || !strings.Contains(a.Final, "D4 triality carrier audit") {
		t.Fatalf("bad final: %s", a.Final)
	}
	res := Generation2NativeThreeSourceCandidateAndD4SU3CarrierFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
