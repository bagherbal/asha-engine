package minimalmodularfamilyhamiltonian

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate411_least_cost_K=%v no_axiom_promoted=%v gate410_no_bundle=%v gate409_U3=%v gate408_scalar_blind=%v charged_moduli=%d no_empirical=%v verdict=%s", x.Executed, x.Gate411LeastCostKGen, x.Gate411NoAxiomPromoted, x.Gate410NoNativeFamilyBundle, x.Gate409TrivialU3Multiplicity, x.Gate408ScalarFlavorBlind, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatHamiltonian(x HamiltonianAxiom) string {
	return fmt.Sprintf("executed=%v name=%s native=%v explicit_axiom=%v trace=%s trace_square=%s eigenvalues=%v distinct=%d hermitian=%v traceless=%v rank=%d minpoly_degree=%d three_level=%v diagonal_only=%v empirical_coefficients=%v verdict=%s reason=%s", x.Executed, x.Name, x.NativeInCurrentAsha, x.ExplicitAxiom, FormatFloat(x.Trace), FormatFloat(x.TraceSquare), x.Eigenvalues, x.DistinctEigenvalues, x.Hermitian, x.Traceless, x.Rank, x.MinimalPolynomialDegree, x.ProvidesThreeLevelOrder, x.DiagonalOnly, x.CoefficientsEmpirical, x.Verdict, x.Reason)
}

func FormatKMS(x KMSState) string {
	return fmt.Sprintf("executed=%v beta=%s weights=[%s] Z=%s positive=%v normalized=%v tracial=%v entropy=%s max_ratio=%s modular_flow_active=%v verdict=%s reason=%s", x.Executed, FormatFloat(x.Beta), formatFloats(x.Weights), FormatFloat(x.PartitionFunction), x.Positive, x.Normalized, x.Tracial, FormatFloat(x.Entropy), FormatFloat(x.MaxWeightRatio), x.ModularFlowActive, x.Verdict, x.Reason)
}

func FormatCompatibility(x CompatibilityAudit) string {
	return fmt.Sprintf("executed=%v family_only=%v commutes_AF=%v commutes_gauge=%v commutes_Y=%v commutes_SU2L=%v commutes_BL=%v Gamma=%v J_if_mirrored=%v first_order_if_DF_broadcast=%v requires_family_axiom=%v verdict=%s reason=%s", x.Executed, x.ActsOnlyOnFamilyFiber, x.CommutesWithAF, x.CommutesWithGaugeCharges, x.CommutesWithHypercharge, x.CommutesWithSU2L, x.CommutesWithBL, x.CompatibleWithGamma, x.JCompatibleIfMirrored, x.FirstOrderUnaffectedIfDFBroadcast, x.RequiresFamilyFiberAxiom, x.Verdict, x.Reason)
}

func FormatMixing(x MixingAudit) string {
	return fmt.Sprintf("executed=%v operators=%s native_noncommuting=%d conditional_noncommuting=%d comm_K_K2=%s comm_K_gauge=%s ckm_native=%v pmns_native=%v ckm_conditional=%v pmns_conditional=%v diagonal_only=%v verdict=%s reason=%s", x.Executed, strings.Join(x.OperatorsAudited, ","), x.NativeNoncommutingPairs, x.ConditionalNoncommutingPairs, FormatFloat(x.CommutatorKWithK2Norm), FormatFloat(x.CommutatorKWithGaugeNorm), x.CKMNative, x.PMNSNative, x.CKMConditional, x.PMNSConditional, x.DiagonalOnly, x.Verdict, x.Reason)
}

func FormatSectorMap(x SectorMapAudit) string {
	return fmt.Sprintf("executed=%v universal_ordering=%v up_native=%v down_native=%v lepton_native=%v sector_maps_needed=%v observed_yukawas_inserted=%v hierarchy_capacity=%v three_masses_conditional=%v verdict=%s reason=%s", x.Executed, x.UniversalFamilyOrdering, x.UpSectorMapNative, x.DownSectorMapNative, x.LeptonSectorMapNative, x.SectorSpecificMapsNeeded, x.ObservedYukawasInserted, x.MassHierarchyCapacity, x.ThreeDistinctMassesPossibleConditionally, x.Verdict, x.Reason)
}

func FormatFirewall(x EmpiricalFirewall) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawas=%v no_sector_amplitudes=%v K_axiom_only=%v no_native_derivation=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaMatricesInserted, x.NoSectorAmplitudesInserted, x.KGenPromotedAsAxiomOnly, x.NoNativeDerivationClaimed, x.Verdict)
}

func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("name=%s status=%s moduli_dim=%d three_masses=%v CKM=%v PMNS=%v native_reduction=%v conditional=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.NativeReduction, x.ConditionalOnly, x.Reason)
}

func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d scenarios=%d best_native_dim=%d native_reduction_below13=%v conditional_hierarchy=%v conditional_CKM_PMNS=%v firewall=%v verdict=%s", x.StartDim, len(x.Scenarios), x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalHierarchy, x.ConditionalCKMPMNS, x.FirewallPreserved, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func CompactStatusTable(a Analysis) string { return strings.Join(Statuses(a), "\n") }

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 412 Registry Audit — Minimal Modular Family Hamiltonian Axiom Consistency Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 412 tests the least-cost Gate 411 extension candidate as an explicit axiom: a centered modular family Hamiltonian `K_gen` on the three-dimensional family fiber. The gate asks whether this axiom is compatible with the existing finite spectral triple and whether it supplies hierarchy or CKM/PMNS capacity without empirical inputs.\n\n")
	b.WriteString("## Prior boundary inherited\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Minimal Hamiltonian axiom\n\n```text\n")
	b.WriteString(FormatHamiltonian(a.Hamiltonian))
	b.WriteString("\n```\n\n")
	b.WriteString("## Nontracial KMS state\n\n```text\n")
	b.WriteString(FormatKMS(a.KMS))
	b.WriteString("\n```\n\n")
	b.WriteString("## Compatibility audit\n\n```text\n")
	b.WriteString(FormatCompatibility(a.Compatibility))
	b.WriteString("\n```\n\n")
	b.WriteString("## CKM / PMNS mixing capacity\n\n```text\n")
	b.WriteString(FormatMixing(a.Mixing))
	b.WriteString("\n```\n\n")
	b.WriteString("## Sector mass-map audit\n\n```text\n")
	b.WriteString(FormatSectorMap(a.SectorMap))
	b.WriteString("\n```\n\n")
	b.WriteString("## Empirical firewall\n\n```text\n")
	b.WriteString(FormatFirewall(a.Firewall))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact\n\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	for _, s := range a.Moduli.Scenarios {
		b.WriteString("\n")
		b.WriteString(FormatScenario(s))
	}
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

func formatFloats(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = FormatFloat(x)
	}
	return strings.Join(parts, ", ")
}
