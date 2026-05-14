package generationaddressfunctor

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate393_domain_blocked=%t gate370_central=%t gate371_N_non_native=%t gate372_dim=%d gate385_oneform=%t no_empirical=%t verdict=%s", x.Executed, x.Gate393DomainNotAdmitted, x.Gate370NativeSupportMapsCentral, x.Gate371NumberOperatorNonNative, x.Gate372ChargedModuliDim, x.Gate385OneFormEdgeSupportDerived, x.NoEmpiricalFlavorValuesImported, x.Verdict)
}

func FormatTarget(x FunctorTarget) string {
	return fmt.Sprintf("domain=%q codomain=%q required=%q success=%q verdict=%s", x.Domain, x.Codomain, x.RequiredNativePattern, x.SuccessfulPattern, x.Verdict)
}

func FormatCandidate(x Candidate) string {
	return fmt.Sprintf("%s source=%q native=%t sealed=%t circular=%t derived=%t central=%t noncentral=%t diagonal=%t rank=%d spectrum=%v commutant_dim=%d breaks_degeneracy=%t mixing=%t reason=%q verdict=%s", x.Name, x.Source, x.Native, x.Sealed, x.Circular, x.DerivedFromExistingLedger, x.Central, x.NonCentral, x.DiagonalOnly, x.Rank, x.Spectrum, x.CommutantDimension, x.BreaksGenerationDegeneracy, x.GivesMixing, x.Reason, x.Verdict)
}

func FormatCandidateAudit(x CandidateAudit) string {
	return fmt.Sprintf("executed=%t native=%d native_noncentral=%d sealed_noncentral=%d central_native=%d verdict=%s", x.Executed, x.NativeCandidateCount, x.NativeNoncentralCount, x.SealedNoncentralCount, x.CentralNativeCount, x.Verdict)
}

func FormatSource(x SourceAudit) string {
	return fmt.Sprintf("%s executed=%t rank=%d spectrum=%v central_only=%t native_noncentral=%t sealed_or_circular_only=%t result=%q verdict=%s", x.Name, x.Executed, x.Rank, x.Spectrum, x.CentralOnly, x.NativeNoncentralFound, x.CircularOrSealedOnly, x.Result, x.Verdict)
}

func FormatNumber(x NumberOperatorAudit) string {
	return fmt.Sprintf("executed=%t status=%q native=%t bridge=%t sealed=%t circular=%t derivation_residual=%v comm_cycle=%.12g comm_mirror=%.12g hypercharge=%t su2l=%t J=%t Gamma=%t DF_edge=%t breaks_triality=%t hierarchy=%t mixing=%t verdict=%s", x.Executed, x.Status, x.Native, x.BridgeCompatible, x.SealedExternalExtension, x.CircularIfUsedAsSolution, x.DerivationResidual, x.CommutatorWithCycle, x.CommutatorWithMirror, x.CommutesWithHypercharge, x.CommutesWithSU2L, x.CommutesWithJ, x.CommutesWithGamma, x.CommutesWithDFEdgeSupport, x.BreaksExactTriality, x.ProducesHierarchy, x.ProducesMixing, x.Verdict)
}

func FormatPair(x PairAudit) string {
	return fmt.Sprintf("%s native_pair=%t sealed_pair=%t comm_norm=%.12g noncommuting=%t simultaneously_diagonalized=%t ckm_capacity=%t reason=%q verdict=%s", x.Name, x.NativePair, x.SealedPair, x.CommutatorNorm, x.Noncommuting, x.SimultaneouslyDiagonalized, x.CKMCapacity, x.Reason, x.Verdict)
}

func FormatTexture(x TextureCapacityAudit) string {
	return fmt.Sprintf("executed=%t native_ops=%d native_noncentral_ops=%d native_noncommuting_pairs=%d sealed_noncommuting_pairs=%d max_native_comm=%.12g max_sealed_comm=%.12g simultaneous=%t ckm_native=%t verdict=%s", x.Executed, x.NativeGenerationOperators, x.NativeNoncentralOperators, x.NativeNoncommutingPairs, x.SealedNoncommutingPairs, x.MaxNativeCommutatorNorm, x.MaxSealedCommutatorNorm, x.SimultaneouslyDiagonalizable, x.CKMCapacityNative, x.Verdict)
}

