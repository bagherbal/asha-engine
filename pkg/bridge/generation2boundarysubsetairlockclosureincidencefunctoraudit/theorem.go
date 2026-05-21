package generation2boundarysubsetairlockclosureincidencefunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE927-BOUNDARY-SUBSET-AIRLOCK-CLOSURE-INCIDENCE-FUNCTOR-AUDIT"
	theoremName = "Gate 927: BoundarySubset AirlockClosure IncidenceFunctor Audit"
)

func Generation2BoundarySubsetAirlockClosureIncidenceFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 927 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "boundary exterior degrees factor through finite boundary subset cardinality", Passed: a.Subset.ExteriorDegreeEqualsCardinal && a.Subset.NativeFiniteSubsetSource && !a.Subset.SelectsAirlockTargetsByItself && containsAll(a.Subset.Supports, []string{SupportExteriorDegreeEqualsSubsetCardinality, SupportLambda1SingletonSubsets, SupportLambda2FullPairSubset, SupportSourceChainSubsetLattice}) && containsAll(a.Subset.Failures, []string{FailureSubsetCardinalityAloneNoTargets}), Detail: FormatSubset(a.Subset)},
			{Name: "airlock flag is typed as a closure ladder candidate but no native closure operator is certified", Passed: a.Ladder.ClosureLadderType && !a.Ladder.NativeClosureOperator && containsAll(a.Ladder.Supports, []string{SupportAirlockFlagClosureLadder, SupportF0PunctureBasepoint, SupportF1MinimalExposedClosure, SupportF2SaturatedFullRightRectangle}) && containsAll(a.Ladder.Failures, []string{FailureAirlockClosureNotNative}), Detail: FormatLadder(a.Ladder)},
			{Name: "Theta_B^Z2 factors through closure level quotient over F0", Passed: a.Factorization.FactorsTheta && a.Factorization.QuotientReconstructsTheta && !a.Factorization.NativeClosureTheorem && containsAll(a.Factorization.Supports, []string{SupportThetaFactorsThroughClosure, SupportDegreeToFlagHasClosureSource, SupportClosureQuotientReconstructsTheta}) && containsAll(a.Factorization.Failures, []string{FailureClosureFunctorCandidateNotNative}), Detail: FormatFactorization(a.Factorization)},
			{Name: "closure levels map empty, singleton, and full-pair activations to F0, F1, and F2 candidates", Passed: a.Basepoint.ClosureSupported && a.Singleton.ClosureSupported && a.FullPair.ClosureSupported && a.Basepoint.MatchesReducedForm && a.Singleton.MatchesReducedForm && a.FullPair.MatchesReducedForm && !a.Basepoint.NativeClosureTheorem && !a.Singleton.NativeClosureTheorem && !a.FullPair.NativeClosureTheorem && containsAll(a.Basepoint.Supports, []string{SupportEmptySubsetMapsF0, SupportBasepointClosureMatchesReduction, SupportDegreeZeroNoActiveAlpha}) && containsAll(a.Singleton.Supports, []string{SupportSingletonClosesF1, SupportOneBoundaryGeneratesF1, SupportThetaOneFromSingletonClosure}) && containsAll(a.FullPair.Supports, []string{SupportFullPairClosesF2, SupportTwoBoundaryGeneratesF2, SupportThetaTwoFromFullPairClosure}), Detail: FormatClosureLevel(a.Basepoint) + " | " + FormatClosureLevel(a.Singleton) + " | " + FormatClosureLevel(a.FullPair)},
			{Name: "cumulative F2/F0 quotient follows from fixed-basepoint closure form while associated-graded F2/F1 remains rejected", Passed: a.Cumulative.FollowsFromBasepoint && a.Cumulative.RejectsAssociatedGraded && !a.Cumulative.NativeBasepointRule && a.Cumulative.CumulativeRank == RankF2OverF0 && a.Cumulative.AssociatedGradedRank == RankF2OverF1 && containsAll(a.Cumulative.Supports, []string{SupportCumulativeF2OverF0FromBase, SupportGradedRejectedByBasepointClosure, SupportTopDegreeClosureOverPuncture}) && containsAll(a.Cumulative.Failures, []string{FailureFixedBasepointQuotientNotNative}), Detail: FormatCumulative(a.Cumulative)},
			{Name: "closure functor is unique under monotone-minimal-saturated-Z2 rules but those rules are not native", Passed: a.Uniqueness.UniqueUnderRules && a.Uniqueness.Monotone && a.Uniqueness.MinimalSingleton && a.Uniqueness.SaturatedFullPair && a.Uniqueness.Z2Invariant && !a.Uniqueness.NativeMinimalSaturation && containsAll(a.Uniqueness.Supports, []string{SupportClosureUniqueRules, SupportMinimalityForcesF1, SupportSaturationForcesF2, SupportZ2ClassInvarianceClosure}) && containsAll(a.Uniqueness.Failures, []string{FailureMinimalSaturationNotNative}), Detail: FormatUniqueness(a.Uniqueness)},
			{Name: "BoundaryActivationMeasure can be rewritten using closure targets but remains non-native", Passed: a.Measure.ClosureSuppliesTargets && a.Measure.MeasureUsesClosure && a.Measure.AlphaReconstructed && !a.Measure.NativeMeasureByClosure && a.Measure.ThetaRankOne == RankF1OverF0 && a.Measure.ThetaRankTwo == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportClosureSuppliesThetaTargets, SupportMeasureUsingAirlockClosure, SupportAlphaViaClosureMeasure}) && containsAll(a.Measure.Failures, []string{FailureMuBNotNativeWithoutClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.InheritedStatus, BoundaryPair, BoundarySubsetLattice, SourceDegreeChain, CardinalityChain, AirlockFlagChain, Z2PunctureClass, ClosureFunctor, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, ThetaOne, ThetaTwo, MeasureViaClosure, AlphaFormula, AssociatedGradedTarget, NextGate, FormatSubset(a.Subset), FormatLadder(a.Ladder), FormatFactorization(a.Factorization), FormatClosureLevel(a.Basepoint), FormatClosureLevel(a.Singleton), FormatClosureLevel(a.FullPair), FormatCumulative(a.Cumulative), FormatUniqueness(a.Uniqueness), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
