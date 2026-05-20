package generation2booleanoctonionicsupportactivationminimalityaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate685Inheritance) string {
	return fmt.Sprintf("rankTraceInherited=%t booleanOctonionicSelection=%t selected=%q dBase=%.15g sSplit=%.15g h72=%d k7=%d trace=%q selection=%q activationUnproved=%t priorFirewall=%t verdict=%q", x.RankSevenTraceInherited, x.BooleanOctonionicSelection, x.SelectedProjector, x.DBase, x.SSplit, x.H72Dimension, x.K7Dimension, x.TraceScalarization, x.ProjectorSelection, x.ActivationStillUnproved, x.PriorFirewallPreserved, x.Verdict)
}

func FormatLadderStep(x ConstraintStep) string {
	return fmt.Sprintf("%d:%s constraints=[%s] carrier=%q dim=%d rank7=%t degenerate=%t uniquePK7=%t witness=%q verdict=%q", x.Index, x.Name, strings.Join(x.Constraints, ", "), x.Carrier, x.CarrierDimension, x.RankSevenRequired, x.Degenerate, x.UniquePK7, x.Witness, x.Verdict)
}

func FormatLadder(x ConstraintLadderAudit) string {
	steps := make([]string, 0, len(x.Steps))
	for _, s := range x.Steps {
		steps = append(steps, FormatLadderStep(s))
	}
	return fmt.Sprintf("rankOnlyDegenerate=%t finiteOnlyDegenerate=%t booleanOnlyDegenerate=%t octonionicOnlyDegenerate=%t combinedSelectsK7=%t weakerDegenerate=%t minimalPair=%t stepVerdicts=[%s] steps=[%s] verdict=%q", x.RankOnlyDegenerate, x.FiniteSupportOnlyDegenerate, x.BooleanOnlyDegenerate, x.OctonionicOnlyDegenerate, x.CombinedSupportSelectsK7, x.AllWeakerSelectorsDegenerate, x.MinimalPairRequired, strings.Join(sortedStepVerdicts(x.Steps), ","), strings.Join(steps, " | "), x.Verdict)
}

func FormatIndependence(x IndependenceAudit) string {
	return fmt.Sprintf("booleanComplement=%d octonionicComplement=%d booleanWitness=%q octonionicWitness=%q booleanImpliesOctonionic=%t octonionicImpliesBoolean=%t neitherRedundant=%t bothRequired=%t verdict=%q", x.BooleanComplementDimension, x.OctonionicComplementDimension, x.BooleanOnlyWitness, x.OctonionicOnlyWitness, x.BooleanImpliesOctonionic, x.OctonionicImpliesBoolean, x.NeitherConditionRedundant, x.BothRequiredToForceK7, x.Verdict)
}

func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("assumptions=[%s] doesNotAssumePK7=%t rankSupportOnly=%t intersectionDimOnly=%t conclusion=%q conditional=%t noncircular=%t verdict=%q", strings.Join(x.Assumptions, "; "), x.DoesNotAssumePK7, x.UsesOnlyRankAndSupport, x.UsesOnlyIntersectionDim, x.ConclusionDerived, x.ConditionalNotAbsolute, x.Noncircular, x.Verdict)
}

func FormatDecomposition(x ActivationDecompositionAudit) string {
	return fmt.Sprintf("response=%q scalar=%q selector=%q trace=%q scalarSelectsRank=%t supportSelectsPK7=%t traceOnly=%t sSplitAloneSelectsProjector=%t activationProved=%t verdict=%q", x.ActiveResponse, x.BoundaryControlScalar, x.ProjectorIdentitySelector, x.TraceScalarization, x.BoundaryScalarSelectsRank, x.SupportSelectorSelectsPK7, x.TraceOnlyScalarizes, x.SSplitAloneSelectsProjector, x.NativeActivationProved, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("claimsSSplitSelectsProjector=%t claimsBoundaryScalarActivatesSieve=%t claimsProjectorActivation=%t claims7=%t claimsBoundary=%t claimsScalarRG=%t claimsHiggs=%t claimsGauge=%t claimsFlavor=%t verdict=%q", x.ClaimsSSplitSelectsProjector, x.ClaimsBoundaryScalarActivatesSieve, x.ClaimsProjectorActivation, x.ClaimsNativeSevenOver72, x.ClaimsBoundaryStressDerivation, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.Verdict)
}
