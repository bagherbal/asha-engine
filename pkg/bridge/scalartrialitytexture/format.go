package scalartrialitytexture

import (
	"fmt"
	"strings"
)

func FormatMatrix3(m Matrix3) string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		rows = append(rows, fmt.Sprintf("[%.6g %.6g %.6g]", m[i][0], m[i][1], m[i][2]))
	}
	return strings.Join(rows, " ")
}

func FormatScalarFlavor(a ScalarFlavorAlignmentAudit) string {
	return fmt.Sprintf("seq=%v source=%q operators=%v higgsSector=%t origin=%t functor=%t genDim=%d trialityEarlier=%t typeCorrect=%t mapDerived=%t verdict=%s", a.TauSequence, a.SourceBundle, a.SourceOperators, a.NativeHiggsSectorObservable, a.ScalarOriginKnown, a.ScalarToTrialityFunctorDerived, a.TrialityCarrierDimension, a.TrialityCarrierDerivedEarlier, a.MapWouldBeTypeCorrect, a.MapActuallyDerived, a.Verdict)
}

func FormatGenerationTexture(a GenerationTextureAudit) string {
	return fmt.Sprintf("name=%q eig=%v distinct=%d breaksS3=%t trace=%d det=%d frob2=%d hermitian=%t scalarMap=%t operator=%t texture=%t conditional=%q verdict=%s", a.CandidateOperatorName, a.Eigenvalues, a.DistinctEigenvalues, a.BreaksS3Degeneracy, a.Trace, a.Determinant, a.FrobeniusNormSquared, a.HermitianDiagonalCapacity, a.ScalarToGenerationMap, a.GenerationOperatorDerived, a.YukawaTextureDerived, a.TextureIfMapExisted, a.Verdict)
}

func FormatNonCommuting(a NonCommutingTextureAudit) string {
	return fmt.Sprintf("gate173Needs=%t gate173Qualified=%t D=%q cycle=%q refl=%q commCycle=%s normC=%.6g commRefl=%s normR=%.6g raw=%t wouldQualify=%t qualified=%t reason=%q CKMcap=%t PMNScap=%t CKM=%t PMNS=%t verdict=%s", a.Gate173NeedsNonCommutingPair, a.Gate173FoundQualifiedPair, a.TauDiagonalName, a.TrialityCycleName, a.TrialityReflectionName, FormatMatrix3(a.CommutatorWithCycle), a.CycleCommutatorNorm, FormatMatrix3(a.CommutatorWithReflection), a.ReflectionCommutatorNorm, a.RawNonCommutingWithTriality, a.PairWouldBeQualifiedIfPullbackHeld, a.PairActuallyQualified, a.ReasonNotQualified, a.CKMPrerequisiteCapacity, a.PMNSPrerequisiteCapacity, a.CKMDerived, a.PMNSDerived, a.Verdict)
}

func FormatPullback(a PullbackObstructionAudit) string {
	return fmt.Sprintf("scalarCarrier=%q trialityCarrier=%q shared=%v missingFunctor=%q missingCompat=%v manual=%q rejected=%t derived=%t reason=%q verdict=%s", a.ScalarBundleCarrier, a.TrialityCarrier, a.KnownSharedStructure, a.MissingFunctor, a.MissingCompatibility, a.ManualDiagonalInsertion, a.ManualInsertionRejected, a.PullbackDerived, a.Reason, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("forcedScalarGen=%t forcedTauTexture=%t yukawaMasses=%t CKM=%t PMNS=%t observedMasses=%t claimedMasses=%t claimedFlavor=%t weakPlane=%t polluted=%t verdict=%s", a.ForcedScalarToGenerationMap, a.ForcedTauDiagonalTexture, a.ImportedYukawaMasses, a.ImportedCKM, a.ImportedPMNS, a.InsertedObservedMasses, a.ClaimedFermionMasses, a.ClaimedFiniteFlavorTheorem, a.ClaimedWeakPlane, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("origin=%t functor=%t tauCapacity=%t texture=%t noncommutingCapacity=%t qualifiedPair=%t CKMPMNS=%t masses=%t weak=%t status=%q next=%q comment=%q", a.ScalarOriginKnown, a.ScalarToTrialityFunctorDerived, a.TauGenerationCapacity, a.GenerationTextureDerived, a.RawNonCommutingCapacity, a.QualifiedTexturePairDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.WeakPlaneDerived, a.Status, a.NextGate, a.Comment)
}
