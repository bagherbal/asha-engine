package generation2boundarydegreetoairlockflagtargetfunctornaturalityanduniquenessaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE926-BOUNDARYDEGREE-TO-AIRLOCKFLAG-TARGETFUNCTOR-NATURALITY-UNIQUENESS-AUDIT"
	theoremName = "Gate 926: BoundaryDegree-to-AirlockFlag TargetFunctor Naturality and Uniqueness Audit"
)

func Generation2BoundaryDegreeToAirlockFlagTargetFunctorNaturalityAndUniquenessAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 926 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}
		checks := []theorem.Check{
			{Name: "order-preserving selector is unique between two active two-level chains under constraints", Passed: a.Order.UniqueUnderOrder && a.Order.OrderPreserving && a.Order.SwappedOrderReversing && !a.Order.NativeOrderTheorem && containsAll(a.Order.Supports, []string{SupportOrderPreservingUnique, SupportSwappedOrderReversing, SupportThetaUniqueOrderPreserving}) && containsAll(a.Order.Failures, []string{FailureOrderPreservationNotNative}), Detail: FormatOrder(a.Order)},
			{Name: "exposure/enclosure typing forces Theta assignment under bridge constraints", Passed: a.Types.ExposureUnique && a.Types.EnclosureUnique && a.Types.ForcesThetaAssignment && !a.Types.NativeTypeTheorem && containsAll(a.Types.Supports, []string{SupportExposureTargetsF1, SupportEnclosureTargetsF2, SupportExposureEnclosureForcesTheta}) && containsAll(a.Types.Failures, []string{FailureExposureEnclosureUniquenessNative}), Detail: FormatTypes(a.Types)},
			{Name: "Theta_B^Z2 is representative-independent at Z2 class level", Passed: a.Z2.RepresentativeIndependent && a.Z2.CommutesWithFlip && !a.Z2.NativeZ2Theorem && a.Z2.RankOne == RankF1OverF0 && a.Z2.RankTwo == RankF2OverF0 && containsAll(a.Z2.Supports, []string{SupportThetaRepresentativeIndependent, SupportThetaCommutesZ2, SupportThetaRanksZ2Invariant}) && containsAll(a.Z2.Failures, []string{FailureZ2IndependenceNotNative}), Detail: FormatZ2(a.Z2)},
			{Name: "associated-graded target F2/F1 is rejected by cumulative enclosure type and alpha rank", Passed: a.Graded.FailsCumulativeType && a.Graded.FailsAlphaRank && a.Graded.RejectedByTypeAndRank && !a.Graded.NativeCumulativeTheorem && a.Graded.AlternativeRank == RankF2OverF1 && a.Graded.RequiredRank == RankF2OverF0 && containsAll(a.Graded.Supports, []string{SupportAssociatedGradedFailsType, SupportAssociatedGradedFailsRank, SupportF2OverF1RejectedByTypeAndRank}) && containsAll(a.Graded.Failures, []string{FailureF2OverF1RejectionNotNative}), Detail: FormatGraded(a.Graded)},
			{Name: "constant and cross-lane alternatives fail reduction, type, and alpha-shape constraints", Passed: a.Alternatives.DegreeZeroAbsent && a.Alternatives.CrossLaneExposureToEnclose && a.Alternatives.CrossLaneEnclosureToExpose && a.Alternatives.ViolatesType && a.Alternatives.FailsAlphaShape && !a.Alternatives.NativeUniquenessTheorem && containsAll(a.Alternatives.Supports, []string{SupportDegreeZeroAbsentByReduction, SupportCrossLaneViolatesType, SupportCrossLaneFailsAlphaShape}) && containsAll(a.Alternatives.Failures, []string{FailureAlternativeRejectionNotNative}), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "unique Theta fixes mu_B target ranks but native measure remains blocked", Passed: a.Measure.SelectorFunctionhood && a.Measure.CrossLaneExclusion && a.Measure.TargetRanksFixed && !a.Measure.NativeMeasure && a.Measure.ThetaRankOne == RankF1OverF0 && a.Measure.ThetaRankTwo == RankF2OverF0 && containsAll(a.Measure.Supports, []string{SupportThetaUniquenessStrengthensMuB, SupportSelectorAndCrossLaneFollowUnique, SupportMuBTargetRanksFixedByUniqueTheta}) && containsAll(a.Measure.Failures, []string{FailureMuBNotNativeWithoutThetaSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.InheritedStatus, SourceChain, TargetChain, ThetaFunctor, ThetaOne, ThetaTwo, SwappedAssignment, AssociatedGradedTarget, ExposureType, EnclosureType, MeasureFormula, AlphaFormula, NextGate, FormatOrder(a.Order), FormatTypes(a.Types), FormatZ2(a.Z2), FormatGraded(a.Graded), FormatAlternatives(a.Alternatives), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
