from asha_figures.projector_universe_zoom6s import compute_zoom6s, render_svg


def test_zoom6s_all_six_action_sectors_visible():
    data = compute_zoom6s(samples=512)
    assert data["validation"]["status"] == "PASS_ZOOM6S_ALL_SIX_ACTION_SECTORS_VISIBLE"
    checks = data["validation"]["checks"]
    assert checks["sector_count"] == 6
    assert checks["expected_keys_present"] is True
    assert checks["central_flow_passed"] is True
    assert checks["all_sector_discs_inside_view"] is True
    assert checks["gauge_pulse_count"] == 12
    assert checks["fermion_node_count"] == 16
    assert checks["yukawa_lane_count"] == 3


def test_zoom6s_render_is_text_free():
    data = compute_zoom6s(samples=256)
    svg = render_svg(data)
    assert "<text" not in svg.lower()
    assert "S_grav" not in svg
    assert "S_Higgs" not in svg
