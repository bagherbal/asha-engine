package generation2boundaryexteriordegreetargetselectionmapaudit

import (
	"strings"
	"testing"
)

func TestGate871ReducedResponseInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Response.ZeroOrderSuppressed || !a.Response.HigherTruncated || !a.Response.DegreeOnePresent || !a.Response.DegreeTwoPresent || a.Response.NativeFunctionalCertified {
		t.Fatalf("bad inherited response: %s", FormatResponse(a.Response))
	}
}

func TestGate871DegreeTargetCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DegreeOne.MapCertified || a.DegreeOne.TargetRank != PiTopRank || a.DegreeOne.ChamberDim != H10Dim {
		t.Fatalf("bad degree-one target: %s", FormatTarget(a.DegreeOne))
	}
	if a.DegreeTwo.MapCertified || a.DegreeTwo.TargetRank != HRminRank || a.DegreeTwo.ChamberDim != H72Dim {
		t.Fatalf("bad degree-two target: %s", FormatTarget(a.DegreeTwo))
	}
}

func TestGate871CrossLaneFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.LinearToHRMinExcludedCandidate || !a.CrossLane.QuadraticToPiTopExcludedCandidate || a.CrossLane.CrossLaneExclusionTheorem {
		t.Fatalf("cross-lane audit overpromoted: %s", FormatCrossLane(a.CrossLane))
	}
	if !containsAll(a.CrossLane.Failures, []string{FailureNoCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion}) {
		t.Fatalf("missing cross-lane failures: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate871AlphaReconstructionButNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.Native {
		t.Fatalf("bad alpha candidate: %s", FormatCandidate(a.Candidate))
	}
}

func TestGate871Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.R3.EligibleForR3 || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 {
		t.Fatalf("firewalls broken: %s | %s | %s", FormatFirewalls(a.Firewalls), FormatR3(a.R3), FormatImpact(a.Impact))
	}
}

func TestGate871Theorem(t *testing.T) {
	res := Generation2BoundaryExteriorDegreeTargetSelectionMapAuditTheorem().Verify()
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
