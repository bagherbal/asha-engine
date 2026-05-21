package generation2boundaryactivationmeasurenativeconstraintsourceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2-GATE922-BOUNDARYACTIVATIONMEASURE-NATIVECONSTRAINT-SOURCE-AUDIT"
	theoremName = "Gate 922: BoundaryActivationMeasure NativeConstraint Source Audit"
)

func Generation2BoundaryActivationMeasureNativeConstraintSourceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default Gate 922 audit", Passed: false, Detail: err.Error()}}, Notes: []string{FinalTruth, Classification, ShortStatus}}
		}

		checks := []theorem.Check{
			{Name: "reduced-response constraint has bridge-strong basepoint deviation source, not native theorem", Passed: a.Ledger.ReducedResponse.SourceStatus == SourceBridgeStrong && a.Ledger.ReducedResponse.BridgeLawful && !a.Ledger.ReducedResponse.Native && containsAll(a.Ledger.ReducedResponse.Supports, []string{SupportReducedBasepointDeviation, SupportLambda0InactiveBasepoint, SupportAlphaStartsAfterActivation}) && containsAll(a.Ledger.ReducedResponse.Failures, []string{FailureNoNativeBasepointDeviation}), Detail: FormatConstraint(a.Ledger.ReducedResponse)},
			{Name: "degree-respecting constraint has native-shape exterior algebra source", Passed: a.Ledger.DegreeRespect.SourceStatus == SourceNativeShapeStrong && a.Ledger.DegreeRespect.Native && a.Ledger.DegreeRespect.BridgeLawful && containsAll(a.Ledger.DegreeRespect.Supports, []string{SupportDegreeNativeExteriorSource, SupportSPowerFromExteriorDegree, SupportDegreeTwoNotSeparateTransport}) && containsAll(a.Ledger.DegreeRespect.Failures, []string{FailureNoFullNativeMeasureFromExteriorDegree}), Detail: FormatConstraint(a.Ledger.DegreeRespect)},
			{Name: "selector functionhood has only exposure/enclosure bridge candidate and is primary native gap", Passed: a.Ledger.SelectorFunctionhood.SourceStatus == SourceBridgeCandidateNotNative && a.Ledger.SelectorFunctionhood.Primary && !a.Ledger.SelectorFunctionhood.Native && a.Ledger.SelectorFunctionhood.BridgeLawful && containsAll(a.Ledger.SelectorFunctionhood.Supports, []string{SupportSelectorExposureSource, SupportDegreeOneTwoSourceTargets}) && containsAll(a.Ledger.SelectorFunctionhood.Failures, []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureExposureEnclosureNotNativeFunctor}), Detail: FormatConstraint(a.Ledger.SelectorFunctionhood)},
			{Name: "cross-lane exclusion is dependent on selector functionhood", Passed: a.Ledger.CrossLaneExclusion.SourceStatus == SourceDependentOnSelector && a.Ledger.CrossLaneExclusion.Dependent && !a.Ledger.CrossLaneExclusion.Native && containsAll(a.Ledger.CrossLaneExclusion.Supports, []string{SupportCrossLaneSourceSelector, SupportCrossLaneDependent}) && containsAll(a.Ledger.CrossLaneExclusion.Failures, []string{FailureNoNativeCrossLaneWithoutSelector}), Detail: FormatConstraint(a.Ledger.CrossLaneExclusion)},
			{Name: "chamber normalization has bridge-strong local/global lane source candidate", Passed: a.Ledger.ChamberNormalization.SourceStatus == SourceBridgeStrong && a.Ledger.ChamberNormalization.BridgeLawful && !a.Ledger.ChamberNormalization.Native && containsAll(a.Ledger.ChamberNormalization.Supports, []string{SupportChamberLocalGlobalSource, SupportH10LocalRightChamber, SupportH72GlobalLambda4Chamber, SupportUniformBoundaryAugmentation}) && containsAll(a.Ledger.ChamberNormalization.Failures, []string{FailureNoNativeLaneLocalityToChamber, FailureNoNativeResponseChamberNormalization}), Detail: FormatConstraint(a.Ledger.ChamberNormalization)},
			{Name: "Z2 representative-independence has strong orientation-class bridge source", Passed: a.Ledger.Z2Independence.SourceStatus == SourceBridgeStrongOrientation && a.Ledger.Z2Independence.BridgeLawful && !a.Ledger.Z2Independence.Native && containsAll(a.Ledger.Z2Independence.Supports, []string{SupportZ2StrongOrientationSource, SupportPhaseGaugeForAlphaTrace, SupportAlphaRankPairZ2Invariant}) && containsAll(a.Ledger.Z2Independence.Failures, []string{FailureNoNativeGlobalPhaseZ2Equivariance}), Detail: FormatConstraint(a.Ledger.Z2Independence)},
			{Name: "positivity is compatibility only and does not select mu_B", Passed: a.Ledger.Positivity.SourceStatus == SourceCompatibilityOnly && a.Ledger.Positivity.Compatible && !a.Ledger.Positivity.Native && containsAll(a.Ledger.Positivity.Supports, []string{SupportPositiveActiveLanes, SupportPositiveMeasureActiveResponse}) && containsAll(a.Ledger.Positivity.Failures, []string{FailurePositivityNotSelectionTheorem, FailurePositivityNotUniqueMuB}), Detail: FormatConstraint(a.Ledger.Positivity)},
			{Name: "native status identifies selector functionhood as primary remaining gap", Passed: a.NativeStatus.InheritedStatus == Gate921ShortStatus && a.NativeStatus.ConstraintsPartlySourced && a.NativeStatus.DegreeRespectStrongestNativeShape && a.NativeStatus.SelectorFunctionhoodPrimaryGap && !a.NativeStatus.NativeBoundaryActivationMeasure && !a.NativeStatus.NativeAlpha && !a.NativeStatus.NativeR3 && containsAll(a.NativeStatus.Supports, []string{SupportSelectorPrimaryGap}) && containsAll(a.NativeStatus.Failures, []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeSelectorFunctionhood, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}) && near(a.Alpha.LinearContribution, AlphaLinear) && near(a.Alpha.QuadraticContribution, AlphaQuad) && near(a.Alpha.Alpha, AlphaB) && !a.Alpha.NativeAlpha && firewallsOK(a.Firewalls), Detail: FormatNativeStatus(a.NativeStatus) + " | " + FormatAlpha(a.Alpha) + " | " + FormatFirewalls(a.Firewalls)},
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

		notes := []string{a.Truth, a.Classification, a.ShortStatus, StrategicConclusion, FormatLedger(a.Ledger), FormatAlpha(a.Alpha), FormatNativeStatus(a.NativeStatus), FormatFirewalls(a.Firewalls), a.Final, NextGate, BoundaryMeasureObject, MeasureFormula, BranchMeasureFormula, BoundaryAlphaFormula, ReducedResponse, ExteriorResponse, StrongestNativeShapeSource, StrongestBridgeSources, PrimaryGapSelectorFunctionhood}
		notes = append(notes, Statuses()...)
		notes = append(notes, Supports()...)
		notes = append(notes, Failures()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
