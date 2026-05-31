from asha_figures.contact_vacuum import build_section2_geometry, validate_section2_geometry


def test_section2_rank_chain_and_log_area_scaling():
    geometry = build_section2_geometry()
    checks = validate_section2_geometry(geometry)
    assert checks["rank_chain_exact"]
    assert checks["log_area_scaling_exact"]
    assert checks["status"] == "PASS_LOG_RANK_TOPOLOGICAL_CONTACT_VACUUM"
