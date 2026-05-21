package generation2boundarysubsetairlockclosureincidencefunctoraudit

import (
	"strings"
	"testing"
)

func TestGate927SubsetLatticeAndClosureLadder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.InheritedStatus != Gate926ShortStatus {
		t.Fatalf("bad inherited status: %s", a.InheritedStatus)
	}
	if !a.Subset.ExteriorDegreeEqualsCardinal || !a.Subset.NativeFiniteSubsetSource || a.Subset.SelectsAirlockTargetsByItself {
		t.Fatalf("bad subset audit: %s", FormatSubset(a.Subset))
	}
	if !a.Ladder.ClosureLadderType || a.Ladder.NativeClosureOperator {
		t.Fatalf("bad ladder audit: %s", FormatLadder(a.Ladder))
	}
}

func TestGate927ClosureFactorizationAndLevels(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Factorization.FactorsTheta || !a.Factorization.QuotientReconstructsTheta || a.Factorization.NativeClosureTheorem {
		t.Fatalf("bad factorization: %s", FormatFactorization(a.Factorization))
	}
	if !a.Basepoint.ClosureSupported || !a.Basepoint.MatchesReducedForm || a.Basepoint.NativeClosureTheorem || a.Basepoint.ClosureTarget != "F_0" {
		t.Fatalf("bad basepoint closure: %s", FormatClosureLevel(a.Basepoint))
	}
	if !a.Singleton.ClosureSupported || !a.Singleton.MatchesReducedForm || a.Singleton.NativeClosureTheorem || a.Singleton.ClosureTarget != "F_1" {
		t.Fatalf("bad singleton closure: %s", FormatClosureLevel(a.Singleton))
	}
	if !a.FullPair.ClosureSupported || !a.FullPair.MatchesReducedForm || a.FullPair.NativeClosureTheorem || a.FullPair.ClosureTarget != "F_2" {
		t.Fatalf("bad full-pair closure: %s", FormatClosureLevel(a.FullPair))
	}
}

func TestGate927CumulativeUniquenessMeasureAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Cumulative.FollowsFromBasepoint || !a.Cumulative.RejectsAssociatedGraded || a.Cumulative.NativeBasepointRule || a.Cumulative.CumulativeRank != RankF2OverF0 || a.Cumulative.AssociatedGradedRank != RankF2OverF1 {
		t.Fatalf("bad cumulative audit: %s", FormatCumulative(a.Cumulative))
	}
	if !a.Uniqueness.UniqueUnderRules || !a.Uniqueness.Monotone || !a.Uniqueness.MinimalSingleton || !a.Uniqueness.SaturatedFullPair || !a.Uniqueness.Z2Invariant || a.Uniqueness.NativeMinimalSaturation {
		t.Fatalf("bad uniqueness audit: %s", FormatUniqueness(a.Uniqueness))
	}
	if !a.Measure.ClosureSuppliesTargets || !a.Measure.MeasureUsesClosure || !a.Measure.AlphaReconstructed || a.Measure.NativeMeasureByClosure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 {
		t.Fatalf("bad measure audit: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	for _, want := range []string{FailureClosureFunctorCandidateNotNative, FailureSingletonToF1ClosureNotNative, FailureFullPairToF2ClosureNotNative, FailureFixedBasepointQuotientNotNative, FailureMuBNotNativeWithoutClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3} {
		if !containsAll(a.Firewalls.List(), []string{want}) {
			t.Fatalf("missing firewall %s", want)
		}
	}
}

func TestGate927Theorem(t *testing.T) {
	res := Generation2BoundarySubsetAirlockClosureIncidenceFunctorAuditTheorem().Verify()
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
	for _, want := range []string{FinalTruth, Classification, ShortStatus, BoundaryPair, BoundarySubsetLattice, SourceDegreeChain, CardinalityChain, AirlockFlagChain, Z2PunctureClass, ClosureFunctor, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, ThetaOne, ThetaTwo, MeasureViaClosure, AlphaFormula, AssociatedGradedTarget, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
