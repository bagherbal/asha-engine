package generation2neutralpunctureairlockvariationalfunctionalaudit

import (
	"strings"
	"testing"
)

func TestGate896RankOnePunctureCandidatesAreOnlyLeptonCells(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Rank
	if !r.OnlyLeptonCellsAreRankOne || !r.BothLeptonCandidatesPass || r.SelectsEPlusUniquely || r.ActiveComplementRank != 7 {
		t.Fatalf("bad rank audit: %s", FormatRankTerm(r))
	}
	if !containsAll(r.RankOneCandidates, []string{PuncturePlus, PunctureMinus}) || !containsAll(r.Failures, []string{FailureRankDoesNotSelectPlus}) {
		t.Fatalf("missing rank candidates/failures: %s", FormatRankTerm(r))
	}
}

func TestGate896AlphaFlagRanksDoNotDistinguishPlusMinus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.AlphaFlag
	if !f.BothReconstructAlphaShape || f.DistinguishesPlusFromMinus || !near(f.PlusAlpha, AlphaB) || !near(f.MinusAlpha, AlphaB) {
		t.Fatalf("bad alpha flag audit: %s", FormatAlphaFlagTerm(f))
	}
	if !containsAll(f.Failures, []string{FailureAlphaFlagDoesNotSelectPlus, FailureAlphaStillSealed}) {
		t.Fatalf("missing alpha flag failures: %s", FormatAlphaFlagTerm(f))
	}
}

func TestGate896BMinusLCompensationDoesNotSelectPlus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	b := a.BMinusL
	if b.PlusLeptonCharge != -1 || b.MinusLeptonCharge != -1 || !b.FullRectangleNeutral || b.DistinguishesPlusFromMinus {
		t.Fatalf("bad B-L audit: %s", FormatBMinusLTerm(b))
	}
	if !containsAll(b.Failures, []string{FailureBMinusLDoesNotSelectPlus}) {
		t.Fatalf("missing B-L failure: %s", FormatBMinusLTerm(b))
	}
}

func TestGate896EdgeAndKernelSelectPlusOnlyWithCircularOrdering(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	e := a.Edge
	if !e.SelectsEPlusAsNullEdge || e.IndependentEdgeOrdering || !e.CircularWithoutOrdering {
		t.Fatalf("bad edge audit: %s", FormatEdgeSupportTerm(e))
	}
	k := a.Kernel
	if !k.MatchesEPlusPuncture || !k.DependsOnPreselectedImage || k.NativeKernelSelector || k.Kernel != LeftKernel {
		t.Fatalf("bad kernel audit: %s", FormatLeftKernelTerm(k))
	}
}

func TestGate896FunctionalObstructionSharpensToEdgeOrdering(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Functional
	if !f.Formulated || !f.EPlusSatisfiesAllTerms || !f.EMinusAlsoSatisfiesRankFlagBL || !f.RequiresOrientedEdgeOrdering || f.NativeFunctional {
		t.Fatalf("bad functional audit: %s", FormatFunctional(f))
	}
	if !strings.Contains(f.NextMissingObject, "OrientedEdgeOrderingFunctional") {
		t.Fatalf("wrong next missing object: %s", FormatFunctional(f))
	}
}

func TestGate896FreezeAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Freeze.Frozen || !a.Freeze.DiagnosticOnly || a.Freeze.CanUpdate || near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff) {
		t.Fatalf("freeze leak: %s", FormatFreeze(a.Freeze))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate896Theorem(t *testing.T) {
	res := Generation2NeutralPunctureAirlockVariationalFunctionalAuditTheorem().Verify()
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
