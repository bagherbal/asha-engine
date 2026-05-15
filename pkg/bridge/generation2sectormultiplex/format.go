package generation2sectormultiplex

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t provenance=%t evaluator=%t branch_tags=%t residual_harness=%t residual_bridge=%t observed_rejected=%t native_blocked=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseDerived, x.Gate457ProvenanceContract, x.Gate458ComparatorHarness, x.Gate459BranchTags, x.Gate460ResidualHarness, x.Gate460ResidualsBridgeOnly, x.Gate460ObservedRejected, x.Gate460NativePromotionBlocked, x.NoObservedValuesImported, x.Verdict)
}

func FormatContract(x MultiplexContract) string {
	return fmt.Sprintf("executed=%t sector_indexed=%t sectors=%s independent_ray=%t no_ray_sharing=%t no_phase_sharing=%t no_branch_tag_sharing=%t provenance=%t complete_tags=%t labelled_bridge_universality=%t native_universality_rejected=%t bridge=%t verdict=%s reason=%s", x.Executed, x.SectorIndexed, strings.Join(x.RequiredSectors, ","), x.IndependentRayPerSector, x.NoImplicitRaySharing, x.NoImplicitPhaseSharing, x.NoImplicitBranchTagSharing, x.RequiresProvenancePerRow, x.RequiresCompleteBranchTags, x.AllowsLabelledBridgeOnlyUniversality, x.RejectsNativeUniversality, x.BridgeOnlyExport, x.Verdict, x.Reason)
}

func FormatDimensions(x DimensionLedger) string {
	return fmt.Sprintf("executed=%t sectors=%s coeffs_per_sector=%d total_kxy=%d native_dim_before=%d native_dim_after=%d universality_reduces_bridge_dof=%t reduction_native=%t independent_rays_native=%t sector_universality_native=%t verdict=%s reason=%s", x.Executed, strings.Join(x.ChargedSectors, ","), x.CoefficientsPerSector, x.TotalKXYChargedCoefficients, x.NativeChargedFlavorDimBefore, x.NativeChargedFlavorDimAfter, x.UniversalityWouldReduceBridgeDOF, x.UniversalityReductionNative, x.IndependentSectorRaysNative, x.SectorRayUniversalityNative, x.Verdict, x.Reason)
}

func FormatRecord(x SectorRecord) string {
	return fmt.Sprintf("sector=%s IK=%.12g Ispec=%.12g sigma=%d sheet=%d numeric=%t sigma_tag=%t sheet_tag=%t provenance=%t bridge=%t observed=%t native_claim=%t shared_from=%s universal=%t labelled=%t native_universal=%t native_theorem=%t", x.Sector, x.IK, x.ISpec, x.CPOddSign, x.C3Sheet, x.HasNumericPair, x.HasCPOddSign, x.HasC3Sheet, x.HasProvenance, x.BridgeOnly, x.ExplicitObservedData, x.NativePromotionClaim, empty(x.SharedFromSector), x.UniversalityClaim, x.UniversalityLabelled, x.NativeUniversalityClaim, x.IndependentNativeTheorem)
}

func FormatEvaluation(x SectorEvaluation) string {
	return fmt.Sprintf("sector=%s accepted=%t evaluated=%t alpha=%.12g C=%.12g phi=%.12g tag=%t provenance=%t bridge=%t native_blocked=%t independent=%t shared_assumption=%t no_physical=%t verdict=%s reason=%s", x.Record.Sector, x.Accepted, x.Evaluated, x.Alpha, x.Cos3Phi, x.Phi, x.CompleteBranchTag, x.ProvenanceComplete, x.BridgeOnlyExport, x.NativePromotionBlocked, x.RayIndependent, x.RaySharedOnlyByAssumption, x.NoPhysicalObservableValue, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted_cases=%d rejected_cases=%d independent_three_sector=%t labelled_bridge_universality=%t missing=%t native_universal=%t unlabelled_universal=%t observed=%t native_promotion=%t contamination=%t all_bridge=%t no_native_obs=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.IndependentThreeSectorAccepted, x.LabelledBridgeUniversalityAccepted, x.MissingSectorRejected, x.NativeUniversalityRejected, x.UnlabelledUniversalityRejected, x.ObservedDataRejected, x.NativePromotionRejected, x.SectorContaminationRejected, x.AllAcceptedBridgeOnly, x.NoNativeObservableExport, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_top_bottom=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_GST=%t no_ray=%t no_sector_universality=%t no_curvefit=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedTopBottomImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoGSTPromotion, x.NoCoefficientRayPromotion, x.NoCrossSectorUniversalityLaw, x.NoCurveFitPromoted, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 461 Registry Audit — Three-Sector Comparator Multiplex / Universality Assumption Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusMultiplexBridgeOnlyValidated + "`\n\n")
	b.WriteString("Gate 461 lifts the Gate 460 branch-resolved residual harness from a single synthetic/null comparator row into the charged-sector ledger `{u,d,e}`. The audit proves that sector-indexed bridge rays may be evaluated independently, while cross-sector coefficient-ray universality is not native ASHA law.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Multiplex contract\n\n")
	b.WriteString(FormatContract(a.Contract) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("charged sectors = {u,d,e}\n")
	b.WriteString("sector row = {sector, I_K, I_spec, sigma_CP, n_C3, provenance, bridge_only}\n")
	b.WriteString("native: K_gen and X_triangle are shared structural geometry\n")
	b.WriteString("bridge: alpha_s, phi_s, sigma_s, n_s are sector-indexed unless an explicit bridge universality assumption is declared\n")
	b.WriteString("forbidden: alpha_u=alpha_d=alpha_e or phi_u=phi_d=phi_e as native law without an independent theorem\n")
	b.WriteString("```\n\n")

	b.WriteString("## Dimension ledger\n\n")
	b.WriteString(FormatDimensions(a.Dimensions) + "\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Sector evaluations\n\n")
	b.WriteString("| Case | Sector | Accepted | alpha | cos(3phi) | phi | Independent | Shared assumption | Verdict | Reason |\n")
	b.WriteString("|---|---|---|---:|---:|---:|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		for _, e := range a.Sieve.CaseEvaluations[c.Name] {
			b.WriteString(fmt.Sprintf("| %s | %s | %t | %.8g | %.8g | %.8g | %t | %t | `%s` | %s |\n", esc(c.Name), esc(e.Record.Sector), e.Accepted, e.Alpha, e.Cos3Phi, e.Phi, e.RayIndependent, e.RaySharedOnlyByAssumption, esc(e.Verdict), esc(e.Reason)))
		}
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
	b.WriteString("The shared ASHA geometry is the structural family axis and triangle support. The sector coefficient rays are not shared by native law. A universality hypothesis can be carried as a labelled bridge stress test, but it cannot reduce the 9 charged K/X/Y coefficients or the 13-moduli flavor firewall.\n\n")

	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")

	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}

func empty(s string) string {
	if s == "" {
		return "∅"
	}
	return s
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	if s == "" {
		return "∅"
	}
	return s
}