func FormatModuliScenario(x ModuliScenario) string {
	return fmt.Sprintf("%s assumption=%q start=%d result=%d native=%t conditional=%t failed=%t distinct_masses=%t ckm=%t ql_separation=%t reason=%q verdict=%s", x.Name, x.AssumptionClass, x.StartingChargedDim, x.ResultingDim, x.Native, x.Conditional, x.Failed, x.DistinctChargedMassesPossible, x.CKMMisalignmentPossible, x.LeptonQuarkSectorSeparation, x.Reason, x.Verdict)
}

func FormatModuli(x ModuliAudit) string {
	return fmt.Sprintf("executed=%t start=%d native_reduction=%t best_native=%d best_conditional=%d verdict=%s", x.Executed, x.StartingChargedDim, x.NativeReductionBelow13, x.BestNativeDim, x.BestConditionalDim, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_empirical_ordering=%t no_manual_assignment=%t no_tau=%t no_N_promoted=%t no_native_claim=%t no_moduli_claim=%t verdict=%s", x.Executed, x.NoYukawaMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoEmpiricalOrderingImported, x.NoManualGenerationAssignment, x.NoCircularTauInserted, x.NoCircularNPromoted, x.NoNativeAddressClaimed, x.NoModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 394 Registry Audit — Native Generation-Address Functor from Triality/Morita Edge Incidence\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Can existing ASHA finite data derive a native noncentral functor into `End(C^3_gen)`?\n\n")
	b.WriteString("```text\n")
	b.WriteString(FormatTarget(a.Target))
	b.WriteString("\n```\n\n")
	b.WriteString("## Prior gate inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Candidate source table\n\n")
	b.WriteString("| Candidate | Native | Sealed | Circular | Central | Noncentral | Diagonal only | Mixing | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---|\n")
	for _, c := range a.Candidates.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %t | %t | %t | `%s` |\n", c.Name, c.Native, c.Sealed, c.Circular, c.Central, c.NonCentral, c.DiagonalOnly, c.GivesMixing, c.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatCandidateAudit(a.Candidates))
	b.WriteString("\n```\n\n")
	b.WriteString("## Source audits\n\n")
	for _, src := range []SourceAudit{a.TrialityBranch, a.MoritaEdge, a.OneFormSupport} {
		b.WriteString("### ")
		b.WriteString(src.Name)
		b.WriteString("\n\n```text\n")
		b.WriteString(FormatSource(src))
		b.WriteString("\n```\n\n")
	}
	b.WriteString("## Fock number-operator audit\n\n```text\n")
	b.WriteString(FormatNumber(a.Number))
	b.WriteString("\n```\n\n")
	b.WriteString("## Noncommuting texture capacity\n\n```text\n")
	b.WriteString(FormatTexture(a.TextureCapacity))
	b.WriteString("\n```\n\n")
	if len(a.TextureCapacity.Pairs) > 0 {
		b.WriteString("### Pair diagnostics\n\n")
		for _, p := range a.TextureCapacity.Pairs {
			b.WriteString("```text\n")
			b.WriteString(FormatPair(p))
			b.WriteString("\n```\n")
		}
		b.WriteString("\n")
	}
	b.WriteString("## Moduli impact table\n\n")
	b.WriteString("| Scenario | Assumption | Resulting dim | Native | Conditional | Failed | CKM capacity | Verdict |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Moduli.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %t | %t | %t | %t | `%s` |\n", s.Name, s.AssumptionClass, s.ResultingDim, s.Native, s.Conditional, s.Failed, s.CKMMisalignmentPossible, s.Verdict))
	}
	b.WriteString("\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	b.WriteString("\n```\n\n")
	b.WriteString("## Firewall status\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Conclusion\n\n")
	b.WriteString(a.Truth)
	b.WriteString("\n\n")
	b.WriteString("## Next gate\n\n```text\n")
	b.WriteString(FormatNext(a.Next))
	b.WriteString("\n```\n")
	return b.String()
}
