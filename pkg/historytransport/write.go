package historytransport

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func WriteBundle(dir string, b Bundle) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "01_inputs.yaml"), []byte(inputsYAML(b.Inputs)), 0o644); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "02_end_vector.json"), b.EndVector); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "03_boundary_running.json"), b.GaugeBoundary); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "04_scalar_transport.json"), b.ScalarTransport); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "05_flavor_transport.json"), b.FlavorTransport); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "06_history_residual.json"), b.HistoryResidual); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "07_summary.md"), []byte(SummaryMarkdown(b)), 0o644); err != nil {
		return err
	}
	return nil
}

func writeJSON(path string, v any) error {
	buf, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	buf = append(buf, '\n')
	return os.WriteFile(path, buf, 0o644)
}

func inputsYAML(in InputSet) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "task_name: %q\n", in.TaskName)
	fmt.Fprintf(&sb, "mu0_name: %q\nmu0_gev: %.10g\n", in.Mu0Name, in.Mu0GeV)
	fmt.Fprintf(&sb, "asha_boundary_law:\n  k_y: %.12g\n  sin2_theta_boundary: %.12g\n  canonical_boundary_relation: %q\n  finite_algebra: %q\n  scalar_carrier: %q\n  bridge_only: true\n", in.ASHABoundary.KY, in.ASHABoundary.Sin2ThetaBoundary, in.ASHABoundary.CanonicalBoundaryRelation, in.ASHABoundary.FiniteAlgebra, in.ASHABoundary.ScalarCarrier)
	fmt.Fprintf(&sb, "measured:\n")
	keys := []string{"G_F", "m_W", "m_Z", "m_H", "alpha_s_MZ"}
	for _, k := range keys {
		m := in.Measured[k]
		fmt.Fprintf(&sb, "  %s:\n    name: %q\n    value: %.12g\n    uncertainty: %.12g\n    unit: %q\n", k, m.Name, m.Value, m.Uncertainty, m.Unit)
		if m.Scale != "" {
			fmt.Fprintf(&sb, "    scale: %q\n", m.Scale)
		}
		if m.Scheme != "" {
			fmt.Fprintf(&sb, "    scheme: %q\n", m.Scheme)
		}
		fmt.Fprintf(&sb, "    source_id: %q\n    role: %q\n    bridge_only: true\n", m.SourceID, m.Role)
	}
	fmt.Fprintf(&sb, "fermions:\n")
	for _, f := range in.Fermions {
		fmt.Fprintf(&sb, "  - name: %q\n    mass_gev: %.12g\n    mass_uncertainty_gev: %.12g\n    input_scale_gev: %.12g\n    target_scale_gev: %.12g\n    mass_at_mz_gev: %.12g\n    scheme: %q\n    transport: %q\n    source_id: %q\n    bridge_only: true\n", f.Name, f.MassGeV, f.MassUncertainty, f.InputScaleGeV, f.TargetScaleGeV, f.MassAtMZGeV, f.Scheme, f.Transport, f.SourceID)
	}
	fmt.Fprintf(&sb, "ckm:\n  s12: %.12g\n  s13: %.12g\n  s23: %.12g\n  delta_rad: %.12g\n  source_id: %q\n", in.CKM.S12, in.CKM.S13, in.CKM.S23, in.CKM.Delta, in.CKM.SourceID)
	fmt.Fprintf(&sb, "cosmology:\n  omega_c_h2: %.12g\n  omega_b_h2: %.12g\n  n_s: %.12g\n  tau: %.12g\n  source_id: %q\n  bridge_only: true\n", in.Cosmology.OmegaCH2, in.Cosmology.OmegaBH2, in.Cosmology.NS, in.Cosmology.Tau, in.Cosmology.SourceID)
	fmt.Fprintf(&sb, "sources:\n")
	for _, s := range in.Sources {
		fmt.Fprintf(&sb, "  - id: %q\n    title: %q\n    url: %q\n    version: %q\n    note: %q\n", s.ID, s.Title, s.URL, s.Version, s.Note)
	}
	fmt.Fprintf(&sb, "warnings:\n")
	for _, w := range in.Warnings {
		fmt.Fprintf(&sb, "  - %q\n", w)
	}
	return sb.String()
}

