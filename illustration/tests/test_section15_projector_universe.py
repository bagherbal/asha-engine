
from asha_figures.projector_universe import compute_flow


def test_programmatic_projector_universe_contract():
    flow = compute_flow(samples=360)
    assert flow["validation"]["status"] == "PASS_PROGRAMMATIC_X4_P4_PROJECTOR_ONE_TRAJECTORY"
    checks = flow["validation"]["checks"]
    assert checks["eta_signature"] == {"positive": 1, "negative": 7, "zero": 0}
    assert checks["pi_x_idempotent_norm"] < 1e-12
    assert checks["pi_p_idempotent_norm"] < 1e-12
    assert checks["pi_x_pi_p_orthogonal_norm"] < 1e-12
    assert checks["pi_sum_identity_norm"] < 1e-12
    assert checks["omega_skew_norm"] < 1e-12
    assert checks["hamiltonian_relative_drift"] < 1e-12
    assert checks["depth_W_strictly_descending"] is True


def test_text_free_geometry_payload_has_expected_samples():
    flow = compute_flow(samples=512)
    assert flow["samples"] == 512
    assert len(flow["visual"]["trajectory"]) == 512
    assert len(flow["visual"]["x_projection_2d"]) == 512
    assert len(flow["visual"]["p_projection_2d"]) == 512
    assert len(flow["visual"]["v_projection_2d"]) == 512
    assert len(flow["contract"]["pi_x_matrix"]) == 8
    assert len(flow["contract"]["pi_p_matrix"]) == 8
