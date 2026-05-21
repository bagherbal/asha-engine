package generation2reducedboundarypairexteriorresponsedegreetargetselectionaudit

import (
	"strings"
	"testing"
)

func TestGate870ReducedResponseShape(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Response.Lambda0Removed || !a.Response.DegreeOnePresent || !a.Response.DegreeTwoPresent || !a.Response.CubicAndHigherVanish || a.Response.NativeFunctionalCertified {
		t.Fatalf("bad reduced response: %s", FormatResponse(a.Response))
	}
}

func TestGate870AlphaReconstructionOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.Native {
		t.Fatalf("bad alpha candidate: %s", FormatCandidate(a.Candidate))
	}
}

func TestGate870DegreeTargetFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Targets.DegreeOne.TargetMapCertified || a.Targets.DegreeTwo.TargetMapCertified || a.Targets.DegreeTargetSelectionTheorem || a.Targets.SharedTargetFunctor {
		t.Fatalf("degree targets overpromoted: %s", FormatTargets(a.Targets))
	}
	if !containsAll(a.Targets.Failures, []string{FailureNoDegreeTargetSelection, FailureNoTypedDegreeOneToPiTopMap, FailureNoTypedDegreeTwoToHRMinMap}) {
		t.Fatalf("missing target failures: %s", FormatTargets(a.Targets))
	}
}

func TestGate870CrossLaneAndNativeFunctionalBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.ZeroOrderSuppressed || !a.CrossLane.CubicAndHigherSuppressed || a.CrossLane.CrossLaneExclusionTheorem || a.CrossLane.NativeReducedFunctional || a.CrossLane.AlphaNative {
		t.Fatalf("cross-lane/native firewall broken: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate870Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 || a.R3.EligibleForR3 {
		t.Fatalf("firewalls broken: %s | %s | %s", FormatFirewalls(a.Firewalls), FormatImpact(a.Impact), FormatR3(a.R3))
	}
}

func TestGate870Theorem(t *testing.T) {
	res := Generation2ReducedBoundaryPairExteriorResponseDegreeTargetSelectionAuditTheorem().Verify()
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
