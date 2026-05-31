from asha_figures.readme_ckm import ckm_contract, validate_contract


def test_ckm_contract_validates():
    c = ckm_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_CKM_ORIENTATION_TRANSPORT"


def test_ckm_boundary_angles_and_transport_coefficients():
    q = ckm_contract().quantities
    a = q["boundary_angles_rad"]
    assert 0.22 < a["theta12_0"] < 0.24
    assert 0.040 < a["theta23_0"] < 0.044
    assert 0.0034 < a["theta13_0"] < 0.0040
    assert 1.12 < a["delta_CKM_0"] < 1.16
    assert q["source_typed_coefficients"]["158"] == "2(72+7)"
    assert q["transport_coefficients"]["D_u_gamma_sign"] == -1
    assert q["transport_coefficients"]["D_d_gamma_sign"] == 1
