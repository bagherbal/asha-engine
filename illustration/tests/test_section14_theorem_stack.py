from asha_figures.readme_theorem_stack import theorem_stack_contract, validate_contract


def test_theorem_stack_contract_validates():
    c = theorem_stack_contract()
    result = validate_contract(c)
    assert result["ok"], result
    assert result["status"] == "PASS_README_THEOREM_LEVEL_STACK"


def test_theorem_stack_exact_readme_order_and_firewalls():
    q = theorem_stack_contract().quantities
    assert q["node_count"] == 12
    assert q["edge_count"] == 11
    assert q["theorem_nodes"][0] == "Lorentzianized phase-space octave"
    assert q["theorem_nodes"][6] == "Yukawa broadcast breaking by e^{-4πN_Q}"
    assert q["theorem_nodes"][-1] == "Vacuum-zero independence theorem"
    assert q["categories"]["conditional_standard_low_energy_law"] == [
        "Low-energy metric dynamics under standard assumptions"
    ]
    assert any("not the locked physical filling" in item for item in q["firewalls"])
    assert any("PMNS" in item for item in q["firewalls"])
