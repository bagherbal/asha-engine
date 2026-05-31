
"""Programmatic ASHA projector-universe phase-space snapshot.

This module does not hand-place the visible trajectory.  It computes the
canonical V8 = X4 ⊕ P4 phase-space flow, applies exact coordinate projectors,
and exports the numerical trajectory plus validation invariants for rendering.
"""

from __future__ import annotations

import json
import math
from dataclasses import dataclass, asdict
from pathlib import Path
from typing import Any

import numpy as np


FIGURE_ID = "asha_projector_universe_x4_p4_one_trajectory"
WIDTH = 2400
HEIGHT = 1350


@dataclass(frozen=True)
class PhaseSpaceContract:
    coordinate_order: tuple[str, ...]
    eta_diagonal: tuple[float, ...]
    omega_matrix: list[list[float]]
    pi_x_matrix: list[list[float]]
    pi_p_matrix: list[list[float]]
    frequencies: tuple[float, ...]
    depth_N: tuple[float, float, float]
    depth_W: tuple[float, float, float]
    amplitudes_raw: tuple[float, ...]
    amplitudes_visual: tuple[float, ...]
    hamiltonian_kind: str


def canonical_matrices() -> dict[str, np.ndarray]:
    """Return η, Ω, Π_X, Π_P for coordinate order x0..x3,p0..p3."""
    I4 = np.eye(4)
    zero = np.zeros((4, 4))
    eta = np.diag([1, -1, -1, -1, -1, -1, -1, -1]).astype(float)
    # Ω = Σ dp_mu ∧ dx^mu, represented in x,p block order.
    omega = np.block([[zero, I4], [-I4, zero]]).astype(float)
    pi_x = np.diag([1, 1, 1, 1, 0, 0, 0, 0]).astype(float)
    pi_p = np.diag([0, 0, 0, 0, 1, 1, 1, 1]).astype(float)
    return {"eta": eta, "omega": omega, "pi_x": pi_x, "pi_p": pi_p}


