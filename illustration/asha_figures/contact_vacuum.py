from __future__ import annotations

from dataclasses import dataclass, asdict
from math import log, pi, sqrt


@dataclass(frozen=True)
class RankRegion:
    key: str
    label: str
    rank: int
    log_area: float
    cx: float
    cy: float
    rx: float
    ry: float
    fill: str
    stroke: str
    opacity: float


@dataclass(frozen=True)
class VacuumGeometry:
    figure: str
    scale_area: float
    regions: list[RankRegion]
    ranks: dict[str, int]
    statements: list[str]


RANKS = {
    "lambda4_chamber": 70,
    "P_B": 56,
    "P_G": 14,
    "K_7": 7,
}


def ellipse_radii_for_rank(rank: int, scale_area: float, aspect: float) -> tuple[float, float, float]:
    """Return (area, rx, ry) with visual area proportional to ln(rank)."""
    log_area = log(rank)
    area = scale_area * log_area
    # area = pi * rx * ry; aspect = rx / ry
    ry = sqrt(area / (pi * aspect))
    rx = aspect * ry
    return area, rx, ry


def build_section2_geometry(width: int = 1800, height: int = 2400, scale_area: float = 165_000.0) -> VacuumGeometry:
    """Build a deterministic topological/logarithmic layout for PB, PG, and K7.

    Contract:
    - P_B, P_G, K_7 are not literal high-dimensional objects.
    - The visible areas are logarithmically scaled by rank.
    - K_7 is the brightest payload and is placed at the intersection.
    """
    p_area, p_rx, p_ry = ellipse_radii_for_rank(RANKS["P_B"], scale_area, aspect=1.64)
    g_area, g_rx, g_ry = ellipse_radii_for_rank(RANKS["P_G"], scale_area, aspect=1.18)
    k_area, k_rx, k_ry = ellipse_radii_for_rank(RANKS["K_7"], scale_area, aspect=1.0)

    # Slightly asymmetric lenses: Boolean support is broad; G2 support is sharper.
    regions = [
        RankRegion(
            key="P_B",
            label="P_B Boolean support lens",
            rank=RANKS["P_B"],
            log_area=log(RANKS["P_B"]),
            cx=790,
            cy=1170,
            rx=p_rx,
            ry=p_ry,
            fill="#CFF8FF",
            stroke="#E8EDF2",
            opacity=0.115,
        ),
        RankRegion(
            key="P_G",
            label="P_G G₂ / octonionic support lens",
            rank=RANKS["P_G"],
            log_area=log(RANKS["P_G"]),
            cx=1040,
            cy=1170,
            rx=g_rx,
            ry=g_ry,
            fill="#D9B45F",
            stroke="#FFE7A3",
            opacity=0.105,
        ),
        RankRegion(
            key="K_7",
            label="K_7 = Im(P_B) ∩ Im(P_G)",
            rank=RANKS["K_7"],
            log_area=log(RANKS["K_7"]),
            cx=930,
            cy=1170,
            rx=k_rx,
            ry=k_ry,
            fill="#F8F0C8",
            stroke="#FFFFFF",
            opacity=0.95,
        ),
    ]
    return VacuumGeometry(
        figure="section2_boolean_g2_contact_vacuum",
        scale_area=scale_area,
        regions=regions,
        ranks=dict(RANKS),
        statements=[
            "U = Im(P_B), rank 56",
            "V = Im(P_G), rank 14",
            "K_7 = U ∩ V, rank 7",
            "Visible areas are proportional to natural logarithm of rank.",
            "This is a topological/logarithmic visualization, not a literal embedding of high-dimensional projectors.",
        ],
    )


def validate_section2_geometry(geometry: VacuumGeometry, tolerance: float = 1e-9) -> dict:
    checks = {}
    observed = {r.key: pi * r.rx * r.ry for r in geometry.regions}
    expected = {r.key: geometry.scale_area * log(r.rank) for r in geometry.regions}
    ratios = {}
    for key in observed:
        ratios[key] = observed[key] / expected[key]
    checks["ranks_exact"] = geometry.ranks == RANKS
    checks["rank_chain_exact"] = (geometry.ranks["P_B"], geometry.ranks["P_G"], geometry.ranks["K_7"]) == (56, 14, 7)
    checks["log_area_scaling_exact"] = all(abs(ratios[k] - 1.0) <= tolerance for k in ratios)
    checks["area_values"] = observed
    checks["expected_area_values"] = expected
    checks["area_ratios_vs_K7"] = {
        "P_B_over_K_7": observed["P_B"] / observed["K_7"],
        "P_G_over_K_7": observed["P_G"] / observed["K_7"],
        "expected_P_B_over_K_7_ln_ratio": log(56) / log(7),
        "expected_P_G_over_K_7_ln_ratio": log(14) / log(7),
    }
    checks["status"] = "PASS_LOG_RANK_TOPOLOGICAL_CONTACT_VACUUM" if (
        checks["ranks_exact"] and checks["rank_chain_exact"] and checks["log_area_scaling_exact"]
    ) else "FAIL_SECTION2_GEOMETRY_CONTRACT"
    return checks


def export_section2_json_dict(geometry: VacuumGeometry) -> dict:
    return {
        "figure": geometry.figure,
        "mathematical_object": "Boolean/G2 contact-vacuum filtration in the Lambda^4 R8 chamber",
        "source_type": "topological/logarithmic diagram; exact ranks, nonliteral high-dimensional projector geometry",
        "ranks": geometry.ranks,
        "scale_rule": "visible ellipse/circle area = scale_area * ln(rank)",
        "validation": validate_section2_geometry(geometry),
        "regions": [asdict(r) for r in geometry.regions],
        "statements": geometry.statements,
    }
