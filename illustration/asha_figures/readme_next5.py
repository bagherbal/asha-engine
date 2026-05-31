from __future__ import annotations

from dataclasses import dataclass, asdict
from math import exp, pi, sqrt, log
from typing import Any

L = 1.0 / (8.0 * pi)
S_SPLIT = 0.0012924448188162962
M_PLANCK_REDUCED_GEV = 2.435e18

@dataclass(frozen=True)
class FigureContract:
    figure_id: str
    title: str
    readme_anchor: str
    status: str
    quantities: dict[str, Any]
    notes: list[str]


def depth_weights() -> list[float]:
    return [exp(-4*pi/3), exp(-2*pi), exp(-8*pi/3)]


def matter_socket_contract() -> FigureContract:
    quantities = {
        "A_F_blocks": ["C", "H", "M3(C)"],
        "block_cells": {"C": 1, "H": 4, "M3(C)": 9},
        "contact_depth_layers": ["Pi1", "Pi2", "Pi3"],
        "N_Q": [1/3, 1/2, 2/3],
        "W_Q": depth_weights(),
        "formal_extension": "A_depth = A_F ⊗ Q_contact^3",
        "operator": "D_Y^depth = D_Y^F ⊗ exp(-4π N_Q)",
        "top_lane_sigma": 0,
        "contact_lane_sigma_examples": {"bottom": 1, "tau": 1},
    }
    return FigureContract(
        figure_id="asha_readme_matter_sockets_product_depth",
        title="Matter Sockets and Product Depth",
        readme_anchor="Matter sockets and product depth",
        status="PASS_README_MATTER_SOCKETS_PRODUCT_DEPTH",
        quantities=quantities,
        notes=[
            "The figure separates finite Standard Model socket structure from contact-depth labels.",
            "The 1,4,9 cell counts visualize C, H, and M3(C); the depth triple is broadcast as a separate three-layer factor.",
            "No observed masses, particle assignments, CKM/PMNS, or flavor theorem are encoded here.",
        ],
    )


def source_alphabet_contract() -> FigureContract:
    quantities = {
        "L": L,
        "S_split": S_SPLIT,
        "finite_source_numbers": [3,4,7,27,56,70,72],
        "typed_coefficients": {
            "158": "2(72+7)",
            "148": "2·72+4",
            "106": "2(56-3)",
        },
    }
    return FigureContract(
        figure_id="asha_readme_locked_constants_source_alphabet",
        title="Locked Constants and Source Alphabet",
        readme_anchor="Locked constants and source alphabet",
        status="PASS_README_LOCKED_CONSTANTS_SOURCE_ALPHABET",
        quantities=quantities,
        notes=[
            "The figure is a coefficient firewall: numbers are allowed only when source-typed.",
            "L and S are displayed as currently locked formula alphabet entries.",
            "This is not a fit board; it is a source-typing ledger for later formula use.",
        ],
    )


def scale_bridge_contract() -> FigureContract:
    exponent = -12*pi + sqrt(3)/2 + 2*S_SPLIT + 148*S_SPLIT*S_SPLIT
    ratio = exp(exponent)
    v = M_PLANCK_REDUCED_GEV * ratio
    quantities = {
        "M_P_reduced_GeV": M_PLANCK_REDUCED_GEV,
        "exponent": exponent,
        "ratio_v_over_Mp": ratio,
        "v_GeV": v,
        "terms": {
            "three_action_quantum": -12*pi,
            "triadic_amplitude": sqrt(3)/2,
            "wall_response": 2*S_SPLIT,
            "augmented_chamber": 148*S_SPLIT*S_SPLIT,
        },
    }
    return FigureContract(
        figure_id="asha_readme_planck_to_electroweak_scale_bridge",
        title="Planck-to-Electroweak Scale Bridge",
        readme_anchor="Planck-to-electroweak scale bridge",
        status="PASS_README_PLANCK_TO_ELECTROWEAK_SCALE_BRIDGE",
        quantities=quantities,
        notes=[
            "The vertical axis is logarithmic: the visual drop is the exponential suppression from the Planck stiffness seal.",
            "M_P is treated as a physical filling / Planck stiffness seal, not a native theorem.",
            "The bridge reduces v to M_P plus source-typed action structure.",
        ],
    )


def higgs_contract() -> FigureContract:
    lam = (3/8) * (1 + L) * (1/3 - S_SPLIT)
    v = scale_bridge_contract().quantities["v_GeV"]
    mh = v * sqrt(2*lam)
    quantities = {
        "lambda_ASHA": lam,
        "v_GeV_from_scale_bridge": v,
        "m_H_GeV": mh,
        "formula_lambda": "3/8(1+L)(1/3-S)",
        "formula_mH": "v sqrt(2 lambda_ASHA)",
        "potential": "lambda(|phi|^2-v^2/2)^2",
    }
    return FigureContract(
        figure_id="asha_readme_higgs_sector_quartic_mass_chain",
        title="Higgs Sector: Quartic and Mass Chain",
        readme_anchor="Higgs sector",
        status="PASS_README_HIGGS_SECTOR_QUARTIC_MASS_CHAIN",
        quantities=quantities,
        notes=[
            "The figure shows the formula chain from L and S to lambda_ASHA and m_H.",
            "The Mexican-hat curve is a visual potential scaffold, while the numeric output follows the README formula.",
            "This is a locked physical-filling chain, not a native scalar theorem.",
        ],
    )


