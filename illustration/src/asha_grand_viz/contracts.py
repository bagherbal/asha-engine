from __future__ import annotations

from dataclasses import dataclass, asdict
from math import exp, pi, sqrt
from typing import Dict, List, Tuple

S_SPLIT = 0.0012924448188162962
L_HISTORY = 1.0 / (8.0 * pi)
PHI = (1.0 + sqrt(5.0)) / 2.0
E = exp(1.0)

COORDINATES = ["x0", "x1", "x2", "x3", "p0", "p1", "p2", "p3"]
ETA_ASHA = [1, -1, -1, -1, -1, -1, -1, -1]
GRADE_DIMENSIONS = [1, 8, 28, 56, 70, 56, 28, 8, 1]
SOURCE_ALPHABET = [3, 4, 7, 27, 56, 70, 72]
SOURCE_ALPHABET_ROLES = {
    3: "contact-depth triple / generations capacity",
    4: "response dimension / H_phi real scalar frame",
    7: "contact vacuum K7 / 1+6 contact chamber",
    27: "cubic depth volume",
    56: "Boolean incidence support rank",
    70: "middle exterior chamber dim Λ4 R8",
    72: "augmented chamber / generation-mixing map count",
}

N_Q = [1.0 / 3.0, 1.0 / 2.0, 2.0 / 3.0]
W_Q = [exp(-4.0 * pi * n) for n in N_Q]

ACTION_SECTORS = [
    {"symbol": "S_grav", "role": "metric dynamics / Planck stiffness firewall"},
    {"symbol": "S_gauge", "role": "inner-fluctuation gauge curvature"},
    {"symbol": "S_Higgs^ASHA", "role": "quartic scalar lane"},
    {"symbol": "S_fermion", "role": "spinor kinetic carrier"},
    {"symbol": "S_Yukawa^ASHA", "role": "sealed socket/depth amplitudes"},
    {"symbol": "S_nu^seesaw", "role": "Majorana/Takagi seesaw lane"},
]

THEOREM_LOCKS = [
    "Lorentzianized phase-space octave",
    "Flat metric projection",
    "Contact-seven source",
    "Contact phase-triple algebra",
    "Product-depth extension",
    "Yukawa broadcast breaking",
    "Relative flavor orientation",
    "Majorana/Takagi seesaw structure",
    "Low-energy metric dynamics",
    "Dimensional obstruction",
    "Vacuum-zero independence",
]

WOUNDS = [
    "M_P^2 Planck stiffness",
    "Λ_cosmo vacuum residual",
    "PMNS/Majorana selector",
    "SocketLaneSelector",
    "MatterContactUniversalityPrinciple",
]

@dataclass(frozen=True)
class FormulaStack:
    L: float
    S: float
    v_over_mp: float
    lambda_asha: float
    higgs_multiplier: float
    A_tau: float
    y_tau: float
    A_t: float
    y_t: float
    A_b: float
    y_b: float
    theta12_0: float
    theta23_0: float
    theta13_0: float
    delta_ckm_0: float
    M_R3_prefactor_without_sqrt_vmp: float
    m2_over_m3: float


def formula_stack() -> FormulaStack:
    S = S_SPLIT
    L = L_HISTORY
    v_over_mp = exp(-12.0 * pi + sqrt(3.0) / 2.0 + 2.0 * S + 148.0 * S * S)
    lambda_asha = 3.0 / 8.0 * (1.0 + L) * (1.0 / 3.0 - S)
    higgs_multiplier = sqrt(2.0 * lambda_asha)
    A_tau = 4.0 * pi / 3.0 + 3.0 / 10.0 + 7.0 / 72.0 - S + (72.0 + 27.0) * S * S / 2.0
    A_t = L - 5.0 * S + 138.0 * S * S
    A_b = 4.0 * pi / 3.0 - 56.0 * S + 106.0 * S * S
    theta12 = 1.0 / 4.0 - 18.0 * S + 158.0 * S * S
    theta23 = L + 5.0 * S / 3.0 - (8.0 - 2.0 * L) * S * S
    theta13 = 72.0 * L * S - 1.5 * S * S
    delta = pi / 3.0 + 71.0 * S + 93.0 * S * S / 4.0
    return FormulaStack(
        L=L,
        S=S,
        v_over_mp=v_over_mp,
        lambda_asha=lambda_asha,
        higgs_multiplier=higgs_multiplier,
        A_tau=A_tau,
        y_tau=exp(-A_tau),
        A_t=A_t,
        y_t=exp(-A_t),
        A_b=A_b,
        y_b=exp(-A_b),
        theta12_0=theta12,
        theta23_0=theta23,
        theta13_0=theta13,
        delta_ckm_0=delta,
        M_R3_prefactor_without_sqrt_vmp=sqrt(2.0 * pi) + 49.0 * S + 90.0 * S * S,
        m2_over_m3=4.0 * L + 10.0 * S,
    )


def contract_dict() -> Dict:
    fs = formula_stack()
    return {
        "source": "ASHA uploaded zip: README, formula_ledger, essential_ontological_tower_map, gate1349, selected gate audits",
        "ontology_order": [
            "1 scalar identity",
            "V8 = X4 ⊕ P4 measurement octave",
            "η_ASHA = η_1,3 ⊕ (-I4), Ω = Σ dpμ∧dxμ",
            "ΠX metric projection to X4",
            "x0 observer reference + contact seven p0 ⊕ Π1 ⊕ Π2 ⊕ Π3",
            "Boolean/G2 projector intersection K7",
            "A_F = C ⊕ H ⊕ M3(C), product-depth extension",
            "inner fluctuations / gauge + Higgs inventory",
            "scale, Higgs, matter, mixing and seesaw bridge/filling lanes",
            "six-sector low-energy action skeleton",
            "firewalled observation and wounds",
        ],
        "coordinates": COORDINATES,
        "eta_diag": ETA_ASHA,
        "grade_dimensions": GRADE_DIMENSIONS,
        "source_alphabet": SOURCE_ALPHABET,
        "source_alphabet_roles": SOURCE_ALPHABET_ROLES,
        "N_Q": N_Q,
        "W_Q": W_Q,
        "action_sectors": ACTION_SECTORS,
        "theorem_locks": THEOREM_LOCKS,
        "wounds": WOUNDS,
        "formula_stack": asdict(fs),
        "claimed_source_types": {
            "theorem": THEOREM_LOCKS,
            "physical_filling": ["scale bridge", "Higgs", "tau/top/bottom anchors", "CKM boundary", "rank-2 neutrino lane"],
            "unresolved_wound": WOUNDS,
        },
    }
