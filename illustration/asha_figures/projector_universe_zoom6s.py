"""Zoomed ASHA projector-universe view with all six low-energy action sectors.

The rendered SVG/PNG is intentionally text-free.  The six visible sectors are
encoded in geometry JSON and manifest as:
S_grav, S_gauge, S_Higgs^ASHA, S_fermion, S_Yukawa^ASHA, S_nu^seesaw.

The central motion is inherited from the programmatic V8 = X4 ⊕ P4 Hamiltonian
flow of the previous figure, then zoomed/cropped around the law chamber so all
six S-sectors remain visible at once.
"""

from __future__ import annotations

import json
import math
from dataclasses import asdict, dataclass
from pathlib import Path
from typing import Any

import numpy as np

from .projector_universe import compute_flow
from .readme_next5 import L, S_SPLIT, M_PLANCK_REDUCED_GEV, scale_bridge_contract
from .readme_neutrinos import neutrino_contract

FIGURE_ID = "asha_projector_universe_zoom6s_all_visible"
WIDTH = 3000
HEIGHT = 1700
CENTER = (WIDTH / 2, HEIGHT / 2)


@dataclass(frozen=True)
class SectorSpec:
    key: str
    action_term: str
    angle_degrees: float
    center: tuple[float, float]
    radius: float
    formula: str
    visual_role: str


def _sector_centers() -> list[tuple[float, float, float]]:
    """Return six centers around a zoomed law chamber.

    Angles are arranged so the six sectors are all fully visible on a single
    view while leaving the X4/P4 trajectory readable at the center.
    """
    cx, cy = CENTER
    rx, ry = 860.0, 545.0
    # top, upper-right, lower-right, bottom, lower-left, upper-left
    angles = [-90.0, -28.0, 28.0, 90.0, 152.0, 208.0]
    out: list[tuple[float, float, float]] = []
    for deg in angles:
        a = math.radians(deg)
        out.append((cx + rx * math.cos(a), cy + ry * math.sin(a), deg))
    return out


def action_sector_specs() -> list[SectorSpec]:
    centers = _sector_centers()
    keys = [
        ("S_grav", "S_grav", "(M_P^2/2) ∫(R − 2Λ)√−g d^4x", "metric projection / spacetime stiffness"),
        ("S_gauge", "S_gauge", "−1/4 Σ_a F^a_{μν}F_a^{μν}", "inner-fluctuation gauge bridges"),
        ("S_Higgs_ASHA", "S_Higgs^ASHA", "λ_ASHA(r^2 − r0^2)^2", "quartic vacuum basin"),
        ("S_fermion", "S_fermion", "Σ ψ̄ iγ^μD_μ ψ", "16-state finite matter carrier"),
        ("S_Yukawa_ASHA", "S_Yukawa^ASHA", "ψ_L Y_f^ASHA Φ ψ_R + h.c.", "family/depth transport lanes"),
        ("S_nu_seesaw", "S_ν^seesaw", "M_ν = −(v^2/2)Y_ν^T M_R^{-1}Y_ν", "seesaw compression lane"),
    ]
    return [
        SectorSpec(
            key=k,
            action_term=t,
            angle_degrees=centers[i][2],
            center=(float(centers[i][0]), float(centers[i][1])),
            radius=168.0,
            formula=f,
            visual_role=r,
        )
        for i, (k, t, f, r) in enumerate(keys)
    ]


