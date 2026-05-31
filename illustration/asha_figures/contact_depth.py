from __future__ import annotations

from dataclasses import dataclass, asdict
from fractions import Fraction
from math import exp, pi


@dataclass(frozen=True)
class BasisPoint:
    key: str
    label: str
    role: str
    x: float
    y: float
    color_role: str
    in_contact: bool


@dataclass(frozen=True)
class PhasePlane:
    key: str
    label: str
    basis: tuple[str, str]
    depth_fraction: str
    depth_value: float
    weight_formula: str
    weight_value: float
    cx: float
    cy: float
    rx: float
    ry: float


@dataclass(frozen=True)
class ContactDepthGeometry:
    figure: str
    source_type: str
    v8_basis: tuple[str, ...]
    observer_reference: str
    contact_basis: tuple[str, ...]
    decomposition: tuple[str, ...]
    q_contact_dim: int
    n_q: tuple[str, str, str]
    w_q: tuple[float, float, float]
    basis_points: list[BasisPoint]
    phase_planes: list[PhasePlane]
    statements: list[str]


V8_BASIS = ("x0", "p0", "x1", "p1", "x2", "p2", "x3", "p3")
OBSERVER_REFERENCE = "x0"
CONTACT_BASIS = ("p0", "x1", "p1", "x2", "p2", "x3", "p3")
DECOMPOSITION = ("R p0", "Pi1=span(x1,p1)", "Pi2=span(x2,p2)", "Pi3=span(x3,p3)")
N_Q_FRACTIONS = (Fraction(1, 3), Fraction(1, 2), Fraction(2, 3))


def _frac_label(f: Fraction) -> str:
    return f"{f.numerator}/{f.denominator}"


def _w_value(f: Fraction) -> float:
    return exp(-4.0 * pi * float(f))


def build_contact_depth_geometry(width: int = 1800, height: int = 2400) -> ContactDepthGeometry:
    """Build README figure geometry for ASHA's contact seven and depth triple.

    Contract:
    - Start with the measurement octave V8 = X4 ⊕ P4.
    - Select x0 as observer-time reference.
    - The contact carrier is exactly the remaining seven basis directions.
    - The contact carrier decomposes as R p0 plus three spatial phase planes.
    - The depth operator is exactly diag(1/3, 1/2, 2/3).
    - The depth weights are exactly exp(-4π n_i) numerically evaluated for labels.
    """
    left_x = 285
    y0 = 575
    pair_gap = 305
    pair_sep = 70
    basis_points: list[BasisPoint] = []
    for i in range(4):
        x_label = f"x{i}"
        p_label = f"p{i}"
        cy = y0 + i * pair_gap
        basis_points.append(BasisPoint(
            key=x_label,
            label=x_label.replace("0", "⁰").replace("1", "¹").replace("2", "²").replace("3", "³"),
            role="observer-time reference" if x_label == "x0" else f"event direction {i}",
            x=left_x - pair_sep,
            y=cy,
            color_role="time" if x_label == "x0" else "space_event",
            in_contact=x_label != OBSERVER_REFERENCE,
        ))
        basis_points.append(BasisPoint(
            key=p_label,
            label=p_label.replace("0", "₀").replace("1", "₁").replace("2", "₂").replace("3", "₃"),
            role="energy response" if p_label == "p0" else f"momentum response {i}",
            x=left_x + pair_sep,
            y=cy,
            color_role="energy" if p_label == "p0" else "space_response",
            in_contact=True,
        ))

    plane_x = 895
    plane_ys = (1030, 1320, 1610)
    depth_labels = tuple(_frac_label(f) for f in N_Q_FRACTIONS)
    weights = tuple(_w_value(f) for f in N_Q_FRACTIONS)
    phase_planes = [
        PhasePlane(
            key=f"Pi{i}",
            label=f"Π{i}",
            basis=(f"x{i}", f"p{i}"),
            depth_fraction=depth_labels[i - 1],
            depth_value=float(N_Q_FRACTIONS[i - 1]),
            weight_formula=f"exp(-4π·{depth_labels[i - 1]})",
            weight_value=weights[i - 1],
            cx=plane_x,
            cy=plane_ys[i - 1],
            rx=240,
            ry=82,
        )
        for i in (1, 2, 3)
    ]

    return ContactDepthGeometry(
        figure="readme_contact_seven_depth_triple",
        source_type="README visual theorem map; exact carrier/depth labels, conceptual placement",
        v8_basis=V8_BASIS,
        observer_reference=OBSERVER_REFERENCE,
        contact_basis=CONTACT_BASIS,
        decomposition=DECOMPOSITION,
        q_contact_dim=3,
        n_q=depth_labels,
        w_q=weights,
        basis_points=basis_points,
        phase_planes=phase_planes,
        statements=[
            "V8 = X4 ⊕ P4 with basis x0,p0,x1,p1,x2,p2,x3,p3.",
            "Selecting x0 as observer-time reference leaves exactly seven contact directions.",
            "V7_contact = R p0 ⊕ Pi1 ⊕ Pi2 ⊕ Pi3.",
            "Pi_i = span(x^i,p_i) for i=1,2,3.",
            "Q_contact^3 ≅ C^3 with N_Q = diag(1/3,1/2,2/3).",
            "W_Q = exp(-4π N_Q) gives the contact-depth weights.",
        ],
    )


