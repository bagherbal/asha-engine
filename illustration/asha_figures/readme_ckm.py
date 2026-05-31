from __future__ import annotations

from dataclasses import dataclass, asdict
from math import pi
from typing import Any

from .readme_next5 import L, S_SPLIT


@dataclass(frozen=True)
class FigureContract:
    figure_id: str
    title: str
    readme_anchor: str
    status: str
    quantities: dict[str, Any]
    notes: list[str]


def ckm_contract() -> FigureContract:
    theta12 = 1/4 - 18*S_SPLIT + 158*S_SPLIT*S_SPLIT
    theta23 = L + (5/3)*S_SPLIT - (8 - 2*L)*S_SPLIT*S_SPLIT
    theta13 = 72*L*S_SPLIT - (3/2)*S_SPLIT*S_SPLIT
    delta = pi/3 + 71*S_SPLIT + (93/4)*S_SPLIT*S_SPLIT
    quantities = {
        "source_constants": {"L": L, "S_split": S_SPLIT},
        "orientation_identity": "V_CKM = U_u^† U_d",
        "interpretation": "masses are depth eigenvalues; mixing is relative orientation",
        "boundary_angles_rad": {
            "theta12_0": theta12,
            "theta23_0": theta23,
            "theta13_0": theta13,
            "delta_CKM_0": delta,
        },
        "boundary_angles_deg": {
            "theta12_0": theta12*180/pi,
            "theta23_0": theta23*180/pi,
            "theta13_0": theta13*180/pi,
            "delta_CKM_0": delta*180/pi,
        },
        "formulas": {
            "theta12_0": "1/4 - 18S + 158S^2",
            "theta23_0": "L + (5/3)S - (8-2L)S^2",
            "theta13_0": "72LS - (3/2)S^2",
            "delta_CKM_0": "π/3 + 71S + (93/4)S^2",
            "Gamma_q(mu)": "(3/(32π^2)) ∫_{ln M_Z}^{ln μ} y_t(t)^2 dt",
            "D_u(mu)": "D_u^0 - Γ_q(μ) + ε_u(μ)",
            "D_d(mu)": "D_d^0 + Γ_q(μ) + ε_d(μ)",
            "theta23(mu)": "θ23^0 + (1/24+2S) Γ_q(μ)",
            "theta13(mu)": "θ13^0 + (1/256) Γ_q(μ)",
        },
        "transport_coefficients": {
            "theta23_gamma_coefficient": 1/24 + 2*S_SPLIT,
            "theta13_gamma_coefficient": 1/256,
            "D_u_gamma_sign": -1,
            "D_d_gamma_sign": +1,
        },
        "source_typed_coefficients": {
            "158": "2(72+7)",
            "71": "locked README CKM phase coefficient",
            "93/4": "locked README CKM phase quadratic coefficient",
        },
        "firewall": "CKM boundary orientation plus top-driven RG transport; not a native flavor theorem",
    }
    return FigureContract(
        figure_id="asha_readme_ckm_orientation_transport",
        title="CKM Mixing: Relative Orientation and Top-Driven Transport",
        readme_anchor="CKM mixing",
        status="PASS_README_CKM_ORIENTATION_TRANSPORT",
        quantities=quantities,
        notes=[
            "The figure separates depth eigenvalues from relative basis orientation U_u^†U_d.",
            "The four CKM boundary parameters are exact README formulas evaluated from locked L and S.",
            "Γ_q is displayed as a transport operator; no unprovided y_t(t) running curve is invented.",
            "This is a boundary/RG transport board, not a native flavor theorem or PMNS selector.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    if c.figure_id != "asha_readme_ckm_orientation_transport":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    q = c.quantities
    a = q["boundary_angles_rad"]
    problems: list[str] = []
    checks = [
        abs(a["theta12_0"] - (1/4 - 18*S_SPLIT + 158*S_SPLIT*S_SPLIT)) < 1e-14,
        abs(a["theta23_0"] - (L + (5/3)*S_SPLIT - (8 - 2*L)*S_SPLIT*S_SPLIT)) < 1e-14,
        abs(a["theta13_0"] - (72*L*S_SPLIT - (3/2)*S_SPLIT*S_SPLIT)) < 1e-14,
        abs(a["delta_CKM_0"] - (pi/3 + 71*S_SPLIT + (93/4)*S_SPLIT*S_SPLIT)) < 1e-14,
        0.22 < a["theta12_0"] < 0.24,
        0.040 < a["theta23_0"] < 0.044,
        0.0034 < a["theta13_0"] < 0.0040,
        1.12 < a["delta_CKM_0"] < 1.16,
        abs(q["transport_coefficients"]["theta23_gamma_coefficient"] - (1/24 + 2*S_SPLIT)) < 1e-14,
        abs(q["transport_coefficients"]["theta13_gamma_coefficient"] - 1/256) < 1e-15,
        q["source_typed_coefficients"]["158"] == "2(72+7)",
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more CKM formula/range/source checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
