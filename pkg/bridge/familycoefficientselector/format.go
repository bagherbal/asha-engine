package familycoefficientselector

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate413_pair_compatible=%v pair_not_native=%v ckm_capacity=%v coefficients_free=%v gate412_K_diagonal=%v gate411_ledger=%v charged_moduli=%d verdict=%s", x.Executed, x.Gate413PairAxiomCompatible, x.Gate413PairNotNative, x.Gate413CKMCapacity, x.Gate413CoefficientsFree, x.Gate412KDiagonalOnly, x.Gate411AxiomLedgerCompiled, x.ChargedModuliDim, x.Verdict)
}
func FormatArena(x SelectorArena) string {
	return fmt.Sprintf("executed=%v K=%s X=%s family_dim=%d generated_alg_dim=%d noncommuting_capacity=%v coefficients_native=%v verdict=%s reason=%s", x.Executed, x.KName, x.ShiftObservableName, x.FamilyBasisDimension, x.GeneratedAlgebraDimension, x.NoncommutingCapacity, x.CoefficientsNative, x.Verdict, x.Reason)
}
func FormatFunctional(x FunctionalAudit) string {
	return fmt.Sprintf("name=%q executed=%v type=%q gauge_compatible=%v empirical_independent=%v unique_ray=%v selects_noncommuting=%v selects_sector_weights=%v selector_native=%v diagnostic=%s verdict=%s reason=%s", x.Name, x.Executed, x.FunctionalType, x.GaugeCompatible, x.EmpiricalIndependent, x.UniqueCoefficientRay, x.SelectsNoncommutingTexture, x.SelectsPhysicalSectorWeights, x.SelectorNative, FormatFloat(x.DiagnosticValue), x.Verdict, x.Reason)
}
func FormatConnection(x ConnectionAudit) string {
	return fmt.Sprintf("executed=%v ansatz=%q sample_curvature_norm=%s YM_minimizer_flat=%v flat_commutes=%v nonzero_curvature_needs_source=%v gauge_compatible_family_only=%v native_connection=%v coefficients_fixed=%v ckm_conditional=%v ckm_angle_predicted=%v verdict=%s reason=%s", x.Executed, x.ConnectionAnsatz, FormatFloat(x.FamilyCurvatureSampleNorm), x.YangMillsMinimizerFlat, x.FlatMinimizerCommutes, x.NonzeroCurvatureRequiresSource, x.GaugeCompatibilityIfFamilyOnly, x.ConnectionNativeInCurrentAsha, x.CoefficientsFixedByCurvature, x.CKMCapacityConditional, x.CKMAnglePredicted, x.Verdict, x.Reason)
}
func FormatCoefficients(x CoefficientImpact) string {
	return fmt.Sprintf("executed=%v sectors=%s coefficients_per_sector=%d total_free=%d topological_values=%d roots_fix=%v trace_fixes=%v curvature_fixes=%v sector_split_native=%v yukawa_imported=%v verdict=%s reason=%s", x.Executed, strings.Join(x.Sectors, ","), x.CoefficientsPerSector, x.TotalFreeTextureCoefficients, x.TopologicalCoefficientValuesFound, x.RootsOfUnityFixCoefficients, x.TraceFixesCoefficients, x.CurvatureFixesCoefficients, x.SectorSplittingNative, x.YukawaDataImported, x.Verdict, x.Reason)
}
func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("scenario=%q status=%s dim=%d masses3=%v ckm=%v pmns=%v coefficients_fixed=%v native_reduction=%v conditional=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.CoefficientsFixed, x.NativeReduction, x.ConditionalOnly, x.Reason)
}
func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d best_native_dim=%d native_reduction=%v conditional_mixing=%v coefficients_free=%v firewall=%v verdict=%s", x.StartDim, x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalMixingCapacity, x.CoefficientsRemainFree, x.FirewallPreserved, x.Verdict)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa_matrices=%v axiom_status=%v no_native_derivation=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaMatricesInserted, x.AxiomStatusPreserved, x.NoNativeDerivationClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%q reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 414 Registry Audit — Family Coefficient Selector / Constrained Connection Curvature Sieve\n\n")
	b.WriteString("## Claim tested\n\nGate 414 tests whether the Gate-413 noncommuting family-pair axiom can be upgraded from texture capacity into coefficient prediction by a native trace, curvature, finite-action, or constrained U(3)_gen connection rule. The gate does not import Yukawa matrices, observed masses, CKM, or PMNS data.\n\n")
	sections := []struct{ title, body string }{
		{"Prior boundary inherited", FormatInheritance(a.Inheritance)},
		{"Selector arena", FormatArena(a.Arena)},
		{"Constrained family connection", FormatConnection(a.Connection)},
		{"Coefficient impact", FormatCoefficients(a.Coefficients)},
		{"Empirical firewall", FormatFirewall(a.Firewall)},
	}
	for _, s := range sections {
		b.WriteString("## " + s.title + "\n\n```text\n" + s.body + "\n```\n\n")
	}
	b.WriteString("## Functional table\n\n```text\n")
	for _, f := range a.Functionals {
		b.WriteString(FormatFunctional(f) + "\n")
	}
	b.WriteString("```\n\n")
	b.WriteString("## Moduli impact\n\n```text\n")
	b.WriteString(FormatModuli(a.Moduli))
	for _, s := range a.Moduli.Scenarios {
		b.WriteString("\n" + FormatScenario(s))
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
