package generation2spatialprojectiveorientationsealminimalityconsequenceaudit

import "fmt"

func FormatInherited(a InheritedGate573Audit) string {
	return fmt.Sprintf("CP2=%t SU3transitive=%t noInvariantPoint=%t general2plus1=%t noNativeP=%t sealNeeded=%t noWeakPlaneNative=%t k7Time=%t flavorEW=%t verdict=%q", a.CP2SpatialBlockCertified, a.SU3ActsTransitively, a.NoSU3InvariantPoint, a.GeneralTwoPlusOneClassified, a.NoNativeRankOneProjector, a.OrientationSealAlreadyNeeded, a.NoWeakPlaneNativeSelection, a.K7ProductTimePreserved, a.FlavorEWBoundaryPreserved, a.Verdict)
}

func FormatSealDefinition(a SealDefinitionAudit) string {
	return fmt.Sprintf("seal=%q datum=%q equiv=%q u=%v P=%v rank=%d trace=%.12g idemResidual=%.3g hermitian=%t breaks=%q minimal=%t native=%t verdict=%q", a.SealName, a.ProjectiveDatum, a.EquivalentProjectorDatum, a.RepresentativeU, a.Projector, a.Rank, a.Trace, a.IdempotentResidual, a.Hermitian, a.BreaksSU3To, a.MinimalForSymmetryBreaking, a.NativeDerived, a.Verdict)
}

func FormatSealedSelector(a SealedSelectorAudit) string {
	return fmt.Sprintf("selector=%q lambdas=(%.12g,%.12g) matrix=%v eig=%v pattern=%q CP1=%q CP0=%q dims=(%d,%d) constructs=%t nativeWithoutSeal=%t verdict=%q", a.SelectorFormula, a.LambdaOne, a.LambdaTwo, a.SelectorMatrix, a.Eigenvalues, a.MultiplicityPattern, a.CriticalCP1, a.CriticalCP0, a.CP1RealDimension, a.CP0RealDimension, a.ConstructsCP1CP0Split, a.NativeWithoutSeal, a.Verdict)
}

func FormatCommutant(a CommutantAudit) string {
	return fmt.Sprintf("pattern=%q commutant=%q dimFormula=%q dim=%d expected=%d gate555match=%t sealedOnly=%t verdict=%q", a.SelectorMultiplicityPattern, a.Commutant, a.DimensionFormula, a.Dimension, a.ExpectedDimension, a.MatchesGate555Formula, a.SealedSupportOnly, a.Verdict)
}

func FormatBasis(a BasisExampleAudit) string {
	return fmt.Sprintf("gauge=%q P=%q selector=%q complement=%q name=%q basisDependent=%t nativeBasis=%t verdict=%q", a.RepresentativeGauge, a.ProjectorMatrix, a.SelectorMatrix, a.ComplementaryPlane, a.ConventionalPlaneName, a.BasisDependent, a.NativeBasisSelection, a.Verdict)
}

func FormatWeakPlaneFirewall(a WeakPlaneFirewallAudit) string {
	return fmt.Sprintf("physicalWeakPlane=%t reqFST=%t reqH=%t reqD=%t reqJ=%t reqGrading=%t reqFirstOrder=%t proven=%t verdict=%q", a.ComplementaryCP1CanBeCalledPhysicalWeakPlane, a.RequiresFiniteSpectralTripleCarrierAction, a.RequiresQuaternionicCompatibility, a.RequiresDiracCompatibility, a.RequiresRealStructureCompatibility, a.RequiresGradingCompatibility, a.RequiresFirstOrderCompatibility, a.CompatibilityProven, a.Verdict)
}

func FormatFlavorGenerationFirewall(a FlavorGenerationFirewallAudit) string {
	return fmt.Sprintf("generation=%t yukawa=%t ckmPmns=%t observed=%t ewDynamics=%t photon=%t wz=%t verdict=%q", a.GenerationHierarchyDerived, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.PhysicalEWDynamicsDerived, a.PhotonDynamicsDerived, a.WZMassesDerived, a.Verdict)
}

func FormatPreviousGateBoundaries(a PreviousGateBoundaryAudit) string {
	return fmt.Sprintf("tauEtaTrace=%t q4Contact=%t pauliQuaternionicSocket=%t gate564565=%t k7Time=%t verdict=%q", a.TauEtaTraceShadowOnly, a.Q4ContactOnly, a.PauliQuaternionicSocketOnly, a.Gate564565BridgeSymbolic, a.K7TimeRoutesSealed, a.Verdict)
}

func FormatMinimality(a MinimalityAudit) string {
	return fmt.Sprintf("withoutPoint=%t withoutP=%t noSelector=%t equiv=%t anySelectorDeterminesP=%t minimal=%t proof=%q verdict=%q", a.WithoutProjectivePoint, a.WithoutRankOneProjector, a.NoCP2ToCP1CP0Selector, a.PointProjectorEquivalence, a.Any2Plus1SelectorDeterminesPU, a.SealIsMinimal, a.ProofSketch, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("A_sufficient=%t B_minimal=%t C_u2u1=%t D_physicalWeakFlavorEW=%t E_required=%q k7OrTime=%t verdict=%q", a.SealSufficient, a.SealMinimal, a.ReducesSymmetryToU2U1, a.DerivesPhysicalWeakFlavorElectroweakData, a.AdditionalTheoremRequired, a.K7OrProductTimeOpened, a.Verdict)
}
