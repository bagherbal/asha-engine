package contactsingletonflavorfunctor

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Inheritance) string {
	return fmt.Sprintf("executed=%t contactSingletons=%t promotable=%d rationalSingletons=%d quarticPrimary=%d rowSemantics=%d contactActionBlocked=%t oneFormEdges=%t chargedModuli=%d noEmpirical=%t (%s)", a.Executed, a.Gate396ContactSingletonsFound, a.Gate396PromotableGenerationCount, a.Gate151RationalSingletons, a.Gate151QuarticPrimaryBlocks, a.Gate151RowSemantics, a.Gate184ContactActionBlocked, a.Gate385OneFormEdgeSupportDerived, a.Gate372ChargedModuliDim, a.NoEmpiricalFlavorValuesImported, a.Verdict)
}

func FormatSingletonBlock(b SingletonBlock) string {
	return fmt.Sprintf("%s eigen=%s dim=%d field=%s exact=%t native=%t row=%t generation=%t (%s)", b.Name, b.Eigenvalue, b.Dimension, b.BaseField, b.ProjectorExact, b.ProjectorNative, b.RowSemantic, b.GenerationSemantic, b.Verdict)
}

func FormatSingletons(a SingletonAlgebraAudit) string {
	rows := make([]string, 0, len(a.Blocks)+1)
	rows = append(rows, fmt.Sprintf("algebra=%s dim=%d exactIdempotents=%d nativeDomain=%t contactAction=%t finiteDiracAction=%t generationSemantics=%t spectrum=%v (%s)", a.Algebra, a.Dimension, a.ExactOrthogonalIdempotents, a.NativeDomainAlgebra, a.ActsOnContactDomain, a.ActsOnFiniteDiracTarget, a.NativeGenerationSemantics, a.Spectrum, a.Verdict))
	for _, b := range a.Blocks {
		rows = append(rows, "  - "+FormatSingletonBlock(b))
	}
	return strings.Join(rows, "\n")
}

func FormatTarget(a FiniteDiracTargetAudit) string {
	return fmt.Sprintf("target=%q algebra=%q J=%t firstOrder=%t EW=%t oneForm=%t JEdges=%d yukawaChannels=%d trivialGen=%t nativeGenOpDim=%d edgeRank=%d uniform=%t (%s)", a.Target, a.FiniteAlgebra, a.JCompatibleRequired, a.FirstOrderRequired, a.HyperchargeSU2Required, a.OneFormEdgeSupportDerived, a.YSymmetrizedEdgeCount, a.MinimalYukawaChannels, a.GenerationsCurrentlyTrivial, a.NativeGenerationOperatorDim, a.EdgeGenerationRank, a.EdgePatternUniform, a.Verdict)
}

func FormatFunctor(c FunctorCandidate) string {
	return fmt.Sprintf("%s domain=%q target=%q native=%t sealed=%t circular=%t contact=%t edge=%t AF=%t J=%t firstOrder=%t EW=%t oneForms=%t central=%t noncentral=%t diagonal=%t mixing=%t choices=%d rank=%d spectrum=%v commutant=%d promotable=%t reason=%q (%s)", c.Name, c.Domain, c.Target, c.Native, c.Sealed, c.Circular, c.DerivedFromContactIdempotents, c.DerivedFromFiniteDiracEdges, c.CompatibleWithAF, c.CompatibleWithJ, c.CompatibleWithFirstOrder, c.CompatibleWithHyperchargeSU2, c.CompatibleWithOneForms, c.CentralOnGeneration, c.NoncentralOnGeneration, c.DiagonalOnly, c.MixingCapacity, c.AssignmentChoices, c.Rank, c.Spectrum, c.CommutantDimension, c.PromotableAsNativeFunctor, c.Reason, c.Verdict)
}

func FormatFunctors(a FunctorAudit) string {
	rows := []string{fmt.Sprintf("executed=%t nativeCandidates=%d nativeActionFunctors=%d nativeNoncentral=%d sealedNoncentral=%d promotable=%d best=%q (%s)", a.Executed, a.NativeCandidateCount, a.NativeActionFunctorCount, a.NativeNoncentralCount, a.SealedNoncentralCount, a.PromotableNativeCount, a.BestNativeCandidate, a.Verdict)}
	for _, c := range a.Candidates {
		rows = append(rows, "  - "+FormatFunctor(c))
	}
	return strings.Join(rows, "\n")
}

func FormatPair(p PairAudit) string {
	return fmt.Sprintf("%s left=%q right=%q native=%t sealed=%t eligible=%t norm=%.12g noncommuting=%t ckm=%t reason=%q (%s)", p.Name, p.Left, p.Right, p.NativePair, p.SealedPair, p.Eligible, p.CommutatorNorm, p.Noncommuting, p.CKMCapacity, p.Reason, p.Verdict)
}

