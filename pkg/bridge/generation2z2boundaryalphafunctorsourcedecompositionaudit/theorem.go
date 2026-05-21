package generation2z2boundaryalphafunctorsourcedecompositionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_Z2_BOUNDARY_ALPHA_FUNCTOR_SOURCE_DECOMPOSITION_AUDIT"
	theoremName = "Gate 912 — Z2 BoundaryAlpha Functor Source Decomposition Audit"
)

func Generation2Z2BoundaryAlphaFunctorSourceDecompositionAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}

		subs := a.Decomposition.Subobjects
		checks := []theorem.Check{
			{Name: "Gate 911 rail inherited without reopening phase, socket-order, or representative-alpha wounds", Passed: !a.Inherited.ReopensPhaseSign && !a.Inherited.ReopensSocketOrder && !a.Inherited.ReopensRepAlpha && !a.Inherited.DerivesAlpha && !a.Inherited.UpdatesOfficialLedger && a.Inherited.SelectedRail == NativeFunctorObject && containsAll(a.Inherited.Supports, []string{StatusGate911Inherited, StatusNoLoopBack}) && containsAll(a.Inherited.Failures, []string{FailureAlphaStillSealed, FailureNoNativeZ2BoundaryAlphaFunctor}), Detail: FormatInherited(a.Inherited)},
			{Name: "sealed BoundaryAlpha_Z2 formula is reproduced as representative-free but non-native", Passed: a.Formula.RepresentativeFree && !a.Formula.Native && near(a.Formula.Alpha, AlphaB) && a.Formula.RankPair == [2]int{RankF1OverF0, RankF2OverF0} && a.Formula.Denominators == [2]int{LinearDenom, QuadDenom} && containsAll(a.Formula.Supports, []string{StatusSealedFormulaReproduced, SupportRankPairRepresentativeFree}) && containsAll(a.Formula.Failures, []string{FailureAlphaStillSealed, FailureNoNativeZ2BoundaryAlphaFunctor}), Detail: FormatFormula(a.Formula)},
			{Name: "native Z2 BoundaryAlpha functor decomposes into exactly five missing required subobjects", Passed: a.Decomposition.RequiredCount == 5 && a.Decomposition.CertifiedCount == 0 && a.Decomposition.MissingCount == 5 && !a.Decomposition.NativeFunctor && !a.Decomposition.AlphaNative && !a.Decomposition.R3Native && allRequiredSubobjectsPresent(subs) && containsAll(a.Decomposition.Supports, []string{StatusDecompositionComplete, SupportNativeZ2AlphaDecomposed, SupportAlphaReducedToFive}) && containsAll(a.Decomposition.Failures, []string{FailureNoNativeZ2BoundaryAlphaFunctor, FailureAlphaStillSealed, FailureNotNativeR3}), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "subobject 1 records reduced B2 response shape without certifying native boundary functional", Passed: len(subs) >= 1 && subs[0].Index == 1 && subs[0].RequiredTheorem == NativeReducedB2Theorem && subs[0].Required && !subs[0].CertifiedNative && subs[0].CorrectShape && containsAll(subs[0].Supports, []string{SupportReducedB2Required, SupportReducedB2CorrectShape, SupportZeroOrderSuppressed, SupportCubicAbsent}) && containsAll(subs[0].Failures, []string{FailureReducedB2NotNativeFunctional, FailureNoNativeReasonEBMinusOne, FailureNoNativeTransportSInB2}), Detail: FormatSubobject(subs[0])},
			{Name: "subobject 2 records degree-to-Z2-flag-class selector with representative-free rank pair 3,7", Passed: len(subs) >= 2 && subs[1].Index == 2 && subs[1].RequiredTheorem == DegreeSelectorTheorem && subs[1].Required && !subs[1].CertifiedNative && containsAll(subs[1].Supports, []string{SupportDegreeSelectorRequired, SupportDegreeOneExposed, SupportDegreeTwoFull, SupportRankPairRepresentativeFree}) && containsAll(subs[1].Failures, []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeLambda1ExposedMap, FailureNoNativeLambda2FullMap}), Detail: FormatSubobject(subs[1])},
			{Name: "subobject 3 records cross-lane exclusion need and blocks wrong linear/quadratic terms", Passed: len(subs) >= 3 && subs[2].Index == 3 && subs[2].RequiredTheorem == CrossLaneTheorem && subs[2].Required && !subs[2].CertifiedNative && containsAll(subs[2].ForbiddenTargets, []string{ForbiddenLinearFull, ForbiddenQuadraticExposed}) && containsAll(subs[2].WrongTerms, []string{WrongLinearFullTerm, WrongQuadraticExposedTerm}) && containsAll(subs[2].Supports, []string{SupportCrossLaneRequired, SupportCrossLanesExcludedIfFunctor}) && containsAll(subs[2].Failures, []string{FailureNoNativeZ2CrossLane, FailureNoNativeLinearDomainExclusion, FailureNoNativeQuadraticFaceExclusion}), Detail: FormatSubobject(subs[2])},
			{Name: "subobject 4 records S_split transport law as missing typed transport into degree-one and degree-two responses", Passed: len(subs) >= 4 && subs[3].Index == 4 && subs[3].RequiredTheorem == SsplitTransportTheorem && subs[3].Required && !subs[3].CertifiedNative && containsAll(subs[3].Supports, []string{SupportSsplitTransportRequired, SupportSsplitFeedsDegreeShape}) && containsAll(subs[3].Failures, []string{FailureNoNativeTransportS, FailureNoTypedSToLambda1, FailureNoTypedS2ToLambda2}), Detail: FormatSubobject(subs[3])},
			{Name: "subobject 5 records H_10 and H_72 denominator chamber typing while preserving activation-law firewall", Passed: len(subs) >= 5 && subs[4].Index == 5 && subs[4].RequiredTheorem == DenominatorTypingTheorem && subs[4].Required && !subs[4].CertifiedNative && containsAll(subs[4].Supports, []string{SupportDenominatorChambersRequired, SupportDenominatorsTyped}) && containsAll(subs[4].Failures, []string{FailureDenominatorNotActivation}), Detail: FormatSubobject(subs[4])},
			{Name: "native alpha, native R3, full A_F descent, generation/flavor/Yukawa, and official-update firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNoNativeZ2BoundaryAlphaFunctor, FailureReducedB2NotNativeFunctional, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeZ2CrossLane, FailureNoNativeTransportS, FailureDenominatorNotActivation, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoNativeYukawaOperator}), Detail: FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatInherited(a.Inherited), FormatFormula(a.Formula), FormatDecomposition(a.Decomposition), FormatFirewalls(a.Firewalls), a.Final}
		for _, s := range a.Decomposition.Subobjects {
			notes = append(notes, FormatSubobject(s))
		}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
