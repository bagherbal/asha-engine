package generation2z2boundaryalphadecomposedfunctorconsolidationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_Z2_BOUNDARYALPHA_DECOMPOSED_FUNCTOR_CONSOLIDATION_AUDIT"
	theoremName = "Gate 918 — Z2 BoundaryAlpha DecomposedFunctor Consolidation and Native-Theorem Gap Audit"
)

func Generation2Z2BoundaryAlphaDecomposedFunctorConsolidationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		checks := []theorem.Check{
			{Name: "Gates 913 through 917 are inherited as five alpha subobjects audited at shape level only", Passed: a.Subobjects.AllAuditedAtShape && !a.Subobjects.NativeTheorem && containsAll(a.Subobjects.Supports, []string{StatusFiveSubobjectsInherited}) && containsAll(a.Subobjects.Failures, []string{FailureReconstructionNotNativeAlpha, FailureAlphaBridgeCandidateNotNative}), Detail: FormatSubobjects(a.Subobjects)},
			{Name: "five shape-level subobjects compose coherently and reconstruct alpha_B as a decomposed bridge candidate", Passed: a.BoundaryAlpha.InternalCoherent && !a.BoundaryAlpha.OpaqueSeal && a.BoundaryAlpha.BridgeCandidate && !a.BoundaryAlpha.NativeTheorem && a.BoundaryAlpha.RankF1OverF0 == RankF1OverF0 && a.BoundaryAlpha.RankF2OverF0 == RankF2OverF0 && a.BoundaryAlpha.RankH10 == RankH10 && a.BoundaryAlpha.RankH72 == RankH72 && near(a.BoundaryAlpha.LinearContribution, AlphaLinear) && near(a.BoundaryAlpha.QuadraticContribution, AlphaQuad) && near(a.BoundaryAlpha.TotalAlpha, AlphaB) && containsAll(a.BoundaryAlpha.Supports, []string{SupportDecomposedFunctorCoherent, SupportFiveSubobjectsCompose, SupportAlphaNoLongerOpaque, SupportAlphaBridgeCandidateForm}) && containsAll(a.BoundaryAlpha.Failures, []string{FailureInternalCoherenceNotNative, FailureReconstructionNotNativeAlpha}), Detail: FormatBoundaryAlpha(a.BoundaryAlpha)},
			{Name: "BoundaryAlpha_Z2 is representative-independent on the Z2 airlock class but not a native airlock functor", Passed: a.RepresentativeIndependence.RankPair == [2]int{RankF1OverF0, RankF2OverF0} && a.RepresentativeIndependence.TauPhiPreservesRankPair && !a.RepresentativeIndependence.PhaseSignEntersAlpha && a.RepresentativeIndependence.CorrectAlphaDomain && !a.RepresentativeIndependence.NativeAirlockFunctor && containsAll(a.RepresentativeIndependence.Supports, []string{SupportBoundaryAlphaRepresentativeFree, SupportPhaseSignAbsentAfterQuotient, SupportZ2AirlockCorrectDomain}) && containsAll(a.RepresentativeIndependence.Failures, []string{FailureRepresentativeIndependenceNotNative}), Detail: FormatRepresentativeIndependence(a.RepresentativeIndependence)},
			{Name: "alpha branch is promoted from opaque seal to decomposed bridge-theorem candidate with explicit theorem obligations", Passed: a.BridgeCandidate.PreviousStatus != a.BridgeCandidate.NewStatus && a.BridgeCandidate.AllVisibleComponents && !a.BridgeCandidate.NativeTheorem && len(a.BridgeCandidate.TheoremObligations) == 5 && containsAll(a.BridgeCandidate.Supports, []string{SupportPromotedFromOpaqueSeal, SupportAlphaTheoremObligationLedger}) && containsAll(a.BridgeCandidate.Failures, []string{FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeSSplitTransportMap, FailureNoNativeResponseChamberNormalization}), Detail: FormatBridgeCandidate(a.BridgeCandidate)},
			{Name: "R3 trace ledger rests on decomposed alpha bridge candidate while operator values remain diagnostic only", Passed: len(a.R3TraceLedger.TraceRows) == 3 && nearLoose(a.R3TraceLedger.NEffOperator, NEffOperator) && nearLoose(a.R3TraceLedger.CYukawaOperator, CYukawaOperator) && nearLoose(a.R3TraceLedger.CHiggsOperator, CHiggsOperator) && a.R3TraceLedger.DiagnosticOnly && !a.R3TraceLedger.OfficialUpdateAllowed && containsAll(a.R3TraceLedger.Supports, []string{SupportR3TraceLedgerOnCandidate, SupportOperatorsRemainReconstructed}) && containsAll(a.R3TraceLedger.Failures, []string{FailureNotNativeR3, FailureNoOfficialNEffUpdateAllowed, FailureNoCYukawaOrCHiggsUpdateAllowed}), Detail: FormatR3TraceLedger(a.R3TraceLedger)},
			{Name: "native R3 blockers are now explicit alpha theorem gaps plus full A_F descent / orientation status", Passed: len(a.NativeGaps.AlphaGaps) == 5 && len(a.NativeGaps.FiniteLayerGaps) == 2 && a.NativeGaps.GenerationFlavorR4OrLater && !a.NativeGaps.NativeR3 && containsAll(a.NativeGaps.Supports, []string{SupportNativeR3BlockersExplicit}) && containsAll(a.NativeGaps.Failures, []string{FailureFullAFDescentStillBlocked, FailureHiggsOrientationSealed, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}), Detail: FormatNativeGaps(a.NativeGaps)},
			{Name: "native alpha, native R3, full A_F, generation/flavor/Yukawa, and official ledger firewalls remain preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.Firewalls.List(), []string{FailureNotNativeR3, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeSSplitTransportMap, FailureNoNativeResponseChamberNormalization, FailureAlphaBridgeCandidateNotNative, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoOfficialNEffUpdateAllowed}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatSubobjects(a.Subobjects), FormatBoundaryAlpha(a.BoundaryAlpha), FormatRepresentativeIndependence(a.RepresentativeIndependence), FormatBridgeCandidate(a.BridgeCandidate), FormatR3TraceLedger(a.R3TraceLedger), FormatNativeGaps(a.NativeGaps), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
