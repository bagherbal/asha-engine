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


def low_energy_action_contract() -> FigureContract:
    quantities = {
        "low_energy_action": "S_ASHA^low = S_grav + S_gauge + S_Higgs^ASHA + S_fermion + S_Yukawa^ASHA + S_nu^seesaw",
        "sector_order": [
            "S_grav",
            "S_gauge",
            "S_Higgs^ASHA",
            "S_fermion",
            "S_Yukawa^ASHA",
            "S_nu^seesaw",
        ],
        "sector_count": 6,
        "sector_source_types": {
            "S_grav": "conditional low-energy metric lane; M_P^2 and Lambda not native",
            "S_gauge": "standard gauge kinetic lane after finite/internal data are installed",
            "S_Higgs^ASHA": "locked ASHA Higgs quartic/proxy lane",
            "S_fermion": "standard fermion kinetic lane over low-energy metric background",
            "S_Yukawa^ASHA": "ASHA Yukawa shape laws plus sealed flavor/orientation data",
            "S_nu^seesaw": "rank-2 normal seesaw bridge lane; PMNS phases unresolved",
        },
        "higgs_action": "S_Higgs^ASHA = ∫ sqrt(-g)[g^{mu nu}(D_mu phi)^dagger(D_nu phi) - (3/8)(1+L)(1/3-S)(|phi|^2-v^2/2)^2] d^4x",
        "yukawa_action": "S_Yukawa^ASHA = -∫ sqrt(-g) sum_f[bar(psi)_{f,L} Y_f^ASHA phi_f psi_{f,R} + h.c.] d^4x",
        "higgs_coefficient": "(3/8)(1+L)(1/3-S)",
        "potential_center": "|phi|^2 - v^2/2",
        "firewalls": [
            "The low-energy action is a skeleton, not a single native derivation of every term.",
            "M_P^2 and Lambda remain source-typed wounds inherited from the gravity section.",
            "Y_f^ASHA uses ASHA shape laws plus bridge/seal data, not a native full flavor theorem.",
            "PMNS CP and Majorana phase selection remain unresolved in the neutrino lane.",
        ],
    }
    return FigureContract(
        figure_id="asha_readme_low_energy_action_skeleton",
        title="Current Low-Energy ASHA Action Skeleton",
        readme_anchor="Current low-energy ASHA action",
        status="PASS_README_LOW_ENERGY_ACTION_SKELETON",
        quantities=quantities,
        notes=[
            "The figure presents the README action as a six-sector source-typed architecture rather than a fully native theorem.",
            "The Higgs coefficient and Yukawa term are rendered as the two ASHA-specific interior equations.",
            "Obsidian firewall marks the unresolved gravity, flavor, and PMNS wounds inherited by the low-energy skeleton.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    problems: list[str] = []
    if c.figure_id != "asha_readme_low_energy_action_skeleton":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    q = c.quantities
    checks = [
        q["sector_count"] == 6,
        q["sector_order"] == [
            "S_grav",
            "S_gauge",
            "S_Higgs^ASHA",
            "S_fermion",
            "S_Yukawa^ASHA",
            "S_nu^seesaw",
        ],
        "S_grav + S_gauge" in q["low_energy_action"],
        "S_Higgs^ASHA" in q["low_energy_action"],
        "S_Yukawa^ASHA" in q["low_energy_action"],
        "S_nu^seesaw" in q["low_energy_action"],
        "(3/8)(1+L)(1/3-S)" in q["higgs_action"],
        "|phi|^2-v^2/2" in q["higgs_action"],
        "Y_f^ASHA" in q["yukawa_action"],
        "sum_f" in q["yukawa_action"],
        len(q["firewalls"]) == 4,
        "M_P^2" in q["firewalls"][1] and "Lambda" in q["firewalls"][1],
        "PMNS" in q["firewalls"][3],
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more low-energy action skeleton checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
