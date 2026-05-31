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


def gravity_contract() -> FigureContract:
    quantities = {
        "phase_space_carrier": "V8 = X4 ⊕ P4",
        "active_metric": "eta_ASHA = eta_1,3 ⊕ (-I4)",
        "active_signature": [1, 7],
        "symplectic_form": "Omega = sum_{mu=0}^3 dp_mu ∧ dx^mu",
        "projection": "Pi_X : V8 -> X4",
        "projected_metric": "g_X = Pi_X^* eta_ASHA Pi_X = eta_1,3",
        "projected_signature": [1, 3],
        "flat_projection_status": "theorem-level",
        "gravitational_action": "S_grav = (M_P^2/2) ∫ (R - 2 Lambda) sqrt(-g) d^4x",
        "assumptions": ["smooth Lorentzian metric g_{mu nu}", "standard low-energy assumptions"],
        "source_types": {
            "M_P^2": "Planck stiffness / metric-response scale",
            "Lambda": "vacuum-boundary residual",
        },
        "firewalls": [
            "ASHA does not yet derive M_P^2 natively.",
            "ASHA does not yet derive Lambda natively.",
            "The curved low-energy action is conditional on standard gravitational assumptions.",
        ],
    }
    return FigureContract(
        figure_id="asha_readme_gravity_metric_projection_firewall",
        title="Gravity: Metric Projection and Planck-Stiffness Firewall",
        readme_anchor="Gravity",
        status="PASS_README_GRAVITY_METRIC_PROJECTION_FIREWALL",
        quantities=quantities,
        notes=[
            "The figure separates the theorem-level flat metric projection from the conditional low-energy gravitational action.",
            "M_P^2 is displayed as a Planck-stiffness / metric-response filling, not as a native ASHA theorem.",
            "Lambda is displayed as a vacuum-boundary residual wound, not as a derived cosmological constant.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    problems: list[str] = []
    if c.figure_id != "asha_readme_gravity_metric_projection_firewall":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    q = c.quantities
    checks = [
        q["phase_space_carrier"] == "V8 = X4 ⊕ P4",
        q["active_signature"] == [1, 7],
        q["projected_signature"] == [1, 3],
        "eta_1,3" in q["projected_metric"],
        "M_P^2/2" in q["gravitational_action"],
        "R - 2 Lambda" in q["gravitational_action"],
        q["source_types"]["M_P^2"] == "Planck stiffness / metric-response scale",
        q["source_types"]["Lambda"] == "vacuum-boundary residual",
        len(q["firewalls"]) == 3,
        q["flat_projection_status"] == "theorem-level",
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more gravity projection/source-type checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
