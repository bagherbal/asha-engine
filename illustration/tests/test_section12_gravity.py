from asha_figures.readme_gravity import gravity_contract, validate_contract


def test_gravity_contract_validates():
    c = gravity_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_GRAVITY_METRIC_PROJECTION_FIREWALL"


def test_gravity_projection_and_source_firewalls():
    q = gravity_contract().quantities
    assert q["active_signature"] == [1, 7]
    assert q["projected_signature"] == [1, 3]
    assert q["flat_projection_status"] == "theorem-level"
    assert q["source_types"]["M_P^2"] == "Planck stiffness / metric-response scale"
    assert q["source_types"]["Lambda"] == "vacuum-boundary residual"
    assert any("does not yet derive Lambda" in item for item in q["firewalls"])
