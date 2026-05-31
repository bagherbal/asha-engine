from asha_figures.readme_neutrinos import neutrino_contract, validate_contract


def test_neutrino_contract_validates():
    c = neutrino_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_NEUTRINO_SEESAW_PMNS_FIREWALL"


def test_neutrino_rank2_and_firewall_contract():
    q = neutrino_contract().quantities
    lane = q["rank2_normal_order_lane"]
    hs = q["heavy_scale_bridge"]
    assert lane["m1"] == "approximately 0"
    assert 0.16 < lane["m2_over_m3"] < 0.18
    assert 0.045 < lane["m3_eV"] < 0.055
    assert 5.0e10 < hs["M_R3_GeV"] < 8.0e10
    assert "Majorana" in q["firewall"]
