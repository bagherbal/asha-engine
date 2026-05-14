package cliffordpullback

import (
	"fmt"
)

func FormatCliffordAction(a CliffordActionAudit) string {
	return fmt.Sprintf("carrier=%q domain=%q codomain=%q realDim=%d complexDim=%d exteriorBasis=%d creationAnnihilation=%t cliffordAction=%t requiresExterior=%t mapsScalarTrace=%t verdict=%s", a.SourceCarrier, a.Domain, a.Codomain, a.RealSpinorDimension, a.ComplexSpinorDimension, a.ExteriorBasisDimension, a.CreationAnnihilationModel, a.CliffordMultiplicationAvailable, a.RequiresExteriorElement, a.MapsScalarTraceFunctional, a.Verdict)
}

func FormatTauEtaPullback(a TauEtaPullbackAudit) string {
	return fmt.Sprintf("seq=%v mags=%v source=%q scalarTrace=%t exteriorForm=%t gradeKnown=%t bladeCoeffs=%t spatialSlots=%t trialitySlots=%t cliffordApplicable=%t indexMap=%t endomorphism=%t hypothetical=%q rejected=%t reason=%q verdict=%s", a.Sequence, a.Magnitudes, a.SourceType, a.ScalarTraceFunctional, a.ExteriorFormElement, a.HomogeneousGradeKnown, a.BasisBladeCoefficientsKnown, a.SpatialSlotLabelsDerived, a.TrialitySlotLabelsDerived, a.CliffordActionApplicable, a.AtiyahSingerIndexMapDerived, a.EndomorphismConstructed, a.HypotheticalSpatialOperator, a.HypotheticalOperatorRejected, a.RejectionReason, a.Verdict)
}

func FormatFunctor(a PullbackFunctorAudit) string {
	return fmt.Sprintf("cliffordFunctor=%t tauInDomain=%t tauToForm=%t tauToIndex=%t scalarToSpinor=%t gaugeProjection=%t normalization=%t pullback=%t verdict=%s", a.CliffordActionFunctorAvailable, a.TauEtaInFunctorDomain, a.TauEtaToExteriorFormDerived, a.TauEtaToIndexClassDerived, a.ScalarBundleToSpinorBundleMapDerived, a.GaugeProjectionMapDerived, a.CanonicalNormalizationDerived, a.PullbackFunctorDerived, a.Verdict)
}

func FormatSpatial(a SpatialEndomorphismSieve) string {
	return fmt.Sprintf("axes=%v planes=%v tauSelector=%t endomorphism=%t projected=%t spectrum=%t conditionalAxis=%q conditionalPlane=%q derivedAxis=%q weakDerived=%t s3Broken=%t verdict=%s", a.SpatialAxes, a.CandidatePureSpatialPlanes, a.TauMagnitudeSelectorCapacity, a.EndomorphismAvailable, a.ProjectedToSpatialModes, a.MatrixSpectrumAvailable, a.UniqueAxisConditionallySeen, a.ComplementPlaneConditionally, a.UniqueAxisDerived, a.WeakPlaneDerived, a.S3DegeneracyBroken, a.Verdict)
}

func FormatTriality(a TrialityEndomorphismSieve) string {
	return fmt.Sprintf("dim=%d tau=%v distinctCapacity=%t endomorphism=%t projected=%t diagonal=%t texture=%t noncommutingPair=%t CKM/PMNS=%t verdict=%s", a.GenerationCarrierDimension, a.TauSignedSpectrum, a.DistinctEigenvalueCapacity, a.EndomorphismAvailable, a.ProjectedToTrialitySectors, a.DiagonalGenerationOperator, a.GenerationTextureDerived, a.NonCommutingTexturePairDerived, a.CKMPMNSDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedForm=%t forcedSpatial=%t forcedTriality=%t inventedEndomorphism=%t importedWeak=%t importedTexture=%t claimedH=%t claimedChirality=%t claimedCKM/PMNS=%t claimedMasses=%t polluted=%t verdict=%s", a.ForcedTauAsExteriorForm, a.ForcedSpatialSlotMap, a.ForcedTrialitySlotMap, a.InventedCliffordEndomorphism, a.ImportedWeakPlane, a.ImportedGenerationTexture, a.ClaimedGlobalH, a.ClaimedPhysicalChirality, a.ClaimedCKMPMNS, a.ClaimedMasses, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("clifford=%t selectorCapacity=%t tauInDomain=%t endomorphism=%t weakConditional=%t weakDerived=%t generationCapacity=%t generationTexture=%t pullback=%t globalH=%t status=%q next=%q comment=%q", a.CliffordActionAvailable, a.TauEtaSelectorCapacity, a.TauEtaInCliffordDomain, a.EndomorphismConstructed, a.WeakPlaneConditionallySeen, a.WeakPlaneDerived, a.GenerationBreakingCapacity, a.GenerationTextureDerived, a.PullbackFunctorDerived, a.GlobalHDerived, a.Status, a.NextGate, a.Comment)
}
