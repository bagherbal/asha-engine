package spin8trialityfunctor

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

func FormatInherited(a InheritedGate246Audit) string {
	return fmt.Sprintf("origin=%t functor=%t tauCapacity=%t texture=%t rawNoncommuting=%t qualified=%t CKMPMNS=%t masses=%t truth=%q", a.ScalarOriginKnown, a.ScalarToTrialityFunctorDerived, a.TauGenerationCapacity, a.GenerationTextureDerived, a.RawNonCommutingCapacity, a.QualifiedTexturePairDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.TruthStatement)
}

func FormatSpin8(a Spin8TrialityAudit) string {
	return fmt.Sprintf("triality=%t reps=%v group=%q vectorToSpinor=%t matrices=%t scalarBundleIsVector=%t traceTripleIsVector=%t verdict=%s", a.AbstractSpin8TrialityAvailable, a.TrialityRepresentations, a.AutomorphismGroup, a.VectorToSpinorFunctorKnown, a.ExplicitMatricesOnSC, a.ScalarBundleIsVectorRep, a.ScalarTraceTripleIsVector, a.Verdict)
}

func FormatScalarSpinor(a ScalarToSpinorFunctorAudit) string {
	return fmt.Sprintf("source=%q carrier=%q required=%q target=%q traceDim=%d genDim=%d dimMatch=%t exteriorRep=%t charRep=%t explicitTriality=%t pullback=%t manual=%q rejected=%t verdict=%s", a.SourceObject, a.SourceCarrier, a.RequiredSourceForTriality, a.TargetCarrier, a.DimensionOfTraceTriple, a.GenerationCarrierDimension, a.DimensionMatch, a.ExteriorOrVectorRepresentativeKnown, a.CharacteristicRepresentativeKnown, a.ExplicitTrialityAutomorphismKnown, a.PullbackFunctorDerived, a.ManualPullback, a.ManualPullbackRejected, a.Verdict)
}

func FormatTexture(a TextureRealizationAudit) string {
	return fmt.Sprintf("candidate=%q eig=%v distinct=%d breaks=%t commCycle=%s normC=%.6g commRefl=%s normR=%.6g raw=%t pullback=%t operator=%t yukawa=%t CKM=%t PMNS=%t verdict=%s", a.CandidateName, a.CandidateEigenvalues, a.DistinctEigenvalues, a.BreaksGenerationDegeneracy, FormatMatrix3(a.CommutatorWithCycle), a.CycleCommutatorNorm, FormatMatrix3(a.CommutatorWithReflection), a.ReflectionCommutatorNorm, a.RawNonCommutingCapacity, a.LawfulPullbackDerived, a.DiagonalOperatorConstructed, a.YukawaTextureDerived, a.CKMDerived, a.PMNSDerived, a.Verdict)
}

func FormatObstruction(a PullbackObstructionAudit) string {
	return fmt.Sprintf("missing=%v typeMismatch=%q why=%q pullback=%t level=%q verdict=%s", a.MissingPieces, a.BindingTypeMismatch, a.WhyTrialityInsufficient, a.PullbackDerived, a.ObstructionLevel, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("connes=%t spin8Matrices=%t forcedMap=%t insertedDTau=%t yukawaMasses=%t CKM=%t PMNS=%t claimedMasses=%t claimedFlavor=%t polluted=%t verdict=%s", a.ImportedConnesAlgebra, a.InventedSpin8Matrices, a.ForcedScalarToSpinorMap, a.InsertedDTauAsTexture, a.ImportedYukawaMasses, a.ImportedCKM, a.ImportedPMNS, a.ClaimedFermionMasses, a.ClaimedFiniteFlavorTheorem, a.PollutedFiniteCore, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("triality=%t dimMatch=%t tauCapacity=%t scalarVector=%t functor=%t Dtau=%t texture=%t CKMPMNS=%t masses=%t status=%q next=%q comment=%q", a.Spin8TrialityAvailable, a.DimensionMatch, a.TauTextureCapacityInherited, a.ScalarTraceIsVectorRep, a.TrialityFunctorDerived, a.DiagonalTextureConstructed, a.QualifiedTextureDerived, a.CKMPMNSDerived, a.FermionMassesDerived, a.Status, a.NextGate, a.Comment)
}
