package tauetamixingpartner

import "fmt"

func FormatInheritance(a Gate261Inheritance) string {
	return fmt.Sprintf("eightVClosed=%t bilinear=%t tauAction=%t textureAlg=%t complement=%t prevPartner=%t prevTexture=%t tau=%v dim=%d commutant=%d offdiag=%d gaps=%v verdict=%q", a.EightVRouteClosed, a.BilinearCarrierDefined, a.TauEtaActionDerived, a.TextureAlgebraDecomposed, a.CommutatorComplementExposed, a.PreviousCanonicalPartnerFound, a.PreviousPhysicalTextureDerived, a.TauEtaEigenvalues, a.TextureAlgebraDimension, a.CommutantDimension, a.OffDiagonalComplementDimension, a.DistinctAbsMixingGaps, a.Verdict)
}

func FormatCandidate(a CandidateAudit) string {
	return fmt.Sprintf("name=%q kind=%q source=%q matrixAvailable=%t matrix=%s canonical=%t selfAdjoint=%t unitary=%t symmetry=%t scalar=%t phase=%t genEndo=%t amplitude=%t bridge=%t actionCoeff=%t commutes=%t commNorm2=%d offdiagSupport=%d populates=%t qualified=%t disq=%q verdict=%q", a.Name, a.Kind, a.Source, a.MatrixAvailable, MatrixString(a.Matrix), a.CanonicalFiniteData, a.SelfAdjoint, a.UnitaryOrOrthogonal, a.PureSymmetryOrLabelAction, a.ScalarOnly, a.PhaseOnly, a.GenerationEndomorphismDerived, a.AmplitudeSourceDerived, a.RequiresRepresentationBridge, a.RequiresActionCoefficient, a.CommutesWithTauEta, a.CommutatorFrobeniusNormSquared, a.OffDiagonalSupportEntries, a.PopulatesMixingComplement, a.QualifiedFiniteMixingPartner, a.Disqualification, a.Verdict)
}

func FormatInventory(a InventoryAudit) string {
	return fmt.Sprintf("candidates=%d matrix=%d canonical=%d rawNoncomm=%d selfAdjRaw=%d hermitianPhaseBasis=%d qualified=%d bgapRejected=%d hopfRejected=%d empirical=%d verdict=%q", a.CandidateCount, a.MatrixCandidates, a.CanonicalFiniteCandidates, a.RawNonCommutingCandidates, a.SelfAdjointRawCandidates, a.HermitianPhaseBasisCandidates, a.QualifiedFiniteMixingPartners, a.ScalarGapCandidatesRejected, a.RepresentationFreePhaseRejected, a.EmpiricalCandidatesUsed, a.Verdict)
}

func FormatTriality(a TrialityPartnerAudit) string {
	return fmt.Sprintf("cycle=%t reflection=%t hermitianReal=%t hermitianImag=%t hermitianBasisDim=%d directions=%d allSymmetry=%t anyAmplitude=%t verdict=%q", a.PermutationCycleNonCommuting, a.ReflectionNonCommuting, a.HermitianRealPartNonCommuting, a.HermitianImaginaryPartNonCommuting, a.HermitianPhaseBasisDimension, a.RawComplementDirectionsTouched, a.AllRawTrialityMapsAreSymmetryAlgebra, a.AnyTrialityMapQualifiedAsAmplitude, a.Verdict)
}

func FormatFinitePhaseGap(a FinitePhaseGapAudit) string {
	return fmt.Sprintf("bgapScale=%t bgapMatrix=%t bgapOffdiag=%t hopfPhase=%t hopfMap=%t hopfOffdiag=%t needsAction=%t needsBridge=%t verdict=%q", a.BGapAvailableAsPositiveScale, a.BGapHasGenerationMatrix, a.BGapCanPopulateOffDiagonalComplement, a.HopfPhasesAvailableAsPhaseLedger, a.HopfPhaseGenerationMapDerived, a.HopfCanPopulateOffDiagonalComplement, a.RequiresFiniteActionFunctional, a.RequiresRepresentationBridge, a.Verdict)
}

func FormatPartnerVerdict(a MixingPartnerVerdict) string {
	return fmt.Sprintf("raw=%t selfAdjBasis=%t qualified=%t texture=%t ckm=%t masses=%t reason=%q next=%q status=%q", a.RawNonCommutingPartnerExists, a.RawSelfAdjointOffDiagonalBasisExists, a.QualifiedFiniteMixingPartnerFound, a.PhysicalYukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Reason, a.NextGate, a.Status)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate261=%t no8v=%t noMasses=%t noAngles=%t noSymToAmp=%t noBGapTexture=%t noHopfTexture=%t noAction=%t empiricalInactive=%t polluted=%t verdict=%q", a.Gate261SourceMapPreserved, a.DoesNotReopenEightVKernelRoute, a.DoesNotUseObservedMasses, a.DoesNotUseObservedMixingAngles, a.DoesNotPromoteSymmetryToAmplitude, a.DoesNotUseBGapAsTextureWithoutMap, a.DoesNotUseHopfPhaseWithoutMap, a.DoesNotClaimFiniteActionFunctional, a.EmpiricalYukawaSealInactive, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate261=%t trialityComplement=%t hermitianBasis=%t bgapRejected=%t hopfRejected=%t qualified=%t texture=%t ckm=%t masses=%t status=%q next=%q comment=%q", a.Gate261Inherited, a.TrialityComplementPopulated, a.HermitianTrialityBasisExposed, a.BGapRejectedAsRepresentationFree, a.HopfRejectedAsRepresentationFree, a.QualifiedMixingPartnerFound, a.PhysicalYukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
