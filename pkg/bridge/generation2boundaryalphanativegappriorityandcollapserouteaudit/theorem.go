package generation2boundaryalphanativegappriorityandcollapserouteaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE919-BOUNDARYALPHA-NATIVEGAP-PRIORITY-COLLAPSEROUTE-AUDIT"
	theoremName = "Gate 919: BoundaryAlpha NativeGap Priority and CollapseRoute Audit"
)

func Generation2BoundaryAlphaNativeGapPriorityAndCollapseRouteAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 919 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "five explicit Gate 918 alpha gaps share a boundary response measure structure", Passed: a.GapIndependence.InheritedStatus == Gate918ShortStatus && len(a.GapIndependence.GapNames) == 5 && a.GapIndependence.InitiallyIndependent && a.GapIndependence.MayCollapseToMeasure && a.GapIndependence.CandidateMasterObject == BoundaryActivationMeasureName && !a.GapIndependence.NativeCertified && containsAll(a.GapIndependence.Supports, []string{SupportFiveGapsShareMeasureStructure, SupportAlphaGapsMayCollapseToMeasure, SupportSingleMasterFunctorCouldGeneratePieces}) && containsAll(a.GapIndependence.Failures, []string{FailureNoNativeBoundaryActivationMeasureCertified}), Detail: FormatGapIndependence(a.GapIndependence)},
			{Name: "native gap priority ranking selects S_split transport first and BoundaryActivationMeasure as strongest collapse route", Passed: a.PriorityRanking.RankingComplete && a.PriorityRanking.HighestPriorityGap == "S_split -> s transport map" && a.PriorityRanking.BoundaryActivationRank == 2 && a.PriorityRanking.CrossLaneDependent && len(a.PriorityRanking.OrderedGaps) == 5 && containsAll(a.PriorityRanking.Supports, []string{SupportNativeGapPriorityRankingCompleted, SupportSSplitTransportHighestPriority, SupportBoundaryActivationMeasureStrongest}) && containsAll(a.PriorityRanking.Failures, []string{FailureNoNativeSSplitTransportMap, FailureNoNativeBoundaryResponseMeasure, FailureNoNativeZ2CrossLaneExclusionTheorem}), Detail: FormatPriorityRanking(a.PriorityRanking)},
			{Name: "formal mu_B collapse route reassembles all five subobjects and reconstructs alpha_B", Passed: a.CollapseRoute.MasterObject == BoundaryActivationMeasureName && a.CollapseRoute.AlternateName == BoundaryResponseFunctorName && a.CollapseRoute.ReassemblesAllFive && !a.CollapseRoute.NativeTheorem && a.CollapseRoute.RankI1 == RankF1OverF0 && a.CollapseRoute.RankI2 == RankF2OverF0 && a.CollapseRoute.RankH1 == RankH10 && a.CollapseRoute.RankH2 == RankH72 && near(a.CollapseRoute.LinearContribution, AlphaLinear) && near(a.CollapseRoute.QuadraticContribution, AlphaQuad) && near(a.CollapseRoute.Alpha, AlphaB) && containsAll(a.CollapseRoute.Supports, []string{SupportBoundaryActivationMeasureReassembles, SupportMuBReconstructsAlpha, SupportNativeAlphaMayNeedMeasureFunctor}) && containsAll(a.CollapseRoute.Failures, []string{FailureMuBFormalNotNative}), Detail: FormatCollapseRoute(a.CollapseRoute)},
			{Name: "master BoundaryActivationMeasure requirements are explicit but uncertified", Passed: a.Requirements.AllRequired && !a.Requirements.AllCertified && len(a.Requirements.Requirements) == 6 && containsAll(a.Requirements.Supports, []string{SupportNativeBoundaryAlphaAsMeasure, SupportMeasureAbsorbsPieces}) && containsAll(a.Requirements.Failures, []string{FailureNoNativeBoundaryActivationMeasureCertified}), Detail: FormatRequirements(a.Requirements)},
			{Name: "alpha status collapses from five gaps to one master candidate gap without native promotion", Passed: a.Promotion.PreviousStatus != a.Promotion.NewStatus && a.Promotion.BridgeCandidate && !a.Promotion.NativeAlpha && !a.Promotion.NativeR3 && !a.Promotion.OfficialUpdate && containsAll(a.Promotion.Supports, []string{SupportAlphaGapsCollapseCandidate}) && containsAll(a.Promotion.Failures, []string{FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}), Detail: FormatPromotion(a.Promotion)},
			{Name: "native measure, native alpha, native R3, full A_F, generation/flavor/Yukawa firewalls remain preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.Firewalls.List(), []string{FailureNoNativeBoundaryActivationMeasureCertified, FailureNoNativeBoundaryResponseMeasure, FailureMuBFormalNotNative, FailureNoNativeSSplitTransportMap, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeResponseChamberNormalization, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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

		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatGapIndependence(a.GapIndependence), FormatPriorityRanking(a.PriorityRanking), FormatCollapseRoute(a.CollapseRoute), FormatRequirements(a.Requirements), FormatPromotion(a.Promotion), FormatFirewalls(a.Firewalls), a.Final, NextGate}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
