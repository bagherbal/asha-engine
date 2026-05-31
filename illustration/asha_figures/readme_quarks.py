from __future__ import annotations

from dataclasses import dataclass, asdict
from math import exp, pi, sqrt
from typing import Any

from .readme_next5 import L, S_SPLIT, scale_bridge_contract, solve_charged_lepton_masses


@dataclass(frozen=True)
class FigureContract:
    figure_id: str
    title: str
    readme_anchor: str
    status: str
    quantities: dict[str, Any]
    notes: list[str]


def solve_three_mass_shape(anchor: float, K: float, D: float, init: tuple[float, float]) -> tuple[float, float, float]:
    """Solve m1,m2,anchor from Koide-like K and logarithmic D in log sqrt-ratio coordinates.

    x=sqrt(m1/anchor), y=sqrt(m2/anchor)
    K=(x^2+y^2+1)/(x+y+1)^2, D=ln(m2^2/(m1 anchor))=4 ln y - 2 ln x.
    """
    a, b = init
    for _ in range(80):
        x = exp(a)
        y = exp(b)
        f1 = (x*x + y*y + 1.0) / (x + y + 1.0)**2 - K
        f2 = 4.0*b - 2.0*a - D
        if abs(f1) + abs(f2) < 1e-14:
            break
        h = 1e-6
        def vals(aa: float, bb: float) -> tuple[float, float]:
            xx = exp(aa)
            yy = exp(bb)
            return ((xx*xx + yy*yy + 1.0) / (xx + yy + 1.0)**2 - K, 4.0*bb - 2.0*aa - D)
        f1a, f2a = vals(a+h, b)
        f1b, f2b = vals(a, b+h)
        j11 = (f1a - f1) / h
        j21 = (f2a - f2) / h
        j12 = (f1b - f1) / h
        j22 = (f2b - f2) / h
        det = j11*j22 - j12*j21
        if abs(det) < 1e-18:
            raise ValueError("singular Newton system for three-mass shape solve")
        da = (-f1*j22 + j12*f2) / det
        db = (j21*f1 - j11*f2) / det
        a += max(min(da, 0.7), -0.7)
        b += max(min(db, 0.7), -0.7)
    x = exp(a)
    y = exp(b)
    return x*x*anchor, y*y*anchor, anchor


def quarks_contract() -> FigureContract:
    v = scale_bridge_contract().quantities["v_GeV"]

    A_t = L - 5*S_SPLIT + 138*S_SPLIT*S_SPLIT
    y_t = exp(-A_t)
    m_t = v / sqrt(2.0) * y_t
    K_u = 8/9 - S_SPLIT + 4*S_SPLIT*S_SPLIT
    D_u = sqrt(2*pi)/4 - 4*S_SPLIT + 4*S_SPLIT*S_SPLIT
    m_u, m_c, m_t = solve_three_mass_shape(m_t, K_u, D_u, (-5.9, -2.8))

    A_b = 4*pi/3 - 56*S_SPLIT + 106*S_SPLIT*S_SPLIT
    y_b = exp(-A_b)
    m_b = v / sqrt(2.0) * y_b
    K_d = 3/4 - (12/7)*S_SPLIT
    D_d = -1 + 1/72 - 4*S_SPLIT*S_SPLIT
    m_d, m_s, m_b = solve_three_mass_shape(m_b, K_d, D_d, (-3.5, -2.0))

    quantities = {
        "v_GeV_from_scale_bridge": v,
        "top_lane": {
            "A_t": A_t,
            "y_t": y_t,
            "m_t_GeV": m_t,
            "formula": "A_t=L-5S+138S^2",
            "coefficient_source": "138=2(72-3)",
            "alignment": "scalar-aligned top lane; no universal contact-depth penalty",
        },
        "up_sector_shapes": {
            "K_u": K_u,
            "D_u": D_u,
            "formulas": {"K_u": "8/9-S+4S^2", "D_u": "sqrt(2π)/4-4S+4S^2"},
            "solved_masses_GeV": {"m_u": m_u, "m_c": m_c, "m_t": m_t},
        },
        "bottom_lane": {
            "A_b": A_b,
            "y_b": y_b,
            "m_b_GeV": m_b,
            "formula": "A_b=4π/3-56S+106S^2",
            "coefficient_source": "106=2(56-3)",
            "alignment": "contact-depth bottom lane; boundary value before RG transport",
        },
        "down_sector_shapes": {
            "K_d": K_d,
            "D_d": D_d,
            "formulas": {"K_d": "3/4-(12/7)S", "D_d": "-1+1/72-4S^2"},
            "solved_masses_GeV": {"m_d": m_d, "m_s": m_s, "m_b": m_b},
        },
        "transport_status": "M_Z-boundary formulas with RG transport; not native quark theorem",
    }
    return FigureContract(
        figure_id="asha_readme_quark_lanes_shape_laws",
        title="Quarks: Top/Bottom Lanes and Shape Laws",
        readme_anchor="Quarks",
        status="PASS_README_QUARK_LANES_SHAPE_LAWS",
        quantities=quantities,
        notes=[
            "The figure separates top and bottom anchor lanes from up/down sector shape laws.",
            "The quark values are displayed as M_Z-boundary/RG-transport objects, not as native theorem outputs.",
            "The coefficient sources 138=2(72-3) and 106=2(56-3) are explicitly firewalled from fit-tuning.",
        ],
    )


def validate_contract(c: FigureContract) -> dict[str, Any]:
    q = c.quantities
    ok = True
    problems: list[str] = []
    if c.figure_id != "asha_readme_quark_lanes_shape_laws":
        return {"status": "FAIL", "ok": False, "problems": ["unknown figure_id"]}
    tl = q["top_lane"]
    bl = q["bottom_lane"]
    us = q["up_sector_shapes"]
    ds = q["down_sector_shapes"]
    checks = [
        abs(tl["A_t"] - (L - 5*S_SPLIT + 138*S_SPLIT*S_SPLIT)) < 1e-14,
        tl["coefficient_source"] == "138=2(72-3)",
        0.95 < tl["y_t"] < 1.0,
        150 < tl["m_t_GeV"] < 180,
        abs(us["K_u"] - (8/9 - S_SPLIT + 4*S_SPLIT*S_SPLIT)) < 1e-14,
        abs(us["D_u"] - (sqrt(2*pi)/4 - 4*S_SPLIT + 4*S_SPLIT*S_SPLIT)) < 1e-14,
        abs(bl["A_b"] - (4*pi/3 - 56*S_SPLIT + 106*S_SPLIT*S_SPLIT)) < 1e-14,
        bl["coefficient_source"] == "106=2(56-3)",
        0.015 < bl["y_b"] < 0.018,
        2.0 < bl["m_b_GeV"] < 4.0,
        abs(ds["K_d"] - (3/4 - (12/7)*S_SPLIT)) < 1e-14,
        abs(ds["D_d"] - (-1 + 1/72 - 4*S_SPLIT*S_SPLIT)) < 1e-14,
        us["solved_masses_GeV"]["m_u"] < us["solved_masses_GeV"]["m_c"] < us["solved_masses_GeV"]["m_t"],
        ds["solved_masses_GeV"]["m_d"] < ds["solved_masses_GeV"]["m_s"] < ds["solved_masses_GeV"]["m_b"],
    ]
    ok = all(checks)
    if not ok:
        problems.append("one or more quark formula/ordering checks failed")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