def validate_contact_depth_geometry(geometry: ContactDepthGeometry, tolerance: float = 1e-15) -> dict:
    expected_weights = tuple(_w_value(f) for f in N_Q_FRACTIONS)
    checks = {
        "v8_basis_exact": geometry.v8_basis == V8_BASIS,
        "observer_reference_exact": geometry.observer_reference == OBSERVER_REFERENCE,
        "contact_basis_exact": geometry.contact_basis == CONTACT_BASIS,
        "contact_basis_count_exact": len(geometry.contact_basis) == 7,
        "observer_removed_from_contact": geometry.observer_reference not in geometry.contact_basis,
        "decomposition_exact": geometry.decomposition == DECOMPOSITION,
        "q_contact_dim_exact": geometry.q_contact_dim == 3,
        "n_q_exact": geometry.n_q == tuple(_frac_label(f) for f in N_Q_FRACTIONS),
        "w_q_values": geometry.w_q,
        "w_q_expected_values": expected_weights,
        "w_q_numerically_exact": all(abs(a - b) <= tolerance for a, b in zip(geometry.w_q, expected_weights)),
        "w_q_strictly_descending": geometry.w_q[0] > geometry.w_q[1] > geometry.w_q[2],
        "phase_planes_exact": [p.basis for p in geometry.phase_planes] == [("x1", "p1"), ("x2", "p2"), ("x3", "p3")],
    }
    checks["status"] = "PASS_README_CONTACT_SEVEN_DEPTH_TRIPLE" if all(
        v for k, v in checks.items() if k not in {"w_q_values", "w_q_expected_values"}
    ) else "FAIL_README_CONTACT_SEVEN_DEPTH_TRIPLE_CONTRACT"
    return checks


def export_contact_depth_json_dict(geometry: ContactDepthGeometry) -> dict:
    return {
        "figure": geometry.figure,
        "mathematical_object": "Contact seven and contact-depth triple",
        "source_type": geometry.source_type,
        "basis": {
            "V8": geometry.v8_basis,
            "observer_reference": geometry.observer_reference,
            "V7_contact": geometry.contact_basis,
            "decomposition": geometry.decomposition,
        },
        "depth_operator": {
            "Q_contact_dim": geometry.q_contact_dim,
            "N_Q_diag": geometry.n_q,
            "W_Q_diag": geometry.w_q,
        },
        "basis_points": [asdict(p) for p in geometry.basis_points],
        "phase_planes": [asdict(p) for p in geometry.phase_planes],
        "statements": geometry.statements,
        "validation": validate_contact_depth_geometry(geometry),
    }
