package generation2basisinvariance

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t gate444_K=%t gen2_zero=%t triangle=%t phase_sealed=%t coeffs_sealed=%t texture_sum_rule=%t gate451_full_triangle=%t no_phase_selector=%t GST_quarantined=%t no_empirical=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate444Generation2Zero, x.Gate445TriangleForced, x.Gate446PhaseQuarantined, x.Gate447CoefficientsSealed, x.Gate450TextureZeroSumRule, x.Gate451FullTrianglePreserved, x.Gate451NoNativePhaseRaySelector, x.Gate451GSTFritzschQuarantined, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatBasisTransformation(x BasisTransformation) string {
	return fmt.Sprintf("%s: formula=%s preserves_K=%t preserves_K_orbit=%t preserves_gen2=%t preserves_zero22=%t preserves_magnitudes=%t preserves_support=%t can_delete_13=%t native_allowed=%t reason=%s", x.Name, x.Formula, x.PreservesKGen, x.PreservesKGenUpToSign, x.PreservesGeneration2Address, x.PreservesStructuralZero22, x.PreservesEdgeMagnitudes, x.PreservesSupport, x.CanDelete13Edge, x.AllowedNativeGauge, x.Reason)
}

func FormatBasisAudit(x BasisGroupAudit) string {
	return fmt.Sprintf("executed=%t exact_group=%s orbit_extension=%s general_U3_rejected=%t allowed_preserve_13=%t allowed_delete_13=%t verdict=%s reason=%s", x.Executed, x.ExactKPreservingGroup, x.KOrbitPreservingExtension, x.GeneralUnitaryRejected, x.AllNativeAllowedPreserve13, x.AnyNativeAllowedDeletes13, x.Verdict, x.Reason)
}

func FormatEdgeMagnitude(x EdgeMagnitude) string {
	return fmt.Sprintf("edge=%s DeltaK=%d triangle=%.0f nearest_neighbor=%.0f rephasing_invariant=%t deleted_in_NN=%t reason=%s", x.Edge, x.DeltaK, x.TriangleMagnitude, x.NearestNeighborMagnitude, x.RephasingInvariant, x.DeletedInNN, x.Reason)
}

func FormatSupportAudit(x SupportAudit) string {
	return fmt.Sprintf("executed=%t triangle=%s NN=%s edge_count=(%d,%d) cycles=(%d,%d) support_invariant=%t can_rephase_to_NN=%t verdict=%s reason=%s", x.Executed, x.TriangleSupportFormula, x.NearestNeighborFormula, x.EdgeCountTriangle, x.EdgeCountNearestNeighbor, x.TriangleCycleCount, x.NearestNeighborCycleCount, x.SupportPatternInvariant, x.CanRephaseToNN, x.Verdict, x.Reason)
}

func FormatGraphInvariant(x GraphInvariant) string {
	return fmt.Sprintf("%s: triangle=%s NN=%s equal=%t invariant=%t reason=%s", x.Name, x.Triangle, x.NearestNeighbor, x.Equal, x.BasisInvariant, x.Reason)
}

func FormatSpectralAudit(x SpectralAudit) string {
	return fmt.Sprintf("executed=%t spectra=(%s vs %s) det_coeff=(%d,%d) comm_norm2=(%d,%d) same_class=%t verdict=%s reason=%s", x.Executed, x.TriangleAdjacencySpectrum, x.NearestNeighborSpectrum, x.TriangleDetLiftCoeff, x.NearestNeighborDetLiftCoeff, x.TriangleCommutatorNorm2, x.NearestNeighborCommutatorNorm2, x.SameInvariantClass, x.Verdict, x.Reason)
}

