package generation2boundarypairfacefiltrationtargetfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate877FlagQuotients(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flag.ValidNestedFlag || a.Flag.RankP != 1 || a.Flag.RankF1 != 4 || a.Flag.RankF2 != 8 {
		t.Fatalf("bad flag: %s", FormatFlag(a.Flag))
	}
	if !a.DegreeOne.MatchesTarget || a.DegreeOne.QuotientRank != 3 || a.DegreeOne.QuotientRank != a.DegreeOne.LayerRank-a.DegreeOne.PunctureRank {
		t.Fatalf("bad degree one quotient: %s", FormatQuotient(a.DegreeOne))
	}
	if !a.DegreeTwo.MatchesTarget || a.DegreeTwo.QuotientRank != 7 || a.DegreeTwo.QuotientRank != a.DegreeTwo.LayerRank-a.DegreeTwo.PunctureRank {
		t.Fatalf("bad degree two quotient: %s", FormatQuotient(a.DegreeTwo))
	}
}

func TestGate877AlphaReconstructionStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Alpha.ReconstructedFromFlagQuotients || !near(a.Alpha.ReconstructedAlpha, AlphaB) || a.Alpha.NativeFunctor {
		t.Fatalf("bad alpha candidate: %s", FormatAlpha(a.Alpha))
	}
	if !containsAll(a.Alpha.Failures, []string{FailureNoNativeDegreeToFlagFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}) {
		t.Fatalf("missing alpha firewalls: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate877CrossLaneConditionalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.ExcludedIfFunctor || a.CrossLane.NativeExclusion {
		t.Fatalf("cross lane promoted: %s", FormatCrossLane(a.CrossLane))
	}
	if !containsAll(a.CrossLane.Failures, []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeDegreeToFlagFunctor}) {
		t.Fatalf("missing cross-lane firewalls: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate877R3AndOfficialLedgerBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.R3.EligibleForR3 || a.R3.EligibleForR4 || a.Impact.CanPromoteToR3 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("promotion leaked: %s | %s", FormatR3(a.R3), FormatImpact(a.Impact))
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.CanUpdate {
		t.Fatalf("ledger freeze broken: %s", FormatLedger(a.Ledger))
	}
}

func TestGate877Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate877Theorem(t *testing.T) {
	res := Generation2BoundaryPairFaceFiltrationTargetFunctorAuditTheorem().Verify()
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
