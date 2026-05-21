package generation2nativereducedboundarypairresponsefunctionalaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NATIVE_REDUCED_BOUNDARYPAIR_RESPONSE_FUNCTIONAL_AUDIT"
	theoremName = "Gate 913 — Native Reduced BoundaryPair Response Functional Audit"
)

func Generation2NativeReducedBoundaryPairResponseFunctionalAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		checks := []theorem.Check{
			{Name: "Gate 912 decomposition inherited and subobject 1 selected without reopening phase or representative wounds", Passed: a.Inherited.SubobjectIndex == 1 && a.Inherited.TotalSubobjects == 5 && a.Inherited.SelectedSubobject == NativeReducedB2Theorem && !a.Inherited.ReopensPhaseSign && !a.Inherited.ReopensSocketOrder && !a.Inherited.ReopensRepresentative && !a.Inherited.DerivesAlpha && !a.Inherited.UpdatesOfficialLedger && containsAll(a.Inherited.Supports, []string{StatusGate912Inherited, StatusSubobjectOneSelected}) && containsAll(a.Inherited.Failures, []string{FailureReducedB2NotNativeFunctional, FailureAlphaStillSealed}), Detail: FormatInherited(a.Inherited)},
			{Name: "rank-two boundary exterior ledger has Lambda^0, Lambda^1, Lambda^2 and Lambda^3 B2 equal zero", Passed: a.Ledger.Rank == BoundaryPairRank && a.Ledger.Lambda0Dim == Lambda0Dim && a.Ledger.Lambda1Dim == Lambda1Dim && a.Ledger.Lambda2Dim == Lambda2Dim && a.Ledger.Lambda3Dim == Lambda3Dim && a.Ledger.Lambda3Zero && a.Ledger.CubicAndHigherVanish && containsAll(a.Ledger.Supports, []string{StatusExteriorLedgerBuilt, SupportCubicAbsentRankTwo, SupportLambda3Zero}), Detail: FormatLedger(a.Ledger)},
			{Name: "reduced response expands exactly to degree-one s exposure sum plus degree-two s^2 enclosure", Passed: a.Expansion.ExactShape && a.Expansion.ConstantRemoved && !a.Expansion.NativeFunctional && len(a.Expansion.DegreeOneTerms) == 2 && len(a.Expansion.DegreeTwoTerms) == 1 && !hasDegree(a.Expansion.ReducedTerms, 0) && near(sumDegree(a.Expansion.ReducedTerms, 1), 2*SBoundary) && near(sumDegree(a.Expansion.ReducedTerms, 2), SBoundary*SBoundary) && containsAll(a.Expansion.Supports, []string{StatusExpansionExact, SupportExactExpansion, SupportDegreeOneExposureSum, SupportDegreeTwoPairEnclosure}) && containsAll(a.Expansion.Failures, []string{FailureReducedB2NotNativeFunctional}), Detail: FormatExpansion(a.Expansion)},
			{Name: "zero-order term is suppressed by E_B minus one while native reason for reduction remains absent", Passed: a.ZeroOrder.RemovedByReduction && a.ZeroOrder.ReducedStartsAtOrderOne && !a.ZeroOrder.NativeReasonForReduction && containsAll(a.ZeroOrder.Supports, []string{StatusZeroOrderSuppressed, SupportZeroOrderSuppressed, SupportStartsAtOrderOne}) && containsAll(a.ZeroOrder.Failures, []string{FailureNoNativeReasonEBMinusOne}), Detail: FormatZeroOrder(a.ZeroOrder)},
			{Name: "cubic and higher response terms are absent by rank-two exterior truncation", Passed: a.Truncation.Rank == BoundaryPairRank && a.Truncation.HighestNonzeroDegree == 2 && a.Truncation.Lambda3Zero && a.Truncation.NoCubicOrHigher && a.Truncation.ExteriorAlgebraNativeFact && containsAll(a.Truncation.Supports, []string{StatusRankTwoTruncation, SupportCubicAbsentRankTwo, SupportLambda3Zero}), Detail: FormatTruncation(a.Truncation)},
			{Name: "multiplicative boundary-pair response is recorded only as a natural exterior activation candidate", Passed: a.Naturality.MultiplicativeCandidate && !a.Naturality.NativeASHAFunctional && !a.Naturality.VariationalPrinciple && !a.Naturality.FunctorialSelectionPrinciple && containsAll(a.Naturality.Supports, []string{StatusNaturalityCandidate, SupportMultiplicativeNatural, SupportReducedNontrivialPart}) && containsAll(a.Naturality.Failures, []string{FailureMultiplicativeNotNative, FailureNoVariationalProductForm}), Detail: FormatNaturality(a.Naturality)},
			{Name: "reduced response supplies alpha power shape but not Z2 targets, coefficients, cross-lane exclusion, or alpha", Passed: a.AlphaShape.SuppliesPowerShape && !a.AlphaShape.SelectsZ2FlagTargets && !a.AlphaShape.DerivesCoefficients && !a.AlphaShape.ProvesCrossLaneExclusion && !a.AlphaShape.DerivesAlpha && containsAll(a.AlphaShape.Supports, []string{StatusAlphaShapeMatched, SupportSuppliesAlphaPowerShape, SupportAlphaPolynomialShapeMatches}) && containsAll(a.AlphaShape.Failures, []string{FailureNoZ2FlagTargets, FailureNoAlphaCoefficients, FailureNoCrossLaneExclusion, FailureReducedResponseNotAlphaAlone, FailureAlphaStillSealed}), Detail: FormatAlphaShape(a.AlphaShape)},
			{Name: "S_split is used as the response parameter but typed native transport is not certified", Passed: a.Transport.UsesSAsParameter && !a.Transport.NativeTransport && !a.Transport.TypedParameterMap && containsAll(a.Transport.Supports, []string{SupportSsplitFeedsShape}) && containsAll(a.Transport.Failures, []string{StatusSsplitFirewall, FailureNoNativeTransportSIntoB2, FailureNoTypedSToExteriorParameter}), Detail: FormatTransport(a.Transport)},
			{Name: "native selection, target selection, alpha, R3, full A_F descent, and Yukawa-spectrum firewalls are preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureReducedB2NotNativeFunctional, FailureNoNativeReasonEBMinusOne, FailureMultiplicativeNotNative, FailureNoVariationalProductForm, FailureNoNativeTransportSIntoB2, FailureNoZ2FlagTargets, FailureNoCrossLaneExclusion, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatLedger(a.Ledger), FormatExpansion(a.Expansion), FormatZeroOrder(a.ZeroOrder), FormatTruncation(a.Truncation), FormatNaturality(a.Naturality), FormatAlphaShape(a.AlphaShape), FormatTransport(a.Transport), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
