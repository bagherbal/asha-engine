package generation2nestedpuncturecomplementtargetselectionfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate876NestedComplementsReconstructTargets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DegreeOne.MatchesTarget || a.DegreeOne.ComplementRank != PiTopRank || a.DegreeOne.AmbientRank-a.DegreeOne.PunctureRank != PiTopRank {
		t.Fatalf("bad degree one complement: %s", FormatComplement(a.DegreeOne))
	}
	if !a.DegreeTwo.MatchesTarget || a.DegreeTwo.ComplementRank != HRMinRank || a.DegreeTwo.AmbientRank-a.DegreeTwo.PunctureRank != HRMinRank {
		t.Fatalf("bad degree two complement: %s", FormatComplement(a.DegreeTwo))
	}
}

func TestGate876AlphaReconstructionStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Candidate.ShapeCoherent || !near(a.Candidate.ReconstructedAlpha, AlphaB) || a.Candidate.NativeFunctor {
		t.Fatalf("bad alpha candidate: %s", FormatCandidate(a.Candidate))
	}
	if !containsAll(a.Candidate.Failures, []string{FailureNestedComplementNotNativeFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}) {
		t.Fatalf("missing alpha firewalls: %s", FormatCandidate(a.Candidate))
	}
}

func TestGate876CrossLaneStillNotNative(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.TypeCandidate || a.CrossLane.NativeExclusion {
		t.Fatalf("cross lane promoted: %s", FormatCrossLane(a.CrossLane))
	}
	if !containsAll(a.CrossLane.Failures, []string{FailureNoFaceVsEnclosureDegreeTheorem, FailureNoNativeCrossLaneExclusionTheorem}) {
		t.Fatalf("missing cross lane failures: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate876R3AndOfficialLedgerBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.R3.EligibleForR3 || a.R3.EligibleForR4 || a.Impact.CanPromoteToR3 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("promotion leaked: %s | %s", FormatR3(a.R3), FormatImpact(a.Impact))
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.CanUpdateOfficial {
		t.Fatalf("ledger freeze broken: %s", FormatLedger(a.Ledger))
	}
}

func TestGate876Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate876Theorem(t *testing.T) {
	res := Generation2NestedPunctureComplementTargetSelectionFunctorAuditTheorem().Verify()
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
