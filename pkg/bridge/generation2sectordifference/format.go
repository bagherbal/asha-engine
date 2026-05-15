package generation2sectordifference

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t inverse=%t provenance=%t branches=%t residual=%t multiplex=%t independent=%t native_universality_rejected=%t contamination_rejected=%t no_observed=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate456InverseDerived, x.Gate457ProvenanceContract, x.Gate459BranchTags, x.Gate460ResidualHarness, x.Gate461SectorMultiplex, x.Gate461IndependentSectorRaysAccepted, x.Gate461NativeUniversalityRejected, x.Gate461SectorContaminationRejected, x.NoObservedValuesImported, x.Verdict)
}

func FormatContract(x InterfaceContract) string {
	return fmt.Sprintf("executed=%t require_u=%t require_d=%t provenance=%t tags=%t eigenbasis=%t relative_dof=%d diagnostics_only=%t reject_ckm_native=%t reject_pmns_charged=%t reject_observed=%t reject_native_relative=%t verdict=%s reason=%s", x.Executed, x.RequiresUSector, x.RequiresDSector, x.RequiresProvenancePerSector, x.RequiresCompleteBranchTags, x.RequiresEigenbasisConvention, x.RelativeRayDimension, x.ExportsRelativeDiagnosticsOnly, x.RejectsCKMAsNativePrediction, x.RejectsPMNSInChargedCKMLedger, x.RejectsObservedMixingByDefault, x.RejectsNativeRelativeRayPromotion, x.Verdict, x.Reason)
}

func FormatSectorRay(x SectorRay) string {
	return fmt.Sprintf("sector=%s alpha=%.12g phi=%.12g sigma=%d sheet=%d provenance=%t branch_tag=%t eigenbasis=%t bridge=%t observed_CKM=%t observed_PMNS=%t native_claim=%t lepton_ledger=%t universality=%t native_universal=%t", x.Sector, x.Alpha, x.Phi, x.CPOddSign, x.C3Sheet, x.HasProvenance, x.HasBranchTag, x.HasEigenbasisConvention, x.BridgeOnly, x.ExplicitObservedCKM, x.ExplicitObservedPMNS, x.NativePromotionClaim, x.UsesLeptonPMNSLedger, x.UniversalityClaim, x.UniversalityNativeClaim)
}

func FormatRelativeRay(x RelativeRay) string {
	return fmt.Sprintf("%s->%s delta_alpha=%.12g delta_phi=%.12g phase_chord=%.12g distance=%.12g complete=%t eigenbasis=%t bridge=%t exports_CKM=%t exports_PMNS=%t native_export=%t verdict=%s reason=%s", x.FromSector, x.ToSector, x.DeltaAlpha, x.DeltaPhi, x.PhaseChord, x.ProjectiveDistance, x.CompleteInputs, x.EigenbasisConventionSet, x.BridgeOnly, x.ExportsCKMEntry, x.ExportsPMNSEntry, x.ExportsNativeObservable, x.Verdict, x.Reason)
}

func FormatSieve(x Sieve) string {
	return fmt.Sprintf("executed=%t accepted=%d rejected=%d valid_ud=%t missing_sector=%t missing_provenance=%t missing_eigenbasis=%t observed=%t native_prediction=%t native_relative=%t lepton_misroute=%t universality_native=%t all_bridge=%t no_native_mixing=%t verdict=%s reason=%s", x.Executed, x.AcceptedCaseCount, x.RejectedCaseCount, x.ValidUDDifferenceAccepted, x.MissingSectorRejected, x.MissingProvenanceRejected, x.MissingEigenbasisRejected, x.ObservedCKMPMNSRejected, x.NativePredictionRejected, x.NativeRelativePromotionRejected, x.LeptonPMNSMisrouteRejected, x.UniversalityNativeRejected, x.AllAcceptedBridgeOnly, x.NoNativeMixingObservableExport, x.Verdict, x.Reason)
}

func FormatFirewall(x CKMFirewall) string {
	return fmt.Sprintf("executed=%t ray_may_feed_CKM_adapter=%t CKM_computed=%t CKM_native=%t PMNS_computed=%t PMNS_native=%t requires_observed_import=%t provenance=%t eigenbasis=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t no_masses=%t no_yukawas=%t no_CKM=%t no_PMNS=%t no_GST=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.RelativeRayMayFeedCKMAdapter, x.CKMMatrixEntryComputed, x.CKMMatrixEntryNative, x.PMNSMatrixEntryComputed, x.PMNSMatrixEntryNative, x.RequiresObservedComparatorImport, x.RequiresSchemeScaleProvenance, x.RequiresEigenvectorGaugeConvention, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.NoObservedMassesImported, x.NoObservedYukawasImported, x.NoObservedCKMImported, x.NoObservedPMNSImported, x.NoGSTPromotion, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 462 Registry Audit — Sector-Difference Invariant / CKM Interface Firewall Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusCKMInterfaceFirewallValidated + "`\n\n")
	b.WriteString("Gate 462 isolates the relative u-d sector ray that a future CKM-facing bridge adapter may inspect. It does not compute CKM entries, PMNS entries, physical masses, Yukawa values, or native mixing predictions.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Interface contract\n\n")
	b.WriteString(FormatContract(a.Contract) + "\n\n")
	b.WriteString("```text\n")
	b.WriteString("required charged sectors: u,d\n")
	b.WriteString("relative ray: Delta_alpha = alpha_d-alpha_u\n")
	b.WriteString("relative phase: Delta_phi = wrap_pi(phi_d-phi_u)\n")
	b.WriteString("projective diagnostic: distance = sqrt(Delta_alpha^2 + (2 sin(Delta_phi/2))^2)\n")
	b.WriteString("forbidden export: CKM_ij, PMNS_ij, masses, Yukawas, GST/Fritzsch relations, native coefficient values\n")
	b.WriteString("```")
	b.WriteString("\n\n")

	b.WriteString("## Sieve\n\n")
	b.WriteString(FormatSieve(a.Sieve) + "\n\n")
	b.WriteString("| Case | Accepted | Verdict | Relative diagnostic | Reason |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, c := range a.Sieve.Cases {
		b.WriteString(fmt.Sprintf("| %s | %t | `%s` | %s | %s |\n", esc(c.Name), c.Accepted, esc(c.Verdict), esc(FormatRelativeRay(c.Relative)), esc(c.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Accepted relative-ray equation\n\n")
	b.WriteString("For the accepted synthetic u-d bridge row only:\n\n")
	b.WriteString("```text\n")
	b.WriteString("Delta_alpha_ud = alpha_d - alpha_u\n")
	b.WriteString("Delta_phi_ud   = wrap_pi(phi_d - phi_u)\n")
	b.WriteString("d_ud           = sqrt(Delta_alpha_ud^2 + 4 sin^2(Delta_phi_ud/2))\n")
	b.WriteString("```")
	b.WriteString("\n\n")
	b.WriteString("This is a comparator diagnostic. It is not a CKM matrix element and not a native ASHA observable.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n")

	b.WriteString("## Firewall\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")

	b.WriteString("## CKM/PMNS boundary\n\n")
	b.WriteString("A CKM-like adapter would need relative u-d diagonalization data plus explicit eigenbasis ordering and phase-gauge conventions. Gate 462 only constructs the relative-ray slot. PMNS requires a neutrino-sector ledger and is rejected in this charged-sector audit.\n\n")

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
