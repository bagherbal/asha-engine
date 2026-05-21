package generation2z2boundaryalphacrosslaneexclusionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_Z2_BOUNDARYALPHA_CROSSLANE_EXCLUSION_AUDIT"
	theoremName = "Gate 915 — Z2 BoundaryAlpha CrossLane Exclusion Audit"
)

func Generation2Z2BoundaryAlphaCrossLaneExclusionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		checks := []theorem.Check{
			{Name: "Gate 914 selector shape is inherited without reopening phase/socket/representative wounds", Passed: a.Inherited.SelectorShapeSupported && !a.Inherited.ReopensPhaseSign && !a.Inherited.ReopensSocketOrder && !a.Inherited.ReopensRepresentative && !a.Inherited.DerivesAlpha && !a.Inherited.UpdatesOfficialLedger && containsAll(a.Inherited.Supports, []string{StatusGate914Inherited, StatusNoLoopBack}), Detail: FormatInherited(a.Inherited)},
			{Name: "exposure/enclosure type separation marks the two false cross-lanes as type-incompatible", Passed: a.TypeSep.ExcludesByType && !a.TypeSep.NativeTheorem && !a.TypeSep.LinearFalseLane.TypeCompatible && !a.TypeSep.QuadraticFalseLane.TypeCompatible && a.TypeSep.LinearFalseLane.FalseContribution == FalseLinearTerm && a.TypeSep.QuadraticFalseLane.FalseContribution == FalseQuadraticTerm && containsAll(a.TypeSep.Supports, []string{SupportCrossLanesExcludedByType, SupportLambda1ExposureOnly, SupportLambda2EnclosureOnly}) && containsAll(a.TypeSep.Failures, []string{FailureTypeSeparationNotNativeFunctorTheorem}), Detail: FormatTypeSeparation(a.TypeSep)},
			{Name: "degree-indexed selector determinism conditionally blocks false targets by functionhood", Passed: a.Determinism.IsFunction && a.Determinism.ExcludesFalseTargets && !a.Determinism.UniqueNativeSelector && a.Determinism.CorrectLinearTarget == Z2ExposedFaceClass && a.Determinism.CorrectQuadraticTarget == Z2FullEnclosureClass && containsAll(a.Determinism.Supports, []string{SupportCrossLanesExcludedIfFunction, SupportSelectorDeterminismBlocksFalse}) && containsAll(a.Determinism.Failures, []string{FailureNoNativeUniqueDegreeSelector}), Detail: FormatDeterminism(a.Determinism)},
			{Name: "rank contamination check detects the wrong alpha response if cross-lanes are admitted", Passed: near(a.Contamination.CorrectAlpha, AlphaB) && a.Contamination.PollutedAlpha > a.Contamination.CorrectAlpha && a.Contamination.FalseDelta > 0 && a.Contamination.MismatchDetected && !a.Contamination.NativeExclusion && near(a.Contamination.PollutedCoefficient, float64(143)/float64(360)) && containsAll(a.Contamination.FalseTerms, []string{FalseLinearTerm, FalseQuadraticTerm}) && containsAll(a.Contamination.Supports, []string{SupportFalseCrossLanesWrongAlpha, SupportActiveAlphaRequiresExclusion}) && containsAll(a.Contamination.Failures, []string{FailureNumericalMismatchNotNativeExclusion}), Detail: FormatContamination(a.Contamination)},
			{Name: "cross-lane exclusion remains compatible with cumulative enclosure F2 over F0 and rejects F2 over F1", Passed: a.Cumulative.F2OverF0Rank == RankF2OverF0 && a.Cumulative.F2OverF1Rank == RankF2OverF1 && a.Cumulative.KeepsCumulativeEnclosure && a.Cumulative.RejectsAssociatedGradedSlice && !a.Cumulative.NativeReasonForChoice && containsAll(a.Cumulative.Supports, []string{SupportCumulativeCompatible, SupportDegreeTwoRemainsF2OverF0}) && containsAll(a.Cumulative.Failures, []string{FailureNoNativeCumulativeReason}), Detail: FormatCumulative(a.Cumulative)},
			{Name: "Z2 quotient compatibility maps correct lanes to correct lanes and false lanes to false lanes without proving native exclusion", Passed: a.Z2.CorrectLanesRepresentativeFree && a.Z2.FalseLanesRepresentativeFree && a.Z2.CorrectMapToCorrect && a.Z2.FalseMapToFalse && !a.Z2.NativeExclusionTheorem && containsAll(a.Z2.Supports, []string{SupportCrossLaneExclusionZ2Compatible, SupportFalseLanesRepresentativeFreeFalse}) && containsAll(a.Z2.Failures, []string{FailureZ2CompatibilityNotNativeExclusion}), Detail: FormatZ2(a.Z2)},
			{Name: "native cross-lane theorem, alpha, R3, full A_F descent, and generation/flavor/Yukawa firewalls remain preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeZ2CrossLaneTheorem, FailureTypeSeparationNotNativeFunctorTheorem, FailureNoNativeUniqueDegreeSelector, FailureNumericalMismatchNotNativeExclusion, FailureNoNativeCumulativeReason, FailureZ2CompatibilityNotNativeExclusion, FailureAlphaStillSealed, FailureDenominatorsSTransportStillExternal, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatTypeSeparation(a.TypeSep), FormatDeterminism(a.Determinism), FormatContamination(a.Contamination), FormatCumulative(a.Cumulative), FormatZ2(a.Z2), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
