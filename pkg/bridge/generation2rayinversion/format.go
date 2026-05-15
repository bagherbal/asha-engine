package generation2rayinversion

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t texture_sum_rule=%t nn_not_gauge=%t ray_dof=%d min_comparators=%d adapter=%t observed_rejected=%t native_promotion_rejected=%t metadata_required=%t no_empirical=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate450TextureZeroSumRule, x.Gate452NearestNeighborNotGauge, x.Gate454RayDOF, x.Gate454MinimumComparators, x.Gate455AdapterFirewallValidated, x.Gate455ObservedValuesRejected, x.Gate455NativePromotionRejected, x.Gate455RequiresMetadata, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatComparators(x ComparatorPair) string {
	return fmt.Sprintf("executed=%t names=%s I_K=%s I_spec=%s domain=%s rank=%d ray_dof=%d local=%t global=%t verdict=%s reason=%s", x.Executed, join(x.Names), x.IKFormula, x.ISpecFormula, x.Domain, x.LocalRank, x.ProjectiveRayDOF, x.SufficientLocally, x.SufficientGlobally, x.Verdict, x.Reason)
}

func FormatInverse(x InverseMap) string {
	return fmt.Sprintf("executed=%t alpha=%s cos3phi=%s phi_branches=%s abs_IK=%t cos_bound=%t generic_branches=%d bridge_only=%t native_export=%t verdict=%s reason=%s", x.Executed, x.AlphaFormula, x.CosThreePhiFormula, x.PhiBranchFormula, x.RequiresAbsIKLessThanOne, x.RequiresCosBound, x.BranchCountGeneric, x.BridgeOnly, x.ExportsNativeRay, x.Verdict, x.Reason)
}

func FormatDomain(x DomainBoundary) string {
	return fmt.Sprintf("executed=%t IK_interval=%s Ispec_bound=%s IK_boundary=%t cos_boundary=%t outside_rejected=%t jacobian=%s caustic=%s verdict=%s reason=%s", x.Executed, x.IKOpenInterval, x.ISpecBoundFormula, x.BoundaryIKUnit, x.BoundaryCosThreePhiUnit, x.BoundaryOutsideRejected, x.JacobianFormula, x.CausticFormula, x.Verdict, x.Reason)
}

func FormatSample(x Sample) string {
	return fmt.Sprintf("%s: I_K=%.6g I_spec=%.6g alpha=%.6g cos3phi=%.6g inside=%t caustic=%t branches=%d orient_without_tag=%t bridge=%t native=%t verdict=%s reason=%s", x.Name, x.IK, x.ISpec, x.Alpha, x.CosThreePhi, x.InsideDomain, x.AtCaustic, x.GenericBranchCount, x.CanOrientWithoutTag, x.AllowedAsBridgeDryRun, x.AllowedAsNativeExport, x.Verdict, x.Reason)
}

func FormatSieve(x BranchSieve) string {
	return fmt.Sprintf("executed=%t valid=%d rejected=%d generic_exists=%t caustic_exists=%t outside_rejected=%t no_orient_without_tag=%t no_native_export=%t global_unique_absent=%t caustic_branch_tag=%t verdict=%s reason=%s", x.Executed, x.ValidDomainCount, x.RejectedDomainCount, x.GenericBranchSampleExists, x.CausticSampleExists, x.OutsideDomainRejected, x.NoSampleCanOrientWithoutTag, x.NoSampleAllowedAsNativeExport, x.GlobalUniqueCoefficientRayAbsent, x.ExplicitBranchTagRequiredAtCaustics, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_native_ray=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoNativeCoefficientRayValue, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 456 Registry Audit — Symbolic Coefficient-Ray Inversion / Branch-Caustic Map\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusBridgeOnlyInversionValidated + "`\n\n")
	b.WriteString("Gate 456 derives the exact symbolic inverse from labelled bridge comparators to the projective coefficient ray. The derivation is bridge-only: it does not import masses, Yukawa values, CKM/PMNS data, GST/Fritzsch assumptions, or fitted coefficient values.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Comparator pair\n\n")
	b.WriteString(FormatComparators(a.Comparators) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Comparators.IKFormula + "\n")
	b.WriteString(a.Comparators.ISpecFormula + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Symbolic inverse\n\n")
	b.WriteString(FormatInverse(a.Inverse) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Inverse.AlphaFormula + "\n")
	b.WriteString(a.Inverse.CosThreePhiFormula + "\n")
	b.WriteString(a.Inverse.PhiBranchFormula + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Domain and caustics\n\n")
	b.WriteString(FormatDomain(a.Domain) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Domain.IKOpenInterval + "\n")
	b.WriteString(a.Domain.ISpecBoundFormula + "\n")
	b.WriteString(a.Domain.JacobianFormula + "\n")
	b.WriteString(a.Domain.CausticFormula + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Dry-run branch sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Sample | I_K | I_spec | alpha | cos(3phi) | In domain | Caustic | Branches | Bridge dry run | Native export | Verdict | Reason |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---|---:|---|---|---|---|\n")
	for _, s := range a.Sieve.Samples {
		b.WriteString(fmt.Sprintf("| %s | %.6g | %.6g | %.6g | %.6g | %t | %t | %d | %t | %t | `%s` | %s |\n", esc(s.Name), s.IK, s.ISpec, s.Alpha, s.CosThreePhi, s.InsideDomain, s.AtCaustic, s.GenericBranchCount, s.AllowedAsBridgeDryRun, s.AllowedAsNativeExport, esc(s.Verdict), esc(s.Reason)))
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
