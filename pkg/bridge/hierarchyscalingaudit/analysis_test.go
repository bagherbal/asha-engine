package hierarchyscalingaudit

import "testing"

func nearlyEqual(a, b, tol float64) bool {
	if a > b {
		return a-b <= tol
	}
	return b-a <= tol
}

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.SInst <= 12 || a.Inputs.STop <= 78 {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestHierarchyTargets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Targets.RhoUnreduced, 2.016725503526116e-17, 1e-29) {
		t.Fatalf("unexpected unreduced target: %s", FormatTargets(a.Targets))
	}
	if !nearlyEqual(a.Targets.RhoReduced, 1.011036233861601e-16, 1e-28) {
		t.Fatalf("unexpected reduced target: %s", FormatTargets(a.Targets))
	}
}

func TestCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !hasCandidate(a.Candidates, "B-gap instanton exponential") || !hasCandidate(a.Candidates, "rank-56 Boolean near miss") {
		t.Fatalf("missing expected candidates: %s", FormatCandidates(a.Candidates))
	}
	if a.Candidates.BestUnreduced.Name != "rank-56 Boolean near miss" {
		t.Fatalf("unexpected best unreduced near miss: %s", FormatCandidate(a.Candidates.BestUnreduced))
	}
	if a.Candidates.BestUnreduced.Promotable {
		t.Fatalf("rank-56 near miss should not be promotable: %s", FormatCandidate(a.Candidates.BestUnreduced))
	}
}

func TestSynthesisAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Synthesis.NativeDerived || !hasFitLaneRejected(a.Synthesis) {
		t.Fatalf("bad synthesis sieve: %s", FormatSynthesis(a.Synthesis))
	}
	if !a.Firewalls.NoHierarchyScalingFactor || !a.Firewalls.F2MomentUnlocked || !a.Firewalls.ArbitraryExponentFittingRejected {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusHierarchyRatioFormalized, StatusTopologicalCandidatesAudited, StatusFailedHierarchyDerived, StatusFailedArbitraryPowerFitting}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := GaugeHierarchyScalingAuditPlanckFactorSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
