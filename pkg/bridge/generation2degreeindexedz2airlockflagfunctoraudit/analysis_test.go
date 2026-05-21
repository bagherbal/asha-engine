package generation2degreeindexedz2airlockflagfunctoraudit

import (
	"strings"
	"testing"
)

func TestGate914InheritedResponseDoesNotReopenClosedWounds(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ReducedShapeCertified || a.Inherited.ReopensPhaseSign || a.Inherited.ReopensSocketOrder || a.Inherited.ReopensRepresentative || a.Inherited.DerivesAlpha || a.Inherited.UpdatesOfficialLedger {
		t.Fatalf("bad inherited response: %s", FormatInherited(a.Inherited))
	}
}

func TestGate914DegreeTargets(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.DegreeOne.Degree != 1 || a.DegreeOne.BoundaryTerm != DegreeOneTerm || a.DegreeOne.Target.Rank != RankF1OverF0 || !a.DegreeOne.Target.RepresentativeFree {
		t.Fatalf("bad degree-one target: %s", FormatDegreeTarget(a.DegreeOne))
	}
	if a.DegreeTwo.Degree != 2 || a.DegreeTwo.BoundaryTerm != DegreeTwoTerm || a.DegreeTwo.Target.Rank != RankF2OverF0 || !a.DegreeTwo.Target.RepresentativeFree {
		t.Fatalf("bad degree-two target: %s", FormatDegreeTarget(a.DegreeTwo))
	}
	if !a.DegreeOne.Selector || !a.DegreeTwo.Selector || a.DegreeOne.LinearSurjection || a.DegreeTwo.LinearSurjection || a.DegreeOne.NativeMap || a.DegreeTwo.NativeMap {
		t.Fatalf("selector/map firewall leaked: %s | %s", FormatDegreeTarget(a.DegreeOne), FormatDegreeTarget(a.DegreeTwo))
	}
}

func TestGate914SelectorNotSurjection(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Typing.Lambda1Dim != Lambda1Dim || a.Typing.Lambda2Dim != Lambda2Dim || a.Typing.ExposedRank != RankF1OverF0 || a.Typing.FullRank != RankF2OverF0 || !a.Typing.DimensionMismatch || !a.Typing.SelectorNotSurjection {
		t.Fatalf("bad selector typing: %s", FormatTyping(a.Typing))
	}
}

func TestGate914CumulativeEnclosureChoice(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Cumulative.F2OverF0Rank != RankF2OverF0 || a.Cumulative.F2OverF1Rank != RankF2OverF1 || !a.Cumulative.SelectsCumulativeEnclosure || !a.Cumulative.RejectsAssociatedGradedSlice || a.Cumulative.NativeReasonForChoice {
		t.Fatalf("bad cumulative choice: %s", FormatCumulative(a.Cumulative))
	}
}

func TestGate914AlphaRankReconstructionAndExternalInputs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.AlphaRanks.ReconstructsRankPair || !near(a.AlphaRanks.Alpha, AlphaB) || a.AlphaRanks.RankPair != [2]int{RankF1OverF0, RankF2OverF0} {
		t.Fatalf("bad alpha rank reconstruction: %s", FormatAlpha(a.AlphaRanks))
	}
	if a.AlphaRanks.NativeAlphaSource || !a.AlphaRanks.DenominatorsExternal || !a.AlphaRanks.STransportExternal {
		t.Fatalf("alpha source firewall leaked: %s", FormatAlpha(a.AlphaRanks))
	}
}

func TestGate914CrossLaneStatus(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.CrossLane.ForbiddenLanes, []string{ForbiddenLinearFull, ForbiddenQuadraticFace}) || !containsAll(a.CrossLane.FalseTerms, []string{FalseLinearFullTerm, FalseQuadraticFaceTerm}) {
		t.Fatalf("bad cross-lane forbidden data: %s", FormatCrossLane(a.CrossLane))
	}
	if !a.CrossLane.WouldFollowFromSelector || a.CrossLane.IndependentNativeTheorem || a.CrossLane.ProvesCrossLaneExclusion {
		t.Fatalf("cross-lane firewall leaked: %s", FormatCrossLane(a.CrossLane))
	}
}

func TestGate914GlobalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeLambda1ExposedMap, FailureNoNativeLambda2FullMap, FailureLambdaKB2NotSurjection, FailureNoNativeCumulativeReason, FailureNoIndependentCrossLane, FailureSelectorRanksNotAlphaSource, FailureDenominatorsSTransportExternal, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator} {
		if !containsAll(a.FirewallsList(), []string{want}) {
			t.Fatalf("missing firewall %s from %v", want, a.FirewallsList())
		}
	}
}

func TestGate914Theorem(t *testing.T) {
	res := Generation2DegreeIndexedZ2AirlockFlagFunctorAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, StrategicConclusion, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s in notes: %s", want, joined)
		}
	}
}
