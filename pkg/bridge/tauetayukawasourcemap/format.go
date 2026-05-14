package tauetayukawasourcemap

import "fmt"

func FormatInheritance(a Gate260Inheritance) string {
	return fmt.Sprintf("eightVClosed=%t directCarrier=%t tauSource=%t textureAlready=%t ckm=%t masses=%t tau=%v dim=%d carrier=%q status=%q next=%q verdict=%q", a.EightVRouteClosed, a.DirectGenerationCarrierOpened, a.TauEtaYukawaSourceCandidate, a.DirectYukawaTextureAlreadyDerived, a.CKMPMNSAlreadyDerived, a.FermionMassesAlreadyDerived, a.TauEtaEigenvalues, a.GenerationDimension, a.SourceCarrierName, a.Gate260Status, a.Gate260NextGate, a.Verdict)
}

func FormatBilinear(a BilinearCarrierAudit) string {
	return fmt.Sprintf("left=%q right=%q domain=%q codomain=%q algebra=%q genDim=%d algDim=%d kinds=%v kindCount=%d channels=%d uses8v=%t operatorCarrier=%t chargeRules=%t lawful=%t amplitudes=%t verdict=%q", a.LeftCarrierName, a.RightCarrierName, a.Domain, a.Codomain, a.TextureAlgebra, a.GenerationDimension, a.TextureAlgebraDimension, a.FermionKinds, a.FermionKindCount, a.Gate25ChannelsSupported, a.UsesEightVKernel, a.OperatorValuedCarrier, a.ChargeSelectionRulesInherited, a.BilinearCarrierLawful, a.YukawaAmplitudeInserted, a.Verdict)
}

func FormatTauEta(a TauEtaActionAudit) string {
	return fmt.Sprintf("tau=%v matrix=%s trace=%d det=%d signedDistinct=%d magDistinct=%d selfAdjoint=%t diagonal=%t left=%t right=%t no8v=%t signedBreak=%t magBreak=%t verdict=%q", a.Eigenvalues, MatrixString(a.Matrix), a.Trace, a.Determinant, a.SignedDistinctEigenvalueCount, a.MagnitudeDistinctCount, a.SelfAdjoint, a.DiagonalInTauBasis, a.ActsOnLeftGenerationIndex, a.ActsOnRightGenerationIndex, a.ActsWithoutEightVKernel, a.BreaksSignedGenerationDegeneracy, a.BreaksMagnitudeDegeneracy, a.Verdict)
}

func FormatTextureAlgebra(a TextureAlgebraAudit) string {
	return fmt.Sprintf("dim=%d commutant=%d offdiag=%d adEigen=%v nonzero=%v absGaps=%v noncommDirs=%t partner=%t trialityQuarantined=%t diagonalSeed=%t verdict=%q", a.TextureAlgebraDimension, a.CommutantDimension, a.OffDiagonalComplementDimension, a.CommutatorEigenvalues, a.NonzeroCommutatorEigenvalues, a.DistinctAbsMixingGaps, a.NonCommutingDirectionsExist, a.CanonicalNonCommutingPartnerSelected, a.TrialitySymmetryMapsQuarantined, a.DiagonalOnlySeed, a.Verdict)
}

func FormatYukawa(a YukawaSourceMapAudit) string {
	return fmt.Sprintf("name=%q expr=%q acts=%t diagonal=%t signed111=%t split=%t mixingByItself=%t physicalTexture=%t needsPartner=%t needsAction=%t needsNorm=%t needsVEV=%t needsEmpirical=%t ckm=%t masses=%t verdict=%q", a.SourceMapName, a.MapExpression, a.ActsOnBilinearCarrier, a.DiagonalGenerationTextureSeed, a.SignedOnePlusOnePlusOneSpectrum, a.CanSplitThreeGenerations, a.CanProduceMixingByItself, a.PhysicalYukawaTextureDerived, a.RequiresSecondNonCommutingOperator, a.RequiresFiniteActionFunctional, a.RequiresKineticNormalization, a.RequiresScalarVEVAmplitude, a.RequiresEmpiricalYukawaSealForNumbers, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatSealLedger(a SealLedgerAudit) string {
	return fmt.Sprintf("finite=%v conditional=%v missing=%v massesUsed=%t mixingUsed=%t empiricalSeal=%t verdict=%q", a.FiniteDerivedItems, a.SealedOrConditionalItems, a.StillMissingItems, a.EmpiricalMassDataUsed, a.ObservedMixingDataUsed, a.EmpiricalYukawaSealActivated, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate260=%t no8v=%t noTauFock=%t noMasses=%t noMixingAngles=%t noDiagToCKM=%t noInventPartner=%t noActionClaim=%t noEmpiricalSeal=%t polluted=%t verdict=%q", a.Gate260EightVNoGoPreserved, a.DoesNotReopenEightVKernelRoute, a.DoesNotRewriteTauEtaAsFockVector, a.DoesNotUseObservedMasses, a.DoesNotUseObservedMixingAngles, a.DoesNotPromoteDiagonalSeedToCKM, a.DoesNotInventNonCommutingPartner, a.DoesNotClaimSpectralAction, a.DoesNotActivateEmpiricalYukawaSeal, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate260=%t bilinear=%t tauAction=%t algebra=%t diagonal=%t complement=%t partner=%t texture=%t ckm=%t masses=%t status=%q next=%q comment=%q", a.Gate260Inherited, a.BilinearCarrierDefined, a.TauEtaActionDerived, a.TextureAlgebraDecomposed, a.DiagonalSourceMapOpened, a.CommutatorComplementExposed, a.CanonicalMixingPartnerFound, a.PhysicalYukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