def compute_zoom6s(samples: int = 2400) -> dict[str, Any]:
    """Compute central X4/P4 motion plus six visible action-sector geometries."""
    base = compute_flow(samples=samples)
    sectors = action_sector_specs()

    v2 = np.array(base["visual"]["v_projection_2d"], dtype=float)
    x2 = np.array(base["visual"]["x_projection_2d"], dtype=float)
    p2 = np.array(base["visual"]["p_projection_2d"], dtype=float)

    cx, cy = CENTER

    def map_points(points: np.ndarray, sx: float, sy: float, ox: float = 0.0, oy: float = 0.0) -> list[list[float]]:
        return [[float(cx + ox + sx * x), float(cy + oy + sy * y)] for x, y in points]

    central = {
        "v8_one_trajectory": map_points(v2, 420.0, 260.0),
        "x4_projection": map_points(x2, 250.0, 150.0, ox=-260.0),
        "p4_projection": map_points(p2, 250.0, 150.0, ox=260.0),
    }

    lam_contract = scale_bridge_contract()
    neutrino = neutrino_contract()
    lam_asha = (3 / 8) * (1 + L) * (1 / 3 - S_SPLIT)

    sector_geometry: dict[str, Any] = {}
    for spec in sectors:
        sx, sy = spec.center
        r = spec.radius
        if spec.key == "S_grav":
            # Warped metric sheet in the X4 projection lane.
            grid_lines = []
            vals = np.linspace(-1.0, 1.0, 11)
            for u in vals:
                row = []
                col = []
                for v in np.linspace(-1.0, 1.0, 48):
                    warp = 0.18 * math.exp(-2.3 * (u*u + v*v))
                    row.append([sx + r * v, sy + 0.62 * r * (u + warp * math.cos(math.pi * v))])
                    col.append([sx + r * u, sy + 0.62 * r * (v + warp * math.sin(math.pi * u))])
                grid_lines.append({"kind": "row", "points": row})
                grid_lines.append({"kind": "column", "points": col})
            sector_geometry[spec.key] = {"grid_lines": grid_lines, "signature": [1, 3], "projected_from": "Π_X η_ASHA Π_X"}
        elif spec.key == "S_gauge":
            # Three finite islands with 12 traveling gauge-like pulses.
            point = [[sx - 95, sy]]
            two = [[sx - 28 + dx * 26, sy + dy * 26] for dy in [-0.5, 0.5] for dx in [-0.5, 0.5]]
            three = [[sx + 82 + dx * 23, sy + dy * 23] for dy in [-1, 0, 1] for dx in [-1, 0, 1]]
            pulses = []
            bridge_targets = two + three[:8]
            for i, target in enumerate(bridge_targets[:12]):
                t = (i + 0.5) / 12.0
                px = point[0][0] * (1 - t) + target[0] * t
                py = point[0][1] * (1 - t) + target[1] * t + 18 * math.sin(math.tau * t)
                pulses.append([px, py])
            sector_geometry[spec.key] = {"point_island": point, "two_by_two": two, "three_by_three": three, "pulse_count": 12, "pulses": pulses}
        elif spec.key == "S_Higgs_ASHA":
            # Exact quartic profile ring family V = λ(r²-r0²)², normalized.
            rings = []
            for i, rho in enumerate(np.linspace(0.05, 1.28, 26)):
                V = lam_asha * ((rho * rho - 0.66) ** 2)
                yoff = 95.0 * (V / (lam_asha * ((1.28 * 1.28 - 0.66) ** 2)) - 0.34)
                rings.append({"rho": float(rho), "V": float(V), "ellipse": [sx, sy + yoff, r * rho, 0.36 * r * rho]})
            sector_geometry[spec.key] = {"lambda_ASHA": float(lam_asha), "r0_squared_visual": 0.66, "rings": rings}
        elif spec.key == "S_fermion":
            # 16 finite matter states arranged in a fully visible lattice.
            nodes = []
            edges = []
            for row in range(4):
                for col in range(4):
                    nodes.append([sx + (col - 1.5) * 48, sy + (row - 1.5) * 48])
            for i in range(16):
                row, col = divmod(i, 4)
                if col < 3:
                    edges.append([i, i + 1])
                if row < 3:
                    edges.append([i, i + 4])
            sector_geometry[spec.key] = {"node_count": 16, "nodes": nodes, "edges": edges}
        elif spec.key == "S_Yukawa_ASHA":
            # Three family/depth transport lanes; raw W values stored, log-visible curves.
            N = np.array([1/3, 1/2, 2/3], dtype=float)
            Wd = np.exp(-4 * math.pi * N)
            lanes = []
            for j, (n, w) in enumerate(zip(N, Wd)):
                lane = []
                for t in np.linspace(0, 1, 96):
                    # Spiral/transport ray, visible compression based on log depth.
                    angle = 2.2 * math.pi * t + j * 0.55
                    rad = 18 + (95 + 28*j) * t
                    lane.append([sx + rad * math.cos(angle) * 0.92, sy + rad * math.sin(angle) * 0.58])
                lanes.append({"N": float(n), "W": float(w), "points": lane})
            sector_geometry[spec.key] = {"lane_count": 3, "lanes": lanes}
        elif spec.key == "S_nu_seesaw":
            # Heavy-to-light seesaw compression using README neutrino contract values.
            MR3 = float(neutrino.quantities["heavy_scale_bridge"]["M_R3_GeV"])
            MR2 = float(neutrino.quantities["heavy_scale_bridge"]["M_R2_GeV"])
            m3 = float(neutrino.quantities["rank2_normal_order_lane"]["m3_eV"])
            m2 = float(neutrino.quantities["rank2_normal_order_lane"]["m2_eV"])
            # Log normalize five nodes from heavy scale to light scale.
            vals = [MR3, MR2, math.sqrt(MR2 * m3), m3, m2]
            logs = np.log10(np.array(vals, dtype=float))
            lo, hi = float(np.min(logs)), float(np.max(logs))
            nodes = []
            for i, lg in enumerate(logs):
                y = sy - 110 + 220 * (1 - (lg - lo) / (hi - lo))
                x = sx + (i - 2) * 34
                nodes.append([x, y])
            sector_geometry[spec.key] = {"nodes": nodes, "M_R3_GeV": MR3, "M_R2_GeV": MR2, "m3_eV": m3, "m2_eV": m2}

    validation = validate_zoom6s(sectors, sector_geometry, base)
    return {
        "figure": FIGURE_ID,
        "width": WIDTH,
        "height": HEIGHT,
        "source": "programmatic_zoom_of_section15_with_six_low_energy_action_sectors",
        "central_contract": base["contract"],
        "central_validation": base["validation"],
        "source_constants": {
            "L": float(L),
            "S_split": float(S_SPLIT),
            "M_P_reduced_GeV": float(M_PLANCK_REDUCED_GEV),
            "v_GeV": float(lam_contract.quantities["v_GeV"]),
        },
        "central_geometry": central,
        "six_action_sectors": [asdict(s) for s in sectors],
        "sector_geometry": sector_geometry,
        "validation": validation,
    }