def solve_charged_lepton_masses(m_tau: float, K: float, D: float) -> tuple[float, float, float]:
    # Solve for x=sqrt(me/mtau), y=sqrt(mmu/mtau). Use stable damped Newton in log-space.
    a, b = -4.08, -1.41
    for _ in range(40):
        x = exp(a); y = exp(b)
        f1 = (x*x + y*y + 1) / (x+y+1)**2 - K
        f2 = 4*b - 2*a - D
        if abs(f1) + abs(f2) < 1e-14:
            break
        h = 1e-6
        def vals(aa: float, bb: float) -> tuple[float, float]:
            xx = exp(aa); yy = exp(bb)
            return ((xx*xx + yy*yy + 1) / (xx+yy+1)**2 - K, 4*bb - 2*aa - D)
        f1a, f2a = vals(a+h, b)
        f1b, f2b = vals(a, b+h)
        j11 = (f1a-f1)/h; j21 = (f2a-f2)/h
        j12 = (f1b-f1)/h; j22 = (f2b-f2)/h
        det = j11*j22 - j12*j21
        da = (-f1*j22 + j12*f2) / det
        db = (j21*f1 - j11*f2) / det
        # Damping avoids accidental overshoot if constants change.
        a += max(min(da, 0.7), -0.7)
        b += max(min(db, 0.7), -0.7)
    x = exp(a); y = exp(b)
    return x*x*m_tau, y*y*m_tau, m_tau


def charged_leptons_contract() -> FigureContract:
    v = scale_bridge_contract().quantities["v_GeV"]
    A_tau = 4*pi/3 + 3/10 + 7/72 - S_SPLIT + 0.5*(72+27)*S_SPLIT*S_SPLIT
    m_tau = v / sqrt(2) * exp(-A_tau)
    K_e = 2/3 - 4*(1-2*L)*S_SPLIT*S_SPLIT
    D_e = sqrt(2*pi) + 2*S_SPLIT - 4*(1-L)*S_SPLIT*S_SPLIT
    m_e, m_mu, m_tau = solve_charged_lepton_masses(m_tau, K_e, D_e)
    quantities = {
        "A_tau": A_tau,
        "m_tau_GeV": m_tau,
        "K_e": K_e,
        "D_e": D_e,
        "solved_masses_GeV": {"m_e": m_e, "m_mu": m_mu, "m_tau": m_tau},
        "shape_laws": {
            "K_e": "2/3 - 4(1-2L)S^2",
            "D_e": "sqrt(2π)+2S-4(1-L)S^2",
        },
    }
    return FigureContract(
        figure_id="asha_readme_charged_lepton_anchor_shape_laws",
        title="Charged Leptons: Tau Anchor and Shape Laws",
        readme_anchor="Charged leptons",
        status="PASS_README_CHARGED_LEPTON_ANCHOR_SHAPE_LAWS",
        quantities=quantities,
        notes=[
            "The tau anchor fixes the absolute charged-lepton scale after v; K_e and D_e fix the shape.",
            "The mass nodes are displayed as a logarithmic ladder because the hierarchy is multiplicative.",
            "This is a locked physical-filling formula stack, not a native theorem or flavor-selector proof.",
        ],
    )


def next5_contracts() -> list[FigureContract]:
    return [
        matter_socket_contract(),
        source_alphabet_contract(),
        scale_bridge_contract(),
        higgs_contract(),
        charged_leptons_contract(),
    ]


def validate_contract(c: FigureContract) -> dict[str, Any]:
    q = c.quantities
    ok = True
    problems: list[str] = []
    if c.figure_id == "asha_readme_matter_sockets_product_depth":
        ok &= q["block_cells"] == {"C": 1, "H": 4, "M3(C)": 9}
        ok &= q["N_Q"] == [1/3, 1/2, 2/3]
        ok &= len(q["W_Q"]) == 3 and all(a > b for a,b in zip(q["W_Q"], q["W_Q"][1:]))
    elif c.figure_id == "asha_readme_locked_constants_source_alphabet":
        ok &= abs(q["L"] - L) < 1e-15
        ok &= q["finite_source_numbers"] == [3,4,7,27,56,70,72]
        ok &= q["typed_coefficients"]["158"] == "2(72+7)"
    elif c.figure_id == "asha_readme_planck_to_electroweak_scale_bridge":
        ok &= abs(q["exponent"] - (-12*pi + sqrt(3)/2 + 2*S_SPLIT + 148*S_SPLIT*S_SPLIT)) < 1e-13
        ok &= 240 < q["v_GeV"] < 252
    elif c.figure_id == "asha_readme_higgs_sector_quartic_mass_chain":
        ok &= 0.12 < q["lambda_ASHA"] < 0.14
        ok &= 124 < q["m_H_GeV"] < 127
    elif c.figure_id == "asha_readme_charged_lepton_anchor_shape_laws":
        masses = q["solved_masses_GeV"]
        ok &= masses["m_e"] < masses["m_mu"] < masses["m_tau"]
        ok &= 0.66 < q["K_e"] < 0.67
        ok &= 2.50 < q["D_e"] < 2.52
    else:
        ok = False
        problems.append("unknown figure_id")
    return {"status": c.status if ok else "FAIL", "ok": bool(ok), "problems": problems}


def export_contract(c: FigureContract) -> dict[str, Any]:
    return asdict(c)
