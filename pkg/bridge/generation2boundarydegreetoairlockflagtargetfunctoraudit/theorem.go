package generation2boundarydegreetoairlockflagtargetfunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE925-BOUNDARYDEGREE-TO-AIRLOCKFLAG-TARGETFUNCTOR-AUDIT"
	theoremName = "Gate 925: BoundaryDegree-to-AirlockFlag TargetFunctor Audit"
)

func Generation2BoundaryDegreeToAirlockFlagTargetFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 925 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "source boundary-degree chain and target airlock flag chain have matching two-level order type", Passed: a.ChainMatch.SourceLevels == 2 && a.ChainMatch.TargetLevels == 2 && a.ChainMatch.MatchingOrderType && !a.ChainMatch.NativeTargetFunctor && containsAll(a.ChainMatch.Supports, []string{SupportSourceDegreesTwoLevel, SupportTargetFlagTwoNonbaseLevels, SupportDegreeAndFlagMatchingOrderType}) && containsAll(a.ChainMatch.Failures, []string{FailureOrderTypeNotNative}), Detail: FormatChain(a.ChainMatch)},
			{Name: "degree-one exposure maps to first non-base Z2 airlock flag quotient at shape level", Passed: a.ExposureTarget.Degree == 1 && a.ExposureTarget.Target == "[F_1/F_0]_{Z2}" && a.ExposureTarget.Rank == RankF1OverF0 && a.ExposureTarget.Z2ClassLevel && !a.ExposureTarget.NativeRule && containsAll(a.ExposureTarget.Supports, []string{SupportExposureTargetsFirstQuotient, SupportDegreeOneMapsByMinimalExposure, SupportThetaOneExposedFace}) && containsAll(a.ExposureTarget.Failures, []string{FailureMinimalExposureLevelNotNative}), Detail: FormatTarget(a.ExposureTarget)},
			{Name: "degree-two enclosure maps to cumulative full Z2 airlock flag quotient at shape level", Passed: a.EnclosureTarget.Degree == 2 && a.EnclosureTarget.Target == "[F_2/F_0]_{Z2}" && a.EnclosureTarget.Rank == RankF2OverF0 && a.EnclosureTarget.Z2ClassLevel && !a.EnclosureTarget.NativeRule && containsAll(a.EnclosureTarget.Supports, []string{SupportEnclosureTargetsFullQuotient, SupportDegreeTwoMapsByFullEnclosure, SupportThetaTwoFullEnclosure}) && containsAll(a.EnclosureTarget.Failures, []string{FailureFullEnclosureLevelNotNative}), Detail: FormatTarget(a.EnclosureTarget)},
			{Name: "associated-graded slice F2/F1 remains rejected for top-degree alpha target", Passed: a.Graded.TopDegreeCumulative && !a.Graded.NativeCumulativeRule && a.Graded.RejectedRank == RankF2OverF1 && a.Graded.CumulativeRank == RankF2OverF0 && containsAll(a.Graded.Supports, []string{SupportAssociatedGradedRejected, SupportTopDegreeSelectsCumulative, SupportFullPairActivationRequiresF2OverF0}) && containsAll(a.Graded.Failures, []string{FailureCumulativeOverGradedNotNative}), Detail: FormatGraded(a.Graded)},
			{Name: "Theta_B^Z2 target-functor shape is order-preserving, Z2-class compatible, and not native", Passed: a.Theta.OrderPreserving && a.Theta.Z2RepresentativeIndependent && a.Theta.ExposureEnclosureTyped && a.Theta.CumulativeTopDegree && !a.Theta.NativeFunctor && containsAll(a.Theta.Supports, []string{SupportThetaShapeDefined, SupportThetaOrderPreserving, SupportThetaZ2RepresentativeIndependent, SupportThetaExposureEnclosureTyped, SupportThetaCumulativeTopDegree}) && containsAll(a.Theta.Failures, []string{FailureThetaShapeNotNative}), Detail: FormatTheta(a.Theta)},
			{Name: "Theta_B^Z2 supplies selector functionhood and cross-lane exclusion only conditionally", Passed: a.Selector.IBEqualsTheta && a.Selector.CrossLaneExcluded && !a.Selector.NativeSelectorFunctionhood && containsAll(a.Selector.Supports, []string{SupportThetaSuppliesSelectorFunctionhood, SupportIBZ2EqualsTheta, SupportCrossLaneFollowsFromTheta}) && containsAll(a.Selector.Failures, []string{FailureSelectorNonNativeWithoutTheta}), Detail: FormatSelector(a.Selector)},
			{Name: "Theta_B^Z2 supplies BoundaryActivationMeasure target ranks while mu_B remains non-native", Passed: a.Measure.TargetRanksSupplied && !a.Measure.NativeMeasure && a.Measure.ThetaRankOne == RankF1OverF0 && a.Measure.ThetaRankTwo == RankF2OverF0 && a.Measure.H10Rank == RankH10 && a.Measure.H72Rank == RankH72 && containsAll(a.Measure.Supports, []string{SupportThetaSuppliesMuBRanks, SupportAlphaReconstructedGivenTheta, SupportMuBNativeGapReducedToTheta}) && containsAll(a.Measure.Failures, []string{FailureMuBNotNativeWithoutTheta, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMeasure(a.Measure) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.InheritedStatus, BoundaryDegreeChain, AirlockFlagChain, Z2PunctureClass, ThetaFunctor, ThetaOne, ThetaTwo, ExposureTarget, EnclosureTarget, MeasureFormula, AlphaFormula, NextGate, FormatChain(a.ChainMatch), FormatTarget(a.ExposureTarget), FormatTarget(a.EnclosureTarget), FormatGraded(a.Graded), FormatTheta(a.Theta), FormatSelector(a.Selector), FormatMeasure(a.Measure), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
