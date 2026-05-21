package generation2airlocksupportclosureoperatorexistenceidempotenceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE930-AIRLOCK-SUPPORT-CLOSURE-OPERATOR-EXISTENCE-IDEMPOTENCE-AUDIT"
	theoremName = "Gate 930: AirlockSupportClosureOperator Existence and Idempotence Audit"
)

func Generation2AirlockSupportClosureOperatorExistenceIdempotenceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 930 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "finite admissible support chain gives bridge-level closure existence", Passed: a.SupportChain.FiniteChain && a.SupportChain.LeastSupportsExist && !a.SupportChain.NativeClosureTheorem && containsAll(a.SupportChain.Supports, []string{SupportClosureExistsOnFiniteFlagChain, SupportEachDemandHasLeastSupport, SupportExistenceFromFiniteChain}) && containsAll(a.SupportChain.Failures, []string{FailureExistenceBridgeNotNative}), Detail: FormatSupportChain(a.SupportChain)},
			{Name: "closure is extensive at demand-support level but not native subspace closure", Passed: a.Extensivity.DemandTyped && a.Extensivity.Cl0ContainsBasepoint && a.Extensivity.Cl1ContainsExposure && a.Extensivity.Cl2ContainsEnclosure && !a.Extensivity.NativeSubspaceClosure && containsAll(a.Extensivity.Supports, []string{SupportClosureExtensiveDemandLevel, SupportCl0ContainsBasepointDemand, SupportCl1ContainsExposureDemand, SupportCl2ContainsFullEnclosureDemand}) && containsAll(a.Extensivity.Failures, []string{FailureExtensivityDemandTypedNotNative}), Detail: FormatExtensivity(a.Extensivity)},
			{Name: "closure is monotone on the candidate chain", Passed: a.Monotonicity.Monotone && a.Monotonicity.SupportInclusion && !a.Monotonicity.NativeActionTheorem && containsAll(a.Monotonicity.Supports, []string{SupportAirlockClosureMonotone, SupportDemandOrderMatchesSupportInclusion, SupportCl0SubsetCl1SubsetCl2}) && containsAll(a.Monotonicity.Failures, []string{FailureMonotonicityCandidateNotNative}), Detail: FormatMonotonicity(a.Monotonicity)},
			{Name: "closure is idempotent on closed admissible supports", Passed: a.Idempotence.ImageInAdmissibleFamily && a.Idempotence.F0Closed && a.Idempotence.F1Closed && a.Idempotence.F2Closed && a.Idempotence.Idempotent && !a.Idempotence.NativeClosureTheorem && containsAll(a.Idempotence.Supports, []string{SupportAirlockClosureIdempotent, SupportF0F1F2ClosedSupports, SupportClosureTwiceNoChange}) && containsAll(a.Idempotence.Failures, []string{FailureIdempotenceCandidateNotNative}), Detail: FormatIdempotence(a.Idempotence)},
			{Name: "minimal non-base demand closes to F1 and rejects jump to F2", Passed: a.Minimality.DemandK == 1 && a.Minimality.ClosureTarget == "F_1" && a.Minimality.LeastNonbaseSupport && a.Minimality.RejectsJumpToF2 && !a.Minimality.NativeRule && containsAll(a.Minimality.Supports, []string{SupportCl1ByLeastNonbaseSupport, SupportSingletonDoesNotJumpToF2, SupportMinimalExposureForced}) && containsAll(a.Minimality.Failures, []string{FailureLeastNonbaseSupportNotNative}), Detail: FormatMinimality(a.Minimality)},
			{Name: "full-pair demand closes to F2 and rejects F1", Passed: a.Saturation.DemandK == 2 && a.Saturation.ClosureTarget == "F_2" && a.Saturation.FullPairSaturation && a.Saturation.RejectsCloseToF1 && !a.Saturation.NativeRule && containsAll(a.Saturation.Supports, []string{SupportCl2ByFullPairSaturation, SupportTopDemandRequiresSaturatedRectangle, SupportFullPairCannotCloseToF1}) && containsAll(a.Saturation.Failures, []string{FailureFullPairSaturationNotNative}), Detail: FormatSaturation(a.Saturation)},
			{Name: "closure descends to Z2 class", Passed: a.Z2.PhaseFlipCommutes && a.Z2.DescendsToZ2Class && !a.Z2.NativePhaseTheorem && containsAll(a.Z2.Supports, []string{SupportAirlockClosureZ2Equivariant, SupportPhaseFlipCommutesWithClosure, SupportClosureDescendsToZ2Class}) && containsAll(a.Z2.Failures, []string{FailureZ2EquivarianceClosureNotNative}), Detail: FormatZ2(a.Z2)},
			{Name: "Theta_B^Z2 is recovered by fixed-base quotient of closure", Passed: a.TargetRecovery.CumulativeTargets && a.TargetRecovery.RejectsAssociatedGraded && !a.TargetRecovery.NativeFixedBaseTheorem && a.TargetRecovery.ThetaOneRank == RankF1OverF0 && a.TargetRecovery.ThetaTwoRank == RankF2OverF0 && containsAll(a.TargetRecovery.Supports, []string{SupportThetaRecoveredFromClosure, SupportFixedBaseQuotientCumulativeTargets, SupportAssociatedGradedRejectedByF0Root}) && containsAll(a.TargetRecovery.Failures, []string{FailureFixedBaseQuotientBridgeNotNative}), Detail: FormatTargetRecovery(a.TargetRecovery)},
			{Name: "BoundaryActivationMeasure is rewritten using closure but native alpha remains blocked", Passed: a.Measure.RewrittenUsingClosure && a.Measure.AlphaReconstructed && !a.Measure.NativeAlpha && a.Measure.ThetaOneRank == RankF1OverF0 && a.Measure.ThetaTwoRank == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportMuBRewrittenUsingClosure, SupportAlphaReconstructedThroughClosure, SupportMuBTargetGapReducedToClosureStatus}) && containsAll(a.Measure.Failures, []string{FailureNoNativeAirlockSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, AdmissibleSupportFamily, Z2AdmissibleSupportFamily, BoundaryDemandChain, ClosureOperatorName, ClosureDefinition, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, MuBViaClosure, AlphaViaClosureOperator, NextGate, FormatSupportChain(a.SupportChain), FormatExtensivity(a.Extensivity), FormatMonotonicity(a.Monotonicity), FormatIdempotence(a.Idempotence), FormatMinimality(a.Minimality), FormatSaturation(a.Saturation), FormatZ2(a.Z2), FormatTargetRecovery(a.TargetRecovery), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
