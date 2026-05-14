package vectorrepresentative8v

import "fmt"

func FormatInherited(a InheritedGate247Audit) string {
	return fmt.Sprintf("triality=%t dimMatch=%t tauCapacity=%t scalarVector=%t functor=%t Dtau=%t texture=%t CKMPMNS=%t masses=%t truth=%q", a.Spin8TrialityAvailable, a.DimensionMatch, a.TauTextureCapacityInherited, a.ScalarTraceIsVectorRep, a.TrialityFunctorDerived, a.DiagonalTextureConstructed, a.QualifiedTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.TruthStatement)
}

func FormatVectorBasis(a VectorBasisAudit) string {
	return fmt.Sprintf("basis=%q dim=%d native=%t octonionic=%t labels=%v complexified=%t verdict=%s", a.BasisName, a.Dimension, a.NativeCarrierKnown, a.RealOctonionicSplitKnown, a.BasisLabels, a.ComplexifiedCarrierReady, a.Verdict)
}

func FormatScalarBundle(a ScalarBundleAudit) string {
	return fmt.Sprintf("source=%q slots=%v tau=%v origin=%t sourceDim=%d target=%q targetDim=%d embeddable=%t operatorsAre8V=%t neutralScalarsBasis=%t verdict=%s", a.SourceBundle, a.SourceTraceSlots, a.TauEta, a.TraceOriginKnown, a.SourceDimension, a.CandidateTargetRepresentation, a.CandidateTargetDimension, a.DimensionallyEmbeddable, a.OperatorsAre8VCoordinates, a.NeutralScalarsAreBasisVectors, a.Verdict)
}

func FormatScalarVectorMap(a ScalarVectorMapAudit) string {
	return fmt.Sprintf("name=%q required=%q derived=%t basisIndependent=%t metric=%t hphiSubspace=%t qztyBasis=%t manual=%q rejected=%t obstruction=%q verdict=%s", a.RequiredMapName, a.RequiredMap, a.NativeMapDerived, a.BasisIndependent, a.MetricOrInnerProductProvided, a.HphiSubspaceOf8VDerived, a.QZTYToBasisDerived, a.ManualAssignment, a.ManualAssignmentRejected, a.Obstruction, a.Verdict)
}

func FormatVTau(a VTauConstructionAudit) string {
	return fmt.Sprintf("candidate=%q coeff=%v basis=%v constructed=%t lawful=%t norm2=%d rank=%d feedsTriality=%t rejectedBecause=%q verdict=%s", a.Candidate, a.Coefficients, a.TargetBasis, a.Constructed, a.LawfulRepresentative, a.WouldHaveNormSquared, a.WouldHaveRank, a.WouldFeedTriality, a.RejectedBecause, a.Verdict)
}

func FormatTriality(a TrialityPreflightAudit) string {
	return fmt.Sprintf("requires8v=%t vtau=%t matrices=%t spinorTexture=%t genCapacity=%t noncommuting=%t CKMPMNS=%t masses=%t verdict=%s", a.Requires8VRepresentative, a.VTauAvailable, a.ExplicitTrialityMatricesKnown, a.SpinorTextureConstructed, a.GenerationBreakingCapacity, a.NonCommutingTextureCapacity, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("connes=%t forcedHphi8v=%t assignedQZTY=%t vtauByHand=%t trialityMatrices=%t yukawa=%t masses=%t CKMPMNS=%t flavorClaim=%t polluted=%t verdict=%s", a.ImportedConnesAlgebra, a.ForcedHphiTo8VMap, a.AssignedQZTYToBasisByHand, a.ConstructedVTauByHand, a.InventedTrialityMatrices, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.ImportedCKMPMNS, a.ClaimedFiniteFlavorTheorem, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("basis8v=%t scalarOrigin=%t embeddable=%t map=%t vtau=%t triality=%t yukawa=%t CKMPMNS=%t masses=%t status=%q next=%q comment=%q", a.Basis8VKnown, a.ScalarTraceOriginKnown, a.DimensionallyEmbeddable, a.ScalarTo8VMapDerived, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