func FormatOperators(a OperatorAudit) string {
	rows := []string{fmt.Sprintf("executed=%t nativeEligible=%d nativeNoncentral=%d nativePairs=%d sealedPairs=%d maxNative=%.12g maxSealed=%.12g nativeCKM=%t (%s)", a.Executed, a.NativeEligibleOperators, a.NativeNoncentralOperators, a.NativeNoncommutingPairs, a.SealedNoncommutingPairs, a.MaxNativeCommutatorNorm, a.MaxSealedCommutatorNorm, a.CKMCapacityNative, a.Verdict)}
	for _, p := range a.Pairs {
		rows = append(rows, "  - "+FormatPair(p))
	}
	return strings.Join(rows, "\n")
}

func FormatScenario(s ModuliScenario) string {
	return fmt.Sprintf("%s assumption=%q start=%d result=%d native=%t conditional=%t failed=%t masses=%t ckm=%t reason=%q (%s)", s.Name, s.AssumptionClass, s.StartingChargedDim, s.ResultingDim, s.Native, s.Conditional, s.Failed, s.ThreeDistinctMassesPossible, s.CKMMisalignmentPossible, s.Reason, s.Verdict)
}

func FormatModuli(a ModuliAudit) string {
	rows := []string{fmt.Sprintf("executed=%t start=%d nativeReduction=%t bestNative=%d bestConditional=%d (%s)", a.Executed, a.StartingChargedDim, a.NativeReductionBelow13, a.BestNativeDim, a.BestConditionalDim, a.Verdict)}
	for _, s := range a.Scenarios {
		rows = append(rows, "  - "+FormatScenario(s))
	}
	return strings.Join(rows, "\n")
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("executed=%t masses=%t ckm=%t pmns=%t ordering=%t manualGen=%t rootsPromoted=%t sealedCycle=%t nativeFlavor=%t moduliClaim=%t (%s)", a.Executed, a.NoMassesImported, a.NoCKMImported, a.NoPMNSImported, a.NoObservedOrderingImported, a.NoManualGenerationAssignment, a.NoContactRootsPromoted, a.NoSealedCyclePromoted, a.NoNativeFlavorClaimed, a.NoModuliReductionClaimed, a.Verdict)
}

func FormatNext(a NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", a.Gate, a.Title, a.Reason, a.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 397 Registry Audit — Contact Rational Singleton to Finite-Dirac Flavor Functor Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can the three exact rational contact singleton idempotent blocks act canonically as finite-Dirac generation labels? Equivalently, does the current ASHA ledger derive an explicit representation `rho: Q^3_contact -> End(H_finite-Dirac)` compatible with `A_F`, `J`, first-order, electroweak charges, and inner-fluctuation one-form support?\n\n")
	b.WriteString("## Previous gates used\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Contact singleton algebra\n\n```text\n")
	b.WriteString(FormatSingletons(a.Singletons))
	b.WriteString("\n```\n\n")
	b.WriteString("## Finite-Dirac/Yukawa edge target\n\n```text\n")
	b.WriteString(FormatTarget(a.Target))
	b.WriteString("\n```\n\n")
	b.WriteString("## Candidate action/functor table\n\n")
	b.WriteString("| Candidate | Native | Sealed | Circular | AF/J/1st-order/one-form compatible | Noncentral | Mixing | Promotable | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Functors.Candidates {
		compat := c.CompatibleWithAF && c.CompatibleWithJ && c.CompatibleWithFirstOrder && c.CompatibleWithOneForms
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.Native, c.Sealed, c.Circular, compat, c.NoncentralOnGeneration, c.MixingCapacity, c.PromotableAsNativeFunctor, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatFunctors(a.Functors))
	b.WriteString("\n```\n\n")
	b.WriteString("## Noncommuting texture capacity\n\n```text\n")
	b.WriteString(FormatOperators(a.Operators))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact\n\n")
	b.WriteString("| Scenario | Class | Result dim | Native | Conditional | 3 masses | CKM | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Moduli.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %t | %t | %t | %t | `%s` |\n", s.Name, s.AssumptionClass, s.ResultingDim, s.Native, s.Conditional, s.ThreeDistinctMassesPossible, s.CKMMisalignmentPossible, s.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	b.WriteString("\n```\n\n")
	b.WriteString("## Firewall status\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Statuses\n\n```text\n")
	b.WriteString(strings.Join(Statuses(a), "\n"))
	b.WriteString("\n```\n\n")
	b.WriteString("## Conclusion\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n## Next gate\n\n```text\n")
	b.WriteString(FormatNext(a.Next))
	b.WriteString("\n```\n")
	return b.String()
}
