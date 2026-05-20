package generation2etarecordtransferranktraceobstructionaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedGate558Audit) string {
	return fmt.Sprintf("algebra=%t name=%q carrier=%q dim=%d ranks=%d+%d split=%q tauTraceOnly=%t noFunctor=%t verdict=%q", a.AlgebraConstructed, a.AlgebraName, a.SourceCarrier, a.SourceDimension, a.SourcePlusRank, a.SourceMinusRank, a.SourceSplit, a.TauEtaTraceValuesOnly, a.NoPreviousTransferFunctor, a.Verdict)
}

func FormatRankSplit(a RankSplit) string {
	return fmt.Sprintf("%s(r+=%d,r-=%d,formal2plus1=%t,canonical=%t)", a.Name, a.PlusRank, a.MinusRank, a.ProducesTwoPlusOne, a.CanonicalInASHA)
}

func FormatRankSplitList(xs []RankSplit) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatRankSplit(x)
	}
	return strings.Join(parts, ", ")
}

func FormatFormalReps(a FormalRepresentationClassification) string {
	return fmt.Sprintf("algebra=%q carrier=%q dim=%d exist=%t idem=%t exhaustive=%t formal2plus1=%t canonical2plus1=%t splits=[%s] verdict=%q", a.Algebra, a.Carrier, a.CarrierDimension, a.UnitalRepresentationsExist, a.EquivalentToComplementaryIdempotents, a.Exhaustive, a.AnyTwoPlusOneFormal, a.AnyCanonicalTwoPlusOne, FormatRankSplitList(a.RankSplits), a.Verdict)
}

func FormatCanonicalChoice(a CanonicalChoiceAudit) string {
	return fmt.Sprintf("W=%t basis=%q genCapacity=%t native2plus1=%t U12=%t genLabels=%t verdict=%q", a.WSpatialCarrierAvailable, a.WSpatialBasisName, a.GenerationCarrierCapacityVisible, a.NativeBasisIndependentReasonFor2Plus1, a.NativeReasonForU12, a.NativeReasonForGenerationLabels, a.Verdict)
}

func FormatTraceRank(a TraceRankPreservationAudit) string {
	return fmt.Sprintf("source=%q ranks=%v targetDim=%d ordinaryTracePossible=%t requiredTargetRanks=%v obstruction=%q verdict=%q", a.SourceCarrier, a.SourceRanks, a.TargetCarrierDimension, a.OrdinaryTracePreservingPossible, a.RequiredTargetRanks, a.Obstruction, a.Verdict)
}

func FormatNormalizedTrace(a NormalizedTracePreservationAudit) string {
	return fmt.Sprintf("sourceNorm=%v targetDim=%d requiredTargetRanks=%v integral=%t obstruction=%q verdict=%q", a.SourceNormalizedTraces, a.TargetDimension, a.RequiredTargetRanks, a.IntegralRanksPossible, a.Obstruction, a.Verdict)
}

func FormatBL(a BLCompatibilityAudit) string {
	return fmt.Sprintf("restricted=%q formalCommutes=%t rankSplit=%t labels=%t U12=%t verdict=%q", a.RestrictedBLOnWSpatial, a.AnyFormalTransferCommutesWithBL, a.BLSuppliesRankSplit, a.BLSuppliesBasisLabels, a.BLSuppliesCanonicalU12, a.Verdict)
}

func FormatSpectralTriple(a SpectralTripleCompatibilityAudit) string {
	return fmt.Sprintf("candidate=%t gamma=%t J=%t D=%t firstOrder=%t passed=%t missing=%v verdict=%q", a.CandidateTransferExists, a.GradingCheckAvailable, a.JCheckAvailable, a.DCheckAvailable, a.FirstOrderCheckAvailable, a.CompatibilityPassed, a.MissingData, a.Verdict)
}

func FormatGeneration(a GenerationCarrierAudit) string {
	return fmt.Sprintf("formalDim3=%t nativeLabels=%t functor=%t unit=%t hierarchy=%t yukawaCKM=%t verdict=%q", a.FormalDim3GenerationCapacityVisible, a.NativeBasisIndependentLabels, a.FunctorFromAetaRec, a.UnitPreservationVerified, a.ProducesGenerationHierarchy, a.ProducesYukawaOrCKMPMNS, a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("weakPlane=%t weakIso=%t higgs=%t hierarchy=%t yukawa=%t ckm=%t observed=%t preserved=%t verdict=%q", a.WeakPlaneSelectionClaimed, a.WeakIsospinIdentificationClaimed, a.HiggsRadialGoldstoneClaimed, a.GenerationHierarchyClaimed, a.YukawaTextureClaimed, a.CKMPMNSClaimed, a.ObservedFlavorImported, a.Preserved, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("formalReps=%t canonical=%t traceRank=%t BLCanonical=%t lawful=%t next=%q verdict=%q", a.FormalRepresentationsExist, a.CanonicalInASHA, a.TraceRankPreservingTransfer, a.BMinusLCanonicalizesTransfer, a.LawfulTransferAvailable, a.MissingNextTheorem, a.Verdict)
}
