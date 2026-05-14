package noncartanflavorvacuum

import "fmt"

func FormatInheritance(a Gate259Inheritance) string {
	return fmt.Sprintf("tau=%t U12=%t cartanScan=%t cartanPlane=%t maxPolarized=%d maxFull=%d survivors=%v status=%q verdict=%q", a.TauEtaRetrieved, a.ConditionalU12WeakPlaneSelected, a.CartanRestrictedScanCompleted, a.CartanNeutral3PlaneDerived, a.CartanMaxPolarizedKernelDim, a.CartanMaxFullKernelDim, a.SurvivingWitnessNames, a.Gate259Status, a.Verdict)
}

func FormatNonCartan(a NonCartanGeneratorAudit) string {
	return fmt.Sprintf("plane=%s pair=%v T3=%q offdiag=%v ladder=%v pauli=%t closed=%t hermitian=%t traceless=%t inside=%t changesDirection=%t changesSpectrum=%t verdict=%q", a.WeakPlaneName, a.ModePair, a.CartanGenerator, a.OffDiagonalGenerators, a.RaisingLoweringGenerators, a.PauliBasisRetrieved, a.LieAlgebraClosed, a.HermitianBasis, a.TracelessBasis, a.ActsInsideSelectedPlane, a.ChangesGaugeDirection, a.ChangesChargeSpectrum, a.Verdict)
}

func FormatDirection(a GaugeDirectionAudit) string {
	return fmt.Sprintf("%s coeff=%v norm=%.12g radius=%.12g eig=%v y+zeros=%d y-zeros=%d sameCartan=%t canIncrease=%t residual=%.3g", a.Name, a.CoefficientsT1T2T3, a.Norm, a.EigenvalueRadius, a.Eigenvalues, a.YPlusHalfZeroMultiplicity, a.YMinusHalfZeroMultiplicity, a.SameAsCartanSpectrum, a.CanIncreaseBeyondCartanKernel, a.Residual)
}

func FormatGaugeOrbit(a GaugeOrbitInvariantAudit) string {
	return fmt.Sprintf("conjugate=%t spectrumInvariant=%t kernelInvariant=%t rotatesOnly=%t cartanMaxP=%d cartanMaxFull=%d nonCartanUpper=%d canEnlarge=%t directions=%d allMatch=%t verdict=%q", a.AllSU2ElementsConjugateToCartan, a.SpectrumInvariantUnderConjugacy, a.KernelDimensionGaugeInvariant, a.OffDiagonalTermsRotateBasisOnly, a.CartanMaxPolarizedKernelDim, a.CartanMaxFullKernelDim, a.NonCartanUpperBoundFullKernelDim, a.NonCartanCanEnlargeKernel, a.DirectionCount, a.AllDirectionsMatchCartanSpectrum, a.Verdict)
}

func FormatEightVClosure(a EightVRouteClosureAudit) string {
	return fmt.Sprintf("usesSurvivors=%t survivors=%d branches=%d inheritedResults=%d invariantInsteadOfScan=%t needsNewRep=%t plane=%t yukawa8v=%t verdict=%q", a.UsesGate259Survivors, a.SurvivorCount, a.BranchCount, a.InheritedBranchEvaluationCount, a.OffDiagonalScanReplacedByInvariant, a.WouldNeedNewRepresentationNotGaugeRotation, a.Neutral3PlaneAvailable, a.YukawaVia8VOpened, a.Verdict)
}

func FormatGeneration(a DirectGenerationCarrierAudit) string {
	return fmt.Sprintf("carrier=%q kind=%q dim=%d source=%q tau=%v signedDistinct=%d magDistinct=%d trace=%d det=%d generation=%t operatorNot8v=%t capacity=%t bypass8v=%t needsTriality=%t verdict=%q", a.CarrierName, a.CarrierKind, a.Dimension, a.SourceGate, a.TauEtaEigenvalues, a.SignedDistinctEigenvalueCount, a.MagnitudeDistinctCount, a.Trace, a.Determinant, a.ActsOnGenerationIndex, a.OperatorSpaceNotVector8V, a.NativeGenerationBreakingCapacity, a.Bypasses8VNeutralKernel, a.RequiresTrialityTransport, a.Verdict)
}

func FormatYukawaSource(a DirectYukawaSourceAudit) string {
	return fmt.Sprintf("source=%t diagonalSeed=%t signed111=%t breaksDeg=%t texture=%t leftRight=%t action=%t norm=%t phase=%t empiricalSeal=%t ckm=%t masses=%t verdict=%q", a.TauEtaSourceMapCandidate, a.GenerationDiagonalTextureSeed, a.OnePlusOnePlusOneSignedSpectrum, a.CanBreakGenerationDegeneracy, a.YukawaTextureDerived, a.RequiresLeftRightBilinearCarrier, a.RequiresFiniteYukawaAction, a.RequiresKineticNormalization, a.RequiresPhaseMixingSource, a.RequiresEmpiricalYukawaSeal, a.CKMPMNSDerived, a.FermionMassesDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("gate259=%t noWpmAsQ=%t noGaugeSpectrum=%t noBranchHand=%t no3=%t noTauFockVector=%t tauGenOperator=%t noYukawaHand=%t noMasses=%t noCKM=%t seal=%t polluted=%t verdict=%q", a.Gate259NoGoPreserved, a.DoesNotTreatWpmAsChargeOperator, a.DoesNotPromoteGaugeRotationToNewSpectrum, a.DoesNotSelectTrialityByHand, a.DoesNotForceKernelDimThree, a.DoesNotRewriteTauEtaAsFockVector, a.UsesTauEtaAsGenerationOperator, a.DoesNotConstructYukawaTextureByHand, a.DoesNotImportObservedMasses, a.DoesNotImportCKMPMNS, a.SpontaneousCarrierSealPreserved, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("gate259=%t nonCartan=%t orbitInvariant=%t plane8v=%t closed8v=%t directGen=%t tauSource=%t yukawa=%t ckm=%t masses=%t status=%q next=%q comment=%q", a.Gate259Inherited, a.NonCartanGeneratorsRetrieved, a.GaugeOrbitInvariantProved, a.EightVNeutral3PlaneDerived, a.EightVRouteClosed, a.DirectGenerationCarrierOpened, a.TauEtaYukawaSourceCandidate, a.DirectYukawaTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
