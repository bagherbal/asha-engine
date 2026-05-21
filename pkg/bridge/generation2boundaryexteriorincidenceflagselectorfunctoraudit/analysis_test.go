package generation2boundaryexteriorincidenceflagselectorfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate879SourceIncidenceNotLinearMap(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Source.SelectorIndexNotGenerator || !a.Source.DegreeOneBeforeDegreeTwo || !a.Source.Lambda3Zero {
		t.Fatalf("bad source incidence: %s", FormatSource(a.Source))
	}
	if !containsAll(a.Source.Failures, []string{FailureNotLinearMap, FailureNoNativeIncidenceFunctor}) {
		t.Fatalf("missing incidence firewalls: %s", FormatSource(a.Source))
	}
}

func TestGate879TargetFlagQuotients(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Flag.Nested || !a.Flag.QuotientsValid || a.Flag.F1OverF0Rank != 3 || a.Flag.F2OverF0Rank != 7 {
		t.Fatalf("bad target flag: %s", FormatFlag(a.Flag))
	}
}

func TestGate879IncidenceSelectors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DegreeOne.SelectedQuotient != "F_1/F_0" || a.DegreeOne.Target != "Pi_top" || a.DegreeOne.QuotientRank != 3 || !a.DegreeOne.SelectorMode || a.DegreeOne.LinearSurjection || a.DegreeOne.NativeFunctor {
		t.Fatalf("bad degree one selector: %s", FormatSelector(a.DegreeOne))
	}
	if a.DegreeTwo.SelectedQuotient != "F_2/F_0" || a.DegreeTwo.Target != "H_R^min" || a.DegreeTwo.QuotientRank != 7 || !a.DegreeTwo.SelectorMode || a.DegreeTwo.LinearSurjection || a.DegreeTwo.NativeFunctor {
		t.Fatalf("bad degree two selector: %s", FormatSelector(a.DegreeTwo))
	}
}

func TestGate879CrossLaneAndAlphaStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.CrossLane.ExcludedIfFunctor || a.CrossLane.NativeExclusion {
		t.Fatalf("bad cross lane: %s", FormatCrossLane(a.CrossLane))
	}
	if !a.Alpha.ReconstructedByIncidence || !near(a.Alpha.ReconstructedAlpha, AlphaB) || a.Alpha.NativeIncidenceFunctor {
		t.Fatalf("bad alpha: %s", FormatAlpha(a.Alpha))
	}
	if !containsAll(a.Alpha.Failures, []string{FailureNoNativeIncidenceFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}) {
		t.Fatalf("missing alpha firewalls: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate879R3AndOfficialLedgerBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.R3.EligibleForR3 || a.R3.EligibleForR4 || a.Impact.CanPromoteToR3 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("promotion leaked: %s | %s", FormatR3(a.R3), FormatImpact(a.Impact))
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.CanUpdate {
		t.Fatalf("official ledger freeze broken: %s", FormatLedger(a.Ledger))
	}
}

func TestGate879Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate879Theorem(t *testing.T) {
	res := Generation2BoundaryExteriorIncidenceFlagSelectorFunctorAuditTheorem().Verify()
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