def validate_zoom6s(sectors: list[SectorSpec], geometry: dict[str, Any], base: dict[str, Any]) -> dict[str, Any]:
    expected = ["S_grav", "S_gauge", "S_Higgs_ASHA", "S_fermion", "S_Yukawa_ASHA", "S_nu_seesaw"]
    keys = [s.key for s in sectors]
    checks: dict[str, Any] = {
        "sector_count": len(sectors),
        "sector_keys": keys,
        "expected_keys_present": keys == expected,
        "central_flow_passed": base["validation"]["status"] == "PASS_PROGRAMMATIC_X4_P4_PROJECTOR_ONE_TRAJECTORY",
        "all_sector_centers_inside_view": all(0 <= s.center[0] <= WIDTH and 0 <= s.center[1] <= HEIGHT for s in sectors),
        "all_sector_discs_inside_view": all(s.radius < s.center[0] < WIDTH - s.radius and s.radius < s.center[1] < HEIGHT - s.radius for s in sectors),
        "gauge_pulse_count": geometry["S_gauge"]["pulse_count"],
        "fermion_node_count": geometry["S_fermion"]["node_count"],
        "yukawa_lane_count": geometry["S_Yukawa_ASHA"]["lane_count"],
        "higgs_lambda_positive": geometry["S_Higgs_ASHA"]["lambda_ASHA"] > 0,
        "neutrino_heavy_above_light": geometry["S_nu_seesaw"]["M_R3_GeV"] > geometry["S_nu_seesaw"]["m3_eV"],
    }
    passed = (
        checks["sector_count"] == 6
        and checks["expected_keys_present"]
        and checks["central_flow_passed"]
        and checks["all_sector_centers_inside_view"]
        and checks["all_sector_discs_inside_view"]
        and checks["gauge_pulse_count"] == 12
        and checks["fermion_node_count"] == 16
        and checks["yukawa_lane_count"] == 3
        and checks["higgs_lambda_positive"]
        and checks["neutrino_heavy_above_light"]
    )
    return {"status": "PASS_ZOOM6S_ALL_SIX_ACTION_SECTORS_VISIBLE" if passed else "FAIL", "checks": checks}


