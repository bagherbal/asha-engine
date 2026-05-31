from asha_figures.readme_next5 import next5_contracts, validate_contract


def test_next5_contracts_validate():
    contracts = next5_contracts()
    assert len(contracts) == 5
    for contract in contracts:
        result = validate_contract(contract)
        assert result["ok"], contract.figure_id
        assert result["status"].startswith("PASS_README_")


def test_next5_order_matches_readme_sequence():
    assert [c.readme_anchor for c in next5_contracts()] == [
        "Matter sockets and product depth",
        "Locked constants and source alphabet",
        "Planck-to-electroweak scale bridge",
        "Higgs sector",
        "Charged leptons",
    ]
