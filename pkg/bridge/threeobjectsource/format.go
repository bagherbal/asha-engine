package threeobjectsource

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate395_spinor_two_not_three=%t gate395_triality_category=%t gate394_central=%t gate371_N_non_native=%t gate365_tau_non_native=%t contact_singletons=%d quartic_blocks=%d gate184_fock_contact_blocked=%t gate372_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate395SpinorTwoNotThree, x.Gate395TrialityCategoryOnly, x.Gate394CentralBroadcast, x.Gate371NumberLadderNonNative, x.Gate365TauKMSNonNative, x.Gate151ContactRationalSingletons, x.Gate151QuarticPrimaryBlocks, x.Gate184FockContactActionBlocked, x.Gate372ChargedModuliDim, x.NoEmpiricalFlavorValuesImported, x.Verdict)
}

func FormatSource(x SourceCandidate) string {
	return fmt.Sprintf("%s source=%q native=%t endogenous=%t sealed=%t circular=%t objects=%d exactly_three=%t family=%d selector=%t gen_semantics=%t contact=%t color=%t scalar_trace=%t finite_DF=%t J=%t first_order=%t EW=%t own_noncentral=%t gen_noncentral=%t diagonal=%t mixing=%t spectrum=%v commutant_dim=%d promotable=%t reason=%q verdict=%s", x.Name, x.Source, x.Native, x.Endogenous, x.Sealed, x.CircularIfPromoted, x.ObjectCount, x.ExactlyThreeObjects, x.FamilyCount, x.RequiresSelector, x.GenerationSemantics, x.ContactSemantics, x.ColorSemantics, x.ScalarTraceSemantics, x.CompatibleWithFiniteDirac, x.CompatibleWithJ, x.CompatibleWithFirstOrder, x.CompatibleWithElectroweak, x.NoncentralInOwnAlgebra, x.NoncentralOnGenerationSpace, x.DiagonalOnly, x.MixingCapacity, x.Spectrum, x.CommutantDimension, x.PromotableAsGenerationSource, x.Reason, x.Verdict)
}

func FormatSources(x SourceAudit) string {
	return fmt.Sprintf("executed=%t native_candidates=%d native_three_sources=%d native_generation_sources=%d promotable=%d native_noncentral_gen=%d sealed_noncentral_gen=%d best_native_three=%q verdict=%s", x.Executed, x.NativeCandidateCount, x.NativeExactlyThreeSourceCount, x.NativeGenerationSourceCount, x.PromotableGenerationSourceCount, x.NativeNoncentralOnGeneration, x.SealedNoncentralOnGeneration, x.BestNativeThreeSource, x.Verdict)
}

func FormatPair(x PairAudit) string {
	return fmt.Sprintf("%s left=%q right=%q native_pair=%t sealed_pair=%t eligible=%t comm_norm=%.12g noncommuting=%t ckm=%t reason=%q verdict=%s", x.Name, x.Left, x.Right, x.NativePair, x.SealedPair, x.Eligible, x.CommutatorNorm, x.Noncommuting, x.CKMCapacity, x.Reason, x.Verdict)
}

func FormatOperators(x OperatorAudit) string {
	return fmt.Sprintf("executed=%t native_eligible_ops=%d native_noncentral_ops=%d native_noncommuting_pairs=%d sealed_noncommuting_pairs=%d max_native_comm=%.12g max_sealed_comm=%.12g ckm_native=%t verdict=%s", x.Executed, x.NativeEligibleOperators, x.NativeNoncentralOperators, x.NativeNoncommutingPairs, x.SealedNoncommutingPairs, x.MaxNativeCommutatorNorm, x.MaxSealedCommutatorNorm, x.CKMCapacityNative, x.Verdict)
}

func FormatModuliScenario(x ModuliScenario) string {
	return fmt.Sprintf("%s assumption=%q start=%d result=%d native=%t conditional=%t failed=%t three_masses=%t ckm=%t reason=%q verdict=%s", x.Name, x.AssumptionClass, x.StartingChargedDim, x.ResultingDim, x.Native, x.Conditional, x.Failed, x.ThreeDistinctMassesPossible, x.CKMMisalignmentPossible, x.Reason, x.Verdict)
}

func FormatModuli(x ModuliAudit) string {
	return fmt.Sprintf("executed=%t start=%d native_reduction=%t best_native=%d best_conditional=%d verdict=%s", x.Executed, x.StartingChargedDim, x.NativeReductionBelow13, x.BestNativeDim, x.BestConditionalDim, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_empirical_ordering=%t no_manual_assignment=%t no_contact_promoted=%t no_color_promoted=%t no_fano_promoted=%t no_tau_N_promoted=%t no_native_flavor=%t no_moduli_claim=%t verdict=%s", x.Executed, x.NoMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoEmpiricalOrderingImported, x.NoManualGenerationAssignment, x.NoContactRootsPromoted, x.NoColorModesPromoted, x.NoFanoTriplePromoted, x.NoTauOrNPromoted, x.NoNativeFlavorClaimed, x.NoModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 396 Registry Audit — Endogenous Three-Object Source Search beyond Spinor Chirality\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can any native ASHA structure produce exactly three addressable objects that can act as finite-Dirac generation labels, rather than merely as contact roots, color modes, Fano branches, or sealed scalar traces?\n\n")
	b.WriteString("## Prior gate inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Source candidate table\n\n")
	b.WriteString("| Candidate | Native | Endogenous | Sealed | Objects | Exactly three | Family | Selector | Generation semantics | Finite-Dirac compatible | Noncentral on generation | Promotable | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Sources.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %d | %t | %d | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.Native, c.Endogenous, c.Sealed, c.ObjectCount, c.ExactlyThreeObjects, c.FamilyCount, c.RequiresSelector, c.GenerationSemantics, c.CompatibleWithFiniteDirac, c.NoncentralOnGenerationSpace, c.PromotableAsGenerationSource, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatSources(a.Sources))
	b.WriteString("\n```\n\n")
	b.WriteString("## Candidate diagnostics\n\n")
	for _, c := range a.Sources.Candidates {
		b.WriteString("```text\n")
		b.WriteString(FormatSource(c))
		b.WriteString("\n```\n")
	}
	b.WriteString("\n## Noncommuting operator capacity\n\n```text\n")
	b.WriteString(FormatOperators(a.Operators))
	b.WriteString("\n```\n\n")
	b.WriteString("### Pair diagnostics\n\n")
	for _, p := range a.Operators.Pairs {
		b.WriteString("```text\n")
		b.WriteString(FormatPair(p))
		b.WriteString("\n```\n")
	}
	b.WriteString("\n## Moduli impact table\n\n")
	b.WriteString("| Scenario | Assumption | Resulting dim | Native | Conditional | Failed | Three masses | CKM capacity | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Moduli.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %t | %t | %t | %t | %t | `%s` |\n", s.Name, s.AssumptionClass, s.ResultingDim, s.Native, s.Conditional, s.Failed, s.ThreeDistinctMassesPossible, s.CKMMisalignmentPossible, s.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	b.WriteString("\n```\n\n")
	b.WriteString("## Firewall status\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Conclusion\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n## Next gate\n\n```text\n")
	b.WriteString(FormatNext(a.Next))
	b.WriteString("\n```\n")
	return b.String()
}
