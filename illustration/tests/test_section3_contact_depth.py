from asha_figures.contact_depth import build_contact_depth_geometry, validate_contact_depth_geometry


def test_section3_contact_depth_contract():
    geometry = build_contact_depth_geometry()
    checks = validate_contact_depth_geometry(geometry)
    assert checks["v8_basis_exact"]
    assert checks["observer_reference_exact"]
    assert checks["contact_basis_exact"]
    assert checks["contact_basis_count_exact"]
    assert checks["observer_removed_from_contact"]
    assert checks["decomposition_exact"]
    assert checks["q_contact_dim_exact"]
    assert checks["n_q_exact"]
    assert checks["w_q_numerically_exact"]
    assert checks["w_q_strictly_descending"]
    assert checks["phase_planes_exact"]
    assert checks["status"] == "PASS_README_CONTACT_SEVEN_DEPTH_TRIPLE"
