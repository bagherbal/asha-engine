package mixededgelaplaciansieve

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate398_no_hphi_id=%t gate399_H_disjoint=%t oneform_edge_support=%t edge_count=%d Hphi_dim=%d pair_degenerate=%t charged_moduli=%d no_empirical=%t verdict=%s", x.Executed, x.Gate398NoCanonicalHphiID, x.Gate399QuaternionicDisjoint, x.Gate385OneFormEdgeSupportDerived, x.Gate385JDoubledEdgeCount, x.Gate37HphiRealDim, x.Gate37PairDegenerateScalarResponse, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatQ4(x Q4Audit) string {
	return fmt.Sprintf("q4=%s degree=%d irreducible_Q=%t contact_primary=%t branch_free=%t verdict=%s", x.Polynomial, x.Degree, x.IrreducibleOverQ, x.ContactPrimary, x.BranchFree, x.Verdict)
}

func FormatArena(x LaplacianArena) string {
	return fmt.Sprintf("formalized=%t object=%s edge_dim=%d Hphi_dim=%d contact_nodes=%d oneform_measure=%t uniform_edge_metric=%t explicit_DF_edge_weights=%t physical_masses_inserted=%t verdict=%s", x.Formalized, x.Object, x.EdgeSupportDimension, x.HphiDimension, x.ContactNodeDimension, x.OneFormEdgeMeasureDerived, x.UniformEdgeMetric, x.ExplicitDFEdgeWeightsDerived, x.PhysicalMassesInserted, x.Verdict)
}

func FormatCandidate(c MixedOperatorCandidate) string {
	return fmt.Sprintf("name=%s formula=%s native=%t sealed=%t circular=%t Hphi_endomorphism=%t contact_compressed=%t oneform_edge=%t gauge=%t J=%t first_order=%t min_degree=%d char=%s pattern=%s pair_deg=%t central=%t q4_exact=%t q4_factor=%t promotable=%t residual=%.12g verdict=%s reason=%s", c.Name, c.Formula, c.Native, c.Sealed, c.Circular, c.HphiEndomorphism, c.ContactCompressed, c.OneFormEdgeDerived, c.GaugeCompatible, c.CompatibleWithJ, c.CompatibleWithFirstOrder, c.MinimalDegree, c.CharacteristicPolynomial, c.EigenvaluePattern, c.PairDegenerate, c.CentralOnHphi, c.Q4ExactMatch, c.Q4FactorMatch, c.PromotableAsQ4Selector, c.Residual, c.Verdict, c.Reason)
}

func FormatMixed(x MixedInvariantAudit) string {
	return fmt.Sprintf("executed=%t native_Hphi_endomorphisms=%d native_q4_matches=%d promotable_native=%d sealed_q4_matches=%d best_native=%s verdict=%s", x.Executed, x.NativeHphiEndomorphismCount, x.NativeQ4MatchCount, x.PromotableNativeCount, x.SealedQ4MatchCount, x.BestNativeCandidate, x.Verdict)
}

func FormatImpact(x IdentityImpact) string {
	return fmt.Sprintf("Hphi_quartic_identified=%t scalar_bundle_sealed=%t oneform_edge_functor=%t yukawa_reduced=%t charged_moduli=%d->%d flavor_firewall=%t higgs_lane_preserved=%t verdict=%s", x.HphiQuarticIdentified, x.ScalarBundleGeometricallySealed, x.OneFormEdgeFunctorDerived, x.YukawaCouplingsReduced, x.ChargedModuliStart, x.ChargedModuliResult, x.FlavorFirewallPreserved, x.HiggsLanePreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%t no_masses=%t no_CKM=%t no_PMNS=%t no_Higgs_inserted=%t no_manual_q4_Hphi=%t no_companion_promoted=%t no_arbitrary_basis=%t no_yukawa_claim=%t no_flavor_reduction_claim=%t verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoObservedHiggsInserted, x.NoManualQ4HphiID, x.NoCompanionOperatorPromoted, x.NoArbitraryBasisMapPromoted, x.NoYukawaCouplingClaimed, x.NoFlavorModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s\nReason: %s\nPrimary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 400 Registry Audit — Non-Quaternionic Scalar Identity / Mixed Edge Laplacian Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 400 tests whether a mixed one-form edge/contact invariant on the four-real scalar carrier `H_phi` supplies the missing basis-free identity selector whose invariant polynomial is the irreducible contact quartic `q4`.\n\n")
	b.WriteString("## Inheritance\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Contact q4 target\n\n```text\n")
	b.WriteString(FormatQ4(a.Q4))
	b.WriteString("\n```\n\n")
	b.WriteString("## Edge Laplacian arena\n\n```text\n")
	b.WriteString(FormatArena(a.Arena))
	b.WriteString("\n```\n\n")
	b.WriteString("## Mixed invariant candidates\n\n")
	b.WriteString("| Candidate | Native | H_phi endomorphism | Minimal degree | Characteristic polynomial | q4 match | Promotable | Verdict |\n")
	b.WriteString("|---|---:|---:|---:|---|---:|---:|---|\n")
	for _, c := range a.Mixed.Candidates {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %d | %s | %t | %t | `%s` |\n", c.Name, c.Native, c.HphiEndomorphism, c.MinimalDegree, c.CharacteristicPolynomial, c.Q4ExactMatch, c.PromotableAsQ4Selector, c.Verdict))
	}
	b.WriteString("\n```text\n")
	for _, c := range a.Mixed.Candidates {
		b.WriteString(FormatCandidate(c))
		b.WriteString("\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Mixed audit summary\n\n```text\n")
	b.WriteString(FormatMixed(a.Mixed))
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
