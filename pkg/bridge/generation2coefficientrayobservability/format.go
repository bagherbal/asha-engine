package generation2coefficientrayobservability

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("executed=%t K=%t triangle=%t phase_sealed=%t coeffs_sealed=%t texture_sum_rule=%t ratios_require_amplitudes=%t full_triangle=%t nn_not_gauge=%t gate453_interface=%t promotion_rejected=%t no_empirical=%t verdict=%s", x.Executed, x.Gate444KGenForced, x.Gate445TriangleForced, x.Gate446PhaseQuarantined, x.Gate447CoefficientsSealed, x.Gate450TextureZeroSumRule, x.Gate450RatiosRequireAmplitudes, x.Gate451FullTrianglePreserved, x.Gate452NearestNeighborNotGauge, x.Gate453EmpiricalInterfaceDefined, x.Gate453PromotionRejected, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatRayModel(x RayModel) string {
	return fmt.Sprintf("executed=%t matrix=%s parameterization=%s scale=%s ray_params=%s projective_dim=%d native_selectors=%s selector_count=%d verdict=%s reason=%s", x.Executed, x.MatrixFormula, x.Parameterization, x.AbsoluteScaleParameter, join(x.RayParameters), x.ProjectiveDimension, join(x.NativeSelectors), x.NativeSelectorCount, x.Verdict, x.Reason)
}

func FormatObservableMap(x ObservableMap) string {
	return fmt.Sprintf("%s: inputs=%s formulae=%s rank=%d residual_dof=%d local=%t oriented=%t empirical=%t allowed=%t native_promotion=%t reason=%s", x.Name, join(x.Inputs), join(x.Formulae), x.Rank, x.ResidualDOF, x.LocallyIdentifiesRay, x.GloballyOriented, x.RequiresEmpirical, x.AllowedByGate453, x.NativePromotion, x.Reason)
}

func FormatRankAudit(x RankAudit) string {
	return fmt.Sprintf("executed=%t spectrum_rank=%d spectrum_residual_dof=%d min_local_scalars=%d min_oriented_scalars=%d jacobian=%s sample=%.12g nonzero=%t spectrum_rejected=%t two_scalar_local=%t cp_branch=%t verdict=%s reason=%s", x.Executed, x.SpectrumOnlyRank, x.SpectrumOnlyResidualDOF, x.MinimumLocalScalars, x.MinimumOrientedScalars, x.GenericJacobianFormula, x.GenericJacobianSample, x.GenericJacobianNonzero, x.SpectrumOnlyRejected, x.TwoScalarLocalWorks, x.CPBranchTagRequired, x.Verdict, x.Reason)
}

func FormatProtocol(x Protocol) string {
	return fmt.Sprintf("executed=%t native_ledger=%t spectrum_comparator=%t local_ray_fit=%t cp_oriented_fit=%t explicit_label=%t sector=%t renormalization=%t cp_branch=%t native_coeff_claim=%t spectrum_only_ray_claim=%t verdict=%s reason=%s", x.Executed, x.AllowsNativeLedger, x.AllowsSpectrumOnlyComparator, x.AllowsLocalRayFit, x.AllowsCPOrientedRayFit, x.RequiresExplicitEmpiricalLabel, x.RequiresSectorTag, x.RequiresRenormalizationTag, x.RequiresBranchTagForCPOrientation, x.AllowsNativeCoefficientClaim, x.AllowsSpectrumOnlyRayClaim, x.Verdict, x.Reason)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("executed=%t no_muon=%t no_charm=%t no_yukawa=%t no_ckm=%t no_pmns=%t no_curvefit=%t no_GST=%t no_native_ray=%t K=%t triangle=%t Y_sealed=%t coeffs_sealed=%t cp_sealed=%t native_dim=%d kxy_dim=%d verdict=%s reason=%s", x.Executed, x.NoObservedMuonMassImported, x.NoObservedCharmMassImported, x.NoObservedYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NoCurveFitPromoted, x.NoGSTPromotion, x.NoNativeCoefficientRayValue, x.KGenStillForced, x.XTriangleStillForced, x.YPhaseStillQuarantined, x.SectorCoefficientsStillSealed, x.CPOrientationStillSealed, x.NativeFlavorDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Primary task: %s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 454 Registry Audit — Coefficient-Ray Observability Rank Audit\n\n")
	b.WriteString("## Verdict\n\n")
	b.WriteString("`" + StatusComparatorProtocolDefined + "`\n\n")
	b.WriteString("Gate 454 does not fit flavor data. It computes the rank of the legal Gate-453 comparator interface. The result is sharp: the ASHA coefficient ray has two projective coordinates, spectrum-only data supplies one scalar shape invariant, and an explicitly labelled K-addressed overlap supplies the second local coordinate. CP-oriented uniqueness still requires an explicit phase-branch tag. No coefficient value is native.\n\n")

	b.WriteString("## Inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Coefficient-ray model\n\n")
	b.WriteString(FormatRayModel(a.Ray) + "\n\n")
	b.WriteString("The bridge matrix is\n\n")
	b.WriteString("```text\n")
	b.WriteString("M(a,b,c)=aK_gen+bX_triangle+cY_phase\n")
	b.WriteString("r=sqrt(b^2+c^2), alpha=a/r, phi=atan2(c,b)\n")
	b.WriteString("```\n\n")
	b.WriteString("Removing the absolute scale leaves two ray coordinates: `alpha` and `phi`.\n\n")

	b.WriteString("## Observable-rank maps\n\n")
	b.WriteString(FormatRankAudit(a.Rank) + "\n\n")
	b.WriteString("| Map | Inputs | Formulae | Rank | Residual DOF | Local ray? | CP-oriented? | Allowed? | Reason |\n")
	b.WriteString("|---|---|---|---:|---:|---|---|---|---|\n")
	for _, m := range a.Rank.Maps {
		b.WriteString(fmt.Sprintf("| %s | %s | `%s` | %d | %d | %t | %t | %t | %s |\n", esc(m.Name), esc(join(m.Inputs)), esc(join(m.Formulae)), m.Rank, m.ResidualDOF, m.LocallyIdentifiesRay, m.GloballyOriented, m.AllowedByGate453, esc(m.Reason)))
	}
	b.WriteString("\n")

	b.WriteString("## Rank calculation\n\n")
	b.WriteString("Spectrum-only normalized eigenvalue data gives the cubic shape invariant\n\n")
	b.WriteString("```text\n")
	b.WriteString("I_spec = 2 cos(3 phi)/(alpha^2+3)^(3/2)\n")
	b.WriteString("```\n\n")
	b.WriteString("This has rank one, so one continuous ray coordinate remains. Adding the K-addressed overlap\n\n")
	b.WriteString("```text\n")
	b.WriteString("I_K = Tr(MK)/sqrt(Tr(M^2)Tr(K^2)) = alpha/sqrt(alpha^2+3)\n")
	b.WriteString("```\n\n")
	b.WriteString("gives the generic Jacobian\n\n")
	b.WriteString("```text\n")
	b.WriteString(a.Rank.GenericJacobianFormula + "\n")
	b.WriteString("```\n\n")
	b.WriteString(fmt.Sprintf("At the audit sample the determinant is `%.12g`, so the local rank is two away from the expected phase-caustic loci `sin(3 phi)=0`.\n\n", a.Rank.GenericJacobianSample))

	b.WriteString("## Comparator protocol\n\n")
	b.WriteString(FormatProtocol(a.Protocol) + "\n\n")
	b.WriteString("Allowed use is strictly empirical and labelled: sector tag, renormalization scale/scheme tag, and a CP branch tag when oriented phase selection is claimed. Spectrum-only fitting cannot identify the ray and cannot be promoted to native law.\n\n")

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
