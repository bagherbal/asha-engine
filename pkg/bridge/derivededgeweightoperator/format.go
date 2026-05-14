package derivededgeweightoperator

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate400_uniform_central=%t gate400_pair_deg=%t gate400_no_q4=%t gate385_edges=%t edge_count=%d gate26_yukawa=%t gate41_hypercharge=%t charged_moduli=%d no_empirical=%t verdict=%s", x.Executed, x.Gate400UniformCentral, x.Gate400PairDegenerateCompression, x.Gate400NoNativeQ4Selector, x.Gate385OneFormEdges, x.Gate385JDoubledEdgeCount, x.Gate26YukawaChannelsDerived, x.Gate41HyperchargeNormalization, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Audit) string {
	return fmt.Sprintf("q4=%s degree=%d irreducible_Q=%t monic=%s verdict=%s", x.Polynomial, x.Degree, x.IrreducibleOverQ, formatFloatSlice(x.MonicCoefficients), x.Verdict)
}

func FormatArena(x EdgeWeightArena) string {
	return fmt.Sprintf("formalized=%t structural_edges=%d J_doubled=%d Hphi_dim=%d native_Y=%t native_B-L=%t native_T3=%t yukawa_amplitudes_used=%t observed_masses_used=%t edges=%s verdict=%s", x.Formalized, x.StructuralEdgeCount, x.JDoubledEdgeCount, x.HphiDimension, x.NativeElectroweakWeights, x.NativeBMinusLWeights, x.NativeT3Weights, x.ExplicitYukawaAmplitudesUsed, x.ObservedMassesUsed, edgeNames(x.Edges), x.Verdict)
}

func FormatCandidate(c WeightedCandidate) string {
	return fmt.Sprintf("name=%s source=%s formula=%s native=%t sealed=%t circular=%t Hphi_endomorphism=%t canonical_Hphi=%t edge_resolved=%t branch_compressed=%t J_real=%t gauge=%t uses_yukawa=%t uses_mass=%t eigen=%s distinct=%d min_degree=%d char=%s residual_q4=%.12g pair_deg=%t central=%t quartic_capacity=%t q4_exact=%t promotable=%t verdict=%s reason=%s", c.Name, c.WeightSource, c.Formula, c.NativeWeights, c.Sealed, c.Circular, c.HphiEndomorphism, c.CanonicalCompressionToHphi, c.EdgeResolved, c.BranchCompressed, c.JRealDoubled, c.GaugeChargeDerived, c.UsesYukawaAmplitudes, c.UsesObservedMasses, formatFloatSlice(c.Eigenvalues), c.DistinctEigenvalues, c.MinimalDegree, c.CharacteristicPolynomial, c.CharacteristicResidualToQ4, c.PairDegenerate, c.CentralOnHphi, c.IrreducibleQuarticCapacity, c.Q4ExactMatch, c.PromotableAsQ4Selector, c.Verdict, c.Reason)
}

func FormatSieve(x WeightSieve) string {
	parts := []string{fmt.Sprintf("executed=%t native_anisotropic=%d native_quartic_capacity=%d canonical_q4_matches=%d sealed_matches=%d best_native=%s best_residual=%.12g verdict=%s", x.Executed, x.NativeAnisotropicCount, x.NativeQuarticCapacityCount, x.CanonicalHphiQ4MatchCount, x.SealedOrNoncanonicalMatches, x.BestNativeCandidate, x.BestNativeQ4Residual, x.Verdict)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("Hphi_q4_identified=%t scalar_bundle_sealed=%t diff_weights=%t canonical_weighted_laplacian=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t higgs_lane_preserved=%t verdict=%s", x.HphiQuarticIdentified, x.ScalarBundleGeometricallySealed, x.DifferentiatedEdgeWeightsFound, x.CanonicalWeightedLaplacianFound, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_yukawa_amplitudes=%t no_manual_q4=%t no_arbitrary_edge_component_map=%t no_affine_fit=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoManualQ4HphiID, x.NoArbitraryEdgeComponentMap, x.NoAffineChargeFitPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 401 Registry Audit — Derived Edge-Weight Operator / Hypercharge Laplacian Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 401 tests whether native electroweak, B-L, or T3-like charges can differentiate the ten J-doubled finite-Dirac one-form edges strongly enough to produce the irreducible contact quartic `q4` as a canonical invariant on the four-real scalar carrier `H_phi`.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Edge-weight arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Weighted candidate table\n\n```text\n")
	b.WriteString(FormatSieve(a.Sieve))
	b.WriteString("\n```\n\n")
	b.WriteString("## Identity / impact audit\n\n```text\n")
	b.WriteString(FormatImpact(a.Impact))
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
