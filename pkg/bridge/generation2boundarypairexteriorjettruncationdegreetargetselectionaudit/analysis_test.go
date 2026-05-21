package generation2boundarypairexteriorjettruncationdegreetargetselectionaudit

import (
	"strings"
	"testing"
)

func TestGate869ExteriorTruncation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.BoundaryPair.Dimension != 2 || a.BoundaryPair.Lambda3Dim != 0 || !a.BoundaryPair.TruncatesAfterSecondDegree || a.BoundaryPair.TruncationDerivesAlphaResponse {
		t.Fatalf("bad exterior truncation: %s", FormatBoundaryPair(a.BoundaryPair))
	}
}

func TestGate869DegreeTargetsRemainUncertified(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Targets.DegreeOne.TargetMapCertified || a.Targets.DegreeTwo.TargetMapCertified || a.Targets.DegreeTargetSelectionTheorem || a.Targets.SharedDegreeFunctor {
		t.Fatalf("degree targets overpromoted: %s", FormatTargets(a.Targets))
	}
	if !containsAll(a.Targets.Failures, []string{FailureNoDegreeTargetSelection, FailureNoDegreeOneToPiTopMap, FailureNoDegreeTwoToHRMinMap}) {
		t.Fatalf("missing failures: %s", FormatTargets(a.Targets))
	}
}

func TestGate869FormalAlphaReconstructionOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.Native || a.Candidate.ExteriorFunctionalCertified {
		t.Fatalf("candidate malformed: %s", FormatCandidate(a.Candidate))
	}
}

func TestGate869ZeroOrderAndCrossLaneFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ZeroCross.ZeroOrderPresent || a.ZeroCross.ZeroOrderContributes || !a.ZeroCross.CubicAndHigherAbsent || !a.ZeroCross.CubicStopDerivedByLambda3B2Zero || a.ZeroCross.ZeroOrderSuppressionTheorem || a.ZeroCross.CrossLaneExclusionTheorem {
		t.Fatalf("bad zero/cross audit: %s", FormatZeroCross(a.ZeroCross))
	}
}

func TestGate869Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 || a.R3.EligibleForR3 {
		t.Fatalf("firewalls broken: %s | %s | %s", FormatFirewalls(a.Firewalls), FormatImpact(a.Impact), FormatR3(a.R3))
	}
}

func TestGate869Theorem(t *testing.T) {
	res := Generation2BoundaryPairExteriorJetTruncationDegreeTargetSelectionAuditTheorem().Verify()
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
