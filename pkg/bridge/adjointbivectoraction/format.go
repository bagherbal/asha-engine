package adjointbivectoraction

import "fmt"

func FormatInherited(a InheritedGate249Audit) string {
	return fmt.Sprintf("8v=%t strategy=%t EW8v=%t kernel=%t kernel3=%t scalarIso=%t vtau=%t triality=%t yukawa=%t truth=%q", a.Carrier8VKnown, a.NeutralKernelStrategy, a.EWDerivationActionDerived, a.NeutralKernelDerived, a.NeutralKernelDim3, a.ScalarNeutralIsomorphism, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.TruthStatement)
}

func FormatCarrier(a CliffordCarrierAudit) string {
	return fmt.Sprintf("sig=%s dim=%d metric=%v basis=%v grade1=%t formula=%q typed=%t verdict=%s", a.Signature, a.VectorDimension, a.MetricDiagonal, a.VectorBasis, a.Grade1CarrierKnown, a.BivectorActionFormula, a.CommutatorActionTyped, a.Verdict)
}

func FormatSimpleBlade(a SimpleBivectorMatrixAudit) string {
	return fmt.Sprintf("blade=%s size=%dx%d skew=%t rank=%d kernel=%d evenKernel=%t entries=%v verdict=%s", a.Blade, a.MatrixRows, a.MatrixCols, a.SkewSymmetric, a.Rank, a.KernelDimension, a.KernelDimensionEven, a.NonzeroEntries, a.Verdict)
}

func FormatKernelParity(a GenericBivectorParityAudit) string {
	return fmt.Sprintf("skew=%t evenRank=%t evenKernel8=%t exact3=%t reason=%q verdict=%s", a.RealBivectorAdjointMatricesSkew, a.SkewRankAlwaysEven, a.Dimension8KernelAlwaysEven, a.Exact3DKernelPossible, a.Reason, a.Verdict)
}

func FormatEWBivectors(a EWBivectorRetrievalAudit) string {
	return fmt.Sprintf("requested=%v scalarT3=%t scalarY=%t T3blade=%t Yblade=%t T3label=%q Ylabel=%q Qblade=%t Zblade=%t manual=%q rejected=%t obstruction=%q verdict=%s", a.RequestedGenerators, a.ScalarT3Available, a.ScalarYPhiAvailable, a.T3Grade2BladeDerived, a.YPhiGrade2BladeDerived, a.T3BladeLabel, a.YPhiBladeLabel, a.QBladeDerived, a.ZBladeDerived, a.ManualBladeAssignment, a.ManualAssignmentRejected, a.Obstruction, a.Verdict)
}

func FormatMatrices(a ExplicitMatrixAudit) string {
	return fmt.Sprintf("formula=%q candidate=%t T3=%t Y=%t Q8v=%t Z8v=%t Qsize=%dx%d QskewIfBivector=%t kernelKnown=%t kernel=%d neutral3=%t failure=%q verdict=%s", a.AdjointActionFormula, a.CanConstructCandidateBlade, a.CanConstructT3Matrix, a.CanConstructYPhiMatrix, a.Q8VConstructed, a.Z8VConstructed, a.Q8VRows, a.Q8VCols, a.Q8VSkewIfBivector, a.Q8VKernelDimensionKnown, a.Q8VKernelDimension, a.Neutral3PlaneDerived, a.BindingFailure, a.Verdict)
}

func FormatScalarPlane(a ScalarNeutralPlaneAudit) string {
	return fmt.Sprintf("slots=%v tau=%v needsQ=%t Qkernel=%t kernel3=%t iso=%t vtau=%t reason=%q verdict=%s", a.NeutralTraceSlots, a.TauEta, a.NeedsQ8VKernel, a.Q8VKernelAvailable, a.KernelDimensionExactly3, a.CanonicalIsomorphism, a.VTauConstructed, a.Reason, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("inventT3=%t inventY=%t assignCharges=%t forceKernel3=%t vtauHand=%t triality=%t yukawa=%t CKM=%t polluted=%t verdict=%s", a.InventedT3Blade, a.InventedYPhiBlade, a.AssignedChargesToGammaBasis, a.ForcedKernelDim3, a.ConstructedVTauByHand, a.InventedTrialityMap, a.InsertedYukawaTexture, a.ClaimedCKMPMNS, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("adjoint=%t candidates=%t ewBlades=%t Q8v=%t kernel=%t kernel3=%t realBivector3=%t vtau=%t triality=%t yukawa=%t status=%q next=%q comment=%q", a.CliffordAdjointAvailable, a.CandidateMatricesComputable, a.EWBivectorsRetrieved, a.Q8VMatrixDerived, a.NeutralKernelDerived, a.NeutralKernelDim3, a.RealBivector3KernelPossible, a.VTauConstructed, a.TrialityUnblocked, a.YukawaTextureDerived, a.Status, a.NextGate, a.Comment)
}
