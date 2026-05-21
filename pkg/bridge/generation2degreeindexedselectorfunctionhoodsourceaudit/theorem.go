package generation2degreeindexedselectorfunctionhoodsourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE923-DEGREEINDEXED-SELECTOR-FUNCTIONHOOD-SOURCE-AUDIT"
	theoremName = "Gate 923: DegreeIndexed Selector Functionhood Source Audit"
)

func Generation2DegreeIndexedSelectorFunctionhoodSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 923 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "degree one is source-typed as single-boundary exposure targeting exposed face quotient", Passed: a.DegreeOne.Degree == 1 && a.DegreeOne.Rank == RankExposureFace && a.DegreeOne.SourceStatus == SourceBridgeTypedNotNative && a.DegreeOne.BridgeTyped && !a.DegreeOne.Native && containsAll(a.DegreeOne.Supports, []string{SupportDegreeOneExposureSource, SupportSingleExposureTargetsFaceQuotient, SupportIBZ2OneExposedFace}) && containsAll(a.DegreeOne.Failures, []string{FailureExposureToF1NotNative}), Detail: FormatLane(a.DegreeOne)},
			{Name: "degree two is source-typed as full boundary-pair enclosure targeting full puncture complement", Passed: a.DegreeTwo.Degree == 2 && a.DegreeTwo.Rank == RankFullEnclosure && a.DegreeTwo.SourceStatus == SourceBridgeTypedNotNative && a.DegreeTwo.BridgeTyped && !a.DegreeTwo.Native && containsAll(a.DegreeTwo.Supports, []string{SupportDegreeTwoEnclosureSource, SupportFullEnclosureTargetsComplement, SupportIBZ2TwoFullEnclosure}) && containsAll(a.DegreeTwo.Failures, []string{FailureEnclosureToF2NotNative}), Detail: FormatLane(a.DegreeTwo)},
			{Name: "degree two selects cumulative enclosure F2/F0 and rejects associated graded F2/F1", Passed: a.CumulativeEnclosure.CumulativeRequired && !a.CumulativeEnclosure.Native && a.CumulativeEnclosure.RankF2OverF0 == RankFullEnclosure && a.CumulativeEnclosure.RankF2OverF1 == RankAssociatedSlice && containsAll(a.CumulativeEnclosure.Supports, []string{SupportDegreeTwoCumulativeNotGraded, SupportFullPairRequiresF2OverF0, SupportF2OverF1Rejected}) && containsAll(a.CumulativeEnclosure.Failures, []string{FailureNoNativeCumulativeTheorem}), Detail: FormatCumulative(a.CumulativeEnclosure)},
			{Name: "exposure/enclosure typing supports selector functionhood but not native theorem", Passed: a.Functionhood.ExposureEnclosureAccepted && a.Functionhood.FunctionalIfExposureAccepted && !a.Functionhood.NativeFunctionhood && a.Functionhood.PrimaryGap == PrimaryGapExposureFunctor && containsAll(a.Functionhood.Supports, []string{SupportExposureEnclosureGivesFunctionhood, SupportEachDegreeUniqueTargetType, SupportIBZ2FunctionalIfAccepted}) && containsAll(a.Functionhood.Failures, []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureSelectorDependsOnBridgeRule}), Detail: FormatFunctionhood(a.Functionhood)},
			{Name: "cross-lane exclusion follows from selector functionhood but remains non-native", Passed: a.CrossLane.ExcludedIfFunctional && !a.CrossLane.NativeExclusion && containsAll(a.CrossLane.Supports, []string{SupportCrossLaneFollowsFunctionhood, SupportFalseTermsBlocked}) && containsAll(a.CrossLane.Failures, []string{FailureCrossLaneNotNativeWithoutSelector}), Detail: FormatCrossLane(a.CrossLane)},
			{Name: "selector functionhood is compatible with Z2 orientation class", Passed: a.Z2Compatibility.CommutesWithPhaseFlip && a.Z2Compatibility.RanksRepresentativeFree && !a.Z2Compatibility.NativeZ2Selector && containsAll(a.Z2Compatibility.Supports, []string{SupportSelectorFunctionhoodZ2Compatible, SupportIBZ2CommutesWithPhaseFlip, SupportSelectorRanksRepresentativeFree}) && containsAll(a.Z2Compatibility.Failures, []string{FailureZ2CompatibilityNotNativeSelector}), Detail: FormatZ2(a.Z2Compatibility)},
			{Name: "selector supplies BoundaryActivationMeasure target ranks while mu_B remains non-native", Passed: a.MuB.Rank1 == RankExposureFace && a.MuB.Rank2 == RankFullEnclosure && a.MuB.RankH10 == RankH10 && a.MuB.RankH72 == RankH72 && near(a.MuB.LinearContribution, AlphaLinear) && near(a.MuB.QuadraticContribution, AlphaQuad) && near(a.MuB.Alpha, AlphaB) && !a.MuB.NativeMuB && containsAll(a.MuB.Supports, []string{SupportSelectorSuppliesMuBRanks, SupportMuBNativeGapReduced}) && containsAll(a.MuB.Failures, []string{FailureMuBStillNotNativeWithoutSelector, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMuB(a.MuB) + " | " + FormatFirewalls(a.Firewalls)},
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

		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.InheritedStatus, StrategicConclusion, NextGate, PunctureClass, BoundaryResponse, SelectorFormula, SelectorOneFormula, SelectorTwoFormula, MuBFormula, BoundaryAlphaFormula, ExposureEnclosureRule, FormatLane(a.DegreeOne), FormatLane(a.DegreeTwo), FormatCumulative(a.CumulativeEnclosure), FormatFunctionhood(a.Functionhood), FormatCrossLane(a.CrossLane), FormatZ2(a.Z2Compatibility), FormatMuB(a.MuB), FormatFirewalls(a.Firewalls), a.Final, PrimaryGapExposureFunctor, SourceBridgeTypedNotNative, SourceDependent, SourceZ2Compatible}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
