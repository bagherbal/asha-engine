package dynamicgenerationlabels

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate394_centrality=%t gate394_native_noncentral=%d gate394_native_noncommuting=%d gate393_domain_admitted=%t gate247_functor=%t gate372_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate394CentralityFirewall, x.Gate394NativeNoncentralOperators, x.Gate394NativeNoncommutingPairs, x.Gate393DomainAdmitted, x.Gate247TrialityFunctorDerived, x.Gate372ChargedModuliDim, x.NoEmpiricalFlavorValuesImported, x.Verdict)
}

func FormatSpinor(x SpinorDecompositionAudit) string {
	return fmt.Sprintf("executed=%t algebra=%q full_dim=%d chiral_split=%v native_sector_count=%d has_three=%t has_8v_inside_spinor=%t labels_derived=%t chirality_spectrum=%v verdict=%s", x.Executed, x.CliffordAlgebra, x.FullSpinorRealDimension, x.ChiralSplit, x.NativeChiralSectorCount, x.HasThreeNativeSectors, x.HasVectorRepresentationInsideSpinorSplit, x.GenerationLabelsDerived, x.ChiralityOperatorSpectrum, x.Verdict)
}

func FormatTriality(x TrialityAudit) string {
	return fmt.Sprintf("executed=%t reps=%v automorphism=%q category_triple=%t acts_on_generation_copies=%t explicit_theta_DF_flavor=%t native_functor_C3=%t vector_rep_native=%t verdict=%s", x.Executed, x.RepresentationTypes, x.AutomorphismGroup, x.CategoryLevelTriple, x.ActsOnPhysicalGenerationCopies, x.ExplicitThetaOnFiniteDiracFlavor, x.NativeFunctorToC3Gen, x.VectorRepSuppliedNatively, x.Verdict)
}

func FormatLabel(x LabelCandidate) string {
	return fmt.Sprintf("%s source=%q native=%t sealed=%t circular=%t dim=%d sectors=%d central=%t noncentral=%t diagonal=%t mixing=%t spectrum=%v commutant_dim=%d labels_derived=%t finite_DF=%t J=%t first_order=%t EW=%t reason=%q verdict=%s", x.Name, x.Source, x.Native, x.Sealed, x.Circular, x.Dimension, x.SectorCount, x.Central, x.NonCentral, x.DiagonalOnly, x.Mixing, x.Spectrum, x.CommutantDimension, x.GenerationLabelsDerived, x.CompatibleWithFiniteDirac, x.CompatibleWithJ, x.CompatibleWithFirstOrder, x.CompatibleWithHyperchargeSU2, x.Reason, x.Verdict)
}

func FormatLabels(x DynamicLabelAudit) string {
	return fmt.Sprintf("executed=%t native_candidates=%d native_generation_labels=%d native_noncentral=%d sealed_noncentral=%d verdict=%s", x.Executed, x.NativeCandidateCount, x.NativeGenerationLabelCount, x.NativeNoncentralCount, x.SealedNoncentralCount, x.Verdict)
}

func FormatPair(x PairAudit) string {
	return fmt.Sprintf("%s native_pair=%t sealed_pair=%t comm_norm=%.12g noncommuting=%t ckm_capacity=%t reason=%q verdict=%s", x.Name, x.NativePair, x.SealedPair, x.CommutatorNorm, x.Noncommuting, x.CKMCapacity, x.Reason, x.Verdict)
}

func FormatOperators(x OperatorCapacityAudit) string {
	return fmt.Sprintf("executed=%t native_ops=%d native_generation_ops=%d native_noncentral_ops=%d native_noncommuting_pairs=%d sealed_noncommuting_pairs=%d max_native_comm=%.12g max_sealed_comm=%.12g ckm_native=%t verdict=%s", x.Executed, x.NativeOperators, x.NativeGenerationOperators, x.NativeNoncentralOperators, x.NativeNoncommutingPairs, x.SealedNoncommutingPairs, x.MaxNativeCommutatorNorm, x.MaxSealedCommutatorNorm, x.CKMCapacityNative, x.Verdict)
}

func FormatModuliScenario(x ModuliScenario) string {
	return fmt.Sprintf("%s assumption=%q start=%d result=%d native=%t conditional=%t failed=%t three_masses=%t ckm=%t reason=%q verdict=%s", x.Name, x.AssumptionClass, x.StartingChargedDim, x.ResultingDim, x.Native, x.Conditional, x.Failed, x.ThreeDistinctMassesPossible, x.CKMMisalignmentPossible, x.Reason, x.Verdict)
}

func FormatModuli(x ModuliAudit) string {
	return fmt.Sprintf("executed=%t start=%d native_reduction=%t best_native=%d best_conditional=%d verdict=%s", x.Executed, x.StartingChargedDim, x.NativeReductionBelow13, x.BestNativeDim, x.BestConditionalDim, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_empirical_ordering=%t no_manual_assignment=%t no_triality_promoted=%t no_N_promoted=%t no_native_claim=%t no_moduli_claim=%t verdict=%s", x.Executed, x.NoMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoEmpiricalOrderingImported, x.NoManualGenerationAssignment, x.NoTrialityLabelsPromoted, x.NoNPromoted, x.NoNativeFlavorClaimed, x.NoModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 395 Registry Audit — Representation-Origin Search for Dynamic Generation Labels\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can native `Cℓ(1,7)` spinor representation theory dynamically derive three generation labels and activate native noncommuting flavor texture capacity?\n\n")
	b.WriteString("## Prior gate inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Spinor decomposition audit\n\n```text\n")
	b.WriteString(FormatSpinor(a.Spinor))
	b.WriteString("\n```\n\n")
	b.WriteString("## Triality category audit\n\n```text\n")
	b.WriteString(FormatTriality(a.Triality))
	b.WriteString("\n```\n\n")
	b.WriteString("## Dynamic label candidate table\n\n")
	b.WriteString("| Candidate | Native | Sealed | Circular | Dimension | Sectors | Central | Noncentral | Mixing | Labels derived | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Labels.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %d | %d | %t | %t | %t | %t | `%s` |\n", c.Name, c.Native, c.Sealed, c.Circular, c.Dimension, c.SectorCount, c.Central, c.NonCentral, c.Mixing, c.GenerationLabelsDerived, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatLabels(a.Labels))
	b.WriteString("\n```\n\n")
	b.WriteString("### Candidate diagnostics\n\n")
	for _, c := range a.Labels.Candidates {
		b.WriteString("```text\n")
		b.WriteString(FormatLabel(c))
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