def compute_flow(samples: int = 1800) -> dict[str, Any]:
    """Compute the canonical Hamiltonian movement of the seed One.

    The seed is visualized as one symplectic orbit whose four x_mu/p_mu
    oscillator pairs encode observer-time plus the three contact-depth lanes.

    Hamiltonian:
        H = 1/2 Σ_mu (p_mu² + (ω_mu x_mu)²)

    Exact solution:
        x_mu(t) = a_mu cos(ω_mu t + φ_mu)
        p_mu(t) = -a_mu ω_mu sin(ω_mu t + φ_mu)
    """
    mats = canonical_matrices()
    depth_N = np.array([1 / 3, 1 / 2, 2 / 3], dtype=float)
    depth_W = np.exp(-4 * math.pi * depth_N)
    frequencies = np.array([1.0, *depth_N], dtype=float)

    # Raw amplitudes obey the ASHA depth-drop scale.  The raw values are saved.
    amplitudes_raw = np.array([1.0, *depth_W], dtype=float)

    # Visual amplitudes use log compression so all three lanes remain visible.
    # This is a rendering transform only; the raw Hamiltonian invariant is
    # computed from the physical/raw amplitudes.
    amplitudes_visual = np.array(
        [1.0, 0.72, 0.52, 0.38], dtype=float
    )

    phases = np.array([0.0, math.tau / 7, 2 * math.tau / 7, 3 * math.tau / 7])
    t = np.linspace(0, 12 * math.pi, samples)

    def make_trajectory(amplitudes: np.ndarray) -> tuple[np.ndarray, np.ndarray, np.ndarray]:
        x = np.zeros((samples, 4), dtype=float)
        p = np.zeros((samples, 4), dtype=float)
        for mu in range(4):
            x[:, mu] = amplitudes[mu] * np.cos(frequencies[mu] * t + phases[mu])
            p[:, mu] = -amplitudes[mu] * frequencies[mu] * np.sin(
                frequencies[mu] * t + phases[mu]
            )
        z = np.hstack([x, p])
        H_components = 0.5 * (p**2 + (frequencies[None, :] * x) ** 2)
        H_total = H_components.sum(axis=1)
        return z, H_components, H_total

    z_raw, H_components_raw, H_total_raw = make_trajectory(amplitudes_raw)
    z_visual, H_components_visual, H_total_visual = make_trajectory(amplitudes_visual)

    # Exact projectors.
    z_x_raw = z_raw @ mats["pi_x"].T
    z_p_raw = z_raw @ mats["pi_p"].T
    z_x_visual = z_visual @ mats["pi_x"].T
    z_p_visual = z_visual @ mats["pi_p"].T

    # Deterministic 2D projections for rendering.  These are saved in manifest.
    # X projection sees observer-time and spatial contact spread.
    X2 = np.array(
        [
            [0.86, -0.42, 0.28, -0.12, 0, 0, 0, 0],
            [0.18, 0.50, -0.74, 0.39, 0, 0, 0, 0],
        ],
        dtype=float,
    )
    # P projection is analogous but acts only on momentum coordinates.
    P2 = np.array(
        [
            [0, 0, 0, 0, 0.70, -0.56, 0.37, -0.22],
            [0, 0, 0, 0, 0.16, 0.61, -0.68, 0.37],
        ],
        dtype=float,
    )
    # Unified phase projector combines X and P lanes into one central orbit.
    V2 = np.array(
        [
            [0.72, -0.26, 0.18, -0.10, 0.00, 0.43, -0.28, 0.16],
            [0.08, 0.35, -0.51, 0.31, 0.78, -0.22, 0.16, -0.10],
        ],
        dtype=float,
    )

    x2 = (z_visual @ X2.T)
    p2 = (z_visual @ P2.T)
    v2 = (z_visual @ V2.T)

    # Normalize deterministic projections for canvas placement.
    def normalize(points: np.ndarray) -> np.ndarray:
        q = points.copy()
        max_abs = np.max(np.abs(q), axis=0)
        max_abs[max_abs == 0] = 1.0
        return q / max_abs

    x2n = normalize(x2)
    p2n = normalize(p2)
    v2n = normalize(v2)

    validation = validate_flow(mats, z_raw, H_total_raw, depth_W)

    contract = PhaseSpaceContract(
        coordinate_order=("x0", "x1", "x2", "x3", "p0", "p1", "p2", "p3"),
        eta_diagonal=tuple(float(v) for v in np.diag(mats["eta"])),
        omega_matrix=mats["omega"].astype(float).tolist(),
        pi_x_matrix=mats["pi_x"].astype(float).tolist(),
        pi_p_matrix=mats["pi_p"].astype(float).tolist(),
        frequencies=tuple(float(v) for v in frequencies),
        depth_N=tuple(float(v) for v in depth_N),
        depth_W=tuple(float(v) for v in depth_W),
        amplitudes_raw=tuple(float(v) for v in amplitudes_raw),
        amplitudes_visual=tuple(float(v) for v in amplitudes_visual),
        hamiltonian_kind="H = 1/2 sum_mu (p_mu^2 + (omega_mu x_mu)^2)",
    )

    return {
        "figure": FIGURE_ID,
        "width": WIDTH,
        "height": HEIGHT,
        "contract": asdict(contract),
        "projection_matrices": {
            "X2": X2.tolist(),
            "P2": P2.tolist(),
            "V2": V2.tolist(),
        },
        "samples": samples,
        "time_domain": [float(t[0]), float(t[-1])],
        "raw": {
            "trajectory": z_raw.tolist(),
            "hamiltonian_total": H_total_raw.tolist(),
            "hamiltonian_components": H_components_raw.tolist(),
        },
        "visual": {
            "trajectory": z_visual.tolist(),
            "x_projection_2d": x2n.tolist(),
            "p_projection_2d": p2n.tolist(),
            "v_projection_2d": v2n.tolist(),
        },
        "validation": validation,
    }