def _path(points: list[list[float]], step: int = 1) -> str:
    pts = points[::step]
    if not pts:
        return ""
    d = f"M {pts[0][0]:.3f} {pts[0][1]:.3f}"
    for x, y in pts[1:]:
        d += f" L {x:.3f} {y:.3f}"
    return d


def render_svg(data: dict[str, Any]) -> str:
    W, H = data["width"], data["height"]
    cx, cy = CENTER
    sectors = data["six_action_sectors"]
    geom = data["sector_geometry"]
    central = data["central_geometry"]

    svg: list[str] = []
    svg.append(f'<svg xmlns="http://www.w3.org/2000/svg" width="{W}" height="{H}" viewBox="0 0 {W} {H}">')
    svg.append('<defs>')
    svg.append('<radialGradient id="bg" cx="50%" cy="50%" r="72%"><stop offset="0%" stop-color="#101018"/><stop offset="55%" stop-color="#050508"/><stop offset="100%" stop-color="#010103"/></radialGradient>')
    svg.append('<radialGradient id="core" cx="50%" cy="50%" r="48%"><stop offset="0%" stop-color="#fffbe8"/><stop offset="32%" stop-color="#f7d277"/><stop offset="100%" stop-color="#f7d277" stop-opacity="0"/></radialGradient>')
    svg.append('<filter id="softGlow" x="-70%" y="-70%" width="240%" height="240%"><feGaussianBlur stdDeviation="4" result="b"/><feMerge><feMergeNode in="b"/><feMergeNode in="SourceGraphic"/></feMerge></filter>')
    svg.append('<filter id="largeGlow" x="-100%" y="-100%" width="300%" height="300%"><feGaussianBlur stdDeviation="15" result="b"/><feMerge><feMergeNode in="b"/><feMergeNode in="SourceGraphic"/></feMerge></filter>')
    svg.append('<linearGradient id="xp" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" stop-color="#ff8d6b"/><stop offset="50%" stop-color="#fff4cc"/><stop offset="100%" stop-color="#72eaff"/></linearGradient>')
    svg.append('</defs>')
    svg.append('<rect width="100%" height="100%" fill="url(#bg)"/>')

    # Zoomed firewall boundary, close enough to fill the view but leaving all six sectors visible.
    for sw, op in [(64, 0.17), (36, 0.25), (12, 0.48)]:
        svg.append(f'<ellipse cx="{cx:.1f}" cy="{cy:.1f}" rx="{W*0.405:.1f}" ry="{H*0.405:.1f}" fill="none" stroke="#000000" stroke-opacity="{op}" stroke-width="{sw}"/>')
    svg.append(f'<ellipse cx="{cx:.1f}" cy="{cy:.1f}" rx="{W*0.392:.1f}" ry="{H*0.392:.1f}" fill="none" stroke="#2a211d" stroke-opacity="0.48" stroke-width="5"/>')

    # Six radial projector lanes.
    for s in sectors:
        sx, sy = s["center"]
        key = s["key"]
        col = {
            "S_grav": "#d9e8ff",
            "S_gauge": "#83f5ff",
            "S_Higgs_ASHA": "#f7d277",
            "S_fermion": "#fff2c1",
            "S_Yukawa_ASHA": "#ff9b7e",
            "S_nu_seesaw": "#b9b7ff",
        }[key]
        svg.append(f'<line x1="{cx:.1f}" y1="{cy:.1f}" x2="{sx:.1f}" y2="{sy:.1f}" stroke="{col}" stroke-opacity="0.12" stroke-width="2"/>')
        svg.append(f'<circle cx="{sx:.1f}" cy="{sy:.1f}" r="{s["radius"]:.1f}" fill="#ffffff" fill-opacity="0.018" stroke="{col}" stroke-opacity="0.23" stroke-width="1.8" filter="url(#softGlow)"/>')
        svg.append(f'<circle cx="{sx:.1f}" cy="{sy:.1f}" r="{s["radius"]*0.62:.1f}" fill="none" stroke="{col}" stroke-opacity="0.08" stroke-width="1"/>')

    # Central frosted lenses and core.
    svg.append(f'<ellipse cx="{cx-92:.1f}" cy="{cy:.1f}" rx="360" ry="230" fill="#fff9ea" fill-opacity="0.040" stroke="#ffdba0" stroke-opacity="0.20" stroke-width="2" filter="url(#softGlow)"/>')
    svg.append(f'<ellipse cx="{cx+92:.1f}" cy="{cy:.1f}" rx="360" ry="230" fill="#74eaff" fill-opacity="0.040" stroke="#74eaff" stroke-opacity="0.20" stroke-width="2" filter="url(#softGlow)"/>')
    svg.append(f'<circle cx="{cx:.1f}" cy="{cy:.1f}" r="124" fill="url(#core)" filter="url(#largeGlow)"/>')

    # Central X4/P4 motion and One trajectory.
    svg.append(f'<path d="{_path(central["x4_projection"], 2)}" fill="none" stroke="#ff8d6b" stroke-opacity="0.30" stroke-width="2" stroke-linejoin="round" stroke-linecap="round"/>')
    svg.append(f'<path d="{_path(central["p4_projection"], 2)}" fill="none" stroke="#74eaff" stroke-opacity="0.30" stroke-width="2" stroke-linejoin="round" stroke-linecap="round"/>')
    svg.append(f'<path d="{_path(central["v8_one_trajectory"], 1)}" fill="none" stroke="#f7d277" stroke-opacity="0.24" stroke-width="13" stroke-linejoin="round" stroke-linecap="round" filter="url(#largeGlow)"/>')
    svg.append(f'<path d="{_path(central["v8_one_trajectory"], 1)}" fill="none" stroke="url(#xp)" stroke-opacity="0.92" stroke-width="2.0" stroke-linejoin="round" stroke-linecap="round"/>')
    ox, oy = central["v8_one_trajectory"][-1]
    svg.append(f'<circle cx="{ox:.2f}" cy="{oy:.2f}" r="22" fill="#fffbe8" fill-opacity="0.11" filter="url(#largeGlow)"/>')
    svg.append(f'<circle cx="{ox:.2f}" cy="{oy:.2f}" r="6.5" fill="#fffbe8" fill-opacity="0.95" filter="url(#softGlow)"/>')

    # Render each sector.
    # S_grav metric sheet
    for line in geom["S_grav"]["grid_lines"]:
        svg.append(f'<path d="{_path(line["points"])}" fill="none" stroke="#d9e8ff" stroke-opacity="0.135" stroke-width="0.9"/>')
    # S_gauge islands + pulses
    g = geom["S_gauge"]
    for group, color in [("point_island", "#fffbe8"), ("two_by_two", "#f7d277"), ("three_by_three", "#83f5ff")]:
        for x, y in g[group]:
            svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="6.4" fill="{color}" fill-opacity="0.80" filter="url(#softGlow)"/>')
    src = g["point_island"][0]
    for target in g["two_by_two"] + g["three_by_three"]:
        svg.append(f'<path d="M {src[0]:.1f} {src[1]:.1f} Q {(src[0]+target[0])/2:.1f} {src[1]-28:.1f}, {target[0]:.1f} {target[1]:.1f}" stroke="#83f5ff" stroke-opacity="0.13" stroke-width="1" fill="none"/>')
    for x, y in g["pulses"]:
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="4.0" fill="#fffbe8" fill-opacity="0.92" filter="url(#softGlow)"/>')
    # S_Higgs quartic rings
    for ring in geom["S_Higgs_ASHA"]["rings"]:
        x, y, rx, ry = ring["ellipse"]
        svg.append(f'<ellipse cx="{x:.1f}" cy="{y:.1f}" rx="{rx:.1f}" ry="{ry:.1f}" fill="none" stroke="#f7d277" stroke-opacity="0.18" stroke-width="0.9"/>')
    # S_fermion 16-node lattice
    f = geom["S_fermion"]
    for a, b in f["edges"]:
        x1, y1 = f["nodes"][a]
        x2, y2 = f["nodes"][b]
        svg.append(f'<line x1="{x1:.1f}" y1="{y1:.1f}" x2="{x2:.1f}" y2="{y2:.1f}" stroke="#fff2c1" stroke-opacity="0.18" stroke-width="1"/>')
    for i, (x, y) in enumerate(f["nodes"]):
        col = "#ff9b7e" if i % 5 == 0 else "#83f5ff" if i % 3 == 0 else "#fff2c1"
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="7.1" fill="{col}" fill-opacity="0.78" filter="url(#softGlow)"/>')
    # S_Yukawa lanes
    for lane in geom["S_Yukawa_ASHA"]["lanes"]:
        svg.append(f'<path d="{_path(lane["points"])}" fill="none" stroke="#ff9b7e" stroke-opacity="0.42" stroke-width="2.1" stroke-linecap="round" stroke-linejoin="round"/>')
        x, y = lane["points"][-1]
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="5.2" fill="#ffcfaa" fill-opacity="0.82" filter="url(#softGlow)"/>')
    # S_nu seesaw compression
    n = geom["S_nu_seesaw"]
    pts = n["nodes"]
    svg.append(f'<path d="{_path(pts)}" fill="none" stroke="#b9b7ff" stroke-opacity="0.32" stroke-width="2.2"/>')
    for i, (x, y) in enumerate(pts):
        rr = 8.0 - min(i, 4) * 0.85
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="{rr:.1f}" fill="#d6d3ff" fill-opacity="0.82" filter="url(#softGlow)"/>')
    # subtle compression cones in neutrino sector
    for i in range(len(pts)-1):
        x1, y1 = pts[i]
        x2, y2 = pts[i+1]
        svg.append(f'<path d="M {x1:.1f} {y1:.1f} C {(x1+x2)/2:.1f} {y1+30:.1f}, {(x1+x2)/2:.1f} {y2-30:.1f}, {x2:.1f} {y2:.1f}" stroke="#d6d3ff" stroke-opacity="0.13" stroke-width="8" fill="none"/>')

    # Six small anchor lights around central core, one per action sector.
    for s in sectors:
        sx, sy = s["center"]
        a = math.atan2(sy - cy, sx - cx)
        ax = cx + math.cos(a) * 178
        ay = cy + math.sin(a) * 178
        svg.append(f'<circle cx="{ax:.1f}" cy="{ay:.1f}" r="5.8" fill="#fffbe8" fill-opacity="0.80" filter="url(#softGlow)"/>')

    # Exactly 13 outer unresolved nodes, still visible but pushed outward by the zoom.
    for k in range(13):
        angle = -math.pi + math.tau * k / 13.0 + 0.10
        x = cx + W * 0.472 * math.cos(angle)
        y = cy + H * 0.458 * math.sin(angle)
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="4.4" fill="#f7d277" fill-opacity="0.50" filter="url(#softGlow)"/>')
        svg.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="22" fill="none" stroke="#f7d277" stroke-opacity="0.055" stroke-width="1"/>')

    svg.append('</svg>')
    return "\n".join(svg)


