package lietrialitypullback

import "fmt"

func FormatInherited(a InheritedGate251Audit) string {
	return fmt.Sprintf("complex8v=%t hermitian=%t oddKernel=%t nativeHerm=%t kernel=%t kernel3=%t trialityArena=%t canonicalTriality=%t J=%t vtau=%t triality=%t yukawa=%t truth=%q", a.Complex8VKnown, a.HermitianWeightCapacity, a.OddComplexKernelCapacity, a.NativeHermitianMatrices, a.ComplexNeutralKernelDerived, a.NeutralKernelDim3, a.ComplexTrialityArena, a.CanonicalTrialityMap, a.RealStructureCompatible, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.TruthStatement)
}

func FormatInfinitesimalTriality(a InfinitesimalTrialityAudit) string {
	return fmt.Sprintf("lie=%q dim=%d outer=%q reps=%t lieAuto=%t needsExplicit=%t explicit=%t canonical=%t verdict=%q", a.Spin8LieAlgebra, a.LieAlgebraDimension, a.TrialityOuterAutomorphism, a.CanPermuteRepresentations, a.ActsOnLieAlgebra, a.RequiresExplicitAutomorphism, a.ExplicitAutomorphismDerived, a.CanonicalWithoutChoice, a.Verdict)
}

func FormatSpinorGenerators(a SpinorGeneratorAudit) string {
	return fmt.Sprintf("required=%v bridge=%t fock=%t scalar=%t so8coords=%t skewHerm=%t suitable=%t obstruction=%q verdict=%q", a.RequiredGenerators, a.BridgeRepresentationsKnown, a.SpinorFockActionKnown, a.ScalarBundleActionKnown, a.AsSO8BivectorCoordinates, a.AsSkewHermitianSpin8Generators, a.SuitableForInfinitesimalTriality, a.Obstruction, a.Verdict)
}

func FormatTranslation(a TranslationAudit) string {
	return fmt.Sprintf("input=%t map=%t pushT3=%t pushY=%t T3v=%t Yv=%t manualRejected=%t obstruction=%q verdict=%q", a.InputSpinorGeneratorsAvailable, a.InfinitesimalTrialityMapKnown, a.CanPushT3To8V, a.CanPushYTo8V, a.T3VectorMatrixDerived, a.YVectorMatrixDerived, a.ManualDictionaryRejected, a.Obstruction, a.Verdict)
}

func FormatHermitianQ(a HermitianQAudit) string {
	return fmt.Sprintf("carrier=%q rule=%q T3v=%t Yv=%t HT3=%t HY=%t Q=%t Z=%t matrices=%t verdict=%q", a.ComplexCarrier, a.HermitianRule, a.T3VectorMatrixDerived, a.YVectorMatrixDerived, a.HT3Constructed, a.HYConstructed, a.Q8vCConstructed, a.Z8vCConstructed, a.HermitianMatricesAvailable, a.Verdict)
}

func FormatNeutralKernel(a NeutralKernelAudit) string {
	return fmt.Sprintf("def=%q Q=%t eig=%t dimKnown=%t dim=%d exact3=%t plane=%t missingQ=%t verdict=%q", a.Definition, a.Q8vCConstructed, a.EigensystemComputed, a.KernelDimensionKnown, a.KernelComplexDimension, a.ExactlyThree, a.ThreePlaneDerived, a.DependsOnMissingQ, a.Verdict)
}

func FormatTransport(a TrialityTransportAudit) string {
	return fmt.Sprintf("arena=%t plane=%t map=%t image=%t Jspinor=%t Jvector=%t commutesJ=%t meaningful=%t obstruction=%q verdict=%q", a.ComplexTrialityArenaKnown, a.Neutral3PlaneAvailable, a.Canonical8vCTo8sCMapDerived, a.NeutralPlaneImageInSpinorKnown, a.RealStructureJKnownOnSpinor, a.RealStructureJKnownOnVector, a.CommutesWithJ, a.TransportPhysicallyMeaningful, a.Obstruction, a.Verdict)
}

func FormatVTau(a VTauAudit) string {
	return fmt.Sprintf("tau=%v needsPlane=%t plane=%t needsFrame=%t frame=%t constructed=%t transport=%t yukawa=%t rejected=%q verdict=%q", a.TauEta, a.NeedsNeutral3Plane, a.Neutral3PlaneAvailable, a.NeedsScalarSlotFrame, a.ScalarSlotFrameDerived, a.Constructed, a.TrialityTransportReady, a.YukawaTextureDerived, a.RejectedBecause, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("so8=%t triality=%t T3v=%t Yv=%t Q=%t force3=%t ignoreJ=%t vtau=%t yukawa=%t CKM=%t polluted=%t verdict=%q", a.InventedSO8Coordinates, a.InventedLieTrialityMap, a.InventedT3VectorMatrix, a.InventedYVectorMatrix, a.InventedQ8vC, a.ForcedKernelDim3, a.IgnoredJCompatibility, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.ClaimedCKMPMNS, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("infTriality=%t spinorEW=%t spinorSO8=%t trialityMap=%t vectorEW=%t Q=%t plane=%t Jtransport=%t vtau=%t triality=%t yukawa=%t status=%q next=%q comment=%q", a.InfinitesimalTrialityCapacity, a.SpinorEWBridgeKnown, a.SpinorSO8Coordinates, a.ExplicitLieTrialityMap, a.VectorEWMatriciesDerived, a.Q8vCConstructed, a.Neutral3PlaneDerived, a.JCompatibleTransport, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
