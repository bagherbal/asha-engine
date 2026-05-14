package complexweightspacekernel

import "fmt"

func FormatInherited(a InheritedGate250Audit) string {
	return fmt.Sprintf("8v=%t adjoint=%t candidates=%t ewBlades=%t Q8v=%t kernel=%t kernel3=%t real3=%t vtau=%t triality=%t yukawa=%t truth=%q", a.Carrier8VKnown, a.CliffordAdjointAvailable, a.CandidateMatricesComputable, a.EWBivectorsRetrieved, a.Q8VMatrixDerived, a.NeutralKernelDerived, a.NeutralKernelDim3, a.RealBivector3KernelPossible, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.TruthStatement)
}

func FormatComplexification(a ComplexificationAudit) string {
	return fmt.Sprintf("real=%s complex=%s dimR=%d dimC=%d underlyingR=%d native=%t liftEvenRank=%t verdict=%s", a.RealCarrierName, a.ComplexCarrierName, a.RealDimension, a.ComplexDimension, a.UnderlyingRealDimension, a.ComplexificationNative, a.EvenRankObstructionLift, a.Verdict)
}

func FormatHermitian(a HermitianPreflightAudit) string {
	return fmt.Sprintf("realSkew=%t conversion=%q realSpectrum=%t oddWeights=%t candidate=%q candidateKernelC=%d Qherm=%t Zherm=%t verdict=%s", a.RealSkewGeneratorAvailable, a.HermitianConversion, a.HermitianOperatorsHaveRealSpectrum, a.OddWeightSpacesAllowed, a.CandidateSimpleBlade, a.CandidateSimpleBladeKernelComplexDim, a.PhysicalQHermitianDerived, a.PhysicalZHermitianDerived, a.Verdict)
}

func FormatCartan(a CartanWeightAudit) string {
	return fmt.Sprintf("required=%v cartan=%t Q8vC=%t Z8vC=%t simultaneous=%t weights=%t manualRejected=%t obstruction=%q verdict=%s", a.RequiredOperators, a.CartanCommutingPairDerived, a.Q8vCMatrixDerived, a.Z8vCMatrixDerived, a.SimultaneouslyDiagonal, a.WeightSpectrumDerived, a.ManualChargeAssignmentRejected, a.Obstruction, a.Verdict)
}

func FormatNeutralKernel(a ComplexNeutralKernelAudit) string {
	return fmt.Sprintf("def=%q computed=%t dimKnown=%t dim=%d exact3=%t oddAllowed=%t missingQ=%t failure=%q verdict=%s", a.Definition, a.Computed, a.DimensionKnown, a.Dimension, a.ExactlyThreeComplexDim, a.OddDimAllowedInPrinciple, a.DependsOnMissingQ8vC, a.BindingFailure, a.Verdict)
}

func FormatTriality(a ComplexTrialityAudit) string {
	return fmt.Sprintf("spin8C=%t modules=%v sameDim=%t outerAuto=%t canonicalUntwisted=%t kernel=%t map=%t Jchecked=%t Jcompatible=%t obstruction=%q verdict=%s", a.Spin8TrialityOverC, a.Modules, a.SameComplexDimension, a.OuterAutomorphismRequired, a.CanonicalUntwistedIsomorphism, a.NeutralKernelAvailable, a.MapNeutralKernelToSpinor, a.RealStructureCompatibilityChecked, a.CompatibleWithJ, a.Obstruction, a.Verdict)
}

func FormatVTau(a VTauAudit) string {
	return fmt.Sprintf("tau=%v needs3plane=%t plane=%t needsFrame=%t frame=%t constructed=%t feedTriality=%t rejected=%q verdict=%s", a.TauEta, a.NeedsNeutral3Plane, a.Neutral3PlaneAvailable, a.NeedsScalarSlotFrame, a.ScalarSlotFrameDerived, a.Constructed, a.WouldFeedTriality, a.RejectedBecause, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("inventQ=%t inventZ=%t weights=%t force3=%t triality=%t ignoreJ=%t vtau=%t yukawa=%t CKM=%t polluted=%t verdict=%s", a.InventedQ8vC, a.InventedZ8vC, a.AssignedComplexWeightsByHand, a.ForcedKernelDim3, a.InventedTrialityIsomorphism, a.IgnoredRealStructure, a.ConstructedVTauByHand, a.InsertedYukawaTexture, a.ClaimedCKMPMNS, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("complex8v=%t hermitian=%t oddKernel=%t hermMatrices=%t kernel=%t kernel3=%t trialityArena=%t canonicalTriality=%t J=%t vtau=%t triality=%t yukawa=%t status=%q next=%q comment=%q", a.Complex8VKnown, a.HermitianWeightCapacity, a.OddComplexKernelCapacity, a.NativeHermitianMatrices, a.ComplexNeutralKernelDerived, a.NeutralKernelDim3, a.ComplexTrialityArena, a.CanonicalTrialityMap, a.RealStructureCompatible, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
