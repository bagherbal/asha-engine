package generation2branchresiduals

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t provenance=%t evaluator=%t branch_ledger=%t cp_sign=%t c3_sheet=%t unique=%t native_CP_absent=%t native_C3_absent=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseDerived, x.Gate457ProvenanceContractDefined, x.Gate458SyntheticHarnessDefined, x.Gate459BranchTagLedgerDefined, x.Gate459RequiresCPOddSign, x.Gate459RequiresC3Sheet, x.Gate459CompleteTagUnique, x.NativeCPSelectorAbsent, x.NativeC3SheetSelectorAbsent, x.NoObservedValuesImported, x.Verdict)
}

func FormatHarness(x Harness) string {
	return fmt.Sprintf("executed=%t inverse=%t evaluator=%t branch_tags=%t complete_tag=%t ray=%t R22=%t comparator_R=%t tag_R=%t synthetic=%t redacted=%t observed_rejected=%t bridge=%t verdict=%s reason=%s", x.Executed, x.ComposesGate456Inverse, x.ComposesGate458Evaluator, x.ComposesGate459BranchTags, x.RequiresCompleteBranchTag, x.ComputesProjectiveRay, x.ComputesTextureZeroResidual, x.ComputesComparatorResiduals, x.ComputesPhaseTagResiduals, x.SyntheticOnly, x.RedactedAllowedUnevaluated, x.ObservedDataRejected, x.BridgeOnlyExport, x.Verdict, x.Reason)
}

func FormatLedger(x ResidualLedger) string {
	return fmt.Sprintf("executed=%t matrix=%q gauge=%q R22=%q RK=%q Rspec=%q Rtag=%q bridge=%t native_observable=%t verdict=%s reason=%s", x.Executed, x.MatrixFormula, x.RayGauge, x.TextureZeroResidual, x.IKResidualFormula, x.ISpecResidualFormula, x.PhaseTagResidualFormula, x.ResidualsBridgeOnly, x.ResidualsNativeObservable, x.Verdict, x.Reason)
}

func FormatRequest(x ResidualRequest) string {
	return fmt.Sprintf("%s IK=%.12g Ispec=%.12g numeric=%t redacted=%t observed=%t sigma=%d hasSigma=%t sheet=%d hasSheet=%t bridge=%t native_claim=%t", x.Name, x.IK, x.ISpec, x.HasNumericPair, x.Redacted, x.ExplicitObservedData, x.CPOddSign, x.HasCPOddSign, x.C3Sheet, x.HasC3Sheet, x.BridgeOnly, x.NativePromotionClaim)
}

func FormatEvaluation(x ResidualEvaluation) string {
	return fmt.Sprintf("%s accepted=%t evaluated=%t redacted=%t alpha=%.12g C=%.12g phi=%.12g a=%.12g b=%.12g c=%.12g R22=%.3g RK=%.3g Rspec=%.3g Rcp=%d Rc3=%d domain=%t phase_domain=%t caustic=%t complete_tag=%t bridge=%t diagnostics=%t verdict=%s reason=%s", x.Request.Name, x.Accepted, x.Evaluated, x.Redacted, x.Alpha, x.Cos3Phi, x.Phi, x.A, x.B, x.C, x.M22Residual, x.IKResidual, x.ISpecResidual, x.CPSignResidual, x.C3SheetResidual, x.ProjectiveDomainOK, x.PhaseCosDomainOK, x.Caustic, x.CompleteBranchTag, x.BridgeOnlyExport, x.ResidualsAreDiagnostics, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d redacted=%t synthetic=%t incomplete=%t caustic=%t observed=%t native=%t projective=%t phase=%t all_bridge=%t diagnostics=%t no_native_obs=%t verdict=%s reason=%s", x.Executed, x.AcceptedCount, x.RejectedCount, x.RedactedPreserved, x.SyntheticInteriorAccepted, x.IncompleteTagRejected, x.CausticRejected, x.ObservedDataRejected, x.NativePromotionRejected, x.ProjectiveDomainRejected, x.PhaseCosDomainRejected, x.AllAcceptedBridgeOnly, x.AllResidualsDiagnosticOnly, x.NoNativeFlavorObservableOut, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_ray=%t no_phase_branch=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoCoefficientRayPromotion, x.NoPhaseBranchPromotion, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 460 Registry Audit — Branch-Resolved Texture Residual Harness / Synthetic Null Phenomenology Map\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusBridgeOnlyResidualExport + "`\n\n")
	b.WriteString("Gate 460 composes the symbolic coefficient-ray inverse, the redacted comparator evaluator, and the complete `{sigma_CP,n_C3}` branch tag into a branch-resolved residual harness. It evaluates only synthetic/null records and exports only bridge diagnostics.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Harness\n\n")
	b.WriteString(FormatHarness(a.Harness) + "\n\n")

	b.WriteString("## Residual ledger\n\n")
	b.WriteString(FormatLedger(a.Ledger) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha = sqrt(3) I_K / sqrt(1-I_K^2)\n")
	b.WriteString("C = cos(3phi) = (3sqrt(3)/2) I_spec / (1-I_K^2)^(3/2)\n")
	b.WriteString("phi = (sigma_CP arccos(C)+2*pi*n_C3)/3\n")
	b.WriteString("M = alpha K_gen + cos(phi) X_triangle + sin(phi) Y_phase\n")
	b.WriteString("R_22 = M_22 = 0\n")
	b.WriteString("R_K = I_K - alpha/sqrt(alpha^2+3)\n")
	b.WriteString("R_spec = I_spec - 2cos(3phi)/(alpha^2+3)^(3/2)\n")
	b.WriteString("```")
	b.WriteString("\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Record | Accepted | Evaluated | Redacted | alpha | C | phi | R22 | RK | Rspec | Rcp | Rc3 | Bridge-only | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|---|---|\n")
	for _, e := range a.Sieve.Evaluations {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %.8g | %.8g | %.8g | %.3g | %.3g | %.3g | %d | %d | %t | `%s` | %s |\n", esc(e.Request.Name), e.Accepted, e.Evaluated, e.Redacted, e.Alpha, e.Cos3Phi, e.Phi, e.M22Residual, e.IKResidual, e.ISpecResidual, e.CPSignResidual, e.C3SheetResidual, e.BridgeOnlyExport, esc(e.Verdict), esc(e.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## Native boundary\n\n")
	b.WriteString("The harness proves only that labelled bridge records are internally consistent with the ASHA structural texture. It does not create a native mass observable, a Yukawa coefficient, a CKM/PMNS phase, or a GST/Fritzsch relation. The selected phase branch is metadata, not law-space.\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
