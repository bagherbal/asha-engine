from __future__ import annotations

from dataclasses import dataclass, asdict
from math import exp, pi, sqrt
from typing import Any

from .readme_next5 import L, S_SPLIT, M_PLANCK_REDUCED_GEV, scale_bridge_contract, charged_leptons_contract


@dataclass(frozen=True)
class FigureContract:
    figure_id: str
    title: str
    readme_anchor: str
    status: str
    quantities: dict[str, Any]
    notes: list[str]


def neutrino_contract() -> FigureContract:
    v = scale_bridge_contract().quantities["v_GeV"]
    A_tau = charged_leptons_contract().quantities["A_tau"]
    rank2_ratio = 4*L + 10*S_SPLIT
    MR3 = (sqrt(2*pi) + 49*S_SPLIT + 90*S_SPLIT*S_SPLIT) * sqrt(v*M_PLANCK_REDUCED_GEV)
    MR2 = MR3 * exp(-4*pi/3) / rank2_ratio
    m3 = v*v * exp(-2*A_tau) / (2*MR3)
    m2 = rank2_ratio * m3
    theta23_hi = pi/4 + 48*S_SPLIT
    theta23_lo = pi/4 - 48*S_SPLIT
    theta12 = pi/6 + 48*S_SPLIT
    theta13 = 4*L - 7*S_SPLIT + 4*S_SPLIT*S_SPLIT
    quantities = {
        "source_constants": {"L": L, "S_split": S_SPLIT, "M_P_reduced_GeV": M_PLANCK_REDUCED_GEV, "v_GeV": v},
        "seesaw_identity": "M_nu = -(v^2/2) Y_nu^T M_R^{-1} Y_nu",
        "rank2_normal_order_lane": {
            "m1": "approximately 0",
            "m2_over_m3": rank2_ratio,
            "formula": "m2 = (4L + 10S) m3",
            "m2_eV": m2*1e9,
            "m3_eV": m3*1e9,
        },
        "heavy_scale_bridge": {
            "M_R3_GeV": MR3,
            "M_R2_GeV": MR2,
            "M_R3_formula": "(sqrt(2π)+49S+90S^2) sqrt(v M_P)",
            "M_R2_formula": "M_R3 exp(-4π/3)/(4L+10S)",
            "m3_formula": "v^2 exp(-2A_tau)/(2M_R3)",
        },
        "pmns_skeleton_rad": {
            "theta23_minus": theta23_lo,
            "theta23_plus": theta23_hi,
            "theta12": theta12,
            "theta13": theta13,
        },
        "pmns_formulas": {
            "theta23_PMNS": "π/4 ± 48S",
            "theta12_PMNS": "π/6 + 48S",
            "theta13_PMNS": "4L - 7S + 4S^2",
        },
        "firewall": "PMNS is not locked; Majorana orientation selector and Majorana phases remain unresolved.",
    }
    return FigureContract(
        figure_id="asha_readme_neutrino_seesaw_pmns_firewall",
        title="Neutrinos: Seesaw Lane and PMNS Firewall",
        readme_anchor="Neutrinos",
        status="PASS_README_NEUTRINO_SEESAW_PMNS_FIREWALL",
        quantities=quantities,
        notes=[
            "The figure separates the Type-I seesaw theorem skeleton from the sealed rank-2 normal-order lane.",
            "M_R3 and M_R2 are displayed as heavy-scale bridge objects, with M_P kept as Planck-stiffness filling.",
            "PMNS angles are shown only as an unresolved skeleton; Majorana phases and the orientation selector remain firewalled.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    if c.figure_id != "asha_readme_neutrino_seesaw_pmns_firewall":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    q = c.quantities
    problems: list[str] = []
    r = q["rank2_normal_order_lane"]["m2_over_m3"]
    hs = q["heavy_scale_bridge"]
    pmns = q["pmns_skeleton_rad"]
    checks = [
        abs(r - (4*L + 10*S_SPLIT)) < 1e-14,
        0.16 < r < 0.18,
        5.0e10 < hs["M_R3_GeV"] < 8.0e10,
        4.0e9 < hs["M_R2_GeV"] < 7.0e9,
        0.045 < q["rank2_normal_order_lane"]["m3_eV"] < 0.055,
        0.007 < q["rank2_normal_order_lane"]["m2_eV"] < 0.010,
        abs(pmns["theta23_plus"] - (pi/4 + 48*S_SPLIT)) < 1e-14,
        abs(pmns["theta23_minus"] - (pi/4 - 48*S_SPLIT)) < 1e-14,
        abs(pmns["theta12"] - (pi/6 + 48*S_SPLIT)) < 1e-14,
        abs(pmns["theta13"] - (4*L - 7*S_SPLIT + 4*S_SPLIT*S_SPLIT)) < 1e-14,
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more neutrino seesaw/rank/PMNS-skeleton checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
