package generation2provenancecontract

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t bridge_only=%t branches=%d branch_tags=%t domain_guard=%t native_selector_absent=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456SymbolicInverseDerived, x.Gate456BridgeOnly, x.Gate456GenericBranchCount, x.Gate456RequiresBranchTags, x.Gate456ComparatorDomainGuard, x.NativeCoefficientSelectorAbsent, x.NoObservedValuesImported, x.Verdict)
}

func FormatRule(x FieldRule) string {
	return fmt.Sprintf("%s required=%t failure=%s reason=%s", x.Name, x.Required, x.FailureCode, x.Reason)
}

func FormatContract(x Contract) string {
	return fmt.Sprintf("executed=%t required_fields=%d sector=%t observable=%t scale=%t scheme=%t source=%t source_version=%t uncertainty=%t dimensionless=%t bridge_only=%t no_native_promotion=%t branch_tag_if_oriented=%t observed_explicit_bridge=%t verdict=%s reason=%s", x.Executed, x.RequiredFieldCount, x.RequiresSector, x.RequiresObservable, x.RequiresScale, x.RequiresScheme, x.RequiresSource, x.RequiresSourceVersion, x.RequiresUncertainty, x.RequiresDimensionless, x.RequiresBridgeOnly, x.RequiresNoNativePromotion, x.RequiresBranchTagIfOriented, x.AllowsObservedOnlyWithExplicitBridgeImport, x.Verdict, x.Reason)
}

func FormatRecord(x ComparatorRecord) string {
	return fmt.Sprintf("%s sector=%s observable=%s kind=%s value=%s scale=%s scheme=%s source=%s version=%s uncertainty=%s dimensionless=%t bridge_only=%t explicit_observed=%t native_claim=%t oriented=%t branch=%s passed=%t verdict=%s reason=%s", x.Name, x.Sector, x.Observable, x.ValueKind, x.ValueExpression, x.Scale, x.Scheme, x.Source, x.SourceVersion, x.Uncertainty, x.Dimensionless, x.BridgeOnly, x.ExplicitObservedImport, x.NativePromotionClaim, x.RequiresOrientedInverse, x.BranchTag, x.Passed, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d symbolic_ok=%t observed_schema_ok=%t missing_sector=%t missing_scale_scheme=%t missing_source_uncertainty=%t native_promotion=%t observed_default=%t branch_tag=%t dimensionful=%t no_native_export=%t verdict=%s reason=%s", x.Executed, x.AcceptedCount, x.RejectedCount, x.CompleteSymbolicDryRunAccepted, x.ExplicitBridgeObservedSchemaAccepted, x.MissingSectorRejected, x.MissingScaleSchemeRejected, x.MissingSourceUncertaintyRejected, x.NativePromotionRejected, x.ObservedDefaultRejected, x.BranchTagMissingRejected, x.DimensionfulComparatorRejected, x.NoAcceptedNativeExport, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_ray=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoCoefficientRayPromotion, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 457 Registry Audit — Empirical Comparator Provenance Contract / Sector-Scheme Ledger\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusTextureComparatorContractValidated + "`\n\n")
	b.WriteString("Gate 457 defines the fail-closed schema that any future empirical texture comparator must satisfy before reaching the Gate 456 symbolic inverse. It evaluates provenance only; it imports no observed flavor values and promotes no coefficient ray to native law-space.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Provenance contract\n\n")
	b.WriteString(FormatContract(a.Contract) + "\n\n")
	b.WriteString("| Field | Required | Failure code | Reason |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, r := range a.Contract.Rules {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s |\n", esc(r.Name), r.Required, esc(r.FailureCode), esc(r.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Contract sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Record | Sector | Observable | Kind | Scale | Scheme | Source | Uncertainty | Dimensionless | Bridge-only | Explicit observed import | Native claim | Branch tag | Passed | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|\n")
	for _, r := range a.Sieve.Records {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s | %s | %s | %t | %t | %t | %t | %s | %t | `%s` | %s |\n", esc(r.Name), esc(r.Sector), esc(r.Observable), esc(r.ValueKind), esc(r.Scale), esc(r.Scheme), esc(r.Source), esc(r.Uncertainty), r.Dimensionless, r.BridgeOnly, r.ExplicitObservedImport, r.NativePromotionClaim, esc(r.BranchTag), r.Passed, esc(r.Verdict), esc(r.Reason)))
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
	if s == "" {
		return "∅"
	}
	return s
}