def write_outputs(out_dir: Path, samples: int = 2400) -> dict[str, str]:
    out_dir.mkdir(parents=True, exist_ok=True)
    data = compute_zoom6s(samples=samples)
    svg = render_svg(data)
    stem = FIGURE_ID
    svg_path = out_dir / f"{stem}.svg"
    png_path = out_dir / f"{stem}.png"
    geometry_path = out_dir / f"{stem}.geometry.json"
    manifest_path = out_dir / f"{stem}.manifest.json"
    note_path = out_dir / f"{stem}_implementation_note.md"

    svg_path.write_text(svg, encoding="utf-8")
    geometry_path.write_text(json.dumps(data, indent=2), encoding="utf-8")
    manifest = {
        "figure": stem,
        "title": "Zoomed Programmatic Projector-Universe: All Six ASHA S-Terms Visible",
        "text_free_render": "<text" not in svg.lower(),
        "six_S_terms": [s["action_term"] for s in data["six_action_sectors"]],
        "mathematical_contract": {
            "central_phase_space": "V8 = X4 ⊕ P4",
            "central_flow": "canonical Hamiltonian oscillator flow on four X/P pairs",
            "projectors": ["Π_X", "Π_P"],
            "six_sector_view": "S_grav + S_gauge + S_Higgs^ASHA + S_fermion + S_Yukawa^ASHA + S_nu^seesaw",
        },
        "validation": data["validation"],
        "central_validation": data["central_validation"],
        "rendering_note": "The image itself is text-free. The six S-term identities and formulas are stored here and in geometry JSON.",
    }
    manifest_path.write_text(json.dumps(manifest, indent=2), encoding="utf-8")
    note_path.write_text(
        "# Zoom6S Programmatic ASHA Projector-Universe Figure\n\n"
        "This figure is a deterministic zoom of the X4/P4 projector-universe snapshot. "
        "It keeps the rendered artwork text-free while making all six low-energy ASHA action sectors visible at once.\n\n"
        "Six visible sectors:\n"
        "1. S_grav — metric projection / spacetime stiffness.\n"
        "2. S_gauge — inner-fluctuation gauge bridges.\n"
        "3. S_Higgs^ASHA — quartic vacuum basin.\n"
        "4. S_fermion — 16-state finite matter carrier.\n"
        "5. S_Yukawa^ASHA — family/depth transport lanes.\n"
        "6. S_nu^seesaw — heavy-to-light seesaw compression.\n\n"
        "The central orbit is computed from the same V8 = X4 ⊕ P4 Hamiltonian flow used in Section 15.\n",
        encoding="utf-8",
    )

    try:
        import cairosvg  # type: ignore
        cairosvg.svg2png(url=str(svg_path), write_to=str(png_path), output_width=WIDTH, output_height=HEIGHT)
    except Exception:
        import shutil
        import subprocess
        if shutil.which("rsvg-convert"):
            subprocess.run(["rsvg-convert", "-w", str(WIDTH), "-h", str(HEIGHT), str(svg_path), "-o", str(png_path)], check=True)
        elif shutil.which("inkscape"):
            subprocess.run(["inkscape", str(svg_path), "--export-filename", str(png_path)], check=True)
        else:
            raise RuntimeError("No SVG-to-PNG converter available")

    return {
        "svg": str(svg_path),
        "png": str(png_path),
        "geometry": str(geometry_path),
        "manifest": str(manifest_path),
        "note": str(note_path),
    }
