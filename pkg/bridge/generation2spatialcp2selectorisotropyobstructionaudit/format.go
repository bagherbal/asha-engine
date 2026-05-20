package generation2spatialcp2selectorisotropyobstructionaudit

import "fmt"

func FormatInherited(a InheritedGate572Audit) string {
	return fmt.Sprintf("CP3=%t BL1plus3=%t CP2=%t noSecond=%t noK7=%t noTime=%t noFlavorEW=%t verdict=%q", a.CP3ProjectiveLawSpace, a.BMinusLProjectiveOnePlus3, a.SpatialCP2Block, a.NoNativeSecondSelector, a.NoCP3ToK7Functor, a.NoPhysicalTimeOpened, a.NoFlavorElectroweakData, a.Verdict)
}

func FormatCarrier(a SpatialCP2CarrierAudit) string {
	return fmt.Sprintf("Wsp=%q CP2=%q cdim=%d rdim=%d BL=%q eigen=%q gate572=%t certified=%t verdict=%q", a.SpatialEigenspace, a.ProjectiveBlock, a.ComplexDimension, a.RealDimension, a.BLRestrictedMatrix, a.BLRestrictedEigenvalue, a.Gate572CriticalStratumMatched, a.CertifiedAsSpatialBlock, a.Verdict)
}

func FormatSymmetry(a SpatialSymmetryAudit) string {
	return fmt.Sprintf("commutant=%q spatialSym=%q traceless=%q dimU3=%d dimSU3=%d BLscalar=%t furtherSelector=%t preferredDirection=%t verdict=%q", a.CommutantFromBMinusL, a.SpatialSymmetry, a.TracelessPart, a.U3Dimension, a.SU3Dimension, a.BMinusLScalarOnWSpatial, a.SuppliesFurtherSelector, a.PreferredSpatialDirection, a.Verdict)
}

func FormatTransitivity(a TransitivityAudit) string {
	return fmt.Sprintf("group=%q space=%q action=%q groupDim=%d stabilizer=%q stabilizerDim=%d quotientDim=%d CP2Dim=%d transitive=%t invariantPoint=%t invariantP=%t verdict=%q", a.Group, a.Space, a.Action, a.GroupRealDimension, a.PointStabilizer, a.PointStabilizerDimension, a.QuotientRealDimension, a.CP2RealDimension, a.ActsTransitively, a.InvariantPointSelected, a.InvariantRankOneProjector, a.Verdict)
}

func FormatSelector(a HermitianSecondSelectorAudit) string {
	return fmt.Sprintf("selector=%q projector=%q lambdas=(%.12g,%.12g) sample=%q matrix=%v eig=%v pattern=%q CP1=%q CP0=%q dims=(%d,%d) idemResidual=%.3g trace=%.12g classifies=%t nativeWithoutU=%t verdict=%q", a.SelectorFormula, a.RankOneProjectorFormula, a.LambdaOne, a.LambdaTwo, a.SamplePoint, a.SampleMatrix, a.SampleEigenvalues, a.MultiplicityPattern, a.CriticalCP1, a.CriticalCP0, a.CP1RealDimension, a.CP0RealDimension, a.ProjectorIdempotentResidual, a.ProjectorTrace, a.ClassifiesCP1CP0Split, a.NativeWithoutU, a.Verdict)
}

func FormatSearch(a NativeSelectorSearchAudit) string {
	return fmt.Sprintf("candidates=%d nativeP=%t nativePoint=%t nativeSecond=%t reason=%q verdict=%q", a.CandidateCount, a.NativeRankOneProjectorFound, a.NativeProjectivePointFound, a.NativeSecondSelectorFound, a.Reason, a.Verdict)
}

func FormatCandidate(a NativeSelectorCandidate) string {
	return fmt.Sprintf("source=%q gate=%q candidate=%q nativeP=%t CP1CP0=%t status=%q reason=%q", a.Source, a.PriorGate, a.Candidate, a.NativePUProvided, a.WouldSelectCP1CP0, a.Status, a.Reason)
}

func FormatSeal(a OrientationSealAudit) string {
	return fmt.Sprintf("seal=%q minimal=%q equiv=%q formula=%q sealed=%t comm=%q dim=%d expected=%d match=%t verdict=%q", a.SealName, a.MinimalDatum, a.EquivalentDatum, a.SealedSelectorFormula, a.SealedNotNative, a.Commutant, a.CommutantDimension, a.ExpectedCommutantDimension, a.CommMatchesGate555Formula, a.Verdict)
}

func FormatWeakPlane(a WeakPlaneRelationAudit) string {
	return fmt.Sprintf("point=%q complement=%q name=%q basisDependent=%t native=%t weakIsospin=%t generation=%t yukawa=%t ckmPmns=%t verdict=%q", a.SealedPoint, a.ComplementaryPlane, a.ConventionalPlaneName, a.BasisDependent, a.NativeDerived, a.WeakIsospinIdentified, a.GenerationHierarchy, a.YukawaTextureDerived, a.CKMPMNSDerived, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("weakIso=%t physicalPlane=%t hierarchy=%t yukawa=%t ckmPmns=%t observed=%t ewDynamics=%t wz=%t photon=%t k7=%t productTime=%t gate564565=%t verdict=%q", a.PromotedToWeakIsospin, a.PromotedToPhysicalWeakPlane, a.GenerationHierarchyDerived, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorDataImported, a.PhysicalElectroweakDynamics, a.WZMassesDerived, a.PhotonDynamicsDerived, a.CP2ToK7FunctorOpened, a.ProductTimeOpened, a.Gate564565BoundaryPreserved, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("A_CP2=%t B_SU3transitive=%t C_invariantPoint=%t D_generalSelector=%t E_nativeP=%t F_seal=%q G_physicalEWFlavor=%t k7OrTime=%t next=%q verdict=%q", a.SpatialCP2Certified, a.SU3Transitive, a.SU3InvariantPointSelected, a.GeneralTwoPlusOneSelector, a.NativeRankOnePU, a.MinimalSeal, a.PhysicalWeakFlavorEWDerived, a.K7OrProductTimeOpened, a.MissingNextTheorem, a.Verdict)
}
