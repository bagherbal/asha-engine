package generation2boundaryexteriordegreeindexedflagquotientselectoraudit

import (
	"strings"
	"testing"
)

func TestGate878DimensionMismatchForcesSelectorTyping(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Mismatch.SelectorNotLinearSurjection || a.Mismatch.Lambda1Dim == a.Mismatch.F1OverF0Rank || a.Mismatch.Lambda2Dim == a.Mismatch.F2OverF0Rank {
		t.Fatalf("bad mismatch ledger: %s", FormatMismatch(a.Mismatch))
	}
	if !containsAll(a.Mismatch.Failures, []string{FailureNotLinearMap}) {
		t.Fatalf("missing linear-surjection firewall: %s", FormatMismatch(a.Mismatch))
	}
}

func TestGate878DegreeIndexedSelectors(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DegreeOne.SelectedQuotient != "F_1/F_0" || a.DegreeOne.Target != "Pi_top" || a.DegreeOne.QuotientRank != 3 || !a.DegreeOne.SelectorMode || a.DegreeOne.LinearSurjection || a.DegreeOne.NativeSelector {
		t.Fatalf("bad degree one selector: %s", FormatSelector(a.DegreeOne))
	}
	if a.DegreeTwo.SelectedQuotient != "F_2/F_0" || a.DegreeTwo.Target != "H_R^min" || a.DegreeTwo.QuotientRank != 7 || !a.DegreeTwo.SelectorMode || a.DegreeTwo.LinearSurjection || a.DegreeTwo.NativeSelector {
		t.Fatalf("bad degree two selector: %s", FormatSelector(a.DegreeTwo))
	}
}

func TestGate878RejectsPureAssociatedGradedSlice(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.WrongSlice.Rejected || a.WrongSlice.RejectedQuotient != "F_2/F_1" || a.WrongSlice.RejectedRank != 4 || a.WrongSlice.RequiredRank != 7 {
		t.Fatalf("bad rejected target: %s", FormatRejected(a.WrongSlice))
	}
	if !containsAll(a.WrongSlice.Supports, []string{SupportWrongF2F1Rejected, SupportDegreeTwoNotPureSlice}) || !containsAll(a.WrongSlice.Failures, []string{FailureDegreeTwoNotF2OverF1}) {
		t.Fatalf("missing F2/F1 statuses: %s", FormatRejected(a.WrongSlice))
	}
}

func TestGate878AlphaReconstructedStillSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Alpha.ReconstructedFromSelectors || !near(a.Alpha.ReconstructedAlpha, AlphaB) || a.Alpha.NativeSelectorFunctor {
		t.Fatalf("bad alpha candidate: %s", FormatAlpha(a.Alpha))
	}
	if !containsAll(a.Alpha.Failures, []string{FailureNoNativeDegreeIndexedSelector, FailureAlphaStillSealed, FailureNoNativeAlphaSource}) {
		t.Fatalf("missing alpha firewalls: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate878R3AndOfficialLedgerBlocked(t *testing.T) {
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

func TestGate878Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate878Theorem(t *testing.T) {
	res := Generation2BoundaryExteriorDegreeIndexedFlagQuotientSelectorAuditTheorem().Verify()
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