func SummaryMarkdown(b Bundle) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "# ASHA History Transport End Calculation v1\n\n")
	fmt.Fprintf(&sb, "Task: `%s`\n\n", TaskName)
	fmt.Fprintf(&sb, "This is an observed-data transport calculation, not a native ASHA derivation.  It asks how the ASHA boundary normalization `k_Y=5/3`, `g1=g2`, and `sin^2(theta_*)=3/8` maps to the measured endpoint at `mu0=M_Z`.\n\n")
	fmt.Fprintf(&sb, "## Status ledger\n\n")
	for _, s := range b.Statuses {
		fmt.Fprintf(&sb, "- `%s`\n", s)
	}
	fmt.Fprintf(&sb, "\n## Phase 1 — End vector at M_Z\n\n")
	fmt.Fprintf(&sb, "- `v = %.12g GeV`\n- `gY = %.12g`\n- `g1 = %.12g`\n- `g2 = %.12g`\n- `g3 = %.12g`\n- `sin2_theta_End = %.12g`\n- `lambda(M_Z) = %.12g`\n", b.EndVector.VGeV, b.EndVector.GY, b.EndVector.G1, b.EndVector.G2, b.EndVector.G3, b.EndVector.Sin2Theta, b.EndVector.Lambda)
	fmt.Fprintf(&sb, "\n## Phase 2 — Gauge boundary running\n\n")
	fmt.Fprintf(&sb, "- `Lambda_12 = %.12g GeV`\n- `g_star = %.12g`\n- `g3(Lambda_12) = %.12g`\n- `Delta_3 = %.12g`\n- `R_3 = %.12g`\n- interpretation: `%s`\n", b.GaugeBoundary.Lambda12GeV, b.GaugeBoundary.GStar, b.GaugeBoundary.G3Lambda, b.GaugeBoundary.Delta3, b.GaugeBoundary.R3, b.GaugeBoundary.Interpretation)
	fmt.Fprintf(&sb, "\n## Phase 3 — Weak-angle transport\n\n")
	fmt.Fprintf(&sb, "- `sin2_theta_boundary = %.12g`\n- `sin2_theta_End = %.12g`\n- `Delta_sin2 = %.12g`\n- `transport_required = %t`\n", b.WeakAngleTransport.Sin2ThetaBoundary, b.WeakAngleTransport.Sin2ThetaEnd, b.WeakAngleTransport.DeltaSin2, b.WeakAngleTransport.TransportRequired)
	fmt.Fprintf(&sb, "\n## Phase 4 — Scalar transport\n\n")
	if b.ScalarTransport.ZeroCrossingScaleGeV != nil {
		fmt.Fprintf(&sb, "- `lambda(Lambda_12) = %.12g`\n- `y_t(Lambda_12) = %.12g`\n- `beta_lambda(M_Z) = %.12g`\n- zero crossing: `%.12g GeV`\n- status: `%s`\n", b.ScalarTransport.LambdaLambda12, b.ScalarTransport.YT_Lambda12, b.ScalarTransport.BetaLambdaMZ, *b.ScalarTransport.ZeroCrossingScaleGeV, b.ScalarTransport.VacuumStabilityStatus)
	} else {
		fmt.Fprintf(&sb, "- `lambda(Lambda_12) = %.12g`\n- `y_t(Lambda_12) = %.12g`\n- `beta_lambda(M_Z) = %.12g`\n- zero crossing: none before Lambda_12 in v1 approximation\n- status: `%s`\n", b.ScalarTransport.LambdaLambda12, b.ScalarTransport.YT_Lambda12, b.ScalarTransport.BetaLambdaMZ, b.ScalarTransport.VacuumStabilityStatus)
	}
	fmt.Fprintf(&sb, "\n## Phase 5 — Flavor transport\n\n")
	fmt.Fprintf(&sb, "- `J_CKM = %.12g`\n- `Koide_Qe = %.12g`\n- convention: `%s`\n", b.FlavorTransport.JCKM, b.FlavorTransport.KoideQe, b.FlavorTransport.Convention)
	for _, p := range b.FlavorTransport.ResidualPatterns {
		fmt.Fprintf(&sb, "- %s\n", p)
	}
	fmt.Fprintf(&sb, "\n## Phase 6 — History residual\n\n")
	fmt.Fprintf(&sb, "The residual vector is nonzero and therefore records history seals:\n\n")
	fmt.Fprintf(&sb, "```text\nR_hist = (Delta_3=%.12g, Delta_sin2=%.12g, lambda_Lambda12=%.12g, J_CKM=%.12g, Koide_Qe=%.12g, Omega_c h^2=%.12g, Omega_b h^2=%.12g)\n```\n\n", b.HistoryResidual.Gauge.Delta3, b.HistoryResidual.WeakAngle.DeltaSin2, b.HistoryResidual.Scalar.LambdaLambda12, b.HistoryResidual.Flavor.JCKM, b.HistoryResidual.Flavor.KoideQe, b.HistoryResidual.Cosmology.OmegaCH2, b.HistoryResidual.Cosmology.OmegaBH2)
	fmt.Fprintf(&sb, "## Firewall\n\n")
	fmt.Fprintf(&sb, "The calculation preserves the ASHA boundary: it does not claim full gauge unification, does not hide thresholds, does not derive Yukawa/flavor data, does not derive Planck cosmology, and does not import observed masses as native finite algebra.\n")
	return sb.String()
}
