package tauetaweakselector

import (
	"fmt"
	"strings"
)

func FormatFloatVec(v []float64) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%.6g", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatIntVec(v []int) string {
	parts := make([]string, len(v))
	for i, x := range v {
		parts[i] = fmt.Sprintf("%d", x)
	}
	return "(" + strings.Join(parts, ",") + ")"
}

func FormatInheritance(a Gate258Inheritance) string {
	return fmt.Sprintf("bminusl=%t scalarReduced=%t weakReduced=%t combinedReduced=%t rescanned=%t gate258Plane=%t status=%q verdict=%q", a.BMinusLLedgerRetrieved, a.ScalarSieveReduced, a.WeakFrameSieveReduced, a.CombinedWitnessSpaceReduced, a.RestrictedTrialityRescanned, a.Gate258Neutral3PlaneDerived, a.Gate258Status, a.Verdict)
}

func FormatTauEta(a TauEtaSelectorAudit) string {
	return fmt.Sprintf("source=%q expr=%q seq=%s mags=%s stable=%t scalarOnly=%t nativePullback=%t seal=%t twoPlusOne=%t signed111=%t uniqueSlot=%d uniqueMag=%d verdict=%q", a.SourceGate, a.SourceExpression, FormatIntVec(a.Sequence), FormatIntVec(a.Magnitudes), a.StableNativeDegrees, a.ScalarTraceFunctionalOnly, a.NativeFockPullbackDerived, a.RequiresSpontaneousCarrierSeal, a.TwoPlusOneMagnitudeSelector, a.OnePlusOnePlusOneSignedSpectrum, a.UniqueTauSlot, a.UniqueMagnitudeValue, a.Verdict)
}

func FormatSpatialTag(a SpatialTagAudit) string {
	return fmt.Sprintf("spatial=%v labels=%v tauSlots=%v uniqueMode=%d uniqueLabel=%q complement=%v complementName=%q conditional=%t seal=%q nativePullback=%t unsealedManual=%t s3Conditional=%t s3Native=%t verdict=%q", a.SpatialModes, a.SpatialLabels, a.TauSlotToSpatialModes, a.UniqueSpatialMode, a.UniqueSpatialLabel, a.ComplementPlaneModes, a.ComplementPlaneName, a.ConditionalAlignmentApplied, a.AlignmentSeal, a.NativeTauToFockPullbackDerived, a.ManualUnsealedAxisAssignment, a.S3DegeneracyConditionallyBroken, a.S3DegeneracyNativelyBroken, a.Verdict)
}

func FormatWeakSieve(a WeakPlaneSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t uniquePlane=%t uniqueOriented=%t orientationDeg=%t survivorNames=%v rejected=%v verdict=%q", a.InputBMinusLSurvivorCount, a.SurvivorCount, a.Reduced, a.UniqueUnorientedPlaneSelected, a.UniqueOrientedFrameSelected, a.OrientationSignDegeneracyLeft, a.SurvivorNames, a.RejectedNames, a.Verdict)
}

func FormatScalarSieve(a TauEtaScalarSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t unique=%t signDeg=%t survivorNames=%v verdict=%q", a.InputBMinusLSurvivorCount, a.SurvivorCount, a.Reduced, a.UniqueSelected, a.SignDegeneracyLeft, a.SurvivorNames, a.Verdict)
}

func FormatCombined(a CombinedWitnessSieveAudit) string {
	return fmt.Sprintf("input=%d survivors=%d reduced=%t unique=%t survivorNames=%v verdict=%q", a.InputBMinusLWitnessCount, a.SurvivingWitnessCount, a.Reduced, a.UniqueOrientation, a.SurvivingWitnessNames, a.Verdict)
}

func FormatRestrictedScan(a RestrictedTrialityAudit) string {
	return fmt.Sprintf("branches=%d results=%d exactPolarized3=%d exactFull3=%d maxPolarized=%d maxFull=%d uniqueBranch=%t selected=%q allScanned=%t afterTauEta=%t verdict=%q", a.BranchCount, a.ResultCount, a.ExactPolarized3PlaneResults, a.ExactFull3KernelResults, a.MaxPolarizedKernelComplexDim, a.MaxFullQ8vCKernelComplexDim, a.UniqueBranchForPolarized3Plane, a.SelectedBranch, a.AllSurvivorsScanned, a.ScannedAfterTauEtaSelector, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate258NoGo=%t tauAudit=%t nativePullbackPreserved=%t ssbAlign=%t selectorNotOutcome=%t beforeKernel=%t importedWeak=%t masses=%t yukawas=%t forcedWeakNoSeal=%t forcedScalar=%t branchHand=%t branchKernel=%t forced3=%t tauAsFock=%t vtau=%t yukawa=%t polluted=%t verdict=%q", a.Gate258NoGoPreserved, a.TauEtaRetrievedFromAudit, a.TauEtaNativePullbackPreserved, a.ConditionalSSBAlignmentUsed, a.TauEtaUsedAsSelectorNotOutcome, a.SelectorAppliedBeforeKernel, a.ImportedObservedWeakPlane, a.ImportedObservedMasses, a.ImportedObservedYukawas, a.ForcedWeakPlaneWithoutSeal, a.ForcedScalarOrientation, a.SelectedTrialityByHand, a.SelectedTrialityByDesiredKernel, a.ForcedKernelDim3, a.TreatedTauEtaAsFiniteFockOperator, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate258=%t tau=%t tauNativePullback=%t conditionalTag=%t weakReduced=%t uniquePlane=%t uniqueEW=%t combinedReduced=%t rescanned=%t plane=%t branch=%t yukawa=%t status=%q next=%q comment=%q", a.Gate258Inherited, a.TauEtaRetrieved, a.TauEtaNativePullbackDerived, a.ConditionalSpatialTagApplied, a.WeakPlaneSieveReduced, a.UniqueUnorientedWeakPlane, a.UniqueFullEWOrientationSelected, a.CombinedWitnessSpaceReduced, a.RestrictedTrialityRescanned, a.Neutral3PlaneDerived, a.TrialityBranchSelected, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
