package generation2empiricaladapter

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t texture_sum_rule=%t full_triangle=%t nn_not_gauge=%t gate453_interface=%t ray_dof=%d spectrum_rank=%d min_local=%d cp_branch=%t native_selector_absent=%t no_empirical=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate450TextureZeroSumRule, x.Gate451FullTrianglePreserved, x.Gate452NearestNeighborNotGauge, x.Gate453EmpiricalInterfaceDefined, x.Gate454ProjectiveRayDOF, x.Gate454SpectrumOnlyRank, x.Gate454MinimumLocalScalars, x.Gate454CPBranchTagRequired, x.Gate454NativeSelectorAbsent, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatSchema(x AdapterSchema) string {
	return fmt.Sprintf("executed=%t name=%s default_value_mode=%s labels=%s allowed_ops=%s rejected_ops=%s observed_default=%t native_coeff_export=%t gst_native=%t ckm_pmns_native=%t verdict=%s reason=%s", x.Executed, x.Name, x.DefaultValueMode, join(x.RequiredLabels), join(x.AllowedOperations), join(x.RejectedOperations), x.AllowsObservedValuesByDefault, x.AllowsNativeCoefficientExport, x.AllowsGSTAsNativeLaw, x.AllowsCKMPMNSAsNativeSelectors, x.Verdict, x.Reason)
}

func FormatRequest(x AdapterRequest) string {
	return fmt.Sprintf("%s: op=%s value_mode=%s comparators=%d labels={sector:%t scale:%t scheme:%t bridge:%t cp_branch:%t} inputs={spectrum:%t masses:%t yukawa:%t ckm:%t pmns:%t} claims={local:%t oriented:%t native_coeff:%t gst_native:%t phase_native:%t} allowed=%t class=%s verdict=%s reason=%s", x.Name, x.Operation, x.ValueMode, x.ComparatorCount, x.HasSectorTag, x.HasRenormalizationScale, x.HasRenormalizationScheme, x.HasBridgeLabel, x.HasCPBranchTag, x.UsesSpectrum, x.UsesMasses, x.UsesYukawa, x.UsesCKM, x.UsesPMNS, x.ClaimsLocalRay, x.ClaimsOrientedRay, x.ClaimsNativeCoefficient, x.ClaimsGSTNative, x.ClaimsPhaseNative, x.Allowed, x.Classification, x.Verdict, x.Reason)
}

func FormatSieve(x AdapterSieve) string {
	return fmt.Sprintf("executed=%t allowed=%d rejected=%d native_ledger=%t local_dry_run=%t oriented_dry_run=%t spectrum_native_rejected=%t missing_metadata_rejected=%t gst_rejected=%t ckm_pmns_rejected=%t observed_rejected=%t native_coeff_rejected=%t forbidden_accepted=%t verdict=%s reason=%s", x.Executed, x.AllowedCount, x.RejectedCount, x.NativeLedgerAllowed, x.LocalRayDryRunAllowed, x.OrientedRayDryRunAllowed, x.SpectrumOnlyNativePromotionRejected, x.MissingMetadataRejected, x.GSTNativePromotionRejected, x.CKMPMNSNativeSelectorRejected, x.ObservedValuesRejectedByDefault, x.NativeCoefficientExportRejected, x.AnyForbiddenAccepted, x.Verdict, x.Reason)
}

func FormatExport(x DryRunExport) string {
	return fmt.Sprintf("executed=%t observed_values=%d dummy_comparators=%d native_exports=%d bridge_exports=%d native_promotion_blocked=%t schema_failures_fail_closed=%t verdict=%s reason=%s", x.Executed, x.ActualObservedValueCount, x.DummyComparatorCount, x.NativeExportCount, x.BridgeExportCount, x.NativePromotionBlocked, x.SchemaFailuresFailClosed, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_curvefit=%t no_GST=%t no_native_ray=%t K=%t triangle=%t texture_sum_rule_bridge=%t Y_sealed=%t coeffs_sealed=%t cp_branch_tagged=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFitPromoted, x.NoGSTPromotion, x.NoNativeCoefficientRayValue, x.KGenStillForced, x.XTriangleStillForced, x.TextureZeroSumRuleStillBridge, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.CPOrientationStillBranchTagged, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 455 Registry Audit — Empirical Texture Adapter Stub / Dry-Run Firewall Test\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusDryRunFirewallValidated + "`\n\n")
	b.WriteString("Gate 455 does not import measured masses, CKM angles, PMNS angles, Yukawa values, or fitted coefficient rays. It defines and executes a dry-run adapter firewall: native structural ledgers and labelled symbolic bridge comparators are accepted, while all native-promotion routes fail closed.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Adapter schema\n\n")
	b.WriteString(FormatSchema(a.Schema) + "\n\n")
	b.WriteString("Required labels:\n\n")
	for _, x := range a.Schema.RequiredLabels {
		b.WriteString("- " + esc(x) + "\n")
	}
	b.WriteString("\n")

	b.WriteString("## Dry-run request sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Request | Operation | Value mode | Comparators | Allowed | Classification | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---:|---|---|---|---|\n")
	for _, r := range a.Sieve.Requests {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %d | %t | %s | `%s` | %s |\n", esc(r.Name), esc(r.Operation), esc(r.ValueMode), r.ComparatorCount, r.Allowed, esc(r.Classification), esc(r.Verdict), esc(r.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Adapter invariants\n\n")
	b.WriteString("The adapter enforces the Gate-454 rank boundary:\n\n")
	b.WriteString("```text\n")
	b.WriteString("spectrum only: rank = 1, residual coefficient-ray DOF = 1\n")
	b.WriteString("local dry run: {I_spec, I_K}, rank = 2, bridge-only\n")
	b.WriteString("oriented dry run: {I_spec, I_K, CP branch tag}, bridge-only\n")
	b.WriteString("native coefficient export: forbidden\n")
	b.WriteString("observed values in default dry-run mode: forbidden\n")
	b.WriteString("```\n\n")

	b.WriteString("## Dry-run export\n\n")
	b.WriteString(FormatExport(a.Export) + "\n\n")

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
