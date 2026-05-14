package familybundleaxiomledger

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate410_no_native_bundle=%v KO_only_signs=%v KMS_external_K=%v ideals_wrong_domain=%v new_axiom_required=%v gate409_U3=%v gate408_scalar_blind=%v charged_moduli=%d no_empirical=%v verdict=%s", x.Executed, x.Gate410NoNativeFamilyBundle, x.Gate410KOTwistOnlySigns, x.Gate410KMSRequiresExternalK, x.Gate410PrimitiveIdealsWrongDomain, x.Gate410RequiresNewAxiom, x.Gate409TrivialU3Multiplicity, x.Gate408ScalarFlavorBlind, x.Gate372ChargedModuliDim, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatCandidate(x AxiomCandidate) string {
	return fmt.Sprintf("name=%s kind=%s data=%s cost=%d native=%v promoted=%v changes_family=%v replaces_C3=%v families3=%v diagonal=%v noncommuting=%v ckm=%v pmns=%v pure=%v empirical_independent=%v risk=%s AF=%v J=%v Gamma=%v first_order=%v gauge=%v new_axiom=%v external_K=%v algebra_ext=%v functor=%v connection=%v empirical_yukawa=%v verdict=%s reason=%s", x.Name, x.Kind, x.MinimalAdditionalData, x.MathematicalCost, x.NativeInCurrentAsha, x.PromotedToTheorem, x.ChangesFamilyCarrier, x.ReplacesTrivialC3, x.ProvidesThreeFamilies, x.ProvidesDiagonalOperator, x.ProvidesTwoNoncommutingOps, x.CKMCapacity, x.PMNSCapacity, x.PureGeometricFormulable, x.EmpiricalIndependentFormulable, x.CurveFittingRisk, x.PreservesAF, x.PreservesJ, x.PreservesGamma, x.PreservesFirstOrder, x.PreservesGaugeCharges, x.RequiresNewAxiom, x.RequiresExternalHamiltonian, x.RequiresAlgebraExtension, x.RequiresFunctor, x.RequiresConnection, x.RequiresEmpiricalYukawas, x.Verdict, x.Reason)
}

func FormatLedger(x AxiomLedger) string {
	return fmt.Sprintf("executed=%v candidates=%d promoted=%d pure_geometric=%d empirical_independent=%d high_risk=%d ckm_capable=%d pmns_capable=%d lowest_cost=%d least_cost=%s verdict=%s", x.Executed, x.CandidatesAudited, x.PromotedAxioms, x.PureGeometricCandidates, x.EmpiricalIndependentCount, x.CurveFittingRiskHigh, x.CKMCapableCandidates, x.PMNSCapableCandidates, x.LowestCost, strings.Join(x.LeastCostNames, ","), x.Verdict)
}

func FormatCapacity(x CapacityAudit) string {
	return fmt.Sprintf("executed=%v diagonal_only=%d noncommuting_candidates=%d native_pairs=%d conditional_pairs=%d ckm_native=%v ckm_conditional=%v pmns_native=%v pmns_conditional=%v verdict=%s reason=%s", x.Executed, x.CandidatesWithDiagonalOnly, x.CandidatesWithNoncommuting, x.NativeNoncommutingPairs, x.ConditionalNoncommutingPairs, x.CKMNative, x.CKMConditional, x.PMNSNative, x.PMNSConditional, x.Verdict, x.Reason)
}

func FormatEmpiricalIndependence(x EmpiricalIndependenceAudit) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa_matrices=%v pure_rule_candidates=%d fitting_candidates=%d source_risk=%v verdict=%s reason=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaMatricesInserted, x.CandidatesCanBePureRules, x.CandidatesCollapseToFitting, x.UnconstrainedFamilySourceRisk, x.Verdict, x.Reason)
}

func FormatCostRow(x CostRow) string {
	return fmt.Sprintf("rank=%d name=%s cost=%d benefit=%s risk=%s next_test=%s", x.Rank, x.Name, x.Cost, x.Benefit, x.MainRisk, x.NextTest)
}

func FormatRanking(x CostRanking) string {
	return fmt.Sprintf("executed=%v rows=%d verdict=%s", x.Executed, len(x.Rows), x.Verdict)
}

func FormatBoundary(x Boundary) string {
	return fmt.Sprintf("executed=%v lawspace_native=%v family_native=%v new_axiom_required=%v current_flavor_complete=%v verdict=%s statement=%s", x.Executed, x.LawSpaceNative, x.FamilyBundleNative, x.NewAxiomRequiredForFamilies, x.CurrentASHAFlavorComplete, x.Verdict, x.Statement)
}

func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("name=%s status=%s moduli_dim=%d masses3=%v ckm=%v pmns=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.Reason)
}

func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d scenarios=%d best_native_dim=%d native_reduction=%v conditional_reduction=%v firewall=%v verdict=%s", x.StartDim, len(x.Scenarios), x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalReductionBelow13, x.FirewallPreserved, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa=%v no_axiom=%v no_external_K=%v no_connection=%v no_algebra_ext=%v no_functor=%v no_moduli_reduction=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaAmplitudesInserted, x.NoAxiomPromoted, x.NoExternalHamiltonianPromoted, x.NoFamilyConnectionPromoted, x.NoAlgebraExtensionPromoted, x.NoFunctorPromoted, x.NoModuliReductionClaimed, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%s reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func CompactStatusTable(a Analysis) string { return strings.Join(Statuses(a), "\n") }

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 411 Registry Audit — Axiom-Candidate Ledger for Nontrivial Family Bundle Extensions\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Gate 411 classifies the minimal mathematical axioms that could extend ASHA into the flavor/family sector after Gate 410 proved no native nontrivial family bundle is derived. It is a ledger and boundary theorem, not a promotion of new axioms and not an empirical Yukawa seal.\n\n")
	b.WriteString("## Prior boundary inherited\n\n```text\n")
	b.WriteString(FormatInheritance(a.Inheritance))
	b.WriteString("\n```\n\n")
	b.WriteString("## Axiom candidate ledger\n\n```text\n")
	b.WriteString(FormatLedger(a.Ledger))
	for _, c := range a.Ledger.Candidates {
		b.WriteString("\n")
		b.WriteString(FormatCandidate(c))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## CKM / PMNS capacity audit\n\n```text\n")
	b.WriteString(FormatCapacity(a.Capacity))
	b.WriteString("\n```\n\n")
	b.WriteString("## Empirical independence check\n\n```text\n")
	b.WriteString(FormatEmpiricalIndependence(a.EmpiricalIndependence))
	b.WriteString("\n```\n\n")
	b.WriteString("## Cost ranking\n\n```text\n")
	b.WriteString(FormatRanking(a.Ranking))
	for _, r := range a.Ranking.Rows {
		b.WriteString("\n")
		b.WriteString(FormatCostRow(r))
	}
	b.WriteString("\n```\n\n")
	b.WriteString("## Epistemological boundary\n\n```text\n")
	b.WriteString(FormatBoundary(a.Boundary))
	b.WriteString("\n```\n\n")
	b.WriteString("## Moduli impact table\n\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	for _, s := range a.Moduli.Scenarios {
		b.WriteString("\n")
		b.WriteString(FormatScenario(s))
	}
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
