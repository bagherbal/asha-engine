package generation2socketorderorientededgeairlockconsolidationaudit

import (
	"strings"
	"testing"
)

func TestGate897SocketPairTypedButUnordered(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SocketOrder
	if !s.PairTyped || s.NativeOrder || s.PlusSelected || !s.RequiresSelector {
		t.Fatalf("bad socket order: %s", FormatSocketOrder(s))
	}
	if !containsAll(s.Characters, []string{RightCharacterPlus, RightCharacterMinus}) || !containsAll(s.Failures, []string{FailureNoNativeSocketOrderSelector, FailureCharacterConjugationNoPlus}) {
		t.Fatalf("missing socket order statuses: %s", FormatSocketOrder(s))
	}
}

func TestGate897SourcesDoNotSelectPlus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	s := a.SocketOrder
	if s.BoundaryDegreeSelectsPlus || s.JOrChiralitySelectsPlus || s.BMinusLSelectsPlus {
		t.Fatalf("source incorrectly selects plus: %s", FormatSocketOrder(s))
	}
	if !containsAll(s.Failures, []string{FailureBoundaryDegreeNoPlus, FailureJChiralityNoPlus, FailureBMinusLNoPlus}) {
		t.Fatalf("missing source failures: %s", FormatSocketOrder(s))
	}
}

func TestGate897EdgeTableHasZ2MirrorAndReducesToSocketOrder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	e := a.EdgeOrdering
	if !e.CurrentFollowsIfPlusExposed || !e.MirrorExistsIfMinusExposed || !e.SameRankPattern || e.SelectsPlusByItself || !e.ReducesToSocketOrder {
		t.Fatalf("bad edge ordering: %s", FormatEdgeOrdering(e))
	}
	if !containsAll(e.Supports, []string{SupportCurrentTableIfPlusExposed, SupportZ2MirrorTableExists, SupportEdgeOrderingReducesToSocket}) || !containsAll(e.Failures, []string{FailureEdgeTableNoPlusByItself, FailureEdgeOrderingCircular}) {
		t.Fatalf("missing edge statuses: %s", FormatEdgeOrdering(e))
	}
}

func TestGate897AirlockExistsAsZ2Family(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.AirlockFamily
	if !f.BothReconstructAlpha || !f.Z2Family || f.OrderedRepresentativeCertified || !near(f.PlusAlpha, AlphaB) || !near(f.MinusAlpha, AlphaB) {
		t.Fatalf("bad airlock family: %s", FormatAirlockFamily(f))
	}
	if !containsAll(f.Failures, []string{FailureNoNativeNeutralPunctureAirlock, FailureNoNativeOrderedAirlock, FailureNoNativeSelectionSigmaPlus}) {
		t.Fatalf("missing family failures: %s", FormatAirlockFamily(f))
	}
}

func TestGate897FreezeAndFirewalls(t *testing.T) {
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

func TestGate897Theorem(t *testing.T) {
	res := Generation2SocketOrderOrientedEdgeAirlockConsolidationAuditTheorem().Verify()
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