def validate_flow(
    mats: dict[str, np.ndarray],
    z_raw: np.ndarray,
    H_total_raw: np.ndarray,
    depth_W: np.ndarray,
) -> dict[str, Any]:
    eta = mats["eta"]
    omega = mats["omega"]
    pi_x = mats["pi_x"]
    pi_p = mats["pi_p"]
    I8 = np.eye(8)

    eig_eta = np.linalg.eigvalsh(eta)
    signature = {
        "positive": int(np.sum(eig_eta > 1e-12)),
        "negative": int(np.sum(eig_eta < -1e-12)),
        "zero": int(np.sum(np.abs(eig_eta) <= 1e-12)),
    }

    checks = {
        "pi_x_idempotent_norm": float(np.linalg.norm(pi_x @ pi_x - pi_x)),
        "pi_p_idempotent_norm": float(np.linalg.norm(pi_p @ pi_p - pi_p)),
        "pi_x_pi_p_orthogonal_norm": float(np.linalg.norm(pi_x @ pi_p)),
        "pi_sum_identity_norm": float(np.linalg.norm(pi_x + pi_p - I8)),
        "omega_skew_norm": float(np.linalg.norm(omega + omega.T)),
        "eta_signature": signature,
        "hamiltonian_relative_drift": float(
            (np.max(H_total_raw) - np.min(H_total_raw)) / max(1e-30, abs(np.mean(H_total_raw)))
        ),
        "depth_W_strictly_descending": bool(depth_W[0] > depth_W[1] > depth_W[2] > 0),
    }

    passed = (
        checks["pi_x_idempotent_norm"] < 1e-12
        and checks["pi_p_idempotent_norm"] < 1e-12
        and checks["pi_x_pi_p_orthogonal_norm"] < 1e-12
        and checks["pi_sum_identity_norm"] < 1e-12
        and checks["omega_skew_norm"] < 1e-12
        and signature == {"positive": 1, "negative": 7, "zero": 0}
        and checks["hamiltonian_relative_drift"] < 1e-12
        and checks["depth_W_strictly_descending"]
    )

    return {
        "status": "PASS_PROGRAMMATIC_X4_P4_PROJECTOR_ONE_TRAJECTORY" if passed else "FAIL",
        "checks": checks,
    }


