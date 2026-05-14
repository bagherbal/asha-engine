package neutraleigenspacekernel

import "fmt"

func FormatInherited(a InheritedGate248Audit) string {
	return fmt.Sprintf("basis8v=%t scalarOrigin=%t embeddable=%t map=%t vtau=%t triality=%t yukawa=%t CKMPMNS=%t masses=%t truth=%q", a.Basis8VKnown, a.ScalarTraceOriginKnown, a.DimensionallyEmbeddable, a.ScalarTo8VMapDerived, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.TruthStatement)
}

func FormatVectorCarrier(a VectorCarrierAudit) string {
	return fmt.Sprintf("rep=%q dim=%d native=%t labels=%v octonionic=%t verdict=%s", a.RepresentationName, a.Dimension, a.NativeCarrierKnown, a.BasisLabels, a.RealOctonionicSplitKnown, a.Verdict)
}

func FormatEWAction(a EWDerivationActionAudit) string {
	return fmt.Sprintf("ops=%v source=%q Qscalar=%t Zscalar=%t Q8v=%t Z8v=%t simultaneousDiag=%t spectrum=%t manual=%q rejected=%t obstruction=%q verdict=%s", a.Operators, a.Source, a.QKnownAsScalarObservable, a.ZKnownAsScalarObservable, a.QMatrixOn8VDerived, a.ZMatrixOn8VDerived, a.SimultaneouslyDiagonal, a.ChargeSpectrumKnown, a.ManualRepresentation, a.ManualRepresentationRejected, a.Obstruction, a.Verdict)
}

func FormatNeutralKernel(a NeutralKernelAudit) string {
	return fmt.Sprintf("definition=%q computed=%t dimKnown=%t dim=%d exact3=%t basisIndependent=%t invariant=%t dependsOnQ=%t verdict=%s", a.KernelDefinition, a.Computed, a.DimensionKnown, a.Dimension, a.ExactlyThreeDimensional, a.BasisIndependent, a.InvariantSubspaceDerived, a.DependsOnMissingQMatrix, a.Verdict)
}

func FormatScalarPlane(a ScalarPlaneIsomorphismAudit) string {
	return fmt.Sprintf("slots=%v tau=%v count=%d neutralDim=%d dimMatch=%t canonicalIso=%t basisPairing=%t qztyBasis=%t obstruction=%q verdict=%s", a.ScalarTraceSlots, a.TauEta, a.ScalarSlotCount, a.NeutralKernelDimension, a.DimensionMatch, a.CanonicalIsomorphism, a.BasisIndependentPairing, a.QZTYToNeutralBasisDerived, a.Obstruction, a.Verdict)
}

func FormatVTau(a VTauNeutralVectorAudit) string {
	return fmt.Sprintf("candidate=%q constructed=%t lawful=%t coeff=%v host=%q norm2=%d triality=%t rejected=%q verdict=%s", a.Candidate, a.Constructed, a.LawfulRepresentative, a.Coefficients, a.HostSubspace, a.WouldHaveNormSquared, a.WouldFeedTriality, a.RejectedBecause, a.Verdict)
}

func FormatTriality(a TrialityPreflightAudit) string {
	return fmt.Sprintf("requires8v=%t neutralVector=%t canRun=%t Dtau=%t genCapacity=%t noncommuting=%t CKMPMNS=%t masses=%t verdict=%s", a.Requires8VVector, a.Neutral8VVectorAvailable, a.TrialityCanRun, a.DiagonalTextureConstructed, a.GenerationBreakingCapacity, a.NonCommutingTextureCapacity, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("Q8v=%t forceKernel3=%t assignSlots=%t vtauByHand=%t trialityMatrices=%t yukawa=%t masses=%t CKMPMNS=%t flavorClaim=%t polluted=%t verdict=%s", a.InventedQActionOn8V, a.ForcedNeutralKernelDim3, a.AssignedScalarSlotsToBasis, a.ConstructedVTauByHand, a.InventedTrialityMatrices, a.InsertedYukawaTexture, a.ImportedObservedMasses, a.ImportedCKMPMNS, a.ClaimedFiniteFlavorTheorem, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("basis8v=%t ewAction=%t kernel=%t kernel3=%t scalarIso=%t vtau=%t triality=%t yukawa=%t CKMPMNS=%t masses=%t status=%q next=%q comment=%q", a.Basis8VKnown, a.EWDerivationActionDerived, a.NeutralKernelDerived, a.NeutralKernelDim3, a.ScalarNeutralIsomorphism, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
