package characteristicpullback

import (
	"fmt"
	"strings"
)

func FormatOrigin(a OperatorOriginTraceAudit) string {
	parts := make([]string, 0, len(a.Records))
	for _, r := range a.Records {
		parts = append(parts, fmt.Sprintf("slot=%d name=%q expr=%q value=%d rational=%s stable=%t family=%q scalar=%t W=%t exteriorCoeff=%t spatialLabel=%t candidate=%q rejected=%t", r.Slot, r.Name, r.Expression, r.Value, r.ExpectedRational, r.StableInteger, r.SourceOperatorFamily, r.LivesOnScalarBundle, r.LivesOnFockCarrierW, r.IsExteriorFormCoefficient, r.SpatialAxisLabelDerived, r.CandidateAxisIfForced, r.ForcedCandidateRejected))
	}
	return fmt.Sprintf("source=%q functional=%q domain=%q seq=%v stable=%t exactOrigins=%t curvatureObs=%t spatialProjectors=%t basisBlades=%t records=[%s] verdict=%s", a.SourceGate, a.FunctionalExpression, a.FunctionalDomain, a.Sequence, a.StableNativeDegrees, a.ExactOperatorOriginsRecovered, a.OperatorsAreCurvatureObservables, a.OperatorsAreSpatialModeProjectors, a.OperatorsAreBasisBlades, strings.Join(parts, " | "), a.Verdict)
}

func FormatSpatialAlignment(a SpatialModeAlignmentAudit) string {
	return fmt.Sprintf("axes=%v slots=%v countCompatible=%t mapDerived=%t QZT3YLinksAxes=%t definitionsUseFock=%t scalarToFockProjection=%t manualCandidate=%q rejected=%t verdict=%s", a.SpatialAxes, a.TraceSlots, a.ModeSlotCountCompatible, a.NativeOperatorToModeMapDerived, a.QZT3YInherentlyLinkToSpatialAxes, a.OperatorDefinitionsUseFockModes, a.ScalarBundleToFockProjectionDerived, a.ManualMapCandidate, a.ManualMapRejected, a.AlignmentVerdict)
}

func FormatCharacteristic(a CharacteristicRepresentativeAudit) string {
	return fmt.Sprintf("characteristicLanguage=%t finiteEta=%t chernRep=%t pontryaginRep=%t gradeKnown=%t bladeLabels=%t normalization=%t representative=%t hypothetical=%q rejected=%t reason=%q verdict=%s", a.CharacteristicClassLanguageAvailable, a.FiniteEtaTraceFunctional, a.ChernCharacterRepresentativeDerived, a.PontryaginFormRepresentativeDerived, a.ExteriorGradeKnown, a.BasisBladeLabelsKnown, a.NormalizationDerived, a.RepresentativeConstructed, a.HypotheticalRepresentative, a.HypotheticalRepresentativeRejected, a.RejectionReason, a.Verdict)
}

func FormatWeak(a WeakPlaneOutcome) string {
	return fmt.Sprintf("conditionalAxis=%q conditionalPlane=%q exteriorRep=%t spatialMap=%t axisTagged=%q selectedPlane=%q S3Broken=%t physicalWeak=%t globalH=%t verdict=%s", a.InheritedConditionalAxis, a.InheritedConditionalPlane, a.ExteriorRepresentativeDerived, a.SpatialModeAlignmentDerived, a.AxisTagged, a.WeakPlaneSelected, a.S3DegeneracyBroken, a.PhysicalWeakPlaneDerived, a.GlobalHSummandDerived, a.Verdict)
}

func FormatGeneration(a GenerationOutcome) string {
	return fmt.Sprintf("seq=%v distinctCapacity=%t characteristicRep=%t trialityMap=%t generationOperator=%t texture=%t CKM/PMNS=%t verdict=%s", a.Sequence, a.DistinctEigenvalueCapacity, a.CharacteristicRepresentative, a.TrialityCarrierMapDerived, a.GenerationOperatorDerived, a.GenerationTextureDerived, a.CKMPMNSDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedModeMap=%t forcedExterior=%t forcedClass=%t forcedTriality=%t importedWeak=%t importedTexture=%t promotedTrace=%t chirality=%t H=%t CKM/PMNS=%t masses=%t polluted=%t verdict=%s", a.ForcedOperatorToModeMap, a.ForcedExteriorRepresentative, a.ForcedCharacteristicClass, a.ForcedTrialityMap, a.ImportedWeakPlane, a.ImportedGenerationTexture, a.PromotedScalarTraceToMatrix, a.ClaimedPhysicalChirality, a.ClaimedGlobalH, a.ClaimedCKMPMNS, a.ClaimedFermionMasses, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("origin=%t stable=%t opModeMap=%t exteriorRep=%t characteristicClass=%t weakConditional=%t weakDerived=%t generationCapacity=%t generationTexture=%t globalH=%t status=%q next=%q comment=%q", a.TauEtaOriginTraced, a.NativeSequenceStable, a.OperatorModeAlignmentDerived, a.ExteriorRepresentativeDerived, a.CharacteristicClassDerived, a.WeakPlaneConditionallyVisible, a.WeakPlaneDerived, a.GenerationBreakingCapacity, a.GenerationTextureDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
