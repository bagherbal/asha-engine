package generation2comparatorevaluation

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t bridge_only=%t branches=%d provenance=%t fields=%d observed_explicit=%t native_selector_absent=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseDerived, x.Gate456BridgeOnly, x.Gate456GenericBranchCount, x.Gate457ProvenanceContractDefined, x.Gate457RequiredFields, x.Gate457ObservedImportExplicitOnly, x.NativeCoefficientRaySelectorAbsent, x.NoObservedValuesImported, x.Verdict)
}

func FormatHarness(x Harness) string {
	return fmt.Sprintf("executed=%t gate457_only=%t synthetic=%t redacted=%t observed_numeric_rejected=%t inverse=%t alpha=%t cos3phi=%t branches=%t domain_guards=%t bridge_only=%t verdict=%s reason=%s", x.Executed, x.AcceptsOnlyGate457ValidInput, x.SyntheticModeAllowed, x.RedactedModeAllowed, x.ObservedNumericModeRejected, x.UsesGate456Inverse, x.ComputesAlpha, x.ComputesCos3Phi, x.ComputesBranchDiagnostics, x.ComputesDomainGuards, x.BridgeOnlyOutput, x.Verdict, x.Reason)
}

func FormatInput(x ComparatorInput) string {
	return fmt.Sprintf("%s sector=%s pair=%s kind=%s IK=%.12g I_spec=%.12g numeric=%t observed=%t bridge_only=%t native_claim=%t branch=%s", x.Name, x.Sector, x.ObservablePair, x.ValueKind, x.IK, x.ISpec, x.HasNumericPair, x.ExplicitObservedImport, x.BridgeOnly, x.NativePromotionClaim, x.BranchTag)
}

func FormatEvaluation(x Evaluation) string {
	return fmt.Sprintf("%s evaluated=%t accepted=%t redacted=%t alpha=%.12g cos3phi=%.12g branches=%d caustic=%t domainIK=%t domainCos=%t bridge_only=%t native_blocked=%t verdict=%s reason=%s", x.Input.Name, x.Evaluated, x.Accepted, x.Redacted, x.Alpha, x.Cos3Phi, x.PhaseBranches, x.Caustic, x.ProjectiveDomainOK, x.PhaseCosDomainOK, x.BridgeOnlyExport, x.NativePromotionBlocked, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d redacted=%t interior=%t caustic=%t observed_rejected=%t IK_domain=%t cos_domain=%t native_promotion=%t all_bridge=%t no_native_ray=%t verdict=%s reason=%s", x.Executed, x.AcceptedCount, x.RejectedCount, x.RedactedAccepted, x.SyntheticInteriorAccepted, x.SyntheticCausticFlagged, x.ObservedValueRejected, x.IKDomainRejected, x.PhaseCosDomainRejected, x.NativePromotionRejected, x.AllAcceptedBridgeOnly, x.NoNativeRayExport, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_ray=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoCoefficientRayPromotion, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 458 Registry Audit — Comparator Ledger Evaluation Harness / Redacted Phenomenology Slot\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusBridgeOnlyExportValidated + "`\n\n")
	b.WriteString("Gate 458 is the first bridge evaluator behind the Gate 457 provenance contract. It evaluates only synthetic comparator pairs and redacted placeholders, applies the Gate 456 inverse where legal, and refuses observed numeric flavor data in this harness.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Harness\n\n")
	b.WriteString(FormatHarness(a.Harness) + "\n\n")
	b.WriteString("The evaluator uses the symbolic formulas:\n\n")
	b.WriteString("```text\n")
	b.WriteString("alpha = sqrt(3) I_K / sqrt(1 - I_K^2)\n")
	b.WriteString("cos(3 phi) = (3 sqrt(3)/2) I_spec / (1 - I_K^2)^(3/2)\n")
	b.WriteString("domain: |I_K| < 1 and |cos(3 phi)| <= 1\n")
	b.WriteString("caustic: sin(3 phi) = 0\n")
	b.WriteString("```\n\n")

	b.WriteString("## Evaluation sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Record | Sector | Kind | Numeric | I_K | I_spec | Accepted | Evaluated | Redacted | Alpha | cos(3phi) | Branches | Caustic | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---|---:|---:|---|---|---|---:|---:|---:|---|---|---|\n")
	for _, e := range a.Sieve.Evaluations {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %t | %.8g | %.8g | %t | %t | %t | %.8g | %.8g | %d | %t | `%s` | %s |\n", esc(e.Input.Name), esc(e.Input.Sector), esc(e.Input.ValueKind), e.Input.HasNumericPair, e.Input.IK, e.Input.ISpec, e.Accepted, e.Evaluated, e.Redacted, e.Alpha, e.Cos3Phi, e.PhaseBranches, e.Caustic, esc(e.Verdict), esc(e.Reason)))
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