def render_svg(flow: dict[str, Any]) -> str:
    """Render a text-free deterministic SVG."""
    W = flow["width"]
    H = flow["height"]
    x2 = np.array(flow["visual"]["x_projection_2d"], dtype=float)
    p2 = np.array(flow["visual"]["p_projection_2d"], dtype=float)
    v2 = np.array(flow["visual"]["v_projection_2d"], dtype=float)

    def map_points(points: np.ndarray, cx: float, cy: float, sx: float, sy: float) -> list[tuple[float, float]]:
        return [(float(cx + sx * x), float(cy + sy * y)) for x, y in points]

    Xpts = map_points(x2, W * 0.29, H * 0.50, W * 0.17, H * 0.25)
    Ppts = map_points(p2, W * 0.71, H * 0.50, W * 0.17, H * 0.25)
    Vpts = map_points(v2, W * 0.50, H * 0.50, W * 0.23, H * 0.32)

    def path_d(points: list[tuple[float, float]], step: int = 1) -> str:
        pts = points[::step]
        if not pts:
            return ""
        d = f"M {pts[0][0]:.3f} {pts[0][1]:.3f}"
        for x, y in pts[1:]:
            d += f" L {x:.3f} {y:.3f}"
        return d

    # Section 1 background lattice from exact coordinates, if present.
    ladder_nodes = []
    ladder_edges = []
    sec1_path = Path("/mnt/data/asha_section1_measurement_ladder_cl_1_7.coordinates.json")
    if sec1_path.exists():
        sec1 = json.loads(sec1_path.read_text())
        for n in sec1["nodes"]:
            # normalize from old 1800x720-ish coordinate system into central upper arch
            x = W * 0.16 + (float(n["x"]) / 1800.0) * W * 0.68
            y = H * 0.08 + (float(n["y"]) / 720.0) * H * 0.42
            ladder_nodes.append((x, y, int(n["grade"]), bool(n["has_time"])))
        # Use only a performance-safe subset of edges for visual clarity.
        node_lookup = {int(n["mask"]): i for i, n in enumerate(sec1["nodes"])}
        for e in sec1.get("edges", [])[::2]:
            a = node_lookup.get(int(e["source"]))
            b = node_lookup.get(int(e["target"]))
            if a is not None and b is not None:
                ladder_edges.append((a, b))

    def polyline_elements(points, color, opacity, width, dash=None):
        dash_attr = f' stroke-dasharray="{dash}"' if dash else ""
        return f'<path d="{path_d(points)}" fill="none" stroke="{color}" stroke-width="{width}" stroke-opacity="{opacity}" stroke-linecap="round" stroke-linejoin="round"{dash_attr}/>'

    # Exactly 13 outside unresolved nodes.
    outside = []
    for k in range(13):
        angle = -math.pi + (2 * math.pi * k / 13.0) + 0.10
        rx, ry = W * 0.47, H * 0.43
        outside.append((W * 0.5 + rx * math.cos(angle), H * 0.5 + ry * math.sin(angle)))

    # 16 matter sockets in a precise 4x4 lattice inside central law chamber.
    matter = []
    for row in range(4):
        for col in range(4):
            matter.append((W * 0.5 + (col - 1.5) * 35, H * 0.5 + (row - 1.5) * 35))

    # three islands: 1, 2x2, 3x3
    islands = {
        "point": [(W * 0.365, H * 0.50)],
        "two": [(W * 0.405 + dx * 22, H * 0.50 + dy * 22) for dy in [-0.5, 0.5] for dx in [-0.5, 0.5]],
        "three": [(W * 0.595 + dx * 22, H * 0.50 + dy * 22) for dy in [-1, 0, 1] for dx in [-1, 0, 1]],
    }

    svg = []
    svg.append(f'<svg xmlns="http://www.w3.org/2000/svg" width="{W}" height="{H}" viewBox="0 0 {W} {H}">')
    svg.append('<defs>')
    svg.append('<radialGradient id="bg" cx="50%" cy="50%" r="75%"><stop offset="0%" stop-color="#0d0e15"/><stop offset="60%" stop-color="#050508"/><stop offset="100%" stop-color="#010103"/></radialGradient>')
    svg.append('<radialGradient id="core" cx="50%" cy="50%" r="50%"><stop offset="0%" stop-color="#fff8d8"/><stop offset="35%" stop-color="#f4c66a"/><stop offset="100%" stop-color="#f4c66a" stop-opacity="0"/></radialGradient>')
    svg.append('<filter id="softGlow" x="-60%" y="-60%" width="220%" height="220%"><feGaussianBlur stdDeviation="5" result="b"/><feMerge><feMergeNode in="b"/><feMergeNode in="SourceGraphic"/></feMerge></filter>')
    svg.append('<filter id="largeGlow" x="-90%" y="-90%" width="280%" height="280%"><feGaussianBlur stdDeviation="14" result="b"/><feMerge><feMergeNode in="b"/><feMergeNode in="SourceGraphic"/></feMerge></filter>')
    svg.append('</defs>')
    svg.append('<rect width="100%" height="100%" fill="url(#bg)"/>')

    # Chaotic exterior faint nebula arcs/noise as deterministic curves.
    for k in range(22):
        a = 0.15 + k * 0.27
        x1 = W * (0.02 + 0.96 * ((math.sin(k * 1.7) + 1) / 2))
        y1 = H * (0.02 + 0.96 * ((math.cos(k * 1.3) + 1) / 2))
        x2c = W * 0.5 + W * 0.55 * math.cos(a)
        y2c = H * 0.5 + H * 0.50 * math.sin(a)
        color = "#4ddfe8" if k % 2 else "#d86a4f"
        svg.append(f'<path d="M {x1:.1f} {y1:.1f} C {W*0.3:.1f} {H*0.1 + (k%7)*70:.1f}, {W*0.7:.1f} {H*0.9-(k%5)*80:.1f}, {x2c:.1f} {y2c:.1f}" fill="none" stroke="{color}" stroke-opacity="0.055" stroke-width="2"/>')

    # Obsidian firewall ring.
    for i, (sw, op) in enumerate([(54, 0.16), (34, 0.22), (14, 0.40)]):
        svg.append(f'<ellipse cx="{W/2:.1f}" cy="{H/2:.1f}" rx="{W*0.445:.1f}" ry="{H*0.405:.1f}" fill="none" stroke="#000000" stroke-opacity="{op}" stroke-width="{sw}"/>')
    svg.append(f'<ellipse cx="{W/2:.1f}" cy="{H/2:.1f}" rx="{W*0.437:.1f}" ry="{H*0.397:.1f}" fill="none" stroke="#2a211d" stroke-opacity="0.46" stroke-width="5"/>')

    # Outside 13 unresolved nodes.
    for x, y in outside:
        svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="23" fill="#f5c36a" fill-opacity="0.035" filter="url(#largeGlow)"/>')
        svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="4.5" fill="#f6dca0" fill-opacity="0.70" filter="url(#softGlow)"/>')
        svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="33" fill="none" stroke="#f6dca0" stroke-opacity="0.08" stroke-width="1"/>')

    # Frosted projector lenses PB and PG.
    svg.append(f'<ellipse cx="{W*0.42:.1f}" cy="{H*0.50:.1f}" rx="{W*0.225:.1f}" ry="{H*0.255:.1f}" fill="#f0f0ff" fill-opacity="0.055" stroke="#f8e7c0" stroke-opacity="0.28" stroke-width="2" filter="url(#softGlow)"/>')
    svg.append(f'<ellipse cx="{W*0.58:.1f}" cy="{H*0.50:.1f}" rx="{W*0.225:.1f}" ry="{H*0.255:.1f}" fill="#70e8ff" fill-opacity="0.055" stroke="#82f4ff" stroke-opacity="0.28" stroke-width="2" filter="url(#softGlow)"/>')
    svg.append(f'<circle cx="{W*0.50:.1f}" cy="{H*0.50:.1f}" r="{H*0.115:.1f}" fill="url(#core)" filter="url(#largeGlow)"/>')

    # Measurement ladder background.
    if ladder_edges:
        for a, b in ladder_edges:
            x1, y1, _, _ = ladder_nodes[a]
            x2e, y2e, _, _ = ladder_nodes[b]
            svg.append(f'<line x1="{x1:.2f}" y1="{y1:.2f}" x2="{x2e:.2f}" y2="{y2e:.2f}" stroke="#d8eaff" stroke-opacity="0.055" stroke-width="0.8"/>')
    for x, y, grade, has_time in ladder_nodes:
        col = "#ff9b7e" if has_time else "#76eeff"
        r = 1.4 + 0.22 * abs(4 - grade)
        svg.append(f'<circle cx="{x:.2f}" cy="{y:.2f}" r="{r:.2f}" fill="{col}" fill-opacity="0.42" filter="url(#softGlow)"/>')

    # X and P projected trajectories.
    svg.append(polyline_elements(Xpts, "#ff8d6b", 0.31, 1.6))
    svg.append(polyline_elements(Ppts, "#74eaff", 0.31, 1.6))
    # Central unified V8 trajectory, multiple glow layers.
    svg.append(polyline_elements(Vpts, "#fff2c1", 0.16, 11))
    svg.append(polyline_elements(Vpts, "#f4c66a", 0.42, 3.2))
    svg.append(polyline_elements(Vpts, "#fdf8e4", 0.84, 1.15))

    # Projector beams between X/P sheets and core.
    for idx in np.linspace(0, len(Xpts) - 1, 24, dtype=int):
        x, y = Xpts[idx]
        svg.append(f'<line x1="{x:.2f}" y1="{y:.2f}" x2="{W*0.5:.2f}" y2="{H*0.5:.2f}" stroke="#ff9b7e" stroke-opacity="0.045" stroke-width="1"/>')
    for idx in np.linspace(0, len(Ppts) - 1, 24, dtype=int):
        x, y = Ppts[idx]
        svg.append(f'<line x1="{x:.2f}" y1="{y:.2f}" x2="{W*0.5:.2f}" y2="{H*0.5:.2f}" stroke="#74eaff" stroke-opacity="0.045" stroke-width="1"/>')

    # Three depth rings around the core.
    for j, rr in enumerate([56, 86, 118]):
        col = ["#fff8d8", "#ffd27a", "#83f5ff"][j]
        svg.append(f'<circle cx="{W*0.5:.1f}" cy="{H*0.5:.1f}" r="{rr}" fill="none" stroke="{col}" stroke-opacity="{0.35 - 0.06*j}" stroke-width="{1.8 - 0.3*j}" filter="url(#softGlow)"/>')

    # 16-state matter lattice.
    for i, (x1, y1) in enumerate(matter):
        for j, (x2m, y2m) in enumerate(matter):
            if j > i and (abs(x1 - x2m) == 35 and y1 == y2m or abs(y1 - y2m) == 35 and x1 == x2m):
                svg.append(f'<line x1="{x1:.1f}" y1="{y1:.1f}" x2="{x2m:.1f}" y2="{y2m:.1f}" stroke="#f2e6c1" stroke-opacity="0.13" stroke-width="1"/>')
    for idx, (x, y) in enumerate(matter):
        col = "#ff9b7e" if idx % 5 == 0 else "#83f5ff" if idx % 3 == 0 else "#f6dca0"
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="5.4" fill="{col}" fill-opacity="0.78" filter="url(#softGlow)"/>')

    # Three-island finite engine, no labels.
    all_island_points = []
    for key, pts in islands.items():
        all_island_points.extend(pts)
        for x, y in pts:
            svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="5.2" fill="#fff2c1" fill-opacity="0.72" filter="url(#softGlow)"/>')
    # island bridge pulses
    for x, y in islands["point"]:
        for x2m, y2m in islands["two"]:
            svg.append(f'<path d="M {x:.1f} {y:.1f} Q {(x+x2m)/2:.1f} {y-35:.1f}, {x2m:.1f} {y2m:.1f}" stroke="#f4c66a" stroke-opacity="0.24" stroke-width="1.2" fill="none"/>')
    for x, y in islands["two"]:
        for x2m, y2m in islands["three"]:
            if abs(y - y2m) < 24:
                svg.append(f'<path d="M {x:.1f} {y:.1f} Q {(x+x2m)/2:.1f} {y+32:.1f}, {x2m:.1f} {y2m:.1f}" stroke="#83f5ff" stroke-opacity="0.20" stroke-width="1.1" fill="none"/>')

    # Mexican-hat-like quartic basin rendered as computed mesh ellipses.
    cx, cy = W * 0.5, H * 0.735
    for k in range(20):
        r = 18 + k * 13
        # quartic profile y offset from V(r)=lambda(r^2-r0^2)^2, normalized
        rn = r / (18 + 19 * 13)
        yoff = 34 * ((rn * rn - 0.50) ** 2 - 0.10)
        svg.append(f'<ellipse cx="{cx:.1f}" cy="{cy + yoff:.1f}" rx="{r*2.0:.1f}" ry="{r*0.42:.1f}" fill="none" stroke="#fff0c8" stroke-opacity="{0.20*(1-k/24):.3f}" stroke-width="1"/>')
    for k in range(24):
        angle = 2 * math.pi * k / 24
        x2b = cx + math.cos(angle) * 445
        y2b = cy + math.sin(angle) * 96
        svg.append(f'<line x1="{cx:.1f}" y1="{cy:.1f}" x2="{x2b:.1f}" y2="{y2b:.1f}" stroke="#fff0c8" stroke-opacity="0.05" stroke-width="1"/>')

    # "One" current position marker at final sample; trajectory point.
    ox, oy = Vpts[-1]
    svg.append(f'<circle cx="{ox:.2f}" cy="{oy:.2f}" r="21" fill="#fff8d8" fill-opacity="0.12" filter="url(#largeGlow)"/>')
    svg.append(f'<circle cx="{ox:.2f}" cy="{oy:.2f}" r="6.5" fill="#fff8d8" fill-opacity="0.96" filter="url(#softGlow)"/>')

    # central vertical/light axis
    svg.append(f'<line x1="{W*0.5:.1f}" y1="{H*0.15:.1f}" x2="{W*0.5:.1f}" y2="{H*0.86:.1f}" stroke="#fff7d0" stroke-opacity="0.22" stroke-width="2" filter="url(#softGlow)"/>')
    for q in np.linspace(0.18, 0.84, 9):
        svg.append(f'<circle cx="{W*0.5:.1f}" cy="{H*q:.1f}" r="4.7" fill="#fff7d0" fill-opacity="0.66" filter="url(#softGlow)"/>')

    svg.append('</svg>')
    return "\n".join(svg)


