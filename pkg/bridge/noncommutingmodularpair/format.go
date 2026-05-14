package noncommutingmodularpair

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%v gate412_K_compatible=%v K_not_native=%v K_diagonal_only=%v no_CKM_PMNS=%v gate411_ledger=%v gate409_trivial=%v gate408_scalar_blind=%v charged_moduli=%d verdict=%s", x.Executed, x.Gate412KGenAxiomCompatible, x.Gate412KGenNotNative, x.Gate412DiagonalOnly, x.Gate412NoCKMPMNS, x.Gate411AxiomLedgerCompiled, x.Gate409FermionCarrierTrivial, x.Gate408ScalarFlavorBlind, x.ChargedModuliDim, x.Verdict)
}
func FormatOperator(x OperatorAxiom) string {
	return fmt.Sprintf("executed=%v K=%s shift=%s order=%d orthogonal=%v native_shift=%v explicit_axiom=%v family_only=%v comm_KS=%s comm_KX=%s noncommuting=%v verdict=%s reason=%s", x.Executed, x.KName, x.ShiftName, x.ShiftOrder, x.ShiftOrthogonal, x.ShiftNativeInCurrentAsha, x.ExplicitAxiom, x.ActsOnlyOnFamilyFiber, FormatFloat(x.KShiftCommutatorNorm), FormatFloat(x.KXCommutatorNorm), x.Noncommuting, x.Verdict, x.Reason)
}
func FormatWeyl(x WeylAudit) string {
	return fmt.Sprintf("executed=%v omega=(%s,%si) clock_order=%d shift_order=%d weyl_residual=%s roots_fingerprint=%v roots_fix_angles=%v verdict=%s reason=%s", x.Executed, FormatFloat(x.OmegaReal), FormatFloat(x.OmegaImag), x.ClockOrder, x.ShiftOrder, FormatFloat(x.WeylRelationResidual), x.RootsOfUnityFingerprint, x.RootsFixPhysicalAngles, x.Verdict, x.Reason)
}
func FormatCompatibility(x CompatibilityAudit) string {
	return fmt.Sprintf("executed=%v family_only=%v commutes_AF=%v commutes_gauge=%v commutes_Y=%v commutes_SU2L=%v commutes_BL=%v Gamma=%v J_mirrored=%v first_order_if_DF_broadcast=%v requires_connection_axiom=%v verdict=%s reason=%s", x.Executed, x.ActsOnlyOnFamilyFiber, x.CommutesWithAF, x.CommutesWithGaugeCharges, x.CommutesWithHypercharge, x.CommutesWithSU2L, x.CommutesWithBL, x.CompatibleWithGamma, x.JCompatibleIfShiftMirrored, x.FirstOrderUnaffectedIfDFBroadcast, x.RequiresFamilyConnectionAxiom, x.Verdict, x.Reason)
}
func FormatTexture(x TextureCapacity) string {
	return fmt.Sprintf("executed=%v native_pairs=%d conditional_pairs=%d comm_KX=%s sample_up_down_comm=%s generated_alg_dim=%d full_M3_capacity=%v ckm_native=%v pmns_native=%v ckm_conditional=%v pmns_conditional=%v coefficients_fixed=%v coefficients_free=%v verdict=%s reason=%s", x.Executed, x.NativeNoncommutingPairs, x.ConditionalNoncommutingPairs, FormatFloat(x.KXCommutatorNorm), FormatFloat(x.SampleUpDownCommutatorNorm), x.GeneratedAlgebraDimension, x.FullM3CapacityConditional, x.CKMNative, x.PMNSNative, x.CKMConditional, x.PMNSConditional, x.CoefficientsFixedTopologically, x.CoefficientsRemainFree, x.Verdict, x.Reason)
}
func FormatScenario(x ModuliScenario) string {
	return fmt.Sprintf("scenario=%q status=%s dim=%d masses3=%v ckm=%v pmns=%v native_reduction=%v conditional=%v reason=%s", x.Name, x.Status, x.ModuliDim, x.ThreeDistinctMassesPossible, x.CKMPossible, x.PMNSPossible, x.NativeReduction, x.ConditionalOnly, x.Reason)
}
func FormatModuli(x ModuliImpact) string {
	return fmt.Sprintf("start_dim=%d best_native_dim=%d native_reduction=%v conditional_ckm_pmns=%v coefficients_free=%v firewall=%v verdict=%s", x.StartDim, x.BestNativeDim, x.NativeReductionBelow13, x.ConditionalCKMPMNSCapacity, x.CoefficientsFree, x.FirewallPreserved, x.Verdict)
}
func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%v no_masses=%v no_ckm=%v no_pmns=%v no_yukawa_matrices=%v pair_axiom_only=%v no_native_derivation=%v verdict=%s", x.Executed, x.NoObservedMassesImported, x.NoCKMImported, x.NoPMNSImported, x.NoYukawaMatricesInserted, x.PairPromotedAsAxiomOnly, x.NoNativeDerivationClaimed, x.Verdict)
}
func FormatNext(x NextStep) string {
	return fmt.Sprintf("gate=%d title=%q reason=%s primary_task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderMarkdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 413 Registry Audit — Second Family Operator / Noncommuting Modular Pair Axiom Sieve\n\n")
	b.WriteString("## Claim tested\n\nGate 413 tests whether adding the smallest complementary family-shift operator to the Gate-412 modular Hamiltonian activates CKM/PMNS-capable noncommuting texture capacity without empirical Yukawa data. The construction is audited as an explicit axiom, not as a native ASHA theorem.\n\n")
	sections := []struct{ title, body string }{
		{"Prior boundary inherited", FormatInheritance(a.Inheritance)},
		{"Complementary operator axiom", FormatOperator(a.Operator)},
		{"Weyl clock/shift fingerprint", FormatWeyl(a.Weyl)},
		{"Compatibility audit", FormatCompatibility(a.Compatibility)},
		{"Texture / mixing capacity", FormatTexture(a.Texture)},
		{"Empirical firewall", FormatFirewall(a.Firewall)},
	}
	for _, s := range sections {
		b.WriteString("## " + s.title + "\n\n```text\n" + s.body + "\n```\n\n")
	}
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
