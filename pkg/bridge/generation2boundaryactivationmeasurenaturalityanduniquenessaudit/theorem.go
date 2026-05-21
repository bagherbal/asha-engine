package generation2boundaryactivationmeasurenaturalityanduniquenessaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE921-BOUNDARYACTIVATIONMEASURE-NATURALITY-UNIQUENESS-AUDIT"
	theoremName = "Gate 921: BoundaryActivationMeasure Naturality and Uniqueness Audit"
)

func Generation2BoundaryActivationMeasureNaturalityAndUniquenessAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 921 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "mu_B is natural on the reduced active B2 response, not arbitrary polynomial input", Passed: a.Domain.InheritedStatus == Gate920ShortStatus && a.Domain.Domain == ReducedResponse && len(a.Domain.ActiveDegrees) == 2 && a.Domain.ActiveDegrees[0] == 1 && a.Domain.ActiveDegrees[1] == 2 && !a.Domain.ArbitraryPolynomial && a.Domain.NaturalOnReducedB2 && !a.Domain.NativeTheorem && containsAll(a.Domain.Supports, []string{SupportMuBNaturalReducedResponse, SupportMuBNoArbitraryPolynomial}) && containsAll(a.Domain.Failures, []string{FailureDomainNaturalityNotNative}), Detail: FormatDomain(a.Domain)},
			{Name: "basepoint reduction is the unique route if alpha has no constant response", Passed: a.Basepoint.UnreducedHasLambda0 && a.Basepoint.ReducedResponseUsed && near(a.Basepoint.AlphaConstantTerm, 0) && a.Basepoint.NoConstantTermForced && a.Basepoint.UniqueIfNoConstantAlpha && !a.Basepoint.NativeTheorem && containsAll(a.Basepoint.Supports, []string{SupportReducedForcesNoConstant, SupportBasepointRemovalUnique}) && containsAll(a.Basepoint.Failures, []string{FailureNoNativeBasepointReduction}), Detail: FormatBasepoint(a.Basepoint)},
			{Name: "degree naturality uniquely assigns S_split^k to exterior degree k", Passed: a.Degree.DegreePowers[1] == 1 && a.Degree.DegreePowers[2] == 2 && near(a.Degree.Coefficients[1], SBoundary) && near(a.Degree.Coefficients[2], SBoundary*SBoundary) && a.Degree.PowerAssignmentUnique && !a.Degree.NativeTheorem && containsAll(a.Degree.Supports, []string{SupportDegreeForcesSPowerK, SupportSPowerAssignmentUnique}) && containsAll(a.Degree.Failures, []string{FailureNoNativeDegreeRespectingMeasure}), Detail: FormatDegree(a.Degree)},
			{Name: "selector functionhood forces one Z2 target per degree and absorbs cross-lane exclusion", Passed: a.Selector.Targets[1] == "[F_1/F_0]_{Z2}" && a.Selector.Targets[2] == "[F_2/F_0]_{Z2}" && a.Selector.FalseTargets[1] == "[F_2/F_0]_{Z2}" && a.Selector.FalseTargets[2] == "[F_1/F_0]_{Z2}" && a.Selector.Ranks[1] == RankI1 && a.Selector.Ranks[2] == RankI2 && a.Selector.FunctionhoodAssumed && a.Selector.UniquePerDegree && a.Selector.CrossLanesExcluded && !a.Selector.NativeTheorem && containsAll(a.Selector.Supports, []string{SupportSelectorFunctionhoodUnique, SupportCrossLaneFromSelectorUniqueness}) && containsAll(a.Selector.Failures, []string{FailureNoNativeUniqueSelector, FailureNoNativeSelectorFunctionhood}), Detail: FormatSelector(a.Selector)},
			{Name: "lane locality forces H10/H72 normalization under the current bridge constraints", Passed: a.Normalization.Chambers[1] == "H_10=H_R^ambient+B_2" && a.Normalization.Chambers[2] == "H_72=Lambda^4 V_8+B_2" && a.Normalization.Ranks[1] == RankH10 && a.Normalization.Ranks[2] == RankH72 && nearLoose(a.Normalization.Weights[1], 0.3) && nearLoose(a.Normalization.Weights[2], float64(7)/72) && a.Normalization.LaneLocalityAccepted && a.Normalization.UniqueGivenLocalGlobal && !a.Normalization.NativeTheorem && containsAll(a.Normalization.Supports, []string{SupportLaneLocalityForcesChambers, SupportNormalizationUniqueGivenChambers}) && containsAll(a.Normalization.Failures, []string{FailureNoNativeLaneLocalityToChamber, FailureNoNativeChamberNormalization}), Detail: FormatNormalization(a.Normalization)},
			{Name: "mu_B is representative-independent on the Z2 airlock class", Passed: a.Z2.RepresentativesExchanged && a.Z2.RanksInvariant && a.Z2.MeasureInvariant && !a.Z2.NativeTheorem && containsAll(a.Z2.Supports, []string{SupportMuBZ2RepresentativeIndependent, SupportPhaseSignNoChangeMeasure}) && containsAll(a.Z2.Failures, []string{FailureZ2InvarianceNotNative}), Detail: FormatZ2(a.Z2)},
			{Name: "standard alternative measures fail the current naturality constraints", Passed: a.Alternatives.UnreducedRejected && a.Alternatives.CrossLaneRejected && a.Alternatives.BareChamberRejected && a.Alternatives.CommonDenominatorRejected && nearLoose(a.Alternatives.PollutedLinearWeight, float64(143)/360) && nearLoose(a.Alternatives.PollutedQuadraticWeight, float64(143)/360) && nearLoose(a.Alternatives.BareLinearWeight, float64(3)/8) && nearLoose(a.Alternatives.BareQuadraticWeight, float64(7)/70) && a.Alternatives.UniqueAmongTested && !a.Alternatives.FullNativeUniqueness && containsAll(a.Alternatives.Supports, []string{SupportAlternativeMeasuresFailConstraints, SupportMuBUniqueAmongTestedMeasures}) && containsAll(a.Alternatives.Failures, []string{FailureUnreducedConstantTerm, FailureCrossLaneAddsFalseTerms, FailureBareChamberBreaksAugmentation, FailureCommonDenominatorBreaksLocality, FailureAlternativeRejectionNotFullNative}), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "alpha reconstruction remains the same while native alpha and native R3 stay blocked", Passed: a.Alpha.Formula == MeasureFormula && a.Alpha.BoundaryFormula == BoundaryAlphaFormula && near(a.Alpha.LinearContribution, AlphaLinear) && near(a.Alpha.QuadraticContribution, AlphaQuad) && near(a.Alpha.Alpha, AlphaB) && !a.Alpha.NativeAlpha && a.NativeStatus.NaturalMeasureCandidate && a.NativeStatus.UniqueUnderConstraints && !a.NativeStatus.NativeMeasure && !a.NativeStatus.NativeAlpha && !a.NativeStatus.NativeR3 && containsAll(a.NativeStatus.Failures, []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeMeasureUniqueness, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatAlpha(a.Alpha) + " | " + FormatNativeStatus(a.NativeStatus) + " | " + FormatFirewalls(a.Firewalls)},
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

		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatDomain(a.Domain), FormatBasepoint(a.Basepoint), FormatDegree(a.Degree), FormatSelector(a.Selector), FormatNormalization(a.Normalization), FormatZ2(a.Z2), FormatAlternatives(a.Alternatives), FormatAlpha(a.Alpha), FormatNativeStatus(a.NativeStatus), FormatFirewalls(a.Firewalls), a.Final, NextGate, BoundaryMeasureObject, MeasureFormula, BranchMeasureFormula, BoundaryAlphaFormula}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