def write_outputs(out_dir: Path, samples: int = 1800) -> dict[str, str]:
    out_dir.mkdir(parents=True, exist_ok=True)
    flow = compute_flow(samples=samples)
    svg = render_svg(flow)
    stem = FIGURE_ID
    svg_path = out_dir / f"{stem}.svg"
    json_path = out_dir / f"{stem}.geometry.json"
    manifest_path = out_dir / f"{stem}.manifest.json"
    svg_path.write_text(svg, encoding="utf-8")

    # Full geometry JSON with sampled trajectory.
    json_path.write_text(json.dumps(flow, indent=2), encoding="utf-8")

    manifest = {
        "figure": stem,
        "title": "Programmatic Projector-Universe Snapshot: X4/P4 Movement of One",
        "text_free_render": True,
        "mathematical_contract": {
            "V8": "X4 ⊕ P4",
            "eta_ASHA": "diag(+1,-1,-1,-1,-1,-1,-1,-1)",
            "Omega": "Σ dp_mu ∧ dx^mu",
            "Pi_X": "diag(1,1,1,1,0,0,0,0)",
            "Pi_P": "diag(0,0,0,0,1,1,1,1)",
            "flow": flow["contract"]["hamiltonian_kind"],
        },
        "validation": flow["validation"],
        "rendering_note": "The PNG/SVG image intentionally contains no text; formulas are preserved here and in geometry JSON.",
    }
    manifest_path.write_text(json.dumps(manifest, indent=2), encoding="utf-8")

    # Convert SVG to PNG using cairosvg if present; fall back to matplotlib not needed.
    png_path = out_dir / f"{stem}.png"
    try:
        import cairosvg  # type: ignore
        cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=WIDTH, output_height=HEIGHT)
    except Exception:
        # Use rsvg-convert or inkscape if available.
        import subprocess, shutil
        if shutil.which("rsvg-convert"):
            subprocess.run(["rsvg-convert", "-w", str(WIDTH), "-h", str(HEIGHT), str(svg_path), "-o", str(png_path)], check=True)
        elif shutil.which("inkscape"):
            subprocess.run(["inkscape", str(svg_path), "--export-filename", str(png_path)], check=True)
        else:
            raise RuntimeError("No SVG-to-PNG converter available")
    return {
        "svg": str(svg_path),
        "png": str(png_path),
        "geometry": str(json_path),
        "manifest": str(manifest_path),
    }
