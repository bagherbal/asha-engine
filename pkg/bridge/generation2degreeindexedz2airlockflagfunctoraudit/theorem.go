package generation2degreeindexedz2airlockflagfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_DEGREE_INDEXED_Z2_AIRLOCK_FLAGFUNCTOR_AUDIT"
	theoremName = "Gate 914 — DegreeIndexed Z2 Airlock FlagFunctor Audit"
)

func Generation2DegreeIndexedZ2AirlockFlagFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		checks := []theorem.Check{
			{Name: "Gate 913 reduced B2 response shape is inherited without reopening closed phase/socket/representative wounds", Passed: a.Inherited.ReducedShapeCertified && !a.Inherited.ReopensPhaseSign && !a.Inherited.ReopensSocketOrder && !a.Inherited.ReopensRepresentative && !a.Inherited.DerivesAlpha && !a.Inherited.UpdatesOfficialLedger && containsAll(a.Inherited.Supports, []string{StatusGate913Inherited, StatusNoLoopBack}), Detail: FormatInherited(a.Inherited)},
			{Name: "degree-one boundary response targets the Z2 exposed-face class as selector shape", Passed: a.DegreeOne.Degree == 1 && a.DegreeOne.BoundaryTerm == DegreeOneTerm && a.DegreeOne.Target.Rank == RankF1OverF0 && a.DegreeOne.Target.RepresentativeFree && a.DegreeOne.Selector && !a.DegreeOne.LinearSurjection && !a.DegreeOne.NativeMap && containsAll(a.DegreeOne.Supports, []string{StatusDegreeOneTarget, SupportDegreeOneTargetsExposed, SupportLambda1SingleExposure, SupportExposedFaceRankThree}) && containsAll(a.DegreeOne.Failures, []string{FailureNoNativeLambda1ExposedMap, FailureNoNativeDegreeToZ2FlagFunctor}), Detail: FormatDegreeTarget(a.DegreeOne)},
			{Name: "degree-two boundary response targets the Z2 full-enclosure class as selector shape", Passed: a.DegreeTwo.Degree == 2 && a.DegreeTwo.BoundaryTerm == DegreeTwoTerm && a.DegreeTwo.Target.Rank == RankF2OverF0 && a.DegreeTwo.Target.RepresentativeFree && a.DegreeTwo.Selector && !a.DegreeTwo.LinearSurjection && !a.DegreeTwo.NativeMap && containsAll(a.DegreeTwo.Supports, []string{StatusDegreeTwoTarget, SupportDegreeTwoTargetsFull, SupportLambda2FullEnclosure, SupportFullEnclosureRankSeven}) && containsAll(a.DegreeTwo.Failures, []string{FailureNoNativeLambda2FullMap, FailureNoNativeDegreeToZ2FlagFunctor}), Detail: FormatDegreeTarget(a.DegreeTwo)},
			{Name: "degree-to-flag object is typed as selector, not vector-space surjection", Passed: a.Typing.Lambda1Dim == Lambda1Dim && a.Typing.Lambda2Dim == Lambda2Dim && a.Typing.ExposedRank == RankF1OverF0 && a.Typing.FullRank == RankF2OverF0 && a.Typing.DimensionMismatch && a.Typing.SelectorNotSurjection && containsAll(a.Typing.Supports, []string{StatusSelectorNotSurjection, SupportSelectorNotLinearSurjection, SupportDimensionMismatchSelector}) && containsAll(a.Typing.Failures, []string{FailureLambdaKB2NotSurjection}), Detail: FormatTyping(a.Typing)},
			{Name: "degree two selects cumulative enclosure F2 over F0 and rejects associated-graded F2 over F1 for alpha target", Passed: a.Cumulative.F2OverF0Rank == RankF2OverF0 && a.Cumulative.F2OverF1Rank == RankF2OverF1 && a.Cumulative.SelectsCumulativeEnclosure && a.Cumulative.RejectsAssociatedGradedSlice && !a.Cumulative.NativeReasonForChoice && containsAll(a.Cumulative.Supports, []string{StatusCumulativeChoice, StatusAssociatedGradedReject, SupportDegreeTwoCumulative, SupportAssociatedGradedRejected}) && containsAll(a.Cumulative.Failures, []string{FailureNoNativeCumulativeReason}), Detail: FormatCumulative(a.Cumulative)},
			{Name: "selected Z2 flag classes reconstruct the alpha rank pair under seal while denominators and S-transport remain external", Passed: a.AlphaRanks.ReconstructsRankPair && near(a.AlphaRanks.Alpha, AlphaB) && a.AlphaRanks.RankPair == [2]int{RankF1OverF0, RankF2OverF0} && a.AlphaRanks.Denominators == [2]int{LinearDenom, QuadraticDenom} && !a.AlphaRanks.NativeAlphaSource && a.AlphaRanks.DenominatorsExternal && a.AlphaRanks.STransportExternal && containsAll(a.AlphaRanks.Supports, []string{StatusAlphaRankReconstruct, SupportSelectorReconstructsRankPair, SupportRankPairFromSelectedClasses}) && containsAll(a.AlphaRanks.Failures, []string{FailureSelectorRanksNotAlphaSource, FailureDenominatorsSTransportExternal, FailureAlphaStillSealed}), Detail: FormatAlpha(a.AlphaRanks)},
			{Name: "cross-lane exclusion is conditional on certified selector and not independently native yet", Passed: containsAll(a.CrossLane.ForbiddenLanes, []string{ForbiddenLinearFull, ForbiddenQuadraticFace}) && containsAll(a.CrossLane.FalseTerms, []string{FalseLinearFullTerm, FalseQuadraticFaceTerm}) && a.CrossLane.WouldFollowFromSelector && !a.CrossLane.IndependentNativeTheorem && !a.CrossLane.ProvesCrossLaneExclusion && containsAll(a.CrossLane.Supports, []string{StatusCrossLaneConditional, SupportCrossLaneWouldFollow}) && containsAll(a.CrossLane.Failures, []string{FailureNoIndependentCrossLane}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "native functor, native alpha, native R3, full A_F descent, generation/flavor/Yukawa firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeLambda1ExposedMap, FailureNoNativeLambda2FullMap, FailureLambdaKB2NotSurjection, FailureNoNativeCumulativeReason, FailureNoIndependentCrossLane, FailureSelectorRanksNotAlphaSource, FailureDenominatorsSTransportExternal, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatDegreeTarget(a.DegreeOne), FormatDegreeTarget(a.DegreeTwo), FormatTyping(a.Typing), FormatCumulative(a.Cumulative), FormatAlpha(a.AlphaRanks), FormatCrossLane(a.CrossLane), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
