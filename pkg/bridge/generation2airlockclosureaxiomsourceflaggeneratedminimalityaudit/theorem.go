package generation2airlockclosureaxiomsourceflaggeneratedminimalityaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE929-AIRLOCK-CLOSURE-AXIOM-SOURCE-FLAG-GENERATED-MINIMALITY-AUDIT"
	theoremName = "Gate 929: AirlockClosure Axiom Source and Flag-Generated Minimality Audit"
)

func Generation2AirlockClosureAxiomSourceFlagGeneratedMinimalityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 929 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "basepoint axiom is sourced by puncture initiality but not native closure operator", Passed: a.Basepoint.PunctureInitiality && a.Basepoint.EmptyBoundarySubset && a.Basepoint.FlagGenerated && !a.Basepoint.NativeClosureTheorem && a.Basepoint.ClosureTarget == "F_0" && containsAll(a.Basepoint.Supports, []string{SupportBasepointByPunctureInitiality, SupportEmptyClosesMinimalSupport, SupportCl0FlagGenerated}) && containsAll(a.Basepoint.Failures, []string{FailurePunctureInitialityNotNative}), Detail: FormatBasepoint(a.Basepoint)},
			{Name: "monotonicity is sourced by support inclusion but not native boundary action", Passed: a.Monotonicity.SupportInclusion && a.Monotonicity.OrdersCompatible && a.Monotonicity.FlagNatural && !a.Monotonicity.NativeActionTheorem && containsAll(a.Monotonicity.Supports, []string{SupportMonotonicityBySupportInclusion, SupportBoundaryAirlockOrdersCompatible, SupportClosureMonotonicityFlagNatural}) && containsAll(a.Monotonicity.Failures, []string{FailureSupportInclusionNotNative}), Detail: FormatMonotonicity(a.Monotonicity)},
			{Name: "minimality is sourced by least same-socket completion but not native", Passed: a.Minimality.SingletonActivation && a.Minimality.ExposedFace && a.Minimality.ForcedByMinimalSupport && !a.Minimality.NativeClosureTheorem && a.Minimality.ClosureTarget == "F_1" && containsAll(a.Minimality.Supports, []string{SupportMinimalityBySameSocketCompletion, SupportSingletonClosesExposedFace, SupportCl1ForcedByMinimalSupport}) && containsAll(a.Minimality.Failures, []string{FailureLeastSameSocketNotNative}), Detail: FormatMinimality(a.Minimality)},
			{Name: "saturation is sourced by full boundary-pair activation but not native", Passed: a.Saturation.FullPairActivation && a.Saturation.TopExteriorDegree && a.Saturation.ForcedBySaturation && !a.Saturation.NativeClosureTheorem && a.Saturation.ClosureTarget == "F_2" && containsAll(a.Saturation.Supports, []string{SupportSaturationByFullPairActivation, SupportTopDegreeClosesFullRectangle, SupportCl2ForcedBySaturatedCompletion}) && containsAll(a.Saturation.Failures, []string{FailureFullRightRectangleNotNative}), Detail: FormatSaturation(a.Saturation)},
			{Name: "fixed-base quotient is sourced by relative activation above puncture", Passed: a.FixedBase.UsesFixedBaseF0 && a.FixedBase.CumulativeQuotient && a.FixedBase.RejectsAssociatedGraded && !a.FixedBase.NativeMeasureTheorem && containsAll(a.FixedBase.Supports, []string{SupportFixedBaseByRelativeActivation, SupportCumulativeQuotientFlagGenerated, SupportAssociatedGradedRejectedByBasepoint}) && containsAll(a.FixedBase.Failures, []string{FailureRelativeActivationQuotientNotNative}), Detail: FormatFixedBase(a.FixedBase)},
			{Name: "Z2 invariance is sourced by class-level support closure", Passed: a.Z2.PhaseFlipExchanges && a.Z2.ClassLevelClosure && a.Z2.RanksInvariant && !a.Z2.NativeGlobalPhaseTheorem && containsAll(a.Z2.Supports, []string{SupportZ2ByClassLevelClosure, SupportPhaseFlipExchangesRepresentatives, SupportClosureQuotientRanksZ2Invariant}) && containsAll(a.Z2.Failures, []string{FailureZ2ClassClosureNotNative}), Detail: FormatZ2(a.Z2)},
			{Name: "closure axiom ledger is flag-generated but native support closure operator is missing", Passed: a.Ledger.FlagGenerated && !a.Ledger.NativeOperatorExists && containsAll(a.Ledger.Supports, []string{SupportAxiomsFlagGenerated}) && containsAll(a.Ledger.Failures, []string{FailureNoNativeAirlockSupportClosure, FailureAxiomsFlagSourcedNotNative}), Detail: FormatLedger(a.Ledger)},
			{Name: "BoundaryActivationMeasure targets are fixed by flag-generated closure but alpha remains bridge candidate", Passed: a.Measure.ThetaReconstructed && a.Measure.TargetsFixed && !a.Measure.NativeAlpha && a.Measure.ThetaOneRank == RankF1OverF0 && a.Measure.ThetaTwoRank == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportThetaReconstructedFromFlagClosure, SupportMeasureTargetsFixedByFlagClosure}) && containsAll(a.Measure.Failures, []string{FailureMuBNotNativeWithoutSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.Inherited, BoundarySubsetChain, AirlockFlagChain, Z2PunctureClass, CandidateOperator, LeastSupportRule, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, AlphaViaClosure, NextGate, FormatBasepoint(a.Basepoint), FormatMonotonicity(a.Monotonicity), FormatMinimality(a.Minimality), FormatSaturation(a.Saturation), FormatFixedBase(a.FixedBase), FormatZ2(a.Z2), FormatLedger(a.Ledger), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