func FormatGaugeArtifactVerdict(x GaugeArtifactVerdict) string {
	return fmt.Sprintf("executed=%t NN_native_gauge_artifact=%t requires_non_native_U3=%t K_address_destroyed=%t texture_zero_destroyed=%t GST_empirical=%t reevaluate_ratios=%t verdict=%s reason=%s", x.Executed, x.NearestNeighborCanBeNativeGaugeArtifact, x.RequiresNonNativeGeneralUnitary, x.KGenAddressDestroyed, x.TextureZeroAddressDestroyed, x.GSTFritzschStillEmpiricalAssumption, x.ReevaluateRatios, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_CKM=%t no_PMNS=%t no_curvefit=%t K_forced=%t gen2_zero=%t X_triangle=%t Y_quarantined=%t coeffs_sealed=%t GST_quarantined=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFit, x.KGenStillForced, x.Generation2ZeroStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.GSTFritzschRelationsQuarantined, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 452 Registry Audit — Family Basis-Invariance / Texture Gauge-Artifact Audit\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 452 closes the loophole left after Gate 451: perhaps the forced closed triangle is only a family-basis artifact, and a legitimate change of basis can reveal a nearest-neighbor Fritzsch/GST chain. The audit keeps `K_gen=diag(-1,0,1)` as the native family address. Under this rule, allowed basis freedoms are rephasings and, at most, orientation reversal of the unordered primitive spectrum. No observed quark, lepton, CKM, PMNS, or Yukawa data is imported.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Native basis group audit\n\n")
	b.WriteString(FormatBasisAudit(a.BasisAudit) + "\n\n")
	b.WriteString("| Transformation | Native allowed? | Preserves K address? | Preserves support? | Can delete 1-3? | Reason |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, t := range a.BasisAudit.Transformations {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %s |\n", esc(t.Name), t.AllowedNativeGauge, t.PreservesKGen || t.PreservesKGenUpToSign, t.PreservesSupport, t.CanDelete13Edge, esc(t.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("```text\n")
	b.WriteString(a.BasisAudit.ExactKPreservingGroup + "\n")
	b.WriteString(a.BasisAudit.KOrbitPreservingExtension + "\n")
	b.WriteString("```\n\n")

	b.WriteString("## Support and harmonic audit\n\n")
	b.WriteString(FormatSupportAudit(a.Support) + "\n\n")
	b.WriteString("| Edge | Delta K | Triangle magnitude | Nearest-neighbor magnitude | Rephasing invariant? | Deleted in NN? | Reason |\n")
	b.WriteString("|---|---:|---:|---:|---|---|---|\n")
	for _, e := range a.Support.Edges {
		b.WriteString(fmt.Sprintf("| %s | %d | %.0f | %.0f | %t | %t | %s |\n", e.Edge, e.DeltaK, e.TriangleMagnitude, e.NearestNeighborMagnitude, e.RephasingInvariant, e.DeletedInNN, esc(e.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("Diagonal rephasing sends `M_ij -> e^{i(alpha_i-alpha_j)} M_ij`. Therefore `|M_ij|` and the zero/nonzero support pattern are invariant. The 1-3 edge is a second KMS harmonic with `DeltaK=2`; it is not removable by a gauge convention.\n\n")

	b.WriteString("## Graph and spectral invariant audit\n\n")
	b.WriteString(FormatSpectralAudit(a.Spectral) + "\n\n")
	b.WriteString("| Invariant | Triangle | Nearest-neighbor chain | Equal? | Reason |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, inv := range a.Spectral.Invariants {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %t | %s |\n", esc(inv.Name), esc(inv.Triangle), esc(inv.NearestNeighbor), inv.Equal, esc(inv.Reason)))
	}
	b.WriteString("\n")
	b.WriteString("```text\n")
	b.WriteString("det(K+epsilon X_triangle)=2 epsilon^3\n")
	b.WriteString("det(K+epsilon X_NN)=0\n")
	b.WriteString("||[K,X_triangle]||_F^2=12\n")
	b.WriteString("||[K,X_NN]||_F^2=4\n")
	b.WriteString("```\n\n")

	b.WriteString("## Gauge-artifact verdict\n\n")
	b.WriteString(FormatGaugeArtifactVerdict(a.Verdict) + "\n\n")
	b.WriteString("A generic `U(3)` family rotation may change displayed texture entries, but only by sending `K_gen` to `UK_genU^dagger`. That destroys the native KMS address and the meaning of the `(2,2)` structural zero. Such a rotation is not an ASHA gauge equivalence; it is a different coordinate system outside the theorem that forced the zero.\n\n")

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
