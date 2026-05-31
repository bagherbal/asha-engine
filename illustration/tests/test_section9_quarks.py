from asha_figures.readme_quarks import quarks_contract, validate_contract


def test_quark_contract_validates():
    c = quarks_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_QUARK_LANES_SHAPE_LAWS"


def test_quark_mass_ordering_and_sources():
    q = quarks_contract().quantities
    up = q["up_sector_shapes"]["solved_masses_GeV"]
    down = q["down_sector_shapes"]["solved_masses_GeV"]
    assert up["m_u"] < up["m_c"] < up["m_t"]
    assert down["m_d"] < down["m_s"] < down["m_b"]
    assert q["top_lane"]["coefficient_source"] == "138=2(72-3)"
    assert q["bottom_lane"]["coefficient_source"] == "106=2(56-3)"
