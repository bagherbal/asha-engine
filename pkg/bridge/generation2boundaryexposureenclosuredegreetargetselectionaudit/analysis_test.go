package generation2boundaryexposureenclosuredegreetargetselectionaudit

import (
	"strings"
	"testing"
)

func TestGate872ReducedResponseInherited(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Response.ZeroOrderSuppressed || !a.Response.HigherTruncated || !a.Response.DegreeOnePresent || !a.Response.DegreeTwoPresent || a.Response.NativeFunctionalCertified {
		t.Fatalf("bad reduced response: %s", FormatResponse(a.Response))
	}
}

func TestGate872ExposureEnclosureCandidates(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Exposure.MapCertified || a.Exposure.Degree != 1 || a.Exposure.Type != "single-boundary exposure" || a.Exposure.TargetRank != PiTopRank || a.Exposure.ChamberDim != H10Dim {
		t.Fatalf("bad exposure candidate: %s", FormatExposureEnclosure(a.Exposure))
	}
	if a.Enclosure.MapCertified || a.Enclosure.Degree != 2 || a.Enclosure.Type != "full boundary-pair enclosure" || a.Enclosure.TargetRank != HRminRank || a.Enclosure.ChamberDim != H72Dim {
		t.Fatalf("bad enclosure candidate: %s", FormatExposureEnclosure(a.Enclosure))
	}
}

func TestGate872CrossLaneTypeFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.ExposureToHRMinExcludedCandidate || !a.CrossLane.EnclosureToPiTopExcludedCandidate || a.CrossLane.CrossLaneExclusionTheorem {
		t.Fatalf("cross-lane overpromoted: %s", FormatCrossLane(a.CrossLane))
	}
	if !containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion}) {
		t.Fatalf("missing cross-lane failures: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate872PunctureRankSevenRole(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Puncture.PunctureRequiredForRankSeven || a.Puncture.AmbientRightRank-a.Puncture.PunctureRank != a.Puncture.ActiveRightRank || a.Puncture.PuncturedEnclosureTheorem {
		t.Fatalf("bad puncture role: %s", FormatPuncture(a.Puncture))
	}
}

func TestGate872AlphaReconstructionButStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.Native || a.Obstruction.AlphaNative {
		t.Fatalf("bad alpha candidate: %s | %s", FormatCandidate(a.Candidate), FormatObstruction(a.Obstruction))
	}
}

func TestGate872Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.R3.EligibleForR3 || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 {
		t.Fatalf("firewalls broken: %s | %s | %s", FormatFirewalls(a.Firewalls), FormatR3(a.R3), FormatImpact(a.Impact))
	}
}

func TestGate872Theorem(t *testing.T) {
	res := Generation2BoundaryExposureEnclosureDegreeTargetSelectionAuditTheorem().Verify()
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
