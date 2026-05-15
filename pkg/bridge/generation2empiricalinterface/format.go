package generation2empiricalinterface

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate444_K=%t gen2_zero=%t triangle=%t phase_sealed=%t coeffs_sealed=%t texture_sum_rule=%t ratios_require_amplitudes=%t gate451_full_triangle=%t no_phase_selector=%t gate452_not_gauge=%t basis_group=%s no_empirical=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate444Generation2Zero, x.Gate445TriangleForced, x.Gate446PhaseQuarantined, x.Gate447CoefficientsSealed, x.Gate450TextureZeroSumRule, x.Gate450RatiosRequireAmplitudes, x.Gate451FullTrianglePreserved, x.Gate451NoNativePhaseRaySelector, x.Gate452NearestNeighborNotGauge, x.Gate452KPreservingBasisGroup, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatNativeInvariant(x NativeInvariant) string {
	return fmt.Sprintf("%s: formula=%s status=%s depends_on=%s requires_empirical=%t predicts_number=%t reason=%s", x.Name, x.Formula, x.NativeStatus, join(x.DependsOn), x.RequiresEmpirical, x.CanPredictNumber, x.Reason)
}

func FormatNativeLedger(x NativeLedger) string {
	return fmt.Sprintf("executed=%t promoted=%s quarantined=%s native_only_predicts_GST=%t verdict=%s reason=%s", x.Executed, join(x.PromotedStructuralObjects), join(x.QuarantinedObjects), x.NativeOnlyPredictsGST, x.Verdict, x.Reason)
}

func FormatEmpiricalInput(x EmpiricalInput) string {
	return fmt.Sprintf("%s: kind=%s allowed=%t required_label=%s native_promotion=%t reason=%s", x.Name, x.Kind, x.Allowed, x.RequiredLabel, x.NativePromotion, x.Reason)
}

func FormatImportContract(x ImportContract) string {
	return fmt.Sprintf("executed=%t allowed_inputs=%d rejected_promotions=%d explicit_label=%t renormalization_tag=%t sector_tag=%t allows_native_claim=%t verdict=%s reason=%s", x.Executed, x.AllowedCount, x.RejectedPromotionCount, x.RequiresExplicitLabel, x.RequiresRenormalizationTag, x.RequiresSectorTag, x.AllowsNativeClaim, x.Verdict, x.Reason)
}

func FormatComparator(x Comparator) string {
	return fmt.Sprintf("%s: formula=%s inputs=%s native_observable=%t allowed=%t reason=%s", x.Name, x.Formula, join(x.RequiresInputs), x.NativeObservable, x.Allowed, x.Reason)
}

func FormatResidualLedger(x ResidualLedger) string {
	return fmt.Sprintf("executed=%t texture_residuals=%t GST_residual=%t native_GST_claim=%t coefficient_fit_native=%t verdict=%s reason=%s", x.Executed, x.AllowsTextureResiduals, x.AllowsGSTResidual, x.AllowsNativeGSTRatioClaim, x.AllowsCoefficientFittingAsNative, x.Verdict, x.Reason)
}

func FormatInterfaceRequest(x InterfaceRequest) string {
	return fmt.Sprintf("%s: operation=%s imports_empirical=%t labelled=%t attempts_promotion=%t allowed=%t reason=%s", x.Name, x.RequestedOperation, x.ImportsEmpirical, x.ExplicitlyLabelled, x.AttemptsPromotion, x.Allowed, x.Reason)
}

func FormatInterfaceSieve(x InterfaceSieve) string {
	return fmt.Sprintf("executed=%t native_only_allowed=%t empirical_fit_allowed=%t promotion_rejected=%t forbidden_accepted=%t verdict=%s reason=%s", x.Executed, x.NativeOnlyAllowed, x.EmpiricalFitAllowed, x.PromotionRejected, x.AnyForbiddenAccepted, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_curve_fit=%t no_GST_promotion=%t K_forced=%t gen2_zero=%t triangle=%t Y_phase_sealed=%t coeffs_sealed=%t GST_quarantined=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFit, x.NoGSTPromotion, x.KGenStillForced, x.Generation2ZeroStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.GSTFritzschRelationsQuarantined, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 453 Registry Audit — Texture-Zero Invariant Ledger / Allowed Empirical Interface\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusEmpiricalInterfaceDefined + "`\n\n")
	b.WriteString("Gate 453 is not a new native flavor prediction. It is the boundary contract that prevents the exact texture-zero sum rule from being misused as a hidden GST/Fritzsch derivation. The gate accepts native structural ledgers and explicitly labelled empirical comparator use, while rejecting any silent promotion of observed masses, CKM/PMNS data, coefficient fits, or nearest-neighbor assumptions into native ASHA geometry.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Native invariant ledger\n\n")
	b.WriteString(FormatNativeLedger(a.Native) + "\n\n")
	b.WriteString("| Invariant | Formula | Native status | Requires empirical input? | Predicts number? | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, inv := range a.Native.Invariants {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %t | %t | %s |\n", esc(inv.Name), esc(inv.Formula), esc(inv.NativeStatus), inv.RequiresEmpirical, inv.CanPredictNumber, esc(inv.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("Promoted structural objects: `" + esc(join(a.Native.PromotedStructuralObjects)) + "`.\n\n")
	b.WriteString("Quarantined objects: `" + esc(join(a.Native.QuarantinedObjects)) + "`.\n\n")

	b.WriteString("## Empirical import contract\n\n")
	b.WriteString(FormatImportContract(a.Contract) + "\n\n")
	b.WriteString("| Input | Kind | Allowed? | Required label | Native promotion? | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, in := range a.Contract.Inputs {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | `%s` | %t | %s |\n", esc(in.Name), esc(in.Kind), in.Allowed, esc(in.RequiredLabel), in.NativePromotion, esc(in.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("Every empirical bridge input must carry a sector tag and a renormalization scale/scheme tag. Without those tags, the interface must reject the request.\n\n")

	b.WriteString("## Comparator and residual ledger\n\n")
	b.WriteString(FormatResidualLedger(a.Residuals) + "\n\n")
	b.WriteString("| Comparator | Formula | Required inputs | Allowed? | Reason |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, c := range a.Residuals.Comparators {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %t | %s |\n", esc(c.Name), esc(c.Formula), esc(join(c.RequiresInputs)), c.Allowed, esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Interface sieve\n\n")
	b.WriteString(FormatInterfaceSieve(a.Sieve) + "\n\n")
	b.WriteString("| Request | Operation | Imports empirical? | Labelled? | Attempts promotion? | Allowed? | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, r := range a.Sieve.Requests {
		b.WriteString(fmt.Sprintf("| %s | %s | %t | %t | %t | %t | %s |\n", esc(r.Name), esc(r.RequestedOperation), r.ImportsEmpirical, r.ExplicitlyLabelled, r.AttemptsPromotion, r.Allowed, esc(r.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	return s
}
