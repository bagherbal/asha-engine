package tauetaspatialtagging

import (
	"fmt"
	"strings"
)

func FormatTauEta(a TauEtaSequenceAudit) string {
	return fmt.Sprintf("source=%q expr=%q seq=%v magnitudes=%v stable=%t scalarTraceOnly=%t three=%t magPattern=%q signedPattern=%q twoPlusOne=%t onePlusOnePlusOne=%t verdict=%s", a.SourceGate, a.SourceExpression, a.Sequence, a.Magnitudes, a.StableNativeDegrees, a.ScalarTraceFunctionalOnly, a.ThreeComponentSignature, a.MagnitudePattern, a.SignedPattern, a.TwoPlusOneMagnitudeSelector, a.OnePlusOnePlusOneSignedSpectrum, a.Verdict)
}

func FormatSpatial(a SpatialMappingAudit) string {
	return fmt.Sprintf("axes=%v planes=%v dimCompatible=%t scalarBundle=%t actsOnW=%t pullback=%t manualAxis=%t magnitudes=%v uniqueMag=%d uniqueIndex=%d uniqueAxisIfMapped=%q complementIfMapped=%q conditionalWeak=%t derivedWeak=%t wouldBreakS3=%t actualBreakS3=%t verdict=%s", a.SpatialAxes, a.CandidatePureSpatialPlanes, a.DimensionCompatible, a.TauEtaActsOnScalarBundle, a.TauEtaActsOnFockW, a.NativePullbackDerived, a.ManualAxisAssignment, a.Magnitudes, a.UniqueMagnitudeValue, a.UniqueMagnitudeIndex, a.UniqueAxisIfMapped, a.ComplementPlaneIfMapped, a.WeakPlaneConditionallySeen, a.WeakPlaneDerived, a.S3DegeneracyWouldBreak, a.S3DegeneracyActuallyBroken, a.Verdict)
}

func FormatPlanes(rows []PlaneTauAudit) string {
	parts := make([]string, 0, len(rows))
	for _, p := range rows {
		parts = append(parts, fmt.Sprintf("%s complement=%s inherited=%t survivesU1=%t selectedIfMapped=%t selectedNative=%t reason=%q verdict=%s", p.Plane, p.ComplementAxis, p.InheritedFromGate240, p.SurvivesU1Twist, p.SelectedIfTauMapped, p.SelectedNatively, p.SelectionReason, p.Verdict))
	}
	return strings.Join(parts, " | ")
}

func FormatGeneration(a GenerationBreakingAudit) string {
	return fmt.Sprintf("dim=%d tau=%v distinct=%d breaksAllThree=%t trialityTooSymmetric=%t pullback=%t canonicalOperator=%t mixing=%t CKM/PMNS=%t capacity=%t textureDerived=%t verdict=%s", a.TrialityCarrierDimension, a.TauValues, a.DistinctEigenvalueCount, a.SignedSpectrumBreaksAllThree, a.ExactTrialityKnownTooSymmetric, a.TauEtaToGenerationPullback, a.CanonicalTrialityOperatorDerived, a.MixingOperatorDerived, a.CKMPMNSDerived, a.CapacitySupported, a.TextureDerived, a.Verdict)
}

func FormatWeak(a WeakOutcomeAudit) string {
	return fmt.Sprintf("inheritsReebFailure=%t tauCanSelect=%t tauPullback=%t conditionalWeak=%t weakDerived=%t physicalLeft=%t globalH=%t orderOne=%t verdict=%s", a.InheritsGate241ReebFailure, a.TauMagnitudeCanSelectAxis, a.TauToSpatialPullbackDerived, a.UniqueWeakPlaneConditionallySeen, a.UniqueWeakPlaneDerived, a.PhysicalLeftHandedDerived, a.GlobalHSummandDerived, a.OrderOneReady, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedTauSpatial=%t forcedAxis=%t importedWeak=%t importedTexture=%t traceToMatrix=%t claimedChirality=%t claimedH=%t claimedCKM/PMNS=%t claimedMasses=%t polluted=%t verdict=%s", a.ForcedTauToSpatialMap, a.ForcedAxisAssignment, a.ImportedSMWeakPlane, a.ImportedGenerationTexture, a.PromotedTraceToSpinorMatrix, a.ClaimedPhysicalChirality, a.ClaimedGlobalH, a.ClaimedCKMPMNS, a.ClaimedMasses, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("tau=%t magSelector=%t spatialPullback=%t weakConditional=%t weakDerived=%t generationCapacity=%t generationTexture=%t globalH=%t status=%q next=%q comment=%q", a.TauEtaRetrieved, a.MagnitudeSelectorCapacity, a.SpatialPullbackDerived, a.WeakPlaneConditionallySeen, a.WeakPlaneDerived, a.GenerationBreakingCapacity, a.GenerationTextureDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
