from __future__ import annotations

from dataclasses import dataclass, asdict
from itertools import combinations
from math import comb
from typing import Iterable

@dataclass(frozen=True)
class LatticeNode:
    mask: int
    grade: int
    index_in_grade: int
    basis: str
    has_time: bool
    x: float
    y: float
    radius: float

@dataclass(frozen=True)
class LatticeEdge:
    source: int
    target: int
    added_generator: int
    warm: bool

GENERATOR_NAMES = ["e₀", "e₁", "e₂", "e₃", "e₄", "e₅", "e₆", "e₇"]
TIER_COUNTS = [comb(8, k) for k in range(9)]
EXPECTED_TIER_COUNTS = [1, 8, 28, 56, 70, 56, 28, 8, 1]


def popcount(mask: int) -> int:
    return int(mask).bit_count()


def mask_to_basis(mask: int) -> str:
    if mask == 0:
        return "1"
    return "∧".join(GENERATOR_NAMES[i] for i in range(8) if mask & (1 << i))


def masks_by_grade(n: int = 8) -> list[list[int]]:
    tiers: list[list[int]] = []
    for k in range(n + 1):
        masks = []
        for combo in combinations(range(n), k):
            mask = 0
            for i in combo:
                mask |= 1 << i
            masks.append(mask)
        # Lexicographic by generator tuple; stable and reproducible.
        tiers.append(masks)
    return tiers


def build_nodes(width: int, margin_x: int, tier_top: int, tier_bottom: int,
                node_radius: float, middle_node_radius: float,
                terminal_node_radius: float) -> list[LatticeNode]:
    tiers = masks_by_grade(8)
    assert [len(t) for t in tiers] == EXPECTED_TIER_COUNTS
    nodes: list[LatticeNode] = []
    inner_w = width - 2 * margin_x
    dy = (tier_bottom - tier_top) / 8
    for grade, masks in enumerate(tiers):
        y = tier_top + grade * dy
        n = len(masks)
        for i, mask in enumerate(masks):
            x = margin_x + inner_w * (i + 0.5) / n
            if grade in (0, 8):
                r = terminal_node_radius
            elif grade == 4:
                r = middle_node_radius
            else:
                r = node_radius
            nodes.append(LatticeNode(
                mask=mask,
                grade=grade,
                index_in_grade=i,
                basis=mask_to_basis(mask),
                has_time=bool(mask & 1),
                x=x,
                y=y,
                radius=r,
            ))
    return nodes


def build_edges(nodes: Iterable[LatticeNode]) -> list[LatticeEdge]:
    by_mask = {n.mask: n for n in nodes}
    edges: list[LatticeEdge] = []
    for node in by_mask.values():
        for bit in range(8):
            if not (node.mask & (1 << bit)):
                target = node.mask | (1 << bit)
                if target in by_mask:
                    edges.append(LatticeEdge(
                        source=node.mask,
                        target=target,
                        added_generator=bit,
                        warm=(bit == 0),
                    ))
    return edges


def validate(nodes: list[LatticeNode], edges: list[LatticeEdge]) -> dict:
    tier_counts = [0] * 9
    for n in nodes:
        tier_counts[n.grade] += 1
    return {
        "node_count": len(nodes),
        "edge_count": len(edges),
        "tier_counts": tier_counts,
        "expected_tier_counts": EXPECTED_TIER_COUNTS,
        "tier_counts_exact": tier_counts == EXPECTED_TIER_COUNTS,
        "expected_node_count": 256,
        "node_count_exact": len(nodes) == 256,
        "expected_edge_count_boolean_lattice_B8": 8 * (2 ** 7),
        "edge_count_exact": len(edges) == 8 * (2 ** 7),
        "time_like_generator": "e₀, η₀₀=+1",
        "space_like_generators": "e₁…e₇, ηᵢᵢ=-1",
    }


def export_json_dict(nodes: list[LatticeNode], edges: list[LatticeEdge]) -> dict:
    return {
        "figure": "section1_measurement_ladder_cl_1_7",
        "mathematical_object": "Boolean/exterior-grade basis ladder for Cℓ(1,7) over an 8-generator carrier",
        "signature": {"positive": 1, "negative": 7, "metric_diagonal": [1, -1, -1, -1, -1, -1, -1, -1]},
        "validation": validate(nodes, edges),
        "nodes": [asdict(n) for n in nodes],
        "edges": [asdict(e) for e in edges],
    }
