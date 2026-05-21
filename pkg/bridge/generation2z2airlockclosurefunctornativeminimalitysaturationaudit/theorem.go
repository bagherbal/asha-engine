package generation2z2airlockclosurefunctornativeminimalitysaturationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE928-Z2-AIRLOCK-CLOSURE-FUNCTOR-NATIVE-MINIMALITY-SATURATION-AUDIT"
	theoremName = "Gate 928: Z2 AirlockClosureFunctor Native Minimality and Saturation Audit"
)

func Generation2Z2AirlockClosureFunctorNativeMinimalitySaturationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 928 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "basepoint axiom forces Cl(0)=F0 under the axiom but is not native", Passed: a.Basepoint.ForcedByAxiom && a.Basepoint.EmptyHasNoActivation && a.Basepoint.MatchesReducedForm && !a.Basepoint.NativeAxiomCertified && containsAll(a.Basepoint.Supports, []string{SupportBasepointForcesCl0, SupportEmptySubsetNoActiveResponse, SupportReducedMatchesBasepointClosure}) && containsAll(a.Basepoint.Failures, []string{FailureBasepointAxiomNotNative}), Detail: FormatBasepoint(a.Basepoint)},
			{Name: "monotonicity is compatible with the boundary chain and airlock flag but is not native", Passed: a.Monotonicity.Monotone && a.Monotonicity.MatchingThreeLevels && !a.Monotonicity.NativeActionTheorem && containsAll(a.Monotonicity.Supports, []string{SupportClosureMustBeMonotone, SupportAirlockFlagThreeLevels}) && containsAll(a.Monotonicity.Failures, []string{FailureMonotonicityNotNative}), Detail: FormatMonotonicity(a.Monotonicity)},
			{Name: "minimal nontrivial closure forces Cl(1)=F1 under the axiom", Passed: a.Minimality.ForcedByMinimality && !a.Minimality.SkipsSaturation && a.Minimality.FirstNonbaseClosure && !a.Minimality.NativeAxiomCertified && containsAll(a.Minimality.Supports, []string{SupportMinimalForcesCl1, SupportSingletonCannotSkipF2, SupportExposureFirstNonbase}) && containsAll(a.Minimality.Failures, []string{FailureMinimalClosureNotNative}), Detail: FormatMinimality(a.Minimality)},
			{Name: "saturation forces Cl(2)=F2 under the axiom", Passed: a.Saturation.ForcedBySaturation && a.Saturation.TopBoundaryDegree && a.Saturation.CannotRemainAtF1 && !a.Saturation.NativeAxiomCertified && containsAll(a.Saturation.Supports, []string{SupportSaturationForcesCl2, SupportTopDegreeClosesF2, SupportFullPairCannotRemainF1}) && containsAll(a.Saturation.Failures, []string{FailureSaturationNotNative}), Detail: FormatSaturation(a.Saturation)},
			{Name: "closure ladder is Z2 representative-independent at class level", Passed: a.Z2.PhaseFlipCommutes && a.Z2.RanksInvariant && a.Z2.RepresentativeFree && !a.Z2.NativeZ2Theorem && containsAll(a.Z2.Supports, []string{SupportClosureZ2Independent, SupportPhaseFlipCommutesClosure, SupportClosureRanksZ2Invariant}) && containsAll(a.Z2.Failures, []string{FailureZ2ClosureNotNative}), Detail: FormatZ2(a.Z2)},
			{Name: "fixed-base quotienting forces cumulative F2/F0 and rejects F2/F1 under the rule", Passed: a.FixedBase.ForcesCumulative && a.FixedBase.RejectsAssociated && !a.FixedBase.NativeRuleCertified && containsAll(a.FixedBase.Supports, []string{SupportFixedBaseForcesCumulative, SupportF2OverF0FromPunctureBase, SupportGradedRejectedByFixedBase}) && containsAll(a.FixedBase.Failures, []string{FailureFixedBaseQuotientNotNative}), Detail: FormatFixedBase(a.FixedBase)},
			{Name: "airlock closure is unique under basepoint-monotone-minimal-saturated-Z2 axioms but native axiom source is missing", Passed: a.Uniqueness.Basepoint && a.Uniqueness.Monotone && a.Uniqueness.Minimal && a.Uniqueness.Saturated && a.Uniqueness.Z2Invariant && a.Uniqueness.UniqueUnderAxioms && !a.Uniqueness.NativeAxiomSource && containsAll(a.Uniqueness.Supports, []string{SupportClosureUniqueUnderAxioms, SupportCl1Cl2Forced, SupportAlternativesFail}) && containsAll(a.Uniqueness.Failures, []string{FailureClosureAxiomsNotNative}), Detail: FormatUniqueness(a.Uniqueness)},
			{Name: "BoundaryActivationMeasure targets are fixed by unique closure candidate but alpha is not native", Passed: a.Measure.UniqueClosureSuppliesTheta && a.Measure.TargetsFixedByClosure && a.Measure.AlphaReconstructed && !a.Measure.NativeAlphaTheorem && a.Measure.ThetaOneRank == RankF1OverF0 && a.Measure.ThetaTwoRank == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportUniqueClosureSuppliesTheta, SupportMeasureTargetsFixed, SupportAlphaViaUniqueClosure}) && containsAll(a.Measure.Failures, []string{FailureAlphaViaClosureNotNative, FailureMuBNotNativeWithoutAxiomSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, BoundarySubsetChain, ExteriorSourceChain, AirlockFlagChain, ClosureFunctor, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, ThetaOne, ThetaTwo, AlphaViaClosure, NextGate, FormatBasepoint(a.Basepoint), FormatMonotonicity(a.Monotonicity), FormatMinimality(a.Minimality), FormatSaturation(a.Saturation), FormatZ2(a.Z2), FormatFixedBase(a.FixedBase), FormatUniqueness(a.Uniqueness), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
