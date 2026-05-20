package generation2sealedspatialcp1fstcompatibilityaudit

import "fmt"

func FormatInherited(a InheritedGate574Audit) string {
	return fmt.Sprintf("seal=%q sufficient=%t minimal=%t split=%t commutantU2U1=%t native=%t physicalWeak=%t flavorEW=%t k7Time=%t priorRequired=%q verdict=%q", a.SealName, a.SealSufficient, a.SealMinimal, a.CP1CP0Split, a.CommutantU2U1, a.SealNative, a.PhysicalWeakPlaneDerived, a.FlavorElectroweakDataDerived, a.K7OrProductTimeOpened, a.AdditionalTheoremRequiredPrior, a.Verdict)
}

func FormatSealedDecomposition(a SealedDecompositionAudit) string {
	return fmt.Sprintf("seal=%q ray=%q P=%v Q=%v rankP=%d trP=%.12g idemP=%.3g rankQ=%d trQ=%.12g idemQ=%.3g orth=%.3g decomp=%q dimUperp=%d dimCu=%d split=%q exists=%t nativeWithoutSeal=%t verdict=%q", a.SealName, a.RepresentativeRay, a.Projector, a.ComplementProjector, a.ProjectorRank, a.ProjectorTrace, a.ProjectorIdempotentError, a.ComplementRank, a.ComplementTrace, a.ComplementIdempotentError, a.OrthogonalityResidual, a.WSpatialDecomposition, a.DimCUperp, a.DimCCu, a.ProjectiveSplit, a.CP1CP0SplitExists, a.NativeWithoutSeal, a.Verdict)
}

func FormatBMinusL(a BMinusLCompatibilityAudit) string {
	return fmt.Sprintf("restriction=%q eig=%.12g commResidual=%.3g commPU=%t commQ=%t vacuous=%t furtherSelector=%t verdict=%q", a.BLRestriction, a.BLSpatialEigenvalue, a.CommutatorResidual, a.CommutesWithPU, a.CommutesWithComplement, a.CompatibilityVacuous, a.SuppliesFurtherSelector, a.Verdict)
}

func FormatCommutant(a CommutantAudit) string {
	return fmt.Sprintf("selector=%q pattern=%q commutant=%q dimFormula=%q dim=%d gate555=%t sealedOnly=%t verdict=%q", a.SelectorFormula, a.MultiplicityPattern, a.Commutant, a.DimensionFormula, a.Dimension, a.MatchesGate555Formula, a.SealedSupportOnly, a.Verdict)
}

func FormatQuaternionic(a QuaternionicSocketComparisonAudit) string {
	return fmt.Sprintf("imH=%t hphi=%t moment=%t oneFormLink=%t source=%q target=%q imHToSuUperp=%t hToEndUperp=%t compatP=%t wSpatialBlocked=%t verdict=%q", a.ImHSocketAvailableElsewhere, a.HPhiDoubletModuleAvailable, a.PauliHopfMomentQuaternionic, a.FiniteOneFormLinkedStructurally, a.SourceCarrier, a.TestedTarget, a.ImHToSuUperpIntertwinerSupplied, a.HToEndUperpModuleSupplied, a.CompatibleWithPU, a.WSpatialTransferBlocked, a.Verdict)
}

func FormatFiniteTriple(a FiniteSpectralTripleCarrierAudit) string {
	return fmt.Sprintf("AF=%t finiteWeakElsewhere=%t uperpCarrier=%t D=%t J=%t grading=%t firstOrder=%t orderOne=%t verdict=%q", a.AFRepresentationStructural, a.FiniteWeakDoubletCarrierExistsElsewhere, a.UperpUsedAsFiniteWeakDoubletCarrier, a.DCompatibilityForUperp, a.JCompatibilityForUperp, a.GradingCompatibilityForUperp, a.FirstOrderCompatibilityForUperp, a.OrderOneCarrierActionProven, a.Verdict)
}

func FormatOneForm(a OneFormHiggsLaneCompatibilityAudit) string {
	return fmt.Sprintf("oneFormScalar=%t imHOnHPhi=%t cp1OneForm=%t cp1Higgs=%t pauliSeparate=%t gate562=%t verdict=%q", a.FiniteOneFormContainsScalarDoublet, a.ImHActsOnHPhiStructurally, a.SealedCP1AppearsInOneFormLane, a.SealedCP1AppearsInHiggsLane, a.PauliRouteSeparateFromWSpatial, a.Gate562BoundaryPreserved, a.Verdict)
}

func FormatWeakPlane(a PhysicalWeakPlaneFirewallAudit) string {
	return fmt.Sprintf("gauge=%q complement=%q name=%q physical=%t fst=%t H=%t D=%t J=%t grading=%t firstOrder=%t verdict=%q", a.RepresentativeGauge, a.Complement, a.ConventionalName, a.CanCallPhysicalWeakPlane, a.FiniteSpectralTripleCompatible, a.QuaternionicCompatible, a.DCompatible, a.JCompatible, a.GradingCompatible, a.FirstOrderCompatible, a.Verdict)
}

func FormatFlavorEW(a FlavorElectroweakFirewallAudit) string {
	return fmt.Sprintf("generation=%t yukawa=%t ckmPmns=%t observed=%t ew=%t photon=%t wz=%t weakIsospin=%t verdict=%q", a.GenerationHierarchyDerived, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.PhysicalEWDynamicsDerived, a.PhotonDynamicsDerived, a.WZMassesDerived, a.WeakIsospinDerived, a.Verdict)
}

func FormatBoundaries(a PreviousGateBoundaryAudit) string {
	return fmt.Sprintf("tauEta=%t q4=%t pauliNotWSpatial=%t gate564565=%t k7Time=%t orientationOnly=%t verdict=%q", a.TauEtaTraceShadowOnly, a.Q4ContactOnly, a.PauliQuaternionicSocketNotWSpatial, a.Gate564565BridgeSymbolic, a.K7TimeRoutesSealed, a.OrientationSealProjectiveOnly, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("A_split=%t B_commBL=%t C_imHAction=%t D_finiteCarrier=%t E_physicalWeak=%t F_flavorEW=%t required=%q verdict=%q", a.SealedCP1SplitExistsAlgebraically, a.CommutesWithBMinusL, a.CarriesNativeOrSealedImHAction, a.PartOfFiniteWeakDoubletCarrier, a.CanBeCalledPhysicalWeakPlane, a.DerivesFlavorOrEWObservedData, a.AdditionalTheoremRequired, a.Verdict)
}
