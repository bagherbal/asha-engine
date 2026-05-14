package spectralgraphedgeadjacency

import (
	"fmt"
	"math"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate400_no_q4=%t gate401_anisotropic=%t gate401_no_weighted_q4=%t oneform_edges=%t J_edges=%d first_order_edge_graph=%t inner_fluctuation_fields=%t moduli_dim=%d no_empirical=%t verdict=%s", x.Executed, x.Gate400NoNativeQ4Selector, x.Gate401AnisotropicWeightsFound, x.Gate401NoNativeWeightedLaplacian, x.Gate385OneFormEdges, x.Gate385JDoubledEdgeCount, x.Gate297FirstOrderEdgeGraph, x.Gate298InnerFluctuationFields, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Audit) string {
	return fmt.Sprintf("poly=%s degree=%d irreducible=%t monic=%s verdict=%s", x.Polynomial, x.Degree, x.IrreducibleOverQ, formatFloatSlice(x.MonicCoefficients), x.Verdict)
}

func FormatArena(x EdgeGraphArena) string {
	nodes := make([]string, len(x.Nodes))
	for i, n := range x.Nodes {
		nodes[i] = n.Name + "/" + n.Kind
	}
	edges := make([]string, len(x.StructuralEdges))
	for i, e := range x.StructuralEdges {
		flags := []string{}
		if e.Yukawa {
			flags = append(flags, "Yukawa")
		}
		if e.Majorana {
			flags = append(flags, "Majorana")
		}
		edges[i] = fmt.Sprintf("%s:%s-%s/%s/%s", e.Name, e.Source, e.Target, e.ScalarBranch, strings.Join(flags, "+"))
	}
	return fmt.Sprintf("formalized=%t structural_edges=%d yukawa_edges=%d J_doubled=%d canonical_endpoint_incidence=%t canonical_orientation=%t canonical_Hphi_quotient=%t uses_charge_weights=%t uses_yukawa=%t uses_mass=%t nodes=[%s] edges=[%s] verdict=%s", x.Formalized, x.StructuralEdgeCount, x.YukawaEdgeCount, x.JDoubledEdgeCount, x.HasCanonicalEndpointIncidence, x.HasCanonicalEdgeOrientation, x.HasCanonicalHphiQuotient, x.UsesGaugeChargeWeights, x.UsesYukawaAmplitudes, x.UsesObservedMasses, strings.Join(nodes, "; "), strings.Join(edges, "; "), x.Verdict)
}

func FormatCandidate(c GraphCandidate) string {
	return fmt.Sprintf("name=%s domain=%s dim=%d formula=%s native=%t sealed=%t circular=%t Hphi_endomorphism=%t canonical_quotient=%t graph=%t J=%t first_order=%t charge_weights=%t yukawa=%t masses=%t components=%d eigen=%s distinct=%d min_degree=%d char=%s min_poly=%s char_residual_q4=%s min_residual_q4=%s pair_deg=%t central=%t quartic_capacity=%t q4_exact=%t promotable=%t verdict=%s reason=%s", c.Name, c.Domain, c.Dimension, c.Formula, c.Native, c.Sealed, c.Circular, c.HphiEndomorphism, c.CanonicalQuotientToHphi, c.EdgeGraphDerived, c.CompatibleWithJ, c.CompatibleWithFirstOrder, c.UsesGaugeWeights, c.UsesYukawaAmplitudes, c.UsesObservedMasses, c.Components, formatFloatSlice(c.Eigenvalues), c.DistinctEigenvalues, c.MinimalDegree, c.CharacteristicPolynomial, c.MinimalPolynomial, formatResidual(c.CharacteristicResidualToQ4), formatResidual(c.MinimalResidualToQ4), c.PairDegenerate, c.CentralOnHphi, c.IrreducibleQuarticCapacity, c.Q4ExactMatch, c.PromotableAsQ4Selector, c.Verdict, c.Reason)
}

func FormatSieve(x GraphSieve) string {
	parts := []string{fmt.Sprintf("executed=%t native_graph=%d native_Hphi=%d native_quartic_capacity=%d canonical_q4_matches=%d sealed_manual=%d best_native=%s best_residual=%s verdict=%s", x.Executed, x.NativeGraphOperatorCount, x.NativeHphiEndomorphismCount, x.NativeQuarticCapacityCount, x.CanonicalHphiQ4MatchCount, x.SealedOrManualQ4Count, x.BestNativeCandidate, formatResidual(x.BestNativeQ4Residual), x.Verdict)}
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatImpact(x Impact) string {
	return fmt.Sprintf("Hphi_q4_identified=%t native_edge_adjacency=%t canonical_graph_quotient=%t yukawa_reduced=%t moduli_start=%d moduli_result=%d flavor_firewall=%t higgs_lane_preserved=%t edge_graph_lane_opened_unsealed=%t verdict=%s", x.HphiQuarticIdentified, x.NativeEdgeAdjacencyFound, x.CanonicalGraphQuotientFound, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.EdgeGraphLaneOpenedButUnsealed, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_ckm=%t no_pmns=%t no_yukawa_amplitudes=%t no_charge_fit=%t no_manual_q4=%t no_arbitrary_graph_quotient=%t no_companion_promoted=%t no_moduli_reduction=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoGaugeChargeFitReused, x.NoManualQ4HphiID, x.NoArbitraryGraphQuotient, x.NoCompanionOperatorPromoted, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 402 Registry Audit — Spectral Graph Edge-Adjacency Operator Search\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 402 tests whether the native adjacency/incidence topology of the finite one-form Dirac edge graph produces a canonical four-real scalar operator whose invariant polynomial is the irreducible contact quartic `q4`. It deliberately avoids gauge-charge weights, numerical Yukawa amplitudes, observed masses, and manual q4 placement.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Edge-graph arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Graph candidate table\n\n```text\n")
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

func formatResidual(x float64) string {
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	if math.IsNaN(x) {
		return "NaN"
	}
	return fmt.Sprintf("%.12g", x)
}
