package generation2boundaryalphaincidenceflagsealclassificationaudit

import (
	"strings"
	"testing"
)

func TestGate880AlphaSealClassification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Seal.Name != SealName || a.Seal.FullName != FullSealName || !a.Seal.ReducedExteriorShape || !a.Seal.DegreeIndexedFlagSelector || a.Seal.NativeFunctor {
		t.Fatalf("bad seal: %s", FormatSeal(a.Seal))
	}
	if !near(a.Seal.Alpha, AlphaB) || a.Seal.RankF1OverF0 != 3 || a.Seal.RankF2OverF0 != 7 {
		t.Fatalf("bad seal reconstruction: %s", FormatSeal(a.Seal))
	}
}

func TestGate880ConditionalChainCoherentOnlyGivenSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Chain.CoherentGivenSeal || !a.Chain.ReducedExteriorToAlpha || !a.Chain.AlphaToSocketMagnitudes || !a.Chain.SocketMagnitudesToYDagY || !a.Chain.YDagYToHAgg || !a.Chain.HAggToNEff {
		t.Fatalf("bad conditional chain: %s", FormatChain(a.Chain))
	}
	if !containsAll(a.Chain.Failures, []string{FailureAlphaStillSealed, FailureConditionalProxyNotR3, FailureNoNativeSectorTraceMagnitude}) {
		t.Fatalf("missing chain firewalls: %s", FormatChain(a.Chain))
	}
}

func TestGate880MissingNativeFunctorFiled(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.MissingTheorem.Name != "BoundaryExteriorIncidenceFlagFunctor" || a.MissingTheorem.Native || !a.MissingTheorem.RequiredForR3 {
		t.Fatalf("bad missing theorem: %s", FormatMissing(a.MissingTheorem))
	}
	if !containsAll(a.MissingTheorem.Failures, []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion}) {
		t.Fatalf("missing native functor firewalls: %s", FormatMissing(a.MissingTheorem))
	}
}

func TestGate880R3AndOfficialLedgerBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.OfficialFrozen || a.Ledger.CanUpdate || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("official freeze leaked: %s | %s", FormatLedger(a.Ledger), FormatImpact(a.Impact))
	}
	if a.Eligibility.EligibleForR3 || a.Eligibility.EligibleForR4 || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		t.Fatalf("promotion leaked: %s | %s", FormatEligibility(a.Eligibility), FormatImpact(a.Impact))
	}
}

func TestGate880Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate880Theorem(t *testing.T) {
	res := Generation2BoundaryAlphaIncidenceFlagSealClassificationAuditTheorem().Verify()
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
