from __future__ import annotations

from dataclasses import dataclass, asdict
from typing import Any


@dataclass(frozen=True)
class FigureContract:
    figure_id: str
    title: str
    readme_anchor: str
    status: str
    quantities: dict[str, Any]
    notes: list[str]


def theorem_stack_contract() -> FigureContract:
    theorem_nodes = [
        "Lorentzianized phase-space octave",
        "Flat metric projection",
        "Contact-seven source structure",
        "Contact phase-triple projector algebra",
        "Product-depth formal spectral extension",
        "Minimal commuting contact-depth matter extension",
        "Yukawa broadcast breaking by e^{-4πN_Q}",
        "Relative flavor orientation theorem",
        "Majorana Takagi and seesaw structure theorems",
        "Low-energy metric dynamics under standard assumptions",
        "Dimensional obstruction theorem",
        "Vacuum-zero independence theorem",
    ]
    quantities = {
        "node_count": len(theorem_nodes),
        "edge_count": len(theorem_nodes) - 1,
        "theorem_nodes": theorem_nodes,
        "direction": "source-typed descent, top-to-bottom / left-to-right acyclic order",
        "categories": {
            "native_or_theorem_law": [
                "Lorentzianized phase-space octave",
                "Flat metric projection",
                "Contact-seven source structure",
                "Contact phase-triple projector algebra",
                "Product-depth formal spectral extension",
                "Minimal commuting contact-depth matter extension",
                "Yukawa broadcast breaking by e^{-4πN_Q}",
                "Relative flavor orientation theorem",
                "Majorana Takagi and seesaw structure theorems",
                "Dimensional obstruction theorem",
                "Vacuum-zero independence theorem",
            ],
            "conditional_standard_low_energy_law": [
                "Low-energy metric dynamics under standard assumptions",
            ],
        },
        "firewalls": [
            "The theorem-level stack is not the locked physical filling list.",
            "It does not by itself derive v/M_P, Higgs mass, CKM numerics, quark/lepton masses, or PMNS phases.",
            "Every arrow is an architectural dependency, not a claim that downstream physical values are native.",
        ],
    }
    return FigureContract(
        figure_id="asha_readme_theorem_level_stack",
        title="Theorem-Level Stack",
        readme_anchor="Theorem-level stack",
        status="PASS_README_THEOREM_LEVEL_STACK",
        quantities=quantities,
        notes=[
            "The figure turns the README bullet list into an acyclic source-typed theorem rail.",
            "The list has exactly twelve theorem/conditional-theorem nodes and eleven forward dependency arrows.",
            "A lower obsidian firewall separates theorem-level architecture from physical filling and remaining wounds.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    problems: list[str] = []
    if c.figure_id != "asha_readme_theorem_level_stack":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    q = c.quantities
    expected = [
        "Lorentzianized phase-space octave",
        "Flat metric projection",
        "Contact-seven source structure",
        "Contact phase-triple projector algebra",
        "Product-depth formal spectral extension",
        "Minimal commuting contact-depth matter extension",
        "Yukawa broadcast breaking by e^{-4πN_Q}",
        "Relative flavor orientation theorem",
        "Majorana Takagi and seesaw structure theorems",
        "Low-energy metric dynamics under standard assumptions",
        "Dimensional obstruction theorem",
        "Vacuum-zero independence theorem",
    ]
    checks = [
        q["node_count"] == 12,
        q["edge_count"] == 11,
        q["theorem_nodes"] == expected,
        q["direction"].startswith("source-typed descent"),
        len(q["categories"]["conditional_standard_low_energy_law"]) == 1,
        "Low-energy metric dynamics under standard assumptions" in q["categories"]["conditional_standard_low_energy_law"],
        len(q["firewalls"]) == 3,
        "not the locked physical filling" in q["firewalls"][0],
        "PMNS" in q["firewalls"][1],
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more theorem-level stack checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
