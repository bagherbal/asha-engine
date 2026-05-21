package generation2boundaryactivationmeasurefunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE920-BOUNDARYACTIVATIONMEASURE-FUNCTOR-AUDIT"
	theoremName = "Gate 920: BoundaryActivationMeasure Functor Audit"
)

func Generation2BoundaryActivationMeasureFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 920 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "mu_B domain is the reduced active boundary-pair response, not an arbitrary polynomial", Passed: a.Domain.InheritedStatus == Gate919ShortStatus && a.Domain.Domain == ReducedResponse && len(a.Domain.NonzeroDegrees) == 2 && !a.Domain.IncludesLambda0 && a.Domain.ReducedBasepoint && !a.Domain.NativeTheorem && containsAll(a.Domain.Supports, []string{SupportMuBDomainReducedResponse, SupportMuBActsOnActiveDegrees, SupportMuBIgnoresLambda0Basepoint}) && containsAll(a.Domain.Failures, []string{FailureNoNativeMuBDomain, FailureReducedResponseBridgeSelected}), Detail: FormatDomain(a.Domain)},
			{Name: "mu_B extracts exterior degree components and inherits powers of S_split from the reduced response", Passed: !a.DegreeExtraction.SeparateS2Transport && a.DegreeExtraction.ExteriorGeneratedS2 && a.DegreeExtraction.DegreePowers[1] == 1 && a.DegreeExtraction.DegreePowers[2] == 2 && near(a.DegreeExtraction.DegreeCoefficients[1], SBoundary) && near(a.DegreeExtraction.DegreeCoefficients[2], SBoundary*SBoundary) && !a.DegreeExtraction.NativeTheorem && containsAll(a.DegreeExtraction.Supports, []string{SupportMuBExtractsByExteriorDegree, SupportDegreeKCarriesSPowerK, SupportSPowerFollowsExteriorResponse}) && containsAll(a.DegreeExtraction.Failures, []string{FailureNoNativeDegreeExtraction, FailureNoNativeSSplitTransportMap}), Detail: FormatDegreeExtraction(a.DegreeExtraction)},
			{Name: "mu_B integrates the degree-indexed Z2 airlock selector and recovers the target rank pair 3,7", Passed: a.Selector.Targets[1] == "[F_1/F_0]_{Z2}" && a.Selector.Targets[2] == "[F_2/F_0]_{Z2}" && a.Selector.Ranks[1] == RankI1 && a.Selector.Ranks[2] == RankI2 && a.Selector.RepresentativeIndependent && !a.Selector.UniqueNativeSelector && containsAll(a.Selector.Supports, []string{SupportMuBIntegratesSelector, SupportMuBRecoversRankPair, SupportSelectorRepresentativeIndependent}) && containsAll(a.Selector.Failures, []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeUniqueSelector}), Detail: FormatSelector(a.Selector)},
			{Name: "mu_B integrates lane-specific boundary-augmented chamber normalization and recovers 3/10 and 7/72", Passed: a.Chambers.Chambers[1] == "H_10=H_R^ambient+B_2" && a.Chambers.Chambers[2] == "H_72=Lambda^4 V_8+B_2" && a.Chambers.Ranks[1] == RankH10 && a.Chambers.Ranks[2] == RankH72 && nearLoose(a.Chambers.Weights[1], 0.3) && nearLoose(a.Chambers.Weights[2], float64(7)/72) && a.Chambers.ExplicitLaneWeights && !a.Chambers.NativeTheorem && containsAll(a.Chambers.Supports, []string{SupportMuBIntegratesChamberNormalization, SupportMuBRecoversCoefficients, SupportLaneNormalizationExplicit}) && containsAll(a.Chambers.Failures, []string{FailureNoNativeChamberNormalization, FailureNoNativeH1H2Reason}), Detail: FormatChambers(a.Chambers)},
			{Name: "cross-lane exclusion is absorbed into measure indexing only if I_B^Z2 is functional", Passed: a.CrossLanes.CorrectTargets[1] == "[F_1/F_0]_{Z2}" && a.CrossLanes.CorrectTargets[2] == "[F_2/F_0]_{Z2}" && a.CrossLanes.FalseTargets[1] == "[F_2/F_0]_{Z2}" && a.CrossLanes.FalseTargets[2] == "[F_1/F_0]_{Z2}" && a.CrossLanes.ExcludedIfFunctional && a.CrossLanes.AbsorbedInIndexing && !a.CrossLanes.FunctionhoodNative && containsAll(a.CrossLanes.Supports, []string{SupportMuBExcludesCrossLanesIfFunctional, SupportCrossLaneAbsorbedInIndexing}) && containsAll(a.CrossLanes.Failures, []string{FailureNoNativeZ2CrossLaneExclusion, FailureSelectorFunctionhoodNotNative}), Detail: FormatCrossLanes(a.CrossLanes)},
			{Name: "mu_B reassembles all five alpha subobjects and reconstructs alpha_B^Z2", Passed: a.Alpha.Formula == MeasureFormula && a.Alpha.BoundaryFormula == BoundaryAlphaMeasureFormula && a.Alpha.RankI1 == RankI1 && a.Alpha.RankI2 == RankI2 && a.Alpha.RankH1 == RankH10 && a.Alpha.RankH2 == RankH72 && near(a.Alpha.LinearContribution, AlphaLinear) && near(a.Alpha.QuadraticContribution, AlphaQuad) && near(a.Alpha.Alpha, AlphaB) && a.Alpha.ReassemblesFive && !a.Alpha.NativeAlpha && containsAll(a.Alpha.Supports, []string{SupportMuBReconstructsAlpha, SupportMeasureReassemblesFive, SupportAlphaAsMeasure}) && containsAll(a.Alpha.Failures, []string{FailureAlphaByMuBNotNative}), Detail: FormatAlpha(a.Alpha)},
			{Name: "BoundaryActivationMeasure remains a formal bridge candidate with native measure, alpha, and R3 firewalls preserved", Passed: a.NativeStatus.BridgeMeasureCandidate && !a.NativeStatus.NativeMeasure && !a.NativeStatus.NativeAlpha && !a.NativeStatus.NativeR3 && len(a.NativeStatus.MissingNativeTheorems) == 6 && containsAll(a.NativeStatus.Failures, []string{FailureNoNativeBoundaryActivationMeasure, FailureMuBFormalNotNative, FailureNoNativeMeasureUniqueness, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls) && containsAll(a.Firewalls.List(), []string{FailureNoNativeBoundaryActivationMeasure, FailureMuBFormalNotNative, FailureNoNativeMeasureUniqueness, FailureNoNativeMuBDomain, FailureNoNativeDegreeExtraction, FailureNoNativeSSplitTransportMap, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeChamberNormalization, FailureNoNativeZ2CrossLaneExclusion, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatNativeStatus(a.NativeStatus) + " | " + FormatFirewalls(a.Firewalls)},
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

		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatDomain(a.Domain), FormatDegreeExtraction(a.DegreeExtraction), FormatSelector(a.Selector), FormatChambers(a.Chambers), FormatCrossLanes(a.CrossLanes), FormatAlpha(a.Alpha), FormatNativeStatus(a.NativeStatus), FormatFirewalls(a.Firewalls), a.Final, NextGate, BoundaryActivationMeasure, BoundaryResponseFunctor, MeasureFormula, BranchMeasureFormula, BoundaryAlphaMeasureFormula}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
