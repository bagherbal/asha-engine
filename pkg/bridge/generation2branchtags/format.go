package generation2branchtags

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t branches=%d provenance=%t harness=%t observed_rejected=%t bridge_only=%t native_CP_selector_absent=%t native_C3_selector_absent=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseDerived, x.Gate456GenericBranchCount, x.Gate457ProvenanceContractDefined, x.Gate458RedactedHarnessDefined, x.Gate458ObservedValuesRejected, x.Gate458BridgeOnly, x.NativeCPSelectorAbsent, x.NativeC3SheetSelectorAbsent, x.NoObservedValuesImported, x.Verdict)
}

func FormatLedger(x BranchLedger) string {
	return fmt.Sprintf("executed=%t cosine=%t cp_sign=%t c3_sheet=%t cosine_branches=%d cp_sign_branches=%d complete_branches=%d bridge_only=%t reject_CKM_PMNS=%t reject_native=%t native_CP_absent=%t native_C3_absent=%t verdict=%s reason=%s", x.Executed, x.RequiresCosineInvariant, x.RequiresCPOddSign, x.RequiresC3Sheet, x.CosineOnlyBranchCount, x.CPOddSignOnlyBranchCount, x.CompleteBranchTagCount, x.BridgeOnly, x.RejectsCKMOrPMNSAsSelector, x.RejectsNativePromotion, x.NativeCPOddSignSelectorAbsent, x.NativeC3SheetSelectorAbsent, x.Verdict, x.Reason)
}

func FormatRequest(x BranchRequest) string {
	return fmt.Sprintf("%s C=%.12g hasC=%t sigma=%d hasSigma=%t sheet=%d hasSheet=%t CKM_PMNS=%t bridge_only=%t native_claim=%t", x.Name, x.Cos3Phi, x.HasCosineInvariant, x.CPOddSign, x.HasCPOddSign, x.C3Sheet, x.HasC3Sheet, x.UsesCKMOrPMNS, x.BridgeOnly, x.NativePromotionClaim)
}

func FormatEvaluation(x BranchEvaluation) string {
	return fmt.Sprintf("%s accepted=%t selected=%t domain=%t phase=%.12g branches=%d bridge_only=%t native_blocked=%t verdict=%s reason=%s", x.Request.Name, x.Accepted, x.Selected, x.DomainOK, x.Phase, x.BranchCount, x.BridgeOnlyExport, x.NativePromotionBlocked, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d cosine_only=%t cp_only=%t complete_pos=%t complete_neg=%t CKM_PMNS=%t native_promotion=%t invalid_tag=%t all_bridge=%t no_native_phase=%t verdict=%s reason=%s", x.Executed, x.AcceptedCount, x.RejectedCount, x.CosineOnlyFlagged, x.CPOddOnlyFlagged, x.CompletePositiveAccepted, x.CompleteNegativeAccepted, x.CKMPMNSSelectorRejected, x.NativePromotionRejected, x.InvalidTagRejected, x.AllAcceptedBridgeOnly, x.NoNativePhaseExport, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_ray=%t no_CP=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoCoefficientRayPromotion, x.NoCPPhasePromotion, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 459 Registry Audit — Oriented Comparator Branch Tag Sieve / CP-Sign Ledger\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusBridgeOnlyBranchTagValidated + "`\n\n")
	b.WriteString("Gate 459 formalizes the orientation metadata required after the Gate 456 inverse and Gate 458 redacted evaluator. It proves that a cosine invariant alone gives six branches, a CP-odd sign gives three, and a complete bridge tag `{sigma_CP,n_C3}` gives one synthetic phase branch. None of these tags is native ASHA law.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Branch ledger\n\n")
	b.WriteString(FormatLedger(a.Ledger) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("C = cos(3 phi)\n")
	b.WriteString("cosine-only branches: phi = (± arccos(C) + 2*pi*n)/3, n=0,1,2\n")
	b.WriteString("sigma_CP = sign(sin(3 phi)) chooses the ± orientation\n")
	b.WriteString("n_C3 in {0,1,2} chooses the residual cubic sheet\n")
	b.WriteString("complete tag: phi = (sigma_CP arccos(C) + 2*pi*n_C3)/3\n")
	b.WriteString("```\n\n")

	b.WriteString("## Branch-tag sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Record | C=cos(3phi) | sigma_CP | n_C3 | Accepted | Selected | Phase | Branches | Bridge-only | Verdict | Reason |\n")
	b.WriteString("|---|---:|---:|---:|---|---|---:|---:|---|---|---|\n")
	for _, e := range a.Sieve.Evaluations {
		b.WriteString(fmt.Sprintf("| %s | %.8g | %s | %s | %t | %t | %.8g | %d | %t | `%s` | %s |\n", esc(e.Request.Name), e.Request.Cos3Phi, maybeInt(e.Request.CPOddSign, e.Request.HasCPOddSign), maybeInt(e.Request.C3Sheet, e.Request.HasC3Sheet), e.Accepted, e.Selected, e.Phase, e.BranchCount, e.BridgeOnlyExport, esc(e.Verdict), esc(e.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## Native no-go boundary\n\n")
	b.WriteString("- `" + StatusFailedNativeCPSelectorAbsent + "`: no native law in the current atlas selects `sign(sin(3phi))`.\n")
	b.WriteString("- `" + StatusFailedNativeC3SelectorAbsent + "`: no native law in the current atlas selects `n_C3 in {0,1,2}`.\n")
	b.WriteString("- CKM/PMNS phases are rejected as hidden branch selectors in this gate.\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func maybeInt(v int, ok bool) string {
	if !ok {
		return "∅"
	}
	return fmt.Sprintf("%d", v)
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
