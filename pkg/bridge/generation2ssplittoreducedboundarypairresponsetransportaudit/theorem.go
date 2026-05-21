package generation2ssplittoreducedboundarypairresponsetransportaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_S_SPLIT_TO_REDUCED_BOUNDARYPAIR_RESPONSE_TRANSPORT_AUDIT"
	theoremName = "Gate 916 — S_split to Reduced BoundaryPair Response Transport Audit"
)

func Generation2SSplitToReducedBoundaryPairResponseTransportAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		checks := []theorem.Check{
			{Name: "Gates 913, 914, and 915 are inherited without deriving alpha or native R3", Passed: a.Inherited.ResponseShapeCertified && a.Inherited.SelectorShapeTyped && a.Inherited.CrossLanesBlocked && !a.Inherited.DerivesAlpha && !a.Inherited.UpdatesOfficialLedger && !a.Inherited.PromotesNativeR3 && containsAll(a.Inherited.Supports, []string{StatusGates913To915Inherited}), Detail: FormatInherited(a.Inherited)},
			{Name: "S_split transport target is the scalar parameter of R_B(s), not alpha or socket magnitude", Passed: a.Target.Source == TransportSource && a.Target.Target == TransportTarget && a.Target.MapCandidate == TransportMapCandidate && !a.Target.TargetsAlphaDirectly && !a.Target.TargetsSocketMag && a.Target.UsesRBOnly && !a.Target.NativeMap && containsAll(a.Target.Supports, []string{SupportTransportTargetIsResponseParameter, SupportSSplitScalarResponseParameter, SupportAlphaUsesSSplitOnlyThroughRB}) && containsAll(a.Target.Failures, []string{FailureNoNativeTsMap}), Detail: FormatTarget(a.Target)},
			{Name: "single uniform insertion into both boundary factors generates s and s squared without separate quadratic transport", Passed: a.Insertion.Insertions == 2 && a.Insertion.ScalarInsertedPerFactor && !a.Insertion.SeparateQuadraticTransport && a.Insertion.QuadraticFromProduct && a.Insertion.ExpandedResponse == ExpandedReducedResponse && !a.Insertion.NativeUniformLaw && containsAll(a.Insertion.Supports, []string{SupportSingleInsertionGeneratesPowers, SupportS2FromExteriorProduct, SupportPowerResponseUniformInsertion}) && containsAll(a.Insertion.Failures, []string{FailureNoNativeUniformInsertionReason}), Detail: FormatInsertion(a.Insertion)},
			{Name: "S_split is compatible with scalar exterior-generator multiplication but scalar type remains sealed", Passed: near(a.Scalar.Parameter, SBoundary) && a.Scalar.Dimensionless && a.Scalar.CanMultiplyGenerators && !a.Scalar.ScalarTypeNative && containsAll(a.Scalar.Supports, []string{SupportSSplitScalarType, SupportDimensionlessMultipliesExteriorGenerator}) && containsAll(a.Scalar.Failures, []string{FailureSSplitScalarTypeSealed, FailureNoTypedSSplitExteriorParameterMap}), Detail: FormatScalar(a.Scalar)},
			{Name: "S_split applies to active generator terms while identity is basepoint removed by reduction", Passed: a.Reduction.TransportAppliesToActive && !a.Reduction.IdentityTransported && a.Reduction.BasepointRemoved && !a.Reduction.NativeBasepointTheorem && a.Reduction.Reduction == ReducedResponse && containsAll(a.Reduction.Supports, []string{SupportTransportAppliesToActiveTerms, SupportIdentityBasepointRemoved}) && containsAll(a.Reduction.Failures, []string{FailureNoNativeBasepointReduction}), Detail: FormatReduction(a.Reduction)},
			{Name: "single-insertion response is compatible with degree selector and does not reopen cross-lane pollution", Passed: a.Selector.DegreeSelectorCompatible && a.Selector.FeedsCorrectAlphaLanes && !a.Selector.ReopensCrossLanePollution && !a.Selector.NativeTransportTheorem && containsAll(a.Selector.Supports, []string{SupportCompatibleWithDegreeSelector, SupportSingleInsertionFeedsCorrectAlphaLanes, SupportTransportDoesNotReopenCrossLanes}) && containsAll(a.Selector.Failures, []string{FailureSelectorCompatibilityNotNative}), Detail: FormatSelector(a.Selector)},
			{Name: "alpha is reconstructed under the transport seal and prior subobjects, but not as native theorem", Passed: near(a.Alpha.AlphaLinear, AlphaLinear) && near(a.Alpha.AlphaQuadratic, AlphaQuad) && near(a.Alpha.AlphaTotal, AlphaB) && a.Alpha.RankPair == [2]int{RankF1OverF0, RankF2OverF0} && a.Alpha.TransportSealAssumed && a.Alpha.PriorSubobjectsAssumed && !a.Alpha.NativeAlphaTheorem && containsAll(a.Alpha.Supports, []string{SupportAlphaReconstructedGivenTransport, SupportTransportWoundReducedToNativeTs}) && containsAll(a.Alpha.Failures, []string{FailureAlphaReconstructionNotNative, FailureAlphaStillSealed}), Detail: FormatAlpha(a.Alpha)},
			{Name: "native transport map, denominator normalization, alpha, R3, full A_F, and generation/flavor/Yukawa firewalls remain preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeTsMap, FailureNoNativeTransportToZ2Airlock, FailureNoTypedSSplitExteriorParameterMap, FailureNoNativeUniformInsertionReason, FailureDenominatorNormalizationExternal, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatTarget(a.Target), FormatInsertion(a.Insertion), FormatScalar(a.Scalar), FormatReduction(a.Reduction), FormatSelector(a.Selector), FormatAlpha(a.Alpha), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
