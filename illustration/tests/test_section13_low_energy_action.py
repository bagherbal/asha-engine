from asha_figures.readme_low_energy_action import low_energy_action_contract, validate_contract


def test_low_energy_action_contract_validates():
    c = low_energy_action_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_LOW_ENERGY_ACTION_SKELETON"


def test_low_energy_action_has_six_terms_and_firewalls():
    q = low_energy_action_contract().quantities
    assert q["sector_count"] == 6
    assert q["sector_order"] == [
        "S_grav",
        "S_gauge",
        "S_Higgs^ASHA",
        "S_fermion",
        "S_Yukawa^ASHA",
        "S_nu^seesaw",
    ]
    assert "(3/8)(1+L)(1/3-S)" in q["higgs_action"]
    assert "Y_f^ASHA" in q["yukawa_action"]
    assert any("M_P^2" in item and "Lambda" in item for item in q["firewalls"])
    assert any("PMNS" in item for item in q["firewalls"])
