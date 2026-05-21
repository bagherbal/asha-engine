package generation2boundarydegreeexposureenclosurefunctoraudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE924-BOUNDARYDEGREE-EXPOSUREENCLOSURE-FUNCTOR-AUDIT"
	theoremName = "Gate 924: BoundaryDegree ExposureEnclosure Functor Audit"
)

func Generation2BoundaryDegreeExposureEnclosureFunctorAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 924 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "Lambda^1 B2 has native exterior shape as single-boundary exposure", Passed: a.DegreeOne.Degree == 1 && a.DegreeOne.BoundaryFactors == 1 && a.DegreeOne.NativeShape && !a.DegreeOne.TopDegree && !a.DegreeOne.NativeTargetMap && containsAll(a.DegreeOne.Supports, []string{SupportLambda1SingleGenerator, SupportDegreeOneActivatesOneFactor, SupportDegreeOneExposureFromExteriorDegree, SupportLambda1ExposureNativeShape}) && containsAll(a.DegreeOne.Failures, []string{FailureExposureLanguageNotTargetFunctor}), Detail: FormatDegree(a.DegreeOne)},
			{Name: "Lambda^2 B2 has native exterior shape as full boundary-pair enclosure", Passed: a.DegreeTwo.Degree == 2 && a.DegreeTwo.BoundaryFactors == 2 && a.DegreeTwo.TopDegree && a.DegreeTwo.NativeShape && !a.DegreeTwo.NativeTargetMap && containsAll(a.DegreeTwo.Supports, []string{SupportLambda2TopPairSpace, SupportDegreeTwoRequiresBothFactors, SupportDegreeTwoFullEnclosureFromTopDegree, SupportLambda2EnclosureNativeShape}) && containsAll(a.DegreeTwo.Failures, []string{FailureEnclosureLanguageNotTargetFunctor}), Detail: FormatDegree(a.DegreeTwo)},
			{Name: "exposure/enclosure contrast is grounded in exterior degree but does not select Z2 targets", Passed: a.Contrast.GroundedInDegree && !a.Contrast.ArbitraryLabels && !a.Contrast.SelectsZ2Targets && containsAll(a.Contrast.Supports, []string{SupportExposureEnclosureGrounded, SupportOneVsTwoFactorDistinguishesDegrees, SupportBoundaryDegreeTypesNotRandom}) && containsAll(a.Contrast.Failures, []string{FailureExteriorContrastDoesNotSelectZ2}), Detail: FormatContrast(a.Contrast)},
			{Name: "top-degree pair activation source-types cumulative enclosure but not native target functor", Passed: a.Cumulative.TopDegreeSourceCandidate && !a.Cumulative.NativeTargetFunctor && a.Cumulative.RankCumulative == RankF2OverF0 && a.Cumulative.RankAssociatedGraded == RankF2OverF1 && containsAll(a.Cumulative.Supports, []string{SupportTopDegreeCumulativeEnclosure, SupportFullEnclosurePointsToF2OverF0, SupportCumulativeHasTopDegreeSource}) && containsAll(a.Cumulative.Failures, []string{FailureNoNativeTopDegreeToF2OverF0}), Detail: FormatCumulative(a.Cumulative)},
			{Name: "selector source is reduced to exterior-degree type while degree-to-flag target functor remains blocked", Passed: a.Selector.StrengthenedByDegreeType && !a.Selector.TargetFunctorNative && containsAll(a.Selector.Supports, []string{SupportExposureEnclosureStrengthensSelector, SupportSelectorReducedToExteriorDegree, SupportIBZ2TargetsCompatible}) && containsAll(a.Selector.Failures, []string{FailureSelectorFunctionhoodTargetFunctor}), Detail: FormatSelector(a.Selector)},
			{Name: "BoundaryActivationMeasure selector input has stronger exterior-degree source but mu_B remains non-native", Passed: a.MuB.SelectorInputExteriorTyped && a.MuB.FunctionhoodGapWeakened && !a.MuB.NativeMuB && containsAll(a.MuB.Supports, []string{SupportMuBSelectorInputExteriorSource, SupportMuBFunctionhoodGapWeakened, SupportAlphaMeasureStrongerSelectorTyping}) && containsAll(a.MuB.Failures, []string{FailureMuBStillNotNativeTargetFunctor, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && firewallsOK(a.Firewalls), Detail: FormatMuB(a.MuB) + " | " + FormatFirewalls(a.Firewalls)},
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
		notes := []string{a.Truth, a.Classification, a.ShortStatus, a.InheritedStatus, BoundaryPair, ExteriorLedger, ReducedResponse, DegreeOneTerm, DegreeTwoTerm, TargetFunctorGap, NextGate, SourceNativeExteriorShape, SourceBridgeTargetBlocked, FormatDegree(a.DegreeOne), FormatDegree(a.DegreeTwo), FormatContrast(a.Contrast), FormatCumulative(a.Cumulative), FormatSelector(a.Selector), FormatMuB(a.MuB), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
